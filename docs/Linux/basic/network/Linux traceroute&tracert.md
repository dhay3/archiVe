# traceroute/tracert

## traceroute

在linux平台上使用，traceroute会将经过的路由记录下台，默认使用UDP 33434端口，但是端口没探测一次就会加以。如果是==UDP/ICMP==目标没有回包(被并不表示目标不可达，只是没有回echo reply。可能是被firewall过滤)，会打印asterisk( * )

## options

- -U

  使用UDP 53端口探测（UDP可能会被firewall过滤）

- -I | --icmp

  使用icmp echo 探测

- -T | --tcp

  使用TCP SYN探测，外绕过防火墙

- `-q <squeries>`

  并发发送数据包可以加快响应速度，如果这个值太大可能会被router或主机过滤

- ` -z <sendwait>`

  发包的速度，默认0，如果在10以内表示s, 以外表示ms

- `-m <max ttl>`

  ttl的最大值，默认30

- `-p <port>`

  指定探测的端口

  如果是UDP或ICMP每探测一个hop，port就会加一

  如果是TCP，每探测一个hop，port不会改变

- `-O <options>`

  指定某些探测方法特有的属性，具体查看list of available methods

## 例子

- DNS探测

  `traceroute -U host`

- ICMP探测

  `traceroute -I host`

- TCP 探测

  `traceroute -Tp 80 host`
