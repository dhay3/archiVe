# Chrony

ref：
[https://chrony.tuxfamily.org/](https://chrony.tuxfamily.org/)
[https://docs.oracle.com/cd/E90981_01/E90982/html/chrony.html](https://docs.oracle.com/cd/E90981_01/E90982/html/chrony.html)

## Digest
chrony 是 NTP 的一个实现，包含 chronyc (CLI to monitor chronyd) 和 chronyd (daemon) 两个软件
## Terms
### measurement
对 ntp server 探测一次
### RTC
real-time clock 是一个集成电路，用于记录计算机的时间，可以使用`hwclock`来查看和设置
### system clock
计算机识别的时间，以Unix的格式显示，从epoch开始计算
### offset
顾名思义，和 NTP server 的偏移补偿值
### drift
存储在本机上的文件记录和system clock 的偏差
### sample
ntp 同步的一次记录
### startum
[https://endruntechnologies.com/products/ntp-time-servers/stratum1?gclid=EAIaIQobChMIv92ih9bg9gIVW9VMAh1JbghUEAAYASAAEgLN1vD_BwE](https://endruntechnologies.com/products/ntp-time-servers/stratum1?gclid=EAIaIQobChMIv92ih9bg9gIVW9VMAh1JbghUEAAYASAAEgLN1vD_BwE)
startum levels define the distance from the reference clock
例如 startum-0 设备是最精准的使用 GPS 传输(GPS芯片)，但是不能被用在网络上，通常会连接到电脑。这台电脑也被称为startum-1
the basic definition of a startum-1 time server is that it be directly linked to a reliable source of UTC time such as GPS
通过网络连接在 startum-1 的设备也被称为 startum-2，以此类推。经过的startum越多，和 NTP server 差值就会越大
## Chronyd
用于和 NTP servers，reference clocks(GPS) 同步，如果没有指定`-f`参数  ，默认读取`/etc/chrony.conf`作为配置文件，日志会由syslog管控
### Optional args

- `-f FILE`

  指定配置文件的位置，默认`/etc/chrony.conf`

- `-q`

  只和 NTP server 同步一次后，退出 chronyd 进程，同步时不会detach from terminal

  ```
  [vagrant@localhost ~]$ sudo chronyd -q
  2022-03-29T03:43:54Z chronyd version 3.4 starting (+CMDMON +NTP +REFCLOCK +RTC +PRIVDROP +SCFILTER +SIGND +ASYNCDNS +SECHASH +IPV6 +DEBUG)
  2022-03-29T03:43:54Z Frequency 66.097 +/- 65.026 ppm read from /var/lib/chrony/drift
  2022-03-29T03:44:20Z System clock wrong by 0.004164 seconds (step)
  2022-03-29T03:44:20Z chronyd exiting
  ```

- `-Q`

  只输出offset，不会和 NTP server 同步

- `-r`

  重启 chronyd 进程

- `-t TIMEOUT`

  和 NTP server 同步超时退出 chronyd 的时间

- `-m`

  chronyd 只会在 RAM 中运行，不会被 swap 到交换分区

### Conf
内容较为复杂具体查看`man chrony.conf`
chrony 默认配置文件(`/etc/chrony.conf`)，可以通过`-f`手动指定

#### Time sources directives

- `server HOSTNAME [option]...`

  指定同步 NTP server，option有如下几个选项

  1. `minpoll POLL`

     this option specifies the minimum interval between requests sent to the server as a power of 2 in seconds. 默认值为6，即 64 sec，取值范围在 -6 到 24。如果minpoll的值太小会增加网络的开销，如果minpoll值太大同步的周期会比较长

  2. `maxpoll POLL`

     this option sepecifies the maximum interval between reqeusts sent to the server as a power of 2 in seconde. 默认值为10，即1024 sec，取值返回和minpoll相同

  3. `iburst`

     with this option，the interval between the first four request sent to the server will be 2 seconds or less instead of the interval specified by the minpoll opton，which allows chronyd to make the first update of the clock shortly after start
     前四次请求NTP server 的间隔在 2 sec，而非使用`minpoll`, 在chonryd 启动后和NTP server 快速同步

  4. `burst`

     with this option，chronyd will shorten the interval between up to four requests to 2 seconds or less when it cannot get a good measurement
     如果NTP server 无响应，一组四次请求之间的间隔在 2 sec 内，会增加网络的开销

  5. `key ID`

     使用MAC( message authentication code )来对 request 和 response 鉴权，防止 mitma

  6. `maxdelay DELAY`

     round-trip 最大的时间，例如`maxdelay 0.3`表示如果 NTP server 在 0.3 sec 内无响应该 server 就会北忽略。默认 3 sec

  7. `offset`

     if packet sent to the source whre on average delayed by 100 microseconds more than packets sent from the source back, the correction would be -0.00005 sec，默认 0.0 sec

  8. `prefer`

     指定当前 server 为高优先级( 对比无改选项的server )

  9. `noselect`

     不会选择改server

  10. `trust`

      如果其他 NTP server 和当前的 server 结果不一致，以当前的 server 为准

  11. `port`

      指定 NTP server 使用的port，默认 123

  12 . `version`	

  ​       指定发送给 NTP server 的 NTP 包的版本

- `pool NAME [options]`

  和 `server`指令块相同，使用的option也相同，但是指定的是ntp pool （通常未domain name）而不是单一的一个服务器

- `peer NAME [options]`

  和`server`指令块相同，但是不使用client/server 模式，使用peer模式(client可以是server)，不安全

- `manual`

  是否支持使用`chronyc settime`命令

- `dumpdir DIRECTORY`

  指定 measurement histories 记录的位置( 当 chronyd 进程退出或执行`chronyc dump`命令 )

#### Source selection

- `combinelimit LIMIT`

  如果 chronyd 设定了多个 sources

- `maxjitter JITTER`

  允许网络抖动的最大时间，默认 1 sec

- `minsources SOURCES`

  如果system clock 需要调整，至少需要 sources 个 sources 网络可达 

- `reselectdist DISTANCE`

  chronyd 会倾向于选择 shortest synchronisation distance 的 source， reselecting source distance 差值与当前 source 相近，就不会 reselect。默认为 100 microseconds

#### System clock

- `driftfile FILE`

  记录和 real time 差值的文件，重启系统会读取改文件中的差值作为补偿

- `makestep THRESHOLD LIMIT`

  chronyd 会周期性和 NTP server 同步，为了 correct any time offset 。如果NTPserver 和当前的system clock offset 超过了 threshold ，会在 limit 次内立即和 NTP server 同步

#### NTP server
当前host作为ntp server

- `allow [all] [subnet]`

  指定当前NTP server允许接入的NTP client的网段。启动chronyd时，默认以 client mode 执行，当指定改选项时可以是 client 也可以是其他 client 的 server，例如

  ```
  allow 1.2.3.4
  deny 1.2.3
  allow 1.2
  ```

  表示允许1.2.0.0/16，拒绝1.2.3.0/24，但是允许1.2.3.4 连接当前的 NTP server

- `deny [all] [subnet]`

  和 allow 显示，逻辑取反

- `bindaddress [address]`

  指定NTP server 监听的IP

- `clientloglimt [limit]`

  the maximum amount of memory that chronyd is allowed to allocate for logging of  client accesses and the state that chronyd as an NTP server needs to support the interleaved mode for its clients

- `noclientlog`

  client accesses are note to be logged

#### Real-time clock (RTC)

- `rtcsync`

  将 system time 周期性的拷贝到 RTC，系统不通周期不同

#### Logging

- `log [option]...`

  指定日志的格式，可以使用如下几个参数

  1. `rawmeasurements`

     生成`measurements.log`日志，具体日志格式查看man page

  2. `measurements`

     只记录有效的measurement

  3. `statistics`

     生成`statistics.log`，具体查看man page

  4. `logdir DIRECOTORY`

     日志文件存储的目录

### Conf example
server 可以从 [http://www.pool.ntp.org/](http://www.pool.ntp.org/>) 中获取指定区域的 ntp pool
```bash
# 64s 内同步一次，在chronyd进程起来的前4次内快速同步
server foo.example.net iburst minpoll 6
server bar.example.net iburst minpoll 6
server baz.example.net iburst minpoll 6
server qux.example.net iburst minpoll 6
# 如果和NTP server的偏差值超过 1 sec，会在3次内同步
makestep 1.0 3
# kernel 同步RTC时间
rtcsync
allow
clientloglimit 100000000
leapsectz right/UTC
driftfile /var/lib/chrony/drift
dumpdir /var/run/chrony
```
```bash
# Use public servers from the pool.ntp.org project.
# Please consider joining the pool (http://www.pool.ntp.org/join.html).
server 10.206.123.1 minpoll 4 maxpoll 6 iburst
server 10.206.132.45 minpoll 4 maxpoll 6 iburst
server 10.137.1.29 minpoll 4 maxpoll 6 iburst
server 10.137.1.30 minpoll 4 maxpoll 6 iburst
server 10.137.17.43 minpoll 4 maxpoll 6 iburst
server 10.137.11.124 minpoll 4 maxpoll 6 iburst
server 10.105.148.143 minpoll 4 maxpoll 6 iburst
server 10.137.31.241 minpoll 4 maxpoll 6 iburst
server 10.182.214.79 minpoll 4 maxpoll 6 iburst
server 10.182.214.74 minpoll 4 maxpoll 6 iburst


# 同步使用的端口
#acquisitionport 1123

# 存储Server时间的本地目录
dumpdir /var/run/chrony

# Ignore stratum in source selection.
stratumweight 0.01

# Record the rate at which the system clock gains/losses time.
driftfile /var/lib/chrony/drift

# 闰秒配置,17h34m消化1s
leapsecmode slew
maxslewrate 1000
smoothtime 400 0.001 leaponly

# In first three updates step the system clock instead of slew
# if the adjustment is larger than 10 seconds.
# makestep 0.1 3

## Command config
bindcmdaddress 127.0.0.1
cmdallow all

## Real Time clock(RTC)
hwclockfile /etc/adjtime
rtcautotrim 10
rtcsync

# Send a message to syslog if a clock adjustment is larger than 0.5 seconds.
logchange 0.1
log measurements statistics tracking
logdir /var/log/chrony
```
## Chronyc
chronyc 是 chronyd 的 CLI，如果没有指定参数，chronyc 会以`chronyc>`的prompt提示用户输入参数（interactive mode）
### Optional args

- `-n`

  avoid DNS lookups

- `-h`

  用于指定运行 chronyd 的 host，等同于mysql中的`-h`参数。默认使用本机的chronyd

- `-p PORT`

  指定使用chronyd 监听的udp端口，默认 323

  ```
  [vagrant@localhost ~]$ ss -lupn
  State      Recv-Q Send-Q                                                         Local Address:Port                                                                        Peer Address:Port              
  UNCONN     0      0                                                                          *:111                                                                                    *:*                  
  UNCONN     0      0                                                                          *:970                                                                                    *:*                  
  UNCONN     0      0                                                                  127.0.0.1:323                                                                                    *:*                  
  UNCONN     0      0                                                                          *:68                                                                                     *:*                  
  UNCONN     0      0                                                                       [::]:111                                                                                 [::]:*                  
  UNCONN     0      0                                                                       [::]:970                                                                                 [::]:*                  
  UNCONN     0      0                                                                      [::1]:323                                                                                 [::]:*                 
  ```

### Positional args
#### System clock

- `tracking`

  显示 system clock 的相关信息

  ```
  [vagrant@localhost ~]$ chronyc tracking
  Reference ID    : 3DEF6465 (061239100101.ctinets.com)
  Stratum         : 2
  Ref time (UTC)  : Fri Mar 25 05:51:45 2022
  System time     : 177187.984375000 seconds slow of NTP time
  Last offset     : -0.001331689 seconds
  RMS offset      : 1495.612915039 seconds
  Frequency       : 1.957 ppm slow
  Residual freq   : +0.011 ppm
  Skew            : 0.235 ppm
  Root delay      : 0.047433343 seconds
  Root dispersion : 0.002741362 seconds
  Update interval : 960.0 seconds
  Leap status     : Normal
  ```

1. Reference ID

   NTP 同步的地址，如果显示的不是IP地址，表示没有和remote source 同步，使用了local mode

2. Startnum

   到 NTP 同步地址的跳数 

3. Ref time

   UTC 对比的时间

4. System time

   机器识别的时间

5. Last offset

   local offset

6. RMS offset

   long-term average

7. Frequency

   the system's clock went rong rate

8. Skew

   Frequency error bound

9. Root delay

   delays to the startum-1 computer

10. Update interval

    the interval between the last two clock updates（时钟同步的周期）

11. Leap status

    偏移值的状态，可以是Normal, Insert second, Delete second, Not synchronised

- `makestep,makestep THRESHOLD LIMIT`

  顾名思义移动一步，即同步一次
  chronyd 通常会自动按照周期同步时间，可以使用`makestep`不带任何参数立即同步。如果使用第二种格式表示在limit内超过threshold才会同步，例如makestep 0.1 1表示在0.1内同步一次。通常和`burst`一起使用，对NTP servers 探测一次，然后使用`makestep`来同步

#### Time source

- `sources [-v]`

  显示当前chronyd使用的NTP sources，`-v`表示verbose，来显示每列表示的含义

  ```
  [vagrant@localhost ~]$ chronyc sources
  210 Number of sources = 4
  MS Name/IP address         Stratum Poll Reach LastRx Last sample               
  ===============================================================================
  ^- 120.211.39.200                2  10   377   818    +13ms[  +13ms] +/-   89ms
  ^* 061239100101.ctinets.com      1  10   377   269  +3479us[+3863us] +/-   32ms
  ^- 119.28.183.184                2  10   327   41m    +27ms[  +27ms] +/-   66ms
  ^- ntp.wdc1.us.leaseweb.net      2  10   377   343   -609us[ -226us] +/-  237ms
  ```

  1. M

     表示source使用的mode，`^`表示remote server，`=`表示peer，`#`表示locally connected reference clock

  2. S

     表示当前source 的状态，可以是`*`表示当前chronyd选中使用的source，`+`表示和被chronyd选中source一起配合使用的source，`-`表示和被chronyd选中的source不一起使用，`？`表示和source失联，或者数据包没有校验通过，`x`source 被认为falseticker（its time is inconsistent with a majority of other sources），`~`source的时间差值巨大

  3. Name/IP address

     the name or the ip address of the source ,or reference Id for reference clocks

  4. Startum

     shows the startum of the sources

  5. Poll

     以2的x次方表示，例如10表示2的10次方秒进行一次同步
     a value of 6 would indicate that a measurement is being made every 64 seconds

  6. Reach

     标识到达source的寄存器，377表示可以到达所有

  7. LastRx

     this column shows how long age the last good sample was received from the source

  8. Last Smaple

     this column shows the offset between the local clock and the source at the last measurement. the number in the square brackets shows the actual measured offset. The number to the left of the square brackets shows the original measurement. the number following +/- indicator shows the margin of error in the measurement
     中括号中的数值表示当前主机和NTP server的实际偏差(drift)，左边数值表示原始偏差，+ 表示比 NTP server 快了，- 表示比 NTP server 慢了

- `sourcestats [-v]`

  显示和NTP server 计算出的偏差详情

  ```
  [vagrant@localhost ~]$ chronyc sourcestats
  210 Number of sources = 4
  Name/IP Address            NP  NR  Span  Frequency  Freq Skew  Offset  Std Dev
  ==============================================================================
  120.211.39.200             15  10  230m     +0.750      0.571    +20ms  1987us
  061239100101.ctinets.com   51  22  424m     -0.005      0.143  -5459ns  1929us
  119.28.183.184             19  12  381m     -0.005      0.228    +25ms  1869us
  ntp.wdc1.us.leaseweb.net   50  29  449m     -0.322      0.337   -209us  5873us
  ```

  1. Name/IP Address

     the name or IP address of the NTP server or reference ID

  2. NP

     number of sample point，drift 和 offset 会从这些sample中计算

  3. NR

     number of rums of residuals having the same sign following the last regression

  4. Span

     interval between the oldest and newest smaples

  5. offset

     和source的估计差值

  6. Frequency

     表示computer clock 比 server 快或慢的时间

  7. Freq Skew

     Frequency 能到的最大值

- `reselect`

  选择使用最优的NTP server

#### NTP Sources

- `activity`

  reports the number of servers and peers that are online and offline

- `ntpdata [ADDRESS]`

  显示NTP server 的信息

  ```
  Remote address  : 61.239.100.101 (3DEF6465)
  Remote port     : 123
  Local address   : 10.0.2.15 (0A00020F)
  Leap status     : Normal
  Version         : 4
  Mode            : Server
  Stratum         : 1
  Poll interval   : 10 (1024 seconds)
  Precision       : -25 (0.000000030 seconds)
  Root delay      : 0.000000 seconds
  Root dispersion : 0.000000 seconds
  Reference ID    : 47505300 (GPS)
  Reference time  : Fri Mar 25 09:12:15 2022
  Offset          : -0.001693405 seconds
  Peer delay      : 0.054338168 seconds
  Peer dispersion : 0.000000075 seconds
  Response time   : 0.000004753 seconds
  Jitter asymmetry: +0.50
  NTP tests       : 111 111 1111
  Interleaved     : No
  Authenticated   : No
  TX timestamping : Kernel
  RX timestamping : Kernel
  Total TX        : 188
  Total RX        : 175
  Total valid RX  : 175
  ```

  具体信息查看 man page

- `add|delete server|peer ADDRESS [OPTION]`

  添加删除ntp server或peer，可以使用的option 参考配置文件

  ```
  [vagrant@localhost ~]$ sudo chronyc  add server 0.hk.pool.ntp.org  minpoll 6 maxpoll 10 key 25
  511 Source already present
  ```

- `burst GOOD/MAX [ADDRESS]`

  make a set of measurements to each of its NTP sources over a short duartion. chronyd 会立即对NTP server做一次measure
  good 表示chronyd想从 NTP server 得到的反馈good的次数
  max 表示chronyd burst的最大次数
  address 表示测试的ntp server，默认对配置的所有 ntp server 进行burst

  ```
  [vagrant@localhost ~]$ sudo chronyc burst 2/3
  200 OK
  ```

  表示chronyd需要获取到所有配置的ntp server 2 次 good measure才会停止探测，否则会一次探测到3次

- `maxdelay | maxpoll | minpoll ADDRESS NUMBER`

  等同于配置文件中相对的指令块

- `offline | online [ADDRESS]`

  NTP server是否被使用

- `refresh`

  强制 chronyd 解析 NTP server

#### Manual time input

- `manual on | off`

  是否可以使用`settime`指令来手动设置时间

- `settime TIME`

  手动设置时间，time的格式可以是如下几种。设置时间后chronyd仍会去同步正确的时间

  ```
  settime 16:30
  settime 16:30:05
  settime Nov 21, 2015 16:30:05
  
  [vagrant@localhost ~]$ date
  Fri Mar 25 09:50:45 UTC 2022
  [vagrant@localhost ~]$ sudo chronyc manual on
  200 OK
  [vagrant@localhost ~]$ sudo chronyc settime 13:00
  200 OK
  Clock was -11344.40 seconds fast.  Frequency change = 0.00ppm, new frequency = -2.02ppm
  [vagrant@localhost ~]$ date
  Fri Mar 25 09:51:00 UTC 2022
  ```

#### NTP access
关联本机 NTP server 具体不展开
#### other daemon commands

- `cyclelogs`

  看例子

  ```
  # mv /var/log/chrony/measurements.log /var/log/chrony/measurements1.log
  # chronyc cyclelogs
  # ls -l /var/log/chrony
  -rw-r--r--   1 root     root            0 Jun  8 18:17 measurements.log
  -rw-r--r--   1 root     root        12345 Jun  8 18:17 measurements1.log
  # rm -f measurements1.log
  ```

- `dump`

  查看chronyd的dumpfile，对应指令块`dumpdir`

- `shutdown`

  关闭chronyd进程

## chronyc with date
可以使用如下的脚本来观察时间的变化
```bash
#!/bin/bash

for ((i=0;i<10;i++))
do
        tput sc; tput civis                     
        echo -ne $(date +'%Y-%m-%d %H:%M:%S')   
        sleep 1
        tput rc                                
done
tput el;tput cnorm                           
```
如果在使用了`date`手动同步时间，如果没有关闭 chronyd 或其他的 NTP 同步进程还是会使用 NTP 来同步，但是需要按照`minpoll`来同步，如果`minpoll`的值太大同步的周期就会很慢
### 0x001 chronyd运行，date手动修改时间
```bash
[root@localhost vagrant]# date
Tue Mar 29 02:08:46 UTC 2022
[root@localhost vagrant]# ps -ef | grep -v grep | grep chronyd
chrony    4423     1  0 Mar28 ?        00:00:00 chronyd
chrony    5135     1  0 02:01 ?        00:00:00 /usr/sbin/chronyd
[root@localhost vagrant]# date
Tue Mar 29 02:09:38 UTC 2022
[root@localhost vagrant]# date -s '2022-03-29'
Tue Mar 29 00:00:00 UTC 2022
[root@localhost vagrant]# date
Tue Mar 29 00:00:01 UTC 2022

[vagrant@localhost ~]$ date
Tue Mar 29 00:36:04 UTC 2022

[root@localhost vagrant]# sudo chronyc sources
210 Number of sources = 4
MS Name/IP address         Stratum Poll Reach LastRx Last sample               
===============================================================================
^- makaki.miuku.net              2   6   377    46   -554ms[ -554ms] +/-   48ms
^+ 139.199.214.202               2   6   337    52   -546ms[ -546ms] +/-   36ms
^? ntp8.flashdance.cx            2   6   373   112   -545ms[ +126ms] +/-  200ms
^? time.neu.edu.cn               1   6     1     3   -719ms[ -540ms] +/-   23ms

[root@localhost vagrant]# sudo chronyc sources
210 Number of sources = 4
MS Name/IP address         Stratum Poll Reach LastRx Last sample               
===============================================================================
^- makaki.miuku.net              2   6   377    52   -679ms[ -679ms] +/-   49ms
^- 139.199.214.202               2   6   177    58  -7733us[-1013ms] +/-   36ms
^? ntp8.flashdance.cx            2   6   355   117  -6809us[ -722ms] +/-  180ms
^* time.neu.edu.cn               1   6     7    10    +90us[ -663ms] +/-   21ms

#会选定ntp server，但是不会同步时间，需要达到minpoll会才会同步会
[root@localhost vagrant]# date
Tue Mar 29 00:07:35 CST 2022
```
### 0x002 chronyd运行，date手动修改时间，chronyc同步
```bash
[root@localhost vagrant]# date -s '2022-03-29'
Tue Mar 29 00:00:00 CST 2022
[root@localhost vagrant]# date
Tue Mar 29 00:00:01 CST 2022
[root@localhost vagrant]# sudo chronyc burst 3/3
200 OK
[root@localhost vagrant]# date
Tue Mar 29 00:00:19 CST 2022
[root@localhost vagrant]# sudo chronyc makestep
200 OK
[root@localhost vagrant]# date
Tue Mar 29 11:28:55 CST 2022
```
### 0x003 重启chronyd
chronyd重启也会自动同步
```bash
[root@localhost vagrant]# date -s '2022-03-29'
Tue Mar 29 00:00:00 CST 2022
[root@localhost vagrant]# date
Tue Mar 29 00:00:20 CST 2022
[root@localhost vagrant]# systemctl restart chronyd
[root@localhost vagrant]# sudo chronyc sources
210 Number of sources = 4
MS Name/IP address         Stratum Poll Reach LastRx Last sample               
===============================================================================
^? tock.ntp.infomaniak.ch        0   6     0     -     +0ns[   +0ns] +/-    0ns
^? ntp6.flashdance.cx            2   6     1     2  -4108us[-41850s] +/-  198ms
^* 79.133.44.136                 1   6    17     6  -4825us[-41850s] +/-  138ms
^+ ntp1.ams1.nl.leaseweb.net     2   6    17    10    +15ms[  +15ms] +/-  205ms
[root@localhost vagrant]# date
Tue Mar 29 11:38:52 CST 2022
```
## note

1. 如果需要立即同步可以使用，`chronyc burst 3/3;chronyc makestep`
1. 在使用NTP时，需要确保timezone是正确的，可以使用`datetimectl`来设置和查看
1. 只有NTP server 都是可达的才会正确同步时间
1. 如果NTP server 的同步周期慢，可以尝试修改`minpoll`的值（其中包括measurement 和 makestep）的周期
