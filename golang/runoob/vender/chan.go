package vender

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // 把 sum 发送到通道 c
}

func TestChan(){
	s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c) // s[:len(s)/2] 表示 从0 - （len(s)/2-1） 位
    fmt.Println(<-c)
    // go sum(s[:len(s)/2], c)
    // go sum(s[len(s)/2:], c)
    // // x, y := <-c, <-c // 从通道 c 中接收
    // x := <-c
    // y := <-c

    // fmt.Println(x, y, x+y)
}