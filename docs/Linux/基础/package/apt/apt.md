# apt 

参考：

https://www.debian.org/doc/manuals/apt-howto/

主要由`apt-get`和`apt-cache`组成，可以直接使用`apt command`

```
cpl in ~/Videos λ apt list --installed nmap
Listing... Done
nmap/focal,now 7.80+dfsg1-2build1 amd64 [installed]

cpl in ~/Videos λ apt list --upgradable nmap 
Listing... Done
```

## apt-get

syntax: `apt-get [options] command`  

### options

- -f  | --fix-broken

  ==当安装时出现缺少依赖时，使用该参数可以自动下载依赖==

  ```
  apt-get -f install
  ```

- -q | --quiet

  supressed output，最多2个q

  ```
  sudo apt-get -qq update
  ```

- -y | --yes

  assume yes as answer to all prompts

- --only-upgrade

  和install coommand一起使用，对指定package升级

  ```
  sudo apt-get --only-upgrade install nmap
  ```

- -o | --options

  set a configuartion option

### command

- update

  从`sources.list`和`sources.list.d`中更新数据源

- upgrade | dist-upgrade

  更新==所有==本系统安装的packages，dist-upgrade在更新时还解决packages之间的冲突

- `install <packages> `

  安装一个或多个packages，如果依赖不存在会自动安装。报名不一定全称(例如`apt-utils_2.0.5_amd64.deb`等价与`apt-utils`)。如果需要安装指定的版本可以通过如下方式

  ```
  apt-get install apt-utils/focal-updates
  ```

  packages name 支持 POSIX regex(basic regex)

- `reinstall <packages>`

  重新安装，等价与`--reinstall`

- `remove <packages> | autoremove`

  卸载packages，但是==不会删除配置文件==

- `purge <packages>`

  卸载packages同时，也==删除配置文件==。等价与`apt-get --purge hhremove`

- `download <packages> `

  下载制定的二进制文件到当前目录

- `clean <packages> | autoreclean`

  删除本地索引库中不能下载和无用的索引

- `changelog <packages>`

  查看packages的changelog

  ```
  sudo apt-get changelog nmap
  ```

- indextargets

  以deb822格式显示，所有的数据源

## apt-cache

### options

- --names-only

  只按照packages name来寻找

  ```
  sudo apt-cache search --names-only nmap
  ```

### command

- gencaches

  生成apt package cache，默认在使用任何command时被调用

- showpkg

  查看指定packages的具体信息

  ```
  sudo apt-cache showpkg nmap
  package:...
  versions:...
  MD5:...
  depends and dependencies:...
  ```

- show

  查看指定packages的具体信息

  ```
  udo apt-cache show nmap
  ...
  md5sum and sha1 adn sha256:...
  homepage:...
  description:...
  ```

- `search <regex>`

  在package names 和 descriptions中，==通过 regex 寻找包==

- ==pkgname==

  查看apt package cache中所有的package name