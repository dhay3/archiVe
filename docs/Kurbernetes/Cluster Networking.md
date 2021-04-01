# Cluster Networking

[check the detail](https://kubernetes.io/docs/concepts/cluster-administration/networking/)

kubernetes 网络主要有5个问题：

1. container to container：由pods和localhost解决
2. pod to pod：这是本篇文章主要解决的问题
3. pod to service：由service解决
4. external to service：由service解决
5. external to pod：由port-forward或service解决

## pod to pod

每个pod在node中都会分配一个唯一的IP，不需要手动处理容器的端口和主机映射。

==在pod中的container共享namespace和IP address 和 MAC address，所以在同一个pod中的container可以通过localhost通信。也意味着共享port pool==

同时遵循如下规则：

1. pods on a node can communicate with all pods on all nodes without NAT
2. agents on a node (e.g. system daemons, kubelet) can communicate with all pods on that node
3. pods in the host network of a node can communicate with all pods on all nodes without NAT