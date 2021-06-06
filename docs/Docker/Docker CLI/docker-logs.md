# docker-logs

## 概述

pattern：`docker logs [options] <containerId>`

docker logs 用于展示指定容器的日志

## 参数

- `-f | --follow`

  以流模式展示日志(日志追踪)

  ```
  $ docker run --name test -d busybox sh -c "while true; do $(echo date); sleep 1; done"
  $ date
  Tue 14 Nov 2017 16:40:00 CET
  $ docker logs -f --until=2s test
  Tue 14 Nov 2017 16:40:00 CET
  Tue 14 Nov 2017 16:40:01 CET
  Tue 14 Nov 2017 16:40:02 CET
  ```

- `--since`

  查看指定时间之后的日志

  ```
  root in /etc/ssh λ docker logs --since 2m test
  Mon Feb  1 11:56:33 UTC 2021
  Mon Feb  1 11:56:34 UTC 2021
  Mon Feb  1 11:56:35 UTC 2021
  Mon Feb  1 11:56:36 UTC 2021
  ```

  