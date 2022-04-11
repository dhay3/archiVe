# ReplicaSet

参考：

https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/

a stable set of replica.

ReplicaSet用于创建和删除pod到指定数目，当需要创建时使用Pod template。

## when to use a ReplicaSet

ReplicaSet确保pod在运行过程中的数量，但是ReplicaSet由Deployment管理。所以更推荐使用Deployment去创建，除非pod不需要自定义更新。

例如：

```

```

