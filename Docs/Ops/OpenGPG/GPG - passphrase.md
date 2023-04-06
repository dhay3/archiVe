ref
[https://unix.stackexchange.com/questions/60213/gpg-asks-for-password-even-with-passphrase](https://unix.stackexchange.com/questions/60213/gpg-asks-for-password-even-with-passphrase)
[https://unix.stackexchange.com/questions/60213/gpg-asks-for-password-even-with-passphrase/415064#415064](https://unix.stackexchange.com/questions/60213/gpg-asks-for-password-even-with-passphrase/415064#415064)
## Digest
passphrase 简单的理解就是 GPG secrect key 的密码，为了防止 secrect key 泄露而导致出现不必要安全问题而设计的机制

```
     --passphrase string
            Use string as the passphrase. This can only be used if only one passphrase is supplied. Obviously, this is of  very
            questionable security on a multi-user system. Don't use this option if you can avoid it.

            Note  that  since Version 2.0 this passphrase is only used if the option --batch has also been given. Since Version
            2.1 the --pinentry-mode also needs to be set to loopback.
```
