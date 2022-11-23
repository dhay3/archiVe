# Linux install

## Digest

syntax

```
install [OPTION]... [-T] SOURCE DEST
install [OPTION]... SOURCE... DIRECTORY
install [OPTION]... -t DIRECTORY SOURCE...
install [OPTION]... -d DIRECTORY...
```

`install` 命令用于拷贝文件的同时设置权限和属性，和 `cp` 不一样的是 `cp` 最多只能保留原文件的一些权限和属性，但是 `install` 可以手动指定

## Optional args

- `-o | --owner=onwer`

  拷贝文件时，同时设置 onwership （只能被 root 使用）

- `-m | --mode=mode`

  拷贝文件时，同时设置 mode permission

  ```
  install -o root -g root -m 755 kubectl  /usr/local/bin/different encryption methods and formattion
  ```

- `-g | --group=group`

  拷贝文件时，同时设置 group ownership
