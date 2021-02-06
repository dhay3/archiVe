### logging.file.name, logging.file.path 不能同时生效

看一下spring的官网方文档

| `logging.file.name` | `logging.file.path` | Example    | Description                                                  |
| :------------------ | :------------------ | :--------- | :----------------------------------------------------------- |
| *(none)*            | *(none)*            |            | Console only logging.                                        |
| Specific file       | *(none)*            | `my.log`   | Writes to the specified log file. Names can be an exact location or relative to the current directory. |
| *(none)*            | Specific directory  | `/var/log` | Writes `spring.log` to the specified directory. Names can be an exact location or relative to the current directory. |

logging.file.name 可以指定路径和log文件的名字

logging.file.path 只可以只当log的路径, 不能指定log的名字, 使用缺省值spring.log

二者只可以存在一个

### #log levels

- 设置日志级别

```yaml
logging:
  level:
    root: debug #指定根日志记录级别, 即全局
    com.chz.controller: trace #指定包日志记录级别
debug: true #输出更多装配信息, 但是日志不会被设为debug level
```

#log groups

- 管理不同包的统一日志级别

```yaml
logging:
  level:
    my: error
  group:
    my: [com.chz.controller, com.chz.entity]
```

### #自定义日志文件

| Logging System          | Customization                                                |
| :---------------------- | :----------------------------------------------------------- |
| Logback                 | `logback-spring.xml`, `logback-spring.groovy`, `logback.xml`, or `logback.groovy` |
| Log4j2                  | `log4j2-spring.xml` or `log4j2.xml`                          |
| JDK (Java Util Logging) | `logging.properties`                                         |
