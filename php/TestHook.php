<?php
/**
* 钩子函数示例及应用
*/

class Test{
	public function register($class){
		//注册事件
		Hook::add($class);
	}

	public function do(){
		Hook::exec();
	}
}

//定义钩子
class Hook{
	private static $_hooklisk = [];

	//添加
	public static function add($class){
		self::$_hooklisk[] = new $class();
	}
	//触发事件
	public static function exec(){
		foreach (self::$_hooklisk as $key => $value) {
			$value->act();
		}
	}
}

class T1{
	public function act(){
		echo 'class name is '. __CLASS__ .PHP_EOL; 
	}
}

class T2{
	public function act(){
		echo 'class name is '. __CLASS__ .PHP_EOL; 
	}
}

$t = new Test();
$t->register('T1');
$t->register('T2');
$t->do();
