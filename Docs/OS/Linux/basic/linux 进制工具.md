# linux 进制工具

> 可以统一使用`hexdump`

## 八进制

od

```
root in /opt λ cat /dev/urandom | head -1  | od   
0000000 030661 167462 022631 123073 112410 031574 107074 030320
0000020 123654 166626 076752 175541 127157 105202 075306 134651
0000040 176305 073242 133673 154510 012231 072056 055741 031051
```

## 十六进制

xdd可以将文件或stdout转为hexdump，获取将hexdump转为文件

将文件转为hexdump，从左到右依次是address，hexdump，plian text

```
root in /opt λ xxd Dockerfile Dockerfile.hex
root in /opt λ cat Dockerfile.hex
00000000: 4152 4720 7665 7273 696f 6e3d 3a6c 6174  ARG version=:lat
00000010: 6573 740a 4652 4f4d 2063 656e 746f 7324  est.FROM centos$
00000020: 6c61 7465 7374 0a52 554e 2075 7365 7261  latest.RUN usera
00000030: 6464 2063 706c 0a55 5345 5220 6370 6c0a  dd cpl.USER cpl.
00000040: 5255 4e20 6d6b 6469 7220 2f6d 7944 6174  RUN mkdir /myDat
00000050: 610a 5255 4e20 6563 686f 2022 6865 6c6c  a.RUN echo "hell
00000060: 6f20 776f 726c 6422 203e 202f 6d79 4461  o world" > /myDa
00000070: 7461 2f67 7265 6574 696e 670a 0a0a       ta/greeting...
```

使用`-r`参数将hexdump转为二进制或明文

```
root in /opt λ xxd -r Dockerfile.hex
ARG version=:latest
FROM centos$latest
RUN useradd cpl
USER cpl
RUN mkdir /myData
RUN echo "hello world" > /myData/greeting

root in /opt λ cat /bin/echo | head -1 | xxd | xxd -r
ELF>P@▒▒@8      @@@▒888▒m▒m ▒{▒{ ▒{ ▒0 X|X| X| ▒TTTDDP▒td▒`▒`▒`DDQ▒tdR▒td▒{▒{ ▒{ /lib64/ld-linux-x86-64.so.2GNUGNUss▒5l▒ŵ,r▒L▒Ƅ,▒5▒$)V`A 59@▒&▒(▒▒M#▒MB#▒Pv▒▒K▒▒▒▒▒F-▒▒▒,cr▒bA▒9▒*Ը▒▒c*?K▒<▒s▒1▒ yj▒▒A▒▒Q)▒9▒▒U▒▒▒▒▒▒ ▒\

                                                                                                    "▒Ck▒K▒▒zc▒ b▒". ▒ 8▒▒▒L▒\(_| X▒▒▒=@!▒▒r▒ ▒▒u!▒▒N▒▒5▒▒ `/libc.so.6fflush__printf_chksetlocal
```

使用`-p`参数只输出16进制内容，使用`-u`把十六进制中的字母大写

```
root in /dev/pts λ cat /dev/mem | tail -1 | xxd -p
cat: /dev/mem: Operation not permitted
678854240767c64424086c67c6442409f666c1e20267668d442405676689
042466b94d7501006689d866e8afa3ffff6683c40c665b665e66c3ea5be0
00f030362f32332f393900fc81

root in /opt λ cat Dockerfile | xxd -u
00000000: 4152 4720 7665 7273 696F 6E3D 3A6C 6174  ARG version=:lat
00000010: 6573 740A 4652 4F4D 2063 656E 746F 7324  est.FROM centos$
00000020: 6C61 7465 7374 0A52 554E 2075 7365 7261  latest.RUN usera
```

使用`-s <seek>`参数从指定seek bytes开始

```
root in /dev/pts λ cat /dev/mem | xxd -s 10 | head -1
0000000a: 00f0 53ff 00f0 53ff 00f0 54ff 00f0 53ff  ..S...S...T...S.
```







