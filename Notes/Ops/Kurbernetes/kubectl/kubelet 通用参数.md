# kubelet 通用参数

参考：

https://kubernetes.io/docs/reference/kubectl/kubectl/

> 所有的通用参数都会被command继承

- --namespace | -n

  指定namespace

  ```
  [root@k8smaster ~]# kubectl get pod -A -n kube-system
  ```

- --insecure-skip-tls-verfiy

  跳过SSL验证，会让https不安全

- --kubeconfig

  指定kubectl读取的配置文件，默认`$HOME/.kube/config`

- --username string

  认证api-server的用户名

- --password string

  认证api-server的密码

- --request-timeout string

  请求超时时间，默认0表示不限制，可以使用单位

  ```
  [root@k8smaster ~]# kubectl --request-timeout 0.01s get pods --all-namespaces
  Unable to connect to the server: net/http: request canceled (Client.Timeout exceeded while awaiting headers)
  ```

- --server | -n string

  api-server的套接字

