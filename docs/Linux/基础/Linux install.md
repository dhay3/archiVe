# Linux install

install 命令用于拷贝文件的同时设置权限和属性

syntax；`install [options] <src> <dest>`

`-m`拷贝的同时设置权限，默认755

```
install -o root -g root -m 755 kubectl  /usr/local/bin/
```

