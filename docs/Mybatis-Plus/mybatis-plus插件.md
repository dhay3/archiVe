# mybatis-plus 插件

### #DO

```java

@Data
@Accessors(chain = true)
@TableName("tbl_employee")
public class Employee implements Serializable {

    private static final long serialVersionUID = 1L;

    @TableId(value = "id", type = IdType.AUTO)
    private Integer id;

    private String lastName;

    private String email;

    private String gender;

    private Integer age;
    //标明该字段是乐观锁的字段
    @Version
    private Integer version;
    //标明该字段是逻辑删除字段
    @TableLogic
    private Integer deleted;
    //自动插入字段
    @TableField(fill = FieldFill.INSERT)
    @JsonDeserialize(using = LocalDateTimeDeserializer.class)
    @JsonFormat(pattern = "yyyy-MM-dd HH:mm:ss",timezone = "GMT+8")
    private LocalDateTime date;
}

```

### #配置类

```java
@Configuration
public class MybatisPlusConf {
    //分页配置
    @Bean
    public PaginationInterceptor paginationInterceptor() {
        PaginationInterceptor paginationInterceptor = new PaginationInterceptor();
        //防止恶意sql注入
        ArrayList<ISqlParser> sqlParsers = new ArrayList<>();
        sqlParsers.add(new BlockAttackSqlParser());
        paginationInterceptor.setSqlParserList(sqlParsers);
        return paginationInterceptor;
    }

    @Bean//乐观锁插件
    public OptimisticLockerInterceptor optimisticLockerInterceptor() {
        return new OptimisticLockerInterceptor();
    }
}
```

```java
@Slf4j
@Component
public class MyMetaObjectHandler implements MetaObjectHandler {
    /*
    insert时自动填充指定字段
     */
    @Override
    public void insertFill(MetaObject metaObject) {
        log.info("start insert fill ....");
        //二者选一即可
        this.strictUpdateFill(metaObject, "date", LocalDateTime.class, LocalDateTime.now()); // 起始版本 3.3.0(推荐使用)
        this.fillStrategy(metaObject, "date", LocalDateTime.now());
    }

    /*
    update时自动填充指定字段
     */
    @Override
    public void updateFill(MetaObject metaObject) {

    }
}
```

## #sql性能分析插件

- yaml

```yaml
spring:
  datasource:
    #    driver-class-name: com.mysql.cj.jdbc.Driver
    #性能分析插件,将driverClass修改为P6SpyDriver
    driver-class-name: com.p6spy.engine.spy.P6SpyDriver
    #    url: jdbc:mysql://localhost:3306/mp?useUnicode=true&characterEncoding=utf-8&serverTimezone=Asia/Shanghai
    #    修改url为jdbc:p6sy:数据库://地址
    url: jdbc:p6spy:mysql://localhost:3306/mp?useUnicode=true&characterEncoding=utf-8&serverTimezone=Asia/Shanghai
```

- 添加`spy.properties`

```properties
#3.2.1以上使用
modulelist=com.baomidou.mybatisplus.extension.p6spy.MybatisPlusLogFactory,com.p6spy.engine.outage.P6OutageFactory
# 自定义日志打印
logMessageFormat=com.baomidou.mybatisplus.extension.p6spy.P6SpyLogger
#日志输出到控制台
appender=com.baomidou.mybatisplus.extension.p6spy.StdoutLogger
# 使用日志系统记录 sql
#appender=com.p6spy.engine.spy.appender.Slf4JLogger
# 设置 p6spy driver 代理
deregisterdrivers=true
# 取消JDBC URL前缀
useprefix=true
# 配置记录 Log 例外,可去掉的结果集有error,info,batch,debug,statement,commit,rollback,result,resultset.
excludecategories=info,debug,result,commit,resultset
# 日期格式
dateformat=yyyy-MM-dd HH:mm:ss
# 实际驱动可多个
#driverlist=com.mysql.cj.jdbc.Driver
# 是否开启慢SQL记录
outagedetection=true
# 慢SQL记录标准 2 秒
outagedetectioninterval=2
```

