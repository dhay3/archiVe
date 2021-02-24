# docker-cp

用于在宿主机和容器之间拷贝文件

文件生成规则参考：

https://docs.docker.com/engine/reference/commandline/cp/

## 容器到宿主机

pattern：`docker cp <containername:><container-path> <host-path>`

```
root in /opt λ docker cp t1:/etc/resolv.conf /opt
root in /opt λ ls
alibabacloud  containerd  Dockerfile  lsd-0.18.0-x86_64-unknown-linux-gnu  resolv.conf  tput_t.sh
root in /opt λ cat resolv.conf
```

## 宿主机到容器

pattern：`docker cp <host-path> <containername:><container-path>`

buxybox中`/opt`目录

```
root in /opt λ docker cp /opt/Dockfile t1:/opt
root in /opt λ docker  exec t1 cat /opt
ARG version=:latest
FROM centos$latest
RUN useradd cpl
USER cpl
RUN mkdir /myData
RUN echo "hello world" > /myData/greeting

```

