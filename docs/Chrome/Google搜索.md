# Google搜索

[TOC]

## 基本逻辑

- "A B"

  强制进行完全匹配

  ```
  "steve jobs"
  ```

- A AND B

  Google默认操作符就是AND，即返回A和B相关的搜索结果。对于常规搜索，它实际上并没有多大区别，但与其它操作符搭配，就会非常有用。

  ```
  分词 AND 教程
  ```

- A OR  B

  搜索结果为A或者B，再或者A和B，也可以使用符号“|”替代OR。

  ```
  A | B
  ```

- -A

  -A表示排除搜索词A，即返回的搜索结果排除与A有关的。

  ```
  苹果 -公司
  ```

- *

  通配符, 匹配任何单个关键字

  ```
  steve * apple
  ```

- #### ()

  通过()对搜索操作符进行分组，以控制搜索的执行方式。

  ```
  (git OR svn) 回退
  ```

以上就是基本的搜索逻辑，可以组合出复杂的查询表达式，缩小搜索范围，解决歧义问题。

Google还提供了属性字段，限制搜索范围，更加精确。

## 属性

- filetype

  将搜索结果限制为特定的文件类型，如：PDF、PPT等。也可以使用“ext:”，作用是相同的。

  ```
  filetype:torrent kali
  ```

- book

  查询书籍

  ```
  book:hacker
  ```

- site

  限制搜索结果在指定的网站。

  ```
  site:juejin.im 深入理解NLP的中文分词
  ```

- intitle

  限定搜索词, 在网页标题里进行查找, 这样更契合网页的主题。

  ```
  intitle:自然语言处理
  ```

- inurl

  限定搜索词，在URL里进行查找。URL的标准格式如下：

  ```
  inurl:juejin
  ```

- intext

  限定搜索词, 在网页内容中查找

  ```
  intext:admin|password
  ```

- #### cache

  返回网页的最新缓存,就可以查看网页历史信息, 一些被删除的页面并不会立即消失, 而是在浏览器的缓存中

  ```
  cache:xuegod.cn
  ```

> 谷歌黑客数据库
>
> https://www.exploit-db.com/google-hacking-database
>
> 关键字
>
> mysql.ini #windows mysql 配置文件
>
> my.cnf #linux mysql 配置文件
>
> login 后台
>
> index.of  
>
> bash_history
>
> 学习
>
> ==intitle:index.of hack|hacking|hacker==

1. intitel:index of  intext:passwd|password|root|admin|config

2. intitle:index.of  ( .bash_history|mysql.ini|my.cnf)

   含有index.of表示该站点目录对我们开放, 可以查看目录下的所有文件信息

   .bash_history表示我们要筛选的文件名, 也可以替换成其他的敏感信息文件, 该文件记录含有历史命令记录, 同时查找mysql数据库账户和密码

3. intext:user.sql intitle:index.of

   查找sql脚本

4. inurl:login|admin intext:后台 

   ==如果没有验证码就进行爆破==

5. site:域名inurl:login|admin|manage|member|admin_login|login_admin|system|login|user|main|cms

   寻找后台登入地址

5. site:域名 intext:管理|后台|登陆|用户名|密码|验证码|系统|帐号|admin|login|sys|managetem|password|username

    查找文本内容

6. site:域名 inurl:aspx|jsp|php|asp

   查找可注入站点

7. ==site:github.com inurl:src intext:password|username|accessKeySecret|accessKeyId==

   github可是好网站, 有不少程序员将密码上传, aliyun Oss账户和密码

8. filetyle:xls inurl:gov username password

   敏感数据收集

9. filetype:inc inurl:config.inc host

10. `"/** MySQL database password */" ext:txt | ext:cfg | ext:env | ext:in`

    

