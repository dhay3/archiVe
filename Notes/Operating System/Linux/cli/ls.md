# Linux ls

## Digest

syntax：`ls [options]... [file]...`

list directory contents

默认按照字母顺序( alphabetically )排列显示

## Optional args

- `-a | --all`

  显示隐藏文件以`.`开头的文件。同样会包含`.`（表示当前目录）和`..`( 表示上级目录 )

- `-A | --almost-all`

  显示隐藏文件，但是不包含表示当前目录的`.`和表示上级目录的`..`

- `-l`

  使用 long listing format

  ```
  drwx------ 2 cpl  cpl      40 Jul 27 20:09  lattedock-Azixsp
  ```

  第一列：标识文件的类型，组与用户权限

  第二列：文件的链接(hard link)数

  第三列：文件的所有者

  第四列：文件所在的组（有root组）

  第五列：文件大小（allocated size）==如果是文件夹不会显示真实的大小==

  第六列：文件最后修改的时间(mtime)

  第七列：文件的名称

  **文件类型**

  - d：目录

    -：普通文件

    l：链接

    s：socket

    p：name pipe

    b：block device

    c：character device

  **文件权限**

  - r：表示“可读”，用数字4表示
  - w：表示“可写”，用数字2表示
  - x：表示“可执行”，用数字1表示

- `-b | --escape`

  将特殊字符用 C style 转义

  ```
  # ls
  '1 2'
  # ls -b
  1\ 2
  ```

- `--block-size=SIZE`

  需要和 `-l` 一起使用，以指定的 SIZE 显示 size 字段

- `-B | --ignore-backups`

  显示`~`结尾的文件，通常表示备份

- `-d | --direcotry`

  只显示目录，不显示文件。通常和`-R`一起使用

- `--group-directories-first`

  先展示目录文件，然后再显示其他文件

- `-h | --human-readable`

  和`-l`一起使用，文件的 size 字段以 K,M,G为单位显示

- `-i | --inode`

  显示文件的 inode

- `-n | --numeric-uid-gid`

  list numeric user and group ids

- `-Q | --quote-name`

  enclose entry names in double quotes

- `-R | --recursize`

  遍历目录的同时也会遍历子目录

- `-Z | --context`

  显示文件关联 selinux 部分的信息 

### sort

- `-r | --reverse`

  当 list 目录时，逆向排序

- `-S`

  ==sort by file size, largest first==

- `-t`

  sort by mtime（newest first）

  默认以 mtime 作为参考

- `-u`

  和`-lt`一起时，使用 atime 作为参考排序，和`-l`一起使用，显示 atime （默认显示 mtime）但是以文件名排序

- `-c`

  和`-lt`一起时，使用 ctime 作为参考排序，和`-l`一起使用，显示 ctime （默认显示 mtime）但是以文件名排序

- `--time=WORD`

  需要和`-t`一起使用，默认使用 mtime, WORD 可以是 atime, ctime ,mtime, birth

- `--sort=WORD`

  以特定字段排序，WORD可以是 none, size, time, version, extension, width

## Tricks

调用`printf  '%s\n' *`也能达到 list 文件的效果(代码实际就是这样写的) 。

可以利用这个做一些 hack

## OpenSource

有一些开源的 CLI 可以很好的替代 ls 做一些客制化，比如 lsd

https://github.com/Peltoche/lsd