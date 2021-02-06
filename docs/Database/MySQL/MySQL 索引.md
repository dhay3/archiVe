# MySQL 索引

参考：

https://juejin.cn/post/6844903909899632654?hmsr=coffeephp.com&utm_medium=coffeephp.com&utm_medium=coffeephp.com&utm_source=coffeephp.com%3Fhmsr%3Dcoffeephp.com&utm_source=coffeephp.com

https://zhuanlan.zhihu.com/p/29118331

https://www.runoob.com/mysql/mysql-index.html

## 概述

> 索引是存储引擎用于快速找到记录的一种数据结构.

这是MySQL官方对于索引的定义,可以看到索引是一种数据结构,那么我们应该怎样理解索引呢?一个常见的例子就是书的目录.我们都已经养成了看目录的习惯,拿到一本书时,我们首先会先去查看他的目录,并且当我们要查找某个内容时,我们会在目录中查找,然后找到该片段对应的页码,再根据相应的页码去书中查找.如果没有索引(目录)的话,我们就只能一页一页的去查找了.

在MySQL中,假设我们有一张如下记录的表:

| id   | name      | age  |
| ---- | --------- | ---- |
| 1    | huyan     | 10   |
| 2    | huiui     | 18   |
| 3    | lumingfei | 20   |
| 4    | chuzihang | 15   |
| 5    | nono      | 21   |

如果我们希望查找到年龄为15的人的名字,在没有索引的情况下我们只能遍历所有的数据去做逐一的对比,那么时间复杂度是O(n).

而如果我们在插入数据的过程中, 额外维护一个数组,将age字段有序的存储.得到如下数组.

```
[10,15,18,20,21]
 |  |  |  |  |
[x1,x4,x2,x3,x5]
复制代码
```

下面的x是模拟数据再磁盘上的存储位置.这个时候如果我们需要查找15岁的人的名字.我们可以对盖数组进行二分查找.众所周知,二分查找的时间复杂度为O(logn).查找到之后再根据具体的位置去获取真正的数据.

PS:MySQL中的索引不是使用的数组,而是使用的B+树(后面讲),这里用数组举例只是因为比较好理解.

## 优点VS缺点

**优点**

如上面所说,索引能帮助我们快速的查找到数据.其次因为索引中的值是顺序储存,那么可以帮助我们进行orderby操作.而且索引中也是存储了真正的值的,因此有一些的查询直接可以在索引中完成(也就是覆盖索引的概念,后面会提到).

- 减少查询需要扫描的数据量(加快了查询速度)
- 减少服务器的排序操作和创建临时表的操作(加快了groupby和orderby等操作)
- 将服务器的随机IO变为顺序IO(加快查询速度).

**缺点**

首先索引也是数据,也需要存储,因此会带来额外的存储空间占用.其次,在插入,更新和删除操作的同时,需要维护索引,因此会带来额外的时间开销.

- 索引占用磁盘或者内存空间
- 减慢了插入更新操作的速度

## 索引存储数据结构

MySQL默认使用InnoDB做为存储引擎我们可以通过`show engines`来查看

```
mysql> show engines;
+--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
| Engine             | Support | Comment                                                        | Transactions | XA   | Savepoints |
+--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
| MEMORY             | YES     | Hash based, stored in memory, useful for temporary tables      | NO           | NO   | NO         |
| MRG_MYISAM         | YES     | Collection of identical MyISAM tables                          | NO           | NO   | NO         |
| CSV                | YES     | CSV storage engine                                             | NO           | NO   | NO         |
| FEDERATED          | NO      | Federated MySQL storage engine                                 | NULL         | NULL | NULL       |
| PERFORMANCE_SCHEMA | YES     | Performance Schema                                             | NO           | NO   | NO         |
| MyISAM             | YES     | MyISAM storage engine                                          | NO           | NO   | NO         |
| InnoDB             | DEFAULT | Supports transactions, row-level locking, and foreign keys     | YES          | YES  | YES        |
| BLACKHOLE          | YES     | /dev/null storage engine (anything you write to it disappears) | NO           | NO   | NO         |
| ARCHIVE            | YES     | Archive storage engine                                         | NO           | NO   | NO         |
+--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
```

从上可看除 MySQL 当前默认的存储引擎是InnoDB,并且在5.7版本所有的存储引擎中只有 InnoDB 是事务性存储引擎，也就是说只有 InnoDB 支持事务。==而InnoDB采用B+Tree做为存储数据结构==

## 索引类型

可以使用`SHOW INDEX FROM table_name;`查看索引详情

<img src="http://songwenjie.vip/blog/180802/Eif26fJiEc.png?imageslim"/>

1. **主键索引 PRIMARY KEY**

   它是一种特殊的唯一索引，不允许有空值。一般是在建表的时候同时创建主键索引。

   注意：一个表只能有一个主键

<img src="http://songwenjie.vip/blog/180802/1c7D2F0f76.png?imageslim"/>

2. **唯一索引 UNIQUE**

   唯一索引列的值必须唯一，但允许有空值。如果是组合索引，则列值的组合必须唯一。

   可以通过`ALTER TABLE table_name ADD UNIQUE (column);`创建唯一索引

<img src="http://songwenjie.vip/blog/180802/DBdFeKE8Fk.png?imageslim"/>

<img src="http://songwenjie.vip/blog/180802/L2jl91b6J6.png?imageslim"/>

   可以通过`ALTER TABLE table_name ADD UNIQUE (column1,column2);`创建唯一组合索引

<img src="http://songwenjie.vip/blog/180802/mihd7Hm5i6.png?imageslim"/>

<img src="http://songwenjie.vip/blog/180802/bJbdFA9AcL.png?imageslim"/>

3. **普通索引 INDEX**

   最基本的索引，它没有任何限制。

   可以通过`ALTER TABLE table_name ADD INDEX index_name (column);`创建普通索引

<img src="http://songwenjie.vip/blog/180802/17CmJIIJhD.png?imageslim"/>

<img src="http://songwenjie.vip/blog/180802/4fA7L6kBBm.png?imageslim"/>

4. **组合索引 INDEX**

   组合索引，即一个索引包含多个列。多用于避免回表查询。

   可以通过`ALTER TABLE table_name ADD INDEX index_name(column1, column2, column3);`创建组合索引

<img src="http://songwenjie.vip/blog/180802/CLGIKiAC6J.png?imageslim"/>

<img src="http://songwenjie.vip/blog/180802/295B9bGi67.png?imageslim"/>

5. **全文索引 FULLTEXT**

   全文索引（也称全文检索）是目前搜索引擎使用的一种关键技术。

   可以通过`ALTER TABLE table_name ADD FULLTEXT (column);`创建全文索引

<img src="http://songwenjie.vip/blog/180802/AjfLLkhdH1.png?imageslim"/>

<img src="http://songwenjie.vip/blog/180802/bA1a1m49cL.png?imageslim"/>

