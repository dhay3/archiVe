# linux ntpq

参考：

https://detailed.wordpress.com/2017/10/22/understanding-ntpq-output/

syntax：`ntpq [options] [host...]`

ntpq是NTP的查询工具有interactive和cli两种方式，会将请求发送到ntp服务器。默认使用localhost，如果没有指定host。==NTP是一个UDP协议，所以连接不是可靠的==

## options

- n | --numeric

  对host不做解析，以数字的格式输出

- -p | --peers

  当前使用的ntp server的信息

  ```
  cpl in /var/lib/ntp λ ntpq -pn 127.0.0.1                 
       remote           refid      st t when poll reach   delay   offset  jitter
  ==============================================================================
  *84.16.73.33     .GPS.            1 u  978 1024    7  222.338  +10.820   1.841
  -116.203.151.74  131.188.3.222    2 u  948 1024    7  262.759  +21.756  13.279
  +139.199.215.251 100.122.36.4     2 u  959 1024    7   33.316   +6.734   9.906
  +5.79.108.34     130.133.1.10     2 u  947 1024    7  231.657  +28.945  12.221
  ```

  1. remote：同步服务器
  2. refid：同步服务器参考的服务器
  3. st：sratum，0表示根服务器，1表示直接参考跟服务器的服务器，以此类推16表示无同步的服务器
  4. t：服务器的种类
  5. when：最后一个数据包接受的时间，如果为`-`表示从未收到
  6. poll：向服务器请求的间隔
  7. reach：到达位移寄存器
  8. delay：往返时延
  9. offset：服务器和主机的偏移量
  10. jitter：估计错误的偏移量

  ```
  * Synchronized to this peer
  # Almost synchronized to this peer
  + Peer selected for possible synchronization
  – Peer is a candidate for selection
  ~ Peer is statically configured 
  ```

  

