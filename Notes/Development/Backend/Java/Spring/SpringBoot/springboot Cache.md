# springboot Cache

### #使用缓存

使用springboot的缓存时添加依赖, 或直接使用redis的starter

```
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-cache</artifactId>
</dependency>
```

并在入口类加入`@EnableCaching`开启缓存功能

```java
@SpringBootApplication
@EnableCaching
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class,args);
    }
}
```

### #几个重要概念, 缓存注解

| 名称           | 解释                                                         |
| -------------- | ------------------------------------------------------------ |
| Cache          | 缓存接口,定义缓存操作. 实现由RedisCache, EhCacheCache, ConcurrentMapCache |
| CahcerManager  | 缓存管理器, 管理各种缓存组件                                 |
| @Cacheable     | 调用方法前会根据参数从缓存中查看是否有缓存,如果有就不会调用方法,直接从缓存中使用 |
| @CacheEvict    | 清空缓存                                                     |
| @CachePut      | 保证方法被调用, 又希望结果被缓存, 与@Cacheable区别在于是否每次都调用方法,常用于更新 |
| @EnableCaching | 开启缓存                                                     |
| @CacheConfig   | 统一配置本类的缓存注解的属性                                 |

### #@Cacheable/@CachePut/@CacheEvict 主要的参数

