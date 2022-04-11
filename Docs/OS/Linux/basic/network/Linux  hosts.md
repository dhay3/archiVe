# Linux  hosts

> `/etc/hosts`文件于windows中的`C:\Windows\System32\drivers\etc\HOST`文件相同，用于定义host

如果只需要内网通讯(无法通过internet访问)，也可以配置该文件

## Exmaple

pattern：`ip domain [aliases...]`

aliases是可变参数

```
[root@cyberpelican etc]# cat hosts
127.0.0.1   localhost cyberpelican.com localhost4 localhost4.localdomain4
::1         localhost localhost.localdomain localhost6 localhost6.localdomain6
192.168.80.139 139.com

----

[root@cyberpelican mail]# ping 139.com
PING 139.com (192.168.80.139) 56(84) bytes of data.
64 bytes from 139.com (192.168.80.139): icmp_seq=1 ttl=64 time=0.311 ms
64 bytes from 139.com (192.168.80.139): icmp_seq=2 ttl=64 time=0.383 ms
64 bytes from 139.com (192.168.80.139): icmp_seq=3 ttl=64 time=0.355 ms
^C

[root@cyberpelican mail]# ping cyberpelican.com
PING localhost (127.0.0.1) 56(84) bytes of data.
64 bytes from localhost (127.0.0.1): icmp_seq=1 ttl=64 time=0.038 ms
64 bytes from localhost (127.0.0.1): icmp_seq=2 ttl=64 time=0.042 ms
^C
```

