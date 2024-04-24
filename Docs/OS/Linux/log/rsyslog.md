## Overview

rsyslog 是 Linux 上的一个工具，用于收集日志

## Configuration

rsyslog 主要的配置文件在 `/etc/rsyslog.conf`, 用来定义 input modules, filters, actions 以及 gloabl directives

rsyslog 支持 3 种配置语法(可以互相组合使用)

1. basic

   基本的配置语法,也被称为 sysklogd 语法，例如

   ```
   mail.info /var/log/mail.log
   mail.err @@server.example.net
   ```

2. advanced

   高级配置语法，也被称为 RainerScript 语法，例如

   ```
   mail.err action(type="omfwd" protocol="tcp" queue.type="linkedList")
   ```

3. obsolete legacy

   被弃用的，生产中不应该使用

### Basic

基本配置语法[^2]

#### Selectors

选择特定的 subsystem 中指定 priority 的日志

selectors 由 2 部分组成，由 period (`.`) 分隔，大小写不敏感。具体信息在 `man 3 syslog` 中

1. facility

   表示 the subsystem that produced the message

   可以是 *auth, authpriv, cron, daemon, ftp, kern, lpr, mail, mark, news, security (same as auth), syslog, user, uucp and local0 through local7.* 其中的 security 已经被弃用了。大多数情况下，所有人都可以使用任意的 facility，但是 kern 除外，只能被 kernel 使用。

   例如 all mail program log with the mail facility if the log using syslog

2. priority

   表示 the severity of the message

   可以是 *debug, info, notice, warning, warn (same as warning), err, error (same as err), crit, alert, emerg, panic (same as emerg).* 其中 warn, error，panic 都被弃用了。wildcard `*` 表示所有的 priority

有几点特殊的需要注意

- asterisk(`*`) 表示所有的 facility 或者是所有的 priority，根据被声明的位置
- 可以使用 `none` 表示不包含所有的 priority
- 多个 facility 可以通过 comma(`,`) 组合，应用于一个 priority。但是 priority 不能通过 comma 组合
- 多个 selectors 可以 semicolon(`;`) 组合，应用于一个 action。selectors 会按照从左到右执行
- 可以在 priority 前使用 equation sign(`=`)，表示只针对单一的 priority，不包括 higher priorities
- 可以在 priority 前使用 exclamation mark(`!`)，表示忽略当前的 priority 以及 higher priorities
- 在 priority 前还可以组合 exclamtion mark 和 equation sign，表示只忽略单一的 prioritys

#### Actions

通常是一个文件，表示对应的 facility.priority 日志最后被存储到那里，可以是

1. regular file

   一个普通的文件，可以使用绝对或者是相对路径

2. named pipes

   由 `mkfifo` 创建的具名管道符

3. terminal and console

   可以是 tty 文件

4. remote machine

   可以是一台远程运行 syslogd 的服务器。需要在 hostname 或者是 IP address 前使用 `@` 表示是一台远程服务器，例如 `@192.168.2.1`

### Basic Examples

#### Example one

```
*.=crit;kern.none   /var/adm/critical
```

上述的配置会将所有的 crit messages (`*.=crit`)记录到 /var/adm/critical 中，但是除所有的 kern 日志外(`;kern.none`)

#### Example two

```
kern.*      /var/adm/kernel
kern.crit     @finlandia
kern.crit     /dev/console
kern.info;kern.!err   /var/adm/kernel-info
```

1. `kern.* /var/adm/kernel`

   所有的 kern 日志都会被记录到 `/var/adm/kernel`

2. `kern.crit @finlandia`

   所有 kern 级别大于等于 crit 的日志，都会被发送到 remote host finlandia

3. `kern.crit /dev/console`

   所有 kern 级别大于等于 crit 的日志，都会被发送到 terminal console。如果有其他的用户登入到同台服务器的 ternimal 就会看到 crit 日志

4. `kern.info;kern.!err /var/adm/kernel-info`

   所有 kern 级别大于 info 小于 err 的日志，都会被记录到 `/var/adm/kernel-info`

## Conversion

大多数配置都可以被 basic 语法描述的详细得当，所以转为 advanced 语法并不是必须的。例如

```
mail.info   /var/log/maillog
```

可以转为 advanced 语法，等价于下面的配置，但并不是必须的

```
if prifilt("mail.info") then {
     action(type="omfile" file="/var/log/maillog")
}
#or
if prifilt("mail.info") then action(type="omfile" file="/var/log/maillog")
#or
mail.info action(type="omfile" file="/var/log/maillog")
```

#### converting Module Load

如果要转换 obsolete legacy 语法中的 `$ModLod` 为 advanced，参考如下

```
#obsolete legacy
$ModLoad module-name
#advanced
module(load="module-name")
```

如果想要声明特定的 Module 使用参数，在 obsolete legacy 语法中需要声明在 `$ModLoad` 后面。而 advanced 语法中更加向调用函数来声明参数

```
#obsolete legacy
$ModLoad imtcp
$InputTCPMaxSession 500
#advanced
module(load="imtcp" maxSessions="500")
```

需要注意的一点是在 obsolete legacy 语法中，参数可以被多次调用。但是 advanced 语法中不支持（逻辑上也应该被禁止，会导致配置文件难以阅读）

```
#obsolete legacy
$ModLoad imtcp
$InputTCPMaxSession 500
...
*.* /var/log/messages
...
$InputTCPMaxSession 200

#advanced
module(load="imtcp" maxSessions="200")
...
*.* /var/log/messages
...
```

#### converting Actions

| basic                          | advanced                                                     |
| ------------------------------ | ------------------------------------------------------------ |
| file path (/var/log/…)         | action(type=”[omfile](https://www.rsyslog.com/doc/configuration/modules/omfile.html)” file=”/var/log…/” …) |
| UDP forwarding (@remote)       | action(type=”[omfwd](https://www.rsyslog.com/doc/configuration/modules/omfwd.html)” target=”remote” protocol=”udp” …) |
| TCP forwarding (@@remote)      | action(type=”[omfwd](https://www.rsyslog.com/doc/configuration/modules/omfwd.html)” target=”remote” protocol=”tcp” …) |
| user notify (`:omusrmsg:user`) | action(type=”[omusrmsg](https://www.rsyslog.com/doc/configuration/modules/omusrmsg.html)” users=”user” …) |
| module name (`:omxxx:..`)      | action(type=”[omxxx](https://www.rsyslog.com/doc/configuration/modules/idx_output.html)” …) |

#### converting Action Chains

**references**

[^1]:https://www.rsyslog.com/doc/configuration/index.html
[^2]:https://www.rsyslog.com/doc/configuration/sysklogd_format.html