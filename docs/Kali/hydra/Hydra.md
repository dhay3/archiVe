# Hydra

> 可以使用带有GUI的xhydra

## 概述

hydra是一个多线程支持多种协议的密码破解工具

pattern：`hydra [option] target service`

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

- `-t task`

  并行运行任务

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

- 

