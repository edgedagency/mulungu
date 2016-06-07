<?php namespace Edged\Mulungu\Idatha\Document;

use Edged\Mulungu\Idatha\Connection\Driver\Interfaces\ConnectionDriver;
use Edged\Mulungu\Idatha\Document\Exception\DocumentException;
use Edged\Mulungu\Idatha\Document\Interfaces\DocumentInterface;
use Carbon\Carbon;
use Edged\Mulungu\Idatha\Event\DocumentDataEvent;
use Edged\Mulungu\Idatha\Event\DocumentEvent;
use Edged\Mulungu\Idatha\Event\DocumentEvents;
use Psr\Log\LoggerInterface;
use Symfony\Component\EventDispatcher\EventDispatcherInterface;

abstract class Document extends \stdClass implements DocumentInterface{

    private $connection;
    protected $collection;
	protected $logger;
    protected $eventDispatcher;
    protected $internalFieldMappings = array("_id"=>"id","_rev"=>"rev","_key"=>"key" );

    private $isCollection = false;
    private $error = false;
    private $totalDocuments = null;
    private $lastKnownResults;

    public function __construct( ConnectionDriver $connection, LoggerInterface $logger, EventDispatcherInterface $eventDispatcher ){
    	$this->connection = $connection;
    	$this->logger = $logger;
        $this->eventDispatcher = $eventDispatcher;
    }
    /**
     * magic method for return model properties
     *
     * @param string $name property name
     * @return null|string property value
     */
    public function __get($name){
        return ( isset( $this->$name ) )?$this->$name:null;
    }
    /**
     * sets class parameter
     *
     * @param string $name parameter name
     * @param string $value parameter value
     */
    public function __set( $name, $value ){
        if( !is_string( $name ) )
            throw new DocumentException("document can't have non string based key. try document->key = mixed");

        $this->$name = $value;
    }

    public function getType(){
       $type = get_class($this);
       $type = mb_strtolower( substr( $type, strrpos( $type, "\\", 1 ) + 1, strlen( $type ) ) );

       return $type;
    }

    public function setConnection( ConnectionDriver $connection ){
        $this->connection = $connection;
        return $this;
    }

    public function getConnection(){
        return $this->connection;
    }

    public function save(){
		//timestamping
		$this->createdDate = Carbon::now()->toIso8601String();
		$this->modifiedDate = Carbon::now()->toIso8601String();

        $this->eventDispatcher->dispatch( DocumentEvents::BEFORE_SAVE , new DocumentEvent( $this ) );


        $results = $this->connection->save( $this );
        $this->setLastKnownResults( $results );


        if( $this->hasError( $results ) ) {
            $this->hydrateFromArray( $results );
            $this->logger->error( "save operation failed", $results );
            return false;
        }

		$this->hydrateFromArray( $results );


        $this->eventDispatcher->dispatch( DocumentEvents::AFTER_SAVE, new DocumentEvent( $this ));

        return true;
    }

    public function update(){

    	$this->modifiedDate = Carbon::now()->toIso8601String();
        $this->eventDispatcher->dispatch(DocumentEvents::BEFORE_UPDATE, new DocumentEvent($this));
        $results = $this->connection->update( $this );

        $this->setLastKnownResults( $results );

        if( $this->hasError( $results ) ) {

            $this->hydrateFromArray( $results );
            $this->logger->error( "update operation failed", $results );

            return null;
        }

        $this->hydrateFromArray( $results );


        $this->eventDispatcher->dispatch(DocumentEvents::AFTER_UPDATE, new DocumentEvent($this));


        return true;
    }

    /**
     * @param $id
     * @param array $data
     * @return bool
     */
    public function patch( $id, array $data  ){
        $data = $this->clean( $data );
        
        $data[ "modifiedDate" ] = Carbon::now()->toIso8601String();

        $this->eventDispatcher->dispatch(DocumentEvents::BEFORE_PATCH, new DocumentDataEvent($data));
        $results = $this->connection->patch( $this->getInternalId( $id ), $data );

        $this->setLastKnownResults( $results );

        if( $this->hasError( $results ) ) {

            $this->hydrateFromArray( $results );
            $this->logger->error( "patch operation failed", $results );

            return null;
        }


        $this->eventDispatcher->dispatch(DocumentEvents::AFTER_PATCH, new DocumentDataEvent($data));


        //merge data set, typical only _id,_rev,_key are returned from save operation
        $mergedResults = array_merge( $data, $results );
        //remove error key, that is not needed on object
        unset( $mergedResults["error"] );

        $this->hydrateFromArray( $mergedResults );

        return true;
    }
    /**
     * @param $id
     * @param bool $count
     * @param null $fields
     * @return $this
     */
    public function find( $id, $count=false, $fields=null ){

        $results = $this->connection->find($this->getCollection(),
            $id, $count, $fields );

        $this->setLastKnownResults( $results );

        if( $this->hasError( $results ) ) {

            $this->hydrateFromArray( $results );
            $this->logger->error( "find all operation failed", $results );

            return null;
        }

        if( array_key_exists( "result", $results) ) {
            if( !empty( $results["result"] ) ) {
                $this->hydrateFromArray($results["result"][0]);
            }
        }

        return $this;
    }

    /**
     * @param bool $count
     * @param null $batchSize
     * @param null $fields
     * @param array $bindVars
     * @return DocumentCollection|null
     */
    public function findAll( $count=false, $batchSize=null, $fields=null, $bindVars=[], $limit=15, $page=0, $sort=[ "fields"=>["createdDate"], "direction"=>ConnectionDriver::SORT_DESC ] ){
        $results = $this->connection->findAll( $this->getCollection(), $count, $batchSize, $fields, $bindVars, $limit, $page, $sort );
        $this->setLastKnownResults( $results );

        if( $this->hasError( $results ) ) {
            $this->hydrateFromArray( $results );
            $this->logger->error( "find all operation failed", $results );

            return null;
        }

        return $this->toCollection( $results, get_called_class() );
    }

