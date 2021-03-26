# kubeadm reset

Performs a best effort revert of changes made by `kubeadm init` or `kubeadm join`.

syntax：`kubeadm reset [flags]`

## follow

https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-reset/#synopsis

1. check
2. 从集群中删除这个节点
3. 删除节点，如果节点是control plane同时会删除etcd中的信息

## flags

- --cert-dir string

  指定证书存放的位置，默认`/etc/kubernetes/pki`

- -f | --force

- --kubeconfig string

  指定kubeconfig file的位置，默认`/etc/kubernetes/admin.conf`