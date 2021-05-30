# docker volume

docker volume用于管理容器卷。

## 创建

volume name 必须uniq。==volume不是创建在当前目录且不能以绝对路径为命令行参数，挂载点由docker自动生成。==`docker run -v`同样可以指定容器卷挂载，且能以绝对路径挂载。 

```
root in /usr/local/\/docker_test λ docker volume create hello
hello
root in /usr/local/\/docker_test λ ls
t1  t2  test.sh
root in /usr/local/\/docker_test λ docker run -itd --name t1 -v hello:/etc --rm busybox
94a996549284a74cb8838af14b580cb84b8a5b317b6d295fc507d94c726e5641
```

## 查看

使用`docker volume inspect <volume-name>`查看，和docker inspect参数相同。

```
root in /usr/local/\/docker_test λ docker volume inspect hello
[
    {
        "CreatedAt": "2021-02-22T16:51:23+08:00",
        "Driver": "local",
        "Labels": {},
        "Mountpoint": "/var/lib/docker/volumes/hello/_data",
        "Name": "hello",
        "Options": {},
        "Scope": "local"
    }
]
root in /usr/local/\/docker_test λ cd /var/lib/docker/volumes/hello/_data
root in /var/lib/docker/volumes/hello/_data λ ls
group  hostname  hosts  localtime  mtab  network  passwd  resolv.conf  shadow
```

使用`docker volume ls`来查看当前宿主机上所有与容器挂载的卷

```
root in / λ docker volume ls
DRIVER    VOLUME NAME
local     e266cda9aaef84862674ff49822e6d6da59fd62e2c04a02f9e579a62cc3f4de2
local     hello
```

## 删除

使用`docker volume prune`删除所有没有与容器挂载的volume

```
root in / λ docker volume prune
WARNING! This will remove all local volumes not used by at least one container.
Are you sure you want to continue? [y/N] y
Deleted Volumes:
e266cda9aaef84862674ff49822e6d6da59fd62e2c04a02f9e579a62cc3f4de2

Total reclaimed space: 12B
```

`docker volume rm`用于删除指定的volume，但是不能删除被容器正在使用的volume

```
root in / λ docker rm hello
Error: No such container: hello
```

