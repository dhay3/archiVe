ac---
createTime: 2024-11-14 18:05
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# How to Secure Your VPS

## 0x01 Preface

> Linux 服务器的安全防护是一个纷繁复杂的巨大课题。无数的网站、APP、服务、甚至线下基础设施都建立在 Linux 的基石之上，这背后牵涉到巨大的经济利益和商业价值，当然也就就意味着黑灰产有巨大的攻击动力。[^1]

在你从 Cloud Provider 购买 VPS 的第一天起，只要 VPS 开机就会受到来自互联网的各种攻击，例如

- 端口、漏洞扫描
- ssh 暴力破解
- 针对应用 CVE 的利用
- etc.

> [!note]
> 没有绝对的安全，只有相对的安全

所以第一件要做的也是必须做的事，就是加强你的 VPS 安全

## 0x02 Full Update

不管是 Kernel 还是 Softwares 都是人开发的，所以一定会有逻辑上的漏洞。这些被发现的漏洞在学术上称为 CVE[^2]，通常开发者会针对 CVE 发布更新。所以为了保证 CVE 不被利用，应该要做 Full Update

### 0x02a LTS OS[^4]

> [!note]
> 通常 LTS OS Term 为 5 years

针对 Long Term Support(LTS) Release OS，只要是在 Term 内发布的更新，OS 可以通过包管理镜像源升级 Kernel 和 Softwares。而如果发布的更新不在 Term 内，OS 就不能通过包管理镜像源获取到更新，但是可以手动安装或者编译安装更新

例如

> Ubuntu 14.04 LTS (Trusty Tahr) End of Standard Support on 2019-04

Ubuntu 14.04 LTS 可以通过如下命令更新 2019-04 前发布的 Kernel 和 Softwares

```
sudo apt update
sudo apt upgrade
```

但是不能通过 `apt` 获取 2019-04 后发布的 Kernel 和 Software

#### Point Release

在 LTS OS 中还有一个 Point Release 的概念，可以看作是 a snapshot of previous LTS to next LTS

### 0x02b Rolling Release OS[^5]

针对 Rolling Release OS，没有 Term 的逻辑，OS 可以随时通过包管理器镜像升级 Kernel 和 Softwares

例如 

Arch 可以通过如下命令更新 Kernel 和 Softwares

```
sudo pacman -Syyu
```

## 0x03 Account

一个安全的账户虽然防止不了 Priviledge Escalation Attack，但是很大程度上可以防止 Burte Force Attack

而决定一个账户是否安全，有 2 个 factors

- username
- password

### 0x03a Username

> [!note]
> 添加运维用户是必须的，但是用户名的随机性非强制要求，但推荐

添加一个新运维用户（避免直接使用 root），用户名要具有随机性，不能使用类似 `amdin`，`backup`，`xray` 等具有指向性的用户名

可以使用一些无规则的用户名，例如

```
useradd solovki5x
```

### 0x03b Password[^2]

一个高强度的密码，可以让 brute force 几乎不可能完成。而衡量密码强度的标准，在 cryptography 中被称为 entropy(中文也叫作 熵)，随机性越高熵值越高，密码强度就越高

我们可以通过 `pwgen` 来生成一个高熵值的密码

```
pwgen -y -s 14 1
```

当然也可以使用类似 1password，bitwarden 之类的软件来生成并记录密码(但是尽量避免使用浏览器插件)

### 0x03c Expire

> [!note]
> 非强制性，但推荐

一个

### 0x03d Policy

密码更新的周期，以及字符数


## 0x04 SSH

通常 Cloud Providers 会提供 2 种登入 VPS 的方式

- VNC
- SSH

通常 VNC 只能从控制台登入，而 SSH 默认对公网开放，方便你管理 VPS，但是这样其他人同样也可以访问你服务器的 SSH

所以你可以在 `lastb` 中看到各种想要爆破你服务器 ssh 登入失败的信息

