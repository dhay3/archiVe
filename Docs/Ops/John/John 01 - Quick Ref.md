# John 01 - Quick Ref

## 0x01 Preface

john the ripper(jumbo) 是一个跨平台的密码爆破工具

要想使用 john 需要提供 password files 以及选择性地指定一个 cracking mode(如果指定了 password files 但是没有指定 cracking mode，john 默认会先使用 single crack 然后使用 wordlist mode 最后使用 incremental mode，这也就意味着效率不高)

例如

```
john password.list
john --wordlist=password.lst --rules passwd
```

john 默认会以如下逻辑运行

- recognize the hash type of the current hash
- generate hashes on the fly for all the passwords in the dictionary
- stop when a generated hash matches the current hash.


## 0x02 Usage

Syntax

```
john [options] [password-files]
```

当没有指定参数时 john 会默认打印出 help messages

可以通过 有多个

```
sudo unshadow /etc/passwd /etc/shadow > mypasswd.txt
```

### 0x02a Password Files

password files 以 `username:hashed-password` 组成，例如

```
root:3bbfcf08b13c49ebaed6aff9e1f5c1e9
user01:0a041b9462caa4a31bac3567e0b6e6fd9100787db2ab433d96f6d178cabfce90
```


### 0x02b cracking modes relevant

#### Single Crack

- `--single`

	使用 single crack mode

	会使用配置文件中的 `[List.Rules:Single]` 作为 mangle rules

#### Wordlist Crack

- `--wordlist=FILE`

	使用 wordlist mode

	可以和 `--stdin` 或者 `--pipe` 一起使用表示从 stdin 和 pipe 中读取 wordlist

- `--rules`

	在 wordlist mode 中启用 mangle rules

	mangle rules 由 `[List.Rules:Wordlist]` 决定

- `--dupe-supression`

	用于过滤 wordlist 中重复的 candidate password

#### Incremental Crack

- `--incremental[=MODE]`

  使用 incremental mode

#### External Crack

- `--external=MODE`

  使用 external mode

### Filter Relevant

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

  

###  Performance Relevant

- `--save-memory=LEVEL`

	告诉 john 不要占用过多的系统内存，LEVEL 1-3(值越高，节省的内存越多，但是处理效率也越慢，通常不会使用 2-3)，只有在非 single crack mode 中生效

	实际上如果使用 `--save-memory=1` 在某些情况下还可以加快破解速度(默认情况下 john 会将 password files 载入内存中，如果 password files 中的条目格式不对，例如没有 username 字段，john 会将其标示为 invalid。而使用 `--save-memory=1` 可以防止这种情况出现)

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

### Single Crack Mode

最简单的模式，通过 `--single` 调用，会使用如下值作为密码

1. login names

2. GECOS[^5][^6]

   通常在第 5 个值，通过 `adduser --comment` 添加

   例如下面例子中的 Ken Hess

   ```
   khess:x:1001:1001:Ken Hess:/home/khess:/bin/bash
   ```

3. users’ home directory

同时会使用大量的 mangling rules 用于生成 variant passwords

#### Example

> [!NOTE]
> here string(`<<<`) 默认会将 carrage return(`\n`) 输出，所以不能使用 `md5sum <<< root` 这种形式获取哈希值

为了方便测试，使用任意哈希算法生成 password file

```
echo -n ToR | sha256sum | cut -d ' ' -f 1 | xargs -i echo -n tor:{} > crack.txt
```

指定使用 single mode

```
john --format=raw-sha256  --single crack.txt
Using default input encoding: UTF-8
Loaded 1 password hash (Raw-SHA256 [SHA256 128/128 AVX 4x])
Warning: poor OpenMP scalability for this hash type, consider --fork=32
Will run 32 OpenMP threads
Press 'q' or Ctrl-C to abort, almost any other key for status
ToR              (tor)
1g 0:00:00:00 DONE (2024-09-24 15:36) 25.00g/s 9600p/s 9600c/s 9600C/s /tor..tor73
Use the "--show --format=Raw-SHA256" options to display all of the cracked passwords reliably
Session completed
```

