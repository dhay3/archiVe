# @Trasanctional

在springboot中无需开启`@EnableTransactionManagement`(springboot自动替我们配置), 只需使用`@Transactional`即可开启事务

### #spring的事务回滚机制

如果方法执行成功, 会隐式提交事务

不管是checked exception或是 unchecked exception如果异常被try-catch了, 那么方法会照常执行,事务不生效

==如果checked exception 抛出异常, 事务是有效的==,这里不同于@ExceptionHandler和aop中处理异常的机制

所以事务只针对unchecked exception 回滚如空指针,数组越界,etc 

如果想要checked exception 也回滚可以在catch中抛出一个RuntimeException()即可

<img src="D:\java资料\我的笔记\springboot\img\4.PNG" alt="4" style="zoom:80%;" />

注意：**上图中有个错误 - ClassNotFoundException不属于运行时异常！**

### #@Transactional 注解属性

- ##### **propagation**  

  缺省值为Propagation.REQUIRED

  - **Propagation.REQUIRED**

    如果当前存在事务, 则加入该事务, 如果当前不存在事务, 则创建一个事务

    如果a方法和b方法都添加了注解, 使用默认传播模式, 则a方法内调用b方法, 会把两个方法的事务合并为一个事务

    思考一个问题, 如果b方法内部抛出异常, 而a方法catch了b方法的异常, 那这个事务还能正常运行吗?

    答案是不行! 会抛出异常：`org.springframework.transaction.UnexpectedRollbackException: Transaction rolled back because it has been marked as rollback-only`，因为当ServiceB中抛出了一个异常以后，ServiceB会把当前的transaction标记为需要rollback。但是ServiceA中捕获了这个异常，并进行了处理，认为当前transaction应该正常commit。此时就出现了前后不一致，也就是因为这样，抛出了前面的UnexpectedRollbackException。

  - **Propagation.SUPPORTS**

    如果当前存在事务, 则加入该事务; 如果当前不存在事务, 则以非事务的方式继续运行

  - **Propagation.MANDATORY**

    如果当前存在事务, 则加入该事务; 如果当前不存在事务, 则抛出异常

  - **Propagation.REQUIRES_NEW**

    重新创建一个新的事务, 如果当前存在事务, 暂停当前事务

    这个属性可以实现:

    类A中的a方法加上默认注解@Transactional(propagation = Propagation.REQUIRED)，类B中的b方法加上注解@Transactional(propagation = Propagation.REQUIRES_NEW)，然后在a方法中调用b方法操作数据库，再在a方法最后抛出异常，会发现a方法中的b方法对数据库的操作没有回滚，因为Propagation.REQUIRES_NEW会暂停a方法的事务。

  - **Propagation.NOT_SUPPORTED**
    以非事务的方式运行，如果当前存在事务，暂停当前的事务。

  - **Propagation.NEVER**
    以非事务的方式运行，如果当前存在事务，则抛出异常。

  - **Propagation.NESTED**
    和 Propagation.REQUIRED 效果一样。

- ##### isolation

  事务的隔离级别, 缺省值为 Isolation.DEFAULT

  - Isolation.DEFAULT
    使用底层数据库默认的隔离级别。

  - Isolation.READ_UNCOMMITTED  
  - Isolation.READ_COMMITTED
  - Isolation.REPEATABLE_READ
  - Isolation.SERIALIZABLE

- timeout 属性
  事务的超时时间，默认值为-1。如果超过该时间限制但事务还没有完成，则自动回滚事务。

- readOnly 属性
  指定事务是否为只读事务，默认值为 false；为了忽略那些不需要事务的方法，比如读取数据，可以设置 read-only 为 true。

- rollbackFor 属性
  用于指定能够触发事务回滚的异常类型，可以指定多个异常类型。

- noRollbackFor 属性
  抛出指定的异常类型，不回滚事务，也可以指定多个异常类型

### #@Transactional事务几点注意

这里面有几点需要大家留意：
A. 一个功能是否要事务，必须纳入设计、编码考虑。不能仅仅完成了基本功能就ok。
B. 如果加了事务，必须做好开发环境测试（测试环境也尽量触发异常、测试回滚），确保事务生效。
C. 以下列了事务使用过程的注意事项，请大家留意。

1.不要在接口上声明@Transactional ，而要在具体类的方法上使用 @Transactional 注解，否则注解可能无效。

2.不要图省事，将@Transactional放置在类级的声明中，放在类声明，会使得所有方法都有事务。故@Transactional应该放在方法级别，不需要使用事务的方法，就不要放置事务，比如查询方法。否则对性能是有影响的。

3.使用了@Transactional的方法，被同一个类里面的方法调用， @Transactional无效。比如有一个类Test，它的一个方法A，A再调用Test本类的方法B（不管B是否public还是private），但A没有声明注解事务，而B有。则外部调用A之后，B的事务是不会起作用的。（经常在这里出错）

