# Mysql内置函数

转自：

https://blog.csdn.net/hellokandy/article/details/82964077

## 系统信息

| 函数                                                  | 作用                             |
| ----------------------------------------------------- | -------------------------------- |
| VERSION()                                             | 查询当前mysql的版本              |
| CONNECTION_ID()                                       | 返回服务器的连接数               |
| DATABASE()                                            | 返回当前数据库名                 |
| USER()、SYSTEM_USER()、SESSION_USER()、CURRENT_USER() | 返回当前用户                     |
| LAST_INSERT_ID()                                      | 返回最近生成的AUTO_INCREEAMENT值 |

## 字符串

| 函数                           | 说明                                                         |
| ------------------------------ | ------------------------------------------------------------ |
| CHAR_LENGTH(s)                 | 返回字符串s的字符数SELECT CHAR_LENGTH('你好123') -- 5        |
| LENGTH(s)                      | 返回字符串s的长度SELECT LENGTH('你好123') -- 9               |
| CONCAT(s1,s2,...)              | 将字符串s1,s2等多个字符串合并为一个字符串SELECT CONCAT('12','34') -- 1234 |
| CONCAT_WS(x,s1,s2,...)         | 同CONCAT(s1,s2,...)函数，但是每个字符串直接要加上xSELECT CONCAT_WS('@','12','34') -- 12@34 |
| INSERT(s1,x,len,s2)            | 将字符串s2替换s1的x位置开始长度为len的字符串SELECT INSERT('12345',1,3,'abc') -- abc45 |
| UPPER(s),UCAASE(S)             | 将字符串s的所有字母变成大写字母SELECT UPPER('abc') -- ABC    |
| LOWER(s),LCASE(s)              | 将字符串s的所有字母变成小写字母SELECT LOWER('ABC') -- abc    |
| LEFT(s,n)                      | 返回字符串s的前n个字符SELECT LEFT('abcde',2) -- ab           |
| RIGHT(s,n)                     | 返回字符串s的后n个字符SELECT RIGHT('abcde',2) -- de          |
| LPAD(s1,len,s2)                | 字符串s2来填充s1的开始处，使字符串长度达到lenSELECT LPAD('abc',5,'xx') -- xxabc |
| RPAD(s1,len,s2)                | 字符串s2来填充s1的结尾处，使字符串的长度达到lenSELECT RPAD('abc',5,'xx') -- abcxx |
| LTRIM(s)                       | 去掉字符串s开始处的空格                                      |
| RTRIM(s)                       | 去掉字符串s结尾处的空格                                      |
| TRIM(s)                        | 去掉字符串s开始和结尾处的空格                                |
| TRIM(s1 FROM s)                | 去掉字符串s中开始处和结尾处的字符串s1SELECT TRIM('@' FROM '@@abc@@') -- abc |
| REPEAT(s,n)                    | 将字符串s重复n次SELECT REPEAT('ab',3) -- ababab              |
| SPACE(n)                       | 返回n个空格                                                  |
| REPLACE(s,s1,s2)               | 将字符串s2替代字符串s中的字符串s1SELECT REPLACE('abc','a','x') --xbc |
| STRCMP(s1,s2)                  | 比较字符串s1和s2                                             |
| SUBSTRING(s,n,len)             | 获取从字符串s中的第n个位置开始长度为len的字符串              |
| MID(s,n,len)                   | 同SUBSTRING(s,n,len)                                         |
| LOCATE(s1,s),POSITION(s1 IN s) | 从字符串s中获取s1的开始位置SELECT LOCATE('b', 'abc') -- 2    |
| INSTR(s,s1)                    | 从字符串s中获取s1的开始位置SELECT INSTR('abc','b') -- 2      |
| REVERSE(s)                     | 将字符串s的顺序反过来SELECT REVERSE('abc') -- cba            |
| ELT(n,s1,s2,...)               | 返回第n个字符串SELECT ELT(2,'a','b','c') -- b                |
| EXPORT_SET(x,s1,s2)            | 返回一个字符串，在这里对于在“bits”中设定每一位，你得到一个“on”字符串，并且对于每个复位(reset)的位，你得到一个 “off”字符串。每个字符串用“separator”分隔(缺省“,”)，并且只有“bits”的“number_of_bits” (缺省64)位被使用。SELECT EXPORT_SET(5,'Y','N',',',4) -- Y,N,Y,N |
| FIELD(s,s1,s2...)              | 返回第一个与字符串s匹配的字符串位置SELECT FIELD('c','a','b','c') -- 3 |
| FIND_IN_SET(s1,s2)             | 返回在字符串s2中与s1匹配的字符串的位置                       |
| MAKE_SET(x,s1,s2)              | 返回一个集合 (包含由“,”字符分隔的子串组成的一个 字符串)，由相应的位在`bits`集合中的的字符串组成。`str1`对应于位0，`str2`对 应位1，等等。SELECT MAKE_SET(1\|4,'a','b','c'); -- a,c |
| SUBSTRING_INDEX                | 返回从字符串str的第count个出现的分隔符delim之后的子串。如果count是正数，返回第count个字符左边的字符串。如果count是负数，返回第(count的绝对值(从右边数))个字符右边的字符串。SELECT SUBSTRING_INDEX('a*b','*',1) -- a SELECT SUBSTRING_INDEX('a*b','*',-1) -- b SELECT SUBSTRING_INDEX(SUBSTRING_INDEX('a*b*c*d*e','*',3),'*',-1) -- c |
| LOAD_FILE(file_name)           | 读入文件并且作为一个字符串返回文件内容。文件必须在服务器上，你必须指定到文件的完整路径名，而且你必须有file权 限。文件必须所有内容都是可读的并且小于max_allowed_packet。 如果文件不存在或由于上面原因之一不能被读出，函数返回NULL。 |

