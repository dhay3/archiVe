

# Dockerfile

参考：

https://docs.docker.com/engine/reference/builder/#arg

http://c.biancheng.net/view/3162.html

## 概述

`docker build`通过Dockerfile创建镜像，默认以文件Dockerfile来构建

**dockerfile**

```dockerfile
#基于centos镜像
FROM centos

#维护人的信息
MAINTAINER The CentOS Project <303323496@qq.com>

#安装httpd软件包
RUN yum -y update
RUN yum -y install httpd

#开启80端口
EXPOSE 80

#复制网站首页文件至镜像中web站点下
ADD index.html /var/www/html/index.html

#复制该脚本至镜像中，并修改其权限
ADD run.sh /run.sh
RUN chmod 775 /run.sh

#当启动容器时执行的脚本文件
CMD ["/run.sh"]
```

**docker build**

```
root in /opt λ docker build -t centos:test .
Sending build context to Docker daemon   2.55MB
Step 1/3 : FROM centos
 ---> 300e315adb2f
Step 2/3 : RUN echo hello-world
 ---> Using cache
 ---> 6f4446b56a1c
Step 3/3 : CMD ["/bin/bash"]
 ---> Using cache
 ---> b0efe6ee149e
Successfully built b0efe6ee149e
Successfully tagged centos:test                                                                                                       /0.2s
root in /opt λ docker images
REPOSITORY    TAG       IMAGE ID       CREATED         SIZE
centos        test      b0efe6ee149e   6 minutes ago   209MB
centos        latest    300e315adb2f   6 weeks ago     209MB
hello-world   latest    bf756fb1ae65   12 months ago   13.3kB       
```

## 格式

### comment

在 Dockerfile中通过`#`来标记注释

### parser directives

patter：`# directive=value1`

escape指令块必须在Dockerfile的开头位置

#### escape

https://docs.docker.com/engine/reference/builder/#escape

转译符指令块，==对windows的镜像非常有用==。如果为指定使用`\`

```
# escape=` (backtick)
FROM microsoft/nanoserver
COPY testfile.txt c:\\
RUN dir c:\

PS C:\John> docker build -t succeeds --no-cache=true .
Sending build context to Docker daemon 3.072 kB
Step 1/3 : FROM microsoft/nanoserver
 ---> 22738ff49c6d
Step 2/3 : COPY testfile.txt c:\
 ---> 96655de338de
Removing intermediate container 4db9acbb1682
Step 3/3 : RUN dir c:\
 ---> Running in a2c157f842f5
 Volume in drive C has no label.
 Volume Serial Number is 7E6D-E0F7
```

如果为使用`escape`指令块

```
PS C:\John> docker build -t cmd .
Sending build context to Docker daemon 3.072 kB
Step 1/2 : FROM microsoft/nanoserver
 ---> 22738ff49c6d
Step 2/2 : COPY testfile.txt c:\RUN dir c:
GetFileAttributesEx c:RUN: The system cannot find the file specified.
PS C:\John>
```

### Instruction

![450977-20190512115951746-136143052](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/450977-20190512115951746-136143052.4nhk2tsxric0.png)

==指令不区别大小写，但是为了方便便是使用 大写，并且严格区分单双引号(只能使用双引号)==

> ENV于ARG的区别与使用
>
> https://blog.justwe.site/post/docker-arg-env/

#### ENV

Dockerfile中使用的环境变量（==构建镜像过程中和之后都生效==），通过`$`或是`${}`来获取值。

只在如下指令中有效

1. `ADD`
2. `COPY`
3. `ENV`
4. `EXPOSE`
5. `FROM`
6. `LABEL`
7. `STOPSIGNAL`
8. `USER`
9. `VOLUME`
10. `WORKDIR`
11. `ONBUILD`

```
FROM busybox
ENV FOO=/bar
WORKDIR ${FOO}   # WORKDIR /bar
ADD . $FOO       # ADD . /bar
COPY \$FOO /quux # COPY $FOO /quux
```

#### **ARG**

和ENV类似，是在==镜像构建过程中==使用的。可以用于所有指令中。

```
FROM nginx:1.13.1-alpine

LABEL maintainer="GPF <5173180@qq.com>"

#https://yeasy.gitbooks.io/docker_practice/content/image/build.html
RUN mkdir -p /etc/nginx/cert \
    && mkdir -p /etc/nginx/conf.d \
    && mkdir -p /etc/nginx/sites

COPY ./nginx.conf /etc/ngixn/nginx.conf
COPY ./conf.d/ /etc/nginx/conf.d/
COPY ./cert/ /etc/nginx/cert/

COPY ./sites /etc/nginx/sites/


ARG PHP_UPSTREAM_CONTAINER=php-fpm
ARG PHP_UPSTREAM_PORT=9000
RUN echo "upstream php-upstream { server ${PHP_UPSTREAM_CONTAINER}:${PHP_UPSTREAM_PORT}; }" > /etc/nginx/conf.d/upstream.conf

VOLUME ["/var/log/nginx", "/var/www"]

