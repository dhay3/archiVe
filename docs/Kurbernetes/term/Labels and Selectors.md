# Labels and Selectors

参考：

https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#motivation

## Label

labels是map类型，用于标识一个kubernetes object(Pods，Services，etc)。

```
[root@k8smaster opt]# kubectl apply -f pod.yaml
pod/label-demo created
[root@k8smaster opt]# cat pod.yaml

apiVersion: v1
kind: Pod
metadata:
  name: label-demo
  labels:
    environment: production
    app: nginx
spec:
  containers:
  - name: nginx
    image: nginx:1.14.2
    ports:
    - containerPort: 80
```

使用上面pod的配置文件创建出来的pod有两个标签，`environment:prodution`和`app:nginx`

```
[root@k8smaster opt]# kubectl get pods --show-labels
NAME                                READY   STATUS    RESTARTS   AGE   LABELS
label-demo                          1/1     Running   0          27s   app=nginx,environment=production
```

## Label selector

label selector用于定义a set of kubernetes objects，当前支持两种selectors：equality-based and set-based。

### Equality-based 

可以使用`=`,`==`,`!=`

```
[root@k8smaster opt]# kubectl get pods -l app=nginx --show-labels
NAME                                READY   STATUS    RESTARTS   AGE    LABELS
label-demo                          1/1     Running   0          16m    app=nginx,environment=production
nginx-deployment-559d658b74-4n625   1/1     Running   0          104m   app=nginx,pod-template-hash=559d658b74
nginx-deployment-559d658b74-h8v75   1/1     Running   0          104m   app=nginx,pod-template-hash=559d658b74
nginx-deployment-559d658b74-xnh79   1/1     Running   0          104m   app=nginx,pod-template-hash=559d658b74
[root@k8smaster opt]# kubectl get pods -l app!=nginx --show-labels
No resources found in default namespace.
```

### Set-based

使用in,notin,exsits，表达式需要在引号内

```
[root@k8smaster opt]# kubectl get pods -l 'app in (nginx)' --show-labels
NAME                                READY   STATUS    RESTARTS   AGE    LABELS
label-demo                          1/1     Running   0          18m    app=nginx,environment=production
nginx-deployment-559d658b74-4n625   1/1     Running   0          106m   app=nginx,pod-template-hash=559d658b74
nginx-deployment-559d658b74-h8v75   1/1     Running   0          106m   app=nginx,pod-template-hash=559d658b74
nginx-deployment-559d658b74-xnh79   1/1     Running   0          106m   app=nginx,pod-template-hash=559d658b74
[root@k8smaster opt]# kubectl get pods -l 'app notin (nginx)' --show-labels
No resources found in default namespace.


[root@k8smaster opt]# kubectl get pods -l 'environment in (production),tier in (frontend)'
```

## service

service也可以使用label来标识

```
selector:
    component: redis
```

通过matchLabels来匹配

```
selector:
  matchLabels:
    component: redis
  matchExpressions:
    - {key: tier, operator: In, values: [cache]}
    - {key: environment, operator: NotIn, values: [dev]}
```

















