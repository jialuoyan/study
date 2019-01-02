## Dockerfile 语法
| 参数   |  含义  |  格式 |
|  --| --| --|  
| FROM |  指定基础的image | FROM image:tag|
| MAINTAINER |  指定镜像创建者信息 | |
| ENV | 设置环境变量 | |
| ADD | 将本地文件添加到镜像 |  ADD  源路径  目标路径  |
| RUN | 运行任何被基础image支持的命令 | |
| RUN | 容器启动时运行的操作。该指令只能在文件中存在一次，如果有多个，则只执行最后一条 | |
| ENTRYPOINT | 设置容器启动时执行的操作。该指令只能在文件中存在一次，如果有多个，则只执行最后一条 | |
| EXPOSE | 指定容器需要暴露的端口 | |
| COPY | 指定容器需要暴露的端口 | |

注意事项：
- ADD 命令： 如果源文件是个归档文件（压缩文件），则docker会自动帮解压，目标路径如果不存在，会自动创建路径
- COPY/ADD 命令： 推荐使用 exec格式用法，如下文例子中json 格式的exec

区别：
- ADD 命令 与COPY 命令区别
    - COPY指令和ADD指令功能和使用方式类似。只是COPY指令不会做自动解压工作
    - ADD命令可以从网络下载，但COPY不行
- RUN CMD 和的 ENTRYPOINT 区别
    - RUN命令执行命令并创建新的镜像层，通常用于安装软件包
    - CMD命令设置容器启动后默认执行的命令及其参数，但CMD设置的命令能够被`docker run`命令后面的命令行参数**替换**，一般用于设置默认命令
    - ENTRYPOINT配置容器启动时的执行命令（docker run 会执行该命令，除非在run 时指定了 `--entrypoint` 会覆盖该命令）

Shell格式和Exec格式运行命令(推荐使用exec格式):
- Shell格式：<instruction> <command>。例如：apt-get install python3
- Exec格式：<instruction> ["executable", "param1", "param2", ...]。例如： ["apt-get", "install", "python3"]



DockerFile参考：
```
FROM print23/php:7.1.17-base-r1
MAINTAINER print23 <print23@126.com>

ENV PROJECT_TYPE="G7-MS" SRC_DIR="/usr/src" PHP_VERSION="7.1.17" PRODUCT_DIR="/usr/local/product"
ENV PHP_INSTALL_DIR ${PRODUCT_DIR}/php-${PHP_VERSION}

# ADD . ${SRC_DIR}

# apk repository
RUN set -xe; \
    apk add --no-cache --virtual .build-deps \
    linux-headers \
    autoconf \
    dpkg-dev dpkg \
    file \
    g++ \
    gcc \
    pcre-dev \
    libc-dev \
    jpeg-dev \
    libpng-dev \
    gd-dev \
    openssl-dev \
    make \
    pkgconf \
    && mkdir -p ${SRC_DIR} \
    && cd ${SRC_DIR} \
    && wget -O LuaJIT-2.0.5.tar.gz http://luajit.org/download/LuaJIT-2.0.5.tar.gz \
    && tar zxf LuaJIT-2.0.5.tar.gz \
    && cd LuaJIT-2.0.5 \
    && make && make install \
    && cp -r /usr/local/include/luajit-2.0 /usr/include/ \
    && cd ${SRC_DIR} \
    && wget -O tengine-2.2.1.tar.gz http://tengine.taobao.org/download/tengine-2.2.1.tar.gz \
    && tar xzf tengine-2.2.1.tar.gz \
    && cd tengine-2.2.1 \
    && ./configure \
        --prefix=${PRODUCT_DIR}/tengine-2.2.1 \
        --pid-path=/var/run/nginx.pid \
        --with-file-aio \
        --with-http_concat_module \
        --with-http_ssl_module \
        --with-http_v2_module \
        --with-http_realip_module \
        --with-http_addition_module \
        --with-http_image_filter_module \
        --with-http_gunzip_module \
        --with-http_gzip_static_module \
        --with-http_auth_request_module \
        --with-http_stub_status_module \
        --with-pcre --without-http_ssi_module \
        --without-http_userid_module \
        --without-http_split_clients_module \
        --without-http_uwsgi_module \
        --without-http_memcached_module \
        --without-http_geo_module \
        --without-http_scgi_module \
        --without-mail_pop3_module \
        --without-mail_imap_module \
        --without-mail_smtp_module \
        --add-module=../ngx_devel_kit-0.3.0 \
        --with-http_lua_module \
        --with-luajit-lib=/usr/local/lib/ \
        --with-luajit-inc=/usr/local/include/luajit-2.0/ \
        --with-lua-inc=/usr/local/include/luajit-2.0/ \
        --with-lua-lib=/usr/local/lib/ \
        --with-ld-opt=-Wl,-rpath,/usr/local/lib \
        --add-module=../nginx-logid-1.0.0 \
        && make && make install \
        && ln -s /usr/local/product/tengine-2.2.1 /usr/local/nginx \
        && mkdir -p /usr/local/nginx/var/tmp \
       && runDeps="$( \
            scanelf --needed --nobanner --format '%n#p' --recursive /usr/local \
                | tr ',' '\n' \
                | sort -u \
                | awk 'system("[ -e /usr/local/lib/" $1 " ]") == 0 { next } { print "so:" $1 }' \
        )" \
        && apk add --no-cache --virtual .nginx-rundeps $runDeps \
        && apk del .build-deps \
        && rm -rf ${SRC_DIR}/*

COPY alpha.sh /alpha.sh

RUN cd ${SRC_DIR} \
    && apk add --no-cache bash logrotate tzdata

COPY --from=print23/php:7.1.17-ext-r1 [ \
"/usr/local/bin/composer", \
"${PHP_INSTALL_DIR}/bin/composer" \
]

COPY run.sh /run.sh

ENTRYPOINT ["/run.sh"]
```