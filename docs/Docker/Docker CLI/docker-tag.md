# docker-tag

docker tag 将本地镜像以指定的标签打包。

pattern：`docker tag <image> <target-images-name>[:version]`

```
root in ~ λ docker tag 491198851f0c busyb                                                                                                   /0.1s
root in ~ λ docker images
REPOSITORY                                                        TAG          IMAGE ID       CREATED        SIZE
busyb                                                             latest       491198851f0c   3 days ago     1.23MB
busybox                                                           latest       491198851f0c   3 days ago     1.23MB
registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub   version1.0   491198851f0c   3 days ago     1.23MB

```

如果需要将镜像推送到私有的registry，打包的名字中必须是registry的hostname和port(如果需要指定)。可以指定推送的版本也可以不指定。

```
root in ~ λ docker tag 491198851f0c registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub:version1.0                              /0.1s
root in ~ λ docker images
REPOSITORY                                                        TAG          IMAGE ID       CREATED        SIZE
busybox                                                           latest       491198851f0c   3 days ago     1.23MB
registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub   version1.0   491198851f0c   3 days ago     1.23MB
nginx                                                             latest       f6d0b4767a6c   5 weeks ago    133MB       
```

==注意生成的镜像images id相同(指针执行同一个物理地址，如果以imagesId删除会删除所有的镜像)==

```
root in ~ λ docker images
REPOSITORY                                                        TAG          IMAGE ID       CREATED        SIZE
busyb                                                             latest       491198851f0c   3 days ago     1.23MB
busybox                                                           latest       491198851f0c   3 days ago     1.23MB
registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub   version1.0   491198851f0c   3 days ago     1.23MB


root in ~ λ docker rmi -f 491198851f0c
Untagged: busyb:latest
Untagged: busybox:latest
Untagged: busybox@sha256:c6b45a95f932202dbb27c31333c4789f45184a744060f6e569cc9d2bf1b9ad6f
Untagged: registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub:version1.0
Untagged: registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub@sha256:74e4a68dfba6f40b01787a3876cc1be0fb1d9025c3567cf8367c659f2187234f
Deleted: sha256:491198851f0ccdd0882cb9323f3856043d4e4c65b773e8eac3e0f6bc979a2ae7
Deleted: sha256:84009204da3f70b09d2be3914e12844ae9db893aa85ef95df83604f95df05187     

root in ~ λ docker images
REPOSITORY   TAG       IMAGE ID       CREATED        SIZE
nginx        latest    f6d0b4767a6c   5 weeks ago    133MB
centos       latest    300e315adb2f   2 months ago   209MB           
```

imageId相同的镜像使用repository+tag来删除

```
root in ~ λ docker rmi byb                                                                         
root in ~ λ docker images
REPOSITORY                                                        TAG          IMAGE ID       CREATED        SIZE
busybox                                                           latest       491198851f0c   3 days ago     1.23MB
byc                                                               latest       491198851f0c   3 days ago     1.23MB
registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub   version1.0   491198851f0c   3 days ago     1.23MB
nginx                                                             latest       f6d0b4767a6c   5 weeks ago    133MB
centos                                                            latest       300e315adb2f   2 months ago   209MB 
```

