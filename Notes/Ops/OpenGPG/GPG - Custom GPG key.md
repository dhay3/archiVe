# GPG - Custom GPG key

首先使用 `--full-gen-key` 来生成初始的 GPG key，并且不想要 passphrase

```
root@v2:~# gpg --pinentry-mode loopback  --full-gen-key
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
Key is valid for? (0) 
Key does not expire at all
Is this correct? (y/N) y

GnuPG needs to construct a user ID to identify your key.

Real name: c4lice
Email address: c4lice@gmail.com
Comment: 
You selected this USER-ID:
    "c4lice <c4lice@gmail.com>"

Change (N)ame, (C)omment, (E)mail or (O)kay/(Q)uit? o
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
gpg: key 0F78C3000DE02E64 marked as ultimately trusted
gpg: revocation certificate stored as '/root/.gnupg/openpgp-revocs.d/5F833DD99960D231758B75630F78C3000DE02E64.rev'
public and secret key created and signed.

pub   rsa3072 2023-04-08 [SC]
      5F833DD99960D231758B75630F78C3000DE02E64
uid                      c4lice <c4lice@gmail.com>
sub   rsa3072 2023-04-08 [E]
```

修改 GPG key，添加 subkey，这里使用 `--pinentry-mode` 同样是为了在生成 subkey 时不使用 passphrase

```
root@v2:~# gpg --pinentry-mode loopback --edit-key c4lice
gpg> addkey
Please select what kind of key you want:
   (3) DSA (sign only)
   (4) RSA (sign only)
   (5) Elgamal (encrypt only)
   (6) RSA (encrypt only)
  (14) Existing key from card
Your selection? 4
RSA keys may be between 1024 and 4096 bits long.
What keysize do you want? (3072) 
Requested keysize is 3072 bits
Please specify how long the key should be valid.
         0 = key does not expire
      <n>  = key expires in n days
      <n>w = key expires in n weeks
      <n>m = key expires in n months
      <n>y = key expires in n years
Key is valid for? (0)
Key does not expire at all
Is this correct? (y/N) y
Really create? (y/N) y
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
```

修改 key 的使用方式，主密钥用于 Sign 和 Certify，Subkey1 用于 encrypt，Subkey2 用于 authentication

```
gpg> key 2

sec  rsa3072/0F78C3000DE02E64
     created: 2023-04-08  expires: never       usage: SC  
     trust: ultimate      validity: ultimate
ssb  rsa3072/51672B06ABD73D06
     created: 2023-04-08  expires: never       usage: E   
ssb* rsa3072/E0F78F9BC03150DF
     created: 2023-04-08  expires: never       usage: S   
[ultimate] (1). c4lice <c4lice@gmail.com>

gpg> change-usage
Changing usage of a subkey.

Possible actions for a RSA key: Sign Encrypt Authenticate 
Current allowed actions: Sign 

   (S) Toggle the sign capability
   (E) Toggle the encrypt capability
   (A) Toggle the authenticate capability
   (Q) Finished

Your selection? A

Possible actions for a RSA key: Sign Encrypt Authenticate 
Current allowed actions: Sign Authenticate 

   (S) Toggle the sign capability
   (E) Toggle the encrypt capability
   (A) Toggle the authenticate capability
   (Q) Finished

Your selection? S

Possible actions for a RSA key: Sign Encrypt Authenticate 
Current allowed actions: Authenticate 

   (S) Toggle the sign capability
   (E) Toggle the encrypt capability
   (A) Toggle the authenticate capability
   (Q) Finished

Your selection? Q

sec  rsa3072/0F78C3000DE02E64
     created: 2023-04-08  expires: never       usage: SC  
     trust: ultimate      validity: ultimate
ssb  rsa3072/51672B06ABD73D06
     created: 2023-04-08  expires: never       usage: E   
ssb* rsa3072/E0F78F9BC03150DF
     created: 2023-04-08  expires: never       usage: A   
[ultimate] (1). c4lice <c4lice@gmail.com>
```

添加其他邮箱信息

> 如果 Key 需要用于 Github GPG Signing 邮箱需要和 `.gitconfig` 中的相同

```
gpg> adduid
Real name: c4lice
Email address: c4lice@yahoo.com
Comment: 
You selected this USER-ID:
    "c4lice <c4lice@yahoo.com>"

Change (N)ame, (C)omment, (E)mail or (O)kay/(Q)uit? o

sec  rsa3072/0F78C3000DE02E64
     created: 2023-04-08  expires: never       usage: SC  
     trust: ultimate      validity: ultimate
ssb  rsa3072/51672B06ABD73D06
     created: 2023-04-08  expires: never       usage: E   
ssb* rsa3072/E0F78F9BC03150DF
     created: 2023-04-08  expires: never       usage: A   
[ultimate] (1)  c4lice <c4lice@gmail.com>
[ unknown] (2). c4lice <c4lice@yahoo.com>
gpg>save
```

