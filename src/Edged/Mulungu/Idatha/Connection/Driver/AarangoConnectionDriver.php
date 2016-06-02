<?php namespace Edged\Mulungu\Idatha\Connection\Driver;

use Edged\Mulungu\Idatha\Connection\Connection;
use Edged\Mulungu\Idatha\Connection\Driver\Dialect\AarangoDialect;
use Edged\Mulungu\Idatha\Connection\Driver\Interfaces\ConnectionDriver;
use Edged\Mulungu\Idatha\Document\Document;
use GuzzleHttp\Exception\ConnectException;
use GuzzleHttp\Exception\RequestException;
use Psr\Log\LoggerInterface;
use Symfony\Component\HttpFoundation\Request;

class AarangoConnectionDriver implements ConnectionDriver{

    private $connection;
    private $databaseName;
    private $response;
    private $request;
    private $logger;

    public function __construct( Connection $connection, $databaseName = null, LoggerInterface $logger = null ){
        $this->databaseName = $databaseName;
        $this->connection = $connection;
        $this->logger = $logger;
    }

    public function setConnection(Connection $connection){
        $this->connection = $connection;
    }

    public function getConnection(){
        return $this->connection;
    }

    public function save( Document $document ){

        if( !$this->collectionExists( $document->getCollection()  ) ){
            $this->saveCollection( $document->getCollection() );
        }

        try {

            $requestURL = sprintf(AarangoDialect::CREATE_DOCUMENT, $this->databaseName, $document->getCollection());

            $this->request  = new \GuzzleHttp\Psr7\Request( Request::METHOD_POST, $requestURL, [], json_encode( $document ) );
            $this->response = $this->connection->send($this->request);
            $contents = $this->response->getBody()->getContents();

            if( $this->logger ){
                $this->logger->debug( "request sent to " . $requestURL );
                $this->logger->debug( "response " . json_encode( $contents ) );
            }

            return json_decode( $contents, true );

        }catch ( RequestException $requestException  ){

            if( $this->logger ){
                $this->logger->error( "Failed to complete insert exception trace [ " . $requestException->getTraceAsString() . " ]" );
            }

            return [ "error"=>true,
                "exception"=>$requestException->getMessage() ,
                "code"=> $requestException->getCode(),
                "stackTrace"=>explode( "#", $requestException->getTraceAsString() ),
                "errorType"=>"Request Error"
            ];

        }catch( ConnectException $connectionException ){

            return [ "error"=>true,
                "exception"=>$connectionException->getMessage(),
                "stackTrace"=>explode( "#", $connectionException->getTraceAsString() ),
                "errorType"=>"Connection Error"
            ];
        }


    }

    public function delete( $id ){
        try{

            $requestURL  = sprintf( AarangoDialect::DELETE_DOCUMENT_BY_ID, $this->databaseName, $id );
            $this->request  = new \GuzzleHttp\Psr7\Request( Request::METHOD_DELETE, $requestURL );
            $this->response = $this->connection->send( $this->request );
            $contents = $this->response->getBody()->getContents();

            return json_decode( $contents, true);

        }catch ( RequestException $requestException  ){

            return [ "error"=>true,
                "exception"=>$requestException->getMessage() ,
                "code"=> $requestException->getCode(),
                "stackTrace"=>explode( "#", $requestException->getTraceAsString() ),
                "errorType"=>"Request Error"
            ];

        }catch( ConnectException $connectionException ){

            return [ "error"=>true,
                "exception"=>$connectionException->getMessage() ,
                "stackTrace"=>$connectionException->getTraceAsString(),
                "errorType"=>"Connection Error"
            ];
        }
    }

    public function update( Document $document ){

   	    try{
	        $requestURL  = sprintf( AarangoDialect::UPDATE_DOCUMENT_BY_ID, $this->databaseName, $document->_id );
            $this->request  = new \GuzzleHttp\Psr7\Request( Request::METHOD_PUT, $requestURL, [], json_encode( $document ) );
	        $this->response = $this->connection->send( $this->request );
            $contents = $this->response->getBody()->getContents();

            return json_decode( $contents, true);

         }catch ( RequestException $requestException  ){

            return [ "error"=>true,
                "exception"=>$requestException->getMessage() ,
                "code"=> $requestException->getCode(),
                "stackTrace"=>explode( "#", $requestException->getTraceAsString() ),
                "errorType"=>"Request Error"
            ];

         }catch( ConnectException $connectionException ){

            return [ "error"=>true,
                "exception"=>$connectionException->getMessage() ,
                "stackTrace"=>$connectionException->getTraceAsString(),
                "errorType"=>"Connection Error"
            ];
        }
    }

    public function patch( $id, $data=[] ){
        try {
            $requestURL = sprintf(AarangoDialect::UPDATE_DOCUMENT_BY_ID, $this->databaseName, $id);

            if( $this->logger ){
                $this->logger->debug(  "sending request " . $requestURL );
            }

            $this->request  = new \GuzzleHttp\Psr7\Request( Request::METHOD_PATCH, $requestURL, [], json_encode( $data ) );
            $this->response = $this->connection->send($this->request);
            $contents = $this->response->getBody()->getContents();

            return json_decode( $contents, true);

        } catch (RequestException $requestException) {

            return ["error" => true,
                "exception" => $requestException->getMessage(),
                "code" => $requestException->getCode(),
                "stackTrace"=>explode( "#", $requestException->getTraceAsString() ),
                "errorType"=>"Request Error"
            ];

        }catch( ConnectException $connectionException ){

            return [ "error"=>true,
                "exception"=>$connectionException->getMessage() ,
                "stackTrace"=>$connectionException->getTraceAsString(),
                "errorType"=>"Connection Error"
            ];
        }
    }


