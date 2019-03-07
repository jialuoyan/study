<?php

namespace app\Lib;

use PhpOffice\PhpSpreadsheet\Spreadsheet;
use PhpOffice\PhpSpreadsheet\Writer\Xlsx;
use PhpOffice\PhpSpreadsheet\Helper\Sample;
use PhpOffice\PhpSpreadsheet\IOFactory;

class XlsExcel{

	/**
     * 获取文件保存路径
     */
    private function getSavePath(){
        $path =  dirname(dirname(__FILE__)).'/tmp/';
        if (!is_dir($path)) {
            if (!is_writable($path) && !@mkdir($path, 0777, true)) {
                throw new \Exception ('Failed to create  folder for Excel');
            }
        }
        return $path;
    }

    /**
     * 渲染表头
     * @param array  $header header
     * @param object $sheet  sheet
     * @param int $startIndex  头部起始行数
     * @return object
     */
    private function _renderHeader(array $header, $sheet, int $startIndex=1){
    	// 输出表头
        $index = 0;
        $kk=0;
        foreach ($header as $val) {
            if($index<26){
                $cellPos = sprintf('%s%s', chr(ord('A') + $index++), $startIndex);
            }else{
                $cellPos = sprintf('%s%s', chr(ord('A') + $kk++), $startIndex);
                $cellPos = "A".$cellPos;
            }
            $sheet->setCellValue($cellPos, $val);
        }
        return $sheet;
    }

    /**
     * 渲染内容
     * @param array  $header header
     * @param array  $data   data
     * @param object $sheet  sheet
     * @param int $startLine  内容起始行数
     * @return object
     */
    private function _renderBodyer($header, $data, $sheet, $startLine=1){
        foreach ($data as $row) {
            $index = 0;
            $kk=0;
            foreach ($header as $key => $val) {
                if($index<26){
                    $cellPos = sprintf('%s%s', chr(ord('A') + $index++), $startLine);
                }else{
                    $cellPos = sprintf('%s%s', chr(ord('A') + $kk++), $startLine);
                    $cellPos = "A".$cellPos;
                }
                $cellval = "";
                if(is_array($row) && isset($row[$key])){
                    $cellval= $row[$key];
                }else if(isset($row->$key)){
                    $cellval=$row->$key;
                }
                // echo $cellPos.'-'.$cellval.PHP_EOL;
                $sheet->setCellValue($cellPos, $cellval);
            }
            $startLine++;
        }
    }

	/**
	 * 创建excel
	 * 
	 */
	public function createExcel($header, $data, $file_name){
	    // Create new PHPExcel object
	    echo date('H:i:s') , " Create new PHPExcel object" , PHP_EOL;

	    $excel = new Spreadsheet();
        $excel->getProperties()->setTitle($file_name);
        $excel->setActiveSheetIndex(0);
        $sheet = $excel->getActiveSheet();
        $this->_renderHeader($header, $sheet, 1);
        if(!empty($data)){
            $this->_renderBodyer($header, $data, $sheet, 2);
        }

        //获取文件存储路径 创建文件目录 添加读写权限
        $path = $this->getSavePath();
        $fullname = $path.$file_name.'.xlsx';

        $writer = IOFactory::createWriter($excel, 'Xlsx');
        $writer->save($fullname);

        echo date('H:i:s') , " Create PHPExcel finish" , PHP_EOL;

        return $fullname;

	}

}