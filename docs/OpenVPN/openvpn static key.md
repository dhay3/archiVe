# openvpn static key

参考：https://openvpn.net/community-resources/static-key-mini-howto/

**优点**

- 设置简单

- 不是用X509 PKI （使用对称加密）

**缺点**

- 只能连接一个客户端和服务端
- key在两端中明文存在，一般通过scp拷贝到另外一台服务器
- 使用预定好的secure channel

### 例子

客户端通过10.8.0.1:1194连接服务端10.8.0.2:1194

1. 生成static key，`--genkey`必须和`--secret`一起使用

   ```
   openvpn --genkey --secret static.key
   ```

2. 服务端配置文件

   ```
   dev tun
   #没有指定remote监听所有IP，通过authentication才会转发数据包
   ifconfig 10.8.0.1 10.8.0.2
   secret static.key
   ```

3. 客户端配置文件

   ```
   remote remote-host
   dev tun
   ifconfig 10.8.0.2 10.8.0.1
   secret static.key
   ```

4. 防火墙配置，入站出站开放1194端口(openvpn默认使用端口)

5. 启动

   ```
   #client
   openvpn --config client.config
   
   #server
   openvpn --config server.config
   ```

6. 检验

   ```
   #注意这里只使用了udp1194端口，所以tcp1194还是可用的
   root in /tmp λ netstat -ln | grep 1194
   udp        0      0 0.0.0.0:1194            0.0.0.0:* 
   
   #会在客户端和服务端都出现一个tun iface，本机10.8.0.2点对点10.8.0.1
   7: tun0: <POINTOPOINT,MULTICAST,NOARP,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UNKNOWN group default qlen 100
       link/none 
       inet 10.8.0.2 peer 10.8.0.1/32 scope global tun0
       
   #client
   ping 10.8.0.1
   
   #server
   ping 10.8.0.2
   ```
   
   在客户端启动nginx，让服务端访问。数据包通过建立的tun通过来交换
   
   ```
   #client
   sudo docker run -itd --rm --name n1 -p 3389:80 nginx
   
   #server
   root in /tmp λ curl -fsSL 10.8.0.2:3389
   ```
   
   

