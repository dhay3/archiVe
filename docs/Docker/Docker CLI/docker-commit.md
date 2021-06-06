# docker-commit

## 概述

docker commit用于从现有的container生成新的镜像。和docker tag不同，生成的imageId与生成容器的镜像的imageId不同。

pattern：`docker commit <containerId> <new-imagename>`

```
root in ~ λ docker commit 27ad5f06766a busybox:cpl
sha256:32e75ece0752ded3d15cfa0777d75de84d627392b13834616fbf52a20167fc41                                                                     /0.2s
root in ~ λ docker images
REPOSITORY                                                        TAG          IMAGE ID       CREATED         SIZE
busybox                                                           cpl          32e75ece0752   4 seconds ago   1.23MB
registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub   version1.0   491198851f0c   3 days ago      1.23MB
busybox                                                           latest       491198851f0c   3 days ago      1.23MB
byc                                                               latest       491198851f0c   3 days ago      1.23MB
nginx                                                             latest       f6d0b4767a6c   5 weeks ago     133MB
centos                                                            latest       300e315adb2f   2 months ago    209MB  
```

## 参数

- `-a`

  指定镜像的作者

  ```
  root in ~ λ docker commit -a 'cyberpelican' c082c5068def bcpl
  sha256:bec4e0e815ac7de870bb5cec4151c6cafe2c68bd01a5e764913e0b486b68fb92
  root in ~ λ docker inspect -f '{{.Author}}' bec4e0e815ac
  cyberpelican
  ```

- `-c`

  按照Dockerfile的格式指定生成的镜像

  ```
  root in ~ λ docker commit --change='CMD top '  c082c5068def  cbb
  sha256:22162bb10d4c16e6e69b0395a10bb807f5029e6dc9e1a5b0b55ebc40dff1cc2c
  root in ~ λ docker images
  REPOSITORY                                                        TAG          IMAGE ID       CREATED         SIZE
  cbb                                                               latest       22162bb10d4c   7 seconds ago   1.23MB
  
  root in ~ λ docker run --name t2 -d cbb
  da045bd29d53f2fe6e27fc1605ef0ae056346a4bac928592ec57ff99404405fa
  root in ~ λ docker ps
  CONTAINER ID   IMAGE     COMMAND            CREATED          STATUS          PORTS     NAMES
  da045bd29d53   cbb       "/bin/sh -c top"   3 seconds ago    Up 3 seconds              t2
  ```

## 镜像删除

参考：

https://www.jianshu.com/p/3c2e0a89d618

如果以现有的镜像（A）commit后，那么生成后的镜像（B）就是A的子依赖。删除A会提示

`Error response from daemon: conflict: unable to delete 491198851f0c (cannot be forced) - image has dependent child images`

我们可以使用`docker inspect`来查看子镜像然后删除

```
root in ~ λ docker inspect -f '{{.RepoTags}}{{.Id}}{{.Parent}}' $(docker images -q)
[cbb:latest]sha256:22162bb10d4c16e6e69b0395a10bb807f5029e6dc9e1a5b0b55ebc40dff1cc2csha256:491198851f0ccdd0882cb9323f3856043d4e4c65b773e8eac3e0f6bc979a2ae7
[busybox:latest]sha256:491198851f0ccdd0882cb9323f3856043d4e4c65b773e8eac3e0f6bc979a2ae7
[nginx:latest]sha256:f6d0b4767a6c466c178bf718f99bea0d3742b26679081e52dbf8e0c7c4c42d74
[centos:latest]sha256:300e315adb2f96afe5f0b2780b87f28ae95231fe3bdd1e16b9ba606307728f55

root in ~ λ docker rmi 22162bb10d4c16e6e69b0395a10bb807f5029e6dc9e1a5b0b55ebc40dff1cc2c
Untagged: cbb:latest
Deleted: sha256:22162bb10d4c16e6e69b0395a10bb807f5029e6dc9e1a5b0b55ebc40dff1cc2c
Deleted: sha256:bb9e1f3638a8d8eb492bc4c1aab392f86ff547aaf1e61bf73889a30af329a1ed
root in ~ λ docker rmi 491198851f0c
Untagged: busybox:latest
Untagged: busybox@sha256:c6b45a95f932202dbb27c31333c4789f45184a744060f6e569cc9d2bf1b9ad6f
Deleted: sha256:491198851f0ccdd0882cb9323f3856043d4e4c65b773e8eac3e0f6bc979a2ae7
Deleted: sha256:84009204da3f70b09d2be3914e12844ae9db893aa85ef95df83604f95df05187

root in ~ λ docker images
REPOSITORY   TAG       IMAGE ID       CREATED        SIZE
nginx        latest    f6d0b4767a6c   5 weeks ago    133MB
centos       latest    300e315adb2f   2 months ago   209MB
```







