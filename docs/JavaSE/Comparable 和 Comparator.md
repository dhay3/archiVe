# Comparable 和 Comparator

==参数和this比较只有 return -1 时才会移动元素, 0和1都不会移动元素==

调用compareTo方法的是第二个元素, 参数是前一个元素 , **即this 是第二个元素**

##### 从小到大,1和0不动,-1往前移动

* this大于obj,返回正整数
* this小于obj,返回负整数
* this等于obj,返回零

##### 从大到小

* this大于obj,返回负整数
* this小于obj,返回正整数
* this等于obj ,返回零

ComparableTo() 从小到大

-ComparableTo() 从大到小

##### #Comparable 是内部比较器

```java
public int compareTo(Person o) {
```

- 从小到小

```java
if (o.age>this.age){
    return -1; 
}else{
    return 1;
}
```

- 从大到小

```java
if (this.age>o.age){
    return -1; 
}else{
    return 1;
}
```

##### #Comparator 是外部比较器

```java
 String[] a = new String[]{"A", "C", "F", "B"};
        Arrays.sort(a, String::compareTo);
```

这里是Comparator的compare方法指向String的compareTo方法, 即compare方法的具体实现是compareTo

```java
public int compare(Student s1, Student s2) {
				int flag;
				// 首选按年龄升序排序
				flag = s1.getAge()-s2.getAge();
				if(flag==0){
					// 再按学号升序排序
					flag = s1.getNum()-s2.getNum();
				}
				return flag;
			}
```

s1代表的就是**this**(从第二个开始),  s2代表的就是this的前一个元素, 如果小于零移动,
