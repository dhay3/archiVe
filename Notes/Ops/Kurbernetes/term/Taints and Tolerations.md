# Taints and Tolerations

参考：

https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/

Node affinity

## Taints

在node上配置，作用于pods，告诉pods该如何keep away from nodes

1. 添加taint 到 node

   ```
   [root@k8smaster ~]# kubectl taint nodes k8sm k1=v1:NoSchedule
   node/k8sm tainted
   ```

2. 删除taint

   ```
   [root@k8smaster ~]# kubectl taint nodes k8sm k1=v1:NoSchedule-
   node/k8sm untainted
   ```

## Tolerations

在PodSpec中配置，作用于pods，告诉pods该如何schedule onto nodes 













