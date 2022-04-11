# Sentinel

## 流控规则

### 阈值类型

- QPS

  当到达阈值时, 再次访问当前资源名, 就会被sentinel拒绝

- 线程数

  只允许指定线程数访问当前资源, 当超过阈值数就会被拒绝

### 流控模式

- 直接

  当超过QPS或是线程数规定的单机阈值时, 直接触发指定的流控效果

- 关联

  当关联资源的资源到达单机阈值时, 就会触发当前资源的控流效果

<img src="..\..\..\imgs\_SpringCloud\Snipaste_2020-09-28_20-25-36.png"/>

  如上图表示当/testB资源到达单机阈值10时, 就会触发/testA的流控效果快速	失败返回`Blocked by Sentinel (flow limiting)`

- 链路

  具体参考:https://github.com/alibaba/Sentinel/issues/1213

### 流控效果

- 快速失败

  当资源被限流时直接返回 `Blocked by Sentinel (flow limiting)`

- warm up

  从一开始的`单机阈值/coldFactor`(coldFactor默认为3)开始, 经过指定的预热时长**逐渐**升至指定的单机阈值

<img src="..\..\..\imgs\_SpringCloud\Snipaste_2020-09-28_20-55-24.png"/>

  这里表示**当流量突然增加**在5sec以内从3单机阈值开始增加阈值到10。期间请求超过阈值时，同样还是返回`Blocked by Sentinel (flow limiting)`

- 排队等待

  当超过超时时间时，将请求放到队列中，等待有空余时间时处理。不会之间将请求返回控流的信息。

## 降级规则

### 熔断策略

- 慢调用比例

  当请求数大于最小请求数，并且慢调用（响应时间大于最大RT）的比率大于阈值。则在接下来的时间内请求会被自动熔断

- 异常比例

  当异常的比例大于阈值时，则会在接下来的熔断时长内熔断服务

- 异常数

  当大于规定的异常数时，则会在接下来的熔断时长内熔断服务

## 热点规则

- 参数索引

  对应热点方法中参数的位置

- 统计窗口时长

- 单机阈值

  在统计窗口时长内的请求阈值

<img src="..\..\..\imgs\_SpringCloud\Snipaste_2020-09-28_21-39-32.png"/>

这里表示请求资源/testHotkey，请求含有请求方法的第0个索引参数，且请求在1sec内请求的次数大于1次，就会触发服务降级。

服务端

- pom

```xml
        <dependency>
            <groupId>com.alibaba.csp</groupId>
            <artifactId>sentinel-parameter-flow-control</artifactId>
        </dependency>
```

- 业务类

  > @SentinelResource针对的是违背sentinel控制台配置的措施

  为了解耦，可以将`@SentinelResource`放在service层

```java
    @GetMapping("/testHotkey")
    /*
    当超过热点规则限定时,就会触发blockHandler指定的方法
    value, 资源名,一般使用请求路径
    blockHandler, 是该服务熔断降级时的处理,方法必须是静态的
    blockHandlerClass, 服务熔断降级处理方法所在的类
     */
    @SentinelResource(value = "/testHotkey",blockHandler = "deal_testHotkey")
    public String testHotkey(@RequestParam(value = "p1",required = false)String p1,
                             @RequestParam(value = "p2",required = false)String p2){
        return "----testHotkey";
    }

    /*
    参数列表需要与绑定的函数参数相同, BlockException用来打印错误信息
     */
    public String deal_testHotkey(String p1, String p2,BlockException exception){
        exception.printStackTrace();
        return "----deal_testHotkey";
    }
```

发送请求`localhost:8401/testHotkey?p1=a`超过阈值，返回降级信息`----deal_testHotkey`, 同样的发送请求`localhost:8401/testHotkey?p1=a&p2=b`超过阈值时同样会返回降级信息。

但是如果发送请求`localhost:8401/testHotkey?p2=b`并不会返回降级信息。

> 注意我们可以通过指定特殊的参数，来规定特殊的限流规则。需要点击添加按钮

<img src="..\..\..\imgs\_SpringCloud\Snipaste_2020-09-28_21-54-57.png"/>

如果请求的参数value为5时，在统计窗口时长内的阈值为200。请求`localhost:8401/testHotKey?p1=5`

## @SentinelResource

使用`@Sentinel`来规定被限流或是降级后资源应该返回什么信息。

