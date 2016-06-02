<?php namespace Edged\Mulungu\Idatha\Query;


class QueryBuilder{

    protected $query;
    protected $collection;
    protected $collections;

    public function __construct( $collection, $collections ){

        $this->collection = $collection;
        $this->collections = $collections;

        $this->query = [];
        $this->query[] = sprintf("FOR %s IN %s", $this->collection, $this->collections );
    }

    public function addFor( $for ){
        return $this;
    }

    public function addFilter( $leftSide, $operations="==", $rightSide ){
        $this->query[] = sprintf("FILTER %s %s %s", $leftSide, $operations, $rightSide );
        return $this;
    }

    public function addLet( $variableName, QueryBuilder $query){
        return $this;
    }

    public function addColumns( array $columns ){
//        $this->query[] = ["RETURN"=>];
//
//        return $this;
    }
    /**
     * @param $id
     * @param array|null $columns
     * @return $this
     */
    public function find( $id, array $columns=null ){
        return $this;
    }

    /**
     * @param array $filters
     * @param array|null $columns
     * @return $this
     */
    public function findAll( array $filters, array $columns=null ){
        return $this;
    }

    public function compile(){
        $compiledQuery = "";

        foreach( $this->query as $index=>$value ){

            switch( $index ) {
                case "RETURN":

                    break;
                default : $compiledQuery .= $value . " ";
            }

        }

        return $compiledQuery;
    }
}