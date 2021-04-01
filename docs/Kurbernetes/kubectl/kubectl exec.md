# kubectl exec

和docker exec 类似，在指定pod中执行命令(可以用于校验pod是否存活)

syntax：`kubectl exec <type/name> [-c container] -- COMMAND`

> kuberntes创建container是不会分配tty，需要使用`-it`参数

```
[root@k8smaster opt]# kubectl exec kube-nginx -- sh -i
sh: 0: can't access tty; job control turned off

[root@k8smaster opt]# kubectl exec -it kube-nginx -- sh
# ls
bin   docker-entrypoint.d   home   media  proc  sbin  tmp
boot  docker-entrypoint.sh  lib    mnt    root  srv   usr
dev   etc                   lib64  opt    run   sys   var
```

获取日期

```
[root@k8smaster opt]# kubectl exec kube-nginx -- date
Thu Apr  1 07:29:43 UTC 2021
```

