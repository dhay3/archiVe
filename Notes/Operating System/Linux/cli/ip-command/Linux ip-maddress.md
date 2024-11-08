# Linux ip-maddress

ref:

https://en.wikipedia.org/wiki/Multicast_address

## 0x1 Digest

用于管理multicast地址

## 0x2 Commands

### ip maddress show

查看多播地址

- `dev NAME`(default)

  查看指定设备的

  ```
  [vagrant@localhost ~]$ ip maddress 
  1:      lo
          inet  224.0.0.1
          inet6 ff02::1
          inet6 ff01::1
  2:      eth0
          link  01:00:5e:00:00:01
          link  33:33:00:00:00:01
          link  33:33:ff:4d:77:d3
          inet  224.0.0.1
          inet6 ff02::1:ff4d:77d3
          inet6 ff02::1
          inet6 ff01::1
  ```

  其中的224.0.0.1为保留IP表示所有组播主机（包括路由器）

### ip maddress add | delete

增加或删除多播地址



