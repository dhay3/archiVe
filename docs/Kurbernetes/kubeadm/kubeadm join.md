# kubeadm join

initializes a Kubernetes worker node and joins it to the cluster.



syntax：`kubeadm join [api-server-endpoint] [flags]`

api-server-endpoint 标识apiserver的套接字

## 流程

 [detail of execute follow](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-join/)

## options



如果discovery information 

- --control-plane

  在当前的node创建一个control plane

- --discovery-file string

  file-base discovery 可以是URL或file。如果是URL，schema必须是HTTPS

- discovery-token string

  token-based discovery。用于从apiserver上拉取用于校验cluster的信息。

- --token

- --discovery-token-ca-cert-hash

  token-based discovery。校验根证书的hash





















