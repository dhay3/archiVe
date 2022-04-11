# Linux 代理设置

参考：

https://zhuanlan.zhihu.com/p/46973701

> 如果使用客户端的代理加密工具，例如v2ray，==需要开启v2ray的允许来自局域网的连接功能==。

## 虚拟机

### 浏览器

1. 如果虚拟机使用bridge模式无需设置
2. 如果虚拟机使用的网络模式是host-only，可以将代理的IP设置为host-only分配给宿主机的IP加代理的端口。同理NAT

### 终端

1. 按照如下设置只对当前shell生效(关闭shell或unset变量后失效)，可以配置进shell配置文件中永久生效。==不会对icmp协议代理，所以用ping校验无效，可以使用curl==

   ```
   #http代理
   export http_proxy=http://127.0.0.1:12333
   export https_proxy=http://127.0.0.1:12333
   
   #socks代理
   export http_proxy=socks5://127.0.0.1:12333
   export https_proxy=socks5://127.0.0.1:12333
   ```

   









