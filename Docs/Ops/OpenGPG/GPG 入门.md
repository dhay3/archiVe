# GPG 入门

参考：

https://www.ruanyifeng.com/blog/2013/07/gpg.html

https://blog.starryvoid.com/archives/348.html

https://www.gnupg.org/documentation/manuals/gnupg/

https://www.linode.com/docs/guides/gpg-keys-to-send-encrypted-messages/

https://www.gnupg.org/documentation/manpage.html

> 在linux中可以通过`man gpg2`来查看具体命令
>
> Linux的distro中一般使用`gungp`
>
> 文件一般存储在`~/.gungp`
>
> `[root@chz .gnupg]# apt-get --only-upgrade install gnupg`

## 概述

GNU Privacy Guard (GPG)，是一种加密软件。用于生成对称和非对称密匙，数字签名。

在Linux上自带该工具

```
[root@cyberpelican bin]# ll | grep gpg2
lrwxrwxrwx.   1 root root          4 Aug 24 07:51 gpg -> gpg2
-rwxr-xr-x.   1 root root     749976 Jul 13  2018 gpg2
```

使用`gpg --version`和`gpg --help`来查看相应内容

## 输入输出

- `--yes`

  默认使用yes做为大多数问题的答案

- `-v`

  输出详细信息，额外使用`-v`参数，输出信息更加详细

  ```
  [root@cyberpelican /]# gpg --list-keys -v -v
  gpg: using PGP trust model
  gpg: key 57666819: accepted as trusted key
  gpg: key 1368766C: accepted as trusted key
  /root/.gnupg/pubring.gpg
  ------------------------
  pub   2048R/57666819 2020-12-09 [expires: 2021-02-17]
  uid                  kikochz (test for gpg) <kikochz@163.com>
  sub   2048R/4D223A22 2020-12-09 [expires: 2021-02-17]
  
  pub   2048R/1368766C 2020-12-09
  uid                  zhangsan <z@163.com>
  sub   2048R/B29AF973 2020-12-09
  ```

- `--armor`

  gpg默认以二进制的格式输出，使用该参数以ASCII形式输出

  ```
  [root@cyberpelican /]# gpg --armor --output test.txt --export kikochz@163.com
  [root@cyberpelican /]# cat test.txt 
  -----BEGIN PGP PUBLIC KEY BLOCK-----
  Version: GnuPG v2.0.22 (GNU/Linux)
  
  mQENBF/QuakBCADo2pmGO9SJDXbjNn3lhhC8+M3dtxwjtgx9gWDDLB1tmN1TR4GM
  JpBjdrv/69xXMeBbfkE8mA86qWUTPfGFXbS5YbOEKGm9fZrGnIdQ6disB3RFvOxB
  m7XFcAMBBGiVT8f43xUf1rVvPcWbEEsAumPirVqgSfGTEVVQ58cJsUIqjGRvSpYi
  ```

## 生成密钥

> 管理密钥查看How to manage your keys

可以使用 `--quick-generate-key`，`--generate-key`，`--full-generate-key` 来生成 GPG key

这里使用`gpg --gen-key`来生成密匙，生成 GPG key 需要 Passphrase用作私钥的密码。然后我们可以做一些随机动作(移动鼠标，敲键盘)，以生成一个随机数。默认生成的密钥对有效期为生成的当天，==我们可以通过`gpg --full-generate-key=`来指定生成的密钥对的有效期==

```
root@chz:~# gpg --gen-key
gpg (GnuPG) 2.2.20; Copyright (C) 2020 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Note: Use "gpg --full-generate-key" for a full featured key generation dialog.

GnuPG needs to construct a user ID to identify your key.

Real name: kikochz
Email address: kikochz@163.com
You selected this USER-ID:
    "kikochz <kikochz@163.com>"

Change (N)ame, (E)mail, or (O)kay/(Q)uit? O
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
gpg: key 2ED4A52162256625 marked as ultimately trusted
gpg: directory '/root/.gnupg/openpgp-revocs.d' created
gpg: revocation certificate stored as '/root/.gnupg/openpgp-revocs.d/52A645FB16733E1A3875EEC92ED4A52162256625.rev'
public and secret key created and signed.

pub   rsa3072 2020-12-10 [SC] [expires: 2022-12-10]
      52A645FB16733E1A3875EEC92ED4A52162256625
uid                      kikochz <kikochz@163.com>
sub   rsa3072 2020-12-10 [E] [expires: 2022-12-10]
```

