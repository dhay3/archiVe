# GPG - manage keys

ref
[https://www.gnupg.org/gph/en/manual/c14.html#AEN25](https://www.gnupg.org/gph/en/manual/c14.html#AEN25)[https://www.gnupg.org/gph/en/manual/c235.html](https://www.gnupg.org/gph/en/manual/c235.html)
[https://lists.gnupg.org/pipermail/gnupg-users/2011-January/040437.html](https://lists.gnupg.org/pipermail/gnupg-users/2011-January/040437.html)
[https://unix.stackexchange.com/questions/31996/how-are-the-gpg-usage-flags-defined-in-the-key-details-listing](https://unix.stackexchange.com/questions/31996/how-are-the-gpg-usage-flags-defined-in-the-key-details-listing)
[https://superuser.com/questions/1371088/what-do-ssb-and-sec-mean-in-gpgs-output](https://superuser.com/questions/1371088/what-do-ssb-and-sec-mean-in-gpgs-output)

https://wiki.archlinux.org/title/GnuPG#Use_a_keyserver[https://wiki.debian.org/Subkeys](https://wiki.debian.org/Subkeys)
[https://www.void.gr/kargig/blog/2013/12/02/creating-a-new-gpg-key-with-subkeys/](https://www.void.gr/kargig/blog/2013/12/02/creating-a-new-gpg-key-with-subkeys/)

## Generate keys

> 需要保存好 passphrase 如果忘了，就恢复不了！！！

GPG 有几种方式来生成 GPG Primary key

- `--quick-generate-key | --quick-gen-key user_id [algo [usage [expire]]]`
- `--generate-key | --gen-key`
- `--full-generate-key | --full-gen-key`
### quick-gen
快速生成 GPG key，必须要指定 user-id，无需指定邮箱
```
root@v2:/home/ubuntu# gpg --quick-gen-key
usage: gpg [options] --quick-generate-key USER-ID [ALGO [USAGE [EXPIRE]]]
root@v2:/home/ubuntu# gpg --quick-gen-key quick-gen
About to create a key for:
    "quick-gen"

Continue? (Y/n) y
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
gpg: key C86BCE670DBE2517 marked as ultimately trusted
gpg: revocation certificate stored as '/root/.gnupg/openpgp-revocs.d/203F2D41E8C37B75C7F0074BC86BCE670DBE2517.rev'
public and secret key created and signed.

pub   rsa3072 2023-04-06 [SC] [expires: 2025-04-05]
      203F2D41E8C37B75C7F0074BC86BCE670DBE2517
uid                      quick-gen
sub   rsa3072 2023-04-06 [E]
```
生产的 GPG key 默认当天失效，如果需要指定日期可以使用如下格式
```
root@v2:/home/ubuntu# gpg --quick-gen-key quick-gen-e rsa cert 2055-01-01
```
### gen
需要提供邮箱
```
root@v2:/home/ubuntu# gpg --gen-key
gpg (GnuPG) 2.2.19; Copyright (C) 2019 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Note: Use "gpg --full-generate-key" for a full featured key generation dialog.

GnuPG needs to construct a user ID to identify your key.

Real name: gen
Name must be at least 5 characters long
Real name: gen-key
Email address: gen-key@qq.com
You selected this USER-ID:
    "gen-key <gen-key@qq.com>"

Change (N)ame, (E)mail, or (O)kay/(Q)uit? o
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
gpg: key 23EF878D76D3CCD2 marked as ultimately trusted
gpg: revocation certificate stored as '/root/.gnupg/openpgp-revocs.d/A9E50583146B978EC2DBA0AB23EF878D76D3CCD2.rev'
public and secret key created and signed.

pub   rsa3072 2023-04-06 [SC] [expires: 2025-04-05]
      A9E50583146B978EC2DBA0AB23EF878D76D3CCD2
uid                      gen-key <gen-key@qq.com>
sub   rsa3072 2023-04-06 [E] [expires: 2025-04-05]
```
和 `--quick-gen-key` 一样生成的 key 默认 expiredate 为当天
### full-gen
以完整的方式生产 GPG key，可以指定 expiredate
```
root@v2:/home/ubuntu# gpg --full-gen-key
gpg (GnuPG) 2.2.19; Copyright (C) 2019 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Please select what kind of key you want:
   (1) RSA and RSA (default)
   (2) DSA and Elgamal
   (3) DSA (sign only)
   (4) RSA (sign only)
  (14) Existing key from card
Your selection? 
RSA keys may be between 1024 and 4096 bits long.
What keysize do you want? (3072) 
Requested keysize is 3072 bits
Please specify how long the key should be valid.
         0 = key does not expire
      <n>  = key expires in n days
      <n>w = key expires in n weeks
      <n>m = key expires in n months
      <n>y = key expires in n years
Key is valid for? (0) 0
Key does not expire at all
Is this correct? (y/N) y

GnuPG needs to construct a user ID to identify your key.

Real name: full-gen
Email address: full-gen@qq.com
Comment: this is a comment
You selected this USER-ID:
    "full-gen (this is a comment) <full-gen@qq.com>"

Change (N)ame, (C)omment, (E)mail or (O)kay/(Q)uit? o
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
gpg: key 14BED08D91E29B8D marked as ultimately trusted
gpg: revocation certificate stored as '/root/.gnupg/openpgp-revocs.d/207AA207060EE9F87F36728F14BED08D91E29B8D.rev'
public and secret key created and signed.

pub   rsa3072 2023-04-06 [SC]
      207AA207060EE9F87F36728F14BED08D91E29B8D
uid                      full-gen (this is a comment) <full-gen@qq.com>
sub   rsa3072 2023-04-06 [E]
```
### Create key without passphrase
在 GPG 2.0 之后如果想生成的密钥不使用 passphrase，就需要使用 `--pinentry-mode loopback` 

```
root@v2:~/.gnupg# gpg --full-gen-key --pinentry-mode loopback
```

## Query keys

GPG key 其实也是一对 asymmetric 秘钥，在 GPG 中有两个 options 分别用于查询公钥和私钥

- `--list-keys | -k | --list-pubic-key`
- `--list-secret-keys | -K`

option 后面可以跟上一个 user-id，如果没有指定 user-id，默认显示所有的 keys
### public keys
查看 GPG public keys 可以通过如下方式
```
root@v2:~# gpg -k test
pub   rsa3072 2023-04-06 [SC] [expires: 2023-04-13]
      49801911B98422051F6AAA86A23605A6E42927AD
uid           [ultimate] tester (this is a comment) <tester@qq.com>
sub   rsa3072 2023-04-06 [E] [expires: 2023-04-13]
```

- `pub` 表示 public key, 后面的部分从左往右分别为 秘钥加密算法bit, 秘钥创建时间，秘钥过期时间。其中的 `[SC]`表示当前 public key 的用途，分别是 S=siging 和 C=certification
- 第二行显示秘钥的 Hash 值(不特指公钥或者私钥，即指纹), 可用做 user-id 唯一标识符
- 第三行 uid 显示秘钥绑定的用户，以及邮箱 和 comment
- `sub` 表示 public subkey，其中的 `[E]` 表示当前 public subkey 的用途，E=encryption

如果需要查看指纹信息可以使用
```
root@v2:~# gpg --fingerprint -k test
pub   rsa3072 2023-04-06 [SC] [expires: 2023-04-13]
      4980 1911 B984 2205 1F6A  AA86 A236 05A6 E429 27AD
uid           [ultimate] tester (this is a comment) <tester@qq.com>
sub   rsa3072 2023-04-06 [E] [expires: 2023-04-13]
```
### secret keys
查看 GPG 私钥可以通过如下方式
```
root@v2:~# gpg -K test
sec   rsa3072 2023-04-06 [SC] [expires: 2023-04-13]
      49801911B98422051F6AAA86A23605A6E42927AD
uid           [ultimate] tester (this is a comment) <tester@qq.com>
ssb   rsa3072 2023-04-06 [E] [expires: 2023-04-13]
```

- `sec` 表示 secret key, 后面的部分从左往右分别为 秘钥加密算法bit, 秘钥创建时间，秘钥过期时间。其中的 `[SC]`表示当前 secret key 的用途，分别是 S=siging 和 C=certification
- 第二行显示秘钥的 Hash 值(不特指公钥或者私钥，即指纹), 可用做 user-id 唯一标识符
- 第三行 uid 显示秘钥绑定的用户，以及邮箱 和 comment
- `ssb` 表示 secrect subkey，其中的 `[E]` 表示当前 secret subkey 的用途，E=encryption
### subkeys
subkey 是 GPG key 的一种扩展，挂在 primary keys 下，即有自己独立的公钥和私钥。可以单独 revoke subkey，而不需要 revoke primary key
GPG 默认没有参数显示可以显示 subkeys hash ，需要使用 `gpg --edit-key` 中 list 来查看
```
gpg> list

sec  rsa3072/A23605A6E42927AD
     created: 2023-04-06  expires: 2023-04-13  usage: SC  
     trust: ultimate      validity: ultimate
ssb  rsa3072/3E95329471EBCD9A
     created: 2023-04-06  expires: 2023-04-13  usage: E
```
## Edit keys 
GPG 提供了一个用于修改 GPG key 的参数，`--edit-key` (不仅仅能做修改的操作也可以查看 key 的具体信息)
例如 
使用如下命令就会进入到 interactive mode
```
root@v2:~# gpg --edit-key test
gpg (GnuPG) 2.2.19; Copyright (C) 2019 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Secret key is available.

sec  rsa3072/A623CF87362224FE
     created: 2023-04-06  expires: 2025-04-05  usage: SC  
     trust: ultimate      validity: ultimate
ssb  rsa3072/ED7BEAB890C8DDE9
     created: 2023-04-06  expires: never       usage: E   
[ultimate] (1). test

gpg> 
```
GPG 提供了一系列的指令供我们来修改 GPG key，我们可以使用 `?` 来查看。这里列举几个常用的

- `adduid`

  用于添加一个额外的 uid，即可以添加额外的邮箱信息

- `deluid`

  用于删除指定额外的 uid，需要使用 `uid` 来中

- `uid n`

  和 `key n` 类似

- `addkey`

  添加一个 额外的 subkey

- `delkey`

  删除指定的 subkey，需要先试用 key 选择 subkey

- `key n`

  选择指定的 subkey，选择 n 可以是数字，表示 subkey 的 index；或者是 subkey hash

  也可以是 `*` 表示全选，`0` 表示全不选( 即只选择 primary key )

- `change-usage`

  修改 key 的用途，可以是 certify, sign, authenticate, encrypt

- `list`

  显示当面操作的 GPG key

- `expire`

  修改 GPG key 有效的日期

- `passwd`

  修改私钥对应的 passphrase

- `quit`

  退出 edit 界面

- `save`

  保存当前的修改并退出，如果不保存修改会被 discarded

### Add user-id

通过 adduid 可以对应 GPG key 添加额外的邮箱信息

```
root@v2:~/.gnupg# gpg --edit-key A3455E25BA6751B4B551951CC6080149806CD3C3
gpg (GnuPG) 2.2.19; Copyright (C) 2019 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Secret key is available.

sec  rsa3072/C6080149806CD3C3
     created: 2023-04-07  expires: never       usage: SCA 
     trust: ultimate      validity: ultimate
ssb  rsa3072/9F5DBAF4C83B7125
     created: 2023-04-07  expires: never       usage: E   
[ultimate] (1). Alice <Alice@gmail.com>

gpg> adduid
Real name: Alicia
Email address: Alicia@gmail.com
Comment: 
You selected this USER-ID:
    "Alicia <Alicia@gmail.com>"

Change (N)ame, (C)omment, (E)mail or (O)kay/(Q)uit? o

sec  rsa3072/C6080149806CD3C3
     created: 2023-04-07  expires: never       usage: SCA 
     trust: ultimate      validity: ultimate
ssb  rsa3072/9F5DBAF4C83B7125
     created: 2023-04-07  expires: never       usage: E   
[ultimate] (1)  Alice <Alice@gmail.com>
[ unknown] (2). Alicia <Alicia@gmail.com>

gpg> save
```

查看是否生效

```
root@v2:~/.gnupg# gpg -k A3455E25BA6751B4B551951CC6080149806CD3C3
pub   rsa3072 2023-04-07 [SCA]
      A3455E25BA6751B4B551951CC6080149806CD3C3
uid           [ultimate] Alicia <Alicia@gmail.com>
uid           [ultimate] Alice <Alice@gmail.com>
sub   rsa3072 2023-04-07 [E]
```

### Change usage

通过 change usage 我们可以对不同的 key 赋予不同的权限，如果没有使用 `key` 或者使用了 `key 0` 就表示当前只选择了 primary key

例如现在一个 key 状态如下，有两个 subkey

```
root@v2:~/.gnupg# gpg -k A3455E25BA6751B4B551951CC6080149806CD3C3
pub   rsa3072 2023-04-07 [SC]
      A3455E25BA6751B4B551951CC6080149806CD3C3
uid           [ultimate] Alicia <Alicia@gmail.com>
uid           [ultimate] Alice <Alice@gmail.com>
sub   rsa3072 2023-04-07 [E]
sub   dsa2048 2023-04-08 [S]
```

现在想要 subkey 2 只做 authentication 的功能，那么我们就先需使用 `key 2` 来选中对应的 subkey

```
root@v2:~/.gnupg# gpg --edit-key A3455E25BA6751B4B551951CC6080149806CD3C3
gpg (GnuPG) 2.2.19; Copyright (C) 2019 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Secret key is available.

sec  rsa3072/C6080149806CD3C3
     created: 2023-04-07  expires: never       usage: SC  
     trust: ultimate      validity: ultimate
ssb  rsa3072/9F5DBAF4C83B7125
     created: 2023-04-07  expires: never       usage: E   
ssb  dsa2048/AD5EC538042CD3DE
     created: 2023-04-08  expires: never       usage: S   
[ultimate] (1). Alicia <Alicia@gmail.com>
[ultimate] (2)  Alice <Alice@gmail.com>

gpg> key 2

sec  rsa3072/C6080149806CD3C3
     created: 2023-04-07  expires: never       usage: SC  
     trust: ultimate      validity: ultimate
ssb  rsa3072/9F5DBAF4C83B7125
     created: 2023-04-07  expires: never       usage: E   
ssb* dsa2048/AD5EC538042CD3DE
     created: 2023-04-08  expires: never       usage: S   
[ultimate] (1). Alicia <Alicia@gmail.com>
[ultimate] (2)  Alice <Alice@gmail.com>
```

然后再使用 `change-usage`

```
gpg> change-usage
Changing usage of a subkey.

Possible actions for a DSA key: Sign Authenticate 
Current allowed actions: Sign Authenticate 

   (S) Toggle the sign capability
   (A) Toggle the authenticate capability
   (Q) Finished

Your selection? S

Possible actions for a DSA key: Sign Authenticate 
Current allowed actions: Sign 

   (S) Toggle the sign capability
   (A) Toggle the authenticate capability
   (Q) Finished

Your selection? A

Possible actions for a DSA key: Sign Authenticate 
Current allowed actions: Sign Authenticate 

   (S) Toggle the sign capability
   (A) Toggle the authenticate capability
   (Q) Finished

Your selection? Q

sec  rsa3072/C6080149806CD3C3
     created: 2023-04-07  expires: never       usage: SC  
     trust: ultimate      validity: ultimate
ssb  rsa3072/9F5DBAF4C83B7125
     created: 2023-04-07  expires: never       usage: E   
ssb* dsa2048/AD5EC538042CD3DE
     created: 2023-04-08  expires: never       usage: A  
[ultimate] (1). Alicia <Alicia@gmail.com>
[ultimate] (2)  Alice <Alice@gmail.com>
```

当然这是里还要 `save`，才会实际生效

```
root@v2:~/.gnupg# gpg -k A3455E25BA6751B4B551951CC6080149806CD3C3
pub   rsa3072 2023-04-07 [SC]
      A3455E25BA6751B4B551951CC6080149806CD3C3
uid           [ultimate] Alicia <Alicia@gmail.com>
uid           [ultimate] Alice <Alice@gmail.com>
sub   rsa3072 2023-04-07 [E]
sub   dsa2048 2023-04-08 [A]
```

### Change passphrase

例如修改 expiredate

```
root@v2:~# gpg --edit-key test
gpg (GnuPG) 2.2.19; Copyright (C) 2019 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Secret key is available.

sec  rsa3072/A623CF87362224FE
     created: 2023-04-06  expires: never       usage: SC  
     trust: ultimate      validity: ultimate
*** This key has been disabled
ssb  rsa3072/ED7BEAB890C8DDE9
     created: 2023-04-06  expires: never       usage: E   
[ultimate] (1). test

gpg> expire
Changing expiration time for the primary key.
Please specify how long the key should be valid.
         0 = key does not expire
      <n>  = key expires in n days
      <n>w = key expires in n weeks
      <n>m = key expires in n months
      <n>y = key expires in n years
Key is valid for? (0) 10
Key expires at Sun 16 Apr 2023 03:08:38 PM CST
Is this correct? (y/N) y

sec  rsa3072/A623CF87362224FE
     created: 2023-04-06  expires: 2023-04-16  usage: SC  
     trust: ultimate      validity: ultimate
*** This key has been disabled
ssb  rsa3072/ED7BEAB890C8DDE9
     created: 2023-04-06  expires: never       usage: E   
[ultimate] (1). test
gpg> save
```
查看对应的 GPG key expiredate 已经从 never 修改到了 t + 10
```
root@v2:~# gpg -k list test
gpg: checking the trustdb
gpg: marginals needed: 3  completes needed: 1  trust model: pgp
gpg: depth: 0  valid:   5  signed:   0  trust: 0-, 0q, 0n, 0m, 0f, 5u
gpg: next trustdb check due at 2023-04-16
pub   rsa3072 2023-04-06 [SC] [expires: 2023-04-16]
      8F8BD0829F6FBBD48716A0A9A623CF87362224FE
uid           [ultimate] test
sub   rsa3072 2023-04-06 [E]
```
### Remove passphrase
有一种比较的情况，就是需要移除 GPG key passphrase，即 passphrase 留空。在 GPG2.0 之前都是正常的，如果在 GPG2.1 之后必须使用 `--pinentry-mode loopback` 才能留空 passphrase，否则就会一直要求你输入 new passphrase

可以使用如下命令

```
root@v2:~# gpg --pinentry-mode loopback  --passwd 35A9B59E9FA8002D4617425821E13026C0DC7502
```

会提示 3 次 enter password，第 1 次输入原始的命令，第 2，3 次直接输入回车

修改完成后可以使用如下命令来校验

```
root@v2:~# gpg --dry-run --passwd tester
gpg (GnuPG) 2.2.19; Copyright (C) 2019 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
```

当然也可以使用 `--edit-key`

```
root@v2:~# gpg --pinentry-mode loopback  --edit-key 35A9B59E9FA8002D4617425821E13026C0DC7502
```

同样会提示 3 次 enter password，第 1 次输入原始的命令，第 2，3 次直接输入回车

## Delete Keys

如果需要删除 GPG key，需要先删除对应的 secret key 才能删除 public key
```
root@v2:~# gpg --delete-secret-keys 203F2D41E8C37B75C7F0074BC86BCE670DBE2517
gpg (GnuPG) 2.2.19; Copyright (C) 2019 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.


sec  rsa3072/C86BCE670DBE2517 2023-04-06 quick-gen

Delete this key from the keyring? (y/N) y
This is a secret key! - really delete? (y/N) y
root@v2:~# gpg --delete-keys 203F2D41E8C37B75C7F0074BC86BCE670DBE2517
gpg (GnuPG) 2.2.19; Copyright (C) 2019 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.


pub  rsa3072/C86BCE670DBE2517 2023-04-06 quick-gen

Delete this key from the keyring? (y/N) y
```
## Export keys
GPG 支持导出 公钥 和 私钥
### public key
导出公钥需要使用 `--export`，默认以 binary 格式输出到 stdout
```
root@v2:~# gpg --export tester
```
如果需要以明文的方式显示，可以使用 `--armor`
```
root@v2:~# gpg --armor --export tester
```
如果需要导入到指定文件，可以使用 `--out`
```
root@v2:~# gpg --out tester.gpg --export tester
```
当然 `--armor` 和 `--out` 也是可以一起使用的
```
root@v2:~# gpg --armor --out tester.gpg --export tester
```
### secret key
导出私钥需要使用 `--export-scret-keys`，和导出公钥的方式类似
```
root@v2:~# gpg --export-scret-keys tester
root@v2:~# gpg --armor --export-scret-keys tester
root@v2:~# gpg --out tester.gpg --export-scret-keys tester
root@v2:~# gpg --armor --out tester.gpg --export-scret-keys tester
```
## Import keys
### public key
将指定的公钥导入当前的 keyring 中
```
root@v2:~# gpg --import tester.pk
gpg: key A23605A6E42927AD: public key "tester (this is a comment) <tester@qq.com>" imported
gpg: Total number processed: 1
gpg:               imported: 1
root@v2:~# gpg -k
gpg: checking the trustdb
gpg: no ultimately trusted keys found
/root/.gnupg/pubring.kbx
------------------------
pub   rsa3072 2023-04-06 [SC] [expires: 2023-04-13]
      49801911B98422051F6AAA86A23605A6E42927AD
uid           [ unknown] tester (this is a comment) <tester@qq.com>
sub   rsa3072 2023-04-06 [E] [expires: 2023-04-13]

root@v2:~# gpg -K
```
这里可以看到私钥里是没有 tester 对应的信息的
### secret key
和导入公钥一样，私钥也可以被导入，同样通过 `--import` option
```
root@v2:~# gpg --import tester.sk
gpg: key A23605A6E42927AD: "tester (this is a comment) <tester@qq.com>" not changed
gpg: key A23605A6E42927AD: secret key imported
gpg: Total number processed: 1
gpg:              unchanged: 1
gpg:       secret keys read: 1
gpg:   secret keys imported: 1
root@v2:~# gpg -K
/root/.gnupg/pubring.kbx
------------------------
sec   rsa3072 2023-04-06 [SC] [expires: 2023-04-13]
      49801911B98422051F6AAA86A23605A6E42927AD
uid           [ unknown] tester (this is a comment) <tester@qq.com>
ssb   rsa3072 2023-04-06 [E] [expires: 2023-04-13]
```
### import from stdin
有一种比较特殊的方式，就是通过 stdin 导入 GPG key
[https://askubuntu.com/questions/1406741/how-to-import-gpg-secret-private-key-from-command-line-stdin](https://askubuntu.com/questions/1406741/how-to-import-gpg-secret-private-key-from-command-line-stdin)
只需要输入 `gpg --import` 然后使用 `ctrl` + `D` 回车即可 
## Generate revocation certificate
revocation certificate 是用于通告其他人对应的 GPG key 不再被使用了，但是对应被 revoked 的 GPG key 仍然可以被用来校验。之前用对应 GPG key 生成的 signatures，但是不能被用来 encrypt messages

可以通过 `--gen-revoke` 来实现
```
root@v2:~# gpg --gen-revoke CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8

sec  rsa3072/0FC2114BEEB4EBF8 2023-03-25 alice (this is a comment) <alice@yahoo.com>

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
ASCII armored output forced.
-----BEGIN PGP PUBLIC KEY BLOCK-----
Comment: This is a revocation certificate

iQG2BCABCgAgFiEEy2bUJKm+TcW5urT9D8IRS+606/gFAmQvspkCHQAACgkQD8IR
S+606/hCXQwAp7lZ+a5Fwb691BWcQIvNUDC3KuCIg3m7YygvrE3cPiblOuERxZF+
Yw9xt8ilLnG1GWqikcmwqZ4yGRGv4JihMvhR7ffV7QSpMLLHiZVX5rY2zCp1Oaks
JwMN7oqHW4YuHz1X6AUoyyIObDYy1zRd4B8E0aSN+MP8loMG+Zra3Wi2TlyuS9hc
950QtGrAlYH0bIZB87HqmZatANfKc04T0A1WyBDuLdvVqMRRBc3lCaQiR9UpDdTc
xXTOJr1I8UIT6rf6jY/PBuPzl4JMog/l/J8ChT/fRI9qQVGz/d8IplC/l8BGuxHE
Si3Pgruy3k+up50X3YglwM/bEAM8NhAddB5ACqh1i2SsbsNDygC/+IWb9QKLdk5P
IK5yitCXF8QFr5xYBBhGMnl7LqTf6mnqfhl/HJJMJEDJsFwzKI1/T9wJK/su1KRN
3hFRE5m1y5BZKAC/0qZZSY3DCGO5VBMYRw2er7/EFCYOhZnZ9edlwpIi1wJnBV2U
MqUCC9oKUqdV
=3Ed7
-----END PGP PUBLIC KEY BLOCK-----
Revocation certificate created.

Please move it to a medium which you can hide away; if Mallory gets
access to this certificate he can use it to make your key unusable.
It is smart to print this certificate and store it away, just in case
your media become unreadable.  But have some caution:  The print system of
your machine might store the data and make it available to others!
```
默认会输出到 stdout 中，可以使用 `--out` 或者 重定向 到指定文件
```
root@v2:~# gpg --out alice.revok --gen-revoke CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
#等价
root@v2:~# gpg --gen-revoke CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8 > alice.revok
```
### Revoke GPG key
有两种方式
#### 方法一
通过 `--edit-key` 中的 revkey
```
root@v2:~# gpg --edit-key CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
gpg (GnuPG) 2.2.19; Copyright (C) 2019 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Secret key is available.

sec  rsa3072/0FC2114BEEB4EBF8
     created: 2023-03-25  expires: never       usage: SC  
     trust: ultimate      validity: ultimate
ssb  rsa3072/35C73F8856F4FB14
     created: 2023-03-25  expires: never       usage: E   
[ultimate] (1). alice (this is a comment) <alice@yahoo.com>

gpg> revkey
Do you really want to revoke the entire key? (y/N) y
Please select the reason for the revocation:
  0 = No reason specified
  1 = Key has been compromised
  2 = Key is superseded
  3 = Key is no longer used
  Q = Cancel
Your decision? 0
Enter an optional description; end it with an empty line:
> 
Reason for revocation: No reason specified
(No description given)
Is this okay? (y/N) y

The following key was revoked on 2023-04-07 by RSA key 0FC2114BEEB4EBF8 alice (this is a comment) <alice@yahoo.com>
sec  rsa3072/0FC2114BEEB4EBF8
     created: 2023-03-25  revoked: 2023-04-07  usage: SC  
     trust: ultimate      validity: revoked
The following key was revoked on 2023-04-07 by RSA key 0FC2114BEEB4EBF8 alice (this is a comment) <alice@yahoo.com>
ssb  rsa3072/35C73F8856F4FB14
     created: 2023-03-25  revoked: 2023-04-07  usage: E   
[ revoked] (1). alice (this is a comment) <alice@yahoo.com>
```
查看状态
```
root@v2:~# gpg -k CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
pub   rsa3072 2023-03-25 [SC] [revoked: 2023-04-07]
      CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
uid           [ revoked] alice (this is a comment) <alice@yahoo.com>
```
#### 方法二
先生成 revocation certificate 用于撤销 GPG key ( priamry key 或者 subkey )，可以通过 `--gen-revoke` 来实现
```
root@v2:~# gpg --gen-revoke CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8

sec  rsa3072/0FC2114BEEB4EBF8 2023-03-25 alice (this is a comment) <alice@yahoo.com>

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
ASCII armored output forced.
-----BEGIN PGP PUBLIC KEY BLOCK-----
Comment: This is a revocation certificate

iQG2BCABCgAgFiEEy2bUJKm+TcW5urT9D8IRS+606/gFAmQvspkCHQAACgkQD8IR
S+606/hCXQwAp7lZ+a5Fwb691BWcQIvNUDC3KuCIg3m7YygvrE3cPiblOuERxZF+
Yw9xt8ilLnG1GWqikcmwqZ4yGRGv4JihMvhR7ffV7QSpMLLHiZVX5rY2zCp1Oaks
JwMN7oqHW4YuHz1X6AUoyyIObDYy1zRd4B8E0aSN+MP8loMG+Zra3Wi2TlyuS9hc
950QtGrAlYH0bIZB87HqmZatANfKc04T0A1WyBDuLdvVqMRRBc3lCaQiR9UpDdTc
xXTOJr1I8UIT6rf6jY/PBuPzl4JMog/l/J8ChT/fRI9qQVGz/d8IplC/l8BGuxHE
Si3Pgruy3k+up50X3YglwM/bEAM8NhAddB5ACqh1i2SsbsNDygC/+IWb9QKLdk5P
IK5yitCXF8QFr5xYBBhGMnl7LqTf6mnqfhl/HJJMJEDJsFwzKI1/T9wJK/su1KRN
3hFRE5m1y5BZKAC/0qZZSY3DCGO5VBMYRw2er7/EFCYOhZnZ9edlwpIi1wJnBV2U
MqUCC9oKUqdV
=3Ed7
-----END PGP PUBLIC KEY BLOCK-----
Revocation certificate created.

Please move it to a medium which you can hide away; if Mallory gets
access to this certificate he can use it to make your key unusable.
It is smart to print this certificate and store it away, just in case
your media become unreadable.  But have some caution:  The print system of
your machine might store the data and make it available to others!
```
默认会输出到 stdout 中，可以使用 `--out` 或者 重定向 到指定文件
```
root@v2:~# gpg --out alice.revok --gen-revoke CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
#等价
root@v2:~# gpg --gen-revoke CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8 > alice.revok
```
导入对应的 vocation certificate
```
root@v2:~# gpg --import alice.revok 
gpg: key 0FC2114BEEB4EBF8: "alice (this is a comment) <alice@yahoo.com>" revocation certificate imported
gpg: Total number processed: 1
gpg:    new key revocations: 1

root@v2:~# gpg -k alice@yahoo.com
pub   rsa3072 2023-03-25 [SC] [revoked: 2023-04-07]
      CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
uid           [ revoked] alice (this is a comment) <alice@yahoo.com>
```
