# john

## 0x01 Overview

Syntax

```
john [options] [password-files]
```

passowrd-files

可以通过 有多个

```
sudo unshadow /etc/passwd /etc/shadow > mypasswd.txt
```

john the ripper(jumbo) 是 Linux 上的一个密码爆破工具

This is how John works by default:

- recognize the hash type of the current hash
- generate hashes on the fly for all the passwords in the dictionary
- stop when a generated hash matches the current hash.

## 0x02 Options

当没有指定参数时 john 会默认打印出 help messages

如果指定了 password files 但是没有给出其他的参数，john 默认会使用所有的 cracking modes 来破解密码（这也就意味着效率不高）

### cracking modes relevant

- `--single`

  使用 single crack mode


- `--wordlist=FILE`

  使用 wordlist mode

  可以和 `--stdin` 或者 `--pipe` 一起使用表示从 stdin 和 pipe 中读取 wordlist

- `--rules`

  在 wordlist mode 中启用 mangle rules

- `--dupe-supression`

  用于过滤重复的 candidate password

- `--incremental[=MODE]`

  使用 incremental MODE

- `--external=MODE`

  使用 external mode

### condition relevant

- `--users=[-]LOGIN|UID[,..]`

  破解(或者不破解)指定用户

  ```
  #只破解用户 cc
  john --users=cc mypasswd.txt
  
  #不破解用户cc 其他用户都破解
  john --users=-cc mypasswd.txt
  
  #破解用户 cc 和 root
  john --users=cc,root mypasswd.txt
  ```

- `--groups=[-]GID[,..]`

  破解(或者不破解)指定用户组，类似 `--users`

- `--shells=[-]SHELL[,..]`

  破解(或者不破解)使用指定 Shell 的用户，类似 `--users`。非常有用，用来过滤 `nologin`

  ```
  john --shells=-nologin mypasswd.txt
  ```

- `--rules=[rule[,...]]`

  wordlist mode 使用的 mangling rules

  

  具体规则可以查看 wordlist rules syntax[^8]

- `--format=[name[,name...]]`

  john 默认会自动检查加密的 Hash 算法，但是可能也不太准（l例如在 kali 中需要使用 `--format=crypt` 来指定所有的加密算法）

  name 的值可以是如下几种

  1. 完整的 Hash 算法名，例如 `--format=cyrpt-md5`
  2. name 中可以使用 wildcard 和 Linux 中的逻辑相同。例如 `--format=mysql-*`

  2. 也可以使用 `--format=@name` 匹配包含 name 子字符的 Hash 算法（标签）。例如 `--format=@ASE` 匹配 ASE256, ASE512 DES-ASE 等
  3. 也可以使用 `--format=#name` 匹配包含 name 子字符的 Hash 算法(完整名字)
  4. 或者使用 `--format=all` 表示所有 Hash 算法

  注意 john 在同一时间默认只会使用一种 Hash 算法来破解，即使你指定了多种 Hash 算法

  

###  processing relevant

- `--node=MIN[-MAX]/TOTAL`

  

- `--fork=N`

  



### mix

- `--list=[WHAT]`

  用于查看详细的帮助信息，WHAT 可以是

  1. help 

  2. rules

  3. formats

     john 支持的所有 formats

  4. format-details

     所有 formats 的信息信息

- `--restore=[NAME]`

  继续一个没有完成破解的 session，默认会从 `$JOHN/john.rec` 中读取

- `--session=NAME`

  指定 session 名字，会生成 `NAME.rec` 文件

- `--status[=NAME]`

  打印出正在运行或者是终止的 session

- `--show`

  查看破解的密码（实际是查看 `$JOHN/john.pot` 中的内容），可以在跑 john 进程时另外起一个窗口使用 `john --show` 来查看

- `--test[=TIME]`

  一般用于校验 john 是否正常运行

- `--stdin`

## 0x03 Cracking Modes

cracking modes 是 john 中重要的一部分，告诉 john 改如何破解密码，主要有几种

1. wordlist mode
2. single crack mode
3. incremental mode
4. external mode
5. Markov mode
6. mask mode
7. subsets mode
8. regex mode
9. stacking of multiple modes

这里只介绍主流的几种模式[^4]

### Wordlist Mode

最常用的模式就是 wordlist mode，顾名思义就是使用 wordlist 来破解

一个好的 wordlist 应该要去重并且按照字母排序，这可以加快 john 的破解效率，可以使用如下命令(将大小转为小写，由 mangling rule 来控制 variant passwords)

```
tr A-Z a-z < wordlist.source | sort -u > wordlist.target
```

### Single Crack Mode

最简单的模式会使用如下值作为密码

1. login names

2. GECOS[^5][^6]

   通常在第 5 个值，通过 `adduser --comment` 添加

   例如下面例子中的 Ken Hess

   ```
   khess:x:1001:1001:Ken Hess:/home/khess:/bin/bash
   ```

3. users’ home directory

同时会使用大量的 mangling rules 用于生成 variant passwords

### Incremental Mode

最强大的模式，会使用各种可能的密码组合

### External Mode



## 0x04 Managling Rules



## 0x05 Examples

https://www.openwall.com/john/doc/EXAMPLES.shtml

## 0x06 GUI

如果想要使用 GUI 可以使用 johnny(必须先安装 john)

如果要在 Arch 上使用 johnny， 具体参考 https://github.com/openwall/johnny/blob/v2.2/INSTALL 中 debain based distro 部分编译方式



**references**

[^1]:https://github.com/openwall/johnny/blob/v2.2/INSTALL
[^2]:https://github.com/openwall/john/blob/bleeding-jumbo/doc/OPTIONS
[^3]:https://www.freecodecamp.org/news/crack-passwords-using-john-the-ripper-pentesting-tutorial/
[^4]:https://www.openwall.com/john/doc/MODES.shtml
[^5]:https://en.wikipedia.org/wiki/Gecos_field
[^6]:https://www.redhat.com/sysadmin/linux-gecos-demystified
[^7]:https://superuser.com/questions/1684358/john-the-ripper-on-kali-linux-it-outputs-no-password-hashes-loaded
[^8]:https://www.openwall.com/john/doc/RULES.shtml