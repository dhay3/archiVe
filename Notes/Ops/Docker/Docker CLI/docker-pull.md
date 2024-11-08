# docker-pull

docker pull默认从docker hub 上拉取镜像。

```
$ docker pull debian

Using default tag: latest
latest: Pulling from library/debian
fdd5d7827f33: Pull complete
a3ed95caeb02: Pull complete
Digest: sha256:e7d38b3517548a1c71e41bffe9c8ae6d6d29546ce46bf62159837aad072c90aa
Status: Downloaded newer image for debian:latest
```

镜像由多个layer组成。layer能被镜像复用。在下载完后docker会将镜像sha256 digest打印出。

docker pull 可以通过digest来拉取镜像，一般用于校验镜像。digest同样可以被用在Dockerfile中。

```
root in ~ λ docker pull busybox@sha256:c6b45a95f932202dbb27c31333c4789f45184a744060f6e569cc9d2bf1b9ad6f
docker.io/library/busybox@sha256:c6b45a95f932202dbb27c31333c4789f45184a744060f6e569cc9d2bf1b9ad6f: Pulling from library/busybox
Digest: sha256:c6b45a95f932202dbb27c31333c4789f45184a744060f6e569cc9d2bf1b9ad6f
Status: Image is up to date for busybox@sha256:c6b45a95f932202dbb27c31333c4789f45184a744060f6e569cc9d2bf1b9ad6f
docker.io/library/busybox@sha256:c6b45a95f932202dbb27c31333c4789f45184a744060f6e569cc9d2bf1b9ad6f
```

可以使用指定registry来拉取镜像。通过docker login 来管理权鉴。

```
root in ~ λ docker pull  registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub:version1.0
version1.0: Pulling from cyberpelican/aliyun_dockerhub
Digest: sha256:74e4a68dfba6f40b01787a3876cc1be0fb1d9025c3567cf8367c659f2187234f
Status: Downloaded newer image for registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub:version1.0
registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub:version1.0
root in ~ λ docker images
REPOSITORY                                                        TAG          IMAGE ID       CREATED        SIZE
busybox                                                           latest       491198851f0c   3 days ago     1.23MB
registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub   version1.0   491198851f0c   3 days ago     1.23MB
nginx                                                             latest       f6d0b4767a6c   5 weeks ago    133MB
centos                                                            latest       300e315adb2f   2 months ago   209MB
```