WORKDIR /usr/share/nginx/html

```

==RUN指令属于构建构过程，CMD属于构建之后的指令==

#### **FROM**

用于指定构建的镜像的来源。可以在同一个Dockerfile中有多行FROM，来生成不同的镜像。构建是通过`docker build -t <name1> -t <name2> .`来创建。如果没有指定版本，默认使用latest

```
ARG  CODE_VERSION=latest
FROM base:${CODE_VERSION}
CMD  /code/run-app

FROM extras:${CODE_VERSION}
CMD  /code/run-extras
```

#### RUN

RUN指令用于在当前构建的==镜像中==执行命令

pattern：

1. `RUN <command>` shell form

   by default is `/bin/sh -c` on Linux or `cmd /S /C` on Windows

2. `RUN ["executable", "param1", "param2"]` exec form

   如果使用这种形式，如果需要获取环境变量就必须指明使用的shell，否则会报错

   ```
   RUN [ "sh", "-c", "echo $HOME" ]
   ```

> CMD和ENTRYPOINT用于替代在`docker run [command]`中的command
>
> 运行一个没有调用ENTRYPOINT或者CMD的docker镜像, 一定返回错误。
>
> 可以通过这两个参数指定运行容器后执行的命令(通常开启一个用于交互的shell)
>
> `CMD ["/bin/bash"]`(如果你以一个Linux ditro为基础镜像就不需要指定，因为大多数构建的镜像中已执行该命令)
>
> 注意如果以shell form PID 为 1的进程是shell(pid为1的进程不会接受到SIGINT)，==如果使用shell form可能会导致容器不能正常退出，所以统一使用exec form来创建Dockerfile==
>
> **如果以exec form执行注意一下后面跟的是不是command string，如果是必须添加`-c`参数**

#### CMD

https://www.cnblogs.com/sparkdev/p/8461576.html

CMD指令用于==镜像构建完成后（容器启动后）==执行命令。

如果Dokcerfile中有多条CMD指令，只有最后一个CMD指令会执行。

如果以pattern2的形式使用CMD，那么CMD和ENTRYPOTION指令都必须是以JSON数组的形式存在。

pattern：

1. `CMD ["executable","param1","param2"]` (*exec* form, this is the preferred form)

   与RUN一样如果需要使用环境变量需要指明shell。

   通过这种方式运行容器，1号进程就是CMD中指定的。所以如果以交互模式进入容器(`-it`)，==不会在stdout中显示（因为先执行命令，然后创建shell）==。

   ```
   root in /opt λ cat Dockerfile
   ARG version=:latest
   FROM centos$latest
   CMD ["top"]
   
   $ docker build -t test1 .
   $ docker run -idt --name test11 test1
   
   root in /opt λ docker exec test11 ps aux
   USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
   root         1  0.0  0.0  49072  3824 pts/0    Ss+  09:13   0:00 top
   root         9  0.0  0.0  44632  3448 ?        Rs   09:15   0:00 ps aux  
   ```

2. `CMD ["param1","param2"]` (as *default parameters to ENTRYPOINT*)

3. `CMD command param1 param2` (*shell* form)

   与exec form不同的是，==启动容器后的1号进程是sh==。

   如果以交互模式进入容器，那么运行的就是CMD指定命令，==会顶替`docker run -it`创建的shell==

该指令会被`docker run [command]`中的command覆盖。

```
FROM tomcat
...
CMD ["catalina.sh","run"]
--------
$ docker run -it tomcat ls -l
```

#### ENTRYPOINT

与CMD类似，但是不会被`docker run [command]`覆盖，后一个ENTRYPOINT会覆盖前一个ENTRYPOINT。可以使用`--entrypotion`来覆盖Dockerfile内的ENTRYPOINT。

pattern：

1. `ENTRYPOINT ["executable", "param1", "param2"]`(*exec* form)

   exec form会在每一个ENTRYPOTION后拼接command

   如果CMD与ENTRYPOINT(exec form)同时存在，ENTRYPOINT会将CMD拼接在后面

   https://zhuanlan.zhihu.com/p/30555962

   ```
   FROM ubuntu:trusty
   ENTRYPOINT ["/bin/ping","-c","3"]
   CMD ["localhost"] 
   
   $ docker build -t ping .
   [truncated]
   
   $ docker run ping
   PING localhost (127.0.0.1) 56(84) bytes of data.
   64 bytes from localhost (127.0.0.1): icmp_seq=1 ttl=64 time=0.025 ms
   64 bytes from localhost (127.0.0.1): icmp_seq=2 ttl=64 time=0.038 ms
   64 bytes from localhost (127.0.0.1): icmp_seq=3 ttl=64 time=0.051 ms
   
   --- localhost ping statistics ---
   3 packets transmitted, 3 received, 0% packet loss, time 1999ms
   rtt min/avg/max/mdev = 0.025/0.038/0.051/0.010 ms
   
   $ docker ps -l
   CONTAINER ID IMAGE COMMAND CREATED
   82df66a2a9f1 ping:latest "/bin/ping -c 3 localhost" 6 seconds ago 
   ```

2. `ENTRYPOINT command param1 param2`(The *shell* form)

   使用shell form 会==忽略CMD和docker run 中的命令==

   ```
   root in /opt λ docker run -it --name test2 test top
   test                                                                                                                                  /0.8s
   root in /opt λ cat Dockerfile
   ARG version=:latest
   FROM centos$latest
   ENTRYPOINT echo test
   CMD echo "cmd test"
   ```

   > 为了能正常退出ENTRYPOINT创建的长命令需要在命令前使用`exec`
   >
   > `ENTRYPOINT exec top -b`

#### LABEL

用于描述镜像metadata，以键值对的形式表示

```
LABEL "com.example.vendor"="ACME Incorporated"
LABEL com.example.label-with-value="foo"
LABEL version="1.0"

