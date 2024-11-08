# springboot 整合Redis

##### #准备工作

- 配置application.yml

```
spring:
  thymeleaf: #thymeleaf
    cache: false
  datasource: #datasource
    driver-class-name: com.mysql.cj.jdbc.Driver
    url: jdbc:mysql://localhost:3306/mp?useUnicode=true&characterEncoding=utf-8&serverTimezone=Asia/Shanghai
    username: root
    password: 12345
    type: com.alibaba.druid.pool.DruidDataSource
    druid:
      initial-size: 5 #初始连接数
      max-active: 10 #最大活动连接
      max-wait: 60000 #从池中取连接(没有闲置连接)的最大等待时间,-1表示无限等待
      min-idle: 5 #最小闲置数,小于min-idle连接池会主动创建新的连接
      time-between-eviction-runs-millis: 60000 #清理线程启动的间隔时间,当线程池中没有可用的连接启动清理线程
      min-evictable-idle-time-millis: 300000 #清理线程保持闲置最小时间
      validation-query: SELECT 1  #用于校验连接
      test-on-borrow: false #请求连接时是否校验,比较消耗性能,一般设置false
      test-on-return: false #归还连接时是否校验,比较消耗性能,一般设置false
      test-while-idle: true #清理线程通过validation-query来校验连接是否正常,如果不正常将从连接池中移除
      pool-prepared-statements: true #存储相同逻辑的sql到连接池的缓存中
      #      filters: stat,wall #监控统计web的statement(com.chz.com.chz.sql),以及防sql注入的wall
      # 关闭如上配置,可以采用自定义的filter
      filter:
        stat:
          enabled: true #状态监控-->stat
          db-type: mysql
          log-slow-sql: true  #记录超过指定时间的sql到日志中
          slow-sql-millis: 2000
        wall:
          enabled: true #防火墙-->wall
          db-type: mysql
          config:
            delete-allow: false #禁止删除
            drop-table-allow: false #禁止删除表
      web-stat-filter:
        enabled: true #开启监控uri,默认false
        url-pattern: /* #添加过滤规则
        exclusions: "*.js,*.gif,*.jpg,*.png,*.css,*.ico,/druid" #忽略过滤
      stat-view-servlet:
        enabled: true #开启视图管理界面,默认false
        url-pattern: /druid/* #视图管理界面uri
        login-username: druid #账号
        login-password: 12345 #密码
  redis:
    database: 0 #redis数据库索引, 默认 0
    host: localhost #ip 默认localhost
    port: 6379 #端口 默认6379
    lettuce: #使用lettuce客户端线程安全, jedis线程不安全
      pool:
        max-active: 8
        max-wait: -1 #-1表示无限制
        max-idle: 8
        min-idle: 0 #为0表示不会主动创建连接
#    timeout: 0 #超过指定时间抛出异常
#        allow: 127.0.0.1 白名单
#        deny:  192.168.1.130黑名单
mybatis: #mybatis
  configuration:
    log-impl: org.apache.ibatis.logging.stdout.StdOutImpl
```

- 引入依赖

```
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-cache</artifactId>
</dependency>
```

在主入口加入@EnableCaching

```java
@EnableCaching
@SpringBootApplication
public class SpringbootMybatisRedisApplication {
    public static void main(String[] args) {
        SpringApplication.run(SpringbootMybatisRedisApplication.class, args);
    }
}

```

##### #pojo

这里埋了一个坑

```java

@Data
@Accessors(chain = true)//fluent会导致json数据无法正确输出
//@Validated自会对controller层生效,不会对mybatis入库数据校验
public class Employee implements Serializable {
    private static final long serialVersionUID = 1L;
    private Integer id;
    private String lastName;
    private String email;
    private String gender;
    private Integer age;
    private Integer version;
    private Integer deleted;
    private LocalDateTime date;

    @Override
    public String toString() {
        return "Employee{" +
                "id=" + id +
                ", lastName='" + lastName + '\'' +
                ", email='" + email + '\'' +
                ", gender='" + gender + '\'' +
                ", age=" + age +
                ", date=" + date +
                '}';
    }
}

```

##### #mapper