gpg生成的文件一般存储在`~/.gnupg`(以二进制的形式存储)，上述的` key 57666819 marked as ultimately trusted`其中的`57666819 `是==用户的Hash字符串，可以用来替换用户的ID==。`Key fingerprint = 5ABA 1A12 C664 473F BDC8  990F 5264 F015 5766 6819`用于批操作

```
[root@cyberpelican .gnupg]# ls
gpg.conf           pubring.gpg   random_seed  S.gpg-agent
private-keys-v1.d  pubring.gpg~  secring.gpg  trustdb.gpg
```

## 修改私钥密码

```
[root@cyberpelican /]# gpg --passwd kikochz@163.com
gpg (GnuPG) 2.0.22; Copyright (C) 2013 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
```

## 查看密钥

使用`gpg --list-keys`来查看所有的钥匙对，==可以指定userID或邮箱来查看指定用户的钥匙对信息==

```
[root@cyberpelican /]# gpg --list-keys kikochz@163.com
/root/.gnupg/pubring.gpg 
------------------------
pub   2048R/57666819 2020-12-09 [expires: 2021-02-17]
uid                  kikochz (test for gpg) <kikochz@163.com>
sub   2048R/4D223A22 2020-12-09 [expires: 2021-02-17]

[root@cyberpelican /]# gpg --list-secret-keys 57666819
/root/.gnupg/secring.gpg
------------------------
sec   2048R/57666819 2020-12-09 [expires: 2021-02-17]
uid                  kikochz (test for gpg) <kikochz@163.com>
ssb   2048R/4D223A22 2020-12-09
```

1. 第一行显示公钥所在文件
2. 第二行显示公钥的特征，生成时间和失效时间 ，`57666819`这里会显示用户ID
3. 第三行显示用户名，注释，用户邮箱(用户ID)
4. 第四行显示私钥的特征，生成时间和失效时间

使用`gpg --fingerprint `来显示指纹和钥匙对。==用于校验密钥的身份（别人可以通过你的指纹查询到你的公钥）==

```
[root@cyberpelican /]# gpg --fingerprint 
/root/.gnupg/pubring.gpg
------------------------
pub   2048R/57666819 2020-12-09 [expires: 2021-02-17]
      Key fingerprint = 5ABA 1A12 C664 473F BDC8  990F 5264 F015 5766 6819
uid                  kikochz (test for gpg) <kikochz@163.com>
sub   2048R/4D223A22 2020-12-09 [expires: 2021-02-17]

pub   2048R/1368766C 2020-12-09
      Key fingerprint = 3109 99FD 0FE9 81B9 9479  A012 1587 94A7 1368 766C
uid                  zhangsan <z@163.com>
sub   2048R/B29AF973 2020-12-09

```

## 删除密钥对

> 需要先删除私钥，才能删除公钥。如果需要批量删除密钥，需要使用fingerprint。

```
[root@cyberpelican /]# gpg --delete-secret-key z@163.com
gpg (GnuPG) 2.0.22; Copyright (C) 2013 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.


sec  2048R/1368766C 2020-12-09 zhangsan <z@163.com>

Delete this key from the keyring? (y/N) y
This is a secret key! - really delete? (y/N) y
[root@cyberpelican /]# gpg --delete-key z@163.com
gpg (GnuPG) 2.0.22; Copyright (C) 2013 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.


pub  2048R/1368766C 2020-12-09 zhangsan <z@163.com>

Delete this key from the keyring? (y/N) y
[root@cyberpelican /]# gpg --list-keys
gpg: checking the trustdb
gpg: 3 marginal(s) needed, 1 complete(s) needed, PGP trust model
gpg: depth: 0  valid:   1  signed:   0  trust: 0-, 0q, 0n, 0m, 0f, 1u
/root/.gnupg/pubring.gpg
------------------------
pub   2048R/8DCFA23D 2020-12-09
uid                  kikochz (kiko) <kikochz@163.com>
sub   2048R/8A08D086 2020-12-09
```

## 导出密匙

