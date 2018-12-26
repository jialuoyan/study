## 普通Java-maven项目docker 容器化

前提：
- 项目使用maven 打包, pom.xml 文件必须在主目录下
- 可以直接用jar 命令运行起来

目录结构
```
keith@Keith:~/web/java/docker/cms$ tree -L 3
.
├── Dockerfile  //增加文件
├── README.md
├── bootstrap
│   ├── pom.xml
│   ├── src
│   │   └── main
│   └── target
│       ├── bootstrap-0.0.1-SNAPSHOT.jar
│       ├── bootstrap-0.0.1-SNAPSHOT.jar.original
│       ├── classes
│       ├── generated-sources
│       ├── maven-archiver
│       └── maven-status
├── checkstyle-google.xml
├── common
│   ├── pom.xml
│   ├── src
│   │   └── main
│   └── target
│       ├── classes
│       ├── common-0.0.1-SNAPSHOT.jar
│       ├── generated-sources
│       ├── maven-archiver
│       └── maven-status
├── core
│   ├── pom.xml
│   ├── src
│   │   └── main
│   └── target
│       ├── classes
│       ├── core-0.0.1-SNAPSHOT.jar
│       ├── generated-sources
│       ├── maven-archiver
│       └── maven-status
├── docker-compose.yml  //增加文件
├── entrypoint.sh    //增加文件
├── etc
│   ├── pom.xml
│   └── src
├── pom.xml
```

### 执行步骤
本地安装maven 步骤省略
`cd 项目目录`
`mvn clean install` //生成可执行文件
`docker build  -f Dockerfile -t registry.***.com/java/frontend:dev-0.0.2 .` // docker-compose 需要


### 需要增加的文件
docker-compose.yml
```
version: "2"
services:
  cms:
    image: registry.***.com/java/frontend:dev-0.0.2
    environment:
      PROFILE__ACTIVE__ENV: ${PROFILE__ACTIVE__ENV} //设置环境变量
    cap_add:
    - SYS_PTRACE
    mem_limit: 2147483648

```

Dockerfile
```
FROM openjdk:8
MAINTAINER print23 <print23@126.com>

# 调整时区设置
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone

# 创建工作目录
RUN mkdir -p /usr/local/cms

# 运行生成的jar程序
#ENTRYPOINT ["java", "-jar", "/usr/local/cms/bootstrap-0.0.1-SNAPSHOT.jar"]

# 容器中服务使用的端口号,和yml文件中server.port保持一致
EXPOSE 80

# 将生成的可执行的jar文件添加到镜像当中
ADD bootstrap/target/bootstrap-0.0.1-SNAPSHOT.jar /usr/local/cms/bootstrap-0.0.1-SNAPSHOT.jar
ADD entrypoint.sh /usr/local/cms/entrypoint.sh

WORKDIR /usr/local/cms

ENTRYPOINT ["/bin/sh"]

CMD ["/usr/local/cms/entrypoint.sh"]

```

entrypoint.sh
//根据传入的环境变量，选择具体的执行环境
```
#!/bin/sh

profileenv=${PROFILE__ACTIVE__ENV}

if [ ! -n "$profileenv" ];  then
   	profileenv="dev"
fi

java -jar /usr/local/cms/bootstrap-0.0.1-SNAPSHOT.jar --spring.profiles.active=$profileenv
```