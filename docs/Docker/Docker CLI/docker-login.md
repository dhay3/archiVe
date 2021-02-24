# docker login

参考：

https://help.aliyun.com/document_detail/198212.html?spm=a2c4g.11186623.6.554.57bb426aopMSLy

## 概述

docker login 用于登入镜像仓库。默认镜像仓库是https://hub.docker.com，可以另外指定。

## 参数

- `-u`

  指定登入的账号

- `-p`

  指定登入的密码

- `--password-stdin`

  从stdin中读取密码

## 绑定aliyun镜像仓库

Docker的镜像地址

`registry.cn-hangzhou.aliyuncs.com/acs/agent:0.8`

- `registry.cn-hangzhou.aliyuncs.com`是Registry的域名。
- `acs`是您所使用的命名空间的名称。
- `agent`是您所使用的仓库的名称。
- `0.8`是镜像标签（Tag）。非必须，默认为latest。

> 登入的用户名同登入aliyun的账号相同，密码为访问凭证和aliyun账号的密码不同。
>
> 在容器镜像服务中访问凭证设置

```
root in ~ λ docker login registry.cn-shenzhen.aliyuncs.com/cyberpelican/aliyun_dockerhub
Username: kikochz
Password:
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store
```

用户的权鉴信息会以Base64编码的信息存储在`$HOME/.docker/config.json`中。可以使用`docker logout`退出并清除权鉴。

可以将本地镜像上传到改镜像仓库。



