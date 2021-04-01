# Actions

参考:

https://www.zsythink.net/archives/2046

https://help.aliyun.com/knowledge_detail/65077.html?spm=a2c4g.11186623.6.893.207318cdvZOd6Q

awk Actions 遵循如下几条规则

- action需要在`{ }`中可以有多个，==采用C#的编程方式(`{}`可以被省略)==，同样的也可以在BEGIN和END中使用

- `{}`中定义的变量没有类型，==如果变量没有定义，直接取空值==

- 可以定义数组方法与C#中相同，但是有一点不同的是如果值不存在默认取空

  ```
  [root@chz t]# awk ' BEGIN{arr[0]=a;arr[1]=b;arr[3]=c;for(i in arr){print i,arr[i]}}'
  0 
  1 
  3 
  ```

- 有多个action时，需要使用`;`分隔

  ```
  [root@chz opt]# df -hT | awk '{ \
  > print $1;
  > print $2;
  > '}
  Filesystem
  Type
  devtmpfs
  devtmpfs
  tmpfs
  tmpfs
  tmpfs
  tmpfs
  tmpfs
  tmpfs
  /dev/mapper/centos-root
  xfs
  /dev/sda1
  xfs
  tmpfs
  tmpfs
  /dev/sr0
  iso9660
  
  ```

## 例子

https://blog.51cto.com/lizhenliang/1764025

**求和**

https://blog.csdn.net/csCrazybing/article/details/52594989

```
[root@chz t]# seq 10  | awk '{sum+=$1}END{print sum}'
55
```

**组合**

```
[root@chz t]# df -hT | awk '{print $1}{print $2}'
Filesystem
Type
devtmpfs
devtmpfs
tmpfs
tmpfs
tmpfs
tmpfs
tmpfs
tmpfs
/dev/mapper/centos-root
xfs
/dev/sda1
xfs
tmpfs
tmpfs
/dev/sr0
iso9660
```

