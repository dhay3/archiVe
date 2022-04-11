# Service

## 概述

### 为什么需要service ?

Pod并不是持久的资源，假设有一个 front-end pod想要调用 back-end pod 中提供的接口，但是 back-end pod 在不断的变化(Pod 的IP 同时也在改变)。front-end 如何知道具体的接口在哪一个IP呢？

这就引出了service。

### 概念

将一组Pods逻辑上抽象成一个service。通常==由selector过滤pods==，通过workload将这些pods划分为一个service。==kubernetes会为service提供一个虚拟的IP和DNS name==。

现在有1个front-end node，1个无状态的back-end node和3个replicas。当back-end node改变时(Pod 的IP 同时也在改变 )，front-end node无需关注Pod 的 IP，只要将请求发送给service 提供的IP。kubernetes会自动将请求发送给Pod。

## service创建配置

1. 假设有一组label为app=Myapp的pods并且监听TCP 9376端口

   ```
   [root@k8smaster opt]# cat service.yaml
   apiVersion: v1
   kind: Service
   metadata:
     name: my-service
   spec:
     selector:
       app: MyApp
     ports:
       - protocol: TCP
         port: 80
         targetPort: 9376
         
   [root@k8smaster opt]# kubectl apply -f service.yaml
   service/my-service created
   
   [root@k8smaster opt]# kubectl get service
   NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
   my-service   ClusterIP   10.107.113.177   <none>        80/TCP    3m39s
   ```

   - kubernetes会根据`spec`下的属性，创建一个对外开放80端口名字为my-service的service
   - kubernetes为这个serivce分配一个IP(Cluster-IP)，用于service proxies

## 代理模式

