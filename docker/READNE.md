
# docker 
docker COMMAND --help 查看具体详情
## docker run [command]
命令格式：docker run [OPTIONS] IMAGE [COMMAND] [ARG...]
Usage: Run a command in a new container
- 常用选项说明
	- -d, --detach=false， 指定容器运行于前台还是后台，默认为false
    - -i, --interactive=false， 打开STDIN，用于控制台交互
    - -t, --tty=false， 分配tty设备，该可以支持终端登录，默认为false
    - -u, --user=""， 指定容器的用户
    - -a, --attach=[]， 登录容器（必须是以docker run -d启动的容器）
    - -w, --workdir=""， 指定容器的工作目录
    - -e, --env=[]， 指定环境变量，容器中可以使用该环境变量
    - -m, --memory=""， 指定容器的内存上限
    - -P, --publish-all=false， 指定容器暴露的端口
    - -p, --publish=[]， 指定容器暴露的端口
    - -h, --hostname=""， 指定容器的主机名
    - -v, --volume=[]， 给容器挂载存储卷，挂载到容器的某个目录
    - --volumes-from=[]， 给容器挂载其他容器上的卷，挂载到容器的某个目录
    - --cidfile=""， 运行容器后，在指定文件中写入容器PID值，一种典型的监控系统用法
    - --cpuset=""， 设置容器可以使用哪些CPU，此参数可以用来容器独占CPU
    - --device=[]， 添加主机设备给容器，相当于设备直通
    - --dns=[]， 指定容器的dns服务器
    - --dns-search=[]， 指定容器的dns搜索域名，写入到容器的/etc/resolv.conf文件
    - --entrypoint=""， 覆盖image的入口点
    - --env-file=[]， 指定环境变量文件，文件格式为每行一个环境变量
    - --expose=[]， 指定容器暴露的端口，即修改镜像的暴露端口
    - --link=[]， 指定容器间的关联，使用其他容器的IP、env等信息
    - --lxc-conf=[]， 指定容器的配置文件，只有在指定--exec-driver=lxc时使用
    - --name=""， 指定容器名字，后续可以通过名字进行容器管理，links特性需要使用名字
    - --net="bridge"， 容器网络设置:
    - bridge 使用docker daemon指定的网桥
    - host //容器使用主机的网络
    - container:NAME_or_ID >//使用其他容器的网路，共享IP和PORT等网络资源
    - none 容器使用自己的网络（类似--net=bridge），但是不进行配置
    - --privileged=false， 指定容器是否为特权容器，特权容器拥有所有的capabilities
    - --restart="no"， 指定容器停止后的重启策略:
    - no：容器退出时不重启
    - on-failure：容器故障退出（返回值非零）时重启
    - always：容器退出时总是重启
    - --rm=false， 指定容器停止后自动删除容器(不支持以docker run -d启动的容器)
    - --sig-proxy=true， 设置由代理接受并处理信号，但是SIGCHLD、SIGSTOP和SIGKILL不能被代理
- 部分示例
	- docker run -t -i -p 8090:80 image // 以image 镜像启动一个容器，将本机的8090端口映射到容器的80端口
	- docker run -p 81:80 --name mynginx -v $PWD/web:/www -d nginx //以nginx 镜像运行一个container ，并将当前目录下的web目录挂载到/www目录，将本机 81端口映射到container 80端口
	- docker run --rm -v $PWD/web:/www -w /data --entrypoint sh eureka-agent:dev-1.0.2 -c "/alpha.sh composer install" //将本目录与容器 /www 目录共享，/data 作为工作目录，容器启动时，执行 /alpha.sh composer install 命令
- 更多信息参考： https://docs.docker.com/engine/reference/commandline/run/

## docker exec [command]
命令格式：docker exec [OPTIONS] IMAGE [COMMAND] [ARG...]
进入运行的容器
- 常用选项说明
	- -d :分离模式: 在后台运行
	- -i :即使没有附加也保持STDIN 打开
	- -t :分配一个伪终端
- 部分示例
	- docker exec -ti image /bin/bash //进入正在运行的image 容器 （docker ps 可查看当前运行了哪些容器）

## docker build [command]
命令格式: docker build [OPTIONS] PATH | URL | -
使用Dockerfile 构建镜像
- 常用选项说明
	- --build-arg=[] :设置镜像创建时的变量
	- -f :指定要使用的Dockerfile路径；
	- -force-rm :设置镜像过程中删除中间容器；
	- --pull :尝试去更新镜像的新版本；
	- --tag, -t: 镜像的名字及标签，通常 name:tag 或者 name 格式；可以在一次构建中为一个镜像设置多个标签。
	- --network: 默认 default。在构建期间设置RUN指令的网络模式
- 部分示例
	-  docker build -t image-screenshot:1.0 .  //根据Dockerfile 生成一个image

## docker rm
删除容器

## docker rmi
删除镜像

## docker ps
查看当前运行容器

## docker kill 
Usage:  docker kill [OPTIONS] CONTAINER [CONTAINER...]
Kill one or more running containers

## docker search
Usage:  docker search [OPTIONS] TERM
Search the Docker Hub for images
- 部分示例
	- docker search php

## docker pull
Usage:  docker pull [OPTIONS] NAME[:TAG|@DIGEST]
Pull an image or a repository from a registry
- 部分示例
	- docker pull php

## docker login
Usage:	docker login [OPTIONS] [SERVER]
- 部分示例
	- docker login -u test -p *** ip

## docker tag
Usage:	docker tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]
- 部分示例
	- docker tag print23/screen:0.0.1  test/screen:0.0.1

## docker push
Usage:	docker push [OPTIONS] NAME[:TAG]
- 部分示例
	- docker push test/screen:0.0.1 (push 前需要先docker login)

## docker stats
查看容器运行状态

## 部分常用命令
```
docker ps -a //查看所有的容器
docker rmi `docker images | grep none | awk '{print $3}'` //删除所有临时生成的 image
docker rm `docker ps -a -q` //删除所有 container
```

# docker-compoae
Define and run multi-container applications with Docker
官方文档地址：https://docs.docker.com/compose/reference/overview/

## docker-compose
- 常用选项说明
    - -f, --file FILE      Specify an alternate compose file  (default: docker-compose.yml)
    - -p, --project-name NAME
- docker-compose.yml 文件
```
webapp:
  image: examples/web
  ports:
  	- 30124:30124/tcp
    - 8124:80/tcp
  volumes:
    - ./:/data/web_data/alpha_framework/apps/web:rw
```
	- environment //环境变量
	- ports //端口号
	- volumes //文件映射目录
	-

- 部分示例
```
docker-compose -f docker-compose.yml -f docker-compose.local.yml up // 容器启动
docker-compose -f docker-compose.yml  down //容器停止
```
