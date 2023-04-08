# GPG - keyserver

ref
https://www.gnupg.org/gph/en/manual/x457.html

[https://keys.openpgp.org/about/usage](https://www.gnupg.org/gph/en/manual/x457.html)

https://wiki.archlinux.org/title/GnuPG#Use_a_keyserver

https://en.wikipedia.org/wiki/Key_server_(cryptographic)#Keyserver_examples

## Digest

GPG 支持将公钥上传到 key server 便于用户查询，也支持从 key server 将对应的公钥下载到本地，以便于方便和其他用户进行加密通信

只要上传到一个 keyserver ，一般会自动和其他 keyserver 同步

可用的公钥服务器主要有如下几个

- https://keyserver.ubuntu.com/

  上传 keys 后，不需要校验邮箱地址，不能被删除。默认使用该 keyserver，可以使用 `--keyserver` 或者是 `dirmngr.conf` 中的 `keyserver` 替换成其他的 keyserver

- https://keys.mailvelope.com/

  上传 keys 后，需要校验邮箱地址，才能被正常使用以及和其他 keyserver 同步，可以被删除

- https://keys.openpgp.org/

  上传 keys 后，需要校验邮箱地址，才能被正常使用以及和其他 keyserver 同步，可以被删除

## 上传公钥

如果我们需要上传公钥到 keyserver 我们需要使用 `--send-keys`

```
root@v2:/home/ubuntu# gpg -v --send-keys 49801911B98422051F6AAA86A23605A6E42927AD
gpg: sending key A23605A6E42927AD to hkps://keys.openpgp.org
```

这里可以看到对应的公钥上传具体那个 keyserver，这样我们就可以在对 keyserver 站点上看到

如果需要上传到指定的 keyserver 可以使用 `--keyserver`

```
root@v2:/home/ubuntu# gpg --keyserver hkps://keyserver.ubuntu.com --send-keys FFF05AF41C85F603A6DA2704B19B91237DAE029E
gpg: sending key B19B91237DAE029E to https://keyserver.ubuntu.com
```

注意这里 keyserver 尽量使用 hkps 协议

## 更新公钥

如果需要从 keyserver 同步本地的公钥，需要使用 `--refresh` 。可以将其想象成类似 `git pull` 的操作

## 查询公钥

可以使用 `--search-keys` 来查询指定的公钥

```
root@v2:/home/ubuntu# gpg --search-keys 49801911B98422051F6AAA86A23605A6E42927AD
gpg: data source: https://keys.openpgp.org:443
(1)       3072 bit RSA key A23605A6E42927AD, created: 2023-04-06
Keys 1-1 of 1 for "49801911B98422051F6AAA86A23605A6E42927AD".  Enter number(s), N)ext, or Q)uit > 
```

如果需要指定查询那个 keyserver 可以使用

```
[vagrant@localhost ~]$ gpg --keyserver hkps://keys.openpgp.org --search-keys FFF05AF41C85F603A6DA2704B19B91237DAE029E
gpg: searching for "FFF05AF41C85F603A6DA2704B19B91237DAE029E" from hkps server keys.openpgp.org
(1)       3072 bit RSA key 7DAE029E, created: 2023-04-08
Keys 1-1 of 1 for "FFF05AF41C85F603A6DA2704B19B91237DAE029E".  Enter number(s), N)ext, or Q)uit > q
```

其中的 data source 表明是从那个 keyserver 同步过来的，同样的这里也尽量使用 hkps 协议

## 导入公钥

导入公钥可以通过 `--recv-keys` 来实现

```
[vagrant@localhost ~]$ gpg -v --keyserver hkps://keys.openpgp.org --recv-keys FFF05AF41C85F603A6DA2704B19B91237DAE029E
gpg: requesting key 7DAE029E from hkps server keys.openpgp.org
Comment: FFF0 5AF4 1C85 F603 A6DA  2704 B19B 9123 7DAE 029E
gpg: armor header: 
gpg: pub  3072R/7DAE029E 2023-04-08  
gpg: key 7DAE029E: no user ID
gpg: Total number processed: 1

[vagrant@localhost ~]$ gpg -k
/home/vagrant/.gnupg/pubring.gpg
--------------------------------
pub   3072R/7DAE029E 2023-04-08 [expires: 2023-04-09]
uid                  Aroul (this key is for testing purpose) <823419326@qq.com>
sub   3072R/E5004D68 2023-04-08 [expires: 2023-04-09]
```

也可以在查询的使用实现

```
oot@v2:~/.gnupg# gpg --search-keys kikochz@163.com
gpg: data source: https://162.213.33.9:443
(1)     kikochz <kikochz@163.com>
          3072 bit RSA key 744491CF5CB7A082, created: 2020-12-11
(2)       2048 bit RSA key 5264F01557666819, created: 2020-12-09
(3)     cyberpelican (let the word out) <kikochz@163.com>
          3072 bit RSA key CC5EEBA5391FCE0A, created: 2020-12-11
(4)     kikochz (kiko) <kikochz@163.com>
          2048 bit RSA key AA96E1168DCFA23D, created: 2020-12-09
(5)     cyberpelican (let the word out) <kikochz@163.com>
          3072 bit RSA key 8A426D3F09705CFF, created: 2020-12-11
Keys 1-5 of 5 for "kikochz@163.com".  Enter number(s), N)ext, or Q)uit > 2
gpg: key 5264F01557666819: public key "kikochz (test for gpg) <kikochz@163.com>" imported
gpg: Total number processed: 1
gpg:               imported: 1
```

在这里输入需要导入公钥的序号，就会自动导入公钥

```
root@v2:~/.gnupg# gpg -k kikochz
pub   rsa2048 2020-12-09 [SC] [expired: 2021-02-17]
      5ABA1A12C664473FBDC8990F5264F01557666819
uid           [ expired] kikochz (test for gpg) <kikochz@163.com>
```

## 撤销证书

虽然证书不能被有效的删除，但是可以被 revoked

首先需要先生成 revoke 证书

```
root@v2:~/.gnupg# gpg --out aroul.reovk --gen-revoke FFF05AF41C85F603A6DA2704B19B91237DAE029E

sec  rsa3072/B19B91237DAE029E 2023-04-08 Aroul (this key is for testing purpose) <823419326@qq.com>

Create a revocation certificate for this key? (y/N) y
Please select the reason for the revocation:
  0 = No reason specified
  1 = Key has been compromised
  2 = Key is superseded
  3 = Key is no longer used
  Q = Cancel
(Probably you want to select 1 here)
Your decision? 
Enter an optional description; end it with an empty line:
> 
Reason for revocation: Key has been compromised
(No description given)
Is this okay? (y/N) y
Revocation certificate created.

Please move it to a medium which you can hide away; if Mallory gets
access to this certificate he can use it to make your key unusable.
It is smart to print this certificate and store it away, just in case
your media become unreadable.  But have some caution:  The print system of
your machine might store the data and make it available to others!
```

导入 revoke 证书

```
root@v2:~/.gnupg# gpg --import aroul.revok 
gpg: key B19B91237DAE029E: "Aroul (this key is for testing purpose) <823419326@qq.com>" revocation certificate imported
gpg: Total number processed: 1
gpg:    new key revocations: 1
gpg: marginals needed: 3  completes needed: 1  trust model: pgp
gpg: depth: 0  valid:   2  signed:   0  trust: 0-, 0q, 0n, 0m, 0f, 2u
gpg: next trustdb check due at 2023-04-09
```

查看本地公钥是否被 revoked

```
root@v2:~/.gnupg# gpg -k Aroul
pub   rsa3072 2023-04-08 [SC] [revoked: 2023-04-08]
      FFF05AF41C85F603A6DA2704B19B91237DAE029E
uid           [ revoked] Aroul (this key is for testing purpose) <823419326@qq.com>
```

上传 key 到 keyserver，过一段时间就会同步

```
root@v2:~/.gnupg# gpg --send-keys FFF05AF41C85F603A6DA2704B19B91237DAE029E
gpg: sending key B19B91237DAE029E to hkp://keyserver.ubuntu.com
```

假设这时其他用户使用 `--search-keys` 查询 key 时，并不能看到 revoked 字段

```
[vagrant@localhost ~]$ gpg --keyserver keyserver.ubuntu.com --search-keys fff05af41c85f603a6da2704b19b91237dae029e
gpg: searching for "fff05af41c85f603a6da2704b19b91237dae029e" from hkp server keyserver.ubuntu.com
(1)     Aroul (this key is for testing purpose) <823419326@qq.com>
          3072 bit RSA key 7DAE029E, created: 2023-04-08
Keys 1-1 of 1 for "fff05af41c85f603a6da2704b19b91237dae029e".  Enter number(s), N)ext, or Q)uit > 1
gpg: requesting key 7DAE029E from hkp server keyserver.ubuntu.com
gpg: key 7DAE029E: public key "Aroul (this key is for testing purpose) <823419326@qq.com>" imported
gpg: Total number processed: 1
gpg:               imported: 1  (RSA: 1)
```

需要导入 key 才能看到

```
[vagrant@localhost ~]$ gpg -k
/home/vagrant/.gnupg/pubring.gpg
--------------------------------
pub   3072R/7DAE029E 2023-04-08 [revoked: 2023-04-08]
uid                  Aroul (this key is for testing purpose) <823419326@qq.com>
```

这个逻辑完全是正确，因为 revoke 只是针对私钥的，公钥任然可以用做签名校验
