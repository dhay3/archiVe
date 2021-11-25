# mtr - mytraceroute

> mtr -rc 1 -Tp 80 taobao.com

## digest

mtr是一个网络诊断工具，结合了traceroute和ping的功能。默认发icmp request以flow形式输出

syntax：`mtr [options]`

## output fields

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

## input mode

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

  指定icmp echo request发包的时间间隔，默认1 sec

- `-c <count> | --report-cycles <count>`

  reportmode下执行tracroute的次数

- `-s <packetsize> | --psize <packetsize>`

  包含IP和ICMP headers，默认36(20+16)bytes

- `-B <num> | --bitparttern <num>`

  payload大小

- `-G <secs> | --gracetime <secs>`

  等待最好一个包的响应时间，默认5secs

- `-Q <num> | --tos <num>`

  指定tos的值

- `-M <mark>`

  对mtr的包mark

## probe mode

- `-u | --udp`

  以udp替代icmp 发包

- `-T | --tcp`

  已tcp替代icmp 发包

- `-P <port> | --port <port>`

  指定tcp/udp对端的端口

## output mode

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

  