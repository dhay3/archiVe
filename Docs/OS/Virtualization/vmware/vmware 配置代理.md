# vmware 配置代理

> 使用host-only是最方便，如果使用bridge换网络环境之后就需要重新配置。

1. 在v2ray中设置，允许来自局域网的连接

2. 设置浏览器`socks5://192.168.10.1:1080`。这里的IP是host_only宿主机的IP，端口是v2ray的代理端口

3. 设置http_proxy和https_proxy

   ```
   export http_proxy=socks5://192.168.10.1:1080
   export https_proxy=socks5://192.168.10.1:1080
   ```

   

