# docker-push

将image或repository推送到指定的registry。

pattern：`docker push <image-name[:version] | imageId>`

```
root in ~ λ docker push  registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub:version1.0
The push refers to repository [registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub]
84009204da3f: Pushed
version1.0: digest: sha256:74e4a68dfba6f40b01787a3876cc1be0fb1d9025c3567cf8367c659f2187234f size: 527 
```

这里推送的镜像不是registry而是镜像名

<img src="C:\Users\82341\AppData\Roaming\Typora\typora-user-images\image-20210222150042546.png" alt="image-20210222150042546" style="zoom:80%;" />

使用docker pull 来校验

```
root in ~ λ docker pull registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub:version1.0
version1.0: Pulling from cyberpelican/aliyun_dockerhub
Digest: sha256:74e4a68dfba6f40b01787a3876cc1be0fb1d9025c3567cf8367c659f2187234f
Status: Downloaded newer image for registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub:version1.0
registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub:version1.0                                                                  /0.4s
root in ~ λ docker images
REPOSITORY                                                        TAG          IMAGE ID       CREATED        SIZE
busybox                                                           latest       491198851f0c   3 days ago     1.23MB
registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub   version1.0   491198851f0c   3 days ago     1.23MB
nginx                                                             latest       f6d0b4767a6c   5 weeks ago    133MB
centos                                                            latest       300e315adb2f   2 months ago   209MB                          /0.0s
```

## 参数

- `-a`
  一次推送所有的tag

  