- `--export`

  导出公钥，如果不使用`--armor`参数，默认以二进制的形式在SDOUT中输出。可以使用`--output`导入到指定文件，持久保存。

  ```
  [root@cyberpelican /]# gpg --armor --export kikochz@163.com
  -----BEGIN PGP PUBLIC KEY BLOCK-----
  Version: GnuPG v2.0.22 (GNU/Linux)
  
  mQENBF/QuakBCADo2pmGO9SJDXbjNn3lhhC8+M3dtxwjtgx9gWDDLB1tmN1TR4GM
  JpBjdrv/69xXMeBbfkE8mA86qWUTPfGFXbS5YbOEKGm9fZrGnIdQ6disB3RFvOxB
  m7XFcAMBBGiVT8f43xUf1rVvPcWbEEsAumPirVqgSfGTEVVQ58cJsUIqjGRvSpYi
  ZhYBIryvM+Q4sOKTk6hcdR0vnTzvbcx3a5McEx5ExuTclFwNiUULd6WL3YFyi69C
  tT3tY2W4SSH+HYMMPTsCmx8c36qmGgpOhF6tGSBF55UigFWx9n9xB14OoMLSaMJm
  tgo8vor6sIdjojzi65ZCxLsigNEhjeKeGdR9ABEBAAG0KGtpa29jaHogKHRlc3Qg
  ```

- `--export-secrect-keys`

  导出私钥，==这是一个危险的操作==。如果在windows上操作，权限会导致拒绝导出私钥。==但是可以使用git解决。==

  ```
  [root@cyberpelican /]# gpg --armor --export-secret-keys kikochz@163.com
  -----BEGIN PGP PRIVATE KEY BLOCK-----
  Version: GnuPG v2.0.22 (GNU/Linux)
  
  lQO+BF/QuakBCADo2pmGO9SJDXbjNn3lhhC8+M3dtxwjtgx9gWDDLB1tmN1TR4GM
  JpBjdrv/69xXMeBbfkE8mA86qWUTPfGFXbS5YbOEKGm9fZrGnIdQ6disB3RFvOxB
  m7XFcAMBBGiVT8f43xUf1rVvPcWbEEsAumPirVqgSfGTEVVQ58cJsUIqjGRvSpYi
  ```

## 导入公钥

> `--search-keys`只会查找在服务器上是否有keys的公钥
>
> `--recv-keys`如果服务器上有keys对应的公钥，就会导入到本地的数据库
>
> 我们可以在`pool.sks-keyservers.net`上查找公钥

- `--import <pub-key file>` 

  将公钥导入到本地的数据库，导入后可以用作校验签名。==也可以导入私钥==

- `--recv-keys <fingerprint>`

  如果没有公钥，可以通过fingerprint导入到本地数据库

  ```
  gpg --keyserver pgpkeys.mit.edu --recv-key DE885DD3
  ```

## 上传密匙

参考：https://www.zhihu.com/question/25705489

> 可用的公钥服务器
>
> - http://keys.gnupg.net/
> - https://pgp.mit.edu/
> - http://pool.sks-keyservers.net/  首选
> - https://sks-keyservers.net/
> - http://zimmermann.mayfirst.org/
> - https://keyserver.ubuntu.com/

使用`--send-keys <key IDs>`和`--keyserver`来将公钥上传到指定公钥服务器(==过一段时间所有的公钥服务器都会同步信息==)。供其他用户加密和解密。这里可以使用指纹和userID作为key ID，如果是指纹需要将空格去除。==通过指纹找到用户的公钥==

> 使用`gpg --refresh-keys`来刷新服务器上的keys，==值得注意的一点是如果keys没有上传成功没有提示消息，所以在上传不了的时候检查一下网络配置是否出现问题==

具体参考：https://askubuntu.com/questions/796271/fetching-ubuntus-openpgp-keys-for-verification-fails-with-not-a-key-id-skippi

