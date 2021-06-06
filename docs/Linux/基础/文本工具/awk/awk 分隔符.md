# awk 分隔符

awk中两种分隔符FS(输入分隔符)，OFS(输出分隔符)

## FS

默认以空格为分隔符，可以通过`-F`参数指定分隔符。==可以使用正则表达式==

```
[root@chz t]# cat t
/logs/tc0001/tomcat/tomcat7.1/conf/catalina.properties:app.env.server.name = demo.example.com
/logs/tc0001/tomcat/tomcat7.2/conf/catalina.properties:app.env.server.name = quest.example.com
/logs/tc0001/tomcat/tomcat7.5/conf/catalina.properties:app.env.server.name = www.example.com
[root@chz t]# cat t | awk -F'[/=]' '{print $NF}'
 demo.example.com
 quest.example.com
 www.example.com
```

这里的`[]`同正则中的相同，如果没有逗号字符串将拼接

```
[root@chz t]# cat /etc/passwd | awk -F : -v 'OFS=#' -v 'ORS=\t' '{print $1 $2}'
rootx	binx	daemonx	admx	lpx	syncx	shutdownx
```

还可以通过BEGIN pattern来指定，同理OFS

```
[root@chz opt]# cat /etc/passwd | awk 'BEGIN{FS=":"}{print $1}'
root
bin
daemon
adm
```

## OFS

默认以空格为输出分隔符，可以通过`-v OFS=`来指定。==可以使用正则表达式==

```
[root@chz t]# cat /etc/passwd | awk -F : -v OFS='#' '{print $1,$2}'
root#x
bin#x
```

