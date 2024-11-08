# kubeadm join

initializes a Kubernetes worker node and joins it to the cluster.

syntax：`kubeadm join [api-server-endpoint] [flags]`

api-server-endpoint 标识apiserver的套接字

## 流程

 [detail of execute follow](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-join/)

1. kubeadm 从control plane 上下载集群的信息。使用bootstrap token 和 CA key hash 校验信息。
2. 信息校验成功后，kubelet启用TLS bootstrapping process。==kubelet通过token获得api server的授权并向control plane提交csr并签名，control plane 为csr生成证书，kubelet从control plane上下载证书==
3. kubeadm配置kubelet并于api server连接。

如果是control plane需要额外的step

## options

- --control-plane

  在当前的node创建一个control plane

- --discovery-file string

  file-base discovery 可以是URL或file。如果是URL，schema必须是HTTPS

  ```
  kubeadm join --discovery-file https://url/file.conf
  kubeadm join --discovery-file path/to/file.conf
  ```

- --discovery-token string

  token-based discovery。用token去api server拉取信息校验集群。

- --tls-bootstrap-token string

  与control plane 约定好用于握手的token

- ==--discovery-token-ca-cert-hash==

  token-based discovery。校验根证书公钥的hash

- ==--token string==

  使用指定的token做discovery-token 和 tls-bootstrap-token，如果这两个参数都没有指定。

- --node-name string

  指定加入到集群后节点的名字

## 验证方式

k8s有两种验证方式，tokenbased

### token based

如果使用token based authentication，需要指定discovery-token-ca-cert-hash来校验control plane的CA 根证书(会将control plane的CA证书下载)。

```
kubeadm join 1.2.3.4:6443 --discovery-token abcdef.1234567890abcdef --discovery-token-ca-cert-hash sha256:1234..cdef 
```

可以通过`kubeadm token create --print-join-command`提示如何加入到control plane

```
[root@k8smaster ~]# kubeadm token create --print-join-command
kubeadm join 192.168.80.201:6443 --token gb2qkr.ky63vcu9dihgydjo     --discovery-token-ca-cert-hash sha256:02d4a8241290e69442aa9ad929784d878d3fb51fa28b1743c38029f00de8ba43
```

### file or https based

```
kubeadm join --discovery-file path/to/file.conf
kubeadm join --discovery-file https://url/file.conf
```

## 安全安装

1. 关闭control plane自动证书签名

   control plane 默认自动对kublet的csr签名生成证书，可以在control plane使用如下命令关闭该功能

   ```
   kubectl delete clusterrolebinding kubeadm:node-autoapprove-bootstrap
   ```

   使用`kubeadm join`就会被拦截(需要在一定时间后才会报错)，需要在control plane手动执行如下命令

   ```
   [root@k8smaster ~]# kubectl get csr
   NAME        AGE     SIGNERNAME                                    REQUESTOR                 CONDITION
   csr-cxb62   4m57s   kubernetes.io/kube-apiserver-client-kubelet   system:bootstrap:gb2qkr   Pending
   
   [root@k8smaster ~]# kubectl certificate approve csr-cxb62
   certificatesigningrequest.certificates.k8s.io/csr-cxb62 approved
   
   [root@k8smaster ~]# kubectl get csr
   NAME        AGE     SIGNERNAME                                    REQUESTOR                 CONDITION
   csr-cxb62   8m25s   kubernetes.io/kube-apiserver-client-kubelet   system:bootstrap:gb2qkr   Approved,Issued
   ```

## troubleshoot

- User "system:anonymous" cannot create resource "nodes" in API group "" at the cluster scope

  https://github.com/kelseyhightower/kubernetes-the-hard-way/issues/438








































