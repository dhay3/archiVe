ref
[https://lists.gnupg.org/pipermail/gnupg-users/2011-January/040437.html](https://lists.gnupg.org/pipermail/gnupg-users/2011-January/040437.html)
[https://unix.stackexchange.com/questions/31996/how-are-the-gpg-usage-flags-defined-in-the-key-details-listing](https://unix.stackexchange.com/questions/31996/how-are-the-gpg-usage-flags-defined-in-the-key-details-listing)
[https://superuser.com/questions/1371088/what-do-ssb-and-sec-mean-in-gpgs-output](https://superuser.com/questions/1371088/what-do-ssb-and-sec-mean-in-gpgs-output)
[https://wiki.debian.org/Subkeys](https://wiki.debian.org/Subkeys)
[https://www.void.gr/kargig/blog/2013/12/02/creating-a-new-gpg-key-with-subkeys/](https://www.void.gr/kargig/blog/2013/12/02/creating-a-new-gpg-key-with-subkeys/)
## Generate keys
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
在 GPG 2.0 之后只有使用了 `--batch` ，`--passphrase` 才会生效。如果在 GPG 2.1 之后还需要使用 `--pinentry-mode loopback`
即如果你不想对 GPG key 设置密码，就必须使用 `--pinentry-mode loopback` 和 `--batch`
那么下面这种方式就是错误的，因为 `--passphrase` 实际不生效
```
root@v2:/home/ubuntu# gpg --passphrase '' --quick-gen-key test
```
需要使用下面这种方式
```
root@v2:/home/ubuntu# gpg --passphrase '' --pinentry-mode loopback --batch --quick-gen-key test
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
- 第二行显示秘钥的 Hash 值(不特指公钥或者私钥), 可用做 user-id 唯一标识符
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
- 第二行显示秘钥的 Hash 值(不特指公钥或者私钥), 可用做 user-id 唯一标识符
- 第三行 uid 显示秘钥绑定的用户，以及邮箱 和 comment
- `ssb` 表示 secrect subkey，其中的 `[E]` 表示当前 secret subkey 的用途，E=encryption
### subkeys
subkey 是 GPG key 的一种扩展，挂在 primary keys 下，可以单独 revoke subkey，而不需要 revoke primary key
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

- `quit`

退出 edit 界面

- `save`

保存当前的修改并退出，如果不保存修改会被 discarded

- `addkey`

添加一个 额外的 subkey

- `key <subkey hash>`

选择 subkey

- `delkey`

删除指定的 subkey，需要先试用 key 选择 subkey

- `list`

显示当面操作的 GPG key

- `expire`

修改 GPG key 有效的日期

- `passwd`

修改密码
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

