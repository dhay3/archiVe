# docker-ps

## 概述

用于展示容器

## 参数

- `-a`

  展示所有，不带`a`参数默认展示运行中的容器

- `--no-trunc`

  显示完整的containerid

  ```
  root in ~ λ docker ps --no-trunc
  CONTAINER ID                                                       IMAGE     COMMAND       CREATED      STATUS          PORTS     NAMES
  d17f6226d6b1539a9adf2434cb9b63851972536dbbb7485c6c57e5134a42cb61   centos    "/bin/bash"   2 days ago   Up 11 minutes   80/tcp    t3        
  ```

- `-q`

  只显示容器的id

  ```
  root in ~ λ docker images -q
  b97242f89c8a
  f6d0b4767a6c
  300e315adb2f      
  ```

- `-s`

  显示容器的占用磁盘的大小和镜像大小

  ```
  root in ~ λ docker ps -s
  CONTAINER ID   IMAGE     COMMAND       CREATED      STATUS          PORTS     NAMES     SIZE
  d17f6226d6b1   centos    "/bin/bash"   2 days ago   Up 13 minutes   80/tcp    t3        435kB (virtual 210MB)   
  ```

  container's writable layer on the host(images size)

- `--filter`

  1. name支持模糊匹配

  ```
  $ docker ps --filter "name=nostalgic"
  
  CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
  715ebfcee040        busybox             "top"               3 seconds ago       Up 1 second                             i_am_nostalgic
  9b6247364a03        busybox             "top"               7 minutes ago       Up 7 minutes                            nostalgic_stallman
  673394ef1d4c        busybox             "top"               38 minutes ago      Up 38 minutes                           nostalgic_shockley
  ```

  2. exited

  ```
  $ docker ps -a --filter 'exited=0'
  
  CONTAINER ID        IMAGE             COMMAND                CREATED             STATUS                   PORTS                      NAMES
  ea09c3c82f6e        registry:latest   /srv/run.sh            2 weeks ago         Exited (0) 2 weeks ago   127.0.0.1:5000->5000/tcp   desperate_leakey
  106ea823fe4e        fedora:latest     /bin/sh -c 'bash -l'   2 weeks ago         Exited (0) 2 weeks ago                              determined_albattani
  48ee228c9464        fedora:20         bash                   2 weeks ago         Exited (0) 2 weeks ago         
  ```

- `--format`

  以go template显示指定的部分

  ```
  root in ~ λ docker ps --format "{{.ID}}: {{.Command}}"
  d17f6226d6b1: "/bin/bash"  
  ```

  