# Linux cp

## Digest

> 可以使用 `rsync` 替代
>
> 需要注意的一点是复制时，以当前的用户作为文件的 owner 而不是原文件的用户

syntax

```
cp [OPTION]... [-T] SOURCE DEST
cp [OPTION]... SOURCE... DIRECTORY
cp [OPTION]... -t DIRECTORY SOURCE...
```

`cp` 用于将指定文件(多个文件)或者目录(多个目录)拷贝到指定文件或者目录

## Optional args

- `-R`

  copy directories recursively

  如果需要复制文件夹，必须要使用该参数

- `-n | --no-clobber`

  如果目标文件存在对文件不覆盖，==默认会覆盖==

  ```
  [vagrant@localhost ~]$ cat a
  true
  [vagrant@localhost ~]$ cat b
  [vagrant@localhost ~]$ cp a b
  [vagrant@localhost ~]$ cat b
  true
  #清空文件 b
  [vagrant@localhost ~]$ true > b
  #指定参数不覆盖目标文件
  vagrant@localhost ~]$ cp -n a b
  [vagrant@localhost ~]$ cat b
  [vagrant@localhost ~]$ 
  ```

- `-i | --interactive`

  如果复制需要将目的文件覆盖，会提示。==默认==

  ```
  [vagrant@localhost ~]$ cp -i a b
  cp: overwrite ‘b’?
  ```

- `-f | --force`

  强制复制即使目的文件存在

- `-u | --update`

  只有在源文件 mtime 比目标文件 mtime 或者是 目标文件缺失的情况下才复制

- `-p | --preserver`

  ==复制文件时保留 mode，owner-ship，timestamps==

  等价于`--preserver=mode,owner-ship,timestamps`

  ```
  ➜  cp 1 3
  ➜  ll
  .rwxrwxr-- root root 0 B Thu Jun 10 11:57:13 2021  1
  .rwxr-xr-- root root 0 B Thu Jun 10 12:00:37 2021  3
  ➜  cp -p 1 4  
  ➜  ll
  .rwxrwxr-- root root 0 B Thu Jun 10 11:57:13 2021  1
  .rwxr-xr-- root root 0 B Thu Jun 10 12:00:37 2021  3
  .rwxrwxr-- root root 0 B Thu Jun 10 11:57:13 2021  4
  ```

- `-a`

  等价于`-dR --preserve`

- `-d | --no-dereference`

  复制时复制 link，而不是复制 link 对应的文件

- `-L | --dereference`

  复制时复制 link 指向的文件而不是link，==默认复制 link 指向的文件==

- `-s | --symbolic-link`

  复制文件或者目录时，创建 link 而不是直接复制文件

## Labs

### 0x01 root 复制非 root 文件

```
[vagrant@localhost ~]$ touch a
[vagrant@localhost ~]$ sudo bash -c bash
[root@localhost vagrant]# ll
total 0
-rw-rw-r--. 1 vagrant vagrant 0 Nov 21 16:56 a
[root@localhost vagrant]# cp a b
total 0
-rw-rw-r--. 1 vagrant vagrant 0 Nov 21 16:56 a
-rw-r--r--. 1 root    root    0 Nov 21 16:58 b
```

这里可以看到的是，文件的 ownership 和 ownergroup 都改变了，变成当前登入的用户。

但是 mode 和 btime 都不一样的

### 0x02 复制含 link 目录

```
[root@localhost vagrant]# tree test
test
├── a
└── a.link -> a

0 directories, 2 files
[root@localhost vagrant]# cp -R test test1 
[root@localhost vagrant]# tree test1
test1
├── a
└── a.link -> a

0 directories, 2 files
```

如果是复制目录默认，会默认复制 link，而不是 link 对应的文件