# curl

## 概述

curl是一个CLI工具用于发送WEB请求。如果不指定schema默认使用http。执行一次curl，默认会使用长连接。只进行一次tcp连接。

所有和boolean相关的选择都以`--option`和`--no-option`的格式来设置

## progress bar

curl 如果没有保存文件或是下载文件，默认不会显示bar。可以使用`-o`参数来显示bar，或是`-#`

## common

- `--create-dirs`

  和`-o`一起使用，生成目录

- `-o <filename>`

  将服务器的回应保存成文件，等同于`wget`命令。

  ```
  $ curl -o example.html https://www.example.com
  ```

  上面命令将`www.example.com`保存成`example.html`。

- `-m | --max-times <sec>`

  执行curl的最长响应时间，一般用于批量脚本

- `-O`参数将服务器回应的==文件(url对应的只能是文件)==保存成文件，并将 URL 的最后部分当作文件名。

  ```
  $ curl -O https://www.example.com/foo/bar.html
  ```

  上面命令将服务器回应保存成文件，文件名为`bar.html`。

- `-Z | --parallel`

  并发响应，可以加快速度

- `--retry-delay <sec>| --retry-max-time <sec>| retry <num>`

  curl重试时间和次数

- `-s`

  将不输出错误和==进度信息(也可以使用`--no-progress-meter`)==。

  ```
  $ curl -s https://www.example.com
  ```

  上面命令一旦发生错误，不会显示错误信息。不发生错误的话，会正常显示运行结果。

  如果想让 curl 不产生任何输出，可以使用下面的命令。

  ```
  $ curl -s -o /dev/null https://google.com
  ```

- `-S`

  指定只输出错误信息，通常与`-s`一起使用。

  ```
  $ curl -Ss -o /dev/null https://google.com
  ```

  上面命令没有任何输出，除非发生错误。

- `-v`参数输出通信的整个过程，用于调试。

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

  `--trace`参数也可以用于调试，还会输出原始的二进制数据。

  ```
  $ curl --trace - https://www.example.com
  ```

- `--trace-time`

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

- `--trace <file>`

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

- `--stderr <file>`

  将stderr重定向到文件中

  ```
  root in /home/ubuntu λ curl --stderr a dafdafaf
  root in /home/ubuntu λ cat a
  curl: (6) Could not resolve host: dafdafaf
  ```

- `-w | --write-out <format>`

  format可以使用`@filename`的格式从文件中读取或者`@-`从stdin中读取。使用`%{variable_name}`来设置变量

  variable_name具体可以使用值查看manual page

  ```
  root in /home/ubuntu λ curl -so /dev/null -w %{url_effective} baidu.com
  http://baidu.com/# 
  ```

## ftp/sftp

- `-a`

  以增量模式上传文件

## ipv6

https://serverfault.com/questions/1026466/ipv6-address-text-notation-with-prefix-inside-or-outside-square-brackets

- `-g | --globoff`

  关闭shell的globbing，url中可以出现`{}[]`。在shell中`[]`表示匹配指定字符，注意区别test command

按照RFC 5952，如果需要表示ipv6地址，必须使用square bracket  (因为ipv6可以使用零分十进制，ip中带有：，如果直接更port就会出错，所以需要和square bracket一起使用)

`curl -6g [ipv6]:port`

## http/https

- `-A <user_agent>`

  参数指定客户端的用户代理标头，即`User-Agent`。curl 的默认用户代理字符串是`curl/[version]`。

  ```
  $ curl -A 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36' https://google.com
  ```

  上面命令将`User-Agent`改成 Chrome 浏览器。

  ```
  $ curl -A '' https://google.com
  ```

  上面命令会移除`User-Agent`标头。

  也可以通过`-H`参数直接指定标头，更改`User-Agent`。

  ```
  $ curl -H 'User-Agent: php/1.0' https://google.com
  ```

- `--cacert <file>`

  指定使用的ca证书

- `-b <cookie | filename>`

  参数用来向服务器发送 Cookie。

  ```
  $ curl -b 'foo=bar' https://google.com
  ```

  上面命令会生成一个标头`Cookie: foo=bar`，向服务器发送一个名为`foo`、值为`bar`的 Cookie。

  ```
  $ curl -b 'foo1=bar;foo2=bar2' https://google.com
  ```

  上面命令发送两个 Cookie。如果没有在参数内容中使用`=`即表示filname

  ```
  $ curl -b cookies.txt https://www.google.com
  ```

  上面命令读取本地文件`cookies.txt`，里面是服务器设置的 Cookie（参见`-c`参数），将其发送到服务器。

