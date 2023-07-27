# Day01 - 位运算

```java
//convert decimal to 32 bit binary
public static void print(int num) {
  for (int i = 31; i >= 0; i--) {
    //和 32 位每一位都做与运算
    System.out.print((num & (1 << i)) == 0 ? "0" : "1");
  }
  System.out.println();
}
```

## & 与运算

相同为 1，不同为 0

```
print(8);
print(10);
print(8&10);
```

输出

```
00000000000000000000000000001000
00000000000000000000000000001010
00000000000000000000000000001000
```

## | 或运算

