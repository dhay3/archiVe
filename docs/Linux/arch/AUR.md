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
   - `-r`编译成功后删除依赖
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
