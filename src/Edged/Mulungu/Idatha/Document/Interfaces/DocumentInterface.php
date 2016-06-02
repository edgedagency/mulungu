<?php namespace Edged\Mulungu\Idatha\Document\Interfaces;

use Edged\Mulungu\Idatha\Connection\Driver\Interfaces\ConnectionDriver;

interface DocumentInterface
{
    public function getType();
    public function setConnection( ConnectionDriver $connection );
    public function getConnection();
}