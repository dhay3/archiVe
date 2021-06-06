# SpringCloud/gateway

参考:

https://www.cnblogs.com/babycomeon/p/11161073.html

[TOC]

## 依赖

```xml
		<dependency>
			<groupId>org.springframework.cloud</groupId>
			<artifactId>spring-cloud-starter-gateway</artifactId>
		</dependency>
```

==Spring Cloud Gateway 是使用 netty+webflux 实现因此不需要再引入 web 模块。==

## 配置

### routes

- id

  路由的id, 可以随便写, 一般设置为路由的`spring.application.name`

- uri

  请求转发的目标地址

- predicate

  断言, 只有断言为`true`才会转发到目标地址, 一般用于匹配路径

As Follow

```yaml
server:
  port: 8080
spring:
  application:
    name: api-gateway
  cloud:
    nacos:
      discovery:
        server-addr: 127.0.0.1:8848 #和zull相同, 同样需要注入到服务注册中心
    gateway:
      enabled: true #开启网关, 缺省为true
      routes:
        - id: edu-service #服务的id, 可以是任意值
          uri: http://localhost:8001 #请求转发的uri
          predicates:
            - Path=/eduservice/** #只有断言为true才会转发到目标地址, 一般用于匹配路径
```

如上配置表示, 配置了一个id为`edu-service`, 

当访问地址`http://localhost:8080/eduservice/teacher/test`

会将请求转发到`http://localhost:8001/eduservice/teacher/test`

和`nginx`的匹配的路径相似

## 动态路由

动态路由可以不用指定固定端口

```yaml
server:
  port: 9527
spring:
  application:
    name: cloud-gateway
  cloud:
    gateway:
      discovery:
        locator:
          enabled: true #动态路由
      routes:
        - id: payment
          uri: lb://cloud-payment-service #动态路由
          predicates:
            - Path=/payment/**
```

## Predicates

具体参考https://docs.spring.io/spring-cloud-gateway/docs/2.2.5.RELEASE/reference/html/#configuration

通过如下代码生成指定格式的日期

```java
public class Test {
    public static void main(String[] args) {
        ZonedDateTime now = ZonedDateTime.now();
        System.out.println(now);
    }
}
```

- After：表示在指定时间之后生效，同理Before，between

  ```yaml
            predicates:
              - Path=/payment/**
              - After=2020-09-26T10:06:14.022329700+08:00[Asia/Shanghai]
  ```

- Cookie：请求中含有指定cookie时才匹配

  ```yaml
            predicates:
              - Path=/payment/**
              - Cookie=name,chz
  ```

  通过`curl localhost:9527/payment/lb --cookie "name=chz"`测试

- Header：请求中含有指定请求头才会匹配

  ```yaml
            uri: lb://cloud-payment-service #动态路由
            predicates:
              - Path=/payment/**
              - Header=X-Request-id,123
  ```

  通过`curl localhost:9527/payment/lb --header "X-Request-Id:123""`测试

## Filter

过滤请求，与Predicates相似，具体参考：https://docs.spring.io/spring-cloud-gateway/docs/2.2.5.RELEASE/reference/html/#gatewayfilter-factories

### Global Filter

```java
@Slf4j
@Component
public class CustomFilter implements GlobalFilter, Ordered {
    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        log.info("**************into filter******************");
        String uname = exchange.getRequest().getQueryParams().getFirst("uname");
        if (StringUtils.isEmpty(uname)){
            log.info("******非法用户*********");
            exchange.getResponse().setStatusCode(HttpStatus.NOT_ACCEPTABLE);
            return exchange.getResponse().setComplete();
        }
        //请求放行
        return chain.filter(exchange);
    }

    @Override
    public int getOrder() {
        return Ordered.HIGHEST_PRECEDENCE;
    }
```

我们通过`curl -G http://localhost:9527/payment/lb?uname=zs`来测试
