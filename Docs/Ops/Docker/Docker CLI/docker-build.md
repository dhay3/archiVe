# Docker build

参考：

https://docs.docker.com/engine/reference/builder/?spm=a2c4g.11186623.2.2.18b12b117B26pG

## 概述

用于从DockerFile构建镜像

pattern：`docker build [option] <path|url>`

path表示读取dockerfile的目录，url表示读取dockerfile的根路径

## 参数

- `-f <relative-path>`

  指定使用dockerfile的相对path的位置。如果是url，那么是相对是项目的根路径。

  ==如果没有指定改参数，默认使用根路径下的DockerFile==

  ```
  $ docker build -f /path/to/a/Dockerfile .
  ```

- `--squash  true|flase`

  生成的镜像会将多层layer变成一层layer。会生成两个镜像一个多层layer，一个一层layer

- `--build-arg <variable>=<name>`

  用于构建镜像时提供变量（==镜像构建完毕后就失效==），同dockerfile中的ARG指令一样

- `--compress true|false`

  生成镜像时，是否使用gzip压缩

- `--rm true|false`

  构建成功后删除中间环节的容器，使用`--rm false`表明不删除中间使用的镜像

- `-t`

  用于重命名镜像的名字

     A better approach is to provide a fully qualified and meaningful repository, name, and tag (where the tag in this context  means
     the qualifier after the ":"). In this example we build a JBoss image for the Fedora repository and give it the version 1.0:

      docker build -t fedora/jboss:1.0 .

  ==如果在Dockerfile中使用多行FROM，需要多个`-t`来指定==

- `--cpu-shares <wight>`

  设置生成镜像cpu使用的权重。根据`/sys/fs/cgroup/cpu/cpu.shares`中的值设定。

  ```
  docker build --cpu-shares 614 #1024的60%
  ```

   For example, consider a system with more than three cores. If you start one
           container {C0} with --cpu-shares 512 running one process, and another container
           {C1} with --cpu-shares 1024 running two processes, this can result in the following
           division of CPU shares:

                  PID    container    CPU    CPU share
                  100    {C0}         0      100% of CPU0
                  101    {C1}         1      100% of CPU1
                  102    {C1}         2      100% of CPU2

- `--cpuset-cpus`

  设置生成的layer使用哪个cpu。

   For example, if you have four memory nodes on your system (0-3), use --cpuset-mems 0,1 to ensure the processes  in  your  Docker container only use memory from the first two memory nodes.







