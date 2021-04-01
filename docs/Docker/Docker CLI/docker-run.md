# docker run 

> 1. 使用`--rm`参数，在容器退出后会自动删除创建的容器。
> 2. ==如果在启动启动容器时没有添加`-it`参数，容器停止后重新启动同样还是会退出。==可以将`docker start container`理解为一个指向`docker run container`的一个指针。
> 3. 想要检查`docker run`为什么不能启动在命令行后面使用`;echo $?`
> 4. 如果容器以`-it`参数运行，可以使用`ctrl+p+q` detach container
> 5. 当docker运行的容器中没有在运行的前台程序，容器就会退出

## 概述

pattern：`docker run [OPTIONS] IMAGE[:TAG|@DIGEST] [COMMAND] [ARG...]`

`docker run`用于定义容器该如何启动

1. foreground or daemon(detach)
2. container identification(name)
3. network settings
4. runtime constraints on CPU and memory

## Exit Status

https://docs.docker.com/engine/reference/run/#exit-status

使用分号命令可以检查容器为什么启动失败

```
$ docker run --foo busybox; echo $?
```

具体启动失败原因使用`docker logs`来查看

## Detached vs foreground

`docker run`默认使用foreground，如果想要以守护线程的形式运行容器使用`-d`参数，如果想要以交互模式启动docker使用`-it`参数（==分配一个tty，docker启动的容器必须要有一个前台运行的进程，否则就会退出==）

> 这里使用docker pull nginx的镜像，无需启动镜像后再启动nginx

```
root in /etc/docker λ docker run -d --name t2 -p 80:80  nginx

root in /etc/docker λ docker run  --name t3 -p 80:80  nginx
```

## Container identification

- `--name`

  用于指定启动的容器的name，如果不指定docker deamon会随机生成一个name

  ```
  root in /etc/docker λ docker run -d --name t2 -p 80:80  nginx
  
  root in /etc/docker λ docker ps
  CONTAINER ID   IMAGE     COMMAND                  CREATED         STATUS         PORTS                NAMES
  cdc2a2a18bec   nginx     "/docker-entrypoint.…"   5 minutes ago   Up 5 minutes   0.0.0.0:80->80/tcp   t2
  2d730fd0a5d7   nginx     "/docker-entrypoint.…"   7 minutes ago   Up 7 minutes   80/tcp               t1                
  ```

## Network settings

https://www.qikqiak.com/k8s-book/docs/7.Docker%E7%9A%84%E7%BD%91%E7%BB%9C%E6%A8%A1%E5%BC%8F.html

docker使用`--network`来指定生成的容器该怎么创建容器。可以使用的value有如下四个。

==在新版本中的docker不只有4个值，可以是自定义的。参考docker network。==

1. none

   生成的容器只有loopback信息，没有其他iface配置，需要手动配置才能联网。

2. host

   生成的容器和宿主机的iface配置相同(使用宿主机的 IP 和端口。但是，容器的其他方面，如文件系统、进程列表等还是和宿主机隔离的)，==如果对网络要求较高推荐采用这种方法，但是由于容器可以访问宿主的网络信息，所以不安全。==

   <img src="D:\asset\note\imgs\_docker\Snipaste_2021-02-23_10-06-14.png" style="zoom:80%;" />

   ```
   root in ~ λ docker run --network=host -itd --name t4 busybox
   61c0db02f00fc529bee34e3f886b736c8008bef8315737a5f88910516958bd03
   root in ~ λ docker exec -it t4 sh
   / # ip a
   1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue qlen 1000
       link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
       inet 127.0.0.1/8 scope host lo
          valid_lft forever preferred_lft forever
       inet6 ::1/128 scope host
          valid_lft forever preferred_lft forever
   2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel qlen 1000
       link/ether 00:16:3e:0a:be:8b brd ff:ff:ff:ff:ff:ff
       inet 172.19.124.44/20 brd 172.19.127.255 scope global dynamic eth0
          valid_lft 315292524sec preferred_lft 315292524sec
       inet6 fe80::216:3eff:fe0a:be8b/64 scope link
          valid_lft forever preferred_lft forever
   3: docker0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue
       link/ether 02:42:2e:49:7b:39 brd ff:ff:ff:ff:ff:ff
       inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
          valid_lft forever preferred_lft forever
       inet6 fe80::42:2eff:fe49:7b39/64 scope link
          valid_lft forever preferred_lft forever
   15: veth03c6e13@if14: <BROADCAST,MULTICAST,UP,LOWER_UP,M-DOWN> mtu 1500 qdisc noqueue master docker0
       link/ether b6:21:2d:4a:28:16 brd ff:ff:ff:ff:ff:ff
       inet6 fe80::b421:2dff:fe4a:2816/64 scope link
          valid_lft forever preferred_lft forever
   17: vethc1b39cc@if16: <BROADCAST,MULTICAST,UP,LOWER_UP,M-DOWN> mtu 1500 qdisc noqueue master docker0
       link/ether ee:2c:d2:0f:a0:d6 brd ff:ff:ff:ff:ff:ff
       inet6 fe80::ec2c:d2ff:fe0f:a0d6/64 scope link
          valid_lft forever preferred_lft forever
   
   ```