- ~~命令行，注意的是需要将`keyserver`参数写前面~~

  ```
  [root@cyberpelican /]# gpg --keyserver hkp://pool.sks-keyservers.net  --send-keys 5ABA1A12C664473FBDC8990F5264F0155766681
  gpg: sending key 57666819 to hkp server keyserver.ubuntu.com
  
  [root@cyberpelican /]# gpg --keyserver  hkp://pool.sks-keyservers.net  --send-keys 57666819
  gpg: sending key 57666819 to hkp server keyserver.ubuntu.com
  ```

  ~~校验网站上使用指纹需要带上`0x`前缀(同样的需要去掉空格)，或是在命令行使用`--search-keys`或`--recv-keys`来校验~~

  ```
  root@chz:~# gpg --keyserver hkp://pool.sks-keyservers.net --send-keys E7BDD87346B4D623FB203FC6DFCFDF52F1627354
  gpg: sending key DFCFDF52F1627354 to hkp://pool.sks-keyservers.net
  root@chz:~# gpg --keyserver hkp://pool.sks-keyservers.net --recv-keys E7BDD87346B4D623FB203FC6DFCFDF52F1627354
  gpg: keyserver receive failed: No data
  root@chz:~# gpg --keyserver hkp://pool.sks-keyservers.net --recv-keys E7BDD87346B4D623FB203FC6DFCFDF52F1627354
  gpg: key DFCFDF52F1627354: "cyberpelican <hostlockdown@gmail.com>" not changed
  gpg: Total number processed: 1
  gpg:              unchanged: 1
  
  root@chz:~# gpg --keyserver hkp://pool.sks-keyservers.net --search-keys E7BDD87346B4D623FB203FC6DFCFDF52F1627354
  gpg: data source: http://4.35.226.103:11371
  (1)     cyberpelican <hostlockdown@gmail.com>
            3072 bit RSA key DFCFDF52F1627354, created: 2020-12-10
  Keys 1-1 of 1 for "E7BDD87346B4D623FB203FC6DFCFDF52F1627354".  Enter number(s), N)ext, or Q)uit > s
  ```

> 参考：gpg How to change the configuration

## 文件加密和解密

> 可以通过`--output`指定加密和解密生成的文件

- `--encrypt`

  需要使用`--recipient`指定接收者的ID(指纹)，可以使用`-r`参数缩写

  ```
  [root@chz Desktop]# gpg -r 52A645FB16733E1A3875EEC92ED4A52162256625 --sign --encrypt gpg.test
  File 'gpg.test.gpg' exists. Overwrite? (y/N) y
  ```

- `--decrypt`

  ```
  [root@chz Desktop]# gpg --decrypt gpg.test.gpg 
  gpg: encrypted with 3072-bit RSA key, ID 780C901925A6006A, created 2020-12-10
        "kikochz <kikochz@163.com>"
  hello world
  gpg: Signature made Thu 10 Dec 2020 06:57:40 AM EST
  gpg:                using RSA key 52A645FB16733E1A3875EEC92ED4A52162256625
  gpg: Good signature from "kikochz <kikochz@163.com>" [ultimate]
  ```

- `--recipient | -r`

  指定接收者的指纹

- `--local-user | -u`

  指定发送者的指纹，==如果没有该参数使用`--default-key`，采用`--list-keys`显示的第一用户的私钥==

  ```
  [root@chz Desktop]# gpg -u E7BDD87346B4D623FB203FC6DFCFDF52F1627354 -r 52A645FB16733E1A3875EEC92ED4A52162256625 --encrypt  gpg.test
  ```

## 生成签名

```
root@v2:/home/ubuntu/test# gpg --sign data
root@v2:/home/ubuntu/test# ls
data  data.gpg
```

会生成一个 `.gpg` 后缀的文件( 包含原始文件和签名 )，默认以 binary 格式存储 

```
root@v2:/home/ubuntu/test# gpg --decrypt data.gpg
this is a test message
gpg: Signature made Sat 25 Mar 2023 05:44:13 PM CST
gpg:                using RSA key CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
gpg: Good signature from "alice (this is a comment) <alice@yahoo.com>" [ultimate]
```

如果需要生成明文的文件，可以使用 `--clearsign` (无须和 `--sign` 一起使用)

```
root@v2:/home/ubuntu/test# gpg --clearsign data
root@v2:/home/ubuntu/test# ls
data  data.asc
```

会生成一个 `.asc` 后缀的文件( 包含原始文件中的内容和签名，以明文显示 )

