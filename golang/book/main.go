/**
study main函数
export GOPATH=/Users/yanjialuo/Documents/study/study/golang/book
*/
package main
// import "vender"

import (
	"gopl.io/shell"
	// "math"
)

func main() {
	// ch3.TestInt()
	// arr := [...]float32{ 10.1, 10.2, 20.1, 20.2,15.7}
	// // q := [...]int{1, 2, 3}
	// fmt.Println(arr);
	// fmt.Println(os.Args[1:])
	// ch3.TestFunc()
	var sh = "ls -lh"
	shell.Execute(sh)

}
