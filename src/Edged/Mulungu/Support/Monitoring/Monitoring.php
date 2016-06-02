<?php namespace Edged\Mulungu\Support\Document\Monitoring;

class Monitoring {

    public function startTransaction( $unique ){
        if (extension_loaded('newrelic')) {
            newrelic_name_transaction ( $unique );
        }
    }

    public function endTransaction(){

    }
}