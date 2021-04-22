# snap

参考：

https://cn.ubuntu.com/blog/what-is-snap-application

https://snapcraft.io/about

## 概述

snap是一款开源的软件包管理器，==支持cross-platform 和 dependency-free==

## action

- download

  下载指定snap文件(`.snap`和`.assert`文件)不会执行安装，到当前目录，可以通过`--edge | --beta | --candidate | --stable`来指定下载的版本

  ```
  root in /opt λ snap download --edge hello-world
  Fetching snap "hello-world"
  Fetching assertions for "hello-world"
  Install the snap with:
     snap ack hello-world_29.assert
     snap install hello-world_29.snap
  root in /opt λ ls
  chkrootkit.tar.gz  containerd  hello-world_29.assert  hello-world_29.snap
  ```

  

