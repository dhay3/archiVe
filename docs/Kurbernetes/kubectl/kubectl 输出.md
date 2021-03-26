# kubectl 输出

> jsonpath-expression
>
> https://kubernetes.io//docs/reference/kubectl/jsonpath/

## 格式

syntax：`kubectl [command] [type] [name] -o <output_format>`

- wide

  额外输出信息

  ```
  [root@k8smaster .kube]# kubectl get nodes -o wide
  NAME        STATUS                     ROLES                  AGE     VERSION   INTERNAL-IP      EXTERNAL-IP   OS-IMAGE                KERNEL-VERSION           CONTAINER-RUNTIME
  k8sm        Ready                      control-plane,master   5d5h    v1.20.4   192.168.80.201   <none>        CentOS Linux 7 (Core)   3.10.0-1062.el7.x86_64   docker://1.13.1
  k8snode01   Ready,SchedulingDisabled   <none>                 3h41m   v1.20.4   192.168.80.202   <none>        CentOS Linux 7 (Core)   3.10.0-1062.el7.x86_64   docker://1.13.1
  ```

- json

  以json的格式输出

  ```
  [root@k8smaster .kube]# kubectl get nodes -o json | less
  ```

- yaml

  ```
  [root@k8smaster .kube]# kubectl get nodes -o yaml | less
  ```

- name

  只输出名字

  ```
  [root@k8smaster .kube]# kubectl get nodes -o name
  node/k8sm
  node/k8snode01
  ```

- custom-columns

  以`header:json-path-expr`的格式执行显示列

  ```
  [root@k8smaster .kube]# kubectl get nodes -o custom-columns=name:.metadata.name
  name
  k8sm
  k8snode01
  
  [root@k8smaster ~]# kubectl get pod -o custom-columns=name:.metadata.name --all-namespaces
  name
  calico-kube-controllers-f95867bfb-rf7p2
  calico-node-fnmwb
  calico-node-ft729
  calico-typha-77db86ffbc-mncfg
  coredns-7f89b7bc75-c9mkr
  coredns-7f89b7bc75-sg49w
  etcd-k8sm
  kube-apiserver-k8sm
  kube-controller-manager-k8sm
  kube-proxy-mlj2z
  kube-proxy-p6dwn
  kube-scheduler-k8sm
  tigera-operator-675ccbb69c-mj7nr
  ```

- go-template

  ```
  [root@k8smaster ~]# kubectl get pod -o go-template --template '{{range .itmes}}{{.metadata.name}}{{"\n"}}{{end}}'
  ```

## 排序

syntax：`kubectl [command] [TYPE] [NAME] --sort-by=<jsonpath_exp>`

```
[root@k8smaster .kube]# kubectl get pod --sort-by={.metadata.name} --all-namespaces
NAMESPACE         NAME                                      READY   STATUS    RESTARTS   AGE
calico-system     calico-kube-controllers-f95867bfb-rf7p2   1/1     Running   0          2d
calico-system     calico-node-fnmwb                         0/1     Running   0          2d
calico-system     calico-node-ft729                         0/1     Running   0          4h
calico-system     calico-typha-77db86ffbc-mncfg             1/1     Running   0          2d
kube-system       coredns-7f89b7bc75-c9mkr                  1/1     Running   0          5d5h
kube-system       coredns-7f89b7bc75-sg49w                  1/1     Running   0          5d5h
kube-system       etcd-k8sm                                 1/1     Running   0          5d5h
```

