```
root@v2:/home/ubuntu/test# cat data.asc
-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA512

this is a test message
-----BEGIN PGP SIGNATURE-----

iQGzBAEBCgAdFiEEy2bUJKm+TcW5urT9D8IRS+606/gFAmQevX0ACgkQD8IRS+60
6/jAIwwAgkmhf4BhmktBovGa4Hvdj4ZYU31/WtDwumckczExYerGWYX5cdFlYxWs
V7M2hNKTNkHHl80HynawrLoTu3g6maFPLEunaPHzVgVQ0FRjrcP9KWhkwiFuOZpG
hgj8c6CShLcFbSc9uLVgO1E80yjILxvyI04DkZmu8Ls5dQhEGndvmiUeQ3/5kwca
HPR9/4cR1u4RqsHTUbX/3OxePD+Jrcz3ipW3rJZAgt9BSo33mtwFUfRlZ7UIRRTd
sPbrKWh2ViuAJxB5wHPf+7og4ZTYREsCYKXGfIQeKRSjDpCQTuniniJ3C+gSe96S
JIBJ5GRG//S1KlRLAnaxW31MFkGgtwzKmtuvon5SXrIFLDNSSS9Lt+z9V/H9d8Nt
YXoF0Ovmptv+lz909pnjpD3z9nvyF1DDk1sCdI4PvA3SxWTP5+Mt63q9kv+NFIOV
YPWGUA3hVbxQDEKMueON9+HF6aUMdOt1IFp72BDadSlOqht70Dr7CMpbsIdVe5EU
Ev8Cbtzc
=MN/S
-----END PGP SIGNATURE-----
```

如果需要单独生成签名文件，可以使用 `--detach-sign`

```
root@v2:/home/ubuntu/test# gpg  --detach-sign --sign data
```

会生成一个 `.sig` 后缀的 binary 文件，表示签名

```
root@v2:/home/ubuntu/test# gpg --decrypt data.sig
gpg: assuming signed data in 'data'
gpg: Signature made Sat 25 Mar 2023 05:34:24 PM CST
gpg:                using RSA key CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
gpg: Good signature from "alice (this is a comment) <alice@yahoo.com>" [ultimate]
```

如果需要以明文的方式存储签名，可以和 `--armor` 一起使用

```
root@v2:/home/ubuntu/test# gpg --armor --detach-sign --sign data
root@v2:/home/ubuntu/test# ls
data  data.asc
root@v2:/home/ubuntu/test# cat data.asc
-----BEGIN PGP SIGNATURE-----

iQGyBAABCgAdFiEEy2bUJKm+TcW5urT9D8IRS+606/gFAmQewxsACgkQD8IRS+60
6/jD5gv2JwZf50QprtDNq/inOVrgfIPl4IzrwkvSJuJuxwLAzTQeMLSV4GPD6S5H
zTvVqCU4LKgudswbHi8A+b0qv8YtJd+B+LqoLYxyFnEw6mfl/94J6mjEehN/sn7k
CdS5jcrBU+OPKFuIfDAtcUM7BMgh3y+pEejAcFNfUv4A9lYrWGYdjsfW7kYA3qF9
tS/bEiF4kKBwSjfNmbmhwtELM5fzKdpFeEAl/ucxytZ9M+UwdFJsrTFqYmSRNnsm
PvLB2ltpL13o8yPwEV8VVt2Z64UD9215plygh9dPY9512fhN0NCH/y7nuEbT/W8m
yUpcOeTTsCIKt9xZ7E0fJ4YatVz00qME58hs65aj0nJSuIHmRE05oFi6GfN9AilI
oVHFO2N75yNg/SUqQhMGTfeQNHBqRE55koNrsBYJE3fXl1PHnWuALJ6ktzo4CEw1
AmvlJ2U2mKrP3hXtFcWk0zUGG4waCiH5/Y+ukDcF/CerdwvijnpDwsLmS4LmR5e8
0/tbyNw=
=q3qN
-----END PGP SIGNATURE-----
```

## 校验签名

> 如果没有指定`-u`发送者，默认使用`--default-key`，采用`--list-keys`显示的第一用户的私钥

- `--sign`

  不对文件加密，对文件签名，表示该文件是由本人发出。如果需要对文件加密使用`--encrypt`对文件加密（私钥加密，使用公钥解密）或是`--symmetric`使用对称加密和解密（passphase）。会生成一个以`.gpg`结尾的二进制文件

  ```
  root@chz:~/Desktop# ls
  gpg.test 360safe_cq.exe  Hack-v3.003-ttf
  root@chz:~/Desktop# gpg --sign  gpg.test
  root@chz:~/Desktop# ls
  360safe_cq.exe  gpg.test  gpg.test.gpg  Hack-v3.003-ttf
  ```

