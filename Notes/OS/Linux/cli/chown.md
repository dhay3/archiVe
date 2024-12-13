# Linux chown

syntax：`chown [option] [owner][:[group]] file...`

chown用于修改文件的所有者和所有组，如果只有用户ID或用户名表示只针对用户

```
chown  root /etc/
```

如果添加colon(`:`)并且带有组ID或组名表示将文件的用户和组都修改为当前用户和当前组

```
chown 1001:1002 /etc/
```

如果colon之后没有带组ID或组名但是有colon就表示将文件的用户和组都修改为当前用户和当前组

```
chown 1001： /etc/
```

如果没有用户ID或用户名表示只修改文件的组所有权为当前组

```
chown :1001 /etc/
```

## options

chown默认不修改链接文件的权限

- -h

  影响链接文件，但不影响实际指向的文件

- -H

  影响链接文件，同时影响实际指向的文件

- -R

  递归