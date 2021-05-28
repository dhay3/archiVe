# qemu-img

## commands

### create

syntax：`qemu-img [-f fmt] filename size`

以指定的size生成名为filename的disk image。==只有在被用到的时候才会分配磁盘空间==。（可以使用`du`或`qemu-img info`来查看）

```
cpl in ~/software/kvm λ qemu-img create -f qcow2 win.img 100M
Formatting 'win.img', fmt=qcow2 size=104857600 cluster_size=65536 lazy_refcounts=off refcount_bits=16
```

同样的也可以使用`dd`命令(==直接分配磁盘空间==)

```bash
cpl in ~/software/kvm λ qemu-img create -f qcow2 win.img 100M
Formatting 'win.img', fmt=qcow2 size=104857600 cluster_size=65536 lazy_refcounts=off refcount_bits=16

cpl in ~/software/kvm λ dd if=/dev/zero of=win2.img bs=100M count=1 
1+0 records in
1+0 records out
104857600 bytes (105 MB, 100 MiB) copied, 0.0974888 s, 1.1 GB/s

cpl in ~/software/kvm λ du -haBM
1M	./win.img
100M	./win2.img
101M	.
```

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

syntax：`qeum-img resize [-f fmt] [--shrink] <filename> [+|- ]size`

修改disk image的大小，如果需要缩小必须使用`--shrink`

```bash
cpl in ~/software/kvm λ qemu-img resize win.img +10M
Image resized.
```

### bench

syntax：`qeum-img bench [-d depth] [-c count] [-s buffer_size] [-w]`

对disk image做IO测试，分别指定深度，IO请求次数，每次IO请求的大小。默认只做read IO，使用`-w`表示做write IO

```bash
qemu-img bench -c 10 -d 0 -s 1M win.img -w
```

## format

qemu支持的file format有如下几种类型

- raw

  默认使用的类型，可以将这种格式的file导入到其他的emulators

- qcow2

  生成的镜像更小，常常用在不支持linux fs的系统(例如windows)

- other

  包括其他一些兼容的格式，例如VMDK，VDI，VHD，qcow1和QED等

