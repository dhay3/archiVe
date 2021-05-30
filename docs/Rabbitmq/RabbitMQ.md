# RabbitMQ

### 安装RabbitMQ



<img src="..\..\..\..\java资料\rbmq\img\1.png"/>

下载erlang和rabbitmq

进入rabbitmq安装目录中的sbin,或是通过shortcut 进入cmd窗口

```
rabbitmq-plugins enable rabbitmq_management
```

rabbitMQ自动开启,默认端口5672

访问`http://localhost:15672/`进入rabbitmq的图形化管理界面

默认账号guest,默认密码 geust

### 添加用户

<img src="..\..\..\..\java资料\rbmq\img\2.PNG"/>

### 添加虚拟host

### <img src="D:\java资料\rbmq\img\3.png" alt="3" style="zoom:60%;" />host一般以斜杠开头

### 用户授权

进入新创建好的virtual host

<img src="D:\java资料\rbmq\img\4.PNG" alt="4" style="zoom:60%;" />

不需要重新开启, rabbitMQ自动开启
