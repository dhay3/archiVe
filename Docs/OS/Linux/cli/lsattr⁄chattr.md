# Linux lsattr/chattr

ref:

https://en.wikipedia.org/wiki/Ext2

https://en.wikipedia.org/wiki/Extended_file_attributes#Linux

https://www.nongnu.org/ext2-doc/ext2.html

http://c.biancheng.net/view/874.html

https://en.wikipedia.org/wiki/Copy-on-write

https://en.wikipedia.org/wiki/Extent_(file_systems)

> 只能被用在 ext2/ext4 file system 中，这也就意味不能在 tmpfs 中使用，即不能在`/tmp`中使用

## Terms

- second extended file system

  通常也被叫做 ==ext2== ，最早是被用来替代 ext 

  大多数 linux distro 都使用 ext2 作为 file system（不过现在大多数都支持 ext4，比 ext2 支持更大的 volume size）

- extended attributes

  通常也被叫做 ==xattr==

- copy on write

  is a resource-managment technique used in computer programming to efficiently implement a “duplicate” or “copy” operation on modifiable resources

  通常也被叫做 ==COW==, 对比 KVM 中的 qcow format 就比较好理解了

- extent

  is a contiguous area of storage reserved for a file in a file system, represented as a range of block numbers, or tracks on count key data devices. ==A file can consist of zero or more extents==

  可以简单的将 extent 理解成一组 blocks

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

