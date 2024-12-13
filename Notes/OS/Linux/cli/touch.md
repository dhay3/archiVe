# Linux touch

## 0x01 Overview

touch 主要有如下几点功能

1. update acess time(atime) and modification times(mtime) of files
2. create empty files

## 0x02 Syntax

```
touch [OPTION]... FILE...
```

FILE 是 positional args 且是变参
- 当 FILE 不存在时，会新建 empty files
- 当 FILE 存在时，同时更新 atime 和 mtime

Linux 默认不能手动修改 change time(ctime)，只有文件的 metadata 修改后自动更新 change time(ctime)[^1]

```
# Update ctime
debugfs -w -R 'set_inode_field /tmp/foo ctime 201001010101' /dev/sda1

# Drop vm cache so ctime update is reflected
echo 2 > /proc/sys/vm/drop_caches
```

brith time(btime) 不能被强制修改。但是可以修改机器的时间(前提需要把 NTP 停了) 或者自己手动编译内核[^2]

```
$ timedatectl set-ntp false
$ timedatectl set-time "1970-01-02 18:00:00"
$ touch a
$ stat a
  File: ‘a’
  Size: 0               Blocks: 0          IO Block: 4096   regular empty file
Device: 801h/2049d      Inode: 67182222    Links: 1
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:user_tmp_t:s0
Access: 1970-01-02 18:00:03.963982454 +0000
Modify: 1970-01-02 18:00:03.963982454 +0000
Change: 1970-01-02 18:00:03.963982454 +0000
 Birth: -
```

## 0x03 Optional args

- `a`
	只修改 atime，默认会修改到当前时间
	```
	$ stat a
	  File: a
	  Size: 6               Blocks: 8          IO Block: 4096   regular file
	Device: 0,35    Inode: 176         Links: 1
	Access: (0644/-rw-r--r--)  Uid: ( 1000/      cc)   Gid: ( 1000/      cc)
	Access: 2024-06-19 09:38:36.336558886 +0800
	Modify: 2024-06-19 09:38:36.336558886 +0800
	Change: 2024-06-19 09:38:36.336558886 +0800
	 Birth: 2024-06-19 09:33:30.176038255 +0800
	
	$ touch -a a
	
	$ stat a
	  File: a
	  Size: 6               Blocks: 8          IO Block: 4096   regular file
	Device: 0,35    Inode: 176         Links: 1
	Access: (0644/-rw-r--r--)  Uid: ( 1000/      cc)   Gid: ( 1000/      cc)
	Access: 2024-06-19 09:38:47.646580194 +0800
	Modify: 2024-06-19 09:38:36.336558886 +0800
	Change: 2024-06-19 09:38:47.646580194 +0800
	 Birth: 2024-06-19 09:33:30.176038255 +0800
	```

- `m`
	只修改 mtime，ctime 同时也会发生改变
	```
	$ stat a
	  File: a
	  Size: 6               Blocks: 8          IO Block: 4096   regular file
	Device: 0,35    Inode: 176         Links: 1
	Access: (0644/-rw-r--r--)  Uid: ( 1000/      cc)   Gid: ( 1000/      cc)
	Access: 2024-06-19 09:38:47.646580194 +0800
	Modify: 2024-06-19 09:38:36.336558886 +0800
	Change: 2024-06-19 09:38:47.646580194 +0800
	 Birth: 2024-06-19 09:33:30.176038255 +0800
	
	$ touch -m a
	
	$ stat a
	  File: a
	  Size: 6               Blocks: 8          IO Block: 4096   regular file
	Device: 0,35    Inode: 176         Links: 1
	Access: (0644/-rw-r--r--)  Uid: ( 1000/      cc)   Gid: ( 1000/      cc)
	Access: 2024-06-19 09:38:47.646580194 +0800
	Modify: 2024-06-19 09:40:45.286808362 +0800
	Change: 2024-06-19 09:40:45.286808362 +0800
	 Birth: 2024-06-19 09:33:30.176038255 +0800
	```

- `-d | --date=<STRING>`
	修改atime和mtime到指定时间，STRING 格式具体查看 man page
  ```
  $ stat a
    File: a
    Size: 6               Blocks: 8          IO Block: 4096   regular file
  Device: 0,35    Inode: 176         Links: 1
  Access: (0644/-rw-r--r--)  Uid: ( 1000/      cc)   Gid: ( 1000/      cc)
  Access: 2024-06-19 09:38:47.646580194 +0800
  Modify: 2024-06-19 09:40:45.286808362 +0800
  Change: 2024-06-19 09:40:45.286808362 +0800
   Birth: 2024-06-19 09:33:30.176038255 +0800
  
  $ touch -d "2020-01-15 10:30:45" a
  
  $ stat a
    File: a
    Size: 6               Blocks: 8          IO Block: 4096   regular file
  Device: 0,35    Inode: 176         Links: 1
  Access: (0644/-rw-r--r--)  Uid: ( 1000/      cc)   Gid: ( 1000/      cc)
  Access: 2020-01-15 10:30:45.000000000 +0800
  Modify: 2020-01-15 10:30:45.000000000 +0800
  Change: 2024-06-19 09:47:07.940944021 +0800
   Birth: 2024-06-19 09:33:30.176038255 +0800
  ```

- `-r | --reference=<FILE>`
	使用指定文件的时间戳修改 atime mtime
  ```
  root in /usr/local λ stat bin
    File: bin
    Size: 4096            Blocks: 8          IO Block: 4096   directory
  Device: fc01h/64513d    Inode: 396189      Links: 2
  Access: (0755/drwxr-xr-x)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2021-04-13 09:57:56.820566390 +0800
  Modify: 2020-03-11 16:52:58.561788635 +0800
  Change: 2020-03-11 16:52:58.561788635 +0800
   Birth: -
   
   root in /usr/local λ stat c
    File: c
    Size: 254             Blocks: 8          IO Block: 4096   regular file
  Device: fc01h/64513d    Inode: 407745      Links: 1
  Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2021-04-13 10:33:27.280482886 +0800
  Modify: 2021-04-13 10:33:26.824487146 +0800
  Change: 2021-04-13 10:33:26.824487146 +0800
   Birth: -
  
  $ touch -r c bin
  
  root in /usr/local λ stat bin
    File: bin
    Size: 4096            Blocks: 8          IO Block: 4096   directory
  Device: fc01h/64513d    Inode: 396189      Links: 2
  Access: (0755/drwxr-xr-x)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2021-04-13 10:33:27.280482886 +0800
  Modify: 2021-04-13 10:33:26.824487146 +0800
  Change: 2021-04-13 10:36:59.422500498 +0800
   Birth: -
  ```

**references**

[^1]:[timestamps - How can I change 'change' date of file? - Unix & Linux Stack Exchange](https://unix.stackexchange.com/questions/36021/how-can-i-change-change-date-of-file)
[^2]:[linux - Change file "Birth date" for ext4 files? - Unix & Linux Stack Exchange](https://unix.stackexchange.com/questions/556040/change-file-birth-date-for-ext4-files)
  

  
 
  
  



