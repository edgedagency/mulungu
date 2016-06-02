<?php namespace Edged\Mulungu\Idatha\Connection\Driver\Interfaces;

use Edged\Mulungu\Idatha\Document\Document;

interface ConnectionDriver {

    const SORT_ASC = "ASC";
    const SORT_DESC = "DESC";

    public function getVersionInfo();
    public function save( Document $document );
    public function find( $collection, $id, $count=false, $fields=null );
    public function delete( $id );
}