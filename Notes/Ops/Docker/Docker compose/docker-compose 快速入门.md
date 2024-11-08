# Docker compose

参考：

https://docs.docker.com/compose/

docker compse是一个定义和启动容器的工具(便于微服务的管理)，通过yaml配置。

使用docker compse需要以下三个步骤：

1. 通过Dockerfile定义app需要的环境
2. 通过docker-compose.yml定义app需要的service
3. 使用`docker-compose up`命令启动所有服务(默认读取当前文件夹中的docker-compose.yml文件，也可以通过`-f`参数指定文件)

docker compose 有继承机制，基础文件叫`docker-compose.yml`。`docker-compose.override.yml`会继承`docker-compose.yml`中内容并对里面的内容进行override，使用`docker-compose up`会读取`docker-compose.override.yml`中的内容。

都可以通过`-f`参数指定，会读取两个配置文件然后组合。

```
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```

## 入门

安装与校验参考：

https://docs.docker.com/compose/install/

==注意这里脚本会自动安装在`/urs/local/bin/`下，需要赋予excute权限==

https://docs.docker.com/compose/gettingstarted/

因为这里使用了官方的镜像，所以需要==替换Dockerfile中构建镜像的下载源==。从Docker hub上查找到镜像中使用的layer是alpine linux，我这里使用aliyun的alpine镜像。

```
root in /usr/local/\/composetest λ cat Dockerfile
FROM python:3.7-alpine
WORKDIR /code
ENV FLASK_APP=app.py
ENV FLASK_RUN_HOST=0.0.0.0
#添加aliyun镜像
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories
RUN apk add --no-cache gcc musl-dev linux-headers
COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt
EXPOSE 5000
COPY . .
CMD ["flask", "run"]
```

校验

```
root in /usr/local/\/composetest λ curl -fsSL localhost:5000
Hello World! I have been seen 1 times.
```

## 环境变量

> 变量的使用顺序
>
> 1. Compose file
> 2. Shell environment variables
> 3. Environment file
> 4. Dockerfile
> 5. Variable is not defined

- var for docker-compose.yml

  docker-compose.yml中可以使用当前shell的环境变量

  ```
  web:
    image: "webapp:${TAG}" #变量通过shell变量的形式获取
  ```

  也可以将变量存在当前目录中的`.env`文件中，或通过`--env-file`参数来指定文件。

  ```
  docker-compose --env-file ./config/.env.dev up 
  ```

- var for container

  可以通过environment属性来设置==容器中使用的变量==。和`docker run -e KEY=VALUE`一样。

  ```
  web:
    environment:
      - DEBUG=1
  ```

  也可以通过env_file属性将环境变量以文件的形式传给==容器==。和`docker run --env-file=File`一样。

  ```
  web:
    env_file:
      - web-variables.env
  ```

- var for docker-compose run

  使用`-e`参数指定docker-compose run是的环境变量

  ```
  er-compose run -e DEBUG=1 web python console.py
  ```

## profile

和springboot一样，docker也可以通过profile在特定环境下开启服务。

```
version: "3.9"
#services会从 docker hub 上下载镜像 
services:
  frontend:
    image: frontend
    profiles: ["frontend"]

  phpmyadmin:
    image: phpmyadmin
    depends_on:
      - db
    profiles:
      - debug

  backend:
    image: backend

  db:
    image: mysql
```

fontend 只会在frontend profile激活的情况下运行，phpmyadmin只会在debug profile激活的情况下运行。没有指定profile的service always be enabled

通过`docker-compose --profile <profile-name> up`来指定激活的profile，通过多个`--profile`来指定启动环境。或是设置COMPOSE_PROFILES环境变量。

```
$ docker-compose --profile frontend --profile debug up
$ COMPOSE_PROFILES=frontend,debug docker-compose up
```

## Network

`docker-compose up`运行时

1. 会生成一个以项目名命名的Network

   ```
   root in /usr/local/\/composetest λ docker network ls
   NETWORK ID     NAME                  DRIVER    SCOPE
   80b9b205cd4c   bridge                bridge    local
   1c52adfeb5c5   composetest_default   bridge    local
   6b19b36f4b9c   help                  bridge    local
   81078d4d17f9   host                  host      local
   6cc6d934173d   none                  null      local
   ```

2. container分别加入到这个Network

   ```
   root in /usr/local/\/composetest λ docker-compose ps
          Name                      Command               State           Ports
   -------------------------------------------------------------------------------------
   composetest_redis_1   docker-entrypoint.sh redis ...   Up      6379/tcp
   composetest_web_1     flask run                        Up      0.0.0.0:5000->5000/tcp
   root in /usr/local/\/composetest λ docker ps
   CONTAINER ID   IMAGE             COMMAND                  CREATED       STATUS       PORTS                    NAMES
   c87eb8f691ab   composetest_web   "flask run"              2 hours ago   Up 2 hours   0.0.0.0:5000->5000/tcp   composetest_web_1
   2ff25c2cfd4d   redis:alpine      "docker-entrypoint.s…"   3 hours ago   Up 2 hours   6379/tcp                 composetest_redis_1
   
   root in /usr/local/\/composetest λ docker network inspect -f '{{json .Containers}}' composetest_default
   {"2ff25c2cfd4de351c7e19de0ae99118ece313401fa1967a4aee1b21edfeaa040":{"Name":"composetest_redis_1","EndpointID":"6834629e2c2d28737654ecfbbac8b278ba730e3c2eff9f7190f6c0593d1069f4","MacAddress":"02:42:ac:17:00:02","IPv4Address":"172.23.0.2/16","IPv6Address":""},"c87eb8f691ab59c9c0bd9e3d2c65baa4ba198f1a820175ab23a5baafefeddf46":{"Name":"composetest_web_1","EndpointID":"de208ed9f1779482094e311348a165c5ae239378a9ecf5144a7c4740508b49fc","MacAddress":"02:42:ac:17:00:03","IPv4Address":"172.23.0.3/16","IPv6Address":""}}
   ```

   























