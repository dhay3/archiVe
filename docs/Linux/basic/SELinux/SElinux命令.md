# SElinux命令



## getenforce、sestatus、setforce

- getenforce

  显示当前工作模式

  ```
  [root@chz opt]# getenforce 
  Enforcing
  ```

- sestatus

  显示当前工作模式和策略

  ```
  [root@chz /]# sestatus 
  SELinux status:                 enabled
  SELinuxfs mount:                /sys/fs/selinux
  SELinux root directory:         /etc/selinux
  Loaded policy name:             targeted
  Current mode:                   enforcing
  Mode from config file:          enforcing
  Policy MLS status:              enabled
  Policy deny_unknown status:     allowed
  Max kernel policy version:      31
  ```

- setenforce

  设置工作模式

  0：宽容模式

  1：强制模式

  如果想要置于disable需要修改配置文件

  ```
  [root@chz /]# setenforce 0
  [root@chz /]# getenforce 
  Permissive
  ```

## seinfo

```
#安装
[root@localhost ~]# yum -y install setroubleshoot
[root@localhost ~]# yum -y install setools-console
[root@chz html]# seinfo

#列出当前SELinux中所有属性
Statistics for policy file: /sys/fs/selinux/policy
Policy Version & Type: v.31 (binary, mls)

   Classes:           130    Permissions:       272
   Sensitivities:       1    Categories:       1024
   Types:            4792    Attributes:        253
   Users:               8    Roles:              14
   Booleans:          316    Cond. Expr.:       362
   Allow:          107360    Neverallow:          0
   Auditallow:        157    Dontaudit:       10020
   Type_trans:      18129    Type_change:        74
   Type_member:        35    Role allow:         39
   Role_trans:        416    Range_trans:      5899
   Constraints:       143    Validatetrans:       0
   Initial SIDs:       27    Fs_use:             32
   Genfscon:          103    Portcon:           614
   Netifcon:            0    Nodecon:             0
   Permissives:         0    Polcap:              5

```

### 参数

- -u 

  列出所有user

  ```
  [root@chz yum]# seinfo -u
  
  Users: 8
     sysadm_u
     system_u
     xguest_u
     root
     guest_u
     staff_u
     user_u
     unconfined_u
  ```

- -r

  列出所有role

  ```
  [root@chz yum]# seinfo -r
  
  Roles: 14
     auditadm_r
     dbadm_r
     guest_r
     staff_r
     user_r
     logadm_r
     object_r
     secadm_r
     sysadm_r
     system_r
     webadm_r
     xguest_r
     nx_server_r
     unconfined_r
  ```

- -t

  列出所有type

  ```
  [root@chz yum]# seinfo -t|more
  
  Types: 4792
     bluetooth_conf_t
     cmirrord_exec_t
     colord_exec_t
     container_auth_t
     foghorn_exec_t
     jacorb_port_t
     pki_ra_exec_t
     pki_ra_lock_t
     sosreport_t
     squid_script_exec_t
     etc_runtime_t
     fenced_tmp_t
     git_session_t
     glance_port_t
     osad_log_t
  ```

- -b

  默认显示默认策略(targeted)中的布尔值即规则

  ```
  [root@cyberpelican www]# seinfo -b|more
  
  Conditional Booleans: 316
     auditadm_exec_content
     cdrecord_read_content
     cvs_read_shadow
     fcron_crond
     glance_api_can_network
     gluster_export_all_rw
     httpd_dontaudit_search_dirs
     httpd_manage_ipa
     httpd_run_ipa
     httpd_run_stickshift
  ```

## sesearch

seinfo只能看到所有==规则的名称==，如果想要知道规则的具体内容，就需要使用sesearch

- --allow

  显示允许的规则

  ```
  [root@cyberpelican www]# sesearch --allow -t httpd_t|more
  Found 991 semantic av rules:
     allow sosreport_t domain : packet_socket getattr ; 
     allow mojomojo_script_t httpd_t : unix_stream_socket { ioctl read write getattr ac
  cept } ; 
  ```

- --neverallow

  显示不允许的规则

- --all

  显示所有规则

- -s

  显示和指定==主体的类型（subject type）==的规则

  ```
  [root@cyberpelican www]# sesearch --allow -s httpd_t|more
  Found 1916 semantic av rules:
     allow httpd_t init_var_run_t : dir { getattr search open } ; 
     allow daemon init_t : pppox_socket { ioctl read write getattr getopt
   setopt } ; 
  ```

