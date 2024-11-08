# ipvsadm

ipvsadm用于创建和管理LVS规则，和iptables非常类似

syntax：`ipvsadm command [protocal] <service-address>`

service-address由三部分组成：IP address，port number，protocol

## command

- -L list virtual server table

  ```
  [root@VM-0-4-centos ~]# ipvsadm -L -n
  IP Virtual Server version 1.2.1 (size=4096)
  Prot LocalAddress:Port Scheduler Flags
    -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
  TCP  207.175.44.110:80 rr
    -> 192.168.10.1:80              Masq    1      0          0
    -> 192.168.10.2:80              Masq    1      0          0
    -> 192.168.10.3:80              Masq    1      0          0
    -> 192.168.10.4:80              Masq    1      0          0
    -> 192.168.10.5:80              Masq    1      0          0
  ```

- -Z zero the packet，byte and rate count

- -A  add a virtual service，这里必须指定
- -a add real server to a virtual service

- -E edit a virtual service

- -e Edit a real server in a virtual service

- -D delete a virtual service

  ```
  [root@VM-0-4-centos ~]# ipvsadm -D -t 207.175.44.110:80
  ```

- -d Remove a real server from a virtual service

- -R 将-S导出的规则载入

  ```
  [root@VM-0-4-centos ~]# ipvsadm -R < file
  ```

- -S 将LVS rules输出到stdout

  ```
  [root@VM-0-4-centos ~]# ipvsadm -S
  -A -t 8.135.0.171:http -s rr
  -a -t 8.135.0.171:http -r 192.168.10.1:http -m -w 1
  -a -t 8.135.0.171:http -r 192.168.10.2:http -m -w 1
  -a -t 8.135.0.171:http -r 192.168.10.3:http -m -w 1
  -a -t 8.135.0.171:http -r 192.168.10.4:http -m -w 1
  -a -t 8.135.0.171:http -r 192.168.10.5:http -m -w 1
  ```

## parameters

- -t | --tcp-service service-address

  使用tcp，service-address 使用 host[:port]格式。port可以被忽略，表示任意port，但是必须和`-p`一起使用

- -u

  使用udp，等同-t

- -r | --real-server service-address

  指定关联的real server

- -s | --scheduler scheduling-method-alogrithm

  指定连接到real servers的算法，支持如下几种算法：

  1. rr：round robin
  2. wrr：weighted round robin
  3. lc：least-connection
  4. wlc：weighted least-connection
  5. lblc：locality-based least-connection
  6. lblcr： locality-Based  least-connection
  7. dh：destination hashing
  8. sh：source hashing 
  9. sed：shortest Expeccted delay
  10. nq：never queue

- -p | --persitent [timeout]

  表明virtual service is persistent。同一个客户端发起的请求在timeout(默认300)内都往同一个real server发送。类似于有时限的hash一致算法。

- -g | --gatewaying 使用网关转发请求
- -i | --ipip  使用通道转发请求
- -m | --masquerading  使用伪装转发请求

- -w | --weight weight 指定分配的权重

- -y 指定最小连接数

- -n numberic output，必须在-L之后

  ```
  [root@VM-0-4-centos ~]# ipvsadm -Ln
  IP Virtual Server version 1.2.1 (size=4096)
  Prot LocalAddress:Port Scheduler Flags
    -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
  TCP  8.135.0.171:80 rr
  ```

## 例子

1. 创建一个tcp virtual IP，使用round robin 访问real server

   ```
   [root@VM-0-4-centos ~]# ipvsadm -A -t 8.135.0.171 -s rr
   Zero port specified for non-persistent service
   ```

2. 添加规则

   ```
   ipvsadm -a -t 8.135.0.171:80 -r 192.168.10.1:80 -m
   ipvsadm -a -t 8.135.0.171:80 -r 192.168.10.2:80 -m
   ipvsadm -a -t 8.135.0.171:80 -r 192.168.10.3:80 -m
   ipvsadm -a -t 8.135.0.171:80 -r 192.168.10.4:80 -m
   ipvsadm -a -t 8.135.0.171:80 -r 192.168.10.5:80 -m
   ```

3. 因为需要让virtual server转发数据包，所以要开启内核ip_forward参数。==还需要让配置文件生效==

   ```
   echo "1" > /proc/sys/net/ipv4/ip_forward
   ```

   







