# vagrant box

## 0x0 Overview

用于管理 vagrant boxes(在 vagrant 的概念里，boxes 不等价于虚拟机，boxes 更像是 ISO + vagrantfile) 提供了一组 subcommands

- `add`
- `list`
- `outdated`
- `prune`
- `remove`
- `repackage`
- `update`

## 0x1 Subcommands

### 0x10 add

This adds a box with the given address to Vagrant. The address can be one of three things:

> 直接理解成下载 ISO

- A shorthand name from the [public catalog of available Vagrant images](https://vagrantcloud.com/boxes/search), such as "hashicorp/bionic64".
- File path or HTTP URL to a box in a [catalog](https://vagrantcloud.com/boxes/search). For HTTP, basic authentication is supported and `http_proxy` environmental variables are respected. HTTPS is also supported.
- URL directly a box file. In this case, you must specify a `--name` flag (see below) and versioning/updates will not work.

#### Optional args

- `--box-version <VALUE>`

  指定下载的 box 版本，默认下载最新版本的

- `--clean`

  下载 box 前删除之前下载失败的 cache

- `--force`

  强制重新下载 box

- `--name <VALUE>`

  指定下载出来的 box 名字

### 0x11 list

显示当前系统中添加 boxes

```
(base) 0x00 in ~ λ vagrant box list           
centos/7          (virtualbox, 2004.01)
lhq/kylin_v10_sp2 (virtualbox, 0.1)
```

### 0x12 outdated

查看 boxes 是否是最新版本的，需要在有 `Vagrantfile` 的目录下

```
(base) 0x00 in ~/Hypervisor/vagrant-machines/centos7 λ vagrant box outdated
Checking if box 'centos/7' version '2004.01' is up to date...
```

#### optional args

- `--global`

  查看所有的 boxes，需要在有 `Vagrantfile` 的目录下

  ```
  (base) 0x00 in ~/Hypervisor/vagrant-machines/centos7 λ vagrant box outdated --global
  * 'lhq/kylin_v10_sp2' for 'virtualbox' (v0.1) is up to date
  * 'centos/7' for 'virtualbox' (v2004.01) is up to date
  ```

### 0x13 prune

删除老版本的 boxes

```
(base) 0x00 in ~/Hypervisor/vagrant-machines/centos7 λ vagrant box prune            
The following boxes will be kept...
centos/7          (virtualbox, 2004.01)
lhq/kylin_v10_sp2 (virtualbox, 0.1)

Checking for older boxes...
No old versions of boxes to remove...

```

### 0x13 remove

删除指定的 boxes (所有版本)

```
(base) 0x00 in ~/Hypervisor/vagrant-machines/centos7 λ vagrant box remove lhq/kylin_v10_sp2
Removing box 'lhq/kylin_v10_sp2' (v0.1) with provider 'virtualbox'...
```

### 0x14 update

升级 box

```
(base) 0x00 in ~/Hypervisor/vagrant-machines/centos7/.vagrant λ vagrant box update
==> default: Checking for updates to 'centos/7'
    default: Latest installed version: 2004.01
    default: Version constraints: 
    default: Provider: virtualbox
    default: Architecture: "amd64"
==> default: Box 'centos/7' (v2004.01) is running the latest version.
```

**references**

[^1]:https://developer.hashicorp.com/vagrant/docs/cli/box