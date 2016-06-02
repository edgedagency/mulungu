<?php
/**
 * Created by PhpStorm.
 * User: KOGA
 * Date: 15/03/28
 * Time: 11:10 PM
 */

namespace Edged\Mulungu\Support\Document;


class BarCodeGenerator {

    public function oneDimensional( $fileName, $author, $title, $subject, $keyWords, $mode  ){

        $pdf = new \TCPDF(PDF_PAGE_ORIENTATION, PDF_UNIT, "A5", true, 'UTF-8', false);
        $pdf->SetCreator(PDF_CREATOR);
        $pdf->SetAuthor($author);
        $pdf->SetTitle($title);
        $pdf->SetSubject($subject);
        $pdf->SetKeywords($keyWords);

// set default header data
        //$pdf->SetHeaderData(PDF_HEADER_LOGO, PDF_HEADER_LOGO_WIDTH, $title , $author);

// set header and footer fonts
        $pdf->setHeaderFont(Array(PDF_FONT_NAME_MAIN, '', PDF_FONT_SIZE_MAIN));
        $pdf->setFooterFont(Array(PDF_FONT_NAME_DATA, '', PDF_FONT_SIZE_DATA));

// remove default header/footer
        $pdf->setPrintHeader(false);
        $pdf->setPrintFooter(false);
// set default monospaced font
        $pdf->SetDefaultMonospacedFont(PDF_FONT_MONOSPACED);

// set margins
        $pdf->SetMargins(PDF_MARGIN_LEFT, PDF_MARGIN_TOP, PDF_MARGIN_RIGHT);
        $pdf->SetHeaderMargin(PDF_MARGIN_HEADER);
        $pdf->SetFooterMargin(PDF_MARGIN_FOOTER);

// set auto page breaks
        $pdf->SetAutoPageBreak(TRUE, PDF_MARGIN_BOTTOM);

// set image scale factor
        $pdf->setImageScale(PDF_IMAGE_SCALE_RATIO);

// set some language-dependent strings (optional)
        if (@file_exists(dirname(__FILE__).'/lang/eng.php')) {
            require_once(dirname(__FILE__).'/lang/eng.php');
            $pdf->setLanguageArray($l);
        }


// set a barcode on the page footer
        $pdf->setBarcode(date('Y-m-d H:i:s'));

// set font
        $pdf->SetFont('helvetica', '', 11);

// add a page
        $pdf->AddPage();

// print a message
//        $txt = "You can also export 1D barcodes in other formats (PNG, SVG, HTML). Check the examples inside the barcodes directory.\n";
//        $pdf->MultiCell(70, 50, $txt, 0, 'J', false, 1, 125, 30, true, 0, false, true, 0, 'T', false);
//        $pdf->SetY(30);

// -----------------------------------------------------------------------------

        $pdf->SetFont('helvetica', '', 10);

// define barcode style
        $style = array(
            'position' => '',
            'align' => 'C',
            'stretch' => true,
            'fitwidth' => true,
            'cellfitalign' => '',
            'border' => true,
            'hpadding' => 'auto',
            'vpadding' => 'auto',
            'fgcolor' => array(0,0,0),
            'bgcolor' => false, //array(255,255,255),
            'text' => true,
            'font' => 'helvetica',
            'fontsize' => 8,
            'stretchtext' => 4
        );

// PRINT VARIOUS 1D BARCODES

// CODE 39 - ANSI MH10.8M-1983 - USD-3 - 3 of 9.
//        $pdf->Cell(0, 0, 'CODE 39 - ANSI MH10.8M-1983 - USD-3 - 3 of 9', 0, 1);
        $pdf->write1DBarcode('KAT03-1F14', 'C128', '', '', '', 18, 0.4, $style, 'N');

        $pdf->Ln();

        $pdf->lastPage();
//Close and output PDF document
        $pdf->Output( $fileName , $mode );

    }

    public function outputHTML( $code, $width=2, $height=30, $color="black"){
        $barcode = new \TCPDFBarcode( $code, 'C128');
        return $barcode->getBarcodeHTML( $width, $height, $color );
    }

    public function outputPNG( $code, $width=2, $height=30, $color = array(0,0,0) ){
        $barcode = new \TCPDFBarcode( $code, 'C128');
        $barcode->getBarcodePNG( $width, $height, $color);
    }
}