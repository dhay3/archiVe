# kubectl port-forward

用于转发多个或一个本地端口到pod，这个命令需要安装`socat`

syntax：`kubectl port-forward type/source_name [localport:]remote_port ...`

- 监听本地5000端口，随机端口转发

  ```
  kubectl port-forward pod/mypod :5000
  ```

- 监听本地10086端口，单端口转发

  ```
  kubectl port-forward pod/mypod 10086:5000
  ```

- 多端口转发

  本地监听5000和6000，转发到pod中的5000和6000

  ```
  kubectl port-forward pod/mypod 5000 6000
  ```

- 由workload配置规则转发

  本地监听5000和6000，deployment中配置了转发规则，由deployment自动转发

  ```
  kubectl port-forward deployment/mydeployment 5000 6000
  ```

==port-forward默认只监听localhost，如果想要监听所有IP需要使用`--address`参数==

```
#默认监听localhost，所以其他IP请求port-forward不会监听
[root@k8smaster opt]# kubectl port-forward pod/kube-nginx 10086:80
Forwarding from 127.0.0.1:10086 -> 80
Forwarding from [::1]:10086 -> 80

[root@k8smaster ~]# ss -lnpt | grep 10086
LISTEN     0      128    127.0.0.1:10086                    *:*                   users:(("kubectl",pid=114338,fd=8))
LISTEN     0      128      [::1]:10086                 [::]:*                   users:(("kubectl",pid=114338,fd=9))

C:\Users\82341>curl -I 192.168.80.201:10086
curl: (7) Failed to connect to 192.168.80.201 port 10086: Connection refused


#监听所有的IP
[root@k8smaster opt]kubectl port-forward --address=0.0.0.0 pod/kube-nginx 10086:80
Forwarding from 0.0.0.0:10086 -> 80
Handling connection for 10086

[root@k8smaster ~]# ss -lnpt | grep 10086
LISTEN     0      128          *:10086                    *:*                   users:(("kubectl",pid=114033,fd=8))

C:\Users\82341>curl -I 192.168.80.201:10086
curl: (7) Failed to connect to 192.168.80.201 port 10086: Connection refused

C:\Users\82341>curl -I 192.168.80.201:10086
HTTP/1.1 200 OK
Server: nginx/1.19.6
Date: Thu, 01 Apr 2021 08:36:31 GMT
Content-Type: text/html
Content-Length: 612
Last-Modified: Tue, 15 Dec 2020 13:59:38 GMT
Connection: keep-alive
ETag: "5fd8c14a-264"
Accept-Ranges: bytes
```

