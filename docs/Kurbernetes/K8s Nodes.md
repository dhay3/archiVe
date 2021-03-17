# K8s Nodes

Node是一个由control plane 管理的节点，包含kubelet，kube-proxy，pods。可以是一台虚拟机或是物理机

```mermaid
classDiagram
class Node{
+kubelet
+kube-proxy
+pods
}
```

