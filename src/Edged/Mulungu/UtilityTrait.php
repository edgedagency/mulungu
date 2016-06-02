<?php namespace Edged\Mulungu;

use Symfony\Component\EventDispatcher\Event;

trait UtilityTrait {

    public $logger;
    public $eventDispatcher;

    public function log( $message, $level="DEBUG", $context=null ){
        if( isset( $this->logger ) ){
            $this->logger->log( $level, $message );
        }
    }

    public function dispatch( $name, Event $event ){
        if( isset( $this->eventDispatcher ) ) {
            return $this->eventDispatcher->dispatch($name, $event);
        }
    }
}