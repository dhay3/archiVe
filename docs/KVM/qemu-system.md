# qemu-system

参考：

https://wiki.archlinux.org/title/QEMU#Tips_and_tricks

> `kvm` binary is replaced by `qeum-system`

`qemu-system`根据具体宿主机 的系统架构来命名，例如`qemu-system-x86_64`。如果manual page查看不方便，可以重定向后查看。

syntax：`qemu-system-x86_64 [options] [disk_iamge]`

- `-enable-kvm`

  kvm虚拟化支持

- `-smp [cpus=]n[,cores=cores][,threads=threads]`

  cpus设置虚拟机的CPUs，默认1。

  cores设置CPU的核数

  threads设置单核线程

- `-boot [order=drives][,menu=on|off]`

  虚拟机引导的顺序drives包括floppy(a)，hard disk(c)，CD-ROM(d)，network(n)，会自动识别，是否开启boot menu

- `-m [size=]megs`

  虚拟机的RAM

- `-k language`

  keyboard layout(例如 us，cn)

- `-name string`

  虚拟机的名字

- `-cdrom file`

  `-fda/fdb file`

  `-hda/-hdb file`

  使用file作为引导盘

- `-drive [file=file][,format=f]`

  虚拟机使用的磁盘和格式

- `-display vnc=<IP:Port>[,<optargs>]`

  当宿主机没有图形化功能时，可以设置vnc。其他人可以通过`vnc://IP:Port`连接虚拟机，例如：

  ```
  qemu-system-x86_64 -hda disk01.img -cdrom os.iso -m 512 -boot d -vnc IP:1
  ```

  

## 安装虚拟机

安装时可以将hard disk img 写在flash drive(格式化成ext4)上保证宿主机和虚拟机完全隔离。

```
cpl in /sharing/vm λ qemu-system-x86_64 -enable-kvm -smp cores=2,threads=2 -cdrom Win10_21H1_Chinese\(Simplified\)_x64.iso -boot order=a,menu=on -drive file=win_disk,format=qcow2 -m 4G
```

安装完成后就可以取消`-cdrom`参数，正常启动虚拟机

## 网络
