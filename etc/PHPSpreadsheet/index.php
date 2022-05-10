<?php

ini_set('memory_limit','4096M');

use PhpOffice\PhpSpreadsheet\Spreadsheet;

require __DIR__ . '/../Header.php';

function getRandomString($length = 6) {
    $characters = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
    $string = '';

    for ($i = 0; $i < $length; $i++) {
        $string .= $characters[mt_rand(0, strlen($characters) - 1)];
    }

    return $string;
}

$spreadsheet = new Spreadsheet();
$sheet = $spreadsheet->getActiveSheet();

for ($r=1; $r <= 102400; $r++) {
    for ($c=1; $c <= 50; $c++) {
        $sheet->setCellValueByColumnAndRow($c, $r, getRandomString());
    }
}

$helper->write($spreadsheet, './PhpSpreadsheet_r102400xc50.xlsx', ['Xlsx']);