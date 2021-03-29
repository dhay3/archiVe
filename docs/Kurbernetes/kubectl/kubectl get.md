# kubectl get

https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#get

获取k8s resource的信息

syntax：`kubectl get [type] [flags]`

## flag

- --all-namespaces | -A

  展示所有namespaces的信息。默认输出defaults namespace。

  ```
  [root@k8smaster ~]# kubectl get pod -A
  NAMESPACE         NAME                                      READY   STATUS    RESTARTS   AGE
  calico-system     calico-kube-controllers-f95867bfb-rf7p2   1/1     Running   0          2d18h
  calico-system     calico-node-fnmwb                         0/1     Running   0          2d18h
  ```

- --show-labels

  同时打印label

  ```
  [root@k8smaster ~]# kubectl  get pod -n kube-system  --show-labels
  NAME                           READY   STATUS    RESTARTS   AGE     LABELS
  coredns-7f89b7bc75-c9mkr       1/1     Running   0          5d23h   k8s-app=kube-dns,pod-template-hash=7f89b7bc75
  coredns-7f89b7bc75-sg49w       1/1     Running   0          5d23h   k8s-app=kube-dns,pod-template-hash=7f89b7bc75
  ```

- ==--selector | -l==

  根据lable去过滤。支持`==`，`=`，`!=`

  ```
  [root@k8smaster ~]# kubectl get node -l  beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64
  NAME        STATUS                     ROLES                  AGE    VERSION
  k8sm        Ready                      control-plane,master   6d2h   v1.20.4
  ```

- --watch | -w

  监控资源的变更情况

- --watch-only

  只展示变更的资源信息

- --no-headers

  不输出tabel的header

- --output | -o

  格式化 输出

- --show-kind

  同时打印resource的type

  ```
  [root@k8smaster ~]# kubectl get pod -A --show-kind --no-headers  --server-print=false
  calico-system     pod/calico-kube-controllers-f95867bfb-rf7p2   2d19h
  ```

- --server-print

  打印READY,，STATUS，RESTARTS。缺省值