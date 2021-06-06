# kubeadm upgrade

是一个复合的命令

## kubeadm upgrade plan

用于检查集群的组件是否可以升级

syntax：`kubeadm upgrade plan [flags]`

```
[root@k8smaster ~]# kubeadm upgrade plan
...
Upgrade to the latest version in the v1.20 series:

COMPONENT                 CURRENT    AVAILABLE
kube-apiserver            v1.20.4    v1.20.5
kube-controller-manager   v1.20.4    v1.20.5
kube-scheduler            v1.20.4    v1.20.5
kube-proxy                v1.20.4    v1.20.5
CoreDNS                   1.7.0      1.7.0
etcd                      3.4.13-0   3.4.13-0
...
```

- ignore-preflight-errors

  指定不展示的错误信息

  ```
  kubeadm upgrade plan --ignore-preflight-errors all
  ```

## kubeadm upgrade apply

升级集群(所有的组件)到指定版本

syntax：`kubeadm upgrad apply [version] [flags]`

```
kubeadm upgrade apply -f  v1.20.5
```

- --dry-run 

  模拟

- -f 

  force

- -y

  yes

## kubeadm upgrade diff

查看升级的更新

syntax：`kubeadm upgrade diff [version] [flags]`

```
[root@k8smaster ~]# kubeadm upgrade diff   v1.20.5
[upgrade/diff] Reading configuration from the cluster...
[upgrade/diff] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -o yaml'
--- /etc/kubernetes/manifests/kube-apiserver.yaml
+++ new manifest
@@ -41,7 +41,7 @@
     - --service-cluster-ip-range=10.96.0.0/12
     - --tls-cert-file=/etc/kubernetes/pki/apiserver.crt
     - --tls-private-key-file=/etc/kubernetes/pki/apiserver.key
-    image: registry.aliyuncs.com/google_containers/kube-apiserver:v1.20.4
+    image: registry.aliyuncs.com/google_containers/kube-apiserver:v1.20.5
     imagePullPolicy: IfNotPresent
     livenessProbe:
       failureThreshold: 8
--- /etc/kubernetes/manifests/kube-controller-manager.yaml
```

## kubeadm upgrade node

升级集群中的node包括kubelet的配置，etcd，如果节点是control plane同样也会被升级

syntax：`kubeadm upgrade node [flags]`

- --dry-run

- --etcd-upgrade 

  升级etcd，默认true



































