# qemu-img/虚拟磁盘

参考：

https://wiki.archlinux.org/title/QEMU#Tips_and_tricks

https://qemu.readthedocs.io/en/latest/tools/qemu-img.html

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

## Optional Args

- `-h`

  with or without a command, shows help and lists the supported formats

- `-p`

  display progress bar

- `-q`

  quiet mode - do not print any output

- `-S SIZE`

  指定 qemu-img create 生成的 disk image 大小。如果没有suffix，默认使用byte

- `--object OBJECTDEF`

  

## Positional Args

只介绍几个commands，具体可以查看 man page

### create

syntax：` create [--object OBJECTDEF] [-q] [-f FMT] [-b BACKING_FILE [-F BACKING_FMT]] [-u] [-o OPTIONS] FILENAME [SIZE]` 

cearte the new disk image FILENAME of size SIZE and format FMT

```
cpl in ~/software/kvm λ qemu-img create -f qcow2 win.img 100M
Formatting 'win.img', fmt=qcow2 size=104857600 cluster_size=65536 lazy_refcounts=off refcount_bits=16
```

- `-f FMT`

  生成的 disk image 格式

- `-b BACKING_FILE`

  只会记录与 BACKING_FILE 之间的 differ

也可以对现有的disk image进行备份(只有qcow2才支持)

```bash
#qemu-img create -f qcow2 -o backing_file=src,backing_fmt=src_fmt bak_file

cpl in /sharing/vm λ qemu-img create -f qcow2 -o backing_file=raw,backing_fmt=raw raw.bak
Formatting 'raw.bak', fmt=qcow2 cluster_size=65536 extended_l2=off compression_type=zlib size=10485760 backing_file=raw backing_fmt=raw lazy_refcounts=off refcount_bits=16
```

具体支持的options，查看qemu-img note章节

### amend

syntax：` amend [--object OBJECTDEF] [--image-opts] [-p] [-q] [-f FMT] [-t CACHE] [--force] -o OPTIONS FILENAME`

Amends the image format specific OPTIONS for the image file FILENAME. Not all file formats support this operation

### bench

syntax：`bench [-c  COUNT] [-d DEPTH] [-f FMT] [--flush-interval=FLUSH_INTERVAL] [-i AIO]  [-n] [--no-drain] [-o OFFSET] [--pattern=PATTERN] [-q] [-s BUFFER_SIZE]  [-S STEP_SIZE] [-t CACHE] [-w] [-U] FILENAME`

run a simple sequential I/O benchmark on the specified image. If `-w` is specified, a write test is performed, otherwise a read test is performed

```
➜  ISO qemu-img bench -c 10 -d 10 -s 1M win.img -w
Sending 10 write requests, 1048576 bytes each, 10 in parallel (starting at offset 0, step size 1048576)
Run completed in 0.005 seconds.
```

- `-c COUNT`

  会执行 COUNT 次 IO 请求，

- `-s BUFFER_SIZE`

  每次IO请求 BUFFER_SIZE byte，

- `-d DEPTH`

  IO请求时，DEPTH 个并发线程

- `--flush-interval=FLUSH_INTERVAL`

  FLUSH_INTERVAL 会在新请求前清空磁盘

### check

syntax：`check [--object OBJECTDEF] [--image-opts] [-q] [-f FMT] [--output=OFMT] [-r [leaks | all]] [-T SRC_CACHE] [-U] FILENAME`

Perform a consistency check on the edisk image FILENAME. only the formats qcow2, qed, parallels, vhdx, vmdk and vdi support consistency checks

```
➜  ISO qemu-img check --output json win.img
{
    "image-end-offset": 262144,
    "total-clusters": 1600,
    "check-errors": 0,
    "filename": "win.img",
    "format": "qcow2"
}
```

check-errors 通常包括如下几种状态码

1. 0

   check completed, the image is now consistent

2. 1

   check not completed because of internal errors

3. 2

   check completed, image is corrupted

4. 64

   checks are not supported by the image format

- `--output`

  指定输出的格式，可以是 human 或者 json

- `-r [leaks | all]`

  qemu-img tries to repair any inconsistencies found during the check，`-r leaks` repairs only cluster leaks, whereas `-r all` fixes all kinds of errors, with a higher risk of choosing the wrong fix or hiding corruption that has already occurred



### snapshot

- `-a`

  回滚到某一快照

- `-c`

  生成快照

- `-d`

  删除快照

- `-l`

  显示image所有的快照

### compare

syntax：`compare [--object OBJECTDEF] [--image-opts] [-f FMT] [-F FMT] [-T SRC_CACHE] [-p] [-q] [-s] [-U] FILENAME1 FILENAME2`

check if two images have the same content

```
➜  ISO qemu-img compare win.img win2.img 
Images are identical.
```

包含如下几种退出值

1. 0

   images are identical

2. 1

   iamges differ

3. 2

   error on opening an image

4. 3

   error on checking a sector allocation

5. 4

   error on reading data

- `-f`

  first image format

- `-F`

  second image format

### dd

syntax：`dd [--image-opts] [-U] [-f FMT] [-O OUTPUT_FMT] [bs=BLOCK_SIZE] [count=BLOCKS] [skip=BLOCKS] if=INPUT of=OUTPUT`

dd copies from INPUT file to OUTPUT file coverting it from FMT format to OUTPUT_FMT

和 linux dd 类似，可以替代 convert 的部分功能

- `bs=BLOCK_SIZE`

  define the block size

- `count=BLOCKS`

  sets the number of input blocks to copy

- `if=INPUT`

  sets the input file

- `of=OUTPUT`

  sets the output file

- `skip=BLOCKS`

  sets the number of input blocks to skip

### info

syntax：`info [--object OBJECTDEF] [--image-opts] [-f FMT] [--output=OFMT] [--backing-chain] [-U] FILENAME`

give information about the disk image FILENAME.

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

syntax：` convert [--object OBJECTDEF] [--image-opts] [--target-image-opts] [--target-is-zero] [--bitmaps [--skip-broken-bitmaps]] [-U] [-C] [-c] [-p] [-q] [-n] [-f FMT] [-t CACHE] [-T SRC_CACHE] [-O OUTPUT_FMT] [-B BACKING_FILE [-F BACKING_FMT]] [-o OPTIONS] [-l SNAPSHOT_PARAM] [-S SPARSE_SIZE] [-r RATE_LIMIT] [-m NUM_COROUTINES] [-W] FILENAME [FILENAME2 [...]] OUTPUT_FILENAME`

Convert the disk image FILENAME or a snapshot SNAPSHOT_PARAM to disk image OUTPUT_FILENAME using format OUTPUT_FMT

```
➜  ISO qemu-img convert -f qcow2 -O raw t1.img t1.raw
➜  ISO qemu-img info t1.raw
image: t1.raw
file format: raw
virtual size: 100 MiB (104857600 bytes)
disk size: 4 KiB
```

- `-f`

  输入 image 的格式

- `-O`

  希望转成的目标格式

- `-c`

  compressed，将 OUTPUT_FLIENAME 压缩，只有 qcow 和 qcow2 支持

- `-B`

  force the output image to be created as a copu on write image of the specifed base image

- `-W`

  如果多线程处理器，可以加速处理

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

