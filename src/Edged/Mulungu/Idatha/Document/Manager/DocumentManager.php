<?php namespace Edged\Mulungu\Idatha\Document\Manager;

use Edged\Mulungu\Idatha\Connection\Driver\Interfaces\ConnectionDriver;
use Edged\Mulungu\Idatha\Document\Document;
use Psr\Log\LoggerInterface;

class DocumentManager
{
    private $connectionDriver;
	private $logger;
	
    public function __construct( ConnectionDriver $connectionDriver, LoggerInterface $logger = null ){
        $this->connectionDriver = $connectionDriver;
        $this->logger = $logger;
    }

    public function save( Document $document ){

    }

    public function update( Document $document ){

    }

    public function delete( $id ){

    }
}