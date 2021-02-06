# Linux readlink

参考：

https://blog.csdn.net/diabloneo/article/details/7173438

readlink用于找出链接指向的对象

```
root in /var/spool/cron/crontabs λ readlink $(readlink /usr/bin/awk)
/usr/bin/gawk 
```

使用`-f`参数，递归找出真正调用的对象

```
root in /var/spool/cron/crontabs λ readlink -f /usr/bin/awk
/usr/bin/gawk   
```

