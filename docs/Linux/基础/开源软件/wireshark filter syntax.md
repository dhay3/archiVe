# wireshark filter syntax

参考：https://www.wireshark.org/docs/man-pages/wireshark-filter.html

> [cheat sheet](https://www.comparitech.com/net-admin/wireshark-cheat-sheet/)

## oprator

- eq, ==    Equal ，精确匹配  
- ne, !=    Not Equal    
- gt, >     Greater Than   
-  lt, <     Less Than    
- ge, >=    Greater than or Equal to    l
- e, <=    Less than or Equal to

### logical expression

- and, &&   Logical AND    
- or,  ||   Logical OR    
- not, !    Logical NOT

### search and match operators

- contains 使用模糊查询匹配
- matches,~ 使用regex匹配，忽略大小写

例如：匹配URL

```
http contains "https://www.wireshark.org"
```

### slice operator

和golang一样wireshark也可以使用slice，例如

```
eth.src[0:3] == 00:00:83
llc[0] eq aa
#也可以是数组
frame[4] == 0xff
```

### membership operator

可以使用curly bracket来描述范围，例如

```
tcp.port in {80 443 8080}
```

等价于

```
tcp.port == 80 or tcp.port == 443 or tcp.port == 8080
```

也可以通过这种方式来过滤请求指定的方法

```
http.request.method in {"HEAD" "GET"}
```

### Type conversions

wireshark 支持将字符转为ASCii的方式表达过滤式

```
http.request.method == "GET"
#等价
http.request.method == 47.45.54
```















