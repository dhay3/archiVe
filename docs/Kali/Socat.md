# Socat

参考：

https://www.hi-linux.com/posts/61543.html

> 实验前关闭防火墙

Socat是一款socket工具，可以看作是netcat的加强版

syntax：`socat [options] <address> <address> `

1. `-`表示==STDIO=stdin + stdout==
2. 大小写不敏感
3. 两个地址可以组合，`open:read.html\!\!open:write.txt,create,append`，由于`!`表示历史命令，需要转义
4. address options以逗号隔开

## 参数

- -d

  debug，d越多打印信息越多

- `-lf <filename>`

  将debug信息存储到指定位置，默认stdout

  ```
  root in /opt λ socat -ddd -lf socat.log - /opt/a.sh
  ```

- `-T <timeout>`

  指定连接在多长时间后断开

- -u 

  指定第一地址只做读
  
  -U
  
  指定第一个地址只做写
  
  ==默认双向通信==



## Address Types

- STDERR | STDOUT | STDIN

- PTY

  生成一个pseudo terminal

- `CREATE:<filename>`

  使用create函数创建文件

- `EXEC:<command-line>`

  执行命令，默认寻找`$PATH`

- `TCP:<host>:<port> | UDP:<host>:<port>`

- `IP-SENDTO:<host>:<protocol> | IP4-SENDTO:<host>:<protocol>`

  打开socket

- `IP-RECV:<protocol> | IP4-RECV:<protocol>`

- `INTERFACE:<interface>`

  使用指定的iface

- `OPEN:<filename>`

  打开文件

- `OPENSSL:<host>:<port>`

  使用SSL连接

- `OPENSSL-LISTEN:<port>`

- PIPE 

  生成一个匿名管道符

- `PIPE:<filename>`

  生成一个具名管道符

- `PROXY:<proxy>:<hostname>:<port>`

## Address Options

具体查看`^\s+<address-types> option group`

## 例子

1. 监听所有IP的8080端口

   ```
   root in /opt λ socat - tcp-listen:8080
   ```
   
2. 建立TCP连接

   ```
   #192.168.80.200
   root in /opt λ socat - tcp:192.168.80.140:8080
   hello world
   
   #192.168.80.140
   [root@chz ~]# nc -lkp 8080
   hello world
   ```

3. 查看文件

   file可以被省略

   ```
   root in /opt λ socat - file:/opt/a.sh
   #!/bin/bash
   exec 3 < /dev/tcp/8.135.0.171/22
   #timeout 1 cat &< 3
   ```

4. 执行命令，也可以建立shell，`socat TCP-LISTEN:7005,fork,reuseaddr EXEC:/bin/bash,pty,stderr`

   200主机执行的结果返回给140:8080

   ```
   #192.168.80.200
   root in /opt λ socat  exec:ls tcp:192.168.80.140:8080
   #或者，以在pty上的格式输出
   root in /opt λ socat  exec:ls,pty  tcp:192.168.80.140:8080
   
   
   #192.168.80.140
   [root@chz ~]# nc -lkp 8080
   a.sh
   b
   bak.xml
   Blasting_dictionary-master
   burpsuite pro
   busybox.tar
   containerd
   file
   hydra.restore
   jdk-14.0.2
   jdk-14.0.2_linux-x64_bin.tar
   kubectl
   lsd_0.18.0_amd64.deb
   packer_1.7.0_linux_amd64.zip
   ```

5. 连接主机，并分配pty

   ```
   root in /opt λ socat - exec:'ssh -l root 192.168.80.140',pty
   root@192.168.80.140's password:
   ```

6. 转存信息到文件

   ```
   socat -u TCP4-LISTEN:3334,reuseaddr,fork \
          OPEN:/tmp/in.log,creat,append
   ```

7. 端口转发

   监听192.168.80.200的8080端口，将请求转发到192.168.80.140:80的nginx服务。不会保持长连接

   ```
   #192.168.80.200
   root in /opt λ socat tcp-listen:8080,bind=192.168.80.200,reuseaddr tcp:192.168.80.140:80
   
   #测试
   C:\Users\82341>curl -I 192.168.80.200:8080
   HTTP/1.1 200 OK
   Server: nginx/1.18.0
   Date: Tue, 06 Apr 2021 10:50:50 GMT
   Content-Type: text/html
   Content-Length: 612
   Last-Modified: Thu, 29 Oct 2020 15:25:17 GMT
   Connection: keep-alive
   ETag: "5f9adedd-264"
   Accept-Ranges: bytes
   ```

8. 文件传输，也可以结合`-u | -U`

   ```
   #192.168.80.200
   root in /opt λ socat open:a.sh  tcp-listen:8080
   
   #192.168.80.140
   [root@chz /]# socat tcp:192.168.80.200:8080 open:ttt,create
   [root@chz /]# cat ttt
   #!/bin/bash
   exec 3 < /dev/tcp/8.135.0.171/22
   #timeout 1 cat &< 3
   ```

9. NAT

   ```
   #192.168.80.200 监听8080 和 9000，往本机8080发送请求能转到本机的9000端口，反之9000端口能到8080
   root in /opt λ socat tcp-listen:8080 tcp-listen:9000
   
   #192.168.80.140 将远程8080端口接收到的请求转发到本机的80端口
   [root@chz /]# socat tcp:192.168.80.200:8080 tcp:192.168.80.140:80
   
   #9000没有服务所以报错
   C:\Users\82341>curl -I 192.168.80.200:8080
   curl: (7) Failed to connect to 192.168.80.200 port 8080: Connection refused
   
   C:\Users\82341>curl -I 192.168.80.200:9000
   HTTP/1.1 200 OK
   Server: nginx/1.18.0
   Date: Tue, 06 Apr 2021 11:27:13 GMT
   Content-Type: text/html
   Content-Length: 612
   Last-Modified: Thu, 29 Oct 2020 15:25:17 GMT
   Connection: keep-alive
   ETag: "5f9adedd-264"
   Accept-Ranges: bytes
   ```

- 伪造web服务器

  ```
  #192.168.80.200
  root in /opt λ cat read.html
    File: read.html
    <!DOCTYPE html>
    <html>
    <head>
    <title>Page Title</title>
    </head>
    <body>
    <h1>This is a Heading</h1>
    <p>This is a paragraph.</p>
    </body>
    </html>
  
  root in /opt λ socat open:read.html\!\!open:write.html,create,append tcp-listen:8080
  
  root in /opt λ cat write.html
    File: write.html
    GET / HTTP/1.1
    Host: 192.168.80.200:8080
    User-Agent: curl/7.55.1
    Accept: */*
  
  #测试
  C:\Users\82341>curl 192.168.80.200:8080
  <!DOCTYPE html>
  <html>
  <head>
  <title>Page Title</title>
  </head>
  <body>
  
  <h1>This is a Heading</h1>
  <p>This is a paragraph.</p>
  
  </body>
  </html>
  ```

  







