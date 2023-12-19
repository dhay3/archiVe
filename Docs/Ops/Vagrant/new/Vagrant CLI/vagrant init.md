# vagrant init

## 0x0 Overview

syntax

```
vagrant init [name[url]]
```

用于生成 Vagrantfile

## 0x1 Optional args

- `--box-version`

  Vagrantfile 中 box 使用的 version

## 0x2 Examples

```
 vagrant init --box-version '> 0.1.5' hashicorp/bionic64
```

**references**

[^1]:https://developer.hashicorp.com/vagrant/docs/cli/init