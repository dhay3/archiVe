ref：
[https://github.com/curl/curl](https://github.com/curl/curl)
[https://curl.se/docs/faq.html](https://curl.se/docs/faq.html)
[https://curl.se/libcurl/c/curl_easy_getinfo.html](https://curl.se/libcurl/c/curl_easy_getinfo.html)
[https://curl.se/docs/sslcerts.html](https://curl.se/docs/sslcerts.html)

https://blog.cloudflare.com/a-question-of-timing/

https://curl.se/mail/archive-2015-12/0011.html

## 0x1 Digest
syntax：`curl [options] [URL...]`
curl是一个基于libcurl的CLI工具，支持模拟多种协议请求

## 0x2 Terms
### progress meter
在使用 curl 时通常会显示 progress meter ——即进度，显示相关数据，单位以byte（所以可以通过这种方式来计算下载的速率，等同于wget）。但是服务端回显数据到请求端时不会显示进度条，如果想要显示进度条，可以使用shell redirct`>`或`-o`将内容重定向到文件
```bash
root in /home/ubuntu λ curl baidu.com
<html>
<meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
</html>
root in /home/ubuntu λ curl baidu.com > a
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    81  100    81    0     0   5062      0 --:--:-- --:--:-- --:--:--  5062
```
如果想要以进度条的方式替代meter，可以使用`-#`或`--progress-bar`
```bash
root in /home/ubuntu λ curl --progress-bar baidu.com > a
######################################################################################################################################### 100.0%
```
如果需要关闭progress meter 可以使用`-s`或`--silent`

### protocol

curl 支持如下多种协议，只介绍其中几种比较特殊的

#### dict

具体查看：https://github.com/dhay3/archive/blob/master/Docs/Net/0x1%20Digest.md

lets your lookup words using online dictionaries

```
cpl in ~ λ curl dict.org/d:hello
```

### file

具体查看：

read or write local files，curl 在linux上不支持远程读取文件但是在windows上支持

### gopher

### ldap





## 0x3 positional args

curl 只有一种 positional args，即 URL。通常URL需要指定协议，但是如果没有指定协议（protocol://），curl默会根据URL来猜测使用的协议，通常是http，但是如果URL中包含`ftp.`就会使用ftp协议，同理其他类似的URL。针对`multipart/form-data`curl会复用tcp连接
### globbing
curl可以使用globbing(为了方便记忆我这个功能叫做globbing，内置的表示即使在windows上可以使用，区别与linux中的globbing)
```bash
http://site.{one,two,three}.com
#等价于linux中的{1..100}
ftp://ftp.example.com/file[1-100].txt
http://example.com/archive[1996-1999]/vol[1-4]/part{a,b,c}.html
```
当使用curl globbing 时，URL通常需要加上双引号，否则可能会被 shell globbing 解析

```
 "http://[fe80::3%25eth0]/"
```



## 0x4 optional args
所有和boolean相关的参数，都可以添加前缀`no`，例如`--option`，如果需要取反可以使用`--no-option`（after 7.19.0）
### common optoinal args

-  `--create-dirs`
和`-o`一起使用，生成目录 
- `-g | --globoff`

关闭url globbing parser的功能，如果是IPv6的地址需要指定改参数，因为IPv6的地址默认是在`[]`中的

-  `-o <filename>`
将服务器的回应保存成文件，等同于`wget`命令。 
上面命令将`www.example.com`保存成`example.html`。 
```
$ curl -o example.html https://www.example.com
```

-  `-m | --max-times <sec>`
执行curl的最长响应时间，一般用于批量脚本 
-  `-O`参数将服务器回应的==文件(url对应的只能是文件)==保存成文件，并将 URL 的最后部分当作文件名。 
上面命令将服务器回应保存成文件，文件名为`bar.html`。 
```
$ curl -O https://www.example.com/foo/bar.html
```

-  `-Z | --parallel`
并发响应，可以加快速度 
-  `--retry-delay <sec>| --retry-max-time <sec>| retry <num>`
curl重试时间和次数 
-  `-s`
将不输出错误和进度信息(也可以使用)`--no-progress-meter`。 
下面命令一旦发生错误，不会显示错误信息。不发生错误的话，会正常显示运行结果。
如果想让 curl 不产生任何输出，可以使用下面的命令(只限于输出在stdout的内容)。  
```
$ curl -s https://www.example.com
$ curl -s -o /dev/null https://google.com
```

-  `-S`
指定只输出错误信息，通常与`-s`一起使用。 
下面命令没有任何输出，除非发生错误。 
```
$ curl -Ss -o /dev/null https://google.com
```

-  `-v`参数输出通信的整个过程，用于调试。 
`--trace`参数也可以用于调试，还会输出原始的二进制数据。  
```
cpl in /sharing/conf λ curl -v baidu.com
*   Trying 220.181.38.251:80...
* Connected to baidu.com (220.181.38.251) port 80 (#0)
> GET / HTTP/1.1
> Host: baidu.com
> User-Agent: curl/7.77.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Thu, 12 Aug 2021 03:01:16 GMT
< Server: Apache
< Last-Modified: Tue, 12 Jan 2010 13:48:00 GMT
< ETag: "51-47cf7e6ee8400"
< Accept-Ranges: bytes
< Content-Length: 81
< Cache-Control: max-age=86400
< Expires: Fri, 13 Aug 2021 03:01:16 GMT
< Connection: Keep-Alive
< Content-Type: text/html
< 
<html>
<meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
</html>
* Connection #0 to host baidu.com left intact
```
```
$ curl --trace - https://www.example.com
```

-  `--trace-time`
和`-v`结合一起使用时，输出timestamp  
```
cpl in /sharing/conf λ curl -v --trace-time baidu.com
10:49:41.305257 *   Trying 220.181.38.148:80...
10:49:41.342109 * Connected to baidu.com (220.181.38.148) port 80 (#0)
10:49:41.342252 > GET / HTTP/1.1
10:49:41.342252 > Host: baidu.com
10:49:41.342252 > User-Agent: curl/7.77.0
10:49:41.342252 > Accept: */*
10:49:41.342252 > 
10:49:41.379880 * Mark bundle as not supporting multiuse
10:49:41.379938 < HTTP/1.1 200 OK
10:49:41.379960 < Date: Thu, 12 Aug 2021 02:49:43 GMT
10:49:41.379981 < Server: Apache
10:49:41.380000 < Last-Modified: Tue, 12 Jan 2010 13:48:00 GMT
10:49:41.380022 < ETag: "51-47cf7e6ee8400"
10:49:41.380042 < Accept-Ranges: bytes
10:49:41.380066 < Content-Length: 81
10:49:41.380085 < Cache-Control: max-age=86400
10:49:41.380107 < Expires: Fri, 13 Aug 2021 02:49:43 GMT
10:49:41.380128 < Connection: Keep-Alive
10:49:41.380144 < Content-Type: text/html
10:49:41.380163 <
```

-  `--trace <file>`
输出到示详细的信息，可以使用`-`表示到stdout  
```
cpl in /sharing/conf λ curl --trace - baidu.com
== Info:   Trying 220.181.38.251:80...
== Info: Connected to baidu.com (220.181.38.251) port 80 (#0)
=> Send header, 73 bytes (0x49)
0000: 47 45 54 20 2f 20 48 54 54 50 2f 31 2e 31 0d 0a GET / HTTP/1.1..
0010: 48 6f 73 74 3a 20 62 61 69 64 75 2e 63 6f 6d 0d Host: baidu.com.
0020: 0a 55 73 65 72 2d 41 67 65 6e 74 3a 20 63 75 72 .User-Agent: cur
0030: 6c 2f 37 2e 37 37 2e 30 0d 0
......
```

-  `--stderr <file>`
将stderr重定向到文件中  
```
root in /home/ubuntu λ curl --stderr a dafdafaf
root in /home/ubuntu λ cat a
curl: (6) Could not resolve host: dafdafaf
```

-  `-w | --write-out <format>`
format可以使用`@filename`的格式从文件中读取或者`@-`从stdin中读取。使用`%{variable_name}`来设置变量
variable_name具体可以使用值查看manual page  
```
root in /home/ubuntu λ curl -so /dev/null -w %{url_effective} baidu.com
http://baidu.com/#
```

### ftp/sftp

-  `-a`
以增量模式上传文件 
### dns

- `--dns-servers <addresses>`

使用指定的 dns 服务器而不是系统默认的
### ssl

- `-k`参数指定跳过 SSL 检测。 

 默认每个SSL链接会去校验是否安全，这个参数不会检查服务器的 SSL 证书是否正确。 
```
$ curl -k https://www.example.com
```

- `--cacert <file>`
指定使用的ca证书（即服务器回送的证书），可以现在浏览器上把证书下载下来
- `--cert-type <type>`

ca证书的类型，默认`PEM`

- `--cert`

### http

-  `-A <user_agent>`
参数指定客户端的用户代理标头，即`User-Agent`。curl 的默认用户代理字符串是`curl/[version]`。
下命令将`User-Agent`改成 Chrome 浏览器。 
```bash
$ curl -A 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36' https://google.com
```
下命令会移除`User-Agent`标头。
```bash
$ curl -A '' https://google.com
```
也可以通过`-H`参数直接指定标头，更改`User-Agent`。  
```


$ curl -H 'User-Agent: php/1.0' https://google.com
```

-  `-b <cookie | filename>`
参数用来向服务器发送 Cookie。 
上面命令会生成一个标头`Cookie: foo=bar`，向服务器发送一个名为`foo`、值为`bar`的 Cookie。 
上面命令发送两个 Cookie。如果没有在参数内容中使用`=`即表示filname 
上面命令读取本地文件`cookies.txt`，里面是服务器设置的 Cookie（参见`-c`参数），将其发送到服务器。 
```
$ curl -b 'foo=bar' https://google.com
$ curl -b 'foo1=bar;foo2=bar2' https://google.com
$ curl -b cookies.txt https://www.google.com
```

-  `-c|--cookie-jar`
将服务器返回的cookie保存到本地 
上面命令将服务器的 HTTP 回应所设置 Cookie 写入文本文件`cookies.txt`。 
```
$ curl -c cookies.txt https://www.google.com
```

-  `-d`
参数用于发送 POST 请求的数据体。 
使用`-d`参数以后，HTTP 请求会自动加上标头`Content-Type : application/x-www-form-urlencoded`。并且会自动将请求转为 POST 方法，因此可以省略`-X POST`。
`-d`参数可以读取本地文本文件的数据，向服务器发送。 
上面命令读取`data.txt`文件的内容，作为数据体向服务器发送。 
```
$ curl -d'login=emma＆password=123'-X POST https://google.com/login
# 或者
$ curl -d 'login=emma' -d 'password=123' -X POST  https://google.com/login

$ curl -d '@data.txt' https://google.com/login
```

-  `--data-urlencode`
参数等同于`-d`，发送 POST 请求的数据体，区别在于会自动将发送的数据进行 URL 编码。 
上面代码中，发送的数据`hello world`之间有一个空格，所以需要进行 URL 编码。 
```
$ curl --data-urlencode 'comment=hello world' https://google.com/login
```

-  `-D | --dump-header <filename>`
将接受到的header写入文件  
```
root in /home/ubuntu λ curl -D a baidu.com
<html>
<meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
</html>
root in /home/ubuntu λ cat a
HTTP/1.1 200 OK
Date: Tue, 28 Sep 2021 05:16:20 GMT
Server: Apache
Last-Modified: Tue, 12 Jan 2010 13:48:00 GMT
ETag: "51-47cf7e6ee8400"
Accept-Ranges: bytes
Content-Length: 81
Cache-Control: max-age=86400
Expires: Wed, 29 Sep 2021 05:16:20 GMT
Connection: Keep-Alive
Content-Type: text/html
```

-  `-G | --get`
当和`-d`一起使用是，使用该参数强制发送GET请求 
-  `-x`
指定 HTTP 请求的代理。 
上面命令指定 HTTP 请求通过`myproxy.com:8080`的 socks5 代理发出。
如果没有指定代理协议，默认为 HTTP。 
上面命令中，请求的代理使用 HTTP 协议。 
```
$ curl -x socks5://james:cats@myproxy.com:8080 https://www.example.com
$ curl -x james:cats@myproxy.com:8080 https://www.example.com
```

-  `--socks5 <host:port>`
使用socks5代理，如果没有指定port默认是1080。默认会覆盖`-x`参数 
-  `-X`
指定 HTTP 请求的方法。 
上面命令对`https://www.example.com`发出 POST 请求。也可以使用环境变量http_proxy(这个是特殊变量必须小写)和HTTPS_PROXY 
```
$ curl -X POST https://www.example.com
```

-  `-0`
强制使用HTTP1.0发送请求，不保证响应的也是HTTP1.0  
```
cpl in /tmp λ curl -0v baidu.com
*   Trying 220.181.38.148:80...
* Connected to baidu.com (220.181.38.148) port 80 (#0)
> GET / HTTP/1.0
> Host: baidu.com
> User-Agent: curl/7.77.0
> Accept: */*
>
```

-  `--http1.1 | --http2`
使用http1.1发送请求，http2需要目标libcurl支持http2.0 
-  `-e`
参数用来设置 HTTP 的标头`Referer`，表示请求的来源。 
上面命令将`Referer`标头设为`https://google.com?q=example`。
`-H`参数可以通过直接添加标头`Referer`，达到同样效果。  
```
curl -e 'https://google.com?q=example' https://www.example.com
curl -H 'Referer: https://google.com?q=example' https://www.example.com
```

-  `-H`
添加 HTTP 请求的标头。 
上面命令添加 HTTP 标头`Accept-Language: en-US`。 
上面命令添加两个 HTTP 标头。 
上面命令添加 HTTP 请求的标头是`Content-Type: application/json`，然后用`-d`参数发送 JSON 数据。 
```
$ curl -H 'Accept-Language: en-US' https://google.com
$ curl -H 'Accept-Language: en-US' -H 'Secret-Message: xyzzy' https://google.com
$ curl -d '{"login": "emma", "pass": "123"}' -H 'Content-Type: application/json' https://google.com/login
```

-  `-F`
用来向服务器上传二进制文件。 
上面命令会给 HTTP 请求加上标头`Content-Type: multipart/form-data`，然后将文件`photo.png`作为`file`字段上传。
`-F`参数可以指定 MIME 类型。 
上面命令指定 MIME 类型为`image/png`，否则 curl 会把 MIME 类型设为`application/octet-stream`。
`-F`参数也可以指定文件名。 
上面命令中，原始文件名为`photo.png`，但是服务器接收到的文件名为`me.png`。 
```
$ curl -F 'file=@photo.png' https://google.com/profile
$ curl -F 'file=@photo.png;type=image/png' https://google.com/profile
$ curl -F 'file=@photo.png;filename=me.png' https://google.com/profile
```

-  `-f | --fail`
如果抓取指定网站的内容不显示，会默认显示一个错误页面，使用改参数内容将不显示，而是返回exit code。但是一般的网站重定向后会展示一个非默认的错误页面，所以不显示exit code  
```
cpl in /tmp λ curl -fL http://ftp.cn.debian.org/debian/aa
curl: (22) The requested URL returned error: 404
cpl in /tmp λ curl -L http://ftp.cn.debian.org/debian/aa 
<html>
<head><title>404 Not Found</title></head>
<body>
<center><h1>404 Not Found</h1></center>
<hr><center>openresty</center>
</body>
</html>
```

-  `-i`
参数打印出服务器回应的 HTTP 响应头。 
上面命令收到服务器回应后，先输出服务器回应的响应头，然后空一行，再输出网页的源码。 
```bash
$ curl -i https://www.example.com
```

-  `-I|--head`
参数向服务器发出 HEAD 请求，然会将服务器返回的 HTTP 标头打印出来。 
上面命令输出服务器对 HEAD 请求的回应。
```
$ curl -I https://www.example.com
$ curl --head https://www.example.com
```

-  `--keepalive-time <seconds>`
发送keepalive信号的间隔，默认60sec，tcp连接保活 
-  `--no-keepalive`
不使用keepalive 
-  `-L`
会让 HTTP 请求跟随服务器的重定向（http status code是3xx）。curl 默认不跟随重定向。  
```
$ curl -L -d 'tweet=hi' https://api.twitter.com/tweet
```

-  `--limit-rate`
用来限制 HTTP 请求和回应的带宽，模拟慢网速的环境。  
```
$ curl --limit-rate 200k https://google.com
```

-  `-u`
用来设置服务器认证的用户名和密码。 
上面命令设置用户名为`bob`，密码为`12345`，然后将其转为 HTTP 标头`Authorization: Basic Ym9iOjEyMzQ1`。
curl 能够识别 URL 里面的用户名和密码。 
上面命令能够识别 URL 里面的用户名和密码，将其转为上个例子里面的 HTTP 标头。 
上面命令只设置了用户名，执行后，curl 会提示用户输入密码。 
```
$ curl -u 'bob:12345' https://google.com/login
$ curl https://bob:12345@google.com/login
$ curl -u 'bob' https://google.com/login
```

-  `-T | --upload <file>`
上传文件到remote url，如果是https使用PUT方法  
```
cpl in /sharing/conf λ curl -LT cus-alias http://ftp.cn.debian.org/debian
<html>
<head><title>405 Not Allowed</title></head>
<body>
<center><h1>405 Not Allowed</h1></center>
<hr><center>openresty</center>
</body>
</html>
```
## 0x5 Timing

![2022-04-17_22-05](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220417/2022-04-17_22-05.75xnwr9nn780.webp)



## 0x6 Exit code

> man `/exit codes`


在使用curl时，exit code会记录在括号中，例如

```
curl: (56) Recv failure: Connection reset by peer

curl: (52) Empty reply from server
```

## 常用命令

- 校验时耗
```bash
for cnt in {1..100}; do
curl -SsLo /dev/null -w "%{remote_ip} %{time_namelookup} %{time_connect} %{time_appconnect} %{time_redirect} %{time_pretransfer} %{time_starttransfer} %{time_total}\n" https://taobao.api.xixikf.cn
done

for ($i=0;$i -lt 100;$i++){
cmd /r curl -SsLo 1 -w "%{remote_ip} %{time_namelookup} %{time_connect} %{time_appconnect} %{time_redirect} %{time_pretransfer} %{time_starttransfer} %{time_total}\n" https://taobao.api.xixikf.cn
}

for /l %x in (1 ,1 ,100) do curl -SsLo 1 -w "%{remote_ip} %{time_namelookup} %{time_connect} %{time_appconnect} %{time_redirect} %{time_pretransfer} %{time_starttransfer} %{time_total}\n" https://taobao.api.xixikf.cn
```

- 打流
```bash
chz_ips=(baidu.com 10.76.25.33:80);MAX_WAIT=2; \
for cnt in {1..100}; do
for idx in "${!chz_ips[@]}"; do
rst=$(curl -w "%{remote_ip} %{remote_port} %{time_namelookup} %{time_connect} %{time_starttransfer} %{time_total}" -so /dev/null -m$MAX_WAIT "${chz_ips[idx]}")
ret_code=$?
timestamp=$(date +"%F %T")
if (( ${ret_code} == 0 || ${ret_code} == 52 ));then
awk -v time="${timestamp}" '{printf("> %s Connected to %24s:%-5s%2snamelookup [%s]%2sconnect [%s]%2sstarttransfer [%s]%2stotal [%s]\n",time,$1,$2," ",$3," ",$4," ",$5," ",$6)}' <<< "${rst}"
else
awk -v time="${timestamp}" -v red="\033[31m" -v normal="\033[0;39m" '{printf("* %s%s Failed to Connect to %16s:%-5s%s\n",red,time,$1,$2,normal)}' <<< "${rst}"
fi
done
done; \
unset chz_ips MAX_WAIT rst ret_code timestamp
```

-  重定向stoud，stderr中的内容
```bash
curl -o /dev/null -svm1 >&2 <host>
```

- 不输出所有的错误正常的信息，只显示verbose信息
```bash
url -fsSvkL -o /dev/null https://baidu.com
```