- `-c|--cookie-jar`

  将服务器返回的cookie保存到本地

  ```
  $ curl -c cookies.txt https://www.google.com
  ```

  上面命令将服务器的 HTTP 回应所设置 Cookie 写入文本文件`cookies.txt`。

- `-d`

  参数用于发送 POST 请求的数据体。

  ```
  $ curl -d'login=emma＆password=123'-X POST https://google.com/login
  # 或者
  $ curl -d 'login=emma' -d 'password=123' -X POST  https://google.com/login
  ```

  使用`-d`参数以后，HTTP 请求会自动加上标头`Content-Type : application/x-www-form-urlencoded`。并且会自动将请求转为 POST 方法，因此可以省略`-X POST`。

  `-d`参数可以读取本地文本文件的数据，向服务器发送。

  ```
  $ curl -d '@data.txt' https://google.com/login
  ```

  上面命令读取`data.txt`文件的内容，作为数据体向服务器发送。

- `--data-urlencode`

  参数等同于`-d`，发送 POST 请求的数据体，区别在于会自动将发送的数据进行 URL 编码。

  ```
  $ curl --data-urlencode 'comment=hello world' https://google.com/login
  ```

  上面代码中，发送的数据`hello world`之间有一个空格，所以需要进行 URL 编码。

- `-D | --dump-header <filename>`

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

- `-k`参数指定跳过 SSL 检测。

  ```
  $ curl -k https://www.example.com
  ```

  上面命令不会检查服务器的 SSL 证书是否正确。

- `-G | --get`

  当和`-d`一起使用是，使用该参数强制发送GET请求

- `-x`

  指定 HTTP 请求的代理。

  ```
  $ curl -x socks5://james:cats@myproxy.com:8080 https://www.example.com
  ```

  上面命令指定 HTTP 请求通过`myproxy.com:8080`的 socks5 代理发出。

  如果没有指定代理协议，默认为 HTTP。

  ```
  $ curl -x james:cats@myproxy.com:8080 https://www.example.com
  ```

  上面命令中，请求的代理使用 HTTP 协议。

- `--socks5 <host:port>`

  使用socks5代理，如果没有指定port默认是1080。默认会覆盖`-x`参数

- `-X`

  指定 HTTP 请求的方法。

  ```
  $ curl -X POST https://www.example.com
  ```

  上面命令对`https://www.example.com`发出 POST 请求。也可以使用环境变量http_proxy(这个是特殊变量必须小写)和HTTPS_PROXY

- `-0`

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

- `--http1.1 | --http2`

  使用http1.1发送请求，http2需要目标libcurl支持http2.0

- `-e`

  参数用来设置 HTTP 的标头`Referer`，表示请求的来源。

  ```
  curl -e 'https://google.com?q=example' https://www.example.com
  ```

  上面命令将`Referer`标头设为`https://google.com?q=example`。

  `-H`参数可以通过直接添加标头`Referer`，达到同样效果。

  ```
  curl -H 'Referer: https://google.com?q=example' https://www.example.com
  ```

- `-H`

  添加 HTTP 请求的标头。

  ```
  $ curl -H 'Accept-Language: en-US' https://google.com
  ```

  上面命令添加 HTTP 标头`Accept-Language: en-US`。

  ```
  $ curl -H 'Accept-Language: en-US' -H 'Secret-Message: xyzzy' https://google.com
  ```

  上面命令添加两个 HTTP 标头。

  ```bash
  $ curl -d '{"login": "emma", "pass": "123"}' -H 'Content-Type: application/json' https://google.com/login
  ```

  上面命令添加 HTTP 请求的标头是`Content-Type: application/json`，然后用`-d`参数发送 JSON 数据。

- `-F`

  用来向服务器上传二进制文件。

  ```
  $ curl -F 'file=@photo.png' https://google.com/profile
  ```

  上面命令会给 HTTP 请求加上标头`Content-Type: multipart/form-data`，然后将文件`photo.png`作为`file`字段上传。

  `-F`参数可以指定 MIME 类型。

  ```
  $ curl -F 'file=@photo.png;type=image/png' https://google.com/profile
  ```

  上面命令指定 MIME 类型为`image/png`，否则 curl 会把 MIME 类型设为`application/octet-stream`。

  `-F`参数也可以指定文件名。

  ```
  $ curl -F 'file=@photo.png;filename=me.png' https://google.com/profile
  ```

  上面命令中，原始文件名为`photo.png`，但是服务器接收到的文件名为`me.png`。

