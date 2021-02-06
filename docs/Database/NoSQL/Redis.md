# Redis

[TOC]

参考链接：

Http://redis.io/

http://www.redis.cn/

## 介绍

**Remote Dictionary Server(远程字典服务器)**

它可以用作数据库、缓存和消息中间件。 它支持多种类型的数据结构，如 字符串（strings）， 散列（hashes）， 列表（lists）， 集合（sets），有序集合（sorted sets)，并通过 Redis哨兵（Sentinel）和自动 分区（Cluster）提供高可用性（high availability）。

### 为什么使用Redis

- 性能

<img src="..\..\..\imgs\_redis\4.PNG"/>

- 并发

<img src="..\..\..\imgs\_redis\5.PNG"/>

## 安装

### 方法一/ wget

```shell
$ wget http://download.redis.io/releases/redis-5.0.8.tar.gz
$ tar xzf redis-5.0.8.tar.gz
$ cd redis-5.0.8
$ make  执行redis的Makefile文件
```

<img src="..\..\..\imgs\_redis\2.PNG"/>

- 如果出现如上所示错误,需要安装gcc(c语言编辑器)

  ```shell
  $ yum install gcc
  ```

<img src="..\..\..\imgs\_redis\3.PNG"/>

- 出现如上错误编译库出错, 将make 命令改为`make MALLOC=libc`或是`make distclean`

  ```shell
  $ src/redis-server 开启redis
  ```

### 方法二/ Docker

```shell
docker pull redis:version
```

```shell
docker run -d redis 运行redis
```

- `Docker`安装后的目录， 运行`Redis`最好也在该目录下

<img src="..\..\..\imgs\_redis\6.PNG"/>

## 开启Redis

###  服务端

- 前台启动`Redis`

  进入安装目录`redis-server`,默认端口6379

- 后台启动`Redis`, 需要修改配置文件

  复制源文件中的`redis.conf`,到自定义目录**修改redis 默认配置**, 通过修改过的`redis.conf`来启动Redis

  <img src="..\..\..\imgs\_redis\7.PNG" alt="7" style="zoom:60%;" />

### 客户端

- 开启客户端， `redis-cli `

- 测试

  ```shell
  redis-cli 
  redis> set foo bar
  OK
  redis> get foo
  "bar"
  ```

- 查看redis是否运行

  ```shell
  ps -ef|grep redis
  ```

- 关闭客户端

  ```shell
  redis-cli shutdown
  redis-cli -p 6379 shutdown 指定端口关闭
  ```

  `shutdown`退出关闭redis客户端，`exit`退出redis(==直接exit并不会退出redis==)

  <font style='color:red'>**注意！！！**在使用shutdown，redis会将缓存中的数据贯入硬盘中，生成dump.rdb所以会还原数据</font>

## 常识

- 默认16个数据库，类似数组下表从零开始，初始默认使用零号库，使用`SELECT <dbid>`命令在连接上指定数据库id
- `Dbsize`查看当前数据库的key的数量
- `Flushdb`清空当前库
- `Flushall`通杀全部库
- 统一密码管理，16个库都是同样密码，要么都OK要么一个也连接不上

## 五大数据结构

- **String**（字符串）

  一个key对应一个value。redis的`string`可以包含任何数据，比如jpg图片或者序列化的对象 。

- **Hash**(散列)

  hash是一个string类型的`field`和`value`的映射表，hash特别适合用于存储对象。类似Java里面的`Map<String,Object>`

- **Lists**（列表）

  Redis 列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素导列表的头部（左边）或者尾部（右边）。它的底层实际是个链表

-  **Set**（集合）

  Redis的Set是string类型的无序集合。它是通过HashTable实现实现的

- **Zset**（有序集合)

  Redis zset 和 set 一样也是string类型元素的集合,且不允许重复的成员。不同的是每个元素都会关联一个double类型的分数。redis正是通过分数来为集合中的成员进行从小到大的排序。

## 指令集

参考链接：

http://www.redis.cn/commands.html

http://doc.redisfans.com/

- key(键)
  - `keys *` 获取所有的key
  - `exsits key`
  - `move key db`
  - `expire key`
  - `ttl key` 永久有效返回-1，不存在或过期返回-2
  - `type key`

- String(字符串)

  - `set`,`get`
  - `del`
  - `incr` 自增1，`incrby`增加
  - `meset`，`mget`
  - `getrange`， `setrange`

  - `getset`

- List(列表)

  如果键不存在，创建新的链表；
  如果键已存在，新增内容；
  如果值全移除，对应的键也就消失了。

<img src="..\..\..\imgs\_redis\29.png"/>

  - `lpush`(从左边压入元素)，`rpush`(从右边压入元素), `lrange`（从左往右打印）

    `lrange 0 10`会打印11元素， ==左闭右闭==

    `lrange 0 -1` 这里的-1表示倒数第一个元素，同理 -2 ...

  - `lpop`(左边第一个元素)，`rpop`(右边第一个元素)

  - `lindex`(从左边开始是第一个元素， 下标从0开始)

  - `lrem`移除元素

- Set(集合)
  - `sadd`，`smember`显示散列元素
  - `scard`
  - `srem`

