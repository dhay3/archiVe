# Linux lsattr/chattr

ref:

https://en.wikipedia.org/wiki/Ext2

https://en.wikipedia.org/wiki/Extended_file_attributes#Linux

https://www.nongnu.org/ext2-doc/ext2.html

http://c.biancheng.net/view/874.html

https://en.wikipedia.org/wiki/Copy-on-write

## Terms

- second extended file system

  通常也被叫做 ==ext2== ，最早是被用来替代 ext 

  大多数 linux distro 都使用 ext2 作为 file system（不过现在大多数都支持 ext4，比 ext2 支持更大的 volume size）

- extended attributes

  通常也被叫做 ==xattr==

## lsattr

### Digest

syntax: `lsattr [options] [files...]`

lists the file attributes on a second extended file system 

### Optional args

- `-R`

  recursively list attributes of directories and their contents

- `-l`

  以 long names 格式显示

## chattr

### Digest

syntax：`chattr [-RVf] [-v version] [-p project] [mode] files...`

change the file attributes on second extended file system

和 windows 上的 `assoc` 类似，`+` 表示添加，`-` 表示减去，`=` 表示等于（赋值）

### Attributes

attributes 有如下几种

1. a

   append only

   文件只能做 append，不能做其他修改和清空操作；如果是文件只能新建和修改，但是不能删除。只有root或有CAP_LINUX_IMMUTABLE属性的进程才能设置和清除

2. A

   no atime updates

   atime 不能被修改，几乎等价于不能读写

3. c

   compressed

   如果读文件会返回非压缩的内容，如果写文件会会在存储到 disk 前压缩

   不能和 C 一起使用

4. C

   no copy on write

   不允许 copy on write (COW, the copy operation is deferred until the first write)

   对比 KVM 中的 qcow2 就知道什么含义了

5. d

   no dump

   文件不会被 dump 程序做 core dump

6. e

   extent format

   it may not removed using chattr

7. E

   encrypted by the file system

   it may not removed or set using chattr

7. F

   case-insenstive directory lookups

8. i

   immutable

9. j

   data journaling

10. m

    don’t compress

11. P

    project hierarchy

12. s

    secure deletion

13. S

    synchronous updates

14. t

    no tail-merging

15. T

    top of directory hierarchy

16. u

    undeleteable

17. x

    direct access for files

### Optional args

- `-R`

  recursively change attributes of directories and their contents