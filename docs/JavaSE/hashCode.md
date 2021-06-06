# hashCode

- hashcode能大大降低对象比较次数，提高查找效率！
- 如果两个对象相同，那么它们的hashCode值一定要相同
- 如果两个对象的hashCode相同，它们并不一定相同（这里说的对象相同指的是用eqauls方法比较）。 
- equals()相等的两个对象，hashcode()一定相等；equals()不相等的两个对象，却并不能证明他们的hashcode()不相等。 

==所以一般先通过判断hash值再通过equals(),判断对象内容是否相同,同时也可以优化内存结构==

```java
System.out.println(a1.hashCode()==a2.hashCode()||a1.equals(a2));
```

