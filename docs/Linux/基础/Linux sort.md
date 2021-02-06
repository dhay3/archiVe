# Linux sort

默认按照ascii表比较，`-f`将小写字母转为大写字母比较，`-r`逆序输出，`-b`忽略每行开头的空格，`-M`按照月份比较，`-o`将比较后的内容持久化。可以将管道符内的内容做为标准输入。

```
root in /opt λ cat test
c
d
a
f
g            
root in /opt λ sort test
a
c
d
f
g  
```

- `-n`

  按照数字的值比较，而不是ascii码

  ```
  root in /opt λ sort -n test
  2
  11
  33
  254                                                                                                                                                            /0.0s
  root in /opt λ cat test
  11
  2
  33
  254                   
  ```

  