# docker port

pattern：`docker port <containerID>`

用于展示指定容器开放的端口情况，==不会展示`--expose`(指对宿主机开放)指定的端口==

```
root in /etc/ssh λ docker run -itd --expose 10-20 -p 800:80 -p 4430:443/tcp --name t2 busybox
01988200e30e47752ee81b11f76300ea739f55a3df00e8e8448cd6d42a19125f                                                                                               /0.6s
root in /etc/ssh λ docker port t2
443/tcp -> 0.0.0.0:4430
80/tcp -> 0.0.0.0:800     
```