```
$ lastb | tac
...
vnc      ssh:notty    104.243.34.203   Wed Nov 13 01:28 - 01:28  (00:00)
vnc      ssh:notty    104.243.34.203   Wed Nov 13 01:28 - 01:28  (00:00)
vnc      ssh:notty    104.243.34.203   Wed Nov 13 01:28 - 01:28  (00:00)
hadoop   ssh:notty    104.243.34.203   Wed Nov 13 01:28 - 01:28  (00:00)
hadoop   ssh:notty    104.243.34.203   Wed Nov 13 01:28 - 01:28  (00:00)
hadoop   ssh:notty    104.243.34.203   Wed Nov 13 01:29 - 01:29  (00:00)
hadoop   ssh:notty    104.243.34.203   Wed Nov 13 01:29 - 01:29  (00:00)
hadoop   ssh:notty    104.243.34.203   Wed Nov 13 01:29 - 01:29  (00:00)
hadoop   ssh:notty    104.243.34.203   Wed Nov 13 01:29 - 01:29  (00:00)
admin    ssh:notty    104.243.34.203   Wed Nov 13 01:29 - 01:29  (00:00)
admin    ssh:notty    104.243.34.203   Wed Nov 13 01:29 - 01:29  (00:00)
admin    ssh:notty    104.243.34.203   Wed Nov 13 01:29 - 01:29  (00:00)
admin    ssh:notty    104.243.34.203   Wed Nov 13 01:29 - 01:29  (00:00)
admin    ssh:notty    104.243.34.203   Wed Nov 13 01:29 - 01:29  (00:00)
admin    ssh:notty    104.243.34.203   Wed Nov 13 01:29 - 01:29  (00:00)
admin    ssh:notty    104.243.34.203   Wed Nov 13 01:30 - 01:30  (00:00)
admin    ssh:notty    104.243.34.203   Wed Nov 13 01:30 - 01:30  (00:00)
admin    ssh:notty    104.243.34.203   Wed Nov 13 01:30 - 01:30  (00:00)
admin    ssh:notty    104.243.34.203   Wed Nov 13 01:30 - 01:30  (00:00)
admin    ssh:notty    104.243.34.203   Wed Nov 13 01:30 - 01:30  (00:00)
admin    ssh:notty    104.243.34.203   Wed Nov 13 01:30 - 01:30  (00:00)
root     ssh:notty    104.243.34.203   Wed Nov 13 01:30 - 01:30  (00:00)
root     ssh:notty    104.243.34.203   Wed Nov 13 01:30 - 01:30  (00:00)
root     ssh:notty    104.243.34.203   Wed Nov 13 01:31 - 01:31  (00:00)
root     ssh:notty    104.243.34.203   Wed Nov 13 01:31 - 01:31  (00:00)
root     ssh:notty    104.243.34.203   Wed Nov 13 01:31 - 01:31  (00:00)
root     ssh:notty    104.243.34.203   Wed Nov 13 01:31 - 01:31  (00:00)
root     ssh:notty    104.243.34.203   Wed Nov 13 01:31 - 01:31  (00:00)
root     ssh:notty    104.243.34.203   Wed Nov 13 01:32 - 01:32  (00:00)
```

### Port

22 是 SSH 的默认监听端口，如果你不修改，相信你会在 `lastb` 看到更多的登入失败信息 

我们可以通过修改 `/etc/ssh/sshd_config` 中的 `Port` 指令来修改默认端口，例如使用 65522

```
Port 6
```



### Deny Root

### Authentication Method

user certi publickey and keyboardinteractive

## Sudo/Polkit


## 0x05 Firewall

最重要也是最复杂的就是网络策略

## 0x06 Audit

## 0x07 PAM

## 0x08 Apparmor/SELinux

## 0x04 Miscellaneous

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [Security - ArchWiki](https://wiki.archlinux.org/title/Security)
- [【第 4 章】安全防护篇 | Project X](https://xtls.github.io/document/level-0/ch04-security.html)
- [Understanding Entropy: The Key to Secure Cryptography and Randomness | Netdata](https://www.netdata.cloud/blog/understanding-entropy-the-key-to-secure-cryptography-and-randomness/)
- [40 Linux Server Hardening Security Tips \[2024 edition\] - nixCraft](https://www.cyberciti.biz/tips/linux-security.html)

***References***

[^1]:[【第 4 章】安全防护篇 | Project X](https://xtls.github.io/document/level-0/ch04-security.html)
[^2]:[Common Vulnerabilities and Exposures - Wikipedia](https://en.wikipedia.org/wiki/Common_Vulnerabilities_and_Exposures)
[^3]:[Security - ArchWiki#Passwords](https://wiki.archlinux.org/title/Security#Passwords)
[^4]:[LTS - Ubuntu Wiki](https://wiki.ubuntu.com/LTS)
[^5]:[What is a Rolling Release Distribution?](https://itsfoss.com/rolling-release/)