java -D 

参考:

https://blog.csdn.net/fenglongmiao/article/details/80511512

java程序启动参数 -D是用来做什么的呢？去查询了一下官方解释：

*Set a system property value. If value is a string that contains spaces, you must enclose the string in double quotes:*

```hljs
java -Dfoo="some string" SomeClass
```

也就是说-D是用来在启动一个java程序时设置系统属性值的。如果该值是一个字符串且包含空格，那么需要包在一对双引号中。

何为系统属性值呢？也就是在System类中通过getProperties()得到的一串系统属性。

下面我们来写个测试方法就知道了！

```java
public class SystemProperty {
    public static void main(String[] args){
        System.out.print(System.getProperty("dubbo.token"));
    }
}
```

在运行改程序时加上JVM参数-Ddubbo.token="666" 或者 -Ddubbo.token=666，那么运行之后你可以看到控制台输出了666！

一点值得注意的是，需要设置的是JVM参数而不是program参数，注意看下图