> 如果无需load balance 可以使用[headless service](https://kubernetes.io/docs/concepts/services-networking/service/#headless-services)

每一个node中都有一个kube-proxy为service提供虚拟IP。kubernetes提供三种proxy mode

### user space proxy mode

1. kube-proxy 监听control plane用于添加和删除service的endpoint objects

2. 随机打开一个proxy port，连接通过这个proxy port发向backend pods

3. iptables根据clusterIP和port，==将流量重定向到proxy port==

4. kube-proxy使用round-robin 访问backend pods

5. 如果选择的第一个pod没有响应，会尝试其他的pod

   ![Snipaste_2021-03-25_17-32-57](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2021-03-25_17-32-57.5n26i0ypt180.png)

### iptables proxy mode

> 默认使用这个proxy mode

1. kube-proxy 监听control plane用于添加和删除service的endpoint objects

2. iptables根据clusterIP和port将==流量重定向到backend pods==

3. kube-proxy随机选择一个backend pod

4. 如果第一个选择的pod没有响应，连接直接会失败

   ![Snipaste_2021-03-25_17-48-09](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2021-03-25_17-48-09.r8bndzf52dc.png)
   
   IPVS proxy mode

IPVS模式基于netfilter使用LVS规则，需要保证IPVS module可用

1. 比iptbales proxy mode延迟更加低。
2. 比其他的proxy mode 有更高的网络吞吐量
3. 支持多种负载算法
   - `rr`: round-robin
   - `lc`: least connection (smallest number of open connections)
   - `dh`: destination hashing
   - `sh`: source hashing
   - `sed`: shortest expected delay
   - `nq`: never queue

![Snipaste_2021-03-26_17-58-07](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2021-03-26_17-58-07.7emhfb2znns0.png)

​	==virtual server 和 real server是LVS中的术语==

## 服务发现

kubernetes支持两种发现方式：enviroment  varibales 和 DNS

### Enviroment variables

将端口和IP通过环境变量的方式导出，必须先创建service，调用service的pod才能发现

### DNS

使用dns add-on(CoreDNS)，例如：

有一个namespace为my-ns的service叫my-service。control plane 和 DNS server就会生成一条为my-service-my-ns 的DNS记录。在my-ns的命名空间中的pods，就可以通过my-service或my-service.my-ns来找到这个service。其他命名空间中的pods可以通过my-service.my-ns来找到这个service。

## 服务暴露

> 也可以通过[Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/)来暴露服务

kubernetes允许你修改`ServiceType`属性(默认为ClusterIP)来修改kubernetes service的type，支持如下几种：

### ClusterIP

暴露一个cluster-internal IP，这个service只能在cluster中使用，缺省值。

### NodePort

通过暴露pod的port做为一个service，==同时也会创建ClusterIP==。更加个性化的负载均衡方案。

==集群外和集群内都可以通过`<NodeIP>:<NodePort>`来访问，集群内部可以通过`<ClusterIP>:<Port>`来访问。==

可以通过`--nodeport-addresses`来指定kube-proxy只代理指定CIDR的IP。

control plane可以通过`--service-node-port-range`来指定分配的端口(需要30000-32767)。可以通过`.spec.ports[*].nodePort`属性来查看分配的端口。如果想要指定特定的端口，可以设置`nodePort`属性，需要辨别端口是否冲突。

可以通过`<NodeIP>:spec.ports[*].nodePort`和`.spec.clusterIP:spec.ports[*].port`来查看。

例如：

```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  type: NodePort
  selector:
    app: MyApp
  ports:
    - port: 8080
      targetPort: 80
      nodePort: 30007
```

port表示service暴露的port，targetPort表示pod监听的端口，nodePort表示==集群外部和集群内部都==可以通过`nodeIP:nodePort`访问到targetPort中的服务。

例如，service中有一个Pod A,Nginx监听了80端口。==集群内部==可以通过`ClusterIP:8080`端口访问到Pod A中的Nginx，集群外部可以通过`nodeIP:nodePort`访问到Nginx。

### LoadBalancer

参考:

https://help.aliyun.com/document_detail/182217.html?spm=5176.11065259.1996646101.searchclickresult.35784ec2aDl3Rp

通过云服务商的load balancer将service暴露。NodePort和ClusterIP负载均衡规则都是自动生成的，但是云服务商的负载均衡需要手动创建。

例如：

```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: MyApp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9376
  clusterIP: 10.0.171.239
  type: LoadBalancer
status:
  loadBalancer:
    ingress:
    - ip: 192.0.2.127
```

默认如果暴露了多个端口，所有的端口的protocol必须相同。可以设置kube-apiserver的`MixedProtocolLBService`

如果需要创建[内网负载均衡](https://kubernetes.io/docs/concepts/services-networking/service/#internal-load-balancer)

### ExternalName

将service映射到externalName对应的值，例如：externalName的值为`foo.bar.example.com`，就生成一条CNAME的记录对应这个值。

使用这种方式会导致，HTTP和HTTPS的host值错误，不推荐使用。

### 多端口暴露

如果service需要暴露多个端口，需要通过name属性来区别端口，必须是小写。

```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: MyApp
  ports:
    - name: http
      protocol: TCP
      port: 80
      targetPort: 9376
    - name: https
      protocol: TCP
      port: 443
      targetPort: 9377
```

## troubleshoot

https://kubernetes.io/docs/tasks/debug-application-cluster/debug-service/

当前有一个Deployment

```
[root@k8smaster opt]# cat deployment.yaml 
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deploy
  labels:
    deploy: v1
spec:
    selector:
        matchLabels:
            pod: v1
    template:
        metadata:
            labels:
                pod: v1
        spec:
            containers:
                - name: nginx
                  image: nginx
                  ports:
                    - containerPort: 80
                    - containerPort: 10086
```

缩放Deployment规格

```
[root@k8smaster opt]# kubectl scale deployment deploy --replicas=2
deployment.apps/deploy scaled
[root@k8smaster opt]# kubectl get pod -o wide
NAME                      READY   STATUS        RESTARTS   AGE   IP               NODE        NOMINATED NODE   READINESS GATES
deploy-78cfbdd995-9k9fn   1/1     Running       0          13m   192.168.16.138   k8smaster   <none>           <none>
deploy-78cfbdd995-c4mkg   1/1     Running       0          10m   192.168.16.139   k8smaster   <none>           <none>
deploy-78cfbdd995-jrb6r   0/1     Terminating   0          10m   <none>           k8smaster   <none>           <none>
kube-nginx                1/1     Running       0          17d   192.168.16.132   k8smaster   <none> 
```

按照label获取deployment中所有的Pod IP

```
[root@k8smaster opt]# kubectl get -l pod=v1 pod  -o go-template='{{range .items}}{{printf "%s\n" .status.podIP }}{{end}}' -A
192.168.16.138
192.168.16.139
```



## 例子0x001

1. 手动创建pod，

   ```
   [root@k8smaster opt]# cat pod.yaml
   apiVersion: v1
   kind: Pod
   metadata:
     name: kube-nginx
     labels:
        k1: v1
   spec:
     containers:
     - name: nginx
       image: nginx
   [root@k8smaster opt]# kubectl get pod -o wide --show-labels
   NAME         READY   STATUS    RESTARTS   AGE   IP               NODE        NOMINATED NODE   READINESS GATES   LABELS
   kube-nginx   1/1     Running   0          26m   192.168.16.132   k8smaster   <none>           <none>            k1=v1
   ```

2. 创建service

   ```
   [root@k8smaster opt]# cat service.yaml
   apiVersion: v1
   kind: Service
   metadata:
     name: my-service
   spec:
     type: NodePort
     selector:
       k1: v1
     ports:
       - name: nginx
         nodePort: 30001
         protocol: TCP
         port: 80
         targetPort: 80
   
   [root@k8smaster opt]# kubectl get service
   NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
   kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP        29m
   my-service   NodePort    10.108.67.230   <none>        80:30001/TCP   26m
   ```

3. cluster 外部主机访问`nodeIP:nodePort`

   ```
   C:\Users\82341>curl -I 192.168.80.201:30001
   HTTP/1.1 200 OK
   Server: nginx/1.19.6
   Date: Thu, 01 Apr 2021 07:52:40 GMT
   Content-Type: text/html
   Content-Length: 612
   Last-Modified: Tue, 15 Dec 2020 13:59:38 GMT
   Connection: keep-alive
   ETag: "5fd8c14a-264"
   Accept-Ranges: bytes
   ```

4. cluster 内部主机访问，`nodeIP:nodePort`和`clusterIP:port`都可以访问地到

   ```
   [root@k8snode01 ~]# curl -I 192.168.80.201:30001
   HTTP/1.1 200 OK
   Server: nginx/1.19.6
   Date: Thu, 01 Apr 2021 07:51:06 GMT
   Content-Type: text/html
   Content-Length: 612
   Last-Modified: Tue, 15 Dec 2020 13:59:38 GMT
   Connection: keep-alive
   ETag: "5fd8c14a-264"
   Accept-Ranges: bytes
   
   [root@k8snode01 ~]# curl -I 10.108.67.230:80
   HTTP/1.1 200 OK
   Server: nginx/1.19.6
   Date: Thu, 01 Apr 2021 07:55:32 GMT
   Content-Type: text/html
   Content-Length: 612
   Last-Modified: Tue, 15 Dec 2020 13:59:38 GMT
   Connection: keep-alive
   ETag: "5fd8c14a-264"
   Accept-Ranges: bytes
   ```

## 例子0x002







































