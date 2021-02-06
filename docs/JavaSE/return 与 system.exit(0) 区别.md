# return 与 system.exit(0) 区别

```java
public class ReturnTest {
    public static void main(String[] args) {
            f1(1);
            f2();
    }
    public static void f1(int i){
        if (1==i){
          System.out.println("不能为1");
           // System.exit(0);
            return;
        }
        System.out.println(i);
    }
    public static void f2(){
        System.out.println("调用方法2");
    }
}
```

使用return;

结果:

```
不能为1
调用方法2
```

return 后的本方法的代码不会执行, 但是其他的方法会执行

---

使用System.exit(0);

结果

```
Process finished with exit code 0
```

退出jvm, 之后的所有代码将不会被执行