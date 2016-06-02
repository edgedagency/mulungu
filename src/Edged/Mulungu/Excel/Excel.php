<?php namespace Edged\Mulungu\Excel;

use Symfony\Component\ExpressionLanguage\ExpressionLanguage;

class Excel {

    const EXCEL_2007 = 'Excel2007';
    const PHP_OUT = 'php://output';

    const NEW_ROW = 1;

    private $expressionLanguage;
    private $document;

    public function __construct(){
        $this->expressionLanguage = new ExpressionLanguage();
        $this->document = new \PHPExcel();
    }

	private function addSheetColumnItem( $worksheet, $column, $row, $item, $evalute=false, $dataset=[] ){
        $item = ($evalute)?$this->processExpression( $item, $dataset ):$item;
		$worksheet->setCellValueByColumnAndRow( $column, $row, $item );
	}

    public function create( $configuration, $saveAs, $worksheetLabel, $type=Excel::EXCEL_2007 ){

        $worksheet = new \PHPExcel_Worksheet( $this->document, $worksheetLabel);

        $this->document->addSheet( $worksheet, 0 );

        //creating sheet headers
        $sheetHeadingColumnIndex = 0;
        foreach( $configuration["item"] as $sheetHeaderItemKey=>$sheetHeaderItemValue){
            if( is_array( $sheetHeaderItemValue ) ){
                foreach(  $sheetHeaderItemValue[ "item" ] as $sheetHeaderSubItemKey=>$sheetHeaderSubItemValue ){
					$this->addSheetColumnItem( $worksheet, $sheetHeadingColumnIndex, 1, $sheetHeaderSubItemKey );
					$sheetHeadingColumnIndex ++;
                }
            }else{
				$this->addSheetColumnItem( $worksheet, $sheetHeadingColumnIndex, 1, $sheetHeaderItemKey);
				$sheetHeadingColumnIndex ++;
            }
        }

        //initialize variables
        $this->processSheetItems( $worksheet, 0, 2, $configuration );

        $objWriter = \PHPExcel_IOFactory::createWriter( $this->document, $type );
        $objWriter->save( $saveAs );
    }

    private function processSheetItems( $worksheet, $column, $row, $configuration ){

        $initColumn = $column;
        $initRow = $row;

        foreach( $configuration["data"] as $data){

            $recursed = false;
            $rowBeforeRecursion = 0;
            $recurseColumCount = 0;
            $recurserowCount = 0;

            $column = $initColumn;


            foreach( $configuration["item"] as $key=>$value ){
                if( is_array( $value ) ){
                    //obtain data needed from expressive declarations
                    $value["data"] = $this->processExpression( $value["data"], $data );
                    //process subconfigurations, maintain column and shift rows down if configured to do so

                    $recursed = true;
                    $rowBeforeRecursion = $row + 1;
                    $recurseColumCount = count( $value['item'] );
                    $$recurserowCount = count(  $value["data"] );


                    $this->processSheetItems( $worksheet, $column, $rowBeforeRecursion, $value );

                }else{

                    if( $recursed === true and $recurseColumCount > 0) {
                        //reset row to where it was before recursion,
                        $row    = $rowBeforeRecursion-1;
                        $column = $column + $recurseColumCount;
                        $recurseColumCount = 0;
                    }

                    $this->addSheetColumnItem( $worksheet, $column, $row, $value, true, $data);
                    $column++;
                }
            }

            if( $recursed ) {
                $recursed = false;
                $row += $rowBeforeRecursion + $recurserowCount + 1;
                $recurserowCount = 0;
                $rowBeforeRecursion = 0;
                $column = 0;
            }else{
                $row ++;
            }
        }
    }

    private function processExpression( $expression, $expressionData){
        try {
            return $this->expressionLanguage->evaluate( $expression, $expressionData);
        }catch(\Exception $ex ){
            return $ex->getMessage();
        }
    }
}