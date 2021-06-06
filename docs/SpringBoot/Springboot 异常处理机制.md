# Springboot 异常处理机制

在接收一个异常后, springboot会发送请求到默认的 BasicErrorController 

<img src="..\..\..\..\java资料\我的笔记\springboot\img\10.PNG"/>

这是里面默认的两个映射

<img src="..\..\..\..\java资料\我的笔记\springboot\img\11.png"/>

第一个用于接收请求中accept包含 text/html的请求, 第二个用于接收没有text/html的请求

<img src="..\..\..\..\java资料\我的笔记\springboot\img\12.png"/>

然后通过`java.servlet.error.status_code`这个key获取到对应的错误状态码, 这里看到attributes对应的是一个`ConcurrentHashMap<>()`

<img src="D:\java资料\我的笔记\springboot\img\30.png" style="zoom:60%;" />

通过`getstatus()`中的`Httpstatus.valueof()`将对应的值转为`HttpStatus`枚举类

<img src="D:\java资料\我的笔记\springboot\img\31.png" style="zoom:80%;" />

**所以用于捕捉异常的方法不能设置为枚举类**, 就会发出`java.lang.ClassCastException: org.springframework.http.HttpStatus cannot be cast to java.lang.Integer`

但是可以设置为如下格式, 或者直接给出对应的值

<img src="..\..\..\..\java资料\我的笔记\springboot\img\32.png"/>

接下来获取错误的信息

<img src="..\..\..\..\java资料\我的笔记\springboot\img\13.PNG"/>

通过getErrorAttributes 获取到错误信息

<img src="..\..\..\..\java资料\我的笔记\springboot\img\14.PNG"/>

这里可以看到, 放入了什么信息, 有日期, 状态码, 异常的信息,根据请求域的不同获取请求的servletPath

将这些信息放入`Map<String, Object> model`中

<img src="..\..\..\..\java资料\我的笔记\springboot\img\15.PNG"/>

然后通过resolverErrorView解析视图, 可以看到调用了resolve方法

<img src="..\..\..\..\java资料\我的笔记\springboot\img\17.PNG"/>

这里可以清楚的看到通过thymeleaf模板跳转到了 error/viewName, 也就是templates/error/viewName.html

然后将map和view一起返回

getProvider()会检查视图是不是在templates文件下, 如果不在就访问静态资源下的 static/error/viewName.html ,如果静态资源下没有, 就返回null

<img src="..\..\..\..\java资料\我的笔记\springboot\img\18.PNG"/>

接下来就是一种特殊情况, templates和静态资源路径下找不到对应的视图, 就会进入resolve()

<img src="..\..\..\..\java资料\我的笔记\springboot\img\19.PNG"/>

调用get方法

<img src="..\..\..\..\java资料\我的笔记\springboot\img\20.PNG"/>

一路追踪

<img src="..\..\..\..\java资料\我的笔记\springboot\img\21.PNG"/>

发现`org.springframework.http.HttpStatus$Series`是一个枚举类

<img src="D:\java资料\我的笔记\springboot\img\22.PNG" style="zoom:60%;" />

在通过算法, 得出statutsCode = 400 可以对应 4xx的viewName , 同理 5xx

<img src="..\..\..\..\java资料\我的笔记\springboot\img\23.PNG"/>

如果 4xx 和 5xx 也没有

就通过`return (modelAndView != null) ? modelAndView : new ModelAndView("error", model);`

来判断 ,如果modelAndView就带着model跳转到内置的`Whitelabel Error Pag`e(这里并不是很清楚, 具体的实现)

---

这里用postman 模拟一下另一个映射

<img src="D:\java资料\我的笔记\springboot\img\24.PNG" style="zoom:80%;" />

<img src="..\..\..\..\java资料\我的笔记\springboot\img\25.PNG"/>

会返回json数据类型的到前台
