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

### config

检查compose文件

```
root in /usr/local/\/composetest λ docker-compose config
services:
  redis:
    image: redis:alpine
  web:
    build:
      context: /usr/local/\/composetest
    environment:
      FLASK_ENV: development
    ports:
    - published: 5000
      target: 5000
    volumes:
    - /usr/local/\/composetest:/code:rw
version: '3.9'
```

### up

==Builds==, (re)creates, ==starts==, and attaches to containers for a service.

pattern：`docker-compose up [options] [service]`

如果没有指定service，默认构建和运行docker-compose.yml中所有的service。==同时生成一个包含所有服务的镜像==

- `-d`

  以detach mode 运行容器

- `--quiet-pull`

  不会打印拉取镜像的进度

- `--no-deps`

  不会启动和有依赖(depens_on属性)的服务

- `--no-build`

  不构建镜像，即使不存在

- `--no-start`

  构建完容器后不会自动启动

- `--build`

  启动容器前构建镜像

- `-force-recreate`

  强制生成镜像和容器，即使没有任何变更。

- `--remove-orphans`

  删除没有在compose file中定义但是由docker-compose产生的容器。

### donw

停止和删除由`doker-compose up` 启动的容器，网络，卷，镜像，服务。

默认只会定义在`docker-compose.yml`中的网络和容器。需要在`docker-compose.yml`所在的文件中运行。

syntax：`docker-compose down [options]`

- `-rmi <type>`

  执行默认动作时，同时删除镜像。type只能是all(所有被服务调用的镜像)，local(删除没有tag的镜像)

  ```
  root in /usr/local/\/composetest λ docker-compose down --rmi all
  Removing network composetest_default
  WARNING: Network composetest_default not found.
  Removing image composetest_web
  Removing image redis:alpine
  
  root in /usr/local/\/composetest λ docker-compose down --rmi local
  Removing network composetest_default
  WARNING: Network composetest_default not found.
  Removing image composetest_web
  WARNING: Image composetest_web not found.
  ```

- `--rmove-orphan`

  删除没有在compose file中定义但是由docker-compose产生的容器

### build

build service or rebuild service。如果Dockerfile改变了，需要使用build指令重新生成镜像。

```
root in /usr/local/\/composetest λ docker-compose build
redis uses an image, skipping
Building with native build. Learn about native build in Compose here: https://docs.docker.com/go/compose-native-build/
Building web
Sending build context to Docker daemon   7.68kB
```

### run

运行指定的service

```
root in /usr/local/\/composetest λ docker-compose run web
Creating network "composetest_default" with the default driver
Creating composetest_web_run ... done
 * Serving Flask app "app.py" (lazy loading)
 * Environment: development
 * Debug mode: on
 * Running on http://0.0.0.0:5000/ (Press CTRL+C to quit)
 * Restarting with stat
 * Debugger is active!
 * Debugger PIN: 657-039-674

root in /usr/local/\/composetest λ cat docker-compose.yml
version: "3.9"
services:
  web:
    build: .
    ports:
      - "5000:5000"
    volumes:
      - .:/code
    environment:
      FLASK_ENV: development
  redis:
    image: "redis:alpine"
root in /usr/local/\/composetest λ docker ps
CONTAINER ID   IMAGE             COMMAND       CREATED              STATUS              PORTS      NAMES
9468fd6eb642   composetest_web   "flask run"   About a minute ago   Up About a minute   5000/tcp   composetest_web_run_37d6ddd7dab9
```



