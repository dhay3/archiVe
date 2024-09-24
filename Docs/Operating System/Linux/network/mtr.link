# Linux mtr

ref:

https://www.linode.com/docs/guides/diagnosing-network-issues-with-mtr/

https://www.cnblogs.com/alongdidi/p/mtr.html

## Digest

syntax

```
mtr [-4|-6] [-F FILENAME] [--report] [--report-wide] [--xml] [--gtk]
[--curses] [--displaymode MODE] [--raw] [--csv]  [--json]  [--split]
[--no-dns] [--show-ips] [-o FIELDS] [-y IPINFO] [--aslookup] [-i IN‐
TERVAL] [-c COUNT] [-s PACKETSIZE] [-B BITPATTERN]  [-G GRACEPERIOD]
[-Q TOS] [--mpls] [-I NAME] [-a ADDRESS] [-f FIRST-TTL] [-m MAX-TTL]
[-U MAX-UNKNOWN] [--udp] [--tcp] [--sctp]  [-P PORT]  [-L LOCALPORT]
[-Z TIMEOUT] [-M MARK] HOSTNAME
```

my traceroute（mtr） 是一个网络诊断工具，结合了 `ping` 和 `traceroute`

发包的原理和 traceroute 一样，都是通过 3 层 IP ttl 来实现的

具体查看[Linux traceroute]()

需要注意一点的是早版本的 mtr 是不支持 TCP/UDP probe mode 的

## Output

mtr 默认按照如下格式输出

```
Start: Thu Nov 25 15:15:20 2021
HOST: gonda033059000078.na175     Loss%   Snt   Last   Avg  Best  Wrst StDev
  1.|-- ???                       100.0     1    0.0   0.0   0.0   0.0   0.0
```

1. 跳数
2. IP或domain name
3. loss%丢包率
4. snt 发包数
5. last 最后一个包的时延
6. avg 平均时延
7. best 最佳时延
8. wrst 最差时延
9. stdev 标准delta

## Optional args

### Probe mode args

- `-u | --udp`

  以udp替代icmp 发包

- `-T | --tcp`

  已tcp替代icmp 发包

- `-P <port> | --port <port>`

  指定tcp/udp对端的端口

### Input args

- `-4 | -6`

  只使用IPv4 或 IPv6

- `-F <file>| --filename <file> `

  从文件中读取hostname

- `-n | --no-dns`

  解析domain name

- `-b | --show-ips`

  同时显示主机名和ip

- `-o <fields>`

  以指定的顺序输出指定的fieds

  ```
  L := loss ratio
  D := dropped packets
  R := received packets
  S := sent packets
  N := newest rtt(ms)
  B := min/bes rtt(ms)
  A := Average rtt(ms)
  W := max/worst rtt(ms)
  V := standard deviation
  G := geometric mean
  J := current jitter
  X := worst jitter
  I := interrival jitter
  ```

  例如：`mtr -o "LSD NBAW X"`

- `-i <secs> | --interval <secs>`

  指定 icmp echo request 发包的时间间隔，默认1 sec

  只有 root 用户才可以指定 0 到 1 之间的值

- `-c <count> | --report-cycles <count>`

  指定发包的次数，等价与`traceroute -q <cnt>`

- `-s <packetsize> | --psize <packetsize>`

  包含IP和ICMP headers，默认36(20+16)bytes

- `-B <num> | --bitparttern <num>`

  payload大小

- `-G <secs> | --gracetime <secs>`

  等待最后一个包的响应时间，默认5secs

- `-Q <num> | --tos <num>`

  指定tos的值

- `-m <num> | --max-ttl <num>`

  指定 max ttl 的值

- `-M <mark>`

  对 mtr 的包 mark

### Output args

- `-r | --report`

  以report mode运行，需要和`-c n`一起使用。在n次数后退出

  ```
  #mtr -r -c 1 taobao.com
  Start: Thu Nov 25 15:15:20 2021
  HOST: gonda033059000078.na175     Loss%   Snt   Last   Avg  Best  Wrst StDev
    1.|-- ???                       100.0     1    0.0   0.0   0.0   0.0   0.0
    2.|-- 10.231.20.189              0.0%     1    9.2   9.2   9.2   9.2   0.0
  ```

- `-x  | --xml`

  以xml形式输出

  ```
  #mtr -x -c 1 taobao.com
  <MTR SRC="gonda033059000078.na175" DST="taobao.com" TOS="0x0" PSIZE="64" BITPATTERN="0x00" TESTS="1">
      <HUB COUNT="1" HOST="???">
          <Loss%> 100.0%</Loss%>
          <Snt>     1</Snt>
          <Last>   0.0</Last>
          <Avg>   0.0</Avg>
          <Best>   0.0</Best>
          <Wrst>   0.0</Wrst>
          <StDev>   0.0</StDev>
      </HUB>
      <HUB COUNT="2" HOST="10.231.18.165">
          <Loss%>  0.0%</Loss%>
          <Snt>     1</Snt>
          <Last>   5.1</Last>
          <Avg>   5.1</Avg>
          <Best>   5.1</Best>
          <Wrst>   5.1</Wrst>
          <StDev>   0.0</StDev>
      </HUB>
  ```