| 名称                               | 解释                                                         |
| ---------------------------------- | ------------------------------------------------------------ |
| value                              | 缓存的名字, 在sprig配置文件中定义, 必须指定至少一个<br/> 例如:<br/>@Cacheable(value="cache1")或者@Cacheable(value={"c1","c2"}) |
| key                                | 缓存的key, 可以为空, 如果指定要按照SpEL 表达式编写, <br>如果不指定,则按照方法的所有参数进行组合<br>一般采用唯一的值做为key, 可以主键,唯一的id |
| condition                          | 缓存的条件, 可以为空, 使用SpEL编写, 返回true或false                                                                         只有为true才会进行缓存/清空缓冲                                                                                                         例如: @Cacheable(value="testcache",condition="#userName.length()>2") |
| unless                             | 否定缓存。当条件结果为TRUE时，就不会缓存。<br/>@Cacheable(value=”testcache”,unless=”#userName.length()>2”) |
| allEntries<br/>(@CacheEvict )      | 是否清空所有缓存内容，缺省为 false，如果指定为 true，<br/>则方法调用后将立即清空所有缓存<br/>例如：<br/>@CachEvict(value=”testcache”,allEntries=true) |
| beforeInvocation<br/>(@CacheEvict) | 是否在方法执行前就清空，缺省为 false，如果指定为 true，<br/>则在方法还没有执行的时候就清空缓存，缺省情况下，如果方法<br/>执行抛出异常，则不会清空缓存<br/>例如：<br/>@CachEvict(value=”testcache”，beforeInvocation=true) |
| keyGenerator                       | 缓存数据时key生成策略                                        |

可以将value理解为redis中HSET的key, key为HSET中的field

### #SpEL上下文数据

| 名称        | 位置     | 描述                       | 实例                 |
| :---------- | -------- | -------------------------- | -------------------- |
| methodName  | root对象 | 当前被调用的方法名         | \#root.methodname    |
| method      | root对象 | 当前被调用的方法           | \#root.method.name   |
| target      | root对象 | 当前被调用的目标对象实例   | #root.target         |
| targetClass | root对象 | 当前被调用的目标对象的类   | #root.targetClass    |
| args        | root对象 | 当前被调用的方法的参数列表 | #root.args[0]        |
| caches      | root对象 | 当前方法调用使用的缓存列表 | #root.caches[0].name |

!!! **注意**

1.当我们要使用root对象的属性作为key时我们也可以将“#root”省略，因为Spring默认使用的就是root对象的属性。 如:*

```
@Cacheable(key = "targetClass + methodName +#p0")
```

2.使用方法参数时我们可以直接使用“#参数名”或者“#p参数index”。

==存的是参数对应的值==

 如：

```java
@Cacheable(value="users", key="#id")
@Cacheable(value="users", key="#p0")
```

- #### @Cacheable

  配置了函数的返回值将被加入缓存。同时在查询时，会先从缓存中获取(参数的值对应的key)，若不存在才再发起对数据库的访问。该注解主要有下面几个参数：

  ```java
      @Cacheable(value = "emp" ,key = "targetClass + methodName +#p0")
      public List<NewJob> queryAll(User uid) {
          return newJobDao.findAllByUid(uid);
      }
  ```

  此处的`User`实体类一定要实现序列化`public class User implements Serializable`，否则会报`java.io.NotSerializableException`异常。

  到这里，你已经可以运行程序检验缓存功能是否实现。

  - `value`、`cacheNames`：两个等同的参数（cacheNames为Spring 4新增，作为value的别名），用于指定缓存存储的集合名。由于Spring 4中新增了`@CacheConfig`，因此在Spring 3中原本必须有的value属性，也成为非必需项了；
  - `key`：缓存对象存储在Map集合中的key值，非必需，缺省按照函数的所有参数组合作为key值，若自己配置需使用SpEL表达式，比如：`@Cacheable(key = "#p0")`：使用函数第一个参数作为缓存的key值，更多关于SpEL表达式的详细内容可参考https://docs.spring.io/spring/docs/current/spring-framework-reference/integration.html#cache；
  - `condition`：缓存对象的条件，非必需，也需使用SpEL表达式，只有满足表达式条件的内容才会被缓存，比如：`@Cacheable(key = "#p0", condition = "#p0.length() < 3")`，表示只有当第一个参数的长度小于3的时候才会被缓存；
  - `unless`：另外一个缓存条件参数，非必需，需使用SpEL表达式。它不同于condition参数的地方在于它的判断时机，该条件是在函数被调用之后才做判断的，所以它可以通过对result进行判断；
  - `keyGenerator`：用于指定key生成器，非必需。若需要指定一个自定义的key生成器，我们需要去实现`org.springframework.cache.interceptor.KeyGenerator`接口，并使用该参数来指定；
  - `cacheManager`：用于指定使用哪个缓存管理器，非必需。只有当有多个时才需要使用；
  - `cacheResolver`：用于指定使用那个缓存解析器，非必需。需通过org.springframework.cache.interceptor.CacheResolver接口来实现自己的缓存解析器，并用该参数指定；

- #### @CacheConfig

  当我们需要缓存的地方越来越多，你可以使用`@CacheConfig(cacheNames = {"myCache"})`注解来统一指定`value`的值，这时可省略`value`，如果你在你的方法依旧写上了`value`，那么依然以方法的`value`值为准。

  ```java
  @CacheConfig(cacheNames = {"myCache"})
  public class BotRelationServiceImpl implements BotRelationService {
      @Override
      @Cacheable(key = "targetClass + methodName +#p0")//此处没写value
      public List<BotRelation> findAllLimit(int num) {
          return botRelationRepository.findAllLimit(num);
      }
      .....
  }
  ```

- #### @CachePut

  它与`@Cacheable`不同的是，它每次都会真实调用函数(即访问数据库)，所以主要用于数据==修改和添加==操作上。当我们更新数据的时候，应该使用`@CachePut`进行缓存数据的==更新==，否则将查询到脏数据(即如果不采用, 数据库中数据更新, 但是缓存中数据没有更新), 因为缓存中存的是返回值, 所以==要再查询==一遍将查询结果缓存

  如果不是采用唯一的key, 拿@Cacheput的key和value要与@Cacheable的一致

  ```java
    @CachePut(key = "#p0.sno")
      public Student update(Student student) {
          this.studentMapper.update(student);
          //这里的查询访问的数据库
          return this.studentMapper.queryStudentBySno(student.getSno());
      }
  ```
  
- #### @CacheEvict

  ```java
      @Cacheable(value = "emp",key = "#p0.id")
      public NewJob save(NewJob job) {
          newJobDao.save(job);
          return job;
      }
   
      //清除value中的一条缓存，key为要清空的数据
      @CacheEvict(value="emp",key="#id")
      public void delect(int id) {
          newJobDao.deleteAllById(id);
      }
   
      //方法调用后清空value中的所有缓存
      @CacheEvict(value="accountCache",allEntries=true)
      public void delectAll() {
          newJobDao.deleteAll();
      }
   
      //方法调用前清空value中的所有缓存
      @CacheEvict(value="accountCache",beforeInvocation=true)
      public void delectAll() {
          newJobDao.deleteAll();
      }
  ```

- #### @Caching

有时候我们可能组合多个Cache注解使用，此时就需要@Caching组合多个注解标签了。

```java
  @Caching(cacheable = {
            @Cacheable(value = "emp",key = "#p0"),
            ...
    },
    put = {
            @CachePut(value = "emp",key = "#p0"),
            ...
    },evict = {
            @CacheEvict(value = "emp",key = "#p0"),
            ....
    })
    public User save(User user) {
        ....
    }
```

==下面讲到的整合第三方缓存组件都是基于上面的已经完成的步骤，所以一个应用要先做好你的缓存逻辑，再来整合其他cache组件。==

不使用其他缓存默认采用spring的简单缓存
