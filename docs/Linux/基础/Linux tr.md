# Linux tr

用于删除字符，这里删除`\000`（null）和`\n`换行符

```
root in ~ λ cat /proc/version 
  File: /proc/version
  Linux version 5.7.0-kali1-amd64 (devel@kali.org) (gcc version 9.3.0 (Debian 9.3.0-14), GNU ld (GNU Binutils for Debian) 2.34) #1 SMP Debian 5.7.6-1kali2 (2020-07-01)
root in ~ λ cat /proc/version | tr '\000' '\n'
Linux version 5.7.0-kali1-amd64 (devel@kali.org) (gcc version 9.3.0 (Debian 9.3.0-14), GNU ld (GNU Binutils for Debian) 2.34) #1 SMP Debian 5.7.6-1kali2 (2020-07-01)
```

