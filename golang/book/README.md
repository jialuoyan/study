### Makefile 使用命令
```
make  //生成可执行文件
make run  //运行main
make clean
```

2018-12-20
### go 知识点
- go 语言不允许使用无用的局部变量，所以在有的时候，使用 `_` 来表示

os
```
os.Args[1:] //获取多个参数，从1 开始获取 ，例如 go run main.go aaa bbb ,获取到的是 aaa bbb
os.Exit(1) //退出
```

strings
```
strings.Join(array, string $glue) // 将一个数组的值转化为字符串
strings.HasPrefix(string, str) bool //str 是否在string里面
```

time
```
time.Now().Format("2006-01-02 15:04:05") //格式化当前时间,记忆方式 6-1-2-3-4-5

```

net/http
```
http.Get(url)
```
