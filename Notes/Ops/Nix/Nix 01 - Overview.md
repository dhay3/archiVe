---
createTime: 2024-09-10 14:57
tags:
  - "#hash1"
  - "#hash2"
---

# Nix 01 - Overview

## 0x01 Preface

> Nix is a _purely functional package manager_.

[Nix](https://github.com/NixOS/nix) 是一个(跨平台)纯功能性的包管理工具(Nix project 还有 Nix lanuage，NixOS 等，不要和 Nix 混淆。具体可以看 [Glossary — nix.dev  documentation](https://nix.dev/reference/glossary))。

逻辑上和 Docker 有点类似，有如下几点特性：

- Ad hoc shell environments

	In a Nix shell environment, you can immediately use any program packaged with Nix, without installing it permanently.

	可以在沙盒中运行程序，而无需安装程序

- Reproducible interpreted scripts

	A trivial script with non-trivial dependencies can be reproduced

	针对一些需要其他依赖的脚本，可以复现完整的过程(不包括安装)

- Declarative shell environments with `shell.nix`

	

## 0x02 Installation

在 Arch-based distros 中安装 Nix 非常简单，只需要执行 `pacman -S nix` 即可(默认会以 Multi-user 的模式安装)。但是你现在还不能直接使用 Nix，还需要执行如下指令

```
sudo systemclt start --enable-now nix-daemon
sudo usermod -aG nixbld $USER
```


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***References***

- [Welcome to nix.dev — nix.dev  documentation](https://nix.dev/)
- [Introduction - Nix Reference Manual](https://hydra.nixos.org/build/271691842/download/1/manual/)

***FootNotes***


