<?php namespace Edged\Mulungu\Support\HTTP;

use GuzzleHttp\Client;
use Psr\Http\Message\ResponseInterface;
use Psr\Log\LoggerInterface;

class HTTPSupport {

    private static $promises = [];
    private  $logger;

    public function __construct( LoggerInterface $logger = null ){
        $this->logger = $logger;
    }

    public function downloadAndSave( $fromURL, $toFile ){
        if( $this->logger ){
            $this->logger->debug( "Downloading from " . $fromURL . " to location " . $toFile );
        }

        $client = new Client();
        $client->get($fromURL,['save_to' => $toFile ]);
    }
}