> 修改value值或是请求路径的流控规则或是降级都会对资源生效。
>
> 如果修饰的函数参数使用@RequstParam或是@Pathvarible时，blokHandler函数也要带上

```java
//controller
@RestController
public class RateLimitController {
    @GetMapping("/rateLimit/customerBlockHandler")
    @SentinelResource(value = "customerBlockHandler",
            blockHandlerClass = CustomerBlockHandler.class,
            blockHandler = "handlerException")
    public ResponseBo customerBlockHandler() {
        return ResponseBo.ok();
    }
}
```

> 如果blockHandler与@SentinelResource标注的函数不在同一个类中需要使用static，反之可以不用static

```java
public class CustomerBlockHandler {
    public static ResponseBo handlerException(BlockException exception){
        return ResponseBo.error();
    }
}
```

## 整合Feign

- pom

  ```xml
  <dependency>
      <groupId>org.springframework.cloud</groupId>
      <artifactId>spring-cloud-starter-openfeign</artifactId>
  </dependency>
  ```

- FeignClient

  `@FeignClient`修饰的接口可以不用注入到IOC, 但是实现类必须注入到IOC中

  ```java
  @FeignClient(value = "nacos-payment-provider",fallback = PaymentFeignServiceImpl.class,configuration = PaymentFeignConfiguration.class)
  public interface PaymentFeignService {
      @GetMapping(value = "/paymentSQL/{id}")
       ResponseBo paymentSQL(@PathVariable("id") Long id);
  }
  
  public class PaymentFeignServiceImpl implements PaymentFeignService {
      @Override
      public ResponseBo paymentSQL(Long id) {
          return ResponseBo.error().data("msg","feign 调用失败");
      }
  }
  
  @Configuration
  public class PaymentFeignConfiguration {
      @Bean
      public PaymentFeignServiceImpl paymentService(){
          return new PaymentFeignServiceImpl();
      }
  }
  ```

- controller

  ```java
  @RestController
  public class CircleBreakerFeignController {
  
      @Autowired
      private PaymentFeignService paymentFeignService;
      @GetMapping(value = "/consumer/paymentSQL/{id}")
      @SentinelResource(value ="/consumer/paymentSQL/{id}",blockHandler = "consumerFallback")
      public ResponseBo paymentSQL(@PathVariable("id") Long id) {
          return paymentFeignService.paymentSQL(id);
      }
  
      public ResponseBo consumerFallback(@PathVariable("id") Long id, BlockException e){
          return ResponseBo.error().data("blockhandler","sentinel");
      }
  }
  ```

  > 如果调用的api同时存在Feign fallback 和 Sentinel 流控或是降级规则时， Feign fallbakc生效

## 规则持久化

具体参考：

https://github.com/alibaba/spring-cloud-alibaba/wiki/Sentinel

https://github.com/alibaba/spring-cloud-alibaba/blob/master/spring-cloud-alibaba-examples/sentinel-example/sentinel-core-example/readme-zh.md

- pom

```xml
<dependency>
    <groupId>com.alibaba.csp</groupId>
    <artifactId>sentinel-datasource-nacos</artifactId>
</dependency>

```

- 配置文件

其中`nacos`，`zk`，`apollo`，`redis` 这4种类型的使用需要加上对应的依赖`sentinel-datasource-nacos`, `sentinel-datasource-zookeeper`, `sentinel-datasource-apollo`, `sentinel-datasource-redis`

```yaml
  cloud:
    sentinel:
      datasource: 
        ds1:
          nacos: #对应数据源
            server-addr: localhost:8848
            dataId: cloudalibaba-sentinel-service
            groupId: DEFAULT_GROUP
            data-type: json
            rule-type: flow
```

data-type表示数据的类型支持json或xml（xml需要添加jackson-dataformat-xml依赖）。`rule-type` 配置表示该数据源中的规则属于哪种类型的规则(`flow`，`degrade`，`authority`，`system`, `param-flow`, `gw-flow`, `gw-api-group`)`flow`表示限流

- nacos配置文件

<img src="..\..\..\imgs\_SpringCloud\Snipaste_2020-09-30_12-28-00.png"/>

resource：表示资源名

limitApp：表示来源应用

grade：阈值类型，0表示线程数，1表示QPS

count：单机阈值

strategy：流控模式，0表示直接，1表示关联，2表示链路

controlBehavior：流控效果，0表示快速失败，1表示warm up，2表示排队等待

clusterMode：是否集群





