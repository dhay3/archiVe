# Linux 磁盘管理：e2label

## volume-label

volume label通常也叫做volume name，也就是我们说的卷名

## e2label

syntax：`e2label device [volume-label]`

e2label用于修改`ext2/ext3/ext4`fs label

如果没有给定volume-label，e2label会使用device当前的volume

```
cpl in ~ λ sudo e2label /dev/sdb1 
e2label: Bad magic number in super-block while trying to open /dev/sdb1
/dev/sdb1 contains a iso9660 file system labelled 'Kali Live'
cpl in ~ λ sudo e2label /dev/sdb2 
e2label: Bad magic number in super-block while trying to open /dev/sdb2
/dev/sdb2 contains a vfat file system
```

如果给了volume-label就是修改，==修改前必须umount==

```
cpl in ~ λ sudo e2label /dev/sdb3 kali-persistence
Recovering journal.
cpl in ~ λ sudo e2label /dev/sdb3                 
kali-persistence
```

如果filesystem是ntfs的话，请使用ntfslabel

```
cpl in ~ λ sudo ntfslabel /dev/sdb5 win-store
cpl in ~ λ sudo ntfslabel /dev/sdb5          
win-store
```