## 时间

| 函数                                                         | 说明                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| CURDATE(),CURRENT_DATE()                                     | 返回当前日期SELECT CURDATE() ->2014-12-17                    |
| CURTIME(),CURRENT_TIME                                       | 返回当前时间SELECT CURTIME() ->15:59:02                      |
| NOW(),CURRENT_TIMESTAMP(),LOCALTIME(),SYSDATE(),LOCALTIMESTAMP() | 返回当前日期和时间SELECT NOW() ->2014-12-17 15:59:02         |
| UNIX_TIMESTAMP()                                             | 以UNIX时间戳的形式返回当前时间SELECT UNIX_TIMESTAMP() ->1418803177 |
| UNIX_TIMESTAMP(d)                                            | 将时间d以UNIX时间戳的形式返回SELECT UNIX_TIMESTAMP('2011-11-11 11:11:11') ->1320981071 |
| FROM_UNIXTIME(d)                                             | 将UNIX时间戳的时间转换为普通格式的时间SELECT FROM_UNIXTIME(1320981071) ->2011-11-11 11:11:11 |
| UTC_DATE()                                                   | 返回UTC日期SELECT UTC_DATE() ->2014-12-17                    |
| UTC_TIME()                                                   | 返回UTC时间SELECT UTC_TIME() ->08:01:45 (慢了8小时)          |
| MONTH(d)                                                     | 返回日期d中的月份值，1->12SELECT MONTH('2011-11-11 11:11:11') ->11 |
| MONTHNAME(d)                                                 | 返回日期当中的月份名称，如JanyarySELECT MONTHNAME('2011-11-11 11:11:11') ->November |
| DAYNAME(d)                                                   | 返回日期d是星期几，如Monday,TuesdaySELECT DAYNAME('2011-11-11 11:11:11') ->Friday |
| DAYOFWEEK(d)                                                 | 日期d今天是星期几，1星期日，2星期一SELECT DAYOFWEEK('2011-11-11 11:11:11') ->6 |
| WEEKDAY(d)                                                   | 日期d今天是星期几， 0表示星期一，1表示星期二                 |
| WEEK(d)，WEEKOFYEAR(d)                                       | 计算日期d是本年的第几个星期，范围是0->53SELECT WEEK('2011-11-11 11:11:11') ->45 |
| DAYOFYEAR(d)                                                 | 计算日期d是本年的第几天SELECT DAYOFYEAR('2011-11-11 11:11:11') ->315 |
| DAYOFMONTH(d)                                                | 计算日期d是本月的第几天SELECT DAYOFMONTH('2011-11-11 11:11:11') ->11 |
| QUARTER(d)                                                   | 返回日期d是第几季节，返回1->4SELECT QUARTER('2011-11-11 11:11:11') ->4 |
| HOUR(t)                                                      | 返回t中的小时值SELECT HOUR('1:2:3') ->1                      |
| MINUTE(t)                                                    | 返回t中的分钟值SELECT MINUTE('1:2:3') ->2                    |
| SECOND(t)                                                    | 返回t中的秒钟值SELECT SECOND('1:2:3') ->3                    |
| EXTRACT(type FROM d)                                         | 从日期d中获取指定的值，type指定返回的值SELECT EXTRACT(MINUTE FROM '2011-11-11 11:11:11') ->11type可取值为：MICROSECOND SECOND MINUTE HOUR DAY WEEK MONTH QUARTER YEAR SECOND_MICROSECOND MINUTE_MICROSECOND MINUTE_SECOND HOUR_MICROSECOND HOUR_SECOND HOUR_MINUTE DAY_MICROSECOND DAY_SECOND DAY_MINUTE DAY_HOUR YEAR_MONTH |
| TIME_TO_SEC(t)                                               | 将时间t转换为秒SELECT TIME_TO_SEC('1:12:00') ->4320          |
| SEC_TO_TIME(s)                                               | 将以秒为单位的时间s转换为时分秒的格式SELECT SEC_TO_TIME(4320) ->01:12:00 |
| TO_DAYS(d)                                                   | 计算日期d距离0000年1月1日的天数SELECT TO_DAYS('0001-01-01 01:01:01') ->366 |
| FROM_DAYS(n)                                                 | 计算从0000年1月1日开始n天后的日期SELECT FROM_DAYS(1111) ->0003-01-16 |
| DATEDIFF(d1,d2)                                              | 计算日期d1->d2之间相隔的天数SELECT DATEDIFF('2001-01-01','2001-02-02') ->-32 |
| ADDDATE(d,n)                                                 | 计算其实日期d加上n天的日期                                   |
| ADDDATE(d，INTERVAL expr type)                               | 计算起始日期d加上一个时间段后的日期SELECT ADDDATE('2011-11-11 11:11:11',1) ->2011-11-12 11:11:11 (默认是天)SELECT ADDDATE('2011-11-11 11:11:11', INTERVAL 5 MINUTE) ->2011-11-11 11:16:11 (TYPE的取值与上面那个列出来的函数类似) |
| DATE_ADD(d,INTERVAL expr type)                               | 同上                                                         |
| SUBDATE(d,n)                                                 | 日期d减去n天后的日期SELECT SUBDATE('2011-11-11 11:11:11', 1) ->2011-11-10 11:11:11 (默认是天) |
| SUBDATE(d,INTERVAL expr type)                                | 日期d减去一个时间段后的日期SELECT SUBDATE('2011-11-11 11:11:11', INTERVAL 5 MINUTE) ->2011-11-11 11:06:11 (TYPE的取值与上面那个列出来的函数类似) |
| ADDTIME(t,n)                                                 | 时间t加上n秒的时间SELECT ADDTIME('2011-11-11 11:11:11', 5) ->2011-11-11 11:11:16 (秒) |
| SUBTIME(t,n)                                                 | 时间t减去n秒的时间SELECT SUBTIME('2011-11-11 11:11:11', 5) ->2011-11-11 11:11:06 (秒) |
| DATE_FORMAT(d,f)                                             | 按表达式f的要求显示日期dSELECT DATE_FORMAT('2011-11-11 11:11:11','%Y-%m-%d %r') ->2011-11-11 11:11:11 AM |
| TIME_FORMAT(t,f)                                             | 按表达式f的要求显示时间tSELECT TIME_FORMAT('11:11:11','%r') 11:11:11 AM |
| GET_FORMAT(type,s)                                           | 获得国家地区时间格式函数select get_format(date,'usa') ->%m.%d.%Y (注意返回的就是这个奇怪的字符串(format字符串)) |

