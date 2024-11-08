# Linux sysctl cli

ref

https://www.kernel.org/doc/html/latest/admin-guide/sysctl/index.html

https://wiki.archlinux.org/index.php/sysctl

## Digest

syntax：`sysctl [options] [variable[=value]] [...]`
用于配置 linux runtime kernel parameters，可用的参数均在 `/proc/sys/`下

## Dir

`/proc/sys`下的目录一般包含如下几个文件，分别配置不同的系统参数

| dir     | brief                                                        |
| ------- | ------------------------------------------------------------ |
| abi/    | execution domains & personalities                            |
| debug/  | kernel debugging interfaces                                  |
| dev/    | device specific information (eg dev/cdrom/info)              |
| fs/     | specific filesystems filehandle, inode, dentry and quota tuning binfmt_misc <[Kernel Support for miscellaneous Binary Formats (binfmt_misc)](https://www.kernel.org/doc/html/latest/admin-guide/binfmt-misc.html)> |
| kernel/ | global kernel info / tuning miscellaneous stuff              |
| net/    | networking stuff, for documentation look in: <Documentation/networking/> |
| proc/   | <empty>                                                      |
| sunrpc/ | SUN Remote Procedure Call (NFS)                              |
| vm/     | memory management tuning buffer and cache management         |
| user/   | Per user per user namespace limits                           |

## Options

- `-n | --values`

打印参数是不打印参数名

- `-w | --write`

修改参数值，不会永久生效

- `-p FILE | --load FILE`

加载内核配置参数（默认读取`/etc/sysctl.conf`），==修改配置文件后需要使用该命令让配置生效==

- `-a | --all`

查看所有当前使用的内核参数

- `--system`

从所有的系统配置文件位置读取并加载内核配置，包含如下几个位置

   1. /run/sysctl.d/*.conf
   1. /etc/sysctl.d/*.conf
   1. /usr/local/lib/sysctl.d/*.conf
   1. /usr/lib/sysctl.d/*.conf
   1. /lib/sysctl.d/*.conf
   1. /etc/sysctl.conf
## Example
设置一个系统参数，并让其立即生效（只在当前启动中生效，重启后失效）
```bash
[root@chz etc]# sysctl net.ipv4.icmp_echo_ignore_all=1;sysctl -p
net.ipv4.icmp_echo_ignore_all = 1
```
使用上述方法并不会持久化内核参数，如果需要持久化需要将内容写入到`/etc/sysctl.conf`
```bash
[root@chz etc]# vim /etc/sysctl.conf
[root@chz etc]# cat /etc/sysctl.conf | grep icmp_echo
net.ipv4.icmp_echo_ignore_all = 1
[root@chz etc]# sysctl -p
```
