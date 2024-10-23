---
createTime: 2024-10-23 13:44
tags:
  - "#hash1"
  - "#hash2"
---

# Docker 01 - Installation

## 0x01 Preface

Docker 根据是否有 GUI 可以将安装方式分为 2 种

1. 有 GUI —— Docker Desktop[^1]
2. 无 GUI —— Docker Engine

## 0x02 Docker Desktop

> [!NOTE]
> Docker Desktop 自带 engine（可以在 settings 中的 builders 面板中看到），即安装 Docker Desktop 无须安装 Engine 就可以使用，亦或者本地安装的 engine 可以和 Docker Desktop 自带的 Engine 同时运行(两者之间互相隔离)

安装方式请看，Docker Desktop 官方文档[^1]

## 0x03 Docker Engine

官方为 engine 提供了不同方式的安装，主要两种

1. Package Manager —— 通过包管理器安装
2. Generic —— 通用安装方式

这里主要介绍 Generic 安装，Package Manager 安装的具体可以按照 Distro 查看 Engine Install[^2]

### 0x03a Generic

> [!NOTE]
> 官方并不推荐使用 Generic 这种方式安装，因为不能自动升级，就会导致漏洞不能及时修复

Ceneric 通过二进制的方式安装

#### Download Binary Archive

按照 CPU 架构从 [Index of linux/static/stable/](https://download.docker.com/linux/static/stable/) 下载对应的版本。例如

```
wget https://download.docker.com/linux/static/stable/x86_64/docker-27.3.1.tgz
```

#### Extract tarball

解压

```
tar -xvzf docker-27.3.1.tgz
```

解压后的目录如下

```
tree docker
docker
├── containerd
├── containerd-shim-runc-v2
├── ctr
├── docker
├── dockerd
├── docker-init
├── docker-proxy
└── runc
```

#### Copy Binaries to PATH

将 binaries 拷贝到 PATH(如果直接运行，可能会出现异常)

```
sudo cp docker/* /usr/bin/
```

#### Add Docker User Group

添加 docker 用户组

```
groupadd docker
```

#### Test Docker Daemon

测试 docker 守护进程

```
sudo dockerd
```

#### Verify Docker 

验证 docker 是否正常启用

```
docker run hello-world
```

#### Systemd Unit file

为了方便管理 docker 可以注册一个 unit file

```
cat >> /usr/lib/systemd/system/dockerd.service << EOF
[Unit]
After=network-online.target
Wants=network-online.target 
[Service]
ExecStart=/usr/bin/dockerd
ExecReload=/bin/kill -s HUP $MAINPID
RestartSec=2
Restart=always
[Install]
WantedBy=multi-user.target
EOF
```

## 0x04 Trouble Shooting

### 0x04a invalid userland-proxy-path: userland-proxy is enabled, but userland-proxy-path is not set

执行 `dockerd` 报错，需按照 [Copy Binaries to PATH](#Copy%20Binaries%20to%20PATH) 处理

### 0x04b could not change group /var/run/docker.sock to docker: group docker not found

执行 `useradd docker`

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***References***

- [Install Docker Engine | Docker Docs](https://docs.docker.com/engine/install/)
- [Binaries | Docker Docs](https://docs.docker.com/engine/install/binaries/)
- [Install | Docker Docs](https://docs.docker.com/engine/install/)


***FootNotes***

[^1]:[Docker Desktop | Docker Docs](https://docs.docker.com/desktop/)
[^2]:[Install | Docker Docs](https://docs.docker.com/engine/install/)


