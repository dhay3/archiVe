# mybatis-plus 中的LocalDateTime, LocalDate, LocalTime

**#Mybatis 从3.4.5 开始，默认支持 JSR-310（日期和时间 API）** 

 **即java.time.* 下的时间类自动类型转换**

本文使用的依赖

```xml
		<dependency>
            <groupId>com.alibaba</groupId>
            <artifactId>druid-spring-boot-starter</artifactId>
            <version>1.1.17</version>
        </dependency>
        <dependency>
            <groupId>com.baomidou</groupId>
            <artifactId>mybatis-plus-boot-starter</artifactId>
            <version>3.3.1</version>
        </dependency>
```

查询时, 会发现如下错误

```java
org.springframework.dao.InvalidDataAccessApiUsageException: Error attempting to get column 'date' from result set.  Cause: java.sql.SQLFeatureNotSupportedException
; null; nested exception is java.sql.SQLFeatureNotSupportedException
```

查看pom文件发现, mybatis版本高于3.4.5

<img src="D:\java资料\我的笔记\springboot\img\2.PNG" alt="2" style="zoom:60%;" />

与数据库相关的依赖处理mybatis还有druid, 最后发现是druid的依赖出错了

修改为如下版本:

```xml
		<dependency>
            <groupId>com.alibaba</groupId>
            <artifactId>druid-spring-boot-starter</artifactId>
            <version>1.1.22</version>
        </dependency>
```

