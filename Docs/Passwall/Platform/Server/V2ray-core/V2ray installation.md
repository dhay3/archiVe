# V2ray installation

ref:

https://www.v2ray.com/chapter_00/install.html

安装的 v2ray 需要注意以下文件的

- `/usr/bin/v2ray/v2ray`：V2Ray 程序；
- `/usr/bin/v2ray/v2ctl`：V2Ray 工具；
- `/etc/v2ray/config.json`：配置文件；
- `/usr/bin/v2ray/geoip.dat`：IP 数据文件
- `/usr/bin/v2ray/geosite.dat`：域名数据文件

## 手动安装

通过`uname -a`或是`arch`来查看 CPU 架构

1. aarch == arm
2. x86_64/x86_32 == x86

下载对应架构的 [压缩包](https://github.com/v2ray/v2ray-core/releases)

## 自动安装脚本

不推荐

gitbook 里给出的自动安装脚本已过时，替换成 https://github.com/v2fly/fhs-install-v2ray 仓库下的脚本

## Docker

v2ray 还可以通过 Docker 的方式来安装

```
#stable版本
docker pull v2ray/official
#开发版本
docker pull v2ray/dev

#进入容器
docker run -it --name v2ray v2ray/official sh

#直接启动 v2ray
docker run -it --name v2ray v2ray/official
```

