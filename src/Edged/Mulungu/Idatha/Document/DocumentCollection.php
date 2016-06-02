<?php namespace Edged\Mulungu\Idatha\Document;

use Edged\Mulungu\Idatha\Document\Interfaces\DocumentInterface;
use Psr\Log\LoggerInterface;

class DocumentCollection{

    protected $entries;
	private $logger;

    public function __construct( LoggerInterface $logger = null ){
    	$this->logger = $logger;
    }

    public function isCollection(  ){
        return true;
    }

    public function isEmpty(){
    	
    	if( !isset( $this->entries ) )
    		return true;
    	else if( empty( $this->entries ) )
    		return true;
    	
    	return false;
    }
    
    public function toArray( $root=null, $mapInternalAttributes=false ){
        $entries = [];

        if( isset( $root ) )
            $entries[ $root ] = [];

        if( isset( $this->entries ) && !empty( $this->entries ) ){
            foreach( $this->entries as $entry ){
                if( isset( $root ) ){
                    array_push( $entries[ $root ] , $entry->toArray( null, $mapInternalAttributes ) );
                }else{
                    array_push( $entries , $entry->toArray( null, $mapInternalAttributes ) );
                }
            }
        }

        return $entries;
    }

    public function add( DocumentInterface $document ){
        if( !isset( $this->entries ) )
            $this->entries = [];

        array_push( $this->entries , $document );
    }

    public function first(){

        if( isset( $this->entries ) )
            return $this->entries[ 0 ];

        return null;
    }
} 