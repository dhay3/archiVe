# docker-exec 

> 通常使用`-it`参数代开一个互动的shell

## 概述

pattern：`docker exec [options] containerId command`

docker exec 用于在指定运行的容器中运行命令，需要以`shell`执行命令。否则不会运行。

```
docker exec -ti my_container "echo a && echo b" #不运行
docker exec -ti my_container sh -c "echo a && echo b" #运行
```

==注意，如果跟的是command_string，必须添加`-c`参数==

## 参数

- `-d`

  以守护进程的模式运行

- `-i`

  保持stdin打开的状态

- `-t`

  分配tty

- `--workdir`

  指定运行指令的目录

  ```
  $ docker exec -it -w /root ubuntu_bash pwd
  /root
  ```

  