需要注意的是如果没有指定 format，john 会自己推测 password files 使用的 hash type（大概率不准）。所以通常需要使用 `--format` 指定 hash type

```
Warning: detected hash type "gost", but the string is also recognized as "xxxx"
```


### 0x03a Wordlist Mode

最常用的模式就是 wordlist mode，通过 `--wordlist | --stdin` 调用，顾名思义就是使用 wordlist 来破解。wordlist 通常为明文，以 carrage return 分隔

一个好的 wordlist 应该要去重并且按照字母排序，这可以加快 john 的破解效率，可以使用如下命令(将大小转为小写，由 mangling rule)(`--rules`) 来控制 variant passwords)

```
tr A-Z a-z < wordlist.source | sort -u > wordlist.target
```

#### Example

为了方便测试，使用任意哈希算法生成 password file

```
echo -n toor | md5sum | cut -d ' ' -f 1 | xargs -i echo -n root:{} > crack.txt
```

指定字典

```
john --format=raw-md5 --wordlist=common_roots.txt crack.txt
Using default input encoding: UTF-8
Loaded 1 password hash (Raw-MD5 [MD5 128/128 AVX 4x3])
Warning: no OpenMP support for this hash type, consider --fork=32
Press 'q' or Ctrl-C to abort, almost any other key for status
toor             (root)
1g 0:00:00:00 DONE (2024-09-24 16:29) 50.00g/s 230400p/s 230400c/s 230400C/s tlah..willow
Use the "--show --format=Raw-MD5" options to display all of the cracked passwords reliably
Session completed

```

同样的和 single mode 相同，也需要通过 `--format` 指定 hash type

### Incremental Mode

最强大的模式，会使用各种可能的密码组合(也就意味着几乎不会停止)，所以通常只用在 3 字符组合的密码。如果想要使用该模式，通常需要和 ` ` 一起使用

#### Example

为了方便测试，使用任意哈希算法生成 password file

```
echo -n passw0rd | md5sum | cut -d ' ' -f 1 | xargs -i echo -n root:{} > crack.txt
```

指定使用 incremental mode

```
john --format=raw-md5 --incremental crack.txt
Using default input encoding: UTF-8
Loaded 1 password hash (Raw-MD5 [MD5 128/128 AVX 4x3])
Warning: no OpenMP support for this hash type, consider --fork=32
Press 'q' or Ctrl-C to abort, almost any other key for status
passw0rd         (root)
1g 0:00:00:00 DONE (2024-09-24 17:10) 2.272g/s 930763p/s 930763c/s 930763C/s password..pashie11
Use the "--show --format=Raw-MD5" options to display all of the cracked passwords reliably
Session completed
```

当然也需要指定 hash type

### External Mode



## 0x04 Managling Rules



## 0x05 Examples

https://www.openwall.com/john/doc/EXAMPLES.shtml

## 0x06 GUI

如果想要使用 GUI 可以使用 johnny(必须先安装 john)

如果要在 Arch 上使用 johnny， 具体参考 https://github.com/openwall/johnny/blob/v2.2/INSTALL 中 debain based distro 部分编译方式

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***References***

[^1]:https://github.com/openwall/johnny/blob/v2.2/INSTALL
[^2]:https://github.com/openwall/john/blob/bleeding-jumbo/doc/OPTIONS
[^3]:https://www.freecodecamp.org/news/crack-passwords-using-john-the-ripper-pentesting-tutorial/
[^4]:https://www.openwall.com/john/doc/MODES.shtml
[^5]:https://en.wikipedia.org/wiki/Gecos_field
[^6]:https://www.redhat.com/sysadmin/linux-gecos-demystified
[^7]:https://superuser.com/questions/1684358/john-the-ripper-on-kali-linux-it-outputs-no-password-hashes-loaded
[^8]:https://www.openwall.com/john/doc/RULES.shtml

***FootNotes***