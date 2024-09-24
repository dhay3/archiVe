

# PKGBUILD

## 0x00 Overveiw

PKGBUILD 是一个文件（其实就是 bash 脚本，也就意味着可以自定义变量以及函数），配合 `makepkg` 生成 Arch 包

## 0x01 Variables

> 只记录常用的 variables 
>
> 具体查看官方文档

PKGBUILD 提供了多个变量，告诉 `makepkg` 该如何生成 Arch 包

一个最简单的 PKGBUILD 需要有 `pkgname`, `pkgver`, `pkgrel` 以及 `arch` 变量

- `pkgname(array)`

  包的名字

- `pkgver`

  包的版本，用来决定包是否需要更新

  not allowed to contain colons(:), forward  slashes(/), hyphens(-) or whitespace

  例如 2.7.1

- `pkgrel`

  当前版本包的 release number，用来决定包是否需要更新

- `pkgdesc`

  包的描述信息

  Try to keep the description to one line of text and to not use  the package’s name

- `url`

  上游地址，一般对应包的官网或者是 offical repo

- `license(array)`

  包使用的 license(s)

  例如 license=(‘GPL’ ‘FDL’)

- `install`

  安装脚本，必须和 PKGBUILD 在同目录下，会被 `makepkg` 拷贝到包中

- `changelog`

  变更日志，必须和 PKGBUILD 在同目录下，会被 `makepkg` 拷贝到包中

- `source(array)`

  源码地址(或者是包)，可以是对应文件的下载地址或者是本地文件，如果是压缩文件会被自动解压，必须和 PKGBUILD 在同目录下
  
  如果针对架构包不同，可以使用 `source_x86_64=()` 来指定，但是必须要和 `cksums_x86_64=()` 起义使用

- `noextract(array)`

  和 `source` 相同，但是对应的文件不会被自动解压

- `cksum(array)`

  对应 `source` 的 checksum(in the same order)

  如果值为 SKIP 就不会做 hash 校验

- `md5sums, sha1sums, sha224sums, sha256sums, sha384sums,    sha512sums, b2sums (arrays)`

  和 `cksum` 相同，但是使用特定的哈希算法

- `arch(array)`

  定义包可以在什么架构的 CPU 上运行

  例如 arch=('i686' 'x86_64')).

  如果值是 ‘any’ 表示可以可以在所有的架构运行

- `depends(array)`

  包正常运行的依赖，可以使用比较符号

  例如 depends=(‘nmap >= 1.0.0’)

- `makedepends(array)`

  构建包需要的以来，通常被用在源码构建的时候

- `optdepends(array)`

  可选的依赖

- `confilcts(array)`

  冲突的以来

这些 variables 都可以通过 bash 的语法来取值，例如 `$pkgname`, `$pkgver` 以及自定义的变量

除此外 `makepkg` 还可以使用如下变量

- `srcdir`

  This contains the directory where makepkg extracts, or  copies, all source files.

  `makepkg` 做 extract compression 或者是拷贝 source 时的目录， 无需知道具体的值(实际上是 PKGBUILD 同层级目录下的 src 目录)

  例如

  因为 `makepkg` 会自动解压，如果想要进入解压后的目录，可以使用 `cd $srcdir/$pkgname-$pkgversion`

- `pkgdir`

  This contains the directory where makepkg bundles the  installed package. This directory will become the root directory of your built  package. This variable should only be used in the package() function.

  表示最后包中内容安装的目录(实际上是 PKGBUILD 同层级目录下的 pkg 目录)，直接将 `$pkgdir` 理解为一个对应系统根目录的索引

  例如

  `cp "$srcdir/hello-world.sh" "$pkgdir"`

  会将 `hello-world.sh` 放到 `pkg` 目录下

  如果运行 `makepkg -si`， `hello-world.sh` 会被直接放到根目录下

  

  `cp "$srcdir/hello-world.sh" "$pkgdir/usr/bin"`

  会将 `hello-world.sh` 放到 `pkg/usr/bin` 目录下(需要先在 pkg 下创建对应的层级目录，所以可以使用 `install -D` 来替换 `cp`)

  如果运行 `makepkg -si`， `hello-world.sh` 会被直接放到 `/usr/bin/` 目录下

