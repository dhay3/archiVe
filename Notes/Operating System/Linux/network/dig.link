# Linux dig

ref:

https://www.rfc-editor.org/rfc/rfc1035.txt

## Digest

syntax:

```
dig  [@server]  [-b  address] [-c class] [-f filename] [-k filename]
       [-m]  [-p  port#]  [-q  name]  [-t  type]   [-v]   [-x   addr]   [-y
       [hmac:]name:key] [ [-4] | [-6] ] [name] [type] [class] [queryopt...]
```

dig 和 nslookup 一样，用于查询 DNS 记录。还可使用开源的 dog，比较容易。

通常使用`dig @server name type`来查询记录

## Terms

1. server

   is the name or IP address of the name server to query

2. name

   is the name of the resource record that is to be looked up

3. type

   indicates what type of query is required - ANY, A, MX, SIG, etc.

## Tips

1. dig 会查询`/etc/resolv.conf`中的每一行 nameserve，如果没有得到结果才会去查询本地 host。==注意这一点和正常的 dns 查询不一样。正常的 dns 会先去查询 local host==。

   假设，host 文件中有如下配置

   `baidu.com 127.0.0.1`

   但是 resolve 配置文件中的 dns server 能获取到对应的 dns record。那么 dig 的结果就是 dns record，但是 ping 的结果是 127.0.0.1

2. 如果没有指定任何 positional 和 optional args，dig 默认会去查询根域(“.”)

3. 如果没有指定 type，dig 默认查询 A record

## Output

以查询 root 为例子

```
 <<>> DiG 9.18.3 <<>>
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 2051
;; flags: qr rd ra; QUERY: 1, ANSWER: 13, AUTHORITY: 0, ADDITIONAL: 14
```

第一部分是查询参数和统计

```
;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4000
;; QUESTION SECTION:
;.                              IN      NS
```

第二部是查询，表示查询 root 的 NS record

```
;; ANSWER SECTION:
.                       472473  IN      NS      k.root-servers.net.
.                       472473  IN      NS      l.root-servers.net.
...
```

第三部分是回答，表示 root 对应的 NS record。有13条 NS record，472473 TTL（Time to live）缓存时间（和 IP 的  TTL 需做区分）

```
;; ADDITIONAL SECTION:
k.root-servers.net.     40473   IN      A       193.0.14.129
l.root-servers.net.     40473   IN      A       199.7.83.42
```

第四部分是对第三部分的补充，即 NS record 对应的 A record

```
;; Query time: 10 msec
;; SERVER: 30.30.30.30#53(30.30.30.30) (UDP)
;; WHEN: Fri Jul 15 22:42:22 HKT 2022
;; MSG SIZE  rcvd: 460
```

第五部分是查询过程信息，表示 dns server 是 30.30.30.30:53，回送的报文是 460 byte

## Optional args

- `-4 | -6`

  only IPv4 or IPv6 should be used

  只会显示 IPv4 或者 IPv6 的记录

- `-b address[#port]`

  sets thee source IP address of the query

  使用的指定的IP查询dns server

- `-c class`

  使用指定的 class 查询，默认 IN。其他可使用值具体查考 RFC 

- `-f file`

  batch mode. 批量查询 Domain name

- `-p port`

  指定查询的 DNS server 的端口，默认 53

- `-q name`

  this option specifies the domain name to query

- `-t type`

  this option indicates the resurce record type to query, which can be any valid query type. 默认 A record。其他可使用值具体参考 RCF 

- `-x addr`

  reverse lookups, for mapping addresses to names

  查询 PTR record

## Query options

dig 提供 query options，用来修改查询的方式和查询的结果

1. each query option is identified by a keyword preceded by a plus sign(+)
2. keywords may be abbreviated
3. no to negate the meaning of that keyword

> 下面 no 部分不被展示

### Query args

- `+tcp`

  whether to use TCP when querying name servers

  ```
  cpl in ~ λ dig +tcp baidu.com
  ...
  ;; Query time: 6 msec
  ;; SERVER: 30.30.30.30#53(30.30.30.30) (TCP)
  ;; WHEN: Fri Jul 15 23:22:46 HKT 2022
  ;; MSG SIZE  rcvd: 70
  
  ```

- `+dscp=value`

  this option sets the DSCP code point to be used when sending the query. By default no code point is explicitly set

- `+https[=value]`

  是否使用 DoH，如果使用了该参数端口默认443

  ```
  cpl in ~ λ dig +https @dns.alidns.com  baidu.com
  ```

- `+tls`

  whether to use DNS over TLS(DoT)

- `+recurse`

  是否递归查询DNS

- `+timeout=T`

  sets the timeout for a query to T seconds. The default timeout is 5 seconds. An attempt to set T to less than 1 is silently set to 1

- `+trace`

  从根域开始做DNS查询

### Output args

- `+additional`

  是否展示 additional 部分

- `+answer`

  是否展示 answer 部分

- `+short | +shor`

  只输出最后的结果

  ```
  cpl in ~ λ dig +shor baidu.com
  220.181.38.251
  220.181.38.148
  ```

- `+yaml`

  以 yml 的格式输出内容
