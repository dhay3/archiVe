# Service

## 概述

### 为什么需要service ?

Pod并不是持久的资源，假设有一个 front-end pod想要调用 back-end pod 中提供的接口，但是 back-end pod 在不断的变化(Pod 的IP 同时也在改变)。front-end 如何知道具体的接口在哪一个IP呢？

这就引出了service。

### 概念

将一组Pods逻辑上抽象成一个service。通常由selector过滤pods，通过workload将这些pods划分为一个service。==kubernetes会为service提供一个虚拟的IP和DNS name==。

现在有1个front-end node，1个无状态的back-end node和3个replicas。当back-end node改变时(Pod 的IP 同时也在改变 )，front-end node无需关注Pod 的 IP，只要将请求发送给service 提供的IP。kubernetes会自动将请求发送给Pod。

## 创建

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

## virtual IPs and service proxies

每一个node中都有一个kube-proxy为service提供虚拟IP。kubernetes提供三种proxy mode

### user space proxy mode

1. kube-proxy 监听control plane用于添加和删除service的endpoint objects

2. 随机打开一个proxy port，连接通过这个proxy port发向backend pods

3. iptables根据clusterIP和port，==将流量重定向到proxy port==

4. kube-proxy使用round-robin 访问backend pods

5. 如果选择的第一个pod没有响应，会尝试其他的pod

   ![](D:\asset\note\imgs\_Kubernetes\Snipaste_2021-03-25_17-32-57.png)

### iptables proxy mode

> 默认使用这个proxy mode

1. kube-proxy 监听control plane用于添加和删除service的endpoint objects

2. iptables根据clusterIP和port将==流量重定向到backend pods==

3. kube-proxy随机选择一个backend pod

4. 如果第一个选择的pod没有响应，连接直接会失败

   ![](D:\asset\note\imgs\_Kubernetes\Snipaste_2021-03-25_17-48-09.png)

### IPVS proxy mode























