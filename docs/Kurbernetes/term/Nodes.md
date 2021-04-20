# Nodes

参考：

https://kubernetes.io/docs/concepts/architecture/nodes/

Kubernetes runs your workload by placing containers into Pods to run on Nodes.

Node可以是物理机，也可以是虚拟机。每个Node都有kubelet，container runtime 和 kube-proxy，并且由control plane管理。

==如果Node需要更新或替换，必须先从control plane中删除。==

## Management

有两种主要的途径将Node添加到control plane的管理：

### self-registration of nodes

kubelet代理node主动注册到control plane，==这是默认的方式==。主动注册需要在kubelet启动时添加如下参数：

- --kubeconfig，指定kubeconfig
- --cloud-provider 
- --register-node 默认true
- --register-with-taints
- --node-ip
- --node-lables
- --node-status-update-frequency

### manual node administration

手动添加node到control plane

将--register-node 设置为false

---

可以使用`systemctl status kubelet -l`查看kubelet的启动参数

```
[root@k8snode01 ~]# systemctl status -l kubelet
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/usr/lib/systemd/system/kubelet.service; enabled; vendor preset: disabled)
  Drop-In: /usr/lib/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Wed 2021-03-31 09:57:34 CST; 37min ago
     Docs: https://kubernetes.io/docs/
 Main PID: 110579 (kubelet)
    Tasks: 14
   CGroup: /system.slice/kubelet.service
           └─110579 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --config=/var/lib/kubelet/config.yaml --network-plugin=cni --pod-infra-container-image=registry.aliyuncs.com/google_containers/pause:3.2
```

## Node  name

两个不同的Node可以有相同的名字(默认主机的hostname)，但是kubernetes会将名字相同的节点认定为同一个对象。

## status

可以通过`kubectl describe node/nodename`来查看node 状态，有4个属性：Address，Conditions，Capacity And Allocate，info

### Address

通常提供如下几个属性

- HostName：nodename
- ExternalIP：集群外可访问的IP
- InternalIP：集群内访问的IP

```
Addresses:
  InternalIP:  192.168.80.201
  Hostname:    k8sm
```

### Conditions

conditions用于描述运行的node

- Ready：true代表node健康，反之false
- DiskPressure：true代表node磁盘可用空间不足，反之false
- MemoryPressure：true代表node内存空间不足，反之false
- PIDPressure：true代表node打开的进程过多，反之false
- NetworkUnavailable：true代表node网络配置有问题，反之false

```
Conditions:
  Type                 Status  LastHeartbeatTime                 LastTransitionTime                Reason                       Message
  ----                 ------  -----------------                 ------------------                ------                       -------
  NetworkUnavailable   False   Mon, 29 Mar 2021 09:48:09 +0800   Mon, 29 Mar 2021 09:48:09 +0800   CalicoIsUp                   Calico is running on this node
  MemoryPressure       False   Wed, 31 Mar 2021 10:34:33 +0800   Fri, 19 Mar 2021 11:23:53 +0800   KubeletHasSufficientMemory   kubelet has sufficient memory available
  DiskPressure         False   Wed, 31 Mar 2021 10:34:33 +0800   Fri, 19 Mar 2021 11:23:53 +0800   KubeletHasNoDiskPressure     kubelet has no disk pressure
  PIDPressure          False   Wed, 31 Mar 2021 10:34:33 +0800   Fri, 19 Mar 2021 11:23:53 +0800   KubeletHasSufficientPID      kubelet has sufficient PID available
  Ready                True    Wed, 31 Mar 2021 10:34:33 +0800   Fri, 19 Mar 2021 11:23:57 +0800   KubeletReady                 kubelet is posting ready status
```

### Capacity  and Allocate

node CPU和RAM总容量和可分配资源

```
Capacity:
  cpu:                2
  ephemeral-storage:  17394Mi
  hugepages-1Gi:      0
  hugepages-2Mi:      0
  memory:             1863104Ki
  pods:               110
Allocatable:
  cpu:                2
  ephemeral-storage:  16415037823
  hugepages-1Gi:      0
  hugepages-2Mi:      0
  memory:             1760704Ki
  pods:               110
```

### info

general information of node

```
System Info:
  Machine ID:                 fbb74b6620184684961580de92e236c2
  System UUID:                CBAE4D56-039C-B492-3C21-FA02D20D121D
  Boot ID:                    41ece3b2-fd27-472a-a12b-3dc6efad6ec4
  Kernel Version:             3.10.0-1062.el7.x86_64
  OS Image:                   CentOS Linux 7 (Core)
  Operating System:           linux
  Architecture:               amd64
  Container Runtime Version:  docker://1.13.1
  Kubelet Version:            v1.20.4
  Kube-Proxy Version:         v1.20.4
```

## Graceful Node shutdown

https://kubernetes.io/docs/concepts/architecture/nodes/#graceful-node-shutdown



