<?php namespace Edged\Mulungu\Idatha\Event;


use Edged\Mulungu\Idatha\Document\Document;
use Symfony\Component\EventDispatcher\Event;

class DocumentEvent extends Event{

    protected $document;

    public function __construct( Document $document ){
        $this->document = $document;
    }

    public function setDocument( Document $document ){
        $this->document = $document;
    }

    public function getDocument(){
        return $this->document;
    }
}