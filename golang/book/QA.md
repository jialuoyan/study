#### GOPATH 环境变量设置
https://studygolang.com/articles/4273
```
1. 可以设置临时环境变量GOPATH，export 即可
2. 用户级别设置
	vi ~/.bash_profile
	//在文件的结尾加上
	export GOPATH=/usr/local/src/go_path

	//使配置马上生效
	source /etc/profile
3. 全局设置
	vi /etc/profile

	//在文件的结尾加上
	export GOPATH=/usr/local/src/go_path

	//使配置马上生效
	source /etc/profile

```

#### GOPATH，GOROOT 问题
https://blog.csdn.net/quicmous/article/details/80360126
GO 语言提供两个关键环境变量，GOROOT 指向系统安装路径，GOPATH指向工作路径

GOROOT = /usr/local/go
GOPATH = /Users/yanjialuo/Documents/study/study/golang/book

````conf
/usr/local/go    <<--- GOROOT 指向的位置
    --src                 <<--- Go 语言自带的源代码
    --pkg                 <<--- 编译的中间文件放在此文件夹
    --bin                 <<--- 编译的目标文件放在此文件夹
/Users/yanjialuo/Documents/study/study/golang/book  <<--- GOPATH 指向的位置
    --src                 <<--- 项目源代码放置在此文件夹。!!!警告：一个常犯的错误是把 GOPATH 指向此处!!!
        --HelloWorld      <<--- 我们项目源代码所在的文件夹。!!!警告：一个常犯的错误是把 GOPATH 指向此处!!!
        --vendor          <<--- 第三方开源代码文件夹
            --github.com
                --...
    --pkg                 <<--- 编译的中间文件放在此文件夹，Go编译器自动生成此文件夹
    --bin                 <<--- 编译的目标文件放在此文件夹，Go编译器自动生成此文件夹
````