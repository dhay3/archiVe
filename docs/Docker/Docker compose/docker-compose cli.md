# docker-compose cli

> 查看具体command的用法使用`docker-compose help <command>`

### 通用参数

- `-f | --file <compose file>`

  指定读取的docker-compose.yml文件

- `-p | --project-name <project_name>`

  指定生成的项目名。==默认使用当前目录==

- `--profile <profile name>`

  指定docker compose运行的profile

- `--env-file <filepath>`

  指定在docker-compose.yml中使用的环境变量的文件

- `--tls`

  指定使用tls

## subCommand

### up

==Builds==, (re)creates, ==starts==, and attaches to containers for a service.

pattern：`docker-compose up [options] [service]`

如果没有指定service，默认构建和运行docker-compose.yml中所有的service

- `-d`

  以detach mode 运行容器

- `--quiet-pull`

  不会打印拉取镜像的进度

- `--no-deps`

  不会启动和有依赖(depens_on属性)的服务

- `--no-start`

  构建完容器后不会自动启动

- 