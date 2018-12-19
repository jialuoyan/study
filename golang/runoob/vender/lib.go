package vender
import "fmt"

/**
* 输出素数
* @param int floor
*/
func Isprimary(floor int) {
    fmt.Println(floor,"以内的素数有：")
    var C, c int//声明变量
    C=1 /*这里不写入FOR循环是因为For语句执行之初会将C的值变为1，当我们goto A时for语句会重新执行（不是重新一轮循环）*/
    A: for C < floor {
       C++ //C=1不能写入for这里就不能写入
       for c=2; c < C ; c++ {
           if C%c==0 {
               goto A //若发现因子则不是素数
           }
       }
       fmt.Print(C," ")
    }
    fmt.Println()
}

/**
* 输出杨辉三角
* @param int floor
*/
func ShowYangHuiTriangle(floor int) {
  fmt.Println(floor,"阶杨辉三角")
  nums := []int{}
    for i := 0; i < floor; i++ {
        //补空白
        for j := 0; j < (floor - i); j++ {
            fmt.Print(" ")
        }
        for j := 0; j < (i + 1); j++ {
            var length = len(nums)
            var value int
            if j == 0 || j == i {
                value = 1
            } else {
                value = nums[length-i] + nums[length-i-1]
            }
            nums = append(nums, value)
            fmt.Print(value, " ")
        }
        fmt.Println("")
    }
  fmt.Println()
}

/**
* 返回两个数小的
*/
func ReturnMin(num1 int, num2 int) int {
	var result int;
	if (num1 > num2) {
		result = num2
	}else{
		result = num1
	}
	// fmt.Println(result);
	return result;
}

func ReturnMax(num1 int, num2 int) int {
   var result int
   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   // fmt.Println(result);
   return result 
}


