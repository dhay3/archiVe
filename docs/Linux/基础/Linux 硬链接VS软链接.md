# Linux 硬链接VS软链接

在Unix/Linux系统中允许，多个文件名指向同一个inode号码。

## 硬链接

==这就意味着，可以用不同的文件名访问同样的内容，对文件内容修改，会影响到所有文件，但是删除一个文件名，不影响另一个文件名访问。这种情况就被称为hard link==

使用`ln`命令创建硬链接

```
root in /opt/test λ ln bak ibak                                           /0.0s
root in /opt/test λ ll -i
10647 .rw-r--r-- root root 0 B Wed Dec 30 06:56:26 2020  bak
10647 .rw-r--r-- root root 0 B Wed Dec 30 06:56:26 2020  ibak  
root in /opt/test λ stat bak 
  File: bak
  Size: 0         	Blocks: 0          IO Block: 4096   regular empty file
Device: 801h/2049d	Inode: 10647       Links: 2
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Access: 2020-12-30 06:56:26.568767618 -0500
Modify: 2020-12-30 06:56:26.568767618 -0500
Change: 2020-12-30 06:56:40.816767514 -0500
 Birth: -              
```

这里会发现指向的链接有两个

## 软链接

文件A和文件B的inode号码虽然不一样，但是文件A的内容时文件B的路径。读取文件A时，系统会自动访问文件B。这就意味着文件A依赖于文件B，如果删除文件B，打开文件A就会报错，修改文件A会影响文件B。

使用`ln -s`创建软链接

```
root in /opt/test λ ll -i
10647 .rw-r--r-- root root 6 B Wed Dec 30 06:59:53 2020  bak
10647 .rw-r--r-- root root 6 B Wed Dec 30 06:59:53 2020  ibak
10648 lrwxrwxrwx root root 3 B Wed Dec 30 07:04:42 2020  sbak ⇒ bak     /0.0s
root in /opt/test λ 
root in /opt/test λ stat bak
  File: bak
  Size: 11        	Blocks: 8          IO Block: 4096   regular file
Device: 801h/2049d	Inode: 10647       Links: 2
Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
Access: 2020-12-30 07:05:15.124763757 -0500
Modify: 2020-12-30 07:05:12.444763776 -0500
Change: 2020-12-30 07:05:12.444763776 -0500
```

这里只会发现原先链接的硬链接

## inode特殊作用

1. 有时文件名包含特殊字符，无法正常删除，直接删除inode节点，就能起到删除文件的作用。
2. 移动文件或重命名文件，只是改变文件名，不影响inode
3. 打开一个文件以后，系统就以inode号码来识别这个文件，不再考虑文件名。因此，通常来说，系统无法从inode号码得知文件名。

> 第3点使得软件更新变得简单，可以在不关闭软件的情况下进行更新，不需要重启。因为系统通过inode号码，识别运行中的文件，不通过文件名。更新的时候，新版文件以同样的文件名，生成一个新的inode，不会影响到运行中的文件。等到下一次运行这个软件的时候，文件名就自动指向新版文件，旧版文件的inode则被回收。