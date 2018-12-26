### Golang 安装
三种安装方式
```
brew install golang

apt install golang-go

yum install golang
```
安装完成后 可使用 `go help` 查看具体帮助命令


### 使用 govendor 包管理
```
go get -u github.com/kardianos/govendor  //安装在 bin 目录下

govendor init //初始化vendor 目录

govendor add +external //将项目依赖的外部包引入到 vendor 目录

```
