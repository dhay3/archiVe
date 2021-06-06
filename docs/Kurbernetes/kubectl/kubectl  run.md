# kubectl  run

https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#run

在pod中运行指定的image，类似于docker run

syntax：`kubectl run NAME [flags]  `

## flags

- --image

  指定运行的镜像

- --stdin | -i

  保持stdin打开

- --tty | -t

  为这个pod中的所有container分配tty

- --attach

  运行容器，并attach，不能于`--restart=Never`一起使用

- --rm

  退出容器时，pod进入Terminating状态，然后删除pod

  ```
  [root@k8smaster opt]# kubectl get pod
  NAME         READY   STATUS        RESTARTS   AGE
  bbox         0/1     Terminating   0          6m26s
  
  ```

- --cascade

  级联删除的策略，支持background，orphan 或者 foreground

- --dry-run

- --command

  使用指定的command替换默认的指令

  ```
  [root@k8smaster opt]# kubectl run -it bbox --image=busybox --command -- sh
  ```

- --env

  指定command的环境变量

- --expose

  为生成的pod创建service









