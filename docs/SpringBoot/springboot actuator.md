# springboot actuator

### #依赖

```xml
		<dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>
```

### #yml

```yml
  server:
    port: 80 #指定端口 ,缺省值项目端口 ,访问可以不用写出80端口,这是公认值
    servlet:
      context-path: /monitor #指定访问路径, 缺省值 /actuator
  endpoints:
    web:
      exposure:
        include: "*" #web 默认只能访问health 和info ,其他endpoints不对web开放
#    enabled-by-default: false 关闭默认的endpoints配置
#  endpoint: #使用自己设置的endpoints配置
#    health:
#      enabled: true
#    info:
#      enabled: true
```

直接通过

如果没有配置`context-path` 直接通过`localhost:${server.port}/actuator/{endpoints}`访问

如果配置了`context-path` , 像上面的例子就应该访问

`localhost/monitor/actuator/${endpoints}`

这是`spring`官网的原话

```
When a custom management context path is configured, the “discovery page” automatically moves from /actuator to the root of the management context. 
```

### #常见的endpoints

| ID      | Description                                          |
| ------- | ---------------------------------------------------- |
| /beans  | 显示你项目中的所有的bean                             |
| /caches | 显示使用的缓存                                       |
| /env    | 显示所有环境属性                                     |
| /health | 显示应用的健康指标                                   |
| /info   | 显示应用程序的定制信息，这些信息由info打头的属性提供 |

