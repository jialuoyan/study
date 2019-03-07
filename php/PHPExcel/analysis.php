<?php

error_reporting(E_ALL);
set_time_limit(0);

date_default_timezone_set('Asia/Shanghai');

require_once __DIR__ . '/vendor/autoload.php';

require_once __DIR__.'/default.php';

$header = getHeader();
$data = getBodyData();

// $client = new \GuzzleHttp\Client();
// $response = $client->request('GET','httpbin.org/get',['query'=>'','timeout'=>10])
//     ->getBody()->getContents();
// $response = \GuzzleHttp\json_decode($response, true);

// $data[] = [
//     'key1' => $response['origin'],
//     'key2' => $response['url']
// ];
// var_dump($response);


$file_name = '文件名';

$r = (new app\Lib\XlsExcel())->createExcel($header, $data, $file_name);
