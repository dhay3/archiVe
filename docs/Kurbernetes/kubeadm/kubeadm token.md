# kubeadm token

用于管理和生成token。`kube init`生成的token默认在24h之内有效。

## kubeadm token list

展示bootstrap tokens

syntax：`kubeadm token list [flags]`

```
[root@k8smaster ~]# kubeadm token list
TOKEN                     TTL         EXPIRES                     USAGES                   DESCRIPTION                                                               EXTRA GROUPS
dqaogx.binuaecfwxy7bvzg   23h         2021-03-25T14:56:25+08:00   authentication,signing   <none>                                                                    system:bootstrappers:kubeadm:default-node-token
g0ygrj.nyk997pydf7z8r41   <invalid>   2021-03-20T11:23:54+08:00   authentication,signing   The default bootstrap token g               enerated by 'kubeadm init'.   system:bootstrappers:kubeadm:default-node-token
gb2qkr.ky63vcu9dihgydjo   22h         2021-03-25T13:07:55+08:00   authentication,signing   <none>                                                                    system:bootstrappers:kubeadm:default-node-token
r0pk7i.g2havw423x3t3wws   20h         2021-03-25T11:57:02+08:00   authentication,signing   <none>                                                                    system:bootstrappers:kubeadm:default-node-token
s5lbyx.702z9ltlwm6u14n7   23h         2021-03-25T14:57:15+08:00   authentication,signing   <none>                                                                    system:bootstrappers:kubeadm:default-node-token
wfzsx7.dxqh7jh6tlzaubm5   1d          2021-03-26T14:59:42+08:00   authentication,signing   <none>                                                                    system:bootstrappers:kubeadm:default-node-token
```

- -o [text|json|yaml|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file]

  指定展示的格式，默认text。只展示token。

  ```
  [root@k8smaster ~]# kubeadm token list  -o json | egrep '"token"' | cut -d '"' -f 4
  g0ygrj.nyk997pydf7z8r41
  gb2qkr.ky63vcu9dihgydjo
  r0pk7i.g2havw423x3t3wws
  s5lbyx.702z9ltlwm6u14n7
  wfzsx7.dxqh7jh6tlzaubm5
  [root@k8smaster ~]# kubeadm token list  -o json | egrep '"token"' | awk -F '"' '{printf "%s ",$4}'
  r0pk7i.g2havw423x3t3wws s5lbyx.702z9ltlwm6u14n7 wfzsx7.dxqh7jh6tlzaubm5
  ```

## kubeadm token create

可以手动指定token。如果没有指定token，k8s自动生成一个token

syntax：`token create [token]`

```
[root@k8smaster ~]# kubeadm token create
dqaogx.binuaecfwxy7bvzg
```

- --print-join-command

  打印完整的kubeadm join命令

  ```
  [root@k8smaster ~]# kubeadm token create --print-join-command
  kubeadm join 192.168.80.201:6443 --token s5lbyx.702z9ltlwm6u14n7     --discovery-token-ca-cert-hash sha256:02d4a8241290e69442aa9ad929784d878d3fb51fa28b1743c38029f00de8ba43
  ```

- --ttl duration

  指定token的有效失效时间，0表示无期限。默认24h0m0s

  ```
  [root@k8smaster ~]# kubeadm token create --ttl 48h
  ```

- --dry-run

  token不会实际生效

- --description string

  token的description

## kubeadm token generate

随机生成一个token，==与create不同的是不能手动指定==

syntax：`kubeadm token generate [flags]`

```
[root@k8smaster ~]# kubeadm token create --description comments
4obquj.18cdtqcz2nt57gde
[root@k8smaster ~]# kubeadm token list
TOKEN                     TTL         EXPIRES                     USAGES                   DESCRIPTION                                                EXTRA GROUPS
4obquj.18cdtqcz2nt57gde   23h         2021-03-25T15:30:09+08:00   authentication,signing   comments                                                   system:bootstrappers:kubeadm:default-node-token
```

## kubeadm token delete

删除指定token

syntax：`kubeadm token delete [token-value ...]`

```
[root@k8smaster ~]# kubeadm token delete dqaogx.binuaecfwxy7bvzg
bootstrap token "dqaogx" deleted
```

删除所有token。==printf 左边必须带有空格==

```
[root@k8smaster ~]# kubeadm token list  -o json | egrep '"token"' | awk -F '"' '{printf " %s ",$4}' | xargs kubeadm token delete
bootstrap token "r0pk7i" deleted
bootstrap token "s5lbyx" deleted
bootstrap token "wfzsx7" deleted
```

- --dry-run



