---

LABEL multi.label1="value1" \
      multi.label2="value2" \
      other="value3" #使用backslash拼接
```

可以使用`docker image inspect <imagename>`来查看, 会以JSON的形式显示

```
{
  "com.example.vendor": "ACME Incorporated",
  "com.example.label-with-value": "foo",
  "version": "1.0",
  "description": "This text illustrates that label-values can span multiple lines.",
  "multi.label1": "value1",
  "multi.label2": "value2",
  "other": "value3"
}
```

#### EXPOSE

运行容器提供的服务端口，但是在运行容器时并不会应为声明就会开启这个端口的服务。还是需要`docker run -p <host_port>:<container_port>`来指定。该指令只是为了方便管理。

```
EXPOSE 80/udp
EXPOSE 80/udp
```

#### ADD

用于将remote url 或是本地(不是镜像，==而是宿主机==)的文件拷贝到新镜像中。

pattern：

1. `ADD [--chown=<user>:<group>] <src>... <dest>`
2. `ADD [--chown=<user>:<group>] ["<src>",... "<dest>"]`

支持`*`和`?`通配符。必须以运行`docker build`的目录做为src的目录。

```dockerfile
ADD hom* /mydir/
ADD hom?.txt /mydir/
ADD https://xxx.com/html.tar.gz /var/www/html
```

`--chown `只针对linux中有效，表示dest生成的文件所有者

```
ADD --chown=55:mygroup files* /somedir/
ADD --chown=bin files* /somedir/
ADD --chown=1 files* /somedir/
ADD --chown=10:11 files* /somedir/
```

生成目录规则

1. 如果是压缩文件，会自动解压

2. 如果url以back slash 结尾，默认会创建一个文件

   如`ADD https://xxx.com/user/ /var/`将会在var目录下创建一个user目录

3. 如果dest会存在，会自动创建

4. 如果src是一个文件，只会将文件中的内容拷贝到dest，目录本身不会被拷贝到dest。

#### COPY

与ADD使用规则相同，但是不能自动解压缩文件

#### VOLUME

==将容器的目录挂载到宿主机，宿主机上的目录自动生成。==

pattern：

1. `VOLUME ["/data"]`
2. `VOLUME /myvol`

可以通过`docker inspect <contianer_id>`来查看自动的生成的目录位置

```
root in ~ λ docker ps
CONTAINER ID   IMAGE     COMMAND       CREATED       STATUS       PORTS     NAMES
05950b8c4220   test1     "/bin/bash"   2 hours ago   Up 2 hours             test11                                                    /0.1s
root in ~ λ docker inspect test11
[
		...
        "Mounts": [
            {
                "Type": "volume",
                "Name": "e266cda9aaef84862674ff49822e6d6da59fd62e2c04a02f9e579a62cc3f4de2",
                "Source": "/var/lib/docker/volumes/e266cda9aaef84862674ff49822e6d6da59fd62e2c04a02f9e579a62cc3f4de2/_data",
                "Destination": "/myData",
                "Driver": "local",
                "Mode": "",
                "RW": true,
                "Propagation": ""
            }
        ],
       ...
]                          
```

#### USER

指定接下来使用RUN，CMD，ENTRYPOINT的用户(使用前必须先创建用户)。如果未指定默认以root用户。

pattern：`USER <UID>[:<GID>]`

#### WORKDIR

用于指定RUN，CMD，ENTRYPOINT，CPOY和ADD执行时的目录。

```
ENV DIRPATH=/path
WORKDIR $DIRPATH/$DIRNAME
RUN pwd
```

#### ONBUILD

延迟构建命令。在本次构建中不会执行，但是如果有新的镜像以该镜像构建时运行。==任何指令都可以被用作延迟构建命令，但是不能以ONBUILD ONBUILD这种形式==

```
#father
FROM tomcat
...
ONBUILD echo "hello world" \
	   ADD . /app/src
$docker build -t father .
-------
#son 
FROM father #使用父容器
...
$docker build -t son #创建镜像时会触发父容器的ONBUILD
```

#### SHELL

针对`RUN`, `CMD` and `ENTRYPOINT`的shell form指定使用的shell

```
SHELL ["powershell", "-command"]
RUN Write-Host hello
```

