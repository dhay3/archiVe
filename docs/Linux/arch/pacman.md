# pacman

参考：

https://wiki.manjaro.org/index.php/Pacman_Overview

https://wiki.archlinux.org/title/Pacman_(%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87)

package manager utility(pacman)，manjaro和arch都使用它作为包管理器。

syntax：`pacman <operations> [options] [target]`

target通常package name或是URI

> 注意manjaro只能使用official Manajaro repo，但是也可以使用ARU，flatpaks，snaps或appImages

## operations

### -D | --database

==查看和修改pacman的db==

```bash
cpl in ~ λ sudo pacman -Dv
Root      : /
Conf File : /etc/pacman.conf
DB Path   : /var/lib/pacman/
Cache Dirs: /var/cache/pacman/pkg/  
Hook Dirs : /usr/share/libalpm/hooks/  /etc/pacman.d/hooks/  
Lock File : /var/lib/pacman/db.lck
Log File  : /var/log/pacman.log
GPG Dir   : /etc/pacman.d/gnupg/
Targets   : 无
```

### -Q | --query

查询已经安装的pkg，如果没有指定pkg，默认查询所有安装的pkg

- -i | -q

  查看pkg的详细信息，`-q` show less info for query and search 

- -l

  查看pkg安装的所有内容(ps. 查看文件路径)

- -m | -n

  `-m`查看不是通过sync dbs安装的pkg(ps. not offical repo installed pkg，==即从AUR安装的==)，`-n`查看通过sync dbs安装的pkg

- -s

  从pkg name和description中查找，通常和`-q`一起使用

- -u

  list packages can be upgraded

- -t

  列出不被需要的pkg

- -d

  列出所有被作为denpends的pkg

### -S | --sync

sync form remote server 。和apt一样如果一个package存在于多个pkg中，需要指明repo，例如`pacman -S <reponame>/<pkg>`。`pacman -S <pkg>`==同时做upgrade和install==

同时也可以指定版本

`pacman -S "bash>=3.2"`

- -i 

  查看pkg的详细信息

- -c

  删除pacman保存的cache

- -s

  从remote server上查看指定pkg

  ```
  cpl in /etc/pacman.d λ pacman -Ssq firefox
  ```

- -u

  升级==所有==到达生命周期的pkg

- -y

  更新pkg db，通常和`-u`一起使用(==更新所有的pkg==)

  ```
  pacman -Syu
  ```

  使用`-yy`表示强制更新
  
  ```
  pacman -Syy
  ```

### -R | --remove

删除pkg，默认不会删除配置文件(和`apt-get remove`)类似，所有的配置文件以`.pacsave`结尾。使用`--nosave`等价于`apt-get purne`。

- -c

  删除pkg时，删除depends(==所有的依赖==)

- -n | --nosave

  ==删除pkg时同时删除配置文件==

- -s

  删除pkg中不被==其他包==需要的depends

- -u 

  删除pkg(==本包==)不再需要的depends

- -d

  删除依赖

  ```
  cpl in ~/Downloads/gns3-gui/pkg/gns3-gui/usr/bin on master ● λ sudo pacman -Rdd python-distro 
  
  Packages (1) python-distro-1.6.0-1
  
  Total Removed Size:  0.15 MiB
  
  :: Do you want to remove these packages? [Y/n] 
  ```

### -U | --upgrade

安装一个本地包，通常和`makepkg`一起使用安装AUR BUILDPKG

```
pacman -U /path/to/package/package_name-version.pkg.tar.xz
```

- `--overwirete <glob>`

  如果有文件冲突，直接复写冲突的文件，通常和`-U`一起使用

> 如果一个文件不能执行，很有可能就是依赖的问题，需要downgrade

https://wiki.archlinux.org/title/downgrading_packages

可以从`/var/cache/pacman/pkg`对包降级，例如

`pacman -U /var/cahe/pacman/pkg/python-distro-1.6.0-1-any.pkg.tar.zst`