    public function hasError( $results = null ){
        if( is_array( $results ) ) {
            if (array_key_exists("error", $results)) {
                return $results[ "error" ];
            }
        }

        return $this->error;
    }

    public function isCollection(  ){
        return $this->isCollection;
    }

    public function execute( $statement, $hydrate=true ){
    	$results = $this->connection->execute( $statement );
        $this->setLastKnownResults( $results );

        if( $this->hasError( $results ) ) {
            $this->hydrateFromArray( $results );
            $this->logger->error( "execute operation failed", $results );
            return null;
        }

    	return ($hydrate===true)? $this->toCollection( $results, get_called_class() ) : $results;
    }

    public function delete( $id = null ){
        $id = (isset( $id ) )?$id:$this->id;

        $this->eventDispatcher->dispatch(DocumentEvents::BEFORE_DELETE, new DocumentEvent($this));

        $results = $this->connection->delete( $this->getInternalId( $id ) );

        if( $this->hasError( $results ) ) {
            $this->hydrateFromArray( $results );
            $this->logger->error( "delete operation failed", $results );

            return null;
        }

        $this->hydrateFromArray( $results );
        $this->eventDispatcher->dispatch(DocumentEvents::AFTER_DELETE, new DocumentEvent($this));

        return $this;
    }


    public function setCollection( $collection ){
        $this->collection = $collection;

        return $this;
    }

    public function getCollection(){
        if( !isset( $this->collection ) || $this->collection ==null || $this->collection == "" )
            $this->collection = $this->guessCollection();

        return $this->collection;
    }

    public function hydrateFromArray( $data, $mapInternalAttributes=false ){

        $this->logger->debug( " -- hydrating -- " , $data );

        if( is_array( $data ) ){

            if( $mapInternalAttributes ) {
                $this->internalFieldMappings = array_flip($this->internalFieldMappings);
            }

            foreach( $data as $key=>$value ){

                    $key = ((array_key_exists($key, $this->internalFieldMappings)) and $mapInternalAttributes) ? $this->internalFieldMappings[$key] : $key;
                    $value = ( $key == "_id" )? $this->getId( $value ) : $value;

                    if (isset($this->logger)) {
                        $this->logger->debug("adding key " . $key);
                    }

                    $this->$key = $value;
            }

            return $this;

        }else{

            $this->logger->debug( "data used for hydration is not an array, hydration ignored" );

        }
    }

    private function clean( $data ){
        //removes elements that shouldn't be saved
        unset( $data["_key"], $data["_id"], $data["_rev"] );
        return $data;
    }


    private function getId( $value ){

        if( is_null( $value ) or is_array( $value ) )
            return $value;

        return ( !strpos( $value, "/" ))?$this->getCollection() . "/" . $value:$value;
    }

    public function toArray( $root=null, $mapInternalAttributes=false ){

        $this->logger->debug( " -- toArray -- " );


        $document = [];
        $reflection = new \ReflectionObject($this);
        $properties = $reflection->getProperties( \ReflectionProperty::IS_PUBLIC );

        foreach( $properties as $property ) {
            $key = ( array_key_exists( $property->getName(), $this->internalFieldMappings ) && $mapInternalAttributes )? $this->internalFieldMappings[ $property->getName() ] : $property->getName();
            $value = ( ($key=="id") xor ($key=="_id") )?$this->splitInternalId() : $property->getValue( $this );

            if( isset( $this->logger ) ){
                $this->logger->debug( "adding key " . $key );
            }

            $document[ $key ] = $value;
        }
        return ( isset( $root ) )? array( $root=>$document ) : $document;
    }

    public function reset(){
        $this->isCollection = false;
        $this->error = false;
    }

    public function getTotalDocuments(){
        return $this->totalDocuments;
    }

    private function splitInternalId(){
        $documentId = null;
        if (isset($this->_id) and strpos($this->_id, "/") != false){
            list($this->collection, $documentId) = explode('/', $this->_id, 2);
        }else{
            $documentId = $this->id;
        }

        return $documentId;
    }

    private function guessCollection(){
        return sprintf( "%s%s",$this->getType(), "s" );
    }

    private function getInternalId( $id ){
        return ( strpos( $id, '/') )? $id : $this->getCollection()."/".$id;
    }
    
    private function toCollection( $results, $class ){
    	$collection = new DocumentCollection();

        $this->logger->debug( "toCollection ",  $results );

    	if( is_array( $results ) ){
            if( array_key_exists( "result", $results) ) {

                $resultsData =  $results['result'];

                if( empty( $resultsData ) || count( $resultsData ) == 0  ) {
                    return null;
                }

                foreach ( $resultsData as $result) {

                    if (!empty($result)) {

                        if( isset( $this->logger ) ){
                            $this->logger->debug( "result ",  $result );
                        }

                        $class = new $class($this->connection, $this->logger, $this->eventDispatcher);
                        $class->hydrateFromArray( $result );
                        $collection->add($class);
                    }
                }

                $this->isCollection = true;

                return $collection;
            }
    	}

    	return null;
    }

    public function getLastKnownResults(){
        return $this->lastKnownResults;
    }
    public function setLastKnownResults($lastKnownResults){
        $this->lastKnownResults = $lastKnownResults;
    }
}