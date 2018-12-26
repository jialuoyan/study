## 普通前端项目转docker容器化
 公司项目在rancher 上部署的项目，本地想运行起来，尝试后，成功

前提：
- 项目使用node 编译，如果本地不需要编译，可以直接执行第二步
- 代码已经开发完毕

### 目录结构
```
keith@Keith:~/web/java/docker/cms-frontend$ tree -L 2
.
├── Dockerfile   //增加文件
├── README.md
├── build
│   ├── build.js
│   ├── check-versions.js
│   ├── config.js
│   ├── dev-client.js
│   ├── dev-server.js
│   ├── env.js
│   ├── utils.js
│   ├── vue-loader.conf.js
│   ├── webpack.base.conf.js
│   ├── webpack.dev.conf.js
│   └── webpack.prod.conf.js
├── conf
├── config
│   ├── dev.env.js
│   ├── index.js
│   └── prod.env.js
├── default.conf    //增加文件
├── dist   //打包生成文件
│   ├── index.html
│   └── static
├── docker-compose.yml  //增加文件
├── index.html
├── project.yml
```

### 步骤
此处 docker 安装步骤省略

#### 第一步 编译前端代码
- 下载node镜像
```
docker pull node
```
- 编译代码
```
<!--进入目录-->
cd  /usr/share/src/
<!-- 在项目目录下执行，将本地的项目文件，挂载到 容器mynode 的 /www 目录下 -->
docker run -t -i --name mynode -v $PWD:/www  node
<!-- 如果报错提示 mynode 已经存在则使用 `docker start mynode`  启动该容器 -->
<!-- 新开一个窗口 查看容器是否启动 -->
docker ps 
<!-- 进入mynode 容器 -->
docker exec -ti `docker ps | grep mynode | awk '{print $1}'` bash
<!--执行node 编译命令 -->
yarn && npm run build
<!-- 退出容器 -->
exit
编译代码完成(当前目录下已经存在 dist 目录)
docker ps -a  //查看当前pc所有容器（包含未启动）
docker kill mynode  //关闭容器
docker rm mynode  //移除容器
```
####  启动容器
```
<!-- 生成镜像 -->
docker build  -f Dockerfile -t registry.***.com/ued/frontend:dev-0.0.2 . // docker-compose 需要
<!-- 启动镜像-->
docker-compose -f docker-compose.yml  up
```

### 项目中需要增加的文件
由于前后端完全分离，所以考虑结局跨域问题最方便的情况下，直接用nginx 做反向代理，如果前后端在同一服务器则不需要这么麻烦
default.conf
```
server {
    listen       80;
    server_name  -;
    #access_log  /var/log/nginx/host.access.log  main;
    set $gateway "172.22.31.222:8090";

    location ~ ^/api {
        resolver 8.8.8.8;
        #后端访问请求地址
        if ($host ~ "test") {
             set $gateway "172.22.**.** ";
        }

        if ($host ~ "demo") {
             set $gateway "172.16.*.* ";
        }

        if ($host ~ "online") {
             set $gateway "";
        }

        proxy_http_version 1.1;
        proxy_set_header Connection "keep-alive";
        proxy_set_header x-real-ip $remote_addr;
        proxy_set_header x-forwarded-for $proxy_add_x_forwarded_for;
        proxy_next_upstream http_502 http_503 http_504 http_500 error timeout invalid_header;
        proxy_pass http://${gateway};
    }

    location / {
        alias   /usr/share/nginx/html/dist/;
        index  index.html;
        try_files $uri $uri/ /index.html;
    }
}
```
docker-compose.yml
```
version: '2'
services:
  frontend:
    mem_limit: 1073741824
    image: registry.***.com/ued/frontend:dev-0.0.2 //使用上面生成的
    environment:
      DEBUG_MODE: '1'  //配置的容器的环境变量，容器启动后 env 可以查到
    ports: 
    - 8124:80/tcp  //暴露的端口，本地 8124 映射容器80端口
```

Dockerfile
```
FROM nginx
ADD ./ /usr/share/nginx/html
ADD default.conf /etc/nginx/conf.d/default.conf
```