| Attribute                     | lsattr flag | chattr option                                                | Semantics and rationale                                      |
| ----------------------------- | ----------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| No `atime` updates            | `A`         | `+A` to set `-A` to clear                                    | When a file with the `A` attribute set is accessed, its [atime record](https://en.wikipedia.org/wiki/Stat_(Unix)) is not modified. This avoids a certain amount of disk I/O operations. |
| Append only                   | `a`         | `+a` to set `-a` to clear[[note 1\]](https://en.wikipedia.org/wiki/Chattr#cite_note-cap_immutable-4) | A file with the `a` attribute set can only be open in append mode for writing. |
| Compressed                    | `c`         | `+c` to set `-c` to clear[[note 2\]](https://en.wikipedia.org/wiki/Chattr#cite_note-not_honored_ext23-5) | A file with the `c` attribute set is automatically compressed on the disk by the kernel. A read from this file returns uncompressed data. A write to this file compresses data before storing them on the disk. |
| No Copy-on-Write (CoW)        | `C`         | `+C` to set `-C` to clear[[note 3\]](https://en.wikipedia.org/wiki/Chattr#cite_note-cow_only-6) | A file with the `C` attribute will not be subject to Copy-on-Write updates. Updates to these files may not be subject to atomic snapshots, and  may lack some reliability information on some filesystems and kernels.不允许 copy on write (COW, the copy operation is deferred until the first write)对比 KVM 中的 qcow2 就知道什么含义了 |
| Synchronous directory updates | `D`         | `+D` to set `-D` to clear                                    | When a directory with the `D` attribute set is modified, the changes are written synchronously on the disk This is equivalent to the `dirsync` [mount option](https://en.wikipedia.org/wiki/Mount_(Unix)), applied to a subset of the files. |
| No dump                       | `d`         | `+d` to set `-d` to clear                                    | A file with the `d` attribute set is not candidate for backup when the [dump program](https://en.wikipedia.org/wiki/Dump_(Unix)) is run. |
| Compression error             | `E`         | *(unavailable)*                                              | The `E` attribute is used by the experimental compression patches to indicate that a compressed file has a compression error. |
| Extent format                 | `e`         | *(unavailable)*                                              | The `e` attribute indicates that the file is using [extents](https://en.wikipedia.org/wiki/Extent_(file_systems)) for mapping the blocks on disk. |
| Huge file                     | `h`         | *(unavailable)*                                              | The `h` attribute indicates the file is storing its blocks in units of the filesystem blocksize instead of in units of sectors. It means that the file is, or at one time was, larger than 2TB. |
| Indexed directory             | `I`         | *(unavailable)*                                              | The `I` attribute is used by the [htree program](https://en.wikipedia.org/wiki/HTree#Use) code to indicate that a directory is being indexed using hashed trees. |
| ==Immutable==                 | `i`         | `+i` to set `-i` to clear[[note 1\]](https://en.wikipedia.org/wiki/Chattr#cite_note-cap_immutable-4) | A file with the `i` attribute cannot be modified. It cannot be deleted or renamed, no link can be created to this file and no data can be written to the file. When set, prevents, *even the superuser*, from erasing or changing the contents of the file. |
| Data journaling               | `j`         | `+j` to set `-j` to clear[[note 4\]](https://en.wikipedia.org/wiki/Chattr#cite_note-cap_sysres-7) | A file with the `j` attribute has all of its data written to the [ext3](https://en.wikipedia.org/wiki/Ext3) journal before being written to the file itself, if the filesystem is mounted with the `"data=ordered"` or `"data=writeback"` options. When the filesystem is mounted with the `"data=journal"` option all file data is already [journaled](https://en.wikipedia.org/wiki/Journaling_file_system), so this attribute has no effect. |
| Secure deletion               | `s`         | `+s` to set `-s` to clear[[note 2\]](https://en.wikipedia.org/wiki/Chattr#cite_note-not_honored_ext23-5)[[note 5\]](https://en.wikipedia.org/wiki/Chattr#cite_note-not_honored_ext4-8) | When a file with the `s` attribute set is deleted, [its blocks are zeroed](https://en.wikipedia.org/wiki/Data_remanence#Overwriting) and written back to the disk. |
| Synchronous updates           | `S`         | `+S` to set `-S` to clear                                    | When a file with the `S` attribute set is modified,  the changes are written synchronously on the disk; this is equivalent to the 'sync' mount option applied to a subset of the files. This is equivalent to the `sync` [mount option](https://en.wikipedia.org/wiki/Mount_(Unix)), applied to a subset of the files. |
| Top of directory hierarchy    | `T`         | `+T` to set `-T` to clear                                    | A directory with the `T` attribute will be deemed to be the top of directory hierarchies for the purposes of the [Orlov block allocator](https://en.wikipedia.org/wiki/Orlov_block_allocator). This is a hint to the block allocator used by [ext3](https://en.wikipedia.org/wiki/Ext3) and [ext4](https://en.wikipedia.org/wiki/Ext4) that the subdirectories under this directory are not related, and thus should be spread apart for allocation purposes. For example: it is a very good idea to set the `T` attribute on the `/home` directory, so that `/home/john` and `/home/mary` are placed into separate block groups. For directories where this attribute is not set, the Orlov block  allocator will try to group subdirectories closer together where  possible. |
| No tail-merging               | `t`         | `+t` to set `-t` to clear                                    | For those filesystems that support [tail-merging](https://en.wikipedia.org/wiki/Block_suballocation), a file with the `t` attribute will not have a partial block fragment at the end of the file merged with other files. This is necessary for applications such as [LILO](https://en.wikipedia.org/wiki/LILO_(boot_loader)), which reads the filesystem directly and doesn't understand tail-merged files. |
| Undeletable                   | `u`         | `+u` to set `-u` to clear[[note 2\]](https://en.wikipedia.org/wiki/Chattr#cite_note-not_honored_ext23-5) | When a file with the `u` attribute set is deleted, its contents are saved. This allows the user to ask for its [undeletion](https://en.wikipedia.org/wiki/Undeletion). |
| Compression raw access        | `X`         | *(unavailable)*                                              | The `X` attribute is used by the experimental  compression patches to indicate that a raw contents of a compressed file can be accessed directly. |
| Compressed dirty file         | `Z`         | *(unavailable)*                                              | The `Z` attribute is used by the experimental compression patches to indicate a compressed file is "dirty". |
| Version / generation number   | `-v`        | `-v *version*`                                               | File's version/generation number.                            |

### Optional args

- `-R`

  recursively change attributes of directories and their contents