## 条件判断函数

### IF(expr,v1,v2)函数

　　如果表达式expr成立，返回结果v1；否则，返回结果v2。

```sql
SELECT IF(1 > 0,'正确','错误')    

->正确
```

### IFNULL(v1,v2)函数

　　如果v1的值不为NULL，则返回v1，否则返回v2。

```sql
SELECT IFNULL(null,'Hello Word')
->Hello Word
```

### CASE

　　**语法1：**

```sql
CASE 
　　WHEN e1
　　THEN v1
　　WHEN e2
　　THEN e2
　　...
　　ELSE vn
END
```

　　CASE表示函数开始，END表示函数结束。如果e1成立，则返回v1,如果e2成立，则返回v2，当全部不成立则返回vn，而当有一个成立之后，后面的就不执行了。

```sql
SELECT CASE 
　　WHEN 1 > 0
　　THEN '1 > 0'
　　WHEN 2 > 0
　　THEN '2 > 0'
　　ELSE '3 > 0'
　　END
->1 > 0
```

　　**语法2：**

```sql
CASE expr 
　　WHEN e1 THEN v1
　　WHEN e1 THEN v1
　　...
　　ELSE vn
END
```

　　如果表达式expr的值等于e1，返回v1；如果等于e2,则返回e2。否则返回vn。

```sql
SELECT CASE 1 
　　WHEN 1 THEN '我是1'
　　WHEN 2 THEN '我是2'
ELSE '你是谁'
```

##  加密函数

### PASSWORD（str）

该函数可以对字符串str进行加密，一般情况下，PASSWORD(str)用于给用户的密码加密。

```
SELECT PASSWORD('123')
    ->*23AE809DDACAF96AF0FD78ED04B6A265E05AA257
```

### MD5（str）

对字符串计算散列值

```

SELECT md5('123')
    ->202cb962ac59075b964b07152d234b70
```

## 转换

- ASCII(s) 返回字符串s的第一个字符的ASCII码；
- BIN（x） 返回x的二进制编码；
- HEX(x) 返回x的十六进制编码；
- OCT(x) 返回x的八进制编码；
- CONV(x,f1,f2) 返回f1进制数变成f2进制数；

### 转换数据类型

- CAST(x AS type)
- CONVERT(x,type)

　　这两个函数只对BINARY、CHAR、DATE、DATETIME、TIME、SIGNED INTEGER、UNSIGNED INTEGER。

```

SELECT CAST('123' AS UNSIGNED INTEGER) + 1
    ->124
 
SELECT '123' + 1
    ->124 其实MySQL能默认转换
 
SELECT CAST(NOW() AS DATE)
　　->2014-12-18
```

