// test os lib
package vender
import (
    "fmt"
    "os"
    "strings"
    "time"
    "bufio"
)


func TestOslib() {
    var s, s2, s3 string
    var st1 , st2, st3, st4 time.Time
    st1 = time.Now()
    for _, arg := range os.Args[1:] {
        s += s2 + arg + " "
    }
    st2 = time.Now()
    fmt.Println(st1, st2)
    fmt.Printf("%v\n", st1);
    fmt.Println(s)
    fmt.Println("##################")
    st3 = time.Now()
    s3 = strings.Join(os.Args[1:]," ")
    st4 = time.Now()
    fmt.Println(st3, st4)
    fmt.Println(s3)
    fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
    // fmt.Println(os.Args[1:])
}

func TestF() {
   // for i := 0; i < 10; i++ {
   //   fmt.Println(i);
   // }
    t, _ := time.Parse("2006-01-02 15:04:05", "2016-04-20 16:23:00")
    fmt.Println(t.Unix())
    y, m, d := time.Unix(1466344320, 0).Date()
    fmt.Println(y, m, d)

    //format后面的字符串必须是2006-01-02 15:04:05，据说go是这个时间诞生的
    fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
    fmt.Println(time.Now().Format("2006-01-02"))
    fmt.Println(time.Now().Format("20060102"))
                                                                                                                            
    select {
    case <-time.After(5 * time.Second):
        fmt.Println("After 5 second")
    }   
    c := time.Tick(10 * time.Second)
    for now := range c { 
        fmt.Println(now)
    } 
}

func TestRepeat() {
  counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
    }
    // NOTE: ignoring potential errors from input.Err()
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}




