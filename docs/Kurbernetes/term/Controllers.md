# Controllers

参考：

https://kubernetes.io/docs/concepts/architecture/controller/

Controllers are control loops that watch the state of your cluster. Each controller tries to move the curren cluster state closer to the desired state.

## Job

job controller 是kubernetes内建的controller，向一个pod或多个pod颁发task，只执行一次。

## Direct control

用于控制集群外部的状态