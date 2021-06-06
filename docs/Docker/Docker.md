[TOC]

## Docker 三大组件

- ==**Images**==(镜像)：

  Docker 镜像（Image），就相当于是一个 root 文件系统。比如官方镜像 ubuntu:16.04 就包含了完整的一套 Ubuntu16.04 最小系统的 root 文件系统

  DockerFile中通过`FROM`，`RUN`，`CMD`等命令会创建一层[layer](https://stackoverflow.com/questions/31222377/what-are-docker-image-layers)

-  ==**Containers**==(容器)：

  镜像（Image）和容器（Container）的关系，就像是面向对象程序设计中的类和实例一样，镜像是静态的定义，容器是镜像运行时的实体。容器可以被创建、启动、停止、删除、暂停等。

- ==**Repository**==(仓库)：

  仓库可看成一个代码控制中心，用来保存镜像。

![Snipaste_2020-08-19_18-46-53](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-08-19_18-46-53.2jdsoogrvwo0.png)

## 安装

通过`whereis docker`查看docker的安装目录

[官网](https://docs.docker.com/install/linux/docker-ce/centos/)

1. 按照官网步骤安装

   注意国内需要使用阿里云的docker镜像否则安装docker非常慢

   `sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo`

   如果出现无法安装

   `yum clean all && yum makecache`

2. 如果无法运行

   ```
   docker run hello-world
   ```

   配置镜像加速

   ```json
   sudo mkdir -p /etc/docker
   sudo tee /etc/docker/daemon.json <<-'EOF'
   {
     "registry-mirrors": ["https://u6p3j4k6.mirror.aliyuncs.com"]
   }
   EOF
   sudo systemctl daemon-reload
   sudo systemctl restart docker
   ```

3. 重启docker

   ```
   systemctl restart docker //重启
   systemctl start docker  //启动
   ```

4. 验证

   ```
   docker run hello-world     
   ```

5. 设置开机自启

   ```
   systemctl enable docker
   ```

6. [shell completion](https://docs.docker.com/compose/completion/)

### 镜像命令

https://www.runoob.com/docker/docker-tutorial.html

镜像下载位置

`usr/local/docker`

- docker --help

- docker  -v
  显示版本
- docker info 
  查看docker 具体信息

- docker images 

  列出本地镜像

- docker search [IMAGENAME]
  搜索dockerHub上的资源

- docker pull [IMAGENAME]:[TAG]
  默认拉取最新的，如果需要指定版本，例如`docker pull mysql:5.7`

- docker rmi (-f) [IMAGENAME]:[latest]
  删除镜像
  -f 强制删除
  使用`docker rmi -f $(docker images -qa)` 删除所有镜像，$()复合命令

### 容器命令

==容器其实就是一个精简的linux系统==

![tomcat](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/tomcat.7l7hj0fgy1o0.PNG)

==容器重启后不会保留数据==，如需要保留数据使用容器卷

如果返回container id，表示操作成功

<font style='color:red'>如果没用和容器交互docker会自动关闭容器，可以通过ctrl+p+q挂起当前容器</font>

- docker run [option] IMAGE:[TAG]
  运行镜像，生成容器实例
  -it ：容器终端交互
  `docker run -it centos /bin/bash` (/bin/bash可以省略)
  -d ：以守护线程的模式运行
  --name： 给运行的容器起一个别名
  -p：指定外部端口对应容器中的端口

  -v: 挂载容器卷

- docker attach CONTAINERID
  进入容器内部， 如果以这种方式退出容器，容器就会自动关闭

- ==docker exec -it CONTAINERID /bin/bash==

  `docker exec -it redis /bin/bash redis-cli`容器运行成功后，容器运行`redis-cli`命令

  推荐使用

  进入容器内部，如果以这种方式退出容器，==容器不会自动关闭==

- docker ps [option]
  查看正在运行的容器
  -a： 列出所有正在运行的和历史上运行过的
  -l： 显示最近运行的容器
  -n ：显示最近n个运行过的容器
  -q ：静默模式,只显示正在运行的容器编号

  `docker ps CONTAINERID`查看具体的容器

- docker start CONTAINERID
  启动container

- docker restart CONTAINERID
  重启container

- docker stop CONTAINERID
  正常停止container

- docker kill CONTAINERID
  强制停止container

- docker rm [option] CONTAINERID
  删除停止运行的container
  -f ：强制删除

  `docker rm -f $(docker ps -q) `删除所有运行的container

  `docker rm -f $(docker ps -qa) `删除所有container

- docker logs [option] CONTAINERID
  查看docker日志
  -t ：加入时间戳
  -f ：跟随最新的日志打印,不停追加日志
  --tail： 显示最后多少条
  `docker logs -tf --tail 3 1b69245f0f61` 不停追加显示倒数三行日志

- ==docker inspect== CONTAINERID
  查看容器内部细节(json串显示)

  ==注意新版的镜像挂载点的映射在Mounts中==

![Snipaste_2020-08-20_17-26-03](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-08-20_17-26-03.295jhpnnbc5c.png)

- docker cp [OPTIONS] container:src_path dest_path
  docker cp [OPTIONS] dest_path container:src_path 

  容器与宿主机之间文件拷贝

  ```shell
  将主机/www/runoob目录拷贝到容器96f7f14e99ab的/www目录下。
  docker cp /www/runoob 96f7f14e99ab:/www/
  
  将主机/www/runoob目录拷贝到容器96f7f14e99ab中，目录重命名为www。
  docker cp /www/runoob 96f7f14e99ab:/www
  
  将容器96f7f14e99ab的/www目录拷贝到主机的/tmp目录中。
  docker cp  96f7f14e99ab:/www /tmp/
  ```

- docker history CONTAINTERID

  查看镜像变更记录

## 自定义镜像commit

==注意容器不能退出==，否则容器会还原，需要使用`ctrl+p+q`挂起容器

- docker commit -m"" -a"" CONTAINERID [name]:[version]
  提交修改过的docker image到本地仓库
  -m ：消息message
  -a ：作者author
  name： 自定义新容器的image名字
  version: 自定义新容器的version

![commit](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/commit.5bmejq6ok3c0.PNG)

**运行自定义的镜像要指定版本**

![commit2](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/commit2.1dt8q6dxghpc.PNG)

## 数据容器卷

### 数据卷

将容器中的指定文件挂载到宿主机

==容器与宿主机数据共享，容器重启后，宿主机中的数据同样会与容器共享==

```
docker run -it -v [HOST PATH]：[COTAINER PATH]  [CONTAINERID]
-v volume 创建数据容器卷
docker run -it -v /mydata:/data:ro [CONTAINERID]表示容器内的文件只允许read
```

通过`docekr inspect CONTAINERID`查看是否挂载

### 容器卷

![Snipaste_2020-08-20_17-52-33](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-08-20_17-52-33.gdpt3kijy5k.png)

```
docker run -it --volumes-form [parentcontainerid] --name [alis]
```

==只要是同一个镜像创建的容器通过--volumes-form, 容器之间数据共享==，类似于Maven的组合

==只要容器没有关闭完，即使一个容器退出，并不影响容器间数据共享==