4.使用了@Transactional的方法，只能是public，@Transactional注解的方法都是被外部其他类调用才有效，故只能是public。道理和上面的有关联。故在 protected、private 或者 package-visible 的方法上使用 @Transactional 注解，它也不会报错，但事务无效。

5.spring的事务在抛异常的时候会回滚，如果是catch捕获了，事务无效。可以在catch里面加上throw new RuntimeException();

6.最后有个关键的一点：和锁同时使用需要注意：由于Spring事务是通过AOP实现的，所以在方法执行之前会有开启事务，之后会有提交事务逻辑。而synchronized代码块执行是在事务之内执行的，可以推断在synchronized代码块执行完时，事务还未提交，其他线程进入synchronized代码块后，读取的数据不是最新的。
所以必须使synchronized锁的范围大于事务控制的范围，把synchronized加到Controller层或者大于事务边界的调用层！

---

- 一个方法, 一个事务

```java
@Transactional(propagation = Propagation.REQUIRED)
    public void method1() {
        Employee employee = new Employee();
        employee.setLastName("pink floyd").setAge(23).setEmail("123!").setGender("1");
        save(employee);
        method2();
        System.out.println("方法调用");
        if (true){
            throw new RuntimeException("异常.........");
        }
        System.out.println("异常后");
    }
```

事务回滚method1()没有执行

---

- 两个方法, 一个事务

```java
@Transactional(propagation = Propagation.REQUIRED)
    public void method1() {
        Employee employee = new Employee();
        employee.setLastName("pink floyd").setAge(23).setEmail("123!").setGender("1");
        save(employee);
        method2();
        System.out.println("方法调用");
        if (true){
            throw new RuntimeException("异常.........");
        }
        System.out.println("异常后");
    }
    public void method2(){
        Employee employee = new Employee();
        employee.setLastName("oasis").setAge(12).setEmail("321!").setGender("1");
    }
```

method1()和method2()同一个事务回滚, method1()和method2()都没有执行

---

- 两个方法, 两个事务

```java
 @Transactional(propagation = Propagation.REQUIRED)
    public void method1() {
        Employee employee = new Employee();
        employee.setLastName("pink floyd").setAge(23).setEmail("123!").setGender("1");
        save(employee);
        method2();
        System.out.println("方法调用");
        if (true) {
            throw new RuntimeException("异常.........");
        }
        System.out.println("异常后");
    }

    @Transactional(propagation = Propagation.REQUIRES_NEW)
    public void method2() {
        Employee employee = new Employee();
        employee.setLastName("oasis").setAge(12).setEmail("321!").setGender("1");
        save(employee);
    }
```

method1(), method2() 都回滚

在默认的代理模式下，只有目标方法由外部调用，才能被 Spring 的事务拦截器拦截。在同一个类中的两个方法直接调用，是不会被 Spring 的事务拦截器拦截，就像上面的 method1 方法直接调用了同一个类中的 method2方法，method2 方法不会被 Spring 的事务拦截器拦截。==所以只有开启了一个事务==可以使用 AspectJ 取代 Spring AOP 代理来解决这个问题 , 或是另外写一个类将method2放在该类中

---

- 两个类, 两个方法, 两个事务

```java
@Service
public class TransactionService extends ServiceImpl<EmployeeMapper, Employee> implements IEmployeeService {
    @Autowired
    @Qualifier("transactionOtherService")
    TransactionOtherService otherService;

    @Transactional(propagation = Propagation.REQUIRED)
    public void method1() {
        Employee employee = new Employee();
        employee.setLastName("pink floyd").setAge(23).setEmail("123!").setGender("1");
        save(employee);
        otherService.method2();
        System.out.println("方法调用");
        if (true) {
            throw new RuntimeException("异常.........");
        }
        System.out.println("异常后");
    }
```

```java
@Service
public class TransactionOtherService extends ServiceImpl<EmployeeMapper, Employee> implements IEmployeeService {
    @Transactional(propagation = Propagation.REQUIRES_NEW)
    public void method2() {
        Employee employee = new Employee();
        employee.setLastName("oasis").setAge(12).setEmail("321!").setGender("1");
        save(employee);
    }
```

method1() 回滚, method2()成功提交

 因为Propagation.REQUIRES_NEW 会创建一个新的事务，如果当前存在事务，则把当前事务挂起

---

- 两个方法, 去掉method1的事务, 保留method2的事务

```java
//    @Transactional(propagation = Propagation.REQUIRED)
    public void method1() {
        Employee employee = new Employee();
        employee.setLastName("pink floyd").setAge(23).setEmail("123!").setGender("1");
        save(employee);
        method2();
        System.out.println("方法调用");
        if (true) {
            throw new RuntimeException("异常.........");
        }
        System.out.println("异常后");
    }

    @Transactional(propagation = Propagation.REQUIRES_NEW)
    public void method2() {
        Employee employee = new Employee();
        employee.setLastName("oasis").setAge(12).setEmail("321!").setGender("1");
        save(employee);
    }

```

两个方法都正常执行, 且method2()不会开启事务, 只是当作普通方法
