<?php namespace Edged\Mulungu\Request\Helper;

use Edged\Mulungu\Bundle\FrameworkBundle\Extension\ExtensionController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;

class RequestHelper {

	private $request;
    private $posted;
    private $queries;
    private $controller;

	public function __construct( Request $request, ExtensionController $controller ){
		$this->request = $request;
        $this->posted = $this->request->request->all();
        $this->queries = $this->request->query->all();
        $this->controller = $controller;
	}

    public function is( $method ){
        if( $this->request->getMethod() == $method ) {
            return true;
        }

        return false;
    }

    public function post( $key, $default=null ){
        return (array_key_exists($key, $this->posted ))?$this->posted [ $key ]:$default;
    }
    public function query( $key, $default=null ){
        return (array_key_exists($key, $this->queries ))?$this->queries[ $key ]:$default;
    }

	public function replace( $key , $value ){
		if( array_key_exists($key, $this->posted ) ){
			$this->posted[ $key ] = $value;
		}else if( array_key_exists($key, $this->queries ) ){
            $this->queries[ $key ] = $value;
        }
	}
    /*
     * Removes and returns item from data set takes into consideration both posted and query
     */
    public function extract( $key ){
        $extracted = null;
        if( array_key_exists($key, $this->posted ) ){
            $extracted = $this->posted [ $key ];
            unset( $this->posted [ $key ] );
        }else if( array_key_exists($key, $this->queries ) ){
            $extracted = $this->queries [ $key ];
            unset( $this->queries [ $key ] );
        }
        return $extracted;
    }

	public function has( $key ){
        if( array_key_exists($key, $this->posted() ) ){
			return true;
		}else if( array_key_exists($key, $this->queries()) ){
            return true;
        }
		return false;
	}

    public function posted(){
        return $this->posted;
    }

    public function queries(){
        return $this->queries;
    }

    public function isQuery(){
        if( count( $this->queries ) > 0 )
            return true;
    }

    public function isNotQuery(){
        if( count( $this->queries ) <= 0 )
            return true;
    }

    public function render( $format, $template, array $data, Response $response = null ){
        $response;

        switch( $format ) {
            case "html": $response = $this->controller->render( $template , $data ); break;
            case "json": $response = new JsonResponse( $data ); break;
        }

        return $response;
    }
}