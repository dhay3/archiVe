# kubernetes objects

参考：

https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/

kubernetes objects(pods,service,etc)可以通过yaml格式来定义，然后通过`kubectl apply`来创建kubernete objects。

## 配置

### required fields

- apiVersion：创建object的kubernetes API version
- kind：type of objects
- metadata：唯一区别这个objects的属性，包括name，UID，namespace
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
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```

详细的配置可以查看：

https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#-strong-api-overview-strong-

