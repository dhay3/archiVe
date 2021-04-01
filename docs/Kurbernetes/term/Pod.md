# Pod

参考：

https://kubernetes.io/docs/concepts/workloads/pods/

## 概述

pod是kubernetes中最小的管理和调度单位，由一组container组成。不是持久的资源，且不能对pod进行更新，可以使用workload动态创建和销毁pod，以及更新。

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

使用pod template(==也被称为PodSpec==)来创建和管理pods。例如：

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
        #busybox没有bash
        command: ['sh', '-c', 'echo "Hello, Kubernetes!" && sleep 3600']
      restartPolicy: OnFailure
    # The pod template ends here
```

修改pod template对已经创建好的pods不会生效。但是如果是workload创建的pod template，会被替换成新的。

## networking

![](D:\asset\note\imgs\_Kubernetes\Snipaste_2021-03-25_15-36-38.png)

1. kubernetes为pod提供唯一的IP
2. 同一Pod中的container通过localhost通信

3. 同一pod中的container共享namespace和port space

## 生命周期

pod有5个生命周期，通过STATUS字段表示

| Value       | Description                                                  |
| :---------- | :----------------------------------------------------------- |
| `Pending`   | The Pod has been accepted by the Kubernetes cluster, but one or more of the containers has not been set up and made ready to run. This includes time a Pod spends waiting to be scheduled as well as the time spent downloading container images over the network. |
| `Running`   | The Pod has been bound to a node, and all of the containers have been created. At least one container is still running, or is in the process of starting or restarting. |
| `Succeeded` | All containers in the Pod have terminated in success, and will not be restarted. |
| `Failed`    | All containers in the Pod have terminated, and at least one container has terminated in failure. That is, the container either exited with non-zero status or was terminated by the system. |
| `Unknown`   | For some reason the state of the Pod could not be obtained. This phase typically occurs due to an error in communicating with the node where the Pod should be running. |

## 安全策略



## 0x01 例子

1. 创建pod，==注意buxybox没有bash==

   ```
   [root@k8smaster opt]# cat pod.yaml
   apiVersion: v1
   kind: Pod
   metadata:
     name: pod-example
   spec:
     containers:
     - name: busybox #镜像的名字
       image: busybox #对应仓库中的镜像名
       command: ["echo"] #相当于Dockerfile中的CMD
       args: ["Hello World"]
       
   [root@k8smaster opt]# kubectl apply -f pod.yaml
   pod/pod-example configured
   ```

2. 查看状态，这里由于容器会退出，所以状态是CrashLoopBackoff

   ```
   [root@k8smaster opt]# kubectl get pods
   NAME          READY   STATUS             RESTARTS   AGE
   pod-example   0/1     CrashLoopBackOff   4          3m26s
   ```

3. 使用`kubectl logs`查看

   ```
   [root@k8smaster opt]# kubectl logs pod/pod-example
   Hello World
   ```

## 0x02 例子

1. 

