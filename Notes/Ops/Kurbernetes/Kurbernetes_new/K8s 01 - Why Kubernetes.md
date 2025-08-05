---
createTime: 2025-06-03 16:37
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# K8s 01 - Why Kubernetes

## 0x01 Preface

> Kubernetes is a portable, extensible, open source platform for managing containerized workloads and services, that facilitates both declarative configuration and automation. It has a large, rapidly growing ecosystem. Kubernetes services, support, and tools are widely available.

简而言之 K8s/kei:z/ 是一个配置化，自动化的容器管理平台

## 0x02 Why Kubernetes

1. **Service discovery and load balancing**

	Kubernetes can expose a container using the DNS name or using their own IP address. If traffic to a container is high, Kubernetes is able to load balance and distribute the network traffic so that the deployment is stable.

	能通过 DNS 或者 IP 暴露容器的服务，且提供负载均衡

2. **Storage orchestration**

	Kubernetes allows you to automatically mount a storage system of your choice, such as local storages, public cloud providers, and more.

	容器使用的存储可以有多种选择，本地或者云等等

3. **Automated rollouts and rollbacks**

	You can describe the desired state for your deployed containers using Kubernetes, and it can change the actual state to the desired state at a controlled rate. For example, you can automate Kubernetes to create new containers for your deployment, remove existing containers and adopt all their resources to the new container.

	指定容器预期的状态，根据状态调整容器

4. **Self-healing**

	Kubernetes restarts containers that fail, replaces containers, kills containers that don't respond to your user-defined health check, and doesn't advertise them to clients until they are ready to serve.

	容器自愈(自动 kill/replace 无响应的容器)

5. **Secret and configuration management**

	Kubernetes lets you store and manage sensitive information, such as passwords, OAuth tokens, and SSH keys. You can deploy and update secrets and application configuration without rebuilding your container images, and without exposing secrets in your stack configuration.

	安全的密码配置管理，无需重构容器

6. **Batch execution**

	In addition to services, Kubernetes can manage your batch and CI workloads, replacing containers that fail, if desired.

	管理 CI(continuos integration) 防止容器构建失败

7. **Horizontal scaling**

	Scale your application up and down with a simple command, with a UI, or automatically based on CPU usage.

	根据 CPU 的使用率自动扩缩容容器

8. **IPv4/IPv6 dual-stack**

	Allocation of IPv4 and IPv6 addresses to Pods and Services

	分配 v4/v6 双栈容器地址

9. **Designed for extensibility**

	Add features to your Kubernetes cluster without changing upstream source code.

	无需修改 k8s 源码，可以通过插件的形式增加新功能

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [Overview \| Kubernetes](https://kubernetes.io/docs/concepts/overview/)


***References***


