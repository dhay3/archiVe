# @SpringbootApplication

### #注解

- @SpringBootConfiguration

  |- @Configuration

   	|- @Component

​    实际作用就是将主启动类注入到ioc中

- @EnableAutoConfiguration 

  |-@AutoConfigurationPackage  标注主启动所在的包

  |-@Import  相当于`<import resource="..."/>`引入一个配置类

  ​     `AutoConfigurationImportSelector`用来自动扫描出`@Configuration`注解的类

  ​	并将配置类import到ioc中

- @ComponentScan

  扫描范围

  参考https://www.cnblogs.com/xingjia/p/11184876.html

  用于过滤类不加入到ioc中

  例如:

  ```java
  @ComponentScan(value="com.chz",excludeFilters= {
          @Filter(type=FilterType.ANNOTATION,classes= {Controller.class})
  })
  
  ```

  `com.chz`下`@Controller`不被扫描注入到ioc中

  ```java
  @ComponentScan(value="com.chz",
  includeFilters= {
          @Filter(type=FilterType.ASSIGNABLE_TYPE,classes= {BookService.class})
  })
  ```

  会加载`BookService`,以及`BookService`的子类或者其实现类

### #属性

- exclude == excludeName == @EnableAutoConfiguration

  会覆盖原有配置, 不建议配置

- scanBasePackages 等价于 @CompoentScan (basepages = "xxx")

  会覆盖原有配置的@CompoentScan

- proxyBeanMethods 是否能通过方法调用来获取JavaBean