- `-f | --fail`

  如果抓取指定网站的内容不显示，会默认显示一个错误页面，使用改参数内容将不显示，而是返回exit code。==但是一般的网站重定向后会展示一个非默认的错误页面，所以不显示exit code==

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

- `-i`

  参数打印出服务器回应的 HTTP 响应头。

  ```bash
  $ curl -i https://www.example.com
  ```

  上面命令收到服务器回应后，先输出服务器回应的响应头，然后空一行，再输出网页的源码。

- `-I`

  参数向服务器发出 HEAD 请求，然会将服务器返回的 HTTP 标头打印出来。

  ```
  $ curl -I https://www.example.com
  ```

  上面命令输出服务器对 HEAD 请求的回应。

  `--head`参数等同于`-I`。

  ```
  $ curl --head https://www.example.com
  ```

- `--keepalive-time <seconds>`

  发送keepalive信号的间隔，默认60sec，tcp连接保活

- `--no-keepalive`

  不使用keepalive

- `-L`

  会让 HTTP 请求跟随服务器的重定向（==http status code是3xx==）。curl 默认不跟随重定向。

  ```
  $ curl -L -d 'tweet=hi' https://api.twitter.com/tweet
  ```

- `--limit-rate`

  用来限制 HTTP 请求和回应的带宽，模拟慢网速的环境。

  ```
  $ curl --limit-rate 200k https://google.com
  ```

- `-u`

  用来设置服务器认证的用户名和密码。

  ```
  $ curl -u 'bob:12345' https://google.com/login
  ```

  上面命令设置用户名为`bob`，密码为`12345`，然后将其转为 HTTP 标头`Authorization: Basic Ym9iOjEyMzQ1`。

  curl 能够识别 URL 里面的用户名和密码。

  ```
  $ curl https://bob:12345@google.com/login
  ```

  上面命令能够识别 URL 里面的用户名和密码，将其转为上个例子里面的 HTTP 标头。

  ```
  $ curl -u 'bob' https://google.com/login
  ```

  上面命令只设置了用户名，执行后，curl 会提示用户输入密码。

- `-T | --upload <file>`

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

## Exit code

> man `/exit codes`

在使用curl时，exit code会记录在括号中，例如

```
curl: (56) Recv failure: Connection reset by peer

curl: (52) Empty reply from server
```

## 注意点

- 在linux上不支持file协议(UNC path)，在windows上可以使用

- curl不能使用stdout重定向，因为curl会把后面的当做url，可以使用`-o`参数，结合`--no-progress-meter`关闭所有回显

- 可以结合shell的模式扩展

  ```
  root in /home/ubuntu λ curl {baidu,sohu}.com
  ```

## 常用命令

- `curl -o /dev/null -svm1 >&2 <host>`

  将stderr也输出到stdout

- `curl -skvo /dev/null <host>`

  只输出请求和响应的关于http的详细信息，但不输出返回的信息

  ```
  cpl in /tmp λ curl -skvo /dev/null baidu.com
  *   Trying 220.181.38.251:80...
  * Connected to baidu.com (220.181.38.251) port 80 (#0)
  > GET / HTTP/1.1
  > Host: baidu.com
  > User-Agent: curl/7.78.0
  > Accept: */*
  > 
  * Mark bundle as not supporting multiuse
  < HTTP/1.1 200 OK
  < Date: Thu, 12 Aug 2021 04:01:58 GMT
  < Server: Apache
  < Last-Modified: Tue, 12 Jan 2010 13:48:00 GMT
  < ETag: "51-47cf7e6ee8400"
  < Accept-Ranges: bytes
  < Content-Length: 81
  < Cache-Control: max-age=86400
  < Expires: Fri, 13 Aug 2021 04:01:58 GMT
  < Connection: Keep-Alive
  < Content-Type: text/html
  < 
  { [81 bytes data]
  * Connection #0 to host baidu.com left intact
  ```

- `curl -fsSL host`

  重定向，并以exit code的方式输出错误信息，但是不输出进度信息

  ```
  cpl in /tmp λ curl -fsSL -o /dev/null baidu.com
  ```

- `curl -Sso /dev/null -w %{time_connect} 1.1.1.1 `

  统计http statuscode