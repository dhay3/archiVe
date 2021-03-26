# kubeadm config

在`kubeadm init`时kubead上传clustercoonfiguration到kubeadm-config(Map对象)。在`kubeadm join`，`kubeadm reset`，`kubeadm upgrade`都会被读取。可以通过`kubeadm config view`来读取配置文件。

## 配置

### kubeadm config view

kubeadm用于读取kubeadm的配置文件，新版本使用`kubectl get cm -o yaml -n kube-system kubeadm-config`

```
[root@k8smaster ~]# kubeadm config view
Command "view" is deprecated, This command is deprecated and will be removed in a future release, please use 'kubectl get cm -o yaml -n kube-system kubeadm-config' to get the kubeadm config directly.
apiServer:
  extraArgs:
    authorization-mode: Node,RBAC
  timeoutForControlPlane: 4m0s
apiVersion: kubeadm.k8s.io/v1beta2
certificatesDir: /etc/kubernetes/pki
clusterName: kubernetes
controllerManager: {}
dns:
  type: CoreDNS
etcd:
  local:
    dataDir: /var/lib/etcd
imageRepository: registry.aliyuncs.com/google_containers
kind: ClusterConfiguration
kubernetesVersion: v1.20.4
networking:
  dnsDomain: cluster.local
  podSubnet: 192.168.0.0/16
  serviceSubnet: 10.96.0.0/12
scheduler: {}

```

- --kubeconfig string

  指定kubeconfig的配置文件的位置。默认`/etc/kubernetes/admin.conf`

### kubeadm config print  init-defaults

打印kubeadm init 使用的配置

```
[root@k8smaster ~]# kubeadm config print init-defaults  | less
```

注意这里打印的token并不是真实的token，而是`abcdef.0123456789abcdef`只表示占位符。

- --component-configs [KubeletConfiguration | KubeProxyConfiguration]

  额外打印组件的配置，默认不会打印

  ```
  [root@k8smaster ~]# kubeadm config print init-defaults --component-configs KubeletConfiguration
  ```

- --kubeconfig string

  指定配置文件的位置。默认`/etc/kubernetes/admin.conf`

### kubeadm config print join-defaults

打印kubeadm join 使用的配置

```
[root@k8smaster ~]# kubeadm config print join-defaults
apiVersion: kubeadm.k8s.io/v1beta2
caCertPath: /etc/kubernetes/pki/ca.crt
discovery:
  bootstrapToken:
    apiServerEndpoint: kube-apiserver:6443
    token: abcdef.0123456789abcdef
    unsafeSkipCAVerification: true
  timeout: 5m0s
  tlsBootstrapToken: abcdef.0123456789abcdef
kind: JoinConfiguration
nodeRegistration:
  criSocket: /var/run/dockershim.sock
  name: k8smaster
  taints: null
```

注意这里打印的token并不是真实的token，而是`abcdef.0123456789abcdef`只表示占位符。

- --kubeconfig string

  指定配置文件的位置。默认`/etc/kubernetes/admin.conf`

### kubeadm config migrate

将旧的配置对象升级到最新版本支持的。

## 镜像

### kubeadm config images list

列出将要被使用的镜像

```
[root@k8smaster ~]# kubeadm config images list --image-repository registry.aliyuncs.com/google_containers --kubernetes-version 1.20.4
registry.aliyuncs.com/google_containers/kube-apiserver:v1.20.4
registry.aliyuncs.com/google_containers/kube-controller-manager:v1.20.4
registry.aliyuncs.com/google_containers/kube-scheduler:v1.20.4
registry.aliyuncs.com/google_containers/kube-proxy:v1.20.4
registry.aliyuncs.com/google_containers/pause:3.2
registry.aliyuncs.com/google_containers/etcd:3.4.13-0
registry.aliyuncs.com/google_containers/coredns:1.7.0
```

- --image-repository string

  默认`k8s.gcr.io`

- --kubernetes-version string

  默认stable-1

### kubeadm config images pull

拉取镜像

```
[root@k8smaster ~]# kubeadm config images pull --image-repository registry.aliyuncs.com/google_containers --kubernetes-version 1.20.4
```

- --image-repository string

  默认`k8s.gcr.io`

- --kubernetes-version string

  默认stable-1

