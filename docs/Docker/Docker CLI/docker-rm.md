# docker-rm

`docker rm`用于删除容器

- 删除所有停止的容器，可以使用`-f`参数发送SIGKILL，强制停止并删除所有容器。

  ```
  docker rm $(docker ps -a -q)
  ```

- 删除docker自动生成的匿名卷

  ```
  docker rm -v hello
  ```

  