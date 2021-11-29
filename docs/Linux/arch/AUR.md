# AUR

The Arch User Repository(AUR) 是arch的社区仓库(==有可能会有恶意代码==)，用户可以下载AUR并通过`makepkg`来编译AUR package descriptions(PKGBUILDS)，然后可以通过`pacman`来安装。

> 需要确保`base-devel`完整安装，使用如下命令来安装`pacman -S --needed base-devel`

## installing

1. acquire build files

   ```
   #git clone repo
   git clone <git_clone_url>
   
   #download snapshot
   curl or wget 
   ```

2. verify the `PKGBUILD` and accompanying files

   ```
   #build files 通常包含
   cpl in ~/Downloads/typora on master λ ls
    Makefile   PKGBUILD   typora.js
   ```

3. run `makepkg` to compile the source

   ```
   cpl in ~/Downloads/typora on master λ makepkg;la      
    .      .gitignore   pkg        typora-0.10.11-1-x86_64.pkg.tar.zst
    ..     .SRCINFO     PKGBUILD   typora.js
    .git   Makefile     src        typora_0.10.11_amd64.deb 
   ```

   可以使用一些有用的参数

   - `-s`使用pacman自动下载依赖
   - `-i`编译成功后自动使用pacman安装
   - `-r`**编译成功后删除依赖**
   - `-c`编译成功后删除PKGBUILDS

4. run `pacman -U <package_file>` to isntall package

   上述的`typora-0.10.11-1-x86_64.pkg.tar.zst`就是package_file

   ```
   pacman -U typora-0.10.11-1-x86_64.pkg.tar.zst
   ```

5. check

   ```
   cpl in ~ λ pacman -Q typora 
   typora 0.10.11-1
   ```

## upgrading

在PKGBUILD中运行`git pull`即可获取最新的package，然后执行installing步骤

## PKGBUILD

PKGBUILD是一个文件，里面包含了需要依赖文件，需要注意的是由于depends没有指定版本，==默认更新时会从mirrorlists中拉取最新的版本==，可能会导致package不能被正确的安装

```
# Maintainer: Stephan Springer <buzo+arch@Lini.de>
# Contributor: Hyacinthe Cartiaux <hyacinthe.cartiaux@free.fr>
# Contributor: korjjj <korjjj+aur[at]gmail[dot]com>

pkgname=gns3-gui
pkgver=2.2.27
pkgrel=1
pkgdesc='GNS3 network simulator. Graphical user interface package.'
arch=('any')
url='https://github.com/GNS3/gns3-gui'
license=('GPL3')
groups=('gns3')
depends=(
    'desktop-file-utils'
    'python-distro'
    'python-jsonschema'
    'python-psutil'
    'python-pyqt5'
    'python-sentry_sdk'
    'python-setuptools'
    'python-sip'
    'qt5-svg'
    'qt5-websockets'
)
optdepends=(
    'gns3-server: GNS3 backend. Manages emulators such as Dynamips, VirtualBox or Qemu/KVM'
    'xterm: Default terminal emulator for CLI management of virtual instances'
    'wireshark-qt: Live packet capture')
source=("$pkgname-$pkgver.tar.gz::https://github.com/GNS3/$pkgname/archive/v$pkgver.tar.gz"
        'gns3.desktop')
sha256sums=('c5afadd2932703c38fb6d479be6966af65787ddc6abcb3f9b7c0b3f5722e0fd5'
            '51e6db5b47e6af3d008d85e8c597755369fafb75ddb2af9e79a441f943f4c166')

prepare() {
    cd "$pkgname-$pkgver"
    # Arch usually has the latest versions. Patch requirements to allow them.
    sed -i \
        -e 's|^psutil==5\.8\.0$|psutil>=5.8.0|' \
        -e 's|^sentry-sdk==1\.3\.1$|sentry-sdk>=1.3.1|' \
        -e 's|^distro==1\.6\.0$|distro>=1.6.0|' \
        requirements.txt
}

build() {
    cd "$pkgname-$pkgver"
    python setup.py build
}

package() {
  cd "$pkgname-$pkgver"
  python setup.py install --root="$pkgdir" --optimize=1
  install -Dm644 "$srcdir"/gns3.desktop "$pkgdir"/usr/share/applications/gns3.desktop
  install -Dm644 resources/images/gns3_icon_256x256.png "$pkgdir"/usr/share/pixmaps/gns3.png
  install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}

```

