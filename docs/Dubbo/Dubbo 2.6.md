# Dubbo 2.6 安装 dubbo monitor

dubbo 2.6 安装dubbo admin 步骤差不多相同

### 2.6 与 2.7 的不同

2.6的端口是7001, 而 2.7 的是8080

<img src="..\..\imgs\_Dubbo\11.PNG"/>

2.6需要配置dubbo monitor , 而2.7 不用

2.6不需要安装可视化ui

## 安装 dubbo monitor

##### 1) 打包 dubbo monitor

```
mvn package
```

##### 2)修改配置

解压生成的dubbo-monitor-simple-xxx 压缩包

<img src="..\..\imgs\_Dubbo\12.PNG"/>

修改conf中的dubbo.properties

<img src="..\..\imgs\_Dubbo\13.PNG"/>

==dubbo.protocol.port = 7070 监控中心地址==

##### 3) 点击start.bat 

访问localhost:8080 
