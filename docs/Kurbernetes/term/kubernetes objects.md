# kubernetes objects

参考：

https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/

kubernetes objects(pods,service,etc)可以通过yaml格式来定义，然后通过`kubectl apply`来创建kubernete objects。

## 配置

> 严格区分单双引号，可以参考dockerfile

详细的配置可以查看：

https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#-strong-api-overview-strong-

### required fields

- apiVersion：创建object的kubernetes API version
- kind：type of objects
- [metadata][https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta]：唯一区别这个objects的属性，包括name，UID，namespace
- spec：定义objects的属性

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 2 # tells deployment to run 2 pods matching the template
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx #unique
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```

## container block

- containers slice[map]

  指定pod中运行的container

- name string

  创建container后的名字，唯一标识符。

- image string

  based container

- command | args

- ports  slice[map]

  ==只做提示用和docker expose相同==，会在`kubectl get pod -o json`显示

  ```
  [root@k8smaster opt]# kubectl get pod/kube-nginx-80 -o custom-columns=port:.metadata.annotations.*
  port
  192.168.132.52/32,192.168.132.52/32,{"apiVersion":"v1","kind":"Pod","metadata":{"annotations":{},"name":"kube-nginx-80","namespace":"default"},"spec":{"containers":[{"image":"nginx","name":"nginx","ports":[{"containerPort":80}]}]}}
  ```

- workingdir string

  

