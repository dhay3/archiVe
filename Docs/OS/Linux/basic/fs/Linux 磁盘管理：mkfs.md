# Linux 磁盘管理：mkfs

build a linux filesystem，也就是我们说的格式化，通常针对分区。可以被`mkfs.<type>`命令取代

- `mkfs.fat [options] device `

  用于生成FAT fs，可以通过`-F fat-size`来指定生成fat32或fat16，如果没有指定由mkfs.fat在16到32中决定

  ```
  mkfs.fat -F 32 /dev/sda1
  ```

- `mkfs.ext4 | mkfs.ext2`
- `mkfs.ntfs`

