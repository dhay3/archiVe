# Nacos

[TOC]

相当于过去的`Eruka`组件, 提供服务注册与发现功能

链接:

https://github.com/alibaba/nacos

https://nacos.io/zh-cn/docs/quick-start.html

### 启动Nacos

下载并解压, 运行`.\bin\startup.cmd`

访问如下地址栏

<img src="D:\java资料\springcloud\cloud新\img\1.png" style="zoom:50%;" />

账号`nacos`, 密码`nacos`

<img src="..\..\..\..\java资料\springcloud\cloud新\img\2.png"/>

### 生产消费者注册

生产者消费者配置步骤相同, 详细请参考https://nacos.io/zh-cn/docs/quick-start-spring-cloud.html

#### 添加依赖

```xml
<dependency>
    <groupId>com.alibaba.cloud</groupId>
    <artifactId>spring-cloud-starter-alibaba-nacos-config</artifactId>
    <version>${latest.version}</version>
</dependency>
```

#### 配置Nacos

```properties
server.port=8070
spring.application.name=service-provider

spring.cloud.nacos.discovery.server-addr=127.0.0.1:8848
```

#### 修改启动类

```java
@SpringBootApplication
@EnableDiscoveryClient
public class NacosProviderApplication {

	public static void main(String[] args) {
		SpringApplication.run(NacosProviderApplication.class, args);
	}

	@RestController
	class EchoController {
		@RequestMapping(value = "/echo/{string}", method = RequestMethod.GET)
		public String echo(@PathVariable String string) {
			return "Hello Nacos Discovery " + string;
		}
	}
}
```

#### 消费者调用

- 方法一/ 使用`RestTemplate`

  > 必须添加@LoadBalance

```java
@SpringBootApplication
@EnableDiscoveryClient
public class NacosConsumerApplication {

    @LoadBalanced
    @Bean
    public RestTemplate restTemplate() {
        return new RestTemplate();
    }

    public static void main(String[] args) {
        SpringApplication.run(NacosConsumerApplication.class, args);
    }

    @RestController
    public class TestController {

        private final RestTemplate restTemplate;

        @Autowired
        public TestController(RestTemplate restTemplate) {this.restTemplate = restTemplate;}

        @RequestMapping(value = "/echo/{str}", method = RequestMethod.GET)
        public String echo(@PathVariable String str) {
            return restTemplate.getForObject("http://service-provider/echo/" + str, String.class);
        }
    }
}
```

- 方法二 / 使用`Feign`

添加`FeignClient`

```java
@Component
@FeignClient("service-vod")//对应application name
public interface VoDClientFeign {
    /*
    这里的value必须是全路径, 在该FeignClient下的方法的参数必须使用
    @ReqeustParam或是@PathVariable标注, @RequestBody有且只有一个
    */
    @RequestMapping(value = "/echo/{str}", method = RequestMethod.GET)
    public String echo(@PathVariable String str) 

```

启动类

```java
@EnableFeignClients//如果在同一个模块中就无需使用basePackages
@EnableDiscoveryClient//注册服务到nacos
@SpringBootApplication(
        scanBasePackages = {"com.chz.eduservice", "com.chz.servicebase"})
public class EduApplication {
    public static void main(String[] args) {
        SpringApplication.run(EduApplication.class, args);
    }
}
```

调用`FeignClient`提供的接口, 访问指定`application name`的指定接口

```java
@RestController
@CrossOrigin
public class VideoController {
  
    @Autowired
    private VoDClientFeign voDClientFeign;
    
    @RequestMappig
    public void echoByFeign(){
        //"hello world"作为PathVariable被传入到接口中
        voDClientFeign.echo("hello world");
    }
    
```

