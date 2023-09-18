# Intro to the CLI



## What is CLI

Command-line interface(CLI) is the interface you use to configure Cisco devices

## How to connect to a Cisco device?

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230518/2023-05-18_19-42.5uf1rys0j5a8.png)

如果需要连接 Cisco 设备，可以通过 console 口

支持两种类型的端口

- RJ45
- USB mini-B

RJ45 口，连接的线缆大概张这样，一头是 RJ45，一头是 DB9，也被称为 ==rollover cable==

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230518/2023-05-18_19-47.56t3w27x4o3k.webp)

但是现在大多数笔记本或者是 PC 都不支持 DB9 串行口了，所以栏需要转接器

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230518/2023-05-18_19-51.3mgjw8srz6io.webp)

## Mode

有 3 种 Mode

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230518/2023-05-18_21-04.3t3pzcr5e35s.webp)

### User EXEC Mode

通过 Terminal Emulator 连接到 Cisco 设备后，就会接入 User EXEC Mode	

默认

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230518/2023-05-18_19-59.bjj8y6b88tc.webp)

- User EXEC mode is very limimted

- Users can look at some things, but can’t make any changes to the configuration

  一般只有查询的功能

### Priviledge EXEC Mode

如果我们需要更信息的查询信息，或者重启设备就需要进入 Priviledge EXEC Mode

需要使用 `enable` 命令进入

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230518/2023-05-18_20-02.opmzyt35c9c.webp)

- Provides complete access to view the device’s configuration, restart the device, etc

- Cannot change the configuration, but can change the time on the device, save the configuration file, etc

  和 User EXEC Mode 一样，不支持修改配置文件，但是可以修改部分信息

> 不管是在 User EXEC Mode 还是 Priviledge EXEC Mode
>
> 都可以使用 <kbd>?</kbd> 来查看支持的命令，<kbd>TAB</kbd> 来补全命令

### Global Configuration Mode

如果我们需要修改配置文件，就需要进入到 Global Configuration Mode

需要使用 `configure terminal` 进入

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230518/2023-05-18_20-22.6j2b3tauqbcw.webp)

## Configuration file

Cisco 上两个配置文件

- Running-config

  the current, active configuration file on the device. As you enter commands in the CLI, you edit the active configuration

  当前使用的配置文件。在 global configuration mode 下修改配置对应该文件，存储在设备的 RAM 中。如果设备断电重启，配置消失

  使用 `show running-config` 来查看配置文件

- Startup-config

  the configuration file that will be loadded upon restart of the device

  设备启动时加载的配置文件。存储在 ROM 中，设备断电重启，配置不会消失

  使用 `show startup-config` 来查看配置文件

### Saving the configuration

因为 running-config 会在设备断电重启消失，所以为了保存配置还需要将 running-config 中的内容写入到 startup-config。有 3 种方式(均在 priviledge EXEC mode)

1. `write`

2. `write memory`
3. `copy running-config startup-config`

## Password

> 可以通过 `do <COMMAND>` 的方式在，use EXEC mode 或者 global configuration mode 中使用 priviledge mode 中的命令
>
> 例如 `do show running-config`

- `enable password <PASSWORD>`

  以明文的方式存储密码

- `service password-encryption`

  将明文的密码加密，默认使用 Cisco’s proprietary encryption algorithm 加密，但是不安全

  if you enable `service password-encryption`

  - current passwords will be encrypted
  - future passowrds will be encrypted
  - the enable secret will not be encrypted

  if you disable `service password-encryption`

  - current passwords will not be encrypted
  - future passwords will be encrypted
  - the enable secret will not be encrypted

- `enable secret <PASSWORD>`

  使用 MD5 加密密码，如果同时使用了 `enable secret <PASSWORD>` 和 `enable password <PASSWORD>`，`enable secret <PASSWORD>` 会被使用，同时 `service passowrd-encryption` 也会失效

## Cancelling command

在 command 前加 `no`，例如需要取消 `service password-encryption` 就可以通过 `no service password-encryption`

**references**

[^jeremy‘s IT Lab]:https://www.youtube.com/watch?v=IYbtai7Nu2g&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=8
[^running-config vs startup-config]:https://study-ccna.com/running-startup-configuration/