- `-l | --raw`

  以raw格式输出

  ```
  #mtr -l taobao.com
  h 1 10.231.18.253
  p 1 7597
  h 2 11.73.93.61
  p 2 241
  h 3 11.73.44.229
  ```

- `-j | --json`

  以json的格式输出

## Interactive control

如果 mtr 不是以 report mode 运行，还可以使用如下键来控制 mtr

只写几个常用的值，具体查看 man page

- p: pause
- n: toogle DNS on/off
- r: reset all counters
- j: toggle latency/jitter stats

## Cautions

==如果目的 IP 出现在 NAT 网络中，mtr 的结果可能会不准==。例如：

第 2 跳 和 第 3 跳 之前是 fullNAT（同时做 SNAT 和 DNAT） LVS，LVS 根据路由按照 VS 条目转发到 Realserver 上，第 3 跳的回包过 LVS, SNAT 成 3.3.3.3

```
#traceroute
1 1.1.1.1
2 2.2.2.2
3 3.3.3.3
4 3.3.3.3
5 *
6 *
...

#mtr
HOST: localhost                      Loss%   Snt   Last  Avg  Best  Wrst   StDev
   1. 1.1.1.1  		                 0.0%    10    0.3   0.6   0.3   1.2   0.3
   2. 2.2.2.2  		                 0.0%    10    0.4   1.0   0.4   6.1   1.8
   3. 3.3.3.3                 		 0.0%    10    0.8   2.7   0.8  19.0   5.7
```

对比看发现，traceroute 显示是不通的，但是 mtr 显示是通的。碰到这种情况是以 traceroute 的结果为准

## Analyze reports

### Packet loss

分析 mtr 一般只看 2 个字段 loss 和 latency，如果看见某一跳出现 percentage of loss 只能表示==可能==出现丢包(这里的逻辑和 traceroute 中 asterisk 的逻辑相同)。==一般出现丢包只看最后一跳的 percentage of loss==。下面的这个例子就是正常的

```
root@localhost:~# mtr --report www.google.com
HOST: example                       Loss%   Snt   Last  Avg  Best  Wrst   StDev
   1. 63.247.74.43                  0.0%    10    0.3   0.6   0.3   1.2   0.3
   2. 63.247.64.157                50.0%    10    0.4   1.0   0.4   6.1   1.8
   3. 209.51.130.213                0.0%    10    0.8   2.7   0.8  19.0   5.7
   4. aix.pr1.atl.google.com        0.0%    10    6.7   6.8   6.7   6.9   0.1
   5. 72.14.233.56                  0.0%    10    7.2   8.3   7.1  16.4   2.9
   6. 209.85.254.247                0.0%    10   39.1  39.4  39.1  39.7   0.2
   7. 64.233.174.46                 0.0%    10   39.6  40.4  39.4  46.9   2.3
   8. gw-in-f147.1e100.net          0.0%    10   39.6  40.5  39.5  46.7   2.2
```

出现第 2 跳中的现象，概率是对应的路由设备 ICMP rate limitation

对比下面的这个例子就是异常的，中间路由有丢包

```
root@localhost:~# mtr --report www.google.com
HOST: localhost                      Loss%   Snt   Last  Avg  Best  Wrst   StDev
   1. 63.247.74.43                   0.0%    10    0.3   0.6   0.3   1.2   0.3
   2. 63.247.64.157                  0.0%    10    0.4   1.0   0.4   6.1   1.8
   3. 209.51.130.213                60.0%    10    0.8   2.7   0.8  19.0   5.7
   4. aix.pr1.atl.google.com        60.0%    10    6.7   6.8   6.7   6.9   0.1
   5. 72.14.233.56                  50.0%    10    7.2   8.3   7.1  16.4   2.9
   6. 209.85.254.247                40.0%    10   39.1  39.4  39.1  39.7   0.2
   7. 64.233.174.46                 40.0%    10   39.6  40.4  39.4  46.9   2.3
   8. gw-in-f147.1e100.net          40.0%    10   39.6  40.5  39.5  46.7   2.2
```

第3，4 跳显示有 60% 的包丢了。你可以认为是发生了丢包（因为后面的几跳没有出现 0%），但是后面几跳并没有出现 60%。且最后一跳有 40% 的包丢了。说明实际丢包大概率在 40%

上面显示 60% 可能的原因有

1. 中间的路由设备设置了 ICMP 流量限制，导致 10% - 20% 的包到了，但是没有回包。
2. 因为回包的原IP不同，回包路由或者回包ACL可能会导致丢包。

==另外即使最后一跳显示有 40% 的丢包，这也不能说明是源到目中间链路的问题，也有可能是目到源回包中间链路的问题。即可能出现包 100% 能到目的，但是 40% 包不能通过回程路由达到源==

所以为了确认到底是在那里丢的包，==一般需要提供双向的 mtr==

### Latency

latency 一般和 hops,distance 成线性正比，例如

