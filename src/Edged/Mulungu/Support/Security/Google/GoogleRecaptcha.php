<?php namespace Edged\Mulungu\Support\Security\Google;


use GuzzleHttp\Client;
use Psr\Log\LoggerInterface;

class GoogleRecaptcha{

    private $googleRecaptchaSecret;
    private $googleRecaptchaEnabled;
    private $googleRecaptchaVerificationURL;
    private $logger;

    public function __construct( $googleRecaptchaSecret, $googleRecaptchaEnabled=false, $googleRecaptchaVerificationURL='https://www.google.com/recaptcha/api/siteverify', LoggerInterface $logger ){
        $this->googleRecaptchaSecret=$googleRecaptchaSecret;
        $this->googleRecaptchaEnabled=$googleRecaptchaEnabled;
        $this->googleRecaptchaVerificationURL=$googleRecaptchaVerificationURL;
        $this->logger = $logger;
    }

    public function verify( $submitedResponse, $remoteip=null ){

        $clientRequest=[];
        $clientRequest["form_params"] = [
            "secret"=>$this->googleRecaptchaSecret,
            "response"=>$submitedResponse
        ];

        $clientRequestHandler = new Client();
        $clientResponse = $clientRequestHandler->post( $this->googleRecaptchaVerificationURL, $clientRequest);
        $clientResponseContent = $clientResponse->getBody()->getContents();
        $reponse = json_decode( $clientResponseContent, true );

        if(isset( $this->logger ) ){
            $this->logger->debug( "google recaptcha verififcation response", (is_array($reponse))?$reponse:["no response data"] );
        }

        return $reponse[ 'success' ];
    }
}