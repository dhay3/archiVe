# docker-stop

> `ENTRYPOINT` and `CMD` in the *shell* form run as a child process of `/bin/sh -c`, which does not pass signals. This means that the executable is not the container’s PID 1 and does not receive Unix signals.

docker stop用于停止运行中的容器。容器会收到SIGTERM，如果在宽限期后还没有退出会收到SIGKILL。使用`-t`参数指定在多长时间后收到SIGKILL。

```
$ docker stop my_container
```

