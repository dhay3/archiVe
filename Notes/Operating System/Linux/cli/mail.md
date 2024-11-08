# Linux mail

参考：

https://juejin.im/post/6844904018733432840

https://blog.csdn.net/liang19890820/article/details/53115334

https://www.cnblogs.com/hanganglin/p/6510216.html

## 概述

> 发送邮件需要安装sendmail组件
>
> `yum install sendmail && systemctl start sendmail`
>
> 邮件存储在`/var/spool/email`下

`mail` 命令是 Linux 终端发送邮件用的最多的命令。`mailx` 是 `mail` 命令的更新版本，基于 Berkeley Mail 8.1，意在提供 POSIX `mailx` 命令的功能，并支持 MIME、IMAP、POP3、SMTP 和 S/MIME 扩展。mailx 在某些交互特性上更加强大，如缓冲邮件消息、垃圾邮件评分和过滤等。在 Linux 发行版上，`mail` 命令是 `mailx` 命令的软链接。可以运行下面的命令从官方发行版仓库安装 `mail` 命令。

邮件发送失败后将邮件和错误信息回传

```
From MAILER-DAEMON@chz  Sat Nov  7 13:56:55 2020
Return-Path: <MAILER-DAEMON@chz>
Received: from localhost (localhost)
        by chz (8.14.7/8.14.7) id 0A75utnc014403;
        Sat, 7 Nov 2020 13:56:55 +0800
Date: Sat, 7 Nov 2020 13:56:55 +0800
From: Mail Delivery Subsystem <MAILER-DAEMON@chz>
Message-Id: <202011070556.0A75utnc014403@chz>
To: <root@chz>
MIME-Version: 1.0
Content-Type: multipart/report; report-type=delivery-status;
        boundary="0A75utnc014403.1604728615/chz"
Subject: Returned mail: see transcript for details
Auto-Submitted: auto-generated (failure)

This is a MIME-encapsulated message

--0A75utnc014403.1604728615/chz

The original message was received at Sat, 7 Nov 2020 13:56:55 +0800
from localhost [127.0.0.1]

   ----- The following addresses had permanent fatal errors -----
<root@192.168.80.100>
    (reason: 550 Host unknown)

```

## 参数

pattern：`mailx [option] <to-addr>` 

- -s

  指定邮件内容，注意如果内容中有空格需要使用引号

  ```
  [root@chz ~]# mail -s test hostlockdown@gmail.com
  test from linux
  EOT
  ```

  我们能在gmail的垃圾邮箱中看见发送的邮件，默认以当前用户`root@chz.localdomain`做为发件者

- -u

  读取指定用户的邮件，默认读取当前登入的用户

  ```
  [root@cyberpelican mail]# mail -u root
  Heirloom Mail version 12.5 7/5/10.  Type ? for help.
  "/var/mail/root": 1 message 1 new
  >N  1 Mail Delivery System  Sun Nov  8 10:24  72/2512  "Undelivered Mail Returned to Sender"
  & ?
  ```

- -a

  添加附件

  ```
  [root@cyberpelican mail]# mail -s subject -a /opt/file.txt  hostloackdown@gmail.com
  ```

## 案例

> pattern：`mail -s <subject> <toaddr>`

1. 以当前shell做为编辑器，发送邮件，编辑完后换行ctrl+d结束

   ```
   mail -s mail1 kikochz@163.com
   this is a test
   ctrl+d
   ```

2. 以文件标准输入流，做为内容

   ```
   mail -s  mail2 kikochz@163.com < file
   ```

## mail.rc

sendmail默认会以当前登入用户做为发件人(`root@localhost.localdomain`)，==通常会被当作垃圾邮件==。`/etc/mail.rc`用于配置发送端的服务器。添加如下配置，==授权码需要通过短信获取==

```
set from=xxx@163.com  # 发送的邮件地址
set smtp=smtp.163.com  # 发送邮件服务器
set smtp-auth-user=xxx@163.com  # 发件人账号
set smtp-auth-password=xxx  # 发件人密码（smtp授权码）
set smtp-auth=login  # 邮件认证方式
```

这样linux就可以通过指定的邮件地址发送邮件

```
[root@cyberpelican mail]# echo "this is my test mail" | mail -s 'mail test' hostlockdown@gmail.com
```