    public function execute( $statement ){
        try{

            $requestURL  = sprintf( AarangoDialect::QUERY_AQL, $this->databaseName );
            $this->request = new \GuzzleHttp\Psr7\Request( Request::METHOD_POST, $requestURL, [], json_encode( $statement ) );
            $this->response = $this->connection->send( $this->request );
            $contents = $this->response->getBody()->getContents();

            if( $this->logger ){
                $this->logger->debug( "requestURL " . $requestURL );
                $this->logger->debug( "statement  " , $statement );
                $this->logger->debug( "response " . json_encode( $contents ) );
            }

            return json_decode( $contents, true);

        }catch( RequestException $requestException ){


            if( $this->logger ) {
                $this->logger->error( "request uri :" . $requestException->getRequest()->getUri() );
                $this->logger->error( "request body :" . $requestException->getRequest()->getBody()->getContents() );
                $this->logger->error( "message :" .  $requestException->getMessage() );
                $this->logger->error( "stackTrace :" .  $requestException->getTraceAsString() );
            }

            return [ "error"=>true,
                "exception"=>$requestException->getMessage() ,
                "code"=> $requestException->getCode(),
                "stackTrace"=>explode( "#", $requestException->getTraceAsString() ),
                "errorType"=>"Request Error"
            ];

        }catch( ConnectException $connectionException ){
            return [ "error"=>true,
                "exception"=>$connectionException->getMessage() ,
                "stackTrace"=>$connectionException->getTraceAsString(),
                "errorType"=>"Connection Error"
            ];
        }
    }

    public function find( $collection, $id, $count=false, $fields=null ){

        $query["query"] ="FOR item IN $collection ";
        $query["query"].="FILTER item._key == '$id' ";
        $query["query"].="RETURN ". (( is_array( $fields ) )? $this->getQueryParts($fields,"item") : "item");
        $query["count"] = $count;

        $result = $this->execute( $query );

        return $result;
    }

    public function findAll( $collection, $count=false, $batchSize=null, $fields=null, $bindVars=[], $limit=15, $page=0, $sort=[ "fields"=>["createdDate"], "direction"=>ConnectionDriver::SORT_ASC ] ){

        $query["query"] ="FOR item IN $collection ";

        $query["query"] .= " SORT ". $this->getFieldsPrefixed( $sort["fields"], "item" )." ".$sort["direction"];

        if( $page > 0 ) {
            $query["query"] .= sprintf(" LIMIT %s, %s ", ($limit*($page-1)), $limit );
        }



        $query["query"].=" RETURN ". (( is_array( $fields ) )? $this->getQueryParts($fields,"item") : "item");
        $query["count"] = $count;

        if( $count === true and $page > 0 ){
            $query["options"] = [ "fullCount"=>$count ];
        }

        if( isset( $batchSize ) ) {
            $query["batchSize"] = $batchSize;
        }

        $result = $this->execute( $query );

        return $result;
    }

    public function saveCollection( $collection ){
        $requestURL  = sprintf( AarangoDialect::CREATE_COLLECTION, $this->databaseName );

        $this->request  = new \GuzzleHttp\Psr7\Request( Request::METHOD_POST, $requestURL, [], json_encode( ["name"=>$collection]));
        $this->response = $this->connection->send( $this->request );
        $contents = $this->response->getBody()->getContents();

        return json_decode( $contents, true);
    }

    public function collectionExists( $collection ){
        try{
            $requestURL = sprintf( AarangoDialect::FIND_COLLECTION_BY_NAME, $this->databaseName, $collection);
            $this->response = $this->connection->get($requestURL);

            return ($this->response->getStatusCode() == 404) ? false : true;
        }
        catch (RequestException $requestException ) {
            return false;
        }
    }

    public function getFieldsPrefixed( $fields, $prefix ){
        if( is_array( $fields ) ) {
            array_walk( $fields, function (&$value, $index) use ($prefix) {
                $value = sprintf('%s.%s', $prefix, $value);
            });

            return implode( ",", $fields );
        }
    }

    public function getQueryParts( $fields, $prefix ){
        if( is_array( $fields ) ) {
            array_push( $fields, "_rev", "_id", "_key" );

            array_walk($fields, function (&$value, $key) use ($prefix) {
                $value = sprintf('%s:%s.%s', $value, $prefix, $value);
            });
            return "{" . implode(",", $fields) . "}";
        }
    }

    public function getVersionInfo( $details = false ){
        $requestURL = sprintf( AarangoDialect::GET_VERSION, $this->databaseName, (($details)?"?details=true":"") );
        $this->response = $this->connection->get( $requestURL );
        $contents = $this->response->getBody()->getContents();

        return json_decode( $contents, true);
    }
}