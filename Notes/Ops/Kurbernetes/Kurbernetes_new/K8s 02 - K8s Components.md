---
createTime: 2025-06-03 16:22
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# K8s 02 - K8s Components

## 0x01 Preface

> [!NOTE]
> ![[https://www.youtube.com/watch?v=TlHvYWVUZyc]] 这个视频很好地解释了 k8s 各个组件之间的关系以及功能

![](https://kubernetes.io/images/docs/components-of-kubernetes.svg)

一个 k8s 集群通常由一个 control plane 和 一个或者多个 nodes 组成

## 0x02 Control Plane
 
> Manage the worker nodes and the Pods in the cluster
> 
> 整个集群的管理节点，决定了 Node 的调度

由几个组件构成

- kube-apiserver
- etcd
- kube-scheduler
- kube-controller-manager
- cloud-controller-manager(optional)

### 0x02a kube-apiserver

> The API server is a component of the Kubernetes control plane that exposes the Kubernetes API. The API server is the front end for the Kubernetes control plane.
exit

kube-apiserver 就是 k8s 对外暴露的接口，类似于 Shell 是 kernel 对外暴露的接口

### 0x02b etcd

> Consistent and highly-available key value store used as Kubernetes' backing store for all cluster data.

一个符合 CA 原则的 key value 数据库，用于储存 k8s 集群数据

### 0x02c kube-scheduler

> Control plane component that watches for newly created Pods with no assigned node, and selects a node for them to run on.

调度器，为 Pod 分配合理的 Node

### 0x02d kube-controller-manager

> Control plane component that runs controller processes.
> 
> Logically, each controller is a separate process, but to reduce complexity, they are all compiled into a single binary and run in a single process.

提供控制功能的集成组件

### 0x02e cloud-controller-manager

> The cloud controller manager lets you link your cluster into your cloud provider's API, and separates out the components that interact with that cloud platform from components that only interact with your cluster.

提供云功能的集成组件

## 0x03 Node

> The worker node(s) host the Pods that are the components of the application workload.
> 
> 集群的 worker 运行 Pods

由几个组件构成

- kubelet
- kube-proxy(optional)
- container runtime

### 0x03a kubelet

> An agent that runs on each node in the cluster. It makes sure that containers are running in a Pod.

Node 上的 controller，确保容器运行在 Pod 中

### 0x03b kube-proxy

> kube-proxy is a network proxy that runs on each node in your cluster, implementing part of the Kubernetes Service concept.

为 Pod 提供代理，类似于 Nginx，通过 Proxy 可以在集群中互相通信






---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [Kubernetes Components \| Kubernetes](https://kubernetes.io/docs/concepts/overview/components/)
- [Cluster Architecture \| Kubernetes](https://kubernetes.io/docs/concepts/architecture/)

***References***

