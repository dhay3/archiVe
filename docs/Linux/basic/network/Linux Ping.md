#  Linux Ping

> 如果ICMP不通但是TCP可以通可能是路由器策略，主机iptables，内核参数(`/proc/sys/net/ipv4/icmp_ignore_all=1`)

linux ping强制使用ICMP echo(8)发送数据包，默认带有IP首部(最大60byte)

```

```

syntax：`ping [options] destination`

## options

- -A

  以200ms间隔发包，等价于flood

- -f

  flood ping，不回显数据

- -c <count>

  发送指定个数的数据包

- -i <interval>

  每次发包的间隔时间，默认1s

- -n 

  numeric output

- -O

  如果前一个包不可用，一般和-D一起使用

  ```
  
  ```

- -R

  record route

  ```
  
  ```

  

- -r

- -s <packagesize>

- -D

  print timestamp

  ```
  
  ```

- -W <timeout>
