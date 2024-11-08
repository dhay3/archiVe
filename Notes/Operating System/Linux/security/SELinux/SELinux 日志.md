# SELinux 日志

参考：http://c.biancheng.net/view/1155.html

## 概述

当查看特定安全上下文的策略规则时，SELinux 会使用被称为 AVC（Access Vector Cache，访问矢量缓存）的缓存，如果访问被拒绝（也被称为 AVC 拒绝），则会在一个日志文件中记录下拒绝消息。

这些被拒绝的消息可以帮助诊断和解决常规的 SELinux 策略违规行为，至于这些拒绝消息到底被记录在什么位置，则取决于 auditd 和 rsyslogd 守护进程的状态：

- 若 auditd 守护进程正在运行，则拒绝消息将被记录在 `/var/log/audit/audit.log `中，该文件信息非常多，如果手工查看效率非常低下。
- 若 auditd 守护进程没有运行，但 rsyslogd 守护进程正在运行，则拒绝消息会记录到 `/var/log/messages` 中。

启动auditd

```
systemctl start auditd
```

## audit2why

audit2why （等价于audit2allow -w < filename）命令用来分析 audit.log 日志文件，并分析 SELinux 为什么会拒绝进程的访问。也就是说，这个命令显示的都是 SELinux 的拒绝访问信息，而正确的信息会被忽略。命令的格式也非常简单，如下：

```
[root@localhost ~]# audit2why < /var/log/audit/audit.log
type=AVC msg=audit(1370412789.400:858): avc: denied { getattr ) for pid=25624 comm="httpd"  path="/var/www/htirl/index.html"    dev=sda3    ino=918426
scontext=unconfined_u:system_r:httpd_t:s0 tcontext=unconfined_u:object_r:var_t:s0 tclass=file
#这条信息的意思是拒绝7 PID 是 25624的进程访间"/var/uww/html/Index.html",原因是主体的安全上下文和目标的安全上下文不匹配。其中，denied代表拒绝，path指定目标的文件名,scontext代表全体的安全上下文。tcontext代表目标的安全上下文，仔细看看，其实就是主体的安全上下文类型httpd_t和目标的安全上下文类型var_t不匹配导致的
Was caused by:
Missing type enforcement (TE) allow rule.
You can use audit2allow to generate a loadable module to allow this access.
#给你的处理建议是使用audi t2allow命令来再次分析这个曰志文件
```

## audit2allow

audit2allow 命令的作用是分析日志，并提供允许的建议规则或拒绝的建议规则。这么说很难理解，我们还是尝试一下吧，命令如下：

```
[root@chz audit]# audit2allow -a audit.log 


#============= httpd_t ==============

#!!!! This avc is allowed in the current policy
allow httpd_t zabbix_port_t:tcp_socket name_connect;

#============= plymouthd_t ==============
allow plymouthd_t framebuf_device_t:chr_file map;

#============= xdm_t ==============

#!!!! This avc can be allowed using the boolean 'polyinstantiation_enabled'
allow xdm_t admin_home_t:dir create;
allow xdm_t modemmanager_t:dbus send_msg;

#============= zabbix_agent_t ==============
allow zabbix_agent_t devlog_t:sock_file getattr;
allow zabbix_agent_t initctl_t:fifo_file getattr;
allow zabbix_agent_t proc_kcore_t:file getattr;
allow zabbix_agent_t rpm_exec_t:file { execute execute_no_trans };
allow zabbix_agent_t rpm_var_lib_t:file open;

#============= zabbix_t ==============

#!!!! The file '/run/zabbix/zabbix_server_preprocessing.sock' is mislabeled on your system.  
#!!!! Fix with $ restorecon -R -v /run/zabbix/zabbix_server_preprocessing.sock
#!!!! This avc can be allowed using the boolean 'daemons_enable_cluster_mode'
allow zabbix_t self:unix_stream_socket connectto;
allow zabbix_t zabbix_var_run_t:sock_file { create unlink };

```

案例一：

https://www.zabbix.com/forum/zabbix-help/367261-selinux-and-zabbix

生成错误解决方法模块

```
grep zabbix_t /var/log/audit/audit.log | audit2allow -M zabbix_server_custom

semodule -i zabbix_server_custom.pp

 rm zabbix_server_custom.*
```





