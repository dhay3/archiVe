# Zabbix media types

> 注意zabbix 发送邮件还需和action一起使用

参考：

https://www.osyunwei.com/archives/8113.html

https://blog.51cto.com/andyxu/2145196

https://www.cnblogs.com/kevingrace/p/7107408.html

## 概述

*a means of delivering notifications; delivery channel.*

media types 报警媒介。如果zabbix 某项item触发了trigger该怎么告知运维呢？就是通过media types。

zabbix支持的报警媒介有：

- Email：邮件，这是最常用也是最传统的一种报警媒介，邮件报警，zabbix通过配置好的SMTP邮件服务器向用户发送对应的报警信息。
- Script：脚本，当zabbix中的某些监控项出现异常时，也可以调用自定义的脚本进行报警，脚本的使用就比较灵活，具体怎样报警全看你的脚本怎么写。
- SMS：短信，如果想要使用短信报警，则需要依赖短信网关（貌似需要北美的运行商）。
- Jabber：即时通讯服务。
- Ez Texting：商业的，收费的短信服务（北美运营商提供服务）。

**邮件报警有两种情况：**

1. Zabbix服务端只是单纯的发送报警邮件到指定邮箱，发送报警邮件的这个邮箱账号是Zabbix服务端的本地邮箱账号（例如：root@localhost.localdomain），只能发送，不能接收外部邮件。

2. 使用一个可以在互联网上正常收发邮件的邮箱账号（例如：xxx@163.com），通过在Zabbix服务端中设置，使其能够发送报警邮件到指定邮箱。

> 如果发送成功，我们可以在report → action log中查看

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-11_20-12-53.png"/>

## Email创建

administration → media types → create media type

参考：https://www.zabbix.com/documentation/5.0/manual/config/notifications/media/email

### Media type

> 使用Email时无需修改`/etc/mail.rc`，添加完成后，可以点击list右边的test来测试媒介

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-08_15-33-23.png"/>

- SMTP server：指定smtp服务器，`smtp.163.com`。我们可以在网易邮箱设置中pops/smtp/imap最下方找到

- SMTP server port：默认端口25

- SMTP helo：smtp服务器地址后缀`163.com`CBLOOWSQHRTYTRXQ

- SMTP email：该邮箱地址被当成zabbix发送邮件的地址，指定你的邮箱

- connection security：连接的安全性

- authentication：选择认证方式

  > **注意这里的密码是==smtp的授权认证码==**
  
- message format：消息主体的格式

### Message templates

设置默认的消息主体格式

参考：https://www.zabbix.com/documentation/5.0/manual/config/notifications/media#common_parameters

> ==需要当action触发时才能传递信息==

### add user media type

配置用户的media type，即发送邮件方

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-11_16-36-07.png"/>

### Add action

具体参考action创建

### check mailbox

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-11_20-10-26.png"/>

## 无法发送邮件解决方案

参考：https://yq.aliyun.com/articles/503964

## script创建

administration → media types → create media type

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-07_12-00-19.png"/>

通过脚本报警是非常灵活的，我们可以通过脚本实现发送邮件，发送短信

- 配置`/etc/mail.rc`文件

  ```
  set from=xxx@163.com  # 发送的邮件地址
  set smtp=smtp.163.com  # 发送邮件服务器
  set smtp-auth-user=xxx@163.com  # 发件人账号
  set smtp-auth-password=xxx  # 发件人密码（smtp授权码）
  set smtp-auth=login  # 邮件认证方式
  ```

- 编写脚本

  ```
  [root@chz opt]# cat mail.sh 
  #!/bin/bash
  to=$1
  subject=$2
  context=$3
  
  echo -e "$context" | mail -s "$subject" "$to"
  ```

- 测试脚本

  ```
  [root@chz opt]# sh mail.sh hostlockdown@gmail.com sh "mail.sh test"
  ```

  我们可以登入gmail来校验

- 将编写的脚本放在默认的报警脚本目录，通过`/etc/zabbix/zabbix_server.conf`文件查看，并修改文件权限

  ```i
  AlertScriptsPath=/usr/lib/zabbix/alertscripts
  ...
  cd /usr/lib/zabbix/alertscripts
  cp /opt/mail.sh .
  chmod 551 mail.shtt
  ```

- 添加脚本

  macros参考：https://www.zabbix.com/documentation/5.0/manual/appendix/macros/supported_by_location

  Parameter是zabbix built-in 参数称为macros

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-08_14-58-55.png"/>

- 添加脚本后还不具有发送邮件的功能，还需让用户具有使用media type的能力

  administration→users→media

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-08_15-05-42.png"/>

  








