# qemu-img/虚拟磁盘

参考：

https://wiki.archlinux.org/title/QEMU#Tips_and_tricks

> 禁止使用 qemu-img 修改正常运行的 VM 的 disk image

qemu-img用于生成hard disk imag(模拟的虚拟磁盘)，file format不同生成的hard disk image 不同

## ENBF

```
FILENAME := is a disk image filename

FMT := is the disk image format. It is guessed automatically in most cases. See below for a description of the supported disk formats

SIZE := is the disk image size in  bytes.  Optional  suffixes  k  or  K (kilobyte,  1024)  M (megabyte, 1024k) and G (gigabyte, 1024M) and T (terabyte, 1024G) are supported.  b is ignored.

OUTPUT_FILENAME := is the destination disk image filename.

OUTPUT_FMT := is the destination format.

OPTIONS := is a comma separated list of format specific  options  in  a name=value format. Use -o ? for an overview of the options supported by the used format or see the format descriptions below for details.

SNAPSHOT_PARAM := is param used for internal snapshot, format is 'snapshot.id=[ID],snapshot.name=[NAME]' or '[ID_OR_NAME]'
```

## commands

### create

syntax：`qemu-img create [-f fmt] [-o options] filename [size]`

以指定的size生成名为filename的disk image。（可以使用`du`或`qemu-img info`来查看）

```
cpl in ~/software/kvm λ qemu-img create -f qcow2 win.img 100M
Formatting 'win.img', fmt=qcow2 size=104857600 cluster_size=65536 lazy_refcounts=off refcount_bits=16
```

也可以对现有的disk image进行备份(只有qcow2才支持)

```bash
#qemu-img create -f qcow2 -o backing_file=src,backing_fmt=src_fmt bak_file

cpl in /sharing/vm λ qemu-img create -f qcow2 -o backing_file=raw,backing_fmt=raw raw.bak
Formatting 'raw.bak', fmt=qcow2 cluster_size=65536 extended_l2=off compression_type=zlib size=10485760 backing_file=raw backing_fmt=raw lazy_refcounts=off refcount_bits=16
```

具体支持的options，查看qemu-img note章节

### info

syntax：`qemu-img info <filename>`

输出指定disk image的信息

```bash
cpl in ~/software/kvm λ qemu-img info win.img 
image: win.img
file format: qcow2
virtual size: 100 MiB (104857600 bytes)
disk size: 196 KiB
cluster_size: 65536
Format specific information:
    compat: 1.1
    lazy refcounts: false
    refcount bits: 16
    corrupt: false
```

### resize

syntax：`qemu-img resize [-f fmt] [--shrink] <filename> [+|- ]size`

修改disk image的大小，如果需要缩小必须使用`--shrink`(需要先修改虚拟机中的分区和file system，否则数据会丢失)

```bash
cpl in ~/software/kvm λ qemu-img resize win.img +10M
Image resized.
```

修改完后需要进入虚拟机修改file system

### rebase

changing the backing file of an image

### convert

syntax：`qemu-img [-f src_fmt] [-O output_fmt] [-B bak_file] [-l snapshot] [-c] src_filename output_filename`

以指定output_fmt修改disk image或snapshot到output_filename。只有src_fmt为qcow或qcow2的才支持`-c`压缩

```
cpl in /sharing/vm λ qemu-img convert -f raw -O qcow raw raw.cow
cpl in /sharing/vm λ qemu-img info raw.cow 
image: raw.cow
file format: qcow
virtual size: 10 MiB (10485760 bytes)
disk size: 4 KiB
cluster_size: 4096
```

### bench

syntax：`qemu-img bench [-d depth] [-c count] [-s buffer_size] [-w]`

对disk image做IO测试，分别指定深度，IO请求次数，每次IO请求的大小。默认只做read IO，使用`-w`表示做write IO

```bash
qemu-img bench -c 10 -d 0 -s 1M win.img -w
```

## format

qemu支持的file format有如下几种类型

- raw

  默认使用的类型，可以将这种格式的file导入到其他的emulators。直接分配，分配多少占用宿主机就是多少。也可以通过`dd`来生成raw类型的hard disk image

  ```bash
  cpl in /sharing/vm λ dd if=/dev/zero of=dd bs=10M count=1
  1+0 records in
  1+0 records out
  10485760 bytes (10 MB, 10 MiB) copied, 0.00958298 s, 1.1 GB/s
  cpl in /sharing/vm λ ll
  .rw-r--r-- cpl cpl  10 MB Wed Jun  2 18:11:49 2021  dd
  .rw-r--r-- cpl cpl 192 KB Wed Jun  2 18:11:00 2021  qcow2
  .rw-r--r-- cpl cpl  10 MB Wed Jun  2 18:10:55 2021  raw
  ```

- qcow2

  生成的镜像更小，常常用在不支持linux fs的系统(例如windows)。只有虚拟机往扇区中写时才会实际分配。如果使用这种方式替代raw类型的disk image 可能会造成性能上的影响

  ```bash
  cpl in /sharing/vm λ qemu-img create -f raw raw 10M    
  Formatting 'raw', fmt=raw size=10485760
  cpl in /sharing/vm λ qemu-img create -f qcow2 qcow2 10M
  Formatting 'qcow2', fmt=qcow2 cluster_size=65536 extended_l2=off compression_type=zlib size=10485760 lazy_refcounts=off refcount_bits=16
  cpl in /sharing/vm λ ll
  .rw-r--r-- cpl cpl 192 KB Wed Jun  2 18:11:00 2021  qcow2
  .rw-r--r-- cpl cpl  10 MB Wed Jun  2 18:10:55 2021  raw
  ```

- other

  包括其他一些兼容的格式，例如VMDK，VDI，VHD，qcow1和QED等

