# Springcloud Config

## 配置文件

```yaml
server:
  port: 3344
spring:
  application:
    name: cloud-config-center
  cloud:
    config:
      server:
        git:
          uri:  https://github.com/zzyybs/springcloud-config.git
          search-paths: #仓库名
            - springcloud-config
      label: master #分枝
eureka:
  client:
    service-url:
      defaultZone:  http://localhost:7001/eureka
```

---

也可以采用如下形式

```yaml
server:
  port: 3355
spring:
  application:
    name: config-client
  cloud:
    config:
      uri: http://localhost:3344
      label: master #分支
      name: config #配置文件名
      profile: dev #读取后缀名称, 会自动添加 - , 相当于访问http://localhost:334/master/config-dev.yml
eureka:
  client:
    service-url:
      defaultZone: http://localho=st:7001/eureka
```

具体参考https://docs.spring.io/spring-cloud-config/docs/2.2.5.RELEASE/reference/html/#_quick_start

pattern `/{label}/{application}-{profile}.yml`，这里的label是github的分支。

测试访问http://localhost:3344/master/config-dev.yml

## 主启动类

```java
@EnableConfigServer//开启config服务
@SpringBootApplication
public class Config3344Application {
    public static void main(String[] args) {
        SpringApplication.run(Config3344Application.class,args);
    }
}
```

## 测试

添加测试controller

```
@RestController
public class ConfigClientController {
    @Value("${config.info}")
    private String configInfo;

    @GetMapping("/configInfo")
    public String getConfigInfo() {
        return configInfo;
    }
}
```

需要先启动springcloud config 和 eureka 微服务，springcloud config就会像仓库拿取配置文件

> 注意只有config server 会对配置文件修改生效，其他模块需要重启，或是暴露端口才能生效
>
> 但是不能广播，这里需要引入springcloud bus

1. 引入依赖

   ```xml
          <dependency>
               <groupId>org.springframework.boot</groupId>
               <artifactId>spring-boot-starter-actuator</artifactId>
           </dependency>
   ```

2. 暴露所有端口

   ```yaml
   management:
     endpoints:
       web:
         exposure:
           include: "*"
   ```

3. 业务类添加`@ResfreshScop`

   ```java
   @RefreshScope
   @RestController
   public class ConfigClientController {
       @Value("${config.info}")
       private String configInfo;
   
       @GetMapping("/configInfo")
       public String getConfigInfo() {
           return configInfo;
       }
   }
   ```

4. 发送更新请求`curl -x POST "http://localhost:3355/actuator/refresh"`