```java

@Mapper
public interface EmployeeMapper {
    @SelectProvider(type = UserSqlProvider.class)
    @Results(id = "employeeMap",
            value = {
                    @Result(property = "id", column = "id", id = true),
                    @Result(property = "lastName", column = "last_name")
            })
    List<Employee> list(String name, Integer id);

    //使用原生的NOW()函数自动插入时间
    @Insert("INSERT INTO tbl_employee (last_name,email,gender,age,date)VALUES" +
            "(#{lastName},#{email},#{gender},#{age},NOW())")
    boolean add(Employee employee);

//    @UpdateProvider(type = UserSqlProvider.class)
//    boolean update(String name, Integer age, Integer id);

    @Select("SELECT * FROM tbl_employee WHERE id = #{id}")
    @ResultMap("employeeMap")
    Employee get(Integer id);

    @Update("UPDATE tbl_employee SET " +
            "last_name=#{lastName}, email=#{email},gender=#{gender},age=#{age},date=NOW()" +
            "WHERE id=#{id}")
    boolean update(Employee employee);

    @Delete("DELETE FROM tbl_employee WHERE id = #{id}")
    boolean delete(Integer id);

}

```

##### #service

@Cacheable, @CachePut, @CacheEvict 自行百度

```java
/**
 * 如果不指定缓存,spring默认采用内置简单缓存,生产一般要指定具体的缓存
 * 缓存的是方法的返回值
 */
@CacheConfig(cacheNames = {"emps"})//配置统一的cache属性
public interface IEmployeeService {
   
    //添加到数据库中,同时修改缓存中的数据
    @CachePut(key = "#p0.id")
    public Employee add(Employee employee);

    //因为缓存存的是返回值,返回Employee是为了修改缓存中值,避免脏读
    @CachePut(key = "#p0.id")
    Employee update(Employee employee);
    
	 /*
    @Cacheable,在调用方法前会查看value中对应key的缓存,如果又就不会调用函数
       value: 缓存域,可以理解为redis中Hset的key
       key: 缓存键,可以理解为redis中Hset的field,但是这里存的是参数的值
     */
    @Cacheable(key = "#p0")//一般用唯一的值,数据库采用主键
    public Employee get(Integer id);

    @CacheEvict(key = "#p0")//删除时删除value中key为id值的缓存
    public boolean delete(Integer id);
}

```

```java
@Service
public class EmployeeServiceImpl implements IEmployeeService {
    @Autowired
    EmployeeMapper employeeMapper;


    @Override
    public boolean add(Employee employee) {
        return employeeMapper.add(employee);
    }

    @Override
    public Employee update(Employee employee) {
        employeeMapper.update(employee);
        return employeeMapper.get(employee.getId());
    }

    @Override
    public Employee get(Integer id) {
        return employeeMapper.get(id);
    }

    @Override
    public boolean delete(Integer id) {
        return employeeMapper.delete(id);
    }
}

```

##### #RedisTemplate

可以直接使用`spring`提供的`RedisTemplate`,  默认是`StringRedisTemplate`

但是value ==必须是String==, 否则抛出异常

```java
//设置key的序列化器,两者选一,value同理,但是这种的弊端就是每次都需要设置
template.setKeySerializer(RedisSerializer.string());
template.setKeySerializer(new StringRedisSerializer());
```

```java
opsForValue();// 对应redis的set,get
opsForHash(); //对应redis的hset,hget,hmset,hmget
opsForList();// 对应redis的lpush,rpush,lpop,rpop,lrange
opsForSet(); //对应redis 的set
opsForZset(); //对应redis 的zset
```

生产中为了多样化使用需要自定义Redis的配置类

##### #Redis配置类

