# kubelet 通用参数

> 所有的通用参数都会被继承

- --namespace | -n

  指定namespace

  ```
  [root@k8smaster ~]# kubectl get pod -A -n kube-system
  ```

  