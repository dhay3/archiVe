# kubelet

参考：

https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/

kubelet is the primary "node agent" that runs on each node.

kubelet读取PodSpec来管理container的运行和健康状态，kubelet不会管理不是由kubernetes创建的容器。

配置文件存储在`/etc/kubernetes/kubelet.conf`