[TOC]

learning ...

## #ChapterOne

**四种类型**:

- 接口(包括注释), 即`@Xxx`, 实际就是一个接口
- 类(包括enum)
- 数组
- 基本类型

前三种通常被称为引用类型(reference type)

类实例和数组是对象对象(object), 而基本类型的值不是对象

**类的成员**(memeber):

- 域(field)
- 方法(method)
- 成员类(memeber class)
- 成员接口(memeber interface)

**方法的签名**(signature): 由方法的名称和参数类型组成, 不包括方法的返回类型

## #ChapterTwo

### RuleOne

```java
package chapter01.rule01;

import lombok.Data;

/**
 * 静态工厂
 * @author 82341
 */
@Data
public class staticFactoryMethod {
    private String name;
    private Integer age;
    private staticFactoryMethod(){}

    public static staticFactoryMethod factoryMethod(){
        return new staticFactoryMethod();
    }

    public staticFactoryMethod name(String name){
        this.name = name;
        return this;
    }
    public staticFactoryMethod age(Integer age){
        this.age = age;
        return this;
    }

    public static void main(String[] args) {
        //流式API
        new staticFactoryMethod().name("chz").age(23);
    }
}

```

### RuleTwo

```java
package chapter01.rule02;

import lombok.Data;

@Data
public class Person {
    private String name;
    private Integer age;
    private Boolean gender;
    private Double salary;

    private Person(Builder builder) {
        name = builder.name;
        age = builder.age;
        gender = builder.gender;
        salary = builder.salary;
    }

    @Data
    public static class Builder {
        private String name;
        private Integer age;
        private Boolean gender;
        private Double salary;

        public Builder(String name, Integer age) {
            //这里的this指的时静态内部类
            this.name = name;
            this.age = age;
        }

        public Builder gender(Boolean gender) {
            this.gender = gender;
            return this;
        }

        public Builder salary(Double salary) {
            this.salary = salary;
            return this;
        }

        /**
         * 关键的一步, 将Builder获取到的值传给Person
         */
        public Person build() {
            return new Person(this);
        }
    }
    public static void main(String[] args) {
        Person person = new Builder("chz", 23)
                .gender(Boolean.TRUE).salary(200D)
                .build();
        System.out.println(person);
    }
}
```

### RuleThree

```java
package chapter01.rule03;


/**
 * 单例模式一
 */
class Elvis {
    //使用公开的类变量
    public static final Elvis elvis = new Elvis();

    private Elvis() {
        System.out.println("Elvis is unique");
    }

    public void sing() {
        System.out.println("sing out of tune");
    }
}

/**
 * 单例模式二
 */
class Bowie {
    private static final Bowie bowie = new Bowie();

    private Bowie() {
        System.out.println("rain man");
    }

    public static Bowie getInstance() {
        return bowie;
    }

    public void sing() {
        System.out.println("sing out of tune");
    }

}

/**
 * 单例模式三, 推荐
 * enum无偿提供序列化
 * 反编译的Enum对象自动实现Serializable
 */
public enum Teacher {
    INSTANCE("zs", 22);
    private final String name;
    private final Integer age;

    Teacher(String name, Integer age) {
        this.name = name;
        this.age = age;
    }

    public String getName() {
        return name;
    }

    public Integer getAge() {
        return age;
    }

    public void say() {
        System.out.println("hello world");
    }


    public static void main(String[] args) {
        Elvis.elvis.sing();
        Bowie.getInstance().sing();
        Teacher.INSTANCE.say();
    }
}

```

### #RuleFour

对于工具类, 因该要私有化构造器, 避免工具类实例化, 同时工具类将不能被实例化

### #RuleFive

当创建一个新的实例时, 就将该资源传入到构造器中. 这就是**依赖注入**(dependency injection)

As follows:

```java
class SpellChecker{
    private final Lexicon dictionary;
    public SpellChecker(Lexicon dictionary){
        this.dictionary = Objects.requireNonNull(dictionary);
    }
}
```

### RuleSix

- 对于创建成本高的对象, 要将对象缓存下来

  As follows:

```java
public class RomanNumberals {
    private static final Pattern ROMAN  =  Pattern.compile("");
    static boolean isRoman(String s){
        //创建matcher并判断是否匹配对应的正则
        return  ROMAN.matcher(s).matches();
    }
}
```

这里的`final`作用是将对象存入常量池中, 以免被`gabage collector`回收, 造成不必要的性能浪费

- 要优先使用基本数据类型, 而不是装箱基本类型, 当心无意识的自动装箱. 

  创建多余的对象会降低性能

  As follows:

```java
String s = new String("jetbrain");
Integer i = new Integer(23);
```

但是对于`domain object` 应该使用装箱的基本类型, 防止基本类型的初始化赋值

### RuleSeven

### RuleEight

正常情况下, 未被捕获的异常将会是线程终止, 并打印出`stack trace` , 但是, 如果异常在终结方法中, 则不会如此

(无法证实)

### RuleNine

关闭资源时优先考虑的时`try-with-resource`而不是`try-finally`

