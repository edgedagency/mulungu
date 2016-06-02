<?php namespace Edged\Mulungu\Support;

use Edged\Mulungu\Helper\StringHelper;

class ClassLoader {

    public static $registered = false;
    public static $searchPath = [];

    public static function load( $class ){

        if( isset($class) && !empty( $class ) ) {

            if( strpos( $class, "Extension" ) === 0  ) {
                return self::loadExtensionClass($class);

            }else {
                foreach (self::$searchPath as $path) {
                    $impliedClassPath = $path . DIRECTORY_SEPARATOR . self::getClassDirectory($class) . $class . '.php';
                    if (is_file($impliedClassPath)) {
                        require_once $impliedClassPath;
                        return true;
                    }
                }
            }
        }
        return false;
    }

    private static function loadExtensionClass( $class ){
        foreach (self::$searchPath as $path) {
            $impliedClassPath = $path . DIRECTORY_SEPARATOR . self::getClassDirectory($class) . DIRECTORY_SEPARATOR . $class . '.php';

            if (is_file($impliedClassPath)) {
                require_once $impliedClassPath;
                return true;
            }
        }
    }

    public static function getClassDirectory( $class ){
        if( strpos( $class, "_" ) ) {
            $classDirectoryPart = substr($class, 0, strpos($class, "_"));
            return StringHelper::from_camel_case($classDirectoryPart);
        }
    }

    public static function register(){
        if( !self::$registered ){
            self::$registered = spl_autoload_register( [ '\Edged\Mulungu\Support\ClassLoader','load'] );
        }
    }

    public static function addSearchPath( $searchPath ){
        array_push( self::$searchPath, $searchPath );
    }
} 