如果cache被删除了，可以使用如下方法来降级

先修改`/etc/pacman.conf`

例如这个将时间置位2014.03.30

```
[core]
SigLevel = PackageRequired
Server=https://archive.archlinux.org/repos/2014/03/30/$repo/os/$arch

[extra]
SigLevel = PackageRequired
Server=https://archive.archlinux.org/repos/2014/03/30/$repo/os/$arch

[community]
SigLevel = PackageRequired
Server=https://archive.archlinux.org/repos/2014/03/30/$repo/os/$arch
```

或者是替换`/etc/pacman.d/mirrolist`中的mirror

```
##                                                                              
## Arch Linux repository mirrorlist                                             
## Generated on 2042-01-01                                                      
##
Server=https://archive.archlinux.org/repos/2014/03/30/$repo/os/$arch
```

然后在运用`pacman -Sy python-distro`

## options

- `--noconfirm`

  所有选项都选yes

## remove orphans

pacman没有apt的autoremove的功能，但是可以通过`pacman -Rs $(pacman -Qdtq)`来达到autoremove，也可以使用如下的脚本：

```shell
LOOPFLAG=0
PACMAN=$(which pacman 2> /dev/null)
SUDO=$(which sudo 2> /dev/null)
 
case "$1" in                      
  -l)
  echo -e "
  \r** UNNEEDED DEPENDENCIES **
  \r-> checking dependencies...
  "
  $PACMAN -Qdtq
  if [ "$?" = 1 ]; then
    echo -e "-> Your system doesn't have unneeded dependencies. \n"
  fi 
  ;;
  -r)
  while [ "$LOOPFLAG" = 0 ]
  do
    echo -e "
    \r** UNNEEDED DEPENDENCIES **
    \r-> checking dependencies...
    "
    $PACMAN -Qdtq
    if [ "$?" = 0 ]; then
      echo ""
      echo -n "Remove these packages with pacman? [Y/n] "
      read OPTION 
      if [ "$OPTION" = "y" ] || [ "$OPTION" = "" ]; then
        echo -n "-> "
        if [ -f $SUDO ]; then
          $SUDO $PACMAN -Rn $($PACMAN -Qdtq)
          if [ "$?" != 0 ]; then
            echo -e "-> Dependencies skipped... next dependencies... \n"
          else
            echo -e "-> Unneeded dependencie(s) sucessfully removed. \n"
          fi
        else
          $PACMAN -Rn $($PACMAN -Qdtq)
          echo -e "-> Unneeded dependencie(s) sucessfully removed. \n"
        fi
      elif [ "$OPTION" = "n" ]; then
        exit 0
      fi  
    else
      LOOPFLAG=1
      echo -e "-> Your system doesn't have unneeded dependencies. \n"
    fi
  done
  ;;
  -ra)
  $PACMAN -Qdtq > /dev/null
  if [ "$?" = 1 ]; then
    echo -e "
    \r** UNNEEDED DEPENDENCIES **
    \r-> checking dependencies...
    "
    echo -e "-> Your system doesn't have unneeded dependencies. \n"    
  else  
    echo -e "\n** UNNEEDED DEPENDENCIES - RECURSIVE **"
    echo -n "-> "
    if [ -f $SUDO ]; then
       $SUDO $PACMAN -Rsn $($PACMAN -Qdtq)
    else
       $PACMAN -Rsn $($PACMAN -Qdtq)
    fi
  fi
  ;;
  *)
    echo "Usage: $0 {-l <list> | -r <remove> | -ra <remove all - recursive>}"
esac
exit 0
```



## Pacman is currently in using

https://bbs.archlinux.org/viewtopic.php?id=67729

`rm /var/lib/paman/db.lck`

## symbol lookup error 

https://forum.manjaro.org/t/symbol-lookup-error/73596

可能是没有更新完全导致的，使用`pacman -Syu`更新所有包即可，或是更新未完成的包