```java
package chapter01.rule09;

import java.io.BufferedInputStream;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;

/**
 * 关闭资源时优先考虑的时try-with-resource而不是try-finally
 */
public class TryWithResource {
    public static void main(String[] args) throws IOException {
        FileInputStream fis = null;
        BufferedInputStream bis = null;
        try {
            fis = new FileInputStream("/");
            bis = new BufferedInputStream(fis);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } finally {
            fis.close();
        }
    }
}

```

这里先关闭了 `FileInputStream`, 但是`BufferedInputStream`依赖`FileInputStream`, 如果这样做就会

报出`java.io.IOException`

而正常的关闭顺序是, 由上到下, 由外到里==如果关闭了外层的流, 就可以不用关闭内层的流==, `JAVA`会帮我们以此关闭流, 具体实现可以参考源码

但是使用`try-with-resoures`就可以自动关闭流

```j&amp;#39;a&amp;#39;v
public class TryWithResource {
    public static void main(String[] args) {
        try (FileInputStream fis = new FileInputStream("/");
             BufferedInputStream bis = new BufferedInputStream(fis);
             FileOutputStream fos = new FileOutputStream("")) {
        } catch (Exception e) {
        }
    }
}

```

多说一句, 如果使用`lombok`可以使用`@CleanUp`注解来关闭资源

```java
public class TryWithResource {
    public static void main(String[] args) throws FileNotFoundException {
        @Cleanup FileInputStream fis = new FileInputStream("./jpg");
        @Cleanup BufferedInputStream bis = new BufferedInputStream(fis);
        @Cleanup FileOutputStream fos = new FileOutputStream("./1.jpg");
    }
}

```

## #CharpterThree

### RuleTen

对于既不是`float`也不是`double`类型的基本数据类型, 可以使用`==`进行比较, 对于对象引用域, 可以使用

重写`equals`方法, 对于`float`, 可以使用静态`Float.compare(float,float)`方法, 同理`double`

**对`float`和`double`域进行特殊处理是有必要的**

使用`Float.equals`和`Double.equals`会对值自动装箱, 会导致性能下降

如果存在`null = = null return true;`的情况, 就需要采用`Objects.equals()`

### RuleEleven

- 如果两个对象根据`equals()`方法比较是相等的, 那么调用的这两个对象中的`hashCode`都必须产生相同的整数结果
- 如果两个对象根据`equals()`方法比较是不相等的, 那么调用这两个对象中的`hashCode`有可能相同

如果不重写`hashCode`将会违背, **相同的对象必须具有相同的hashcode**

## #ChapterFour

### RuleTweenty-Four

嵌套类(nested class)

<img src="D:\java资料\我的笔记\EffectiveJava\Snipaste_2020-07-29_13-29-29.png" style="zoom:80%;" />

是指定义在另一个类的内部类, 嵌套类存在的目的应该只为它的外围类(enclosing class)提供服务。

#### 1. 非静态成员类

非静态内部类的每个实例都隐含地与外围类的一个**实例相关联**， 由于需要先初始化外围类才能调用内嵌类， 所以开销比较大

- 由于没有`static` 修饰符，不允许定义类变量， 类方法
- 在实例方法中可以调用外部类的成员变量， 成员方法（包括类相关变量和方法）
- 如果存在同名变量或方法， 通过`enclosingClass.this.member`调用
- 在`enclosing class`中通过，`new` 关键字调用， 在其他类中通过`new enclosingClass.new nestedClass`调用

```java
public class EnclosingClass {
    private int num = 10;
    private static int age = 20;
    //age = age+10; 错误只能在方法体
    public void instanceMethod(){
        System.out.println("instanceMethod invoke");
        //通过new 关键字调用内嵌类
        new NestedClass().nestedMethod();
    }
    public static void staticMethod(){
        System.out.println("staticMethod invoke");
    }
    public class NestedClass{
        private int num = 20;
        //不允许含有类成员变量或类方法
        //private static int cnt = 30;
        public void nestedMethod(){
            //允许内部类实例方法调用外部类实例方法或类方法
            instanceMethod();
            staticMethod();
            //这里的this代表的就是内嵌类的实例
            System.out.println("this.num = "+this.num); //20
            System.out.println("num = "+num); // 20
            //通过 外围类.this.成员变量, 获取内嵌类的实例成员变量
            System.out.println("EnclosingClass.this.num = "+EnclosingClass.this.num); // 10
        }

    }

    public static void main(String[] args) {
        new EnclosingClass().new NestedClass().nestedMethod();
    }
```

#### 2. 静态成员类

==如果**嵌套类的实例**可以在外的外围类的实例之外独立存在，这个嵌套类就必须是静态成员类==， 也就是说嵌套类的实例方法不需要依靠外围类实例方法

- 允许定义类变量，类方法
- 不允许在方法内调用外围类实例方法
- 外围类通过`nestedClass.member`调用方法或是变量

```java
    public static class staticNestClass {
        private int num = 10;
        //可以定义类变量和类方法
        private static int cnt = 30;
        //可以定义实例方法,类方法
        public void nestedMethod() {
            //不允许调用外部类的实例方法
            //instanceMethod();
            staticMethod();
        }

        public static void nestedStaticMethod(){
            System.out.println("nestedStaticMethod");
        }
    }
```

