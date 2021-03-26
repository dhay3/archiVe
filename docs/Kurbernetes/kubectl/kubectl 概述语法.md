# kubectl 概述语法

参考：

https://kubernetes.io/docs/reference/kubectl/overview/

kubectl cli 用来控制k8s集群。kubectl会去读取`~/.kube/config`配置

syntax：`kubectl [command] [type] [names] [flags]`

- command：

  对资源的操作，例如：`create`, `get`, `describe`, `delete`

- type：

  https://kubernetes.io/docs/reference/kubectl/overview/#resource-types

  指定资源的种类。case-insensitive，可以是单数，复数，缩写。以下命令等价

  ```
  [root@k8smaster .kube]# kubectl get node k8sm
  NAME   STATUS   ROLES                  AGE    VERSION
  k8sm   Ready    control-plane,master   5d5h   v1.20.4
  [root@k8smaster .kube]# kubectl get nodes k8sm
  NAME   STATUS   ROLES                  AGE    VERSION
  k8sm   Ready    control-plane,master   5d5h   v1.20.4
  [root@k8smaster .kube]# kubectl get no k8sm
  NAME   STATUS   ROLES                  AGE    VERSION
  k8sm   Ready    control-plane,master   5d5h   v1.20.4
  ```

- name：

  资源的名字

- flags

  optoinal flags

## 特殊

- 如果资源为同一类型，可以通过如下方式展示多个资源

  ```
  [root@k8smaster .kube]# kubectl get no k8sm k8snode01
  NAME        STATUS                     ROLES                  AGE     VERSION
  k8sm        Ready                      control-plane,master   5d5h    v1.20.4
  k8snode01   Ready,SchedulingDisabled   <none>                 3h36m   v1.20.4
  ```

- 如果资源不是同以类型

  ```
  [root@k8smaster .kube]# kubectl get node/k8sm pod/etcd-k8sm -n kube-system
  NAME        STATUS   ROLES                  AGE    VERSION
  node/k8sm   Ready    control-plane,master   5d5h   v1.20.4
  
  NAME            READY   STATUS    RESTARTS   AGE
  pod/etcd-k8sm   1/1     Running   0          5d5h
  ```

## 命令

https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands

https://kubernetes.io/docs/reference/kubectl/overview/#examples-common-operations













