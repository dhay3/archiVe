# Aliyun SLS

## 入门

[https://help.aliyun.com/document_detail/54604.html](https://help.aliyun.com/document_detail/54604.html)
[https://help.aliyun.com/document_detail/28967.html](https://help.aliyun.com/document_detail/28967.html)
## 字段索引
SLS中的字段必须配置索引，才可以被查询。分为以下2种索引

1. 全文索引

   在日志全文查询指定文本文本。例如`cyber`就会在日志中(不管是key还是value)查询cyber文本

2. 字段索引

   可以使用`key:value`的模式查询。如果没有配置字段索引，就不能筛选该字段。例如没有配置hostnat字段的索引，就不能通过`* | select hostname`查询hostname字段

## SQL 语法
在 SLS 中 SQL 分为两部分，两者以竖线（|）分隔，左边是查询语句，右边是分析语句。在使用查询和分析时，必须启用索引否则无法查询

1. 查询语句

   用于指定过滤规则，可以是关键词、数值、数值范围、空格、星号。如果为空格或星号（*），表示无过滤条件。可以理解成普通SQL中的where部分

2. 分析语句

   对查询结果进行计算和统计

### 查询语句
根据配置的索引方式不同，查询语句分为全文查询和字段查询。

1. 全文查询，使用全文检索，根据配置的分词符将日志拆分成多个字段，全日志检索
1. 字段查询，可以使用`key:value`或`key=value`的方式来查询

查询语句还需要注意如下几点

1. 查询语句可以单独使用，但是分析语句必须与查询语句一起使用
1. 在查询语句中`:`等价于`=`
### 分析语句
分析语句，结合了查询（过滤显示指定字段）和SQL计算（聚合函数）的功能。同时有以下几点需要注意

1. 查询语句可以单独使用，但是分析语句必须与查询语句一起使用
1. 分析语句中不需要填写from和where，同时分析语句不区分大小写
1. 在分析语句中，表示字符串的字符必须使用单引号（''）包裹，无符号包裹或被双引号（""）包裹的字符表示字段名或列名。例如：'status'表示字符串status，status或"status"表示日志字段status。
1. 分析默认最多返回100行数据，如果需要返回更多需要使用LIMIT语法
#### 分析语句常用函数

1. IP

   [https://help.aliyun.com/document_detail/63458.html](https://help.aliyun.com/document_detail/63458.html)

2. URL

   [https://help.aliyun.com/document_detail/63452.html](https://help.aliyun.com/document_detail/63452.html)

3. 时间和日期

   [https://help.aliyun.com/document_detail/63451.html](https://help.aliyun.com/document_detail/63451.html#section-uae-8wz-765)

4. 类型转换函数

   [https://help.aliyun.com/document_detail/63456.html](https://help.aliyun.com/document_detail/63456.html?spm=a2c4g.11186623.0.0.35a0753emiKU4g)

#### 模糊查询
模糊查询分为两种，非精确模糊查询和精确模糊查询
##### 非精确模糊查询
出现在查询语句部分，使用`*`和`？`(含义和普通正则一样)，但是有如下几点限制

1. 只会查询100条记录

1. `*`和`？`只能在词的中间或末尾，不能在词的开头

   ```
   status='4*' | select *
   ```
##### 精确模糊查询
出现在分析语句部分，使用 like 或者 正则匹配
[https://help.aliyun.com/document_detail/167995.htm?spm=a2c4g.11186623.0.0.2b034b2daz7wcX#concept-2495512](https://help.aliyun.com/document_detail/167995.htm?spm=a2c4g.11186623.0.0.2b034b2daz7wcX#concept-2495512)
```
#分析语句中的字符串必须要在单引号内
* | select status like '%4%'
```
#### 注意
其实前半部分可以完全忽略，直接按照SQL的语法处理
```
* |
select
  source_ip,
  destination_ip,
  destination_port,
  nat_source_ip,
  hostname
FROM  (
    select
      source_ip,
      destination_ip,
      destination_port,
      hostname,
      case
        when is_prefix_subnet_of('203.119.0.0/16', concat(nat_source_ip, '/16')) then nat_source_ip
      end as nat_source_ip
    FROM      log
  )
where
  nat_source_ip is not null
  and (
    destination_ip = '8.141.190.251'
    or destination_ip = '8.141.190.252'
  )
limit
  1000000
```
### 例子
```
* |
select
  src_idc_abbreviation as s_idc,
  src_app_group_name as s_app_group,
  concat(src_addr, ':', cast(src_port as varchar)) as s_tuple,
  dst_idc_abbreviation as d_idc,
  dst_app_group_name as d_app_group,
  concat(dst_addr, ':', cast(dst_port as varchar)) as d_tuple,
  state,
  label as retrans,
  from_unixtime(ts) as time
where
  src_addr = '33.52.118.170'
order by
  ts
limit
  1000000
```
```
* |
select
  concat(dst_addr, ':', cast(dst_port as varchar)) as d_tuple,
  count(*) as cnt
where
  src_addr = '33.52.118.170'
group by
  dst_addr,
  dst_port
limit
  100000
```