- -t

  显示和指定==目标的类型（object type）==的规则

  ```
  [root@cyberpelican www]# sesearch --allow -t httpd_t|more
  Found 991 semantic av rules:
     allow sosreport_t domain : packet_socket getattr ; 
     allow mojomojo_script_t httpd_t : unix_stream_socket { ioctl read wr
  ite getattr accept } ; 
  ```

==查看subject的domain与object的type是否匹配==

我们知道`httpd_t`与`httpd_sys_content_t`是一对匹配的规则

```
[root@cyberpelican www]# sesearch --allow -s httpd_t -t httpd_sys_content_t|more
Found 29 semantic av rules:
   allow httpd_t httpd_sys_content_t : lnk_file { read getattr } ; 
   allow httpd_t httpd_sys_content_t : dir { ioctl read getattr lock se
arch open } ; 
   allow httpd_t httpd_sys_content_t : file { ioctl read getattr lock m
ap open } ; 

```

可以清楚看到httpd domain允许访问和使用httpd_sys_content_t

## getsebool/setsebool

查看当前Policy中的规则是启用还是关闭

### getsebool

getsebool只有一个用法，显示当前策略的规则开启与关闭的状态

```
[root@chz ~]# getsebool -a | more
abrt_anon_write --> off
abrt_handle_event --> off
abrt_upload_watch_anon_write --> on
antivirus_can_scan_system --> off
antivirus_use_jit --> off
auditadm_exec_content --> on
authlogin_nsswitch_use_ldap --> off
```

### setsebool

设置规则的状态，1表示开启，0表示关闭

```
[root@chz ~]# getsebool -a | grep httpd_tty_comm
httpd_tty_comm --> off
[root@chz ~]# setsebool -PNV httpd_tty_comm=1
[root@chz ~]# getsebool -a | grep httpd_tty_comm
httpd_tty_comm --> on
```



## chcon

修改文件资源的安全性文本

### 参数

- -u

  修改user

  ```
  [root@chz html]# chcon -u root index.html 
  [root@chz html]# ls -Z
  -rw-r--r--. root root root:object_r:var_lib_t:s0       index.html
  ```

- -r

  修改role

  ```
  [root@chz html]# chcon -r object_r index.html 
  ```

- -t

  修改type

  ```
  [root@chz html]# ls -Z 
  -rw-r--r--. root root unconfined_u:object_r:httpd_sys_content_t:s0 index.html
  [root@chz html]# ps -efZ|grep httpd
  system_u:system_r:httpd_t:s0    root       3584      1  0 08:16 ?        00:00:00 /usr/sbin/httpd -DFOREGROUND
  system_u:system_r:httpd_t:s0    apache     3614   3584  0 08:16 ?        00:00:00 /usr/sbin/httpd -DFOREGROUND
  system_u:system_r:httpd_t:s0    apache     3615   3584  0 08:16 ?        00:00:00 /usr/sbin/httpd -DFOREGROUND
  system_u:system_r:httpd_t:s0    apache     3616   3584  0 08:16 ?        00:00:00 /usr/sbin/httpd -DFOREGROUND
  system_u:system_r:httpd_t:s0    apache     3617   3584  0 08:16 ?        00:00:00 /usr/sbin/httpd -DFOREGROUND
  system_u:system_r:httpd_t:s0    apache     3618   3584  0 08:16 ?        00:00:00 /usr/sbin/httpd -DFOREGROUND
  system_u:system_r:httpd_t:s0    apache     3619   3584  0 08:16 ?        00:00:00 /usr/sbin/httpd -DFOREGROUND
  system_u:system_r:httpd_t:s0    apache     3969   3584  0 08:20 ?        00:00:00 /usr/sbin/httpd -DFOREGROUND
  system_u:system_r:httpd_t:s0    apache     3976   3584  0 08:20 ?        00:00:00 /usr/sbin/httpd -DFOREGROUND
  system_u:system_r:httpd_t:s0    apache     3977   3584  0 08:20 ?        00:00:00 /usr/sbin/httpd -DFOREGROUND
  unconfined_u:unconfined_r:unconfined_t:s0-s0:c0.c1023 root 9204 2975  0 11:05 pts/0 00:00:00 grep --color=auto httpd
  [root@chz html]# chcon -t var_t index.html 
  [root@chz html]# ls -Z
  -rw-r--r--. root root unconfined_u:object_r:var_lib_t:s0 index.html
  ```

  由于security context 不匹配，这时httpd process不能访问该文件

  > 这里匹配指并不是字面上的相同，而是policy中配置的映射

