# 配置文件

kubeconfig只是一个通用的名字，并不代表配置文件就叫做kubeconfig，用来定义cluster，namespace，authentication等机制。

kubectl默认会读取`$HOME/.kube/config`，可以使用`--kubeconfig`来指定使用其他的配置文件，也可设置`KUBECONFIG`环境变量。

## 查看配置文件

通过这个配置文件可以观察到有几个属性都是list类型，说明可以配置过个

```
[root@k8smaster .kube]# kubectl config view
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

