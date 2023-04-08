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

导出 ssh 公钥

```
root@v2:~# gpg --out c4lice.spk --armor --export-ssh-key c4lice 
```

