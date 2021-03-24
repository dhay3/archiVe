# kubeadm join

initializes a Kubernetes worker node and joins it to the cluster.



syntax：`kubeadm join [api-server-endpoint] [flags]`

api-server-endpoint 标识apiserver的套接字

## 流程

 [detail of execute follow](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-join/)

1. kubeadm 从api server  上下载集群的信息。使用bootstrap token 和 CA key hash 校验信息。
2. 信息校验成功后，kubelet启用TLS bootstrapping process。通过token获得api server的授权并向control plane提交csr
3. kubeadm配置kubelet并于api server连接

## options



如果discovery information 

- --control-plane

  在当前的node创建一个control plane

- --discovery-file string

  file-base discovery 可以是URL或file。如果是URL，schema必须是HTTPS

- --discovery-token string

  token-based discovery。用token去api server拉取信息校验集群。

- tls-bootstrap-token string

  与control plane 约定好用于握手的token

- ==--discovery-token-ca-cert-hash==

  token-based discovery。校验根证书公钥的hash

- ==--token string==

  使用指定的token做discovery-token 和 tls-bootstrap-token，如果这两个参数都没有指定。

- --node-name string

  指定加入到集群后节点的名字





















