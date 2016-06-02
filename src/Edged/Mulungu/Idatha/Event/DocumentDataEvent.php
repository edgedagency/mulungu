<?php namespace Edged\Mulungu\Idatha\Event;

use Symfony\Component\EventDispatcher\Event;

class DocumentDataEvent extends Event{

    protected $data;

    public function __construct(  array $data ){
        $this->data = $data;
    }

    public function setData( array $data ){
        $this->data = $data;
    }

    public function getData(){
        return $this->data;
    }
}