- Hash(散列)
  - `hset`，`hget`
  - `hmset`,`hmget`
- Zset(有序集合)
  
  - `zadd`,`zrange`

## redis.conf

==`vim`通过/来寻找文件中的内容==

- **NETWORK**

  - 端口

<img src="..\..\..\imgs\_redis\13.PNG"/>

  - 空闲多少时间后断开客户端连接，0表示关闭该功能

<img src="..\..\..\imgs\_redis\14.PNG"/>

  - 每隔多长时间检查一下连接是否还有效

<img src="..\..\..\imgs\_redis\15.PNG"/>

- **GENERAL**

  - 设置redis是否是以守护进程模式运行

<img src="..\..\..\imgs\_redis\11.PNG"/>

  - 护进程模式运行就会生成一个pid文件

<img src="..\..\..\imgs\_redis\10.PNG"/>

  - 设置日志文件的输出等级

<img src="..\..\..\imgs\_redis\12.PNG"/>

  - 设置日志文件名, ==默认不会生成日志文件==

<img src="..\..\..\imgs\_redis\28.PNG"/>

  - 默认关闭系统记录日志文件，如需开启配置如下

<img src="..\..\..\imgs\_redis\17.PNG"/>

  - 一共有16个数据库

<img src="..\..\..\imgs\_redis\16.PNG"/>

## RDB/ Redis DataBase

在指定的时间间隔内将内存中的数据集快照写入磁盘，也就是行话讲的Snapshot快照，它恢复时是将快照文件直接读到内存里。

redis默认开启RDB,当redis断开连接后,数据同样存在, ==但是如果flushdb 或是 flushall可以恢复==

**持久化可以在指定的时间间隔内, 生成数据集的时间点快照。**

**设置快照时间**,在客户端中通过**save** 命令可以达到立即保存的效果（adding a save directive）

<img src="..\..\..\imgs\_redis\21.PNG"/>

==RDB有可能会造成最后一次数据丢失，由于没有到达指定条件， 但是`Redis Serve`关闭==

**设置redis 数据存放位置，即dump.rdb文件**

==默认生成在开启服务器的位置==

<img src="..\..\..\imgs\_redis\20.PNG"/>

- MEMORY MANGEMENT

  内存满了之后redis的策略，默认noeviction（永不过期），生产环境中不能使用`noeviction`

<img src="..\..\..\imgs\_redis\19.PNG"/>

volatile-lru :用LRU算法移除key，只对设置了过期时间的键

allkeys-lru :用LRU算法移除key

volatile-random : 在过期的集合中随机移除key，只对设置了过期时间的键

allkeys-random : 移除随机的key

volatile-ttl: 移除那些TTL值最小的key，即那些将要过期的key

noeviction: 不进行移除 **<font style='color:red'>生产环境中不可采取</font>**

## AOF/ Append Only File

以日志的形式来记录每个写操作，将Redis执行过的所有写指令记录下来(读操作不记录)，只许追加文件但不可以改写文件，redis启动之初会读取该文件重新构建数据，换言之，redis重启的话就根据日志文件的内容将写指令从前到后执行一次以完成数据的恢复工作（==如果执行了flushall会记录flushall，所以还是空的==）

AOF默认关闭，需要设置成yes开启

<img src="..\..\..\imgs\_redis\22.PNG"/>

**<font style='color:red'>如果同时存在RDB和AOF会优先载入AOF</font>**

### 文件出错，恢复数据

如果appendonly.aof文件出错，通过**/usr/local/bin**下的**redis-check-aof**修复

如果dump.rdb文件出错，通过**/usr/local/bin**下的**redis-check-dump**修复

## 主从复制（Master/Slave）

==每次与master断开之后，都需要重新连接，除非你配置进redis.conf文件==

Master用于写,Slave用于读

slave会将Master中的内容复制过来

```shell
SLAEOF 主库IP 主库端口
```

<img src="..\..\..\imgs\_redis\23.PNG"/>

查看主从关系

```shell
查看主从关系
 info replication 
```

<img src="..\..\..\imgs\_redis\24.PNG"/>

#### 1.一主二仆

一个Master两个Slave

<img src="..\..\..\imgs\_redis\25.png"/>

#### 2.薪火相传

下一个salve是上一个的slave,同样的slave还是不能写

<img src="..\..\..\imgs\_redis\26.png"/>

#### 3.反客为主

变为一个单独master

```
SLAVE no one
```

## ==哨兵模式==

反客为主的自动版,能够后台监控主机是否故障,如果故障了根据投票数自动将从库转为主库

#### 1.新建一个sentinel.conf用于配置哨兵

```shell
touch sentinel.conf
```

#### 2.配置sentinel.conf

==这里被监视的必须是主库==

```shell
sentinel monitor 被监视的数据库名(可以自定义) 127.0.0.1 6379 1
最后的1表示主机宕机后,投票得1票即可成为主机

```

#### 3.开启哨兵

```shell
redis-sentinel /myRedis/sentinel.conf
```

==宕机后重新连回来就会变成slave==   以下是自动生成得sentinel.conf中的内容



<img src="..\..\..\imgs\_redis/27.PNG"/>

