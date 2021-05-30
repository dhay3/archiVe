# docker-images

> 复合组件`docker image`
>
> 参考：https://docs.docker.com/engine/reference/commandline/image/

## 概述

用于镜像，默认支持container name模糊查询

```
$ docker images java

REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
java                8                   308e519aac60        6 days ago          824.5 MB
java                7                   493d82594c15        3 months ago        656.3 MB
java                latest              2711b1d6f3aa        5 months ago        603.9 MB
```

## 参数

- `-a`

  展示所有，不带`a`参数默认不展示中间容器

- `--no-trunc`

  显示完整的containerid

  ```
  root in ~ λ docker images --no-trunc
  REPOSITORY   TAG       IMAGE ID                                                                  CREATED       SIZE
  busybox      latest    sha256:b97242f89c8a29d13aea12843a08441a4bbfc33528f55b60366c1d8f6923d0d4   3 weeks ago   1.23MB
  nginx        latest    sha256:f6d0b4767a6c466c178bf718f99bea0d3742b26679081e52dbf8e0c7c4c42d74   3 weeks ago   133MB
  centos       latest    sha256:300e315adb2f96afe5f0b2780b87f28ae95231fe3bdd1e16b9ba606307728f55   8 weeks ago   209MB   /0.0s
  ```

- `-q`

  只显示镜像的id

  ```
  root in ~ λ docker images -q
  b97242f89c8a
  f6d0b4767a6c
  ```

- `--format`

  以go template显示指定的部分

  ```
  root in ~ λ docker images --format "{{.ID}}"
  b97242f89c8a
  f6d0b4767a6c
  300e315adb2f  
  ```

  