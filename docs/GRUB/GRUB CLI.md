ref:
[https://www.gnu.org/software/grub/manual/grub/grub.html#Commands](https://www.gnu.org/software/grub/manual/grub/grub.html#Commands)
[https://www.hackerearth.com/practice/notes/akshaypai94/grub-rescue/](https://www.hackerearth.com/practice/notes/akshaypai94/grub-rescue/)
## Drive/partitio Naming
在GRUB中设备名必须在`()`中

1. `(fd0)`，fd表示是一个 floppy disk，0 表示 driver number(从 0 开始)。整个含义表示 GRUB 会使用整个 floppy disk
1. `(hd0,msdos2)`，hd表示是一个 hard disk，0 表示 dirver numer，msdos 表示分区使用的schem(type)，2表示 partition number（从1开始不是0）。整个含义表示GRUB只使用 hd0 的第二个分区类型为msdos
1. `(hd0,msdos5)`，partition number 为 5 表示 the first extended partition of the first hard disk drive（扩展分区从第5个分区开始计算）

在GRUB中为了进入 disk 或者 partitions，需要使用指令指定。例如进入`fd0`，需要使用`set root=(fd0)`或者`parttoll (fd0) hidden-` ，如果非 rescue mode 下可以使用 TAB 补全。
如果需要指定文件可以使用`(hd0,msdos1)/etc`
## Commands

- `help`

查看帮助信息

- `clear`

clear the screen

- `boot`

引导加载的OS

- `reboot`

重启

- `ls`

list devices or files，如果没有待任何参数会显示 all devices known to GRUB
```bash
grub > ls 
(hd0) (hd0,msdosl)
grub > ls (hd0,msdosl)/boot/grub2
...
```

- `cat FILE`

查看文件内容
```bash
grub > cat (hd0,msdos1)/home/vagrant/a
#!/bin/bash
for ((i=0;i<10;i++))
do
        tput sc; tput civis                    
        echo -ne $(date +'%Y-%m-%d %H:%M:%S')  
        sleep 1
        tput rc                               
done
tput el; tput cnorm                           
```

- `cmp FILE1 FILE2`

比较两个文件

- `date`

打印当前的时间

- `initrd FILE [FILE...]`

Load，in order, all initial ramdisks for a linux kernel image, and set the approriate parameters in the linux setup area in memeory

- `insmode MOUDLE`

inset the dynamic GRUB moudle，通常用来加载normal module，以进入normal grub

- `rmmod MOUDLE`

remove a loaded module

- `linux FILE...`

以32 bit 方式从文件中加载linux kernel，可以在后面拼接kernel command。在执行该命令后需要执行`initrd`，否则不会将kernel载入
```bash
grub> 1inux /boot/vmlinuz-3.10.O-11Z7.el.x86_64
```

- `list_env`

显示block file 中的环境变量，不是当前所有的环境变量

- `lsmod`

显示所有载入的modules

- `normal [FILE]`

enter normal mode and display the GRUB menu.
in normal mode , commands,filesystem modules and cryptography moudules are automatically loaded
正常进入 boot menu

- `normal_exit`

如果不是 nested，直接进入 rescue mode (不会补全命令)

- `search `
- `set [envvar=value]`

设置环境变量，如果没有指定键值对会遍历所有的环境变量

- `unset [envvar]`

取消环境变量的值
## Special env

- `root`

the root device name，在device name 后指定 file names 无意义，在设置了该值每次使用command时无需指定device name，和正常使用 ls 命令一样
```bash
grub >ls /home
vagrant
```
通常由`prefix`环境变量来决定
例如 GRUB 被安装在 hard disk 的第一个分区，prefix的值可能是`(hd0,msdos1)/boot/grub`，root的值`(hd0,msdos1)`

- `prefix`

the location of the `/boot/grub`directory as an absolute file name
通常由GRUB自动设置，值为使用`grub-install`时指定的
## rescue mode to normal mode
[https://www.gnu.org/software/grub/manual/grub/grub.html#GRUB%20only%20offers%20a%20rescue%20shell](https://www.gnu.org/software/grub/manual/grub/grub.html#GRUB%20only%20offers%20a%20rescue%20shell)
如果 grub 自动进入了 rescue mode 一般是 root 或 prefix的值设置错误，可以手动来设置
```bash
grub> normal_exit
grub rescue> set root=(hd0,msdosl)
grub rescue> set prefix=(hd0,msdosl)/boot/grub2
grub rescue> insmod normal
grub rescue > set
prefix=(hd0,msdosl)/boot/grub2
root=(hd0,msdosl)
grub rescue > normal
```
## normal mode to boot manu
```bash
#如果 root 或 prefix 值没有或错误需要配置
grub > set root=(hd0,msdosl)
grub > set prefix=(hd0,msdosl)/boot/grub2
# 加载 linux 内核
grub > linux /boot/vmlinuz-xxxx
# 加载 ram disk，如果不加载就boot进入会hang

```
