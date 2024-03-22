# Linux install

## Digest

syntax

```
install [OPTION]... [-T] SOURCE DEST
install [OPTION]... SOURCE... DIRECTORY
install [OPTION]... -t DIRECTORY SOURCE...
install [OPTION]... -d DIRECTORY...
```

`install` 命令用于拷贝文件的同时设置权限和属性，和 `cp` 不一样的是 `cp` 最多只能保留原文件的一些权限和属性，但是 `install` 可以手动指定（在 PKGBUILD 中常见）

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

- `-D`

  create all leading components of DEST except the last

  拷贝文件的同时，如果 DEST 中有不存在的目录，会被直接创建，但是最后一个层级的内容不会被创建，通常用于复制文件

  ```
  [vagrant@arch opt]$ ls
  hello-world.sh
  [vagrant@arch opt]$ sudo install -D hello-world.sh /opt/a/b/hello-world.sh
  [vagrant@arch opt]$ tree
  .
  ├── a
  │   └── b
  │       └── hello-world.sh
  └── hello-world.sh
  ```

- `-d`

  将 SOURCE 和 DEST 中的所有层级都当成目录，如果 DEST 其他有层级没有对应的目录就会创建目录

  ```
  [vagrant@arch opt]$ ls
  a
  [vagrant@arch opt]$ sudo install -d a /opt/b/c/d
  [vagrant@arch opt]$ tree
  .
  ├── a
  └── b
      └── c
          └── d
  ```

- `-v | --verbose`

  输出详细信息
