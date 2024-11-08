# SSH ssh-keygen

ref

https://man.openbsd.org/ssh-keygen

## Digest

OpenSSH authentication key utility

`ssh-keygen` 用于 生成，管理以及转换 authentication keys for ssh（秘钥对）

在 version 7.6 之后默认使用 SSH2 生成 key，如果不指定生成的算法默认使用 RSA （RSA 原本是有专利的，但是专利过期了所以被开源使用）

### Certificates

`ssh-keygen` 还可以用于 signing of keys to produce certificates that may be used for user or host authentication ( 证书的签发 )

`ssh-keygen` 和 `openssl` 中默认使用 x.509 certificates 不一样，`ssh-keygen` 更加简单

`ssh-keygen` 支持 2 种证书：user 和 host

## Optional args

### gernal args

- `-f filename`

  指定需要使用 key，一般和其他参数一起使用。==可以手动指定该参数，就不会进入交互模式==

- `-v`

  verbose mode, multiple `-v` options increase the verbosity, the maximum is 3

### miscellaneous args

-  `-F hostname | [hostname]:port`

  在 known_hosts 中查看指定的 hostname 的

- `-R hostname | [hostname]:port`

  remove all keys beloging to the specified hostname

- `-H`

  将 known_hosts 中的内容 hash，文件中的原始内容会被写入到 known_hosts.old。==`ssh` 和 `sshd` 同样还是能读取 hash 后的内容==

### keys showing args

- `-l`

  show fingerprint of specified ==public key(如果输入的是私钥也会显示对应公钥的指纹)== file

  可以和 `-v` 一起使用显示 visual fingerprint ASCII art

  ```
  #输入私钥
  [root@localhost .ssh]# ssh-keygen -lf id_rsa
  2048 SHA256:AODQM4vKYPrJbJd11sxGDH4IQ+JJUxswZNxoWRuKkZk 
  #输入公钥
  root@localhost.localdomain (RSA)
  [root@localhost .ssh]# ssh-keygen -lf id_rsa.pub
  2048 SHA256:AODQM4vKYPrJbJd11sxGDH4IQ+JJUxswZNxoWRuKkZk root@localhost.localdomain (RSA)
  ```

- `-E fingerprint_hash`

  指定显示 fingerprints 使用的 hash 算法，可以是 md5 或者 sha256，默认 sha256

- `-e`

  ==输入私钥输出对应的公钥内容==

### keys creation args

- `-t dsa | ecdsa | ecdsa-sk | ed 25519 | ed25519-sk | rsa`

  指定生成秘钥对的算法，如果不指定默认使用 rsa

- `-b bits`

  指定生成 key 的加密 bit。如果加密算法是 RSA 范围在 `[1024,3072]`，如果是 DSA 必须是 1024，如果是 ECSDSA 可以是 256，384，521。如果是 ECDSA-SK, Ed25519, Ed25519-SK 改参数会被忽略 ( 因为加密 bit 固定 )

- `-C comment`

  为 key 添加 comment

### keys modification args

- `-N new_passphrase`

  修改秘钥的 passphrase

- `-p`

  以 interactive 的方式修改 passphrase 

- `-c`

  修改 key comment


- `-m key_format`

### keys signing args

- `-I certificate_identity`

  specify the key identity when singing a public key

## Examples

