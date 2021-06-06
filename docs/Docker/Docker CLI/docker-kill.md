# docker-kill

> `ENTRYPOINT` and `CMD` in the *shell* form run as a child process of `/bin/sh -c`, which does not pass signals. This means that the executable is not the container’s PID 1 and does not receive Unix signals.

`docker kill`默认发送SIGKILL（所以不提供优雅的容器退出方式），可以通过`--signal`指定发送特定的SIG到main process优雅退出容器

```
$ docker kill --signal=SIGHUP  my_container
```

