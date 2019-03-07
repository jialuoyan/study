<?php

function getHeader(){
    return [
        'orgcode'=>'orgcode',
        'orgname'=>'orgname',
        'key1'    =>'测试字段1',
        'key2'    =>'测试字段2',
        'key3'    =>'测试字段3',
        'key4'    =>'测试字段4',
        'key5'    =>'测试字段5',
        'key6'    =>'测试字段6',
        'key7'    =>'测试字段7',
        'key8'    =>'测试字段8',
    ];
}

function getBodyData(){
	 return [
        ['orgcode'=>'2000V0','orgname'=>'机构1','key1'=>10,'key2'=>20,'key3'=>20,'key4'=>30,'key5'=>40,'key6'=>50,'key7'=>60,'key8'=>70],
        ['orgcode'=>'2000VO01','orgname'=>'机构2','key1'=>10,'key2'=>20,'key3'=>20,'key4'=>30,'key5'=>30,'key6'=>30,'key7'=>30,'key8'=>30],
        ['orgcode'=>'2000VO02','orgname'=>'机构3','key1'=>10,'key2'=>20,'key3'=>20,'key4'=>30,'key5'=>30,'key6'=>30,'key7'=>30,'key8'=>30]
    ];
}