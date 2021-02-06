# bootstrap.yml 和 application.yml

Spring Boot 默认支持 properties(.properties) 和 YAML(.yml .yaml ) 两种格式的配置文件，yml 和 properties 文件都属于配置文件，功能一样。

Spring Cloud 构建于 Spring Boot 之上，在 Spring Boot 中有两种上下文，一种是 bootstrap，另外一种是 application，下面列举这两种配置文件的区别

### #加载顺序

若application.yml 和bootstrap.yml 在同一目录下：**bootstrap.yml 先加载 application.yml后加载**

bootstrap.yml 用于应用程序上下文的引导阶段。bootstrap.yml 由父Spring ApplicationContext加载。

### #配置区别

bootstrap.yml 和 application.yml 都可以用来配置参数。

bootstrap.yml 用来程序引导时执行，应用于更加早期配置信息读取。可以理解成系统级别的一些参数配置，这些参数一般是不会变动的。一旦bootStrap.yml 被加载，则内容不会被覆盖。

application.yml 可以用来定义应用级别的， 应用程序特有配置信息，可以用来配置后续各个模块中需使用的公共参数等。

### #属性覆盖问题

启动上下文时，Spring Cloud 会创建一个 Bootstrap Context，作为 Spring 应用的 Application Context 的父上下文。

初始化的时候，Bootstrap Context 负责从外部源加载配置属性并解析配置。这两个上下文共享一个从外部获取的 Environment。**Bootstrap 属性有高优先级，默认情况下，它们不会被本地配置覆盖。**

也就是说如果加载的 application.yml 的内容标签与 bootstrap 的标签一致，application 也不会覆盖 bootstrap，而 application.yml 里面的内容可以动态替换。

### #bootstrap.yml典型的应用场景

- 当使用 Spring Cloud Config Server 配置中心时，这时需要在 bootstrap.yml 配置文件中指定 spring.application.name 和 spring.cloud.config.server.git.uri，添加连接到配置中心的配置属性来加载外部配置中心的配置信息
- 一些固定的不能被覆盖的属性
- 一些加密/解密的场景
  ————————————————
  版权声明：本文为CSDN博主「ThinkWon」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
  原文链接：https://blog.csdn.net/ThinkWon/article/details/100007093
