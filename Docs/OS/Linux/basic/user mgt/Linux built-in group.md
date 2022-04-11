# Linux built-in group

## nobody

nobody是Linux中权限最小的组，例如将一个文件不分配给任何用户时可以使用，通常被指向`/sbin/nologin`

```
cpl in / λ sudo cat /etc/passwd | grep nobody
nobody:x:65534:65534:Nobody:/:/usr/bin/nologin
```

