# nextInt()和nextLine()注意点

```java
        Scanner scanner = new Scanner(System.in);
        System.out.println("before");
        int i = scanner.nextInt();
        String s1 = scanner.nextLine(); //这里会吃换行符
        String s2 = scanner.nextLine();
        System.out.println("after");
```

`nextLine()`会读取\n, 如果要输出,要再加一个`nextLine()`