查看新生成的 key

```
root@v2:~# gpg -k c4lice
pub   rsa3072 2023-04-08 [SC]
      5F833DD99960D231758B75630F78C3000DE02E64
uid           [ultimate] c4lice <c4lice@gmail.com>
uid           [ultimate] c4lice <c4lice@yahoo.com>
sub   rsa3072 2023-04-08 [E]
sub   rsa3072 2023-04-08 [A]
```

创建完之后第一件事就是备份

```
root@v2:~# gpg --out c4lice.ssk --armor --export-secret-subkeys c4lice 
root@v2:~# gpg --out c4lice.sk --armor --export-secret-keys c4lice 
root@v2:~# gpg --out c4lice.pk --armor --export c4lice
```

## Use GPG key for SSH

如果想要使用 GPG key 用于 SSH 参考如下

让 gpg-agent 支持 SSH

```
echo enable-ssh-support >> $HOME/.gnupg/gpg-agent.conf
```

在 `.zshrc` 中添加

```
unset SSH_AGENT_PID
if [ "${gnupg_SSH_AUTH_SOCK_by:-0}" -ne $$ ]; then
    export SSH_AUTH_SOCK="$(gpgconf --list-dirs agent-ssh-socket)"
fi
export GPG_TTY=${TTY:-"$(tty)"}
gpg-connect-agent updatestartuptty /bye >/dev/null
```

查看 subkey keygrip(用于 Authentication 的 key)

```
base) 0x00 in ~ λ gpg --list-keys --with-keygrip
/home/0x00/.gnupg/pubring.kbx
-----------------------------
pub   rsa4096 2021-01-14 [SC]
      A5DD905196EF3973280DA13CB965BC5D279F42ED
      Keygrip = A5786528100EA37404519F98D729D155982C34D8
uid           [ unknown] canihavesomecoffee (GPG key for signing GitHub commits) <git@canihavesome.coffee>
sub   rsa4096 2021-01-14 [E]
      Keygrip = D6208641AD1463510438CC57420A7F75E98E9622

pub   rsa3072 2023-11-08 [SC]
      0D1FBFEA9F499F6A6A01A49D1D8AC28A0E990946
      Keygrip = 71B7748D4CB25884DFD7B7D7B4F875ED7F2BAAB1
uid           [ultimate] HuanzhangCheng (For Github GPG Signing) <62749885+dhay3@users.noreply.github.com>
uid           [ultimate] HuanzhangCheng (For Signing) <hostlockdown@gmail.com>
uid           [ultimate] HuanzhangCheng (For Authentication) <kikochz@163.com>
sub   rsa3072 2023-11-08 [E]
      Keygrip = AE09B1EB7B6CE727315D5B79254E80C313CE4280
sub   rsa3072 2023-11-08 [A]
      Keygrip = 22FE940D6D52158063C18B637976B1666C841072
```

将 keygrip 导入到允许使用 SSH 的配置

```
echo 22FE940D6D52158063C18B637976B1666C841072 >> ~/.gnupg/sshcontrol
```

检查 key 是否已经加入到 SSH 的认证列表

```
(base) 0x00 in ~ λ ssh-add -l           
3072 SHA256:v3Mr+1ILxlWlxrjlo/86hGRTK4dUt85KHVko7sKDk4A (none) (RSA)
```

如果用于 Github 还可以使用如下方式验证(需要将公钥上传至 Github)

```
(base) 0x00 in ~ λ ssh -T git@github.com
Hi dhay3! You've successfully authenticated, but GitHub does not provide shell access.
```

记得备份

```
#导出 GPG 的主秘钥
gpg --armor --export-secret-keys --out master-secret-key.gpg host
#导出 GPG 的副秘钥
gpg --armor --export-secret-subkeys --out sub-secret-key.gpg host
#导出 用于签名和认证的公钥
gpg --armor --export --out public-key.gpg host
#导出 用于 SSH 认证的公钥
gpg --armor --export-ssh-key --out ssh-public-key.gpg host
```

**references**

[^1]:https://gist.github.com/mcattarinussi/834fc4b641ff4572018d0c665e5a94d3
[^2]:https://stackoverflow.com/questions/17846529/could-not-open-a-connection-to-your-authentication-agent
