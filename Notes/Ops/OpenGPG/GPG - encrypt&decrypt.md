# GPG encrypt & decrypt

ref
[https://www.gnupg.org/gph/en/manual/x110.html](https://www.gnupg.org/gph/en/manual/x110.html)

GPG 可以对文件签名，当然也可以使用加密和解密，因为两者只是针对加密方向不同的两种说法。
加密和解密都可以通过 `--out` 或者是 重定向 将内容输出到指定的文件
## Asymmetric Enryption/Decryption
如果 Bobby 需要和 Alice 发送加密信息，Bobby 就需要使用 Alice 的公钥对信息加密，那么 Alice 就可以通过自己的私钥对 Bobby 发送的加密信息解谜
当然如果 Alice 需要发送加密信息给 Bobby，Alice 需要使用 Bobby 的公钥对信息加密，那么 Bobby 就可以通过自己的私钥对 Alice 发送给 Bobby 的加密信息解谜
### Encryption
假设你现在是 Bobboy 要给 Alice 发送如下消息
```
λ ~/ cat message 
this is a secret message from bobby to alice
```
如果需要加密，就需要获取 Alice 的公钥，并导入到 gpg keyring。假设 Alice 已经偷偷告诉你了 ，这里通过 `gpg --armor --export alice`来明文显示 alice 公钥，然后复制直接导入，实际可以通过上传到 keyserver 然后拉取对应的信息或者其他方式传输 
那么你就可以使用如下的命令对 message 文件进行加密
```
λ ~/ gpg -v --local-user bobby -r Alice  -e message
gpg: using PGP trust model
gpg: using subkey C83B7125 instead of primary key 806CD3C3
gpg: C83B7125: There is no assurance this key belongs to the named user

pub  3072R/C83B7125 2023-04-07 Alice <Alice@gmail.com>
 Primary key fingerprint: A345 5E25 BA67 51B4 B551  951C C608 0149 806C D3C3
      Subkey fingerprint: 23FD 983F D6F4 FE3C 45EC  EA3F 9F5D BAF4 C83B 7125

It is NOT certain that the key belongs to the person named
in the user ID.  If you *really* know what you are doing,
you may answer the next question with yes.

Use this key anyway? (y/N) y
gpg: reading from `message'
gpg: writing to `message.gpg'
gpg: RSA/AES256 encrypted for: "C83B7125 Alice <Alice@gmail.com>"
```
上述命令 `-r | --recipient` 表示接受信息人，需要指定 user-id。默认会生成一个 对应文件名的`.gpg` 后缀的加密二进制文件。当然也可以使用 `--armor` 来生产明文的加密文件
```
λ ~/ gpg -v --local-user bobby -r Alice  --armor -e message
gpg: using PGP trust model
gpg: using subkey C83B7125 instead of primary key 806CD3C3
gpg: C83B7125: There is no assurance this key belongs to the named user

pub  3072R/C83B7125 2023-04-07 Alice <Alice@gmail.com>
 Primary key fingerprint: A345 5E25 BA67 51B4 B551  951C C608 0149 806C D3C3
      Subkey fingerprint: 23FD 983F D6F4 FE3C 45EC  EA3F 9F5D BAF4 C83B 7125

It is NOT certain that the key belongs to the person named
in the user ID.  If you *really* know what you are doing,
you may answer the next question with yes.

Use this key anyway? (y/N) y
gpg: reading from `message'
gpg: writing to `message.asc'
gpg: RSA/AES256 encrypted for: "C83B7125 Alice <Alice@gmail.com>"                                                                                                            
```
上述命令会生成一个 `.asc` 后缀的明文文件
```
λ ~/ cat message.asc 
-----BEGIN PGP MESSAGE-----
Version: GnuPG v2.0.22 (GNU/Linux)

hQGMA59duvTIO3ElAQv+I2VQ7MfvHYgv6UKHMypQ51GAQ2WTZeaKsWFgEFVIpfwa
oIktpzygLjMdV/wBORM58kDW2vIVXownurgROVIISJur9Re3L26lssppk0ievWdo
1+yRn0nE9i33CmtDo2LEMBVKq/+0THYlFpUt0xeHafZ9ug1hjowQ8LZoFIu21u6X
sqwS6otl//0jINRNbzkg9eYe6c+3l+1l6D+UADjcoQ6DqCU6tIVkYUfwImXU8yvL
JeOwfbhD/1Dm0X4TbXRDEF5gsycGezXr/sU3oXWRBgS0eKEnDgx1ukzi+V1sRhCL
kAwyEWA0RjeKZZUJlaYppv59iXKqug47UrKvhjoh4e465F+n/I6l7zPmlqDm4o+4
qLR9n1EQx6JbSyBXfyMMJ82l7IEnV4PinH6m+kiAKa2fOBpClTP/W83waAIjjaCU
4HT6fnkfkGmP+ZJbWzM38POzEPdVhId+c3g6airS8jbHLv86caa0lfnbeHSRNYm0
0v7FjHyKoHMZHx1lyWgi0mgByFqUOx11bxyhksiYUOj95Msf4TyCtf9pC+VvPFuu
quq+8cucZW57vMu218aL7oPwkrpYzMjw5fhp3ckg9KChNvobHLPBPK14rxZ6lA2M
kL40mSPixaKEgXlA/oTWLsn+oEm8PMjZKQ==
=uS12
-----END PGP MESSAGE-----
```
当时上面的内容其实并不能证明是 Bobby 发送给 Alice 的，因为谁都可以轻易的获取的 Alice 的公钥。所以也可以和 `--sign` 一起使用加密同时对文件使用 Bobby 的私钥对文件签名
```
λ ~/ gpg -v --local-user bobby -r Alice  --armor --sign  -e message

You need a passphrase to unlock the secret key for
user: "Bobby <bobby@gmail.com>"
2048-bit RSA key, ID 687EE529, created 2023-04-07

gpg: using PGP trust model
gpg: using subkey C83B7125 instead of primary key 806CD3C3
gpg: C83B7125: There is no assurance this key belongs to the named user

pub  3072R/C83B7125 2023-04-07 Alice <Alice@gmail.com>
 Primary key fingerprint: A345 5E25 BA67 51B4 B551  951C C608 0149 806C D3C3
      Subkey fingerprint: 23FD 983F D6F4 FE3C 45EC  EA3F 9F5D BAF4 C83B 7125

It is NOT certain that the key belongs to the person named
in the user ID.  If you *really* know what you are doing,
you may answer the next question with yes.

Use this key anyway? (y/N) y
File `message.asc' exists. Overwrite? (y/N) y
gpg: writing to `message.asc'
gpg: RSA/AES256 encrypted for: "C83B7125 Alice <Alice@gmail.com>"
gpg: RSA/SHA512 signature from: "687EE529 Bobby <bobby@gmail.com>"
```
上述命令会生产一个`.asc` 的加密文件包含签名
注意这里不能使用 `--clearsign` 和 `--detach-sign` 
### Decryption
如果加密的文件没有携带签名，可以使用如下命令来解密。那么会显示如下内容
```
root@v2:~# gpg -d message.asc 
gpg: encrypted with 3072-bit RSA key, ID 9F5DBAF4C83B7125, created 2023-04-07
      "Alice <Alice@gmail.com>"
this is a secret message from bobby to alice
```
如果加密的文件包含 Bobby 的签名，那么同样需要导入 Bobby 的公钥，否则不能正常校验签名。如果签名正常会显示如下内容
```
root@v2:~# vim message.asc
root@v2:~# gpg -d message.asc 
gpg: encrypted with 3072-bit RSA key, ID 9F5DBAF4C83B7125, created 2023-04-07
      "Alice <Alice@gmail.com>"
this is a secret message from bobby to alice
gpg: Signature made Fri 07 Apr 2023 05:02:04 PM CST
gpg:                using RSA key AE075F8C687EE529
gpg: Good signature from "Bobby <bobby@gmail.com>" [unknown]
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: 9CF6 8721 5158 4CAD 207A  F4A9 AE07 5F8C 687E E529
```
## Symmetric Encryption/Decryption
GPG 在创建 key 时，虽然会生成非对称秘钥，但是也可以使用对称秘钥来加密解密。这样就不需要获取对应的公钥了，但是这种方式是不安全的
### Encryption
需要使用 `-c | --symmetric` 来指定表示使用 symmetric 加密的方式
```
λ ~/ gpg -v  --armor --symmetric message                                
gpg: using cipher CAST5
File `message.asc' exists. Overwrite? (y/N) y
gpg: writing to `message.asc'
```
使用上述命令的过程中会要求 Bobby 输入一个用于对称加密的秘钥
```
λ ~/ cat message.asc 
-----BEGIN PGP MESSAGE-----
Version: GnuPG v2.0.22 (GNU/Linux)

jA0EAwMC8x0BBytPGBrkyUPGlWHW5yxfbMndYYPRvkaPIq8DotW5i2dqlEcLtDbs
J/FOPR0oY24xRY/5tHfBiJIwGMyckhZ78czFJRbF9pzvGnOI
=NgMp
-----END PGP MESSAGE-----
```
### Decryption
对称加密和非对称加密一样，可以直接使用 `--decrypt` 对文件解密
```
root@v2:~# gpg --decrypt message.asc
gpg: CAST5 encrypted data
gpg: encrypted with 1 passphrase
this is a secret message from bobby to alice
gpg: WARNING: message was not integrity protected
gpg: Hint: If this message was created before the year 2003 it is
     likely that this message is legitimate.  This is because back
     then integrity protection was not widely used.
gpg: Use the option '--ignore-mdc-error' to decrypt anyway.
gpg: decryption forced to fail!
```
使用上述命令，会要求 Alice 输入对应的对称秘钥
