# mktemp & trap

参考：https://wangdoc.com/bash/mktemp.html

## 临时文件的安全问题

1. 在较早的Linux版中，`/tmp`下创建的文件默认所有人可以读取

2. 如果攻击者知道临时文件的文件名，他可以创建符号链接，链接到临时文件(可以通过修改链接来修改临时文件)，可能导致系统运行异常。
3. 如果脚本意外退出，往往会忽略清理临时文件

==生成临时文件应该遵循下面规则==

- 创建前检查文件是否已经存在。
- 确保临时文件已成功创建。
- 临时文件必须有权限的限制。
- 临时文件要使用不可预测的文件名。
- 脚本退出时，要删除临时文件（使用`trap`命令）。

## mktemp

`mktemp`命令生成的临时文件名是随机的，而且权限是只有用户本人可读写。并不会检查文件是否存在

```
[root@cyberpelican tmp]# mktemp 
/tmp/tmp.1yyTTGdz2w
[root@cyberpelican tmp]# ll
-rw-------. 1 root root  0 Nov 21 10:27 tmp.1yyTTGdz2w
drwx------. 2 root root  6 Oct 25 21:31 tracker-extract-files.0
```

Bash 脚本使用`mktemp`命令的用法如下。`exit 1`防止`mktemp`因为创建文件失败而执行失败

```
[root@cyberpelican opt]# cat test.sh 
#!/bin/bash
TMP=$(mktemp) || exit 1
echo "Our temp file is $TMP"
[root@cyberpelican opt]# ./test.sh 
Our temp file is /tmp/tmp.LiHqLhNfU2
```

### 参数

- `-d`

  创建一个临时目录

  ```
  $ mktemp -d
  /tmp/tmp.Wcau5UjmN6
  ```

- `-t`

  指定创建文件名的模板

  ```
  $ mktemp -t mytemp.XXXXXXX
  /tmp/mytemp.yZ1HgZV
  ```

- `-p`

  指定创建临时文件的目录，如果没有指定使用`/tmp`

  ```
  $ mktemp -p /home/ruanyf/
  /home/ruanyf/tmp.FOKEtvs2H3
  ```





