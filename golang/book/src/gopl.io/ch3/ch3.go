package ch3

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

func TestInt() {

	//整形溢出
	// var u uint8 = 255
	// // fmt.Println(u, u+1, u*u)
	// fmt.Println(u<<2)

	var f float32 = 16777216 // 1 << 24
	// f == f+1
	fmt.Println(f+1)    // "true"!
	var m float64 = 16777216
	// m == m+1
	fmt.Println(m+1)
}

/**
* 测试类入口
*/
func TestFunc() {
	rand.Seed(time.Now().UnixNano())
	nums := []float64{}
	for i := 0; i < 10; i++ {
		nums = append(nums, 500*rand.Float64())
	}
	fmt.Println("init data: ",nums)
	fmt.Println(GetRoundNum(nums))
}

func GetRoundNum(param []float64)float64 {
	sort.Float64s(param)
	fmt.Println("sort data: ",param)
	var l = len(param)/2
	fmt.Printf("slice:%d\n",l)
	return param[l]
}