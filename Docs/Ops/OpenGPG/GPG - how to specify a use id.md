# GPG - how to specify a use id

GPG 中很多参数都需要指定 User ID，User ID 非一个逻辑字段，而是一个集合。允许的格式如下 

## By key ID

```
root@v2:/home/ubuntu# gpg -k
/root/.gnupg/pubring.kbx
------------------------
pub   rsa3072 2023-03-25 [SC] [expires: 2025-03-24]
      D33D866BD87FB869E1F6EC4D0576A7313564DDDC
uid           [ultimate] k
sub   rsa3072 2023-03-25 [E]
```

例如其中的 `D33D866BD87FB869E1F6EC4D0576A7313564DDDC` 就是 key ID

```
root@v2:/home/ubuntu# gpg -k D33D866BD87FB869E1F6EC4D0576A7313564DDDC
pub   rsa3072 2023-03-25 [SC] [expires: 2025-03-24]
      D33D866BD87FB869E1F6EC4D0576A7313564DDDC
uid           [ultimate] k
sub   rsa3072 2023-03-25 [E]
```

## By fingerprint

```
root@v2:/home/ubuntu# gpg --fingerprint
/root/.gnupg/pubring.kbx
------------------------
pub   rsa3072 2023-03-25 [SC] [expires: 2025-03-24]
      D33D 866B D87F B869 E1F6  EC4D 0576 A731 3564 DDDC
uid           [ultimate] k
sub   rsa3072 2023-03-25 [E]
```

例如其中的 `D33D 866B D87F B869 E1F6  EC4D 0576 A731 3564` 就是 fingerprint

```
root@v2:/home/ubuntu# gpg -k "D33D 866B D87F B869 E1F6  EC4D 0576 A731 3564 DDDC"
pub   rsa3072 2023-03-25 [SC] [expires: 2025-03-24]
      D33D866BD87FB869E1F6EC4D0576A7313564DDDC
uid           [ultimate] k
sub   rsa3072 2023-03-25 [E]
```

## By exact mathc on OpenPGP user ID

```
root@v2:/home/ubuntu# gpg -k
gpg: checking the trustdb
gpg: marginals needed: 3  completes needed: 1  trust model: pgp
gpg: depth: 0  valid:   1  signed:   0  trust: 0-, 0q, 0n, 0m, 0f, 1u
/root/.gnupg/pubring.kbx
------------------------
pub   rsa3072 2023-03-25 [SC]
      CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
uid           [ultimate] alice (this is a comment) <alice@yahoo.com>
sub   rsa3072 2023-03-25 [E
```

例如其中的 `alice (this is a comment) <alice@yahoo.com>` 就是 user ID，必须包含 comment

```
root@v2:/home/ubuntu# gpg -k "alice (this is a comment) <alice@yahoo.com>"
pub   rsa3072 2023-03-25 [SC]
      CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
uid           [ultimate] alice (this is a comment) <alice@yahoo.com>
sub   rsa3072 2023-03-25 [E]
```

## By excact match on an email address

还是上面的例子，其中的 `<alice@yahoo.com>` 就是 email address，必须加 double quote

```
root@v2:/home/ubuntu# gpg -k "<alice@yahoo.com>"
pub   rsa3072 2023-03-25 [SC]
      CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
uid           [ultimate] alice (this is a comment) <alice@yahoo.com>
sub   rsa3072 2023-03-25 [E]
```

## By partial match on an email address

还是上面的例子，下面都是 partial of email address

```
root@v2:/home/ubuntu# gpg -k "@yahoo.com"
pub   rsa3072 2023-03-25 [SC]
      CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
uid           [ultimate] alice (this is a comment) <alice@yahoo.com>
sub   rsa3072 2023-03-25 [E]

root@v2:/home/ubuntu# gpg -k "@alice"
pub   rsa3072 2023-03-25 [SC]
      CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
uid           [ultimate] alice (this is a comment) <alice@yahoo.com>
sub   rsa3072 2023-03-25 [E]
```

## By substring match

匹配任意字符串

```
root@v2:/home/ubuntu# gpg -k
gpg: checking the trustdb
gpg: marginals needed: 3  completes needed: 1  trust model: pgp
gpg: depth: 0  valid:   2  signed:   0  trust: 0-, 0q, 0n, 0m, 0f, 2u
gpg: next trustdb check due at 2025-03-24
/root/.gnupg/pubring.kbx
------------------------
pub   rsa3072 2023-03-25 [SC]
      CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
uid           [ultimate] alice (this is a comment) <alice@yahoo.com>
sub   rsa3072 2023-03-25 [E]

pub   rsa3072 2023-03-25 [SC] [expires: 2025-03-24]
      240EF6FAD697C83518F4EA997CD88D19A58AD3EC
uid           [ultimate] alicea
sub   rsa3072 2023-03-25 [E]
```

如果字符串是多个 gpg key 的子串，默认会匹配多个

```
root@v2:/home/ubuntu# gpg -k alice
pub   rsa3072 2023-03-25 [SC]
      CB66D424A9BE4DC5B9BAB4FD0FC2114BEEB4EBF8
uid           [ultimate] alice (this is a comment) <alice@yahoo.com>
sub   rsa3072 2023-03-25 [E]

pub   rsa3072 2023-03-25 [SC] [expires: 2025-03-24]
      240EF6FAD697C83518F4EA997CD88D19A58AD3EC
uid           [ultimate] alicea
sub   rsa3072 2023-03-25 [E
```

