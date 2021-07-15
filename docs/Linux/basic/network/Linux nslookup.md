# Linux nslookup

nslookup用于查询DNS address，有两种模式interactive和non-interactive。当没有指定参数时进入interactive mode，可以查询多个domain。non-interactive只可以查询一个domain可以使用的参数和interactive指令相同，但是按`nslookup [-options] <query name> [ domain_name_server]`格式

## 参数

- server [domain]

  查询时使用的dns，没有指定时默认使用`/etc/resolv.conf`中内容

  ```
   cpl in / λ nslookup baidu.com 1.1.1.1
  Server:         1.1.1.1
  Address:        1.1.1.1#53
  
  Non-authoritative answer:
  Name:   baidu.com
  Address: 220.181.38.148
  Name:   baidu.com
  Address: 39.156.69.79
  ```

- set keyword=[value]

  all：打印出当前使用的配置

  type | q：使用指定DNS记录类型查看

  ```
  cpl in / λ nslookup        
  > set q=RTP
  unknown query type: RTP
  > set q=PTR
  > 127.0.0.1
  Server:         30.14.128.1
  Address:        30.14.128.1#53
  
  1.0.0.127.in-addr.arpa  name = localhost.
  
  
  cpl in / λ nslookup -q=PTR 127.0.0.1
  Server:         30.14.128.1
  Address:        30.14.128.1#53
  
  1.0.0.127.in-addr.arpa  name = localhost.
  ```

  