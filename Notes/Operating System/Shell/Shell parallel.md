# Shell parallel

参考：

https://www.gnu.org/software/bash/manual/bash.html#GNU%20Parallel

https://www.jianshu.com/p/c5a2369fa613

parallel和xargs类似，用于并行运行命令。但是比xargs简便(替代了`-n`和`-P`参数)。

pattern：`command1 | parallel [option] command2`

如果没有指定command2表示`command1 | sh`

这里无需`-print`

```
chz@cyberpelican:/var/www/html$ find . -type f -name '*.html' -print | parallel gzip
gzip: ./index.nginx-debian.html.gz: Permission denied
gzip: ./index.html.gz: Permission denied
gzip: ./dvwa/external/phpids/0.6/tests/coverage/Caching_File.php.html.gz: Permission denied
gzip: ./dvwa/external/phpids/0.6/tests/coverage/Filter_Storage.php.html.gz: Permission denied
gzip: ./dvwa/external/phpids/0.6/tests/coverage/Filter.html.gz: Permission denied
gzip: ./dvwa/external/phpids/0.6/tests/coverage/Caching_Factory.php.html.gz: Permission denied
gzip: ./dvwa/external/phpids/0.6/tests/coverage/Monitor.php.html.gz: Permission denied
gzip: ./dvwa/external/phpids/0.6/tests/coverage/Filter.php.html.gz: Permission denied

```

与xargs不同的是，使用parallel不需要使用`-n`参数，默认每一行输出做为后面命令的命令行参数

```
chz@cyberpelican:/var/www/html$ { echo baidu.com;echo sohu.com; }| parallel traceroute
traceroute to baidu.com (39.156.69.79), 30 hops max, 60 byte packets
 1  192.168.80.2 (192.168.80.2)  0.108 ms  0.058 ms  0.079 ms
 2  * * *
 3  * * *
 4  * * *
 5  * * *
```

## 安装

运行前需要安装`apt install parallel`

## token

- `{}`

  表示将前一个命令的stdout放入到指定的位置

  ```
  printf '%s\n' * | parallel mv {} destdir
  ```

- `{.}`

  自动去除stdout的后缀，并将去除后缀的字符串生成新的字符。例如`foo.jpg`，变成foo

  ```
  chz@cyberpelican:/var/www/html$ ls
  360safe_cq.exe  dvwa  index.html  index.nginx-debian.html
  
  chz@cyberpelican:/var/www/html$ ls * | parallel echo {.}
  360safe_cq
  index
  index.nginx-debian
  
  dvwa:
  about
  CHANGELOG
  config
  COPYING
  docs
  ```

- `{/}`

  等价于basename函数

- `{//}`

  等价于dirname函数

- `{/.}`

  等价于basename函数并去后缀

- `:::`

  将后面的stdout做为命令行参数

  `parallel gzip ::: file1 file2`等价于` (echo file1; echo file2) | parallel gzip`

  ```
  root in /opt λ ls
   1.txt.gz   a         Blasting_dictionary-master   containerd   jdk-14.0.2_linux-x64_bin.tar   t1
   2.txt.gz   bak.xml   burpsuite pro                jdk-14.0.2   lsd_0.18.0_amd64.deb          
  root in /opt λ parallel gzip -d {} ::: *.gz     
  root in /opt λ ls
   1.txt   a         Blasting_dictionary-master   containerd   jdk-14.0.2_linux-x64_bin.tar   t1
   2.txt   bak.xml   burpsuite pro        
  ```

  



