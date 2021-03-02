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

**问题二：**

DNS失效

1. 配置`resolv.conf`

   ```
   nameserver 8.8.8.8
   
   nameserver 8.8.4.4
   ```

2. 配置NIC文件添加，BOOTPROTO需要dhcp才会生效

   ```
   DNS1=8.8.8.8
   DNS2=8.8.4.4
   ```

3. 强制设置file attr

   ```
   chattr +i /etc/resolv.conf
   ```

   

