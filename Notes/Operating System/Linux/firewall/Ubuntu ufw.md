# Ubuntu ufw

## 概述

网络层的防火墙。其原理还是iptables。

`Default: deny (incoming), allow (outgoing), disabled (routed)`默认拒绝所有入流量

使用`systemctl status ufw`查看时不能显示ufw是否生效，需要通过`ufw status`来查看

```
root in ~ λ systemctl status ufw
● ufw.service - Uncomplicated firewall
   Loaded: loaded (/lib/systemd/system/ufw.service; enabled; vendor preset: enabled)
   Active: active (exited) since Thu 2021-01-07 13:57:48 HKT; 5min ago
     Docs: man:ufw(8)
  Process: 1779 ExecStop=/lib/ufw/ufw-init stop (code=exited, status=0/SUCCESS)
  Process: 1780 ExecStart=/lib/ufw/ufw-init start quiet (code=exited, status=0/SUCCESS)
 Main PID: 1780 (code=exited, status=0/SUCCESS)

Jan 07 13:57:48 ubuntu18.04 systemd[1]: Starting Uncomplicated firewall...
Jan 07 13:57:48 ubuntu18.04 systemd[1]: Started Uncomplicated firewall.                               /0.0s
root in ~ λ ufw status
Status: inactive                 
```

使用`ufw enable/disable`打开和关闭ufw

> 需要注意的时，如果没有开启22端口，会将除当前窗口之外的ssh客户端断开

```
root in ~ λ ufw enable
Command may disrupt existing ssh connections. Proceed with operation (y|n)? y
Firewall is active and enabled on system startup                                                      /2.1s
root in ~ λ ufw status verbose  
Status: active
Logging: on (low)
Default: deny (incoming), allow (outgoing), disabled (routed)
New profiles: skip

To                         Action      From
--                         ------      ----
443/tcp                    ALLOW IN    Anywhere                  
443/tcp (v6)               ALLOW IN    Anywhere (v6)   
root in ~ λ ufw disable
Firewall stopped and disabled on system startup  
```

## 策略

### allow

```
ufw allow 443 #允许所有流量访问本机的443端口
ufw allow 443/tcp #只允许tcp流量访问本机的443端口
ufw allow in on eth0 to 192.168.0.1 proto tcp #允许所有通过eth0到192.168.0.1的流量
```

### deny

```
ufw deny proto tcp to any port 80 #拒绝所有tcp流量到达本机的80端口
ufw deny proto tcp from 2001:db8::/32 to any port 25
ufw deny in on eth0 to 224.0.0.1 proto igmp #拒绝所有通过eth0发送到224.0.0.1的流量
```

### reject

让用户知道被拒绝了

```
 ufw reject auth
```

### limit

limit策略能有效防止被暴力破解，默认30s内多次尝试，超过6次，就会触发规则

```
ufw limit ssh/tcp
```

## 增删

增加

```
ufw insert 3 deny to any port 22 from 10.0.0.135 proto tcp
```

删除

```
 ufw delete deny 80/tcp #删除特定规则
 ufw delete 3 #删除指定索引的规则
```











