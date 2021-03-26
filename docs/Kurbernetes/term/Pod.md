# Pod

参考：

https://kubernetes.io/docs/concepts/workloads/pods/

## 概述

pod是kubernetes中最小的管理和调度单位，由一组container组成。不是持久的资源，可以使用workload动态创建和销毁pod。

==container共享storage，network and a specification for how to run the containers.==

类似一组Docker containers共享namespaces 和 shared filesystem volumes。

## 运行模式

- **one container per pod**

  pod是container的外壳，kubernetes不是直接管理container而是通过pod

- **co-located containers per pod**

  多个关联的container在同一个pod

## 如何管理多个容器

kubernetes可以使用controller来自动管理pod，也被称为==workload==。用于pod自动失败修复和推出。

例如当一个Node挂掉了，workload就会为这个Node上的pod创建一个替代的Pod，然后将这些Pod放在一个健康的Node中。

不同的workload功能不同，常见的workload有：

1. Deployment
2. StatefulSet
3. DaemonSet

## pod template

workload使用pod template来创建和管理pods。例如：

```
apiVersion: batch/v1
kind: Job
metadata:
  name: hello
spec:
  template:
    # This is the pod template
    spec:
      containers:
      - name: hello
        image: busybox
        command: ['sh', '-c', 'echo "Hello, Kubernetes!" && sleep 3600']
      restartPolicy: OnFailure
    # The pod template ends here
```

## networking

![](D:\asset\note\imgs\_Kubernetes\Snipaste_2021-03-25_15-36-38.png)

1. kubernetes为pod提供唯一的IP
2. 同一Pod中的container通过localhost通信

3. 同一pod中的container共享namespace和port space











