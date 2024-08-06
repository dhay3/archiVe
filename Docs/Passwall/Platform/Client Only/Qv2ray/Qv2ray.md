# Qv2ray

qv2ray 是一个跨平台的科学上网工具，虽然停更但是好有

https://github.com/Qv2ray/Qv2ray

## Installation

这边推荐使用 appiamge，安装过程中可能会遇到一些问题

## Desktop file

默认存储在 `~/.local/share/applications` 中，可以使用 `fzf` 搜索

## Plugins

参考

https://www.osuix.com/2021/01/12/qv2ray-trojan-go-%E5%AE%89%E8%A3%85%E9%85%8D%E7%BD%AE-for-win10-macos/

## Trouble shooting

### 0x01

运行 appimage 时报错

```
qv2ray: error while loading shared libraries: libcrypt.so.1: cannot open shared object file: No such file or directory
```

参考 https://stackoverflow.com/questions/71187944/dlopen-libcrypt-so-1-cannot-open-shared-object-file-no-such-file-or-directory

安装 `libxcrypt-compat`

### 0x02

Plugins -> open local plugin folder 不能正常使用，导致不能安装 plugins

```
kde-open5: /tmp/.mount_Qv2rayeSJipz/usr/optional/libstdc++/libstdc++.so.6: version `GLIBCXX_3.4.29' not found (required by /usr/lib/libKF5ConfigCore.so.5)
```

参考 https://github.com/Qv2ray/Qv2ray/issues/571

直接将 plugins 放到 qv2ray 对应的配置文件目录。如果找不到，可以使用 `fzf`