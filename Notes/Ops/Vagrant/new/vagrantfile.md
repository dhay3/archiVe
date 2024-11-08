# vagrantfile

## 0x0 Overview

> Vagrantfile 使用 Ruby 语法

Vagrantfile 是用于配置虚拟的文件，决定了 vagrant 如何创建以及修改虚拟机

## 0x1 Lookup path

假设当前目录为 `/home/mitchellh/projects/foo` 当使用任何 vagrant 命令时，会搜索如下路径中的 Vagrantfile 用于构建虚拟机

```
/home/mitchellh/projects/foo/Vagrantfile
/home/mitchellh/projects/Vagrantfile
/home/mitchellh/Vagrantfile
/home/Vagrantfile
/Vagrantfile
```

## 0x2 Vagrantfile template

```

```



**references**

[^1]:https://developer.hashicorp.com/vagrant/docs/vagrantfile