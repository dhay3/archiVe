# Linux网络解决方案

**问题一：**

可以ping通主机，但是不能解析域名，可以尝试以下方法

1. 配置`/etc/resolv.conf`

2. 关闭NetworkManager

   ```
   systemctl stop NetworkManager
   ```

3. 在网络配置文件`ifcfg-ens*`中添加DNS属性

4. 查看是否有默认路由

   ```
   route -n 
   route add -net 0.0.0.0 gw gateway dev device
   ```

   

