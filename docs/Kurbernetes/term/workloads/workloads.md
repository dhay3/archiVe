# workloads

由于pod有不同的生命周期，一个一个去管理pod会比较复杂。这时候就可以使用workload去管理一组pods，也被称为pod控制器。

kubernetes 提供一组内置的workload resources：

1. Deployment 和 ReplicaSet(替代了传统的ReplicationController)。

   管理无状态的应用，在Deployment中的pod是可以被替换的。

2. StatefulSet

   管理有状态的应用

3. DaemonSet

   管理用于node-local的pod

4. Job 和 CronJob