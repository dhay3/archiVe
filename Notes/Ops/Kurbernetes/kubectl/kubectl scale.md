# kubectl scale

对Deployment，ReplicaSet，Replication Controller，or StatefulSet资源大小进行缩放(可以增大与缩小)

syntax：`kubectl sacle [options] --replicas=<COUNT> TYPE RESOURCE_NAME`

将Deployment中的pod缩放到3

```
[root@k8smaster opt]# kubectl scale deployment deploy --replicas=3
deployment.apps/deploy scaled
[root@k8smaster opt]# kubectl get pods
NAME                      READY   STATUS              RESTARTS   AGE
deploy-78cfbdd995-9k9fn   1/1     Running             0          2m37s
deploy-78cfbdd995-c4mkg   0/1     ContainerCreating   0          6s
deploy-78cfbdd995-jrb6r   0/1     ContainerCreating   0          6s
kube-nginx                1/1     Running             0          17d
```

创建的pod的label一样，但是pod的名字不一样

```
[root@k8smaster opt]# kubectl get pod -o wide --show-labels 
NAME                      READY   STATUS    RESTARTS   AGE     IP               NODE        NOMINATED NODE   READINESS GATES   LABELS
deploy-78cfbdd995-9k9fn   1/1     Running   0          12m     192.168.16.138   k8smaster   <none>           <none>            pod-template-hash=78cfbdd995,pod=v1
deploy-78cfbdd995-c4mkg   1/1     Running   0          9m34s   192.168.16.139   k8smaster   <none>           <none>            pod-template-hash=78cfbdd995,pod=v1
deploy-78cfbdd995-jrb6r   1/1     Running   0          9m34s   192.168.16.140   k8smaster   <none>           <none>            pod-template-hash=78cfbdd995,pod=v1
kube-nginx                1/1     Running   0          17d     192.168.16.132   k8smaster   <none>           <none>            k1=v1
```

