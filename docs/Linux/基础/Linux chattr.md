# Linux chattr

参考：http://c.biancheng.net/view/874.html

用于修改文件的属性，和windows上的assoc类似。`+`表示添加属性，`-`表示减去属性

pattern：`chattr [option] [mode] files...`

```
chattr -i /etc/resolv.conf
```

**属性**

1. `a`：文件只能做append，不能做其他修改和清空操作；如果是文件只能新建和修改，但是不能删除。只有root或有CAP_LINUX_IMMUTABLE属性的进程才能设置和清除

2. `A`：文件的modified time 不能被修改

   ```
   root in /usr/local/\ λ stat ssh_payload.sh
     File: ssh_payload.sh
     Size: 147       	Blocks: 8          IO Block: 4096   regular file
   Device: 801h/2049d	Inode: 921620      Links: 1
   Access: (0755/-rwxr-xr-x)  Uid: (    0/    root)   Gid: (    0/    root)
   Access: 2021-02-05 13:07:24.633242042 +0800
   Modify: 2021-02-05 13:05:30.881239225 +0800
   Change: 2021-02-05 13:07:18.717241896 +0800
   ```

3. `c`：文件将会被kernel自动压缩
4. `d`：文件不会被dump程序做core dump
5. `E`：文件加密，不能被chattr设置或清除
6. `i`：文件不能被修改（删除，重命名，写入）。只有root和有CAP_SYS_RESOURCE的进程才能设置或清除该属性
7. `u`：当有该标记的文件被删除时，文件内容还是被保存，以保证后期能够恢复
8. `s`：和 u 相反，删除文件或目录时，会被彻底删除（直接从硬盘上删除，然后用 0 填充所占用的区域），不可恢复。

## 查看

通常使用`lsattr`来查看文件的属性

```
root in /usr/local/\ λlsattr 
--------------e----- ./ssh_payload.sh~
--------------e----- ./compress_log.sh
--------------e----- ./CVE-2021-3156-main
--------------e----- ./compress_log_cron.sh
--------------e----- ./testfile
-----a--c-----e----- ./ssh_payload.sh
--------------e----- ./test
--------------e----- ./test1    
```

## 参数

- `-R`

  递归变更文件的属性