3. container

   与指定的容器共享网络(NIC)。在宿主机上只会生成一个vethxxx。t1与t2的网络信息相同。

   <img src="D:\asset\note\imgs\_docker\Snipaste_2021-02-23_10-11-21.png" style="zoom:80%;" />

   ```
   root in /opt/t λ docker run -itd --name t1 busybox
   e3abf7c34e4e709dbb21fe032eb7a24394a3250a2fcfe523a7e7202d780899d5
   
   root in /opt/t λ docker run -itd --name t2 --network=container:t1 busybox
   e5c7d27d5bf29d0972cec073a1745d4f17e0e13e8e43173dd8f3e12cacfad121
   
   root in /opt/t λ docker ps
   CONTAINER ID   IMAGE     COMMAND   CREATED          STATUS          PORTS     NAMES
   e5c7d27d5bf2   busybox   "sh"      3 seconds ago    Up 2 seconds              t2
   e3abf7c34e4e   busybox   "sh"      53 seconds ago   Up 52 seconds             t1
   
   root in ~ λ docker inspect -f '{{json .NetworkSettings.Networks}}' t1
   {"bridge":{"IPAMConfig":null,"Links":null,"Aliases":null,"NetworkID":"80b9b205cd4c1518c43d2cd74a975440dd78e5766b89c3d3eaf404f4c11bafc3","EndpointID":"f815bde6d256ec4eeb2602212dfd9db8a236369bffa9087d2031642aed865fd5","Gateway":"172.17.0.1","IPAddress":"172.17.0.2","IPPrefixLen":16,"IPv6Gateway":"","GlobalIPv6Address":"","GlobalIPv6PrefixLen":0,"MacAddress":"02:42:ac:11:00:02","DriverOpts":null}}
   root in ~ λ docker ps
   CONTAINER ID   IMAGE     COMMAND   CREATED              STATUS              PORTS     NAMES
   9f42296d6adf   busybox   "sh"      About a minute ago   Up About a minute             t3
   7a1e7cb41549   busybox   "sh"      2 minutes ago        Up 2 minutes                  t2
   2f2fb0f72cfa   busybox   "sh"      3 minutes ago        Up 3 minutes                  t1
   root in ~ λ docker inspect -f '{{json .NetworkSettings.Networks}}' t2
   {}
   #在NetworkMode中显示网络模式
    "NetworkMode": "container:2f2fb0f72cfaa3af75283012aec060e01f0348882356f2db1f419e5a91692890",
   ```

4. bridge 缺省值

   如果以这种模式通信，首先会宿主机上创建一个名为docker0 和一个名为vethxxx的虚拟NIC

   ```
   root in /etc/docker λ ip a
   1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
       link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
       inet 127.0.0.1/8 scope host lo
          valid_lft forever preferred_lft forever
       inet6 ::1/128 scope host
          valid_lft forever preferred_lft forever
   2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
       link/ether 00:16:3e:0a:be:8b brd ff:ff:ff:ff:ff:ff
       inet 172.19.124.44/20 brd 172.19.127.255 scope global dynamic eth0
          valid_lft 315152313sec preferred_lft 315152313sec
       inet6 fe80::216:3eff:fe0a:be8b/64 scope link
          valid_lft forever preferred_lft forever
   3: docker0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
       link/ether 02:42:4a:d4:97:1b brd ff:ff:ff:ff:ff:ff
       inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
          valid_lft forever preferred_lft forever
       inet6 fe80::42:4aff:fed4:971b/64 scope link
          valid_lft forever preferred_lft forever
   239: veth94eba64@if238: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default
       link/ether 52:85:9c:ec:15:18 brd ff:ff:ff:ff:ff:ff link-netnsid 0
       inet6 fe80::5085:9cff:feec:1518/64 scope link
          valid_lft forever preferred_lft forever
          
   root in /etc/docker λ docker run --name net2 -it --network=bridge centos
   [root@1197bf68bc69 /]# ip a
   1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
       link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
       inet 127.0.0.1/8 scope host lo
          valid_lft forever preferred_lft forever
   250: eth0@if251: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
       link/ether 02:42:ac:11:00:05 brd ff:ff:ff:ff:ff:ff link-netnsid 0
       inet 172.17.0.5/16 brd 172.17.255.255 scope global eth0
          valid_lft forever preferred_lft forever
   
   ```

   docker0作为容器的网关，vethxxx的另一端放在新建的容器中，以eth0命名。

   <img src="D:\asset\note\imgs\_docker\Snipaste_2021-01-22_19-50-21.png" style="zoom:80%;" />

   此时容器能与宿主机通信，同时也能访问外网。