```java
//redis配置类
@Configuration
public class RedisConf extends CachingConfigurerSupport {
    /*
    配置自定义RedisTemplate
     */
    @Bean
    public RedisTemplate<String, Object> redisTemplate(RedisConnectionFactory factory) {
        RedisTemplate<String, Object> template = new RedisTemplate<>();
        //配置连接工厂
        template.setConnectionFactory(factory);
        template.setKeySerializer(keySerializer());
        template.setValueSerializer(valueSerializer());
        template.setHashKeySerializer(keySerializer());
        template.setValueSerializer(valueSerializer());
        return template;
    }

    /*
    org.springframework.cache.interceptor包下的
    自定义key的生成策略,对应@Cacheable中的keyGenerator
    实例对象+方法名+参数名
     */
    @Bean
    public KeyGenerator keyGenerator() {
        /*
        target调用缓存方法的实例
        method调用缓存的方法
        params方法的参数
         */
        return (tagert, method, params) -> {
            StringBuilder sb = new StringBuilder();
            sb.append(tagert.getClass().getName())
                    .append(method.getName());
            for (Object param : params) {
                sb.append(param.toString());
            }
            //返回key
            return sb.toString();
        };
    }

    /*
    自定义缓存管理
     */
    @Bean
    public CacheManager cacheManager(RedisConnectionFactory factory) {
        return RedisCacheManager.builder(factory).
                //默认缓存策略
                        cacheDefaults(redisCacheConfiguration(600L)).
                //配置不同缓存域,不同过期时间
                        withInitialCacheConfigurations(RedisCacheConfigurationMap()).
                //更新删除上锁
                        transactionAware().
                        build();
    }

    /*
    配置redis的cache策略
     */
    private RedisCacheConfiguration redisCacheConfiguration(Long sec) {
        return RedisCacheConfiguration.defaultCacheConfig().
                //设置key的序列化,采用stringRedisSerializer
                        serializeKeysWith
                        (RedisSerializationContext.SerializationPair.fromSerializer(keySerializer())).
                //设置value的序列化，采用Jackson2JsonRedis
                        serializeValuesWith
                        (RedisSerializationContext.SerializationPair.fromSerializer(valueSerializer())).
                //设置cache的过期策略
                        entryTtl(Duration.ofSeconds(sec)).
                //不缓存null的值
                        disableCachingNullValues();
    }

    /*
    不同缓存域,不同过期时间,map的key可以被@Cacheable中的value使用
     */
    private Map<String, RedisCacheConfiguration> RedisCacheConfigurationMap() {
        Map<String, RedisCacheConfiguration> redisCacheConfigurationMap = new HashMap<>();
        redisCacheConfigurationMap.put("userInfo", redisCacheConfiguration(3000L));
        redisCacheConfigurationMap.put("otherInfo", redisCacheConfiguration(1000L));
        return redisCacheConfigurationMap;
    }

    /*
    key采用序列化策略
     */
    private RedisSerializer<String> keySerializer() {
        return new StringRedisSerializer();
    }

    /*
    value采用序列化策略
     */
    private RedisSerializer<Object> valueSerializer() {
        Jackson2JsonRedisSerializer<Object> serializer = new Jackson2JsonRedisSerializer<>(Object.class);
        //序列化所有类包括jdk提供的
        ObjectMapper om = new ObjectMapper();
        //设置序列化的域(属性,方法etc)以及修饰范围,Any包括private,public 默认是public的
        //ALL所有方位,ANY所有修饰符
        om.setVisibility(PropertyAccessor.ALL, JsonAutoDetect.Visibility.ANY);
        //enableDefaultTyping 原来的方法存在漏洞,2.0后改用如下配置
        //指定输入的类型
        om.activateDefaultTyping(LaissezFaireSubTypeValidator.instance,
                ObjectMapper.DefaultTyping.NON_FINAL);
        //如果java.time包下Json报错,添加如下两行代码
        om.disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS);
        om.registerModule(new JavaTimeModule());
        
        serializer.setObjectMapper(om);
        return serializer;
    }
}

```

运行结果发现`Cannot deserialize instance of java.time.LocalDateTime out of START_ARRAY`

那是因为LocalDateTime 这一系列比较特殊,  需要指定序列化的方式

```java
	@JsonDeserialize(using = LocalDateTimeDeserializer.class)
    @JsonSerialize(using = LocalDateTimeSerializer.class)
    private LocalDateTime date;
```

或是在`RedisSerializer<Object>valueSerializer()`中添加如下两行代码

```java
        om.disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS);
        om.registerModule(new JavaTimeModule());
```

测试发现, 只会执行一次sql

<img src=".\img\1.PNG"/>

redis整合成功!
