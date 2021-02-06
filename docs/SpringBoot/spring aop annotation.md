# spring aop annotation

### #被代理类

```java
@Component
public class TargetObject {
    public void targetMethod(){
        System.out.println("target method invoke ...");
    }
}
```

### #aspect

可以将pointcut 和 advice 都定义在aspect中

signature就是pointcut的别名

```java
@Aspect//标明该类是一个切面
@Component
public class MyAspect {
    @Pointcut("execution(* com.chz.target.TargetObject.*())")
    //signature
    public void pointcutMethod() {
    }
   @Before("pointcutMethod()")
    public void advice() {
        System.out.println("advice invoke ...");
    }
}
```

---

也可以将pointcut和advice 分离

```java
public class MyPointcut {
    //@Pointcut所在的类无需注入ioc,advice调用该pointcut的signature
    @Pointcut("execution(* com.chz.target.TargetObject.*())")
    public void pointcutMethod() {
    }
}
```

```java
@Aspect
@Component
public class MyAspect {
    @Before("com.chz.pointcut.MyPointcut.pointcutMethod()")
    public void advice() {
        System.out.println("advice invoke ...");
    }
}
```

---

上面的效果等效于

```java
@Aspect
@Component
public class MyAspect {
    @Before("execution(* com.chz.target.TargetObject.*())")
    public void advice() {
        System.out.println("advice invoke ...");
    }
```

---

### #@AfterReturning

可以通过returning属性拿到被代理对象的方法的返回值

```java

    @AfterReturning(pointcut = "com.chz.pointcut.MyPointcut.pointcutMethod()",
    returning = "retVal")
    public void afterReturningMethod(Object retVal) {
        System.out.println(retVal);
        System.out.println("advice invoke ...");
    }
```

### #@AfterThrowing

可以通过throwing属性拿到被代理对象的方法的异常, 

可以指定Exception的类型,当指定类型的异常发生后, 将调用@AfterTrowing所注解的方法

代理的对象的方法, try-catch 异常后将不会调用@AfterThrowing, 如果是抛出异常将会调用

==如果unchecked exception 被捕捉那么@AfterTrowing所注解的方法不会被调用==

```java
    @AfterThrowing(pointcut ="com.chz.pointcut.MyPointcut.pointcutMethod()",
    throwing = "e")
    public void exceptionMethod(Exception e){
        System.out.println(e.getMessage());
        System.out.println("afterThrowing invoke ...");
    }
```

### #@After

相当于finally 在 @AfterReturning之前执行

```java
    @After("com.chz.pointcut.MyPointcut.pointcutMethod()")
    public void AfterMethod(){
        System.out.println("after invoke ...");
    }
```

### #@Around

```java
    @Around("com.chz.pointcut.MyPointcut.pointcutMethod()")
    public Object aroundMethod(ProceedingJoinPoint joinPoint){
        Object proceed =null;
        try {
            //在proceed之前的部分就相当于@Before修饰的方法
            System.out.println("前置通知");
             proceed = joinPoint.proceed();//执行被代理对象的方法,并拿到返回值的返回值
            System.out.println("返回值" + proceed);
            System.out.println("函数名" + joinPoint.getSignature().getName());
            System.out.println("参数个数" + joinPoint.getArgs().length);
            System.out.println("被代理对象"+joinPoint.getTarget());
            System.out.println("被代理对象的全路径名"+joinPoint.getSignature().getDeclaringTypeName());
            //在proceed之后的部分就相当于@AfterReturning修饰的方法
            System.out.println("后置通知");
        } catch (Throwable throwable) {
            throwable.printStackTrace();
            //@AfterThrowing修饰的方法
            System.out.println("异常通知");
        }finally {
            //@After
            System.out.println("最终通知");
        }
        //可以在通知中被代理对象的方法的返回值,要求返回值类型与被代理对象的方法的类型一致
        return proceed;
    }
```

@Before, @After, @AfterReturning 所在的方法同样有一个和`ProceedingJointPoint`类似的属性`JoinPoint`

```java
@Before("execution(* com.chz.target.TargetObject.*())")
//aspectj.lang 下的JoinPoint
public void advice(JoinPoint joinPoint) {
    System.out.println("advice invoke ...");
}
```