- `startdir`

  This contains the absolute path to the directory where  the PKGBUILD is located, which is usually the output of $(pwd) when makepkg is  started. Use of this variable is deprecated and strongly discouraged.

## 0x02 Functions

Functions 用于告诉 `makepkg` how to build and install packages (会被 `makepkg` 自动运行)，所有 bash 或者 系统可以调用的变量或者是命令，都可在 functions 中使用，所以如果需要在 functions 中自定义变量，最好使用 `local` 关键字

一个最简单的 PKGBUILD 需要有 `package()` functions

- `package()`

  The package() function is used to install files into the  directory that will become the root directory of the built package and is run  after all the optional functions listed below. This function is  run inside `$srcdir`.

  

- `build()`

  The optional build() function is used to compile and/or  adjust the source files(最常见的就是解压 deb 包) in preparation to be installed by the package()  function. This function is run inside `$srcdir`.

## 0x03 Tips

1. If you need to create any ==custom variables== for use in your build prcess, it is recommended to prefix their name with an `_`(underscore)
2. To simplify the  maintenance of PKGBUILDs, use the `$pkgname` and `$pkgver` variables when  specifying the download location in `source`
3. It is also possible to change the name of the downloaded file,    which is helpful with weird URLs and for handling multiple source files with    the same name. The syntax is: source=('filename::url').
4. To easily generate cksums, run “makepkg -g  >> PKGBUILD
5. Use `install -D` to copy file into `$pkgdir` without `mkdir` to create directory first

## 0x04 Convert deb package

如果要解压 deb 的包，可以使用 `bsdtar`

## 0x05 Examples

> 更多例子可以在 AUR 中查找

```
# Maintainer: nycex <bernhard / ithnet.com>

_fwname=aic94xx
pkgname=${_fwname}-firmware
pkgver=30
pkgrel=10
pkgdesc="Adaptec SAS 44300, 48300, 58300 Sequencer Firmware for AIC94xx driver"
url="https://storage.microsemi.com/en-us/speed/scsi/linux/${_fwname}-seq-${pkgver}-1_tar_gz.php"
license=('custom')
arch=('any')
source=("${_fwname}-seq-${pkgver}-1.tar.gz::https://download.adaptec.com/scsi/linux/${_fwname}-seq-${pkgver}-1.tar.gz"
        "LICENSE.${_fwname}")
sha256sums=('0608a919b95e65e8fe3c0cbc15f7e559716bda39a6efca863417a65f75e15478'
            '6e0dd2831a66437e87659ed31384f11bdc7720bc539d2efa063fbb7f4ac0e46c')

build() {
    bsdtar xvf "${_fwname}_seq-${pkgver}-1.noarch.rpm"
    chmod 644 "${srcdir}/lib/firmware/${_fwname}-seq.fw"
}

package() {
    install -Dm644 "${srcdir}/lib/firmware/${_fwname}-seq.fw" "${pkgdir}/usr/lib/firmware/${_fwname}-seq.fw"
    install -Dm644 "${srcdir}/LICENSE.${_fwname}" "${pkgdir}/usr/share/doc/${pkgname}/LICENSE.${_fwname}"
    install -Dm644 "${srcdir}/README-94xx.pdf" "${pkgdir}/usr/share/doc/${pkgname}/README-94xx.pdf"
}
```

**references**

[^1]:https://archlinux.org/pacman/PKGBUILD.5.html
[^2]:https://itsfoss.com/create-pkgbuild/
[^3]:https://aur.archlinux.org/cgit/aur.git/tree/PKGBUILD?h=aic94xx-firmware