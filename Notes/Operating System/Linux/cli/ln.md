# ln

## Digest

syntax

```
ln [OPTION]... [-T] TARGET LINK_NAME
ln [OPTION]... TARGET
ln [OPTION]... TARGET... DIRECTORY
ln [OPTION]... -t DIRECTORY TARGET...
```

`ln` 用于创建 link file (快捷文件)

1. soft link
2. hard link (默认)

## Optional args

- `-f | --force`

  remove existing destination files

  如果目的目录中已经有了对应的 link file，默认不能创建。可以通过改参数删除 link 后创建

-  `-i | --interactive`

  prompt whether to remove destination files

  一般和 `-f` 一起使用

- `-s | --symbolic`

  创建 soft link

## soft link vs hard link

1. soft link可以对文件和目录创建，但是hard link只能对文件创建。
2. hard link与源文件的inode相同，但是soft link与源文件的inode不同，由于inode不同，记录的metadata(权限，大小，数据类型，atime/ctime/mtime)不同
3. 由于hard link与源文件的inode相同，所以当源文件被删除时还是可以通过hard link的inode访问源文件，但是soft link不能因为inode不同
4. 如果源文件被删除，向soft link中写入数据会生成一个源文件但是数据不同，向hard link中写入数据不会产生任何影响
