# kubeadm init 

init a k8s control plane.

before read check the [ execute follow][https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-init/]

## init flags

- --apiserver-advertise-address string

  --apiserver-bind port int32

  指定api server的地址和端口。如果没有设置默认使用本机的iface和6443端口

- --dry-run

  用于测试

- --config string

  指定k8s的配置文件位置，默认`~/.kube/config`

- --pod-network-cidr string

  指定pod的CIDR网络，如果指定里control plane会自动分配IP给每一个node

  缺省值：10.96.0.0/12

- --kubernetes-version string

  指定k8s(docker 上下载下来的镜像)的版本，默认stable-1

  ==如果出现找不到镜像有可能是国内的镜像同步组件慢了导致找不到镜像，指定特定版本即可==

  ```
  [root@k8smaster /]# kubeadm init --v=5 --image-repository registry.aliyuncs.com/google_containers --pod-network-cidr 192.168.0.0/16 --kubernetes-version 1.20.4
  ```

- --node-name

  指定control plane node的名字，默认以当前主机hostname为nodename

  ```
  [root@k8smaster /]# kubectl get nodes --all-namespaces 
  NAME   STATUS   ROLES                  AGE   VERSION
  k8sm   Ready    control-plane,master   36s   v1.20.4
  
  [root@k8smaster /]# kubeadm init --v=5 --image-repository registry.aliyuncs.com/google_containers --pod-network-cidr 192.168.0.0/16 --node-name k8sm --kubernetes-version 1.20.4
  ```

- token-ttl duration

  指定token的有效时间，默认24h0m0s。如果设置为0表示永不过期。

- --skip-phases stringSlice

  跳过指定phases

## 无网络初始化

没有网络运行kubeamd，需要提前下载镜像。

检查需要下载的镜像

```
[root@k8smaster manifests]# kubeadm config images list --image-repository registry.aliyuncs.com/google_containers --kubernetes-version 1.20.4
registry.aliyuncs.com/google_containers/kube-apiserver:v1.20.4
registry.aliyuncs.com/google_containers/kube-controller-manager:v1.20.4
registry.aliyuncs.com/google_containers/kube-scheduler:v1.20.4
registry.aliyuncs.com/google_containers/kube-proxy:v1.20.4
registry.aliyuncs.com/google_containers/pause:3.2
registry.aliyuncs.com/google_containers/etcd:3.4.13-0
registry.aliyuncs.com/google_containers/coredns:1.7.0
```

下载镜像

```
[root@k8smaster manifests]# kubeadm config images pull --image-repository registry.aliyuncs.com/google_containers --kubernetes-version 1.20.4
```

