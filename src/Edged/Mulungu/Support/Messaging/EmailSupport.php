<?php namespace Edged\Mulungu\Support\Messaging;

use Pelago\Emogrifier;

class EmailSupport {

    public function inline( $css, $email ){
        $emogrifier = new Emogrifier( $email, $css );
        return $emogrifier->emogrify();
    }
}