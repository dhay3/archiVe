# Linux install

## Digest

syntax

```
install [OPTION]... [-T] SOURCE DEST
install [OPTION]... SOURCE... DIRECTORY
install [OPTION]... -t DIRECTORY SOURCE...
install [OPTION]... -d DIRECTORY...
```

install 和 cp 一样都用于拷贝文件，但是 install 可以在拷贝的同时设置权限和属性

## Optional args

- `-g | --group=GROUP`

  set group ownership

- `-o | --owner=OWNER`

  set ownership

- `-m | --mode=MODE`

  set permission mode

  拷贝的同时设置权限成 755

  ```
  install -o root -g root -m 755 kubectl  /usr/local/bin/different encryption methods and formattion
  ```

- `-p | --presesrver-timestamps`

  复制文件保留源文件件的 atime 和 mtime 值