- `--clearsign`

  以ascII码形式生成文本内容

  ```
  root@chz Desktop]# gpg --clearsign gpg.test
  [root@chz Desktop]# ls
  360safe_cq.exe  gpg.test  gpg.test.asc  gpg.test.gpg  Hack-v3.003-ttf
  [root@chz Desktop]# cat gpg.test.asc 
  -----BEGIN PGP SIGNED MESSAGE-----
  Hash: SHA512
  
  hello world
  -----BEGIN PGP SIGNATURE-----
  
  iQGzBAEBCgAdFiEEUqZF+xZzPho4de7JLtSlIWIlZiUFAl/R3T4ACgkQLtSlIWIl
  ZiWhPwv9E4EpATQ6zH3QvXbAD+r+8xWoSFcrrB9XLuak0/2HwBiMm86UUAQyXM0a
  tdbLP0ajED4ylZJAz61+Z/oRtTZqrn07I/iKwr/ZBo5OS3+12WXwdaCsnmIUiQZt
  SwE1VBWOH20yK1Yn61w3JOcQieX6I36YCMzmhjBe9MRQA+l8fgmkTELjAWX47GAW
  EBSR1Jop3E0fqxZDihzIcVbavRnau/ioAzDHyLMJKB4tglR1We1lyLpWjRh++hlr
  9OysJ//nk4IVvTCZ5HovDljTrlWLLieQXoeVxWicaKXDUsKBfL0arsSqXSW35QcR
  3oyVeXHB9Zwd3rNwD1i+5Q5Z38OLvaQbx8RbH7226i478VFNaMKL73J0fJbUw3UH
  ```

- `--detach-sign`

  将文件内容与与签名单独分开，通常与`--armor`参数一起使用，将内容以ascII码形式显示，文件以`.asc`结尾，签名以`.sig`结尾

> 如果要校验文件签名是否正确，需要导入公钥使用`--import`或是`--recv-keys`

- `--verify <signature>`

  校验签名是否正确(文件需要包含签名和数据)，不会将文件内容输出

  ```
  [root@chz Desktop]# gpg --verify gpg.test.asc
  gpg: Signature made Thu 10 Dec 2020 07:14:02 AM EST
  gpg:                using RSA key 52A645FB16733E1A3875EEC92ED4A52162256625
  gpg: Good signature from "kikochz <kikochz@163.com>" [ultimate]
  gpg: WARNING: not a detached signature; file 'gpg.test' was NOT verified!
  ```

- `--verify <detach-signature> <sign-data>`

  如果是单独的签名文件，需要给出源数据。这里以校验k8s的签名文件为例
  
  ```
  C:\Users\82341\Downloads>gpg --verify repomd.xml.asc  repomd.xml
  gpg: Signature made 2021/3/16 19:17:04 中国标准时间
  gpg:                using RSA key 6A030B21BA07F4FB
  gpg: Good signature from "Google Cloud Packages Automatic Signing Key <gc-team@google.com>" [unknown]
  gpg: WARNING: This key is not certified with a trusted signature!
  gpg:          There is no indication that the signature belongs to the owner.
  Primary key fingerprint: 54A6 47F9 048D 5688 D7DA  2ABE 6A03 0B21 BA07 F4FB
  gpg: Signature made 2021/3/16 19:17:04 中国标准时间
  gpg:                using RSA key 8B57C5C2836F4BEB
  gpg: Good signature from "gLinux Rapture Automatic Signing Key (//depot/google3/production/borg/cloud-rapture/keys/cloud-rapture-pubkeys/cloud-rapture-signing-key-2020-12-03-16_08_05.pub) <glinux-team@google.com>" [unknown]
  gpg: WARNING: This key is not certified with a trusted signature!
  gpg:          There is no indication that the signature belongs to the owner.
  Primary key fingerprint: 59FE 0256 8272 69DC 8157  8F92 8B57 C5C2 836F 4BEB
  ```

## 撤销证书

参考：https://blog.starryvoid.com/archives/348.html

`gpg --gen-revoke <userID>`生成一张"撤销证书"，以备以后密钥作废时，可以请求外部的公钥服务器撤销你的公钥。==需要将文本信息保存==

