# JSONPath

参考：

https://kubernetes.io/docs/reference/kubectl/jsonpath/

https://support.smartbear.com/alertsite/docs/monitors/api/endpoint/jsonpath.html

https://goessner.net/articles/JsonPath/

JSONPath expression用于表示JSON对象中元素的路径

## syntax

JSONPath 可以使用两种方式，同时支持通配符

1. dot-notation

   `$.store.book[0].title`

2. bracket-notation

   `$['store']['book'][0]['title']`

| JSONPath              | 含义                                                   |
| --------------------- | ------------------------------------------------------ |
| `$`                   | 根节点，可以省略默认使用根节点                         |
| `@`                   | 在过滤表示中使用表示当前对象，类似于Java中的this关键字 |
| `.<name> or ['name']` | 子节点                                                 |
| `()`                  | 子表达式                                               |
| `name..`              | 递归搜索name中的元素                                   |
| `*`                   | 通配符                                                 |
| `?(<expression>)`     | 过滤表达式                                             |
| `[start:end]`         | 数组片段，类似于golang中的slice                        |
| `[<name>,<name>]`     | 表示一个或多个子节点或数组中的下标                     |

## 例子

```
[root@k8smaster ~]# kubectl config view
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

- 列出users下的所有元素

  ```
  [root@k8smaster ~]# kubectl config view -o jsonpath="{.users.*}"
  {"name":"kubernetes-admin","user":{"client-certificate-data":"REDACTED","client-key-data":"REDACTED"}}
  ```

- 列出前一个元素子级元素的user

  ```
  #users中的user
  [root@k8smaster ~]# kubectl config view -o jsonpath="{.users..user}"
  {"client-certificate-data":"REDACTED","client-key-data":"REDACTED"}
  
  #根元素中的user
  [root@k8smaster ~]# kubectl config view -o jsonpath="{$..user}"
  {"client-certificate-data":"REDACTED","client-key-data":"REDACTED"} kubernetes-admin[
  ```

- 过滤

  ```
  [root@k8smaster ~]# kubectl config view -o jsonpath='{..users[?(@.name=="kubernetes-admin")].user}'
  {"client-certificate-data":"REDACTED","client-key-data":"REDACTED"}
  ```

  

