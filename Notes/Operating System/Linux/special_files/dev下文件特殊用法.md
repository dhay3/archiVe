# /dev下的文件特殊用法

参考：

http://www.linux-databook.info/?page_id=5108

## tty

通过tty文件让不同tty之间通信

我们在tty1中输入如下命令

```
root# echo "msg from tty1 to tty2" > /dev/tty2
```

tty2端就可接受到信息

```
root# msg from tty1 to tty2
```

这样的方法通用适用于pts

```
root in ~ λ w
 14:46:26 up 20:47,  4 users,  load average: 0.00, 0.00, 0.00
USER     TTY      FROM             LOGIN@   IDLE   JCPU   PCPU WHAT
root     pts/0    115.233.222.34   12:51    1:54m  0.09s  0.09s -zsh
root     pts/1    115.233.222.34   14:39    0.00s  0.31s  0.00s w
root     pts/2    115.233.222.34   12:37    1:47m  0.28s  0.28s -zsh
cpl      pts/3    115.233.222.34   14:45   36.00s  0.06s  0.06s -zsh

root in /dev/pts λ echo "hello to pts/3" >> /dev/pts/3
-----
#pts3
cpl# hello to pts/3
```

## lp0

这是打印机的文件，可以通过该文件快速打印

```
cat test.pdf > /dev/usb/lp0
```

## mem

直接对应系统的RAM

```
root in /dev/pts λ cat /dev/mem | tail -4 | xxd
cat: /dev/mem: Operation not permitted
00000000: 750c 66b8 0d00 0000 66e8 2a73 ffff 660f  u.f.....f.*s..f.
00000010: bec3 665b e921 7366 9066 6864 7c00 00e9  ..f[.!sf.fhd|...
00000020: 57db e932 f166 6877 7c00 00e9 4bdb e972  W..2.fhw|...K..r
00000030: c280 fc89 0f84 14da 6668 fcf0 0000 e938  ........fh.....8
00000040: db66 5566 5766 5666 5366 5266 89c3 678a  .fUfWfVfSfRf..g.


root in /dev/pts λ dd if=/dev/mem bs=256 count=2 | xxd
2+0 records in
2+0 records out
512 bytes copied, 0.000123416 s, 4.1 MB/s
00000000: 53ff 00f0 53ff 00f0 c3e2 00f0 53ff 00f0  S...S.......S...
00000010: 53ff 00f0 54ff 00f0 53ff 00f0 53ff 00f0  S...T...S...S...
00000020: a5fe 00f0 87e9 00f0 70d4 00f0 70d4 00f0  ........p...p...
00000030: 70d4 00f0 70d4 00f0 57ef 00f0 70d4 00f0  p...p...W...p...
00000040: 474d 00c0 4df8 00f0 41f8 00f0 fee3 00f0  GM..M...A.......
```