1. 生成撤销证书

   ```
   [root@cyberpelican /]# gpg --gen-revoke kikochz@163.com > kikochz.revo
   
   sec  2048R/57666819 2020-12-09 kikochz (test for gpg) <kikochz@163.com>
   
   Create a revocation certificate for this key? (y/N) y
   Please select the reason for the revocation:
     0 = No reason specified
     1 = Key has been compromised
     2 = Key is superseded
     3 = Key is no longer used
     Q = Cancel
   (Probably you want to select 1 here)
   Your decision? 0
   Enter an optional description; end it with an empty line:
   > 
   Reason for revocation: No reason specified
   (No description given)
   Is this okay? (y/N) y
   
   You need a passphrase to unlock the secret key for
   user: "kikochz (test for gpg) <kikochz@163.com>"
   2048-bit RSA key, ID 57666819, created 2020-12-09
   
   ASCII armored output forced.
   Revocation certificate created.
   
   Please move it to a medium which you can hide away; if Mallory gets
   access to this certificate he can use it to make your key unusable.
   It is smart to print this certificate and store it away, just in case
   your media become unreadable.  But have some caution:  The print system of
   your machine might store the data and make it available to others!
   -----BEGIN PGP PUBLIC KEY BLOCK-----
   Version: GnuPG v2.0.22 (GNU/Linux)
   Comment: A revocation certificate should follow
   
   iQEfBCABAgAJBQJf0MDbAh0AAAoJEFJk8BVXZmgZ80QH/35ynynapWVO8/Mll7us
   rBrjiP5CRb/6lbCczNUO/wjF2oveHxRSsBfNJMFP1k5yLe9M8CmkBjYvNmYNmlDC
   bw+9k4MLndMyUOTfm+kPF2EGvuC3w+4Ps+s6jz99b0V9qpCvZDAZ8mDMyu7gmTeP
   P2Wq2hKfYPspOAIVMHS1j+JQxc8CH0G8Tdg2/hW+xuC3X0gGvjO0o6y83UMjqTtG
   ```

2. 将撤销证书导入到本地GPG库

   ```
   [root@chz openpgp-revocs.d]# gpg --import kikochz.revo
   ```

3. 将撤销的公钥的发送到公钥服务器

   ```
   [root@chz openpgp-revocs.d]# gpg --send-keys 7C51F673F082FF0C2BA64EF48A426D3F09705CFF
   gpg: sending key 8A426D3F09705CFF to hkp://pool.sks-keyservers.net
   ```

4. 一段时间后查看公钥服务器即可看到对应公钥有 revoked 字样。
   （撤销证书不能发给别人，任何拥有该证书的人都可以进行撤销公钥操作）

   ```
   [root@chz openpgp-revocs.d]# gpg --search-keys 7C51F673F082FF0C2BA64EF48A426D3F09705CFF
   gpg: data source: http://130.206.1.8:11371
   (1)     cyberpelican (let the word out) <kikochz@163.com>
             3072 bit RSA key 8A426D3F09705CFF, created: 2020-12-11 (revoked)
   Keys 1-1 of 1 for "7C51F673F082FF0C2BA64EF48A426D3F09705CFF".  Enter number(s), N)ext, or Q)uit > 
   ```


## gpg.conf

按照manual page 使用`dirmngr`来配置本人为成功，使用`~/.gnupg/gpg.conf`替代，如果使用该配置文件。具体配置查看`man gpg`

~~keyserver-options ca-cert-file 失效~~ ，==如果需要配置TSL(hkps)就必须使用`dirmngr`==

Do not write the 2 dashes, but simply the name of the option and any required arguments. ines with  a  hash ('#') as the first non-white-space character are ignored

```
#default key to sign
default-key E7BDD87346B4D623FB203FC6DFCFDF52F1627354
#which server to be used
keyserver  hkp://pool.sks-keyservers.net
verbose 
armor
```

## dirmngr

https://nova.moe/openpgp-best-practices-keyserver-and-configuration/

在GPG2.1后使用dirmngr来管理keyserver，同样可以去掉two dashes 后配置进文件。文件应该存储在`~/.gnugp/dirmngr.conf`中。`hkp-cacert`证书文件需要以`.pem`结尾

```
keyserver  hkps://pool.sks-keyservers.net
hkp-cacert  
```

启动dirmngr才会生效，`dirmngr --daemon`以守护进程形式启动