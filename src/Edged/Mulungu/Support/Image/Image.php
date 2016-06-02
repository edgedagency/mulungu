<?php Edged\Mulungu\Support\Document\Image;

class Image {

    function generateThumbnail( $from, $to, $width, $height){
        $imagick = new Imagick ( $from );
        $imagick->thumbnailImage( $width, $height );

        return $imagick->writeImage( $to );
    }
}