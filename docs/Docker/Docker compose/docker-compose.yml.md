# docker-compose.yml

> docker-compose中的属性大部分与Dockerfile中的类似

```
version: "3.9"
   
services:
  db: #服务名
    image: postgres #服务拉取的镜像
    environment: #容器中环境变量
      - POSTGRES_DB=postgres
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
  web:
    build: . #web服务由当前目录中的Dockerfile构建
    command: python manage.py runserver 0.0.0.0:8000 #容器启动后执行的命令
    volumes: #将当前目录与/code挂载
      - .:/code
    ports: #容器的8000映射到宿主机的8000端口
      - "8000:8000"
    depends_on: #依赖，要先启动db然后才能启动web
      - db
```

