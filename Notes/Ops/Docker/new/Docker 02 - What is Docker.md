---
createTime: 2024-10-23 17:52
tags:
  - "#hash1"
  - "#hash2"
---

# Docker 02 - What is Docker

## 0x01 Preface

> Docker provides the ability to package and run an application in a loosely isolated environment called a container. The isolation and security lets you run many containers simultaneously on a given host. Containers are lightweight and contain everything needed to run the application, so you don't need to rely on what's installed on the host. You can share containers while you work, and be sure that everyone you share with gets the same container that works in the same way.

简而言之 Docker 可以让你的程序以沙盒的方式运行，且可以让你的程序具有方便移植的特性

## 0x02 Docker Architecture

Docker 使用 C/S 架构

![](https://docs.docker.com/get-started/images/docker-architecture.webp)

### 0x02a Docker Daemon

Docker Daemon(`dockerd`) 是 Docker 中最核心的部分，处理容器的构建、运行以及分发

### 0x02b Docker Client

Docker Client(`docker`) 是 userspace 的管理工具，通过它想 Docker Daemon 下发指令。Daemon 可以和 Client 在同一台机器上，也可以不在同一台机器上

### 0x02c Docker Registries[^1]

Docker Registries 是一个用于存储和分享 Docker Images 的中心化地址(可以是公开的，也可以是私有的)。例如 Docker hub 就是一个公开的 registry，Docker 默认会从 Docker hub 拉取镜像

当然也有一些第三方的 Registries

1. Amazon Elastic Container Registry(ECR)
2. Azure Container Registry(ACR)
3. Google Container Registry(GCR)
4. Github Container Registry(GHCR)

我们常说的 CR 也就指的是 Container Registry

#### Repository

Docker Registry 中有和 Github Repository 一样的逻辑概念(registry 由多个 repositories 组成，一个 repository 由多个 images 组成)。下面这张图就很好的解释了 registry/repository/image 之间的关系

![](https://github.com/dhay3/picx-images-hosting/raw/master/2024-10-25_09-44-47.7paji9jwm.png)


### 0x02d Docker Objects

#### Images[^2]

> An image is a read-only template with instructions for creating a Docker container.

image 就是 container 的一个只读模板，有 2 条准则

1. Images are immutable. Once an image is created, it can't be modified. You can only make a new image or add changes on top of it.
2. Container images are composed of layers. Each layer represented a set of file system changes that add, remove, or modify files.

虽然 images 是 immutable 的，但是你可以通过修改 container ，然后将这个 container 生成 image

#### Containers[^3]

> A container is a runnable instance of an image.

container 就是某个 image 的实例

例如 `docker run -i -t ubuntu /bin/bash` 就会使用 ubuntu image 生成一个 container 实例，会发生如下过程

1. If you don't have the `ubuntu` image locally, Docker pulls it from your configured registry, as though you had run `docker pull ubuntu` manually.
    
2. Docker creates a new container, as though you had run a `docker container create` command manually.
    
3. Docker allocates a read-write filesystem to the container, as its final layer. This allows a running container to create or modify files and directories in its local filesystem.
    
4. Docker creates a network interface to connect the container to the default network, since you didn't specify any networking options. This includes assigning an IP address to the container. By default, containers can connect to external networks using the host machine's network connection.
    
5. Docker starts the container and executes `/bin/bash`. Because the container is running interactively and attached to your terminal (due to the `-i` and `-t` flags), you can provide input using your keyboard while Docker logs the output to your terminal.
    
6. When you run `exit` to terminate the `/bin/bash` command, the container stops but isn't removed. You can start it again or remove it.

##### Why Need Containers

假设有这么一个场景

> Imagine you're developing a killer web app that has three main components - a React frontend, a Python API, and a PostgreSQL database. If you wanted to work on this project, you'd have to install Node, Python, and PostgreSQL.

你如何保证团队里的其他开发，CI/CD 使用相同组件版本，怎么处理不同组件之间的冲突（例如 python 3.8 和 3.12 之前包依赖的不同）。那么你可以通过容器来解决这些问题

- Self-contained. Each container has everything it needs to function with no reliance on any pre-installed dependencies on the host machine.
- Isolated. Since containers are run in isolation, they have minimal influence on the host and other containers, increasing the security of your applications.
- Independent. Each container is independently managed. Deleting one container won't affect any others.
- Portable. Containers can run anywhere! The container that runs on your development machine will work the same way in a data center or anywhere in the cloud!

##### Container VS VM

VM 是一个完整的操作系统，有自己的 kernel，hardware drivers，programs 以及 applications，如果只为了一个应用来使用 VM 会有点 overhead

而 container 可以只为提供应用启动的最基础的环境(例如 nginx container 就没有 gcc)

> **Using VMs and containers together**
> 
> Quite often, you will see containers and VMs used together. As an example, in a cloud environment, the provisioned machines are typically VMs. However, instead of provisioning one machine to run one application, a VM with a container runtime can run multiple containerized applications, increasing resource utilization and reducing costs.

实际上 VMs 会和 containers 一起使用，大多数云服务商提供的服务器都是由 KVM 虚拟的，而一些 SaaS(底层是由 K8s 来管理的)是运行在 VMs 上的

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***References***

- [What is Docker? | Docker Docs](https://docs.docker.com/get-started/docker-overview/)
- [What is a container? | Docker Docs](https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-a-container/)
- [What is an image? | Docker Docs](https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-an-image/)
- [What is a registry? | Docker Docs](https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-a-registry/)


***FootNotes***

[^1]:[What is a registry? | Docker Docs](https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-a-registry/)
[^2]:[What is an image? | Docker Docs](https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-an-image/)
[^3]:[What is a registry? | Docker Docs](https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-a-registry/)