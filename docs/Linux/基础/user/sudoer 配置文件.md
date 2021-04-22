# sudoer 配置文件

> `man sudoers`，`man sudo.conf`

sudoer policy 默认的配置文件在`/etc/sudo.conf`，包含如下四个directives

- Plugin
- Path
- Set
- Debug

## Plugin

提供input/output logging的二进制插件，可以使用绝对路径指定，也可以使用相对路径(`/usr/lib/sudo`)。

```
Plugin sudoers_policy sudoers.so
#等价
Plugin sudoers_policy /usr/lib/sudo/sudoers.so
```

如果没有使用`Plugin`，sudoer会使用默认的security policy等价于

```
Plugin sudoers_policy sudoers.so
Plugin sudoers_io sudoers.so
```

## Path 

用于设置sudo一些特殊参数使用的文件路径。有如下几个值，具体查看manual page

askpass，devsearch，neoxec，plugin_dir，sesh

## Set

用于设置一些前端的设置

- disable_coredump

  对sudo进行coredump，默认关闭coredump

  ```
   Set disable_coredump false
  ```

## Debug

用于debug sudo

## Aliases

有四种aliases可以被用在配置文件中，分别是User_Alias，Runas_Alias，Host_Alias和Cmnd_Alias

## 例子

