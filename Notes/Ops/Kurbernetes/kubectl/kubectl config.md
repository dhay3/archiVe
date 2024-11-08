# kubectl config

查看或修改kubernetes配置

1. 当kubectl 使用`--kubeconfig`时使用指定的配置文件
2. 当设置了`$KUBECONFIG`环境变量使用指定的配置文件，可以有多个冒号分隔
3. 否则默认使用`$HOME/.kube/config`

syntax：`kubectl [flag] config <command> `

## command

[check details here](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#config)

- view

  查看使用的配置文件

  ```
  [root@k8smaster opt]# kubectl config view
  apiVersion: v1
  clusters:
  - cluster:
      certificate-authority-data: DATA+OMITTED
      server: https://192.168.80.201:6443
    name: kubernetes
  contexts:
  - context:
      cluster: kubernetes
      user: kubernetes-admin
    name: kubernetes-admin@kubernetes
  current-context: kubernetes-admin@kubernetes
  kind: Config
  preferences: {}
  users:
  - name: kubernetes-admin
    user:
      client-certificate-data: REDACTED
      client-key-data: REDACTED
  ```

- current-context

  ```
  [root@k8smaster opt]# kubectl config current-context
  kubernetes-admin@kubernetes
  ```

- get-context

  ```
  [root@k8smaster opt]# kubectl config get-contexts
  CURRENT   NAME                          CLUSTER      AUTHINFO           NAMESPACE
            kc
  *         kubernetes-admin@kubernetes   kubernetes   kubernetes-admin
  ```

- set-context

  按照name来区分，如果不存在就创建，如果存在就修改

  ```
  [root@k8smaster opt]# kubectl config set-context  kname --cluster defualt --user kuser
  Context "kname" created.
  ```

- delete-context

  ```
  [root@k8smaster opt]# kubectl config delete-context kname
  deleted context kname from /root/.kube/config
  ```

- use-context

  ```
  [root@k8smaster opt]# kubectl config use-context kc
  Switched to context "kc".
  [root@k8smaster opt]# kubectl config current-context
  kc
  ```

- set-credentials

  [添加或修改用户](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#-em-set-credentials-em-)

- delete-user

  ```
  [root@k8smaster opt]# kubectl config delete-user kubernetes-admin
  deleted user kubernetes-admin from /root/.kube/config
  ```

  



























