## Backend

```
FROM harbor.trjcn.local/library/centos7_zh/jdk8:4.0
LABEL version=1.0
LABEL name=whdxh_backend
WORKDIR /java-app
ENV TIME_ZONE=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone
COPY platform-web-1.0-SNAPSHOT.jar /java-app/platform-web-1.0-SNAPSHOT.jar
COPY run.sh /java-app/run.sh
EXPOSE 8090/TCP
STOPSIGNAL SIGINT
ENTRYPOINT ["sh", "-c", "/java-app/run.sh"]


#!/bin/bash
docker build . --no-cache -t whdxh_backend:uat
   

java -Xms1024m -Xmx1024m -Xss512k -jar -Dserver.port=8090 -Dspring.profiles.active=uat platform-web-1.0-SNAPSHOT.jar 

```

```yaml
whdxh_backend:
    image: whdxh_backend:uat
    container_name: whdxh_backend
    restart: on-failure
    ports: 8090:8090
    volumes:
	    - ./backend/config:/java-app/config
	    - ./backend/logs:/java-app/logs
    #depends_on:
    #  - zucc_redis
     # - mysql
    #networks:
     # - rxapp

```