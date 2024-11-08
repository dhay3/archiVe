# Deployments

参考：

https://kubernetes.io/docs/concepts/workloads/controllers/deployment/

管理Pods和ReplicaSets，通常有如下几种情况可以使用Deployment

1. 创建RepicaSets
2. 更新Pods,只有在`.spec.template`中的内容改变时才会创建一个新的版本
3. 回滚到之前的Deployment，只回滚`.spec.template`中的内容
4. 暂停Deployment
5. 提升Deployment的规格
6. 删除旧的ReplicaSets

例如：

```
apiVersion: apps/v1
kind: Deployment
metadata:
#deployment的名字
  name: nginx-deployment
  #创建dedeployment的标签
  labels:
    app: nginx
spec:
	#创建template replica的数量，默认为1
  replicas: 3
  #选中含有特定标签的pod，必须于template中的label相同
  selector:
    matchLabels:
      app: nginx
  #创建pod的模板
  template:
    metadata:
    #创建pod的标签
      labels:
        app: nginx
    spec:
    #模板使用的镜像
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
        - containerPort: 10086
```

## 创建

> 由deployment创建的replicaset或pod，默认以`[DEPLOYMENT-NAME]-[RANDOM-STRING]`命名

1. 创建一个deployment

   ```
   kubectl apply -f https://k8s.io/examples/controllers/nginx-deployment.yaml
   ```

2. 查看deployment

   ```
   [root@k8smaster ~]# kubectl get deployments
   NAME               READY   UP-TO-DATE   AVAILABLE   AGE
   nginx-deployment   3/3     3            3           3d22h
   
   #详细信息
   [root@k8smaster ~]# kubectl describe deployments/nginx-deployment
   ```

3. 查看replicaset

   ```
   [root@k8smaster ~]# kubectl get rs
   NAME                          DESIRED   CURRENT   READY   AGE
   nginx-deployment-559d658b74   3         3         3       18m
   ```

4. 查看由deployment创建的pods

   ```
   [root@k8smaster ~]# kubectl get pods --show-labels -o wide
   NAME                                READY   STATUS    RESTARTS   AGE   IP               NODE   NOMINATED NODE   READINESS GATES   LABELS
   nginx-deployment-559d658b74-4n625   1/1     Running   0          19m   192.168.132.28   k8sm   <none>           <none>            app=nginx,pod-template-hash=559d658b74
   nginx-deployment-559d658b74-h8v75   1/1     Running   0          19m   192.168.132.27   k8sm   <none>           <none>            app=nginx,pod-template-hash=559d658b74
   nginx-deployment-559d658b74-xnh79   1/1     Running   0          19m   192.168.132.29   k8sm   <none>           <none>            app=nginx,pod-template-hash=559d658b74
   ```

   pod-template-hash标签是由deployment自动生成的

## 更新

1. 修改镜像`nginx:1.14.2`到`nginx:1.16.1`。`--record`记录变更信息

   ```
   kubectl set image deployment/nginx-deployment nginx=nginx:1.16.1 --record
   ```

   或

   使用vim修改

   ```
   kubectl edit deployment nginx-deployment --record
   ```


## 回滚

1. 回滚

   ```
   kubectl rollout undo deployment/nginx-deployment
   ```

2. 查看

   ```
   [root@k8smaster ~]# kubectl rollout undo deployment nginx-deployment --to-revision=1
   ```


## 查看回滚历史

```
[root@k8smaster ~]# kubectl rollout history deployment "nginx-deployment"
deployment.apps/nginx-deployment
REVISION  CHANGE-CAUSE
1         kubectl deployment.apps/nginx-deployment set image deployment.v1.apps/nginx-deployment nginx=nginx:1.16.1 --record=true
```

