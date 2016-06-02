<?php namespace Edged\Mulungu\Excel;

class ExcelCell {

    private $heading;
    private $expression;
    private $expressionDefault;

    public function __construct( $heading, $expression, $expressionDefault ){
        $this->heading = $heading;
        $this->expression = $expression;
        $this->expressionDefault = $expressionDefault;
    }

    /**
     * @return mixed
     */
    public function getHeading()
    {
        return $this->heading;
    }

    /**
     * @param mixed $heading
     */
    public function setHeading($heading)
    {
        $this->heading = $heading;
        return $this;
    }

    /**
     * @return mixed
     */
    public function getExpression()
    {
        return $this->expression;
    }

    /**
     * @param mixed $expression
     */
    public function setExpression($expression)
    {
        $this->expression = $expression;
        return $this;
    }

    /**
     * @return mixed
     */
    public function getExpressionDefault()
    {
        return $this->expressionDefault;
    }

    /**
     * @param mixed $expressionDefault
     */
    public function setExpressionDefault($expressionDefault)
    {
        $this->expressionDefault = $expressionDefault;
        return $this;
    }


}