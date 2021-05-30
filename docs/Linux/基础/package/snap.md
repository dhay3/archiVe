# snap

参考：

https://cn.ubuntu.com/blog/what-is-snap-application

https://snapcraft.io/about

## 概述

snap是一款开源的软件包管理器，==支持cross-platform 和 dependency-free==

用户可以通过snapcraft将程序打包成snap的格式

snap通过channel来区分包，可以选择edge，beta，candidate，stable等

## action

- debug paths

  snap使用的目录

  ```
  root in /home/ubuntu λ snap debug paths
  SNAPD_MOUNT=/snap
  SNAPD_BIN=/snap/bin
  SNAPD_LIBEXEC=/usr/lib/snapd
  ```

- find

  在snap store中查询指定的包，如果publisher是被认证的，会在包名后面有一个绿勾

  ```
  root in /opt λ snap find --narrow hello-world
  Name                             Version              Publisher             Notes    Summary
  hello-world                      6.4                  canonical✓            -        The 'hello-world' of snaps
  ```

  `--narrow`表示只从stable中查询snap

- info

  查看snap的详细信息

  ```
  root in /opt λ snap info hello-world
  name:      hello-world
  summary:   The 'hello-world' of snaps
  publisher: Canonical✓
  store-url: https://snapcraft.io/hello-world
  contact:   snaps@canonical.com
  license:   MIT
  description: |
    This is a simple hello world example.
  snap-id: buPKUD3TKqCOgLEjjHx5kSiCpIs5cMuQ
  channels:
    latest/stable:    6.4 2019-04-17 (29) 20kB -
    latest/candidate: 6.4 2019-04-17 (29) 20kB -
    latest/beta:      6.4 2019-04-17 (29) 20kB -
    latest/edge:      6.4 2019-04-17 (29) 20kB -
  ```

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

- install

  安装指定snap包，第一次下载snap包时会自动下载snap core做为snap的运行环境，可以指定channel

  ```
  root in /opt λ snap install hello-world
  2021-04-22T17:44:32+08:00 INFO Waiting for automatic snapd restart...
  hello-world 6.4 from Canonical✓ installed
  
  #snap通过指定channel来安装指定版本，通过snap info 来查看
  root in /home/ubuntu λ snap isntall --channel esr/stable firefox
  ```

- list

  查看已安装的snaps

  ```
  root in /opt λ snap list
  Name         Version    Rev    Tracking       Publisher   Notes
  core         16-2.49.2  10958  latest/stable  canonical✓  core
  hello-world  6.4        29     latest/stable  canonical✓  jailmode
  ```

- changes

  查看snap的操作

  ```
  root in /home/ubuntu λ snap changes
  ID   Status  Spawn               Ready               Summary
  2    Done    today at 17:43 CST  today at 17:44 CST  Install "hello-world" snap
  3    Done    today at 17:43 CST  today at 17:43 CST  Initialize device
  4    Done    today at 17:47 CST  today at 17:47 CST  Remove "hello-world" snap
  5    Done    today at 17:47 CST  today at 17:47 CST  Install "hello-world" snap
  6    Done    today at 18:30 CST  today at 18:30 CST  Remove "hello-world" snap
  ```

- pack

  将按照格式编写的snapfile打包成snap

### 版本控制

- version

  查看当前OS和snap的系统

  ```
  root in /etc/sudoers.d λ snap version
  snap    2.49.2
  snapd   2.49.2
  series  16
  ubuntu  18.04
  kernel  4.15.0-118-generic
  ```

- refresh

  根据Tacking列来更新snap，如果没有指定snap默认更新所有的snaps。

  ```
  root in /home/ubuntu λ snap refresh hello-world
  snap "hello-world" has no updates available
  #可以指定channel
  ```

- revert

  回退snap版本，如果没有指定`--revision`

  ```
  
  ```

- remove

  删除snap，`--purge`删除snap时同时删除snapshot

  ```
  root in /home/ubuntu λ snap remove --purge  hello-world
  hello-world removed
  ```

### 快照

- forget

  删除一个snapshot

- retore

  回退snap到指定snap

- save

  对当前状态snap保存snapshot

  ```
  root in /etc/sudoers.d λ snap save hello-world
  Set  Snap         Age    Version  Rev  Size    Notes
  2    hello-world  103ms  6.4      29     123B  -
  ```

- saved

  展示当前存储的快照

  ```
  root in /etc/sudoers.d λ snap saved
  Set  Snap         Age    Version  Rev  Size    Notes
  1    hello-world  3d22h  6.4      29     125B  auto
  ```

### 运行

- run

  ```
  root in /etc/sudoers.d λ snap run hello-world
  Hello World!
  ```

  

