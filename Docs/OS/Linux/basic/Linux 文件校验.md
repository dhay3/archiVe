# 文件校验

参考：

https://www.ruanyifeng.com/blog/2019/11/hash-sum.html

> 常用可以不用将hash值存储在文件中直接通过stdin就行
>
> ```
> root in /opt λ md5sum kubectl | md5sum -c 
> kubectl: OK
> ```

## 概述

![Snipaste_2020-12-16_19-09-18](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221121/Snipaste_2020-12-16_19-09-18.2yddoia49xfk.webp)

在下载一些文件，会提供用于校验的文件完整性和是否修改的文件。常见的有：`.sha1`，`.sha256`，`.sig`

## Hash校验

哈希码指的是，文件内容经过哈希函数的计算，会返回一个独一无二的字符串。哪怕原始内容只改动一个字节，哈希码也会完全不同。用户下载软件后，只要计算一下哈希码，再跟作者给出的哈希码比较一下，就会知道软件有没有被改动。

目前，常用的三种哈希函数是 MD5、SHA1 和 SHA256。其中，SHA256 最安全，SHA1 次之，MD5 垫底。一般来说，软件至少会提供其中一种哈希码。

使用`md5sum`，`sha1sum`，`sha256sum`分别校验不同的文件Hash值

默认以text mode，即生成文件的Hash值

```
md5sum inputrc.md 
233bcf56e62879e4429ad7cd2770ffa3  inputrc.md
```

如果想要校验文件的Hash文件，使用`-c`参数，==默认会检查当前目录下是否有文件的hash值与校验文件中的相同==

```
 ┌─────(root)─────(/opt) 
 └> $ md5sum -c inputrc.md.md5 
inputrc.md: OK

 ┌─────(root)─────(/opt) 
 └> $ vim inputrc.md.md5 

 ┌─────(root)─────(/opt) 
 └> $md5sum -c inputrc.md.md5 
inputrc.md: FAILED
md5sum: WARNING: 1 computed checksum did NOT match

root in / λmd5sum -c /opt/k.md5 
md5sum: kubectl: No such file or directory
kubectl: FAILED open or read
md5sum: WARNING: 1 listed file could not be read

```

## 签名验证

Hash码只能保证文件内容没有修改，但是Hash码本身也可能仿冒(Hash Collision)

文件签名能解决这个问题。软件发布时，作者用自己的私钥，对发布文件生成一个签名文件(`.sig`)，用户用作者的公钥验证签名。