```
root@localhost:~# mtr --report www.google.com
HOST: localhost                      Loss%   Snt   Last   Avg  Best  Wrst  StDev
    1. 63.247.74.43                  0.0%    10    0.3   0.6   0.3   1.2   0.3
    2. 63.247.64.157                 0.0%    10    0.4   1.0   0.4   6.1   1.8
    3. 209.51.130.213                0.0%    10    0.8   2.7   0.8  19.0   5.7
    4. aix.pr1.atl.google.com        0.0%    10  388.0 360.4 342.1 396.7   0.2
    5. 72.14.233.56                  0.0%    10  390.6 360.4 342.1 396.7   0.2
    6. 209.85.254.247                0.0%    10  391.6 360.4 342.1 396.7   0.4
    7. 64.233.174.46                 0.0%    10  391.8 360.4 342.1 396.7   2.1
    8. gw-in-f147.1e100.net          0.0%    10  392.0 360.4 342.1 396.7   1.2
```

第 3 跳 到 第 4 跳 latency 显著的升高，且后面几跳同样也很高。这就表明第 4 跳对应的路由设备可能有问题

1. 设备物理性能到达瓶颈，处理包速度慢
2. 路由设备配置错误
3. 第 3 跳 到 第 4 跳 链路拥塞

另外 lantency 也会出现表象的增高，例如

```
root@localhost:~# mtr --report www.google.com
HOST:  localhost                     Loss%   Snt   Last  Avg  Best  Wrst   StDev
    1. 63.247.74.43                  0.0%    10    0.3   0.6   0.3   1.2   0.3
    2. 63.247.64.157                 0.0%    10    0.4   1.0   0.4   6.1   1.8
    3. 209.51.130.213                0.0%    10    0.8   2.7   0.8  19.0   5.7
    4. aix.pr1.atl.google.com        0.0%    10    6.7   6.8   6.7   6.9   0.1
    5. 72.14.233.56                  0.0%    10  254.2 250.3 230.1 263.4   2.9
    6. 209.85.254.247                0.0%    10   39.1  39.4  39.1  39.7   0.2
    7. 64.233.174.46                 0.0%    10   39.6  40.4  39.4  46.9   2.3
    8. gw-in-f147.1e100.net          0.0%    10   39.6  40.5  39.5  46.7   2.2 
```

第 4 到 第 5 跳 latency 显著增高，但是后面几跳 latency 是正常的。第 5 跳为什么会出现这种现象，

1. 概率是策略第 5 跳的回包路径不一样
2. ICMP rate limitation （没能搞明白，为什么会导致表象的升高）

所以访问 google 的实际平均时延是 40.5 ms

## Common reports

### Destination host network configured improperly

```
root@localhost:~# mtr --report www.google.com
HOST:  localhost                     Loss%   Snt   Last  Avg  Best  Wrst  StDev
    1. 63.247.74.43                  0.0%    10    0.3   0.6   0.3   1.2   0.3
    2. 63.247.64.157                 0.0%    10    0.4   1.0   0.4   6.1   1.8
    3. 209.51.130.213                0.0%    10    0.8   2.7   0.8  19.0   5.7
    4. aix.pr1.atl.google.com        0.0%    10    6.7   6.8   6.7   6.9   0.1
    5. 72.14.233.56                  0.0%    10    7.2   8.3   7.1  16.4   2.9
    6. 209.85.254.247                0.0%    10   39.1  39.4  39.1  39.7   0.2
    7. 64.233.174.46                 0.0%    10   39.6  40.4  39.4  46.9   2.3
    8. gw-in-f147.1e100.net         100.0    10    0.0   0.0   0.0   0.0   0.0
```

如果最后一跳 100% loss 且是目的主机，概率是目的主机的问题

1. 如果是 3 层，概率是路由配置错误，或者配置了 ignore_icmp_all，或者是 firewall 策略
2. 如果是 4 层，概率是 firewall 策略

### ICMP rate limitation

```
root@localhost:~# mtr --report www.google.com
HOST:  localhost                     Loss%   Snt   Last  Avg  Best  Wrst   StDev
    1. 63.247.74.43                  0.0%    10    0.3   0.6   0.3   1.2   0.3
    2. 63.247.64.157                 0.0%    10    0.4   1.0   0.4   6.1   1.8
    3. 209.51.130.213                0.0%    10    0.8   2.7   0.8  19.0   5.7
    4. aix.pr1.atl.google.com        0.0%    10    6.7   6.8   6.7   6.9   0.1
    5. 72.14.233.56                 60.0%    10   27.2  25.3  23.1  26.4   2.9
    6. 209.85.254.247                0.0%    10   39.1  39.4  39.1  39.7   0.2
    7. 64.233.174.46                 0.0%    10   39.6  40.4  39.4  46.9   2.3
    8. gw-in-f147.1e100.net          0.0%    10   39.6  40.5  39.5  46.7   2.2
```

如果中间几跳出现丢包，但是最后一跳没有出现丢包。通常是中间的路由设备设置了 ICMP 回包速率（但是也不排除第 4 跳回包的路径不同导致的）

## Examples

```
mtr -rnTP 80 -i0.1 -c 100 host
```



