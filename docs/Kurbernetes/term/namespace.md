# namespace

参考：

https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/

## 概述

> 创建的namespace不能以`kube-`开头，这些namespace为kubernetes system namespaces

namespace用于分隔多用户的cluster resources。==也可以使用label来分别不同得resources。==

==any resource can only be in one namespace== 

kubernetes启动时有四个初始的namespaces：

1. defaults：如果没有指定namespace默认使用该namespace

   ```
   [root@k8smaster ~]# kubectl get node -n defaults
   NAME        STATUS                     ROLES                  AGE   VERSION
   k8sm        Ready                      control-plane,master   6d    v1.20.4
   k8snode01   Ready,SchedulingDisabled   <none>                 22h   v1.20.4
   [root@k8smaster ~]# kubectl get node
   NAME        STATUS                     ROLES                  AGE   VERSION
   k8sm        Ready                      control-plane,master   6d    v1.20.4
   k8snode01   Ready,SchedulingDisabled   <none>                 22h   v1.20.4
   ```

2. kube-system：kubernetes system 创建的namespace

3. kube-public：在该namespace下的资源可以被公共读取和使用

4. kube-node-lease：lease objects的namespace

namespace也是一种resource，可以通过`kubectl get`命令获取

```
[root@k8smaster ~]# kubectl get namespaces
NAME              STATUS   AGE
calico-system     Active   2d19h
default           Active   6d
kube-node-lease   Active   6d
kube-public       Active   6d
kube-system       Active   6d
tigera-operator   Active   2d19h
```

> 并不是所有的resource都在namespace中

- 使用`kubectl api-resources --namespaced=true`来查看所有在namespace中的resource

- 使用`kubectl api-resources --namespaced=flase`来查看所有不在namespace中的resource

==nodes不在namespace中所以指定namespace对node不生效。==

```
[root@k8smaster ~]# kubectl get node -n defaults
NAME        STATUS                     ROLES                  AGE   VERSION
k8sm        Ready                      control-plane,master   6d    v1.20.4
k8snode01   Ready,SchedulingDisabled   <none>                 22h   v1.20.4
[root@k8smaster ~]# kubectl get node -n kube-system
NAME        STATUS                     ROLES                  AGE   VERSION
k8sm        Ready                      control-plane,master   6d    v1.20.4
k8snode01   Ready,SchedulingDisabled   <none>                 22h   v1.20.4
```





