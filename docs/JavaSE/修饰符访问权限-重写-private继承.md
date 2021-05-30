# Java修饰符访问权限/重写/private继承

[TOC]

### 修饰符权限

| 修饰符        | 所在类 | 同一个包内其他类 | 其他包内子类 | 其他包内非子类 |
| ------------- | ------ | ---------------- | ------------ | -------------- |
| private       | √      | ×                | ×            | ×              |
| 缺省(default) | √      | √                | ×            | ×              |
| protected     | √      | √                | √            | ×              |
| public        | √      | √                | √            | √              |

### private

- 修饰类: 只能修饰内部类, 表示该类只能在当前类访问
- 修饰方法: 表示该方法只能在当前类访问

依次的权限修饰符从小到大 private --> default --> protected --> public

Members of a class that are declared private are not inherited by subclasses of that class.
Only members of a class that are declared protected or public are inherited by subclasses declared in a package other than the one in which the class is declared.
摘自:https://docs.oracle.com/javase/specs/jls/se7/html/jls-8.html#jls-8.2

1. 所以private属性不能被继承,但是子类可以通过public方法调用, 是因为子类继承时就和构造器一样super(), 而子类实际是通过super关键字访问父类的属性

2. 子类未指定调用父类的某个构造器,隐式调用父类的无参构造器构造器

3. 静态方法也会被继承,但不能被重写, 同样可以通过子类名调用

> 注意构造方法是不能被继承的, 子类隐式调用父类的构造器

###  重写

子类重写父类注意事项：遵循：”两同、两小、一大“原则

1、两同：方法名相同，形参列表相同

2、两小：子类方法返回值类型比父类方法返回值类型更小或相等，子类方法声明抛出的异常类比父类方法声明抛出的异常类更小或相等

3、一大：子类方法访问权限比父类方法访问权限更大或相等

4、特别注意的是：覆盖的方法和被覆盖的方法，要么都是类方法，要么都是实例方法，不能一个是类方法，一个是实例方法，否则会发生编译错误

 