## Restart policies

docker 使用`--restart=`指定容器退出后的策略(宿主机关机后容器也会退出)。==不能与`--rm`一起使用==

1. no 缺省值

   容器退出时默认不重启

2. on-failure[:max-retries]

   容器非正常退出，重启容器

3. always

   ==不管容器退出的状态都会重启容器==

4. unless-stopped

```
$ docker run --restart=always redis
$ docker run --restart=on-failure:10 redis
```

## Clean up

默认如果容器退出，容器不会从`docker ps`中删除容器，且匿名挂载到宿主机上的文件不会被清除。使用`--rm`在容器退出后清空匿名挂载文件和`docker ps`中的对应的容器。==与`--restart`冲突==

```
root in /etc/docker λ docker run --rm --name net3 centos
```

## Volume

https://docs.docker.com/engine/reference/commandline/run/#mount-volume--v---read-only

使用`-v`参数将==宿主机上的文件映射到容器(会覆盖)==，可以是目录也可以是文件

pattern：`docker run -v [host-src:]container-dest`

如果没有指定host-src，docker自动生成相应的挂载卷(通过`root in ~ λ docker inspect t1 --format="{{json .Mounts}}"`可以查看)。默认挂载的卷使用rw权限，可以添加后缀`:ro`或`:rw`指定权限

```
$ sudo docker run \
  --volume=/:/rootfs:ro \
  --volume=/var/run:/var/run:rw \
  --volume=/sys:/sys:ro \
  --volume=/var/lib/docker/:/var/lib/docker:ro \
  --publish=8080:8080 \
  --detach=true \
  --name=cadvisor \
  google/cadvisor:latest
```

## Workdir

docker 使用`-w`参数来指定进入容器后的工作目录

```
root in ~ λ docker run -itd -w="/etc" --name t3 centos
06dcdb3c3c9c4e008a1b760f1ccf5261b199760964b0fd29ba68ddd1c1f59fd8                                                     /0.6s
root in ~ λ docker exec -it t3 /bin/bash
[root@06dcdb3c3c9c etc]# pwd
/etc
```

## Expose

- `--expose`

  将容器指定的端口暴露给外界，动态绑定

  ```
  root in /etc/docker λ docker run --name t2 -itd --expose=80-3306 centos
  16bada980c206032b1c9008a5f7545180581952dd7ed9031b3da182398d7e6e7                                                                      /0.7s
  root in /etc/docker λ docker ps
  CONTAINER ID   IMAGE     COMMAND       CREATED         STATUS         PORTS         NAMES
  16bada980c20   centos    "/bin/bash"   3 seconds ago   Up 2 seconds   80-3306/tcp   t2
  7fcf4bec3932   centos    "/bin/bash"   4 minutes ago   Up 4 minutes                 t1   
  ```

- `-P`

  对外界暴露所有端口，动态绑定

- `-p [ip:]<host_port>:<container_port>[/proto]`

  将指定的端口对外界暴露，静态绑定，我们可以通过宿主机访问容器。

  可以使用范围表示，但是需要一一对应

  ```
   -p 1234-1236:1234-1236/tcp
  ```

  宿主机可以用范围来表示，容器的一个端口

  ```
  -p 1234-1236:1234/tcp
  ```

## ENV

在创建容器时Docker可以通过`-e`参数设置一些自定义的参数

```
root in ~ λ docker run -e "deep=purple" --name t3 --rm -it centos
[root@660a4357c46c /]# echo $deep
purple
```

## HealthCheck

- `--health-cmd`

  指定检查健康的脚本

- `--health-interval`

  脚本运行的周期

```
$ docker run --name=test -d \
    --health-cmd='stat /etc/passwd || exit 1' \
    --health-interval=2s \
    busybox sleep 1d
```

==由于docker的机制容器需要有前台进程运行才能正常启动容器，所以指定`sleep 1d`==

## User

docker通过`-u`参数指定进入容器后的用户，默认为root

pattern：`-u=[ user | user:group | uid | uid:gid | user:gid | uid:group ]`

```
root in ~ λ docker run -it --name t1 centos
[root@7bc79a71cf37 /]# uid
bash: uid: command not found
[root@7bc79a71cf37 /]# id
uid=0(root) gid=0(root) groups=0(root)
[root@7bc79a71cf37 /]# whoami
root
[root@7bc79a71cf37 /]# exit
exit                                                                                                                /12.8s
root in ~ λ docker run -u="1:1" -it --name t2 centos
bash-4.4$ whoami
bin
bash-4.4$ id
uid=1(bin) gid=1(bin) groups=1(bin)
bash-4.4$
```

## Runtime constraints on resources

https://docs.docker.com/engine/reference/run/#runtime-constraints-on-resources







