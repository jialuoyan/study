package vender

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
    "time"
    "strconv"
)

func TestHttp() {
	url := os.Args[1:][0]
	if ( !strings.HasPrefix(url,"http") ){
		url = "http://"+url
	}
	fmt.Println(url)
	Curl(url)

}

func Curl(url string) {
	resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
        os.Exit(1)
    }
    b, err := ioutil.ReadAll(resp.Body)
    fmt.Println(resp.Status)
    resp.Body.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
        os.Exit(1)
    }
    fmt.Printf("%s\n", b)
}


func TestFetch() {
	urls := []string{}
	url :=  "http://127.0.0.1:8124/index/index"
	num,_ := strconv.Atoi(os.Args[1:][0])
	fmt.Printf("%v\n",num)
	for i := 0; i < num; i++ {
		urls = append(urls,url)
	}

	start := time.Now()
    ch := make(chan string)
    for _, url := range urls {
        go fetch(url, ch) // start a goroutine
    }
    for range urls {
        fmt.Println(<-ch) // receive from channel ch
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close() // don't leak resources
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

