# Spring Aop

### #概念

> AOP(Aspect Orient Programming) 

面向切面编程，是面向对象编程（OOP）的一种补充。

面向对象编程将程序抽象成各个层次的对象，而面向切面编程是将程序抽象成各个切面。

### #术语

- Aspect (切面)

   通常是一个类, 在Aspect中包含Advice和PointCut

- JoinPoint (连接点)

  表示程序中明确定义的点, 典型的包括方法的调用, 对类成员的访问等

- PointCut (切点)

   定义了Advice将要发生的位置, 可以由多个JoinPoint或是单个JoinPoint组成

- Advice (通知/增强)

  定义了PointCut在特定的时机(before, afterReturning, after, afterThrowing, around)具体要做什么

- Target (目标对象)

  被织入Advice的对象

- Weaving (织入)

  关联Advice,PointCut, 具体的时间

- Target (目标对象)

  织入Advice的目标对象, 即被代理的对象

---

<img src="D:\asset\note\imgs\_SpringBoot\29.png"/>

---



### #常用JoinPoint 指示符

切入点指示符用来指示切入点表达式目的

- execution

  用于匹配方法的执行

  `execution( * test.*.mainFunction(..)`

- within

  用于匹配指定类型中的方法执行

  `within(com.xyz.someapp.web..*)`

- args

  用于匹配参数为指定类习的方法的执行

  `args(java.io.Serializable)`

- @annotation

  用于匹配持有指定注解的方法的执行

  `@annotation(com.chz.annotaion.Log)`

- @target

  

### #PointCut

springboot中开启aop无需`@EnableAspectJAutoProxy`

而springMVC中开启aop需要该注解, 或是在xml中配置`<aop:aspectj-autoproxy/>`

>  A pointcut declaration has two parts: a signature comprising a name and any parameters and a pointcut expression that determines exactly which method executions we are interested in

```java
@Pointcut("execution(* transfer(..))") // the pointcut expression
private void anyOldTransfer() {} // the pointcut signature
```

### #PointCut Expressions

​      可以使用 `&&,` `||` and `!` 来组合expressions

- `execution(public * *(..))`

  任意包,修饰符为public的方法, 返回类型任意, 方法名任意, 参数任意

- `execution( * test.*.mainFunction(..))`

  第一个* 表示返回类型为任意

  第二个* 表示test下的任意类

  参数中的 .. 表示任意参数类型

- `execution(* com.xyz.someapp.dao.*.*(..))`

  dao包下所有类的所有方法

- `execution(* test..*.mainFunctionException(..))`

  全路径名中 .. 表示当前包(test)及其子包

- `execution(* *To(..))`

  所有以To结尾的方法

- `execution(* set*(..))`

  所有以set开头的方法

- ` within(com.xyz.service.*)`

  `com.xyz.service`包下所有类的所有方法

- `args(java.io.Serializable)`

  参数为Serializable类型的所有方法

- `  @target(org.springframework.transaction.annotation.Transactional)`

  匹配代理对象被@Transactional 注解的所有方法
