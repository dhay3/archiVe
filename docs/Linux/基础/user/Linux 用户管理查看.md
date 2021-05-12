# Linux 用户管理查看

## id

该命令用于查看指定用户的信息

```
root in /var/spool λ id zs
uid=1001(zs) gid=1001(zs) groups=1001(zs)
root in /var/spool λ id -Z zs #查看SELinux 规则
id: --context (-Z) works only on an SELinux-enabled kernel   
#查看用户的uid和gid
cpl in ~ λ id -g;id -u
1000
1000
```

## 查看所有用户名

`grep`使用`-v`参数表示不匹配

```
root in /var/spool λ cat /etc/passwd | grep -v nologin | grep -v sync |awk -F : '{print $1":"$3":"$4}'
root:0:0
zs:1001:1001         
```