- -R

  递归

## restorcon

==将安全性文本恢复成默认的，但是不会对修改的user或是role生效==

```
[root@chz html]# restorecon -v index.html 
restorecon reset /var/www/html/index.html context root:object_r:var_lib_t:s0->root:object_r:httpd_sys_content_t:s0
```

- -R 

  递归

- -v

  verbose

## semanage

前面讲到，restorecon 命令可以将文件或目录恢复成默认的安全上下文，这就说明==每个文件和目录都有自己的默认安全上下文==，事实也是如此，==为了管理的便捷，系统给所有的系统默认文件和目录都定义了默认的安全上下文==

那么，默认安全上下文该如何查询和修改呢？这就要使用 semanage 命令了。该命令的基本格式如下：

```
[root@localhost ~]# semanage [login|user|port|interface|fcontext|translation] -l
[root@localhost ~]# semanage fcontext [选项] [-first] file_spec
```

其中，fcontext 主要用于安全上下文方面，-l 是查询的意思。除此之外，此命令常用的一些选项及含义，如表 1 所示。

| 选项 | 含义                       |
| ---- | -------------------------- |
| -a   | 添加默认安全上下文配置。   |
| -d   | 删除指定的默认安全上下文。 |
| -m   | 修改指定的默认安全上下文。 |
| -t   | 设定默认安全上下文的类型   |

- 【例 1】查询默认安全上下文。

  ```shell
  [root@localhost ~]# semanage fcontext -l
  ...
  /var/www(/.*)?                                     all files          system_u:object_r:httpd_sys_content_t:s0
  ...
  ```

  这里可以看到`/var/www(/.*)?  `目录下文件安全上下文为`system_u:object_r:httpd_sys_content_t:s0`。==注意采用表达式最优原则，如果没有指定，就采用父级目录的安全性文本。==

  所以，一旦对 `/var/www/ `目录下文件的安全上下文进行了修改，就可以使用` restorecon `命令进行恢复，因为默认安全上下文已经明确定义了。

- 【例 2】修改默认安全上下文。
  那么，可以修改目录的默认安全上下文吗？当然可以，举个例子：

  ```
  [root@localhost ~]# mkdir /www
  \#新建/www/目录，打算用这个目录作为apache的网页主目录，而不再使用/var/www/html/目录
  [root@localhost ~]# ls -Zd /www/
  drwxr-xr-x．root root unconfined_u：object_r：default_t：s0 /www/
  \#而这个目录的安全上下文类型是default_t，那么apache进程当然就不能访问和使用/www/目录了
  ```

  这时我们可以直接设置 /www/ 目录的安全上下文类型为 httpd_sys_content_t，但是为了以后管理方便，我打算修改 /www/ 目录的默认安全上下文类型。先查询一下 /www/ 目录的默认安全上下文类型，命令如下：

  可以使用`egrep`替代

  ```
  [root@localhost ~]# semanage fcontext -l | grep "/www"
  \#查询/www/目录的默认安全上下文
  ```

  查询出了一堆结果，但是并没有 /www/ 目录的默认安全上下文，因==为这个目录是手工建立的，并不是系统默认目录，所以并没有默认安全上下文，需要我们手工设定。==命令如下：

  ```
  [root@localhost ~]# semanage fcontext -a -t httpd_sys_content_t "/www(/.*)?"
  \#这条命令会给/www/目录及目录下的所有内容设定默认安全上下文类型是httpd_sys_content_t
  [root@localhost ~# semanage fcontext -l | grep "/www"
  …省略部分输出…
  /www(/.*)? all files system_u：object_r：httpd_sys_content_t：s0
  \#/www/目录的默认安全上下文出现了
  ```

  这时已经设定好了 /www/ 目录的默认安全上下文。

  ```
  [root@localhost ~]# ls -Zd /www/
  drwxr-xr-x．root root unconfined_u：object_r：default_t：s0 /www/
  \#但是查询发现/www/目录的安全上下文并没有进行修改，那是因为我们只修改了默认安全上下文，而没有修改目录的当前安全上下文
  [root@localhost ~]# restorecon -Rv /www/
  restorecon reset /www context
  unconfined_u：object_r：default_t：s0->unconfined_u：object_r：httpd_sys_content_t：s0
  \#恢复一下/www/目录的默认安全上下文，发现类型已经被修改为httpd_sys_content_t
  ```

  我们可以访问`file:///www/`来测试，默认安全上下文的设定就这么简单。

