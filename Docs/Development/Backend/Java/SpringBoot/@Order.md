# @Order

@Order用于指定Bean的执行顺序(==也会影响加载顺序==), 数字越小优先级越高, 默认`Ordered.LOWEST_PRECEDENCE`

假设有两全局异常处理类处理相同的错误, 谁的优先级高会先调用谁

### handler1

```java
@Slf4j
@Order(Ordered.HIGHEST_PRECEDENCE)
//@Order(Ordered.LOWEST_PRECEDENCE)
@RestControllerAdvice
public class Handler1 {
    public Handler1() throws InterruptedException {
        TimeUnit.SECONDS.sleep(2);
        System.out.println("handler1 注入 ioc ....");
    }
    @ExceptionHandler(NullPointerException.class)
    private String catche(){
        log.warn("handler1 调用");
        return "handler1";
    }
}
```

### handler2

```java
@Slf4j
@Order(Ordered.LOWEST_PRECEDENCE)
//@Order(Ordered.HIGHEST_PRECEDENCE)
@RestControllerAdvice
public class Handler2 {
    public Handler2() {
        System.out.println("handler2 注入 ioc ....");
    }

    @ExceptionHandler(NullPointerException.class)
    private String catche(){
        log.warn("handler2 调用");
        return "handler2";
    }
}
```

### controller

```java
@RestController
public class Controller {
    @GetMapping("/get")
    @ResponseBody
    public void get() {
        throw new NullPointerException();
    }
}

```

先让`Handler1` 注入时休眠`2s`, 此时`Handler1`的`@Order(Ordered.HIGHEST_PRECEDENCE)` 是最高等级

运行发现

<img src="..\..\..\..\java资料\我的笔记\springboot\img\33.png"/>

==说明@order决定pojo注入的ioc的先后顺序==

我们发送请求看一看

<img src="D:\java资料\我的笔记\springboot\img\35.png" style="zoom:200%;" />

然后我们交换优先级, 发送`get`请求

<img src="D:\java资料\我的笔记\springboot\img\34.png" style="zoom:200%;" />

从结果可以明显看出优先级高的异常处理器被调用了

说明@order也决定执行顺序
