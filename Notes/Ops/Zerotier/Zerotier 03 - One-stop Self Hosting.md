---
createTime: 2025-10-10 00:02
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# Zerotier 03 - One-stop Self Hosting

## 0x01 Preface

```
services:
  ztncui:
    image: keynetworks/ztncui
    container_name: ztncui
    restart: unless-stopped
    environment:
      - NODE_ENV=production
      - HTTPS_PORT=3443
      - HTTP_PORT=3000
      - HTTP_ALL_INTERFACES=1
      - MYADDR=8.153.99.116
      - MYDOMAIN=zero.cyberpelican.icu
      - ZTNCUI_PASSWD=zerotier@admin
    volumes:
      - ./zerotier-one:/var/lib/zerotier-one
      - ./ztncui:/opt/key-networks/ztncui/etc
    ports:
      - 9993:9993/udp
      - 3443:3443
      - 3000:3000
    networks:
      - zerotier

networks:
  zerotier:
```

TLS acme.sh nginx 外挂

```
services:
  zerotier:
    image: zerotier/zerotier:latest
    container_name: zerotier
    restart: unless-stopped
    volumes:
      - ./zerotier-one:/var/lib/zerotier-one
    environment:
      - ZT_OVERRIDE_LOCAL_CONF=true
      - ZT_ALLOW_MANAGEMENT_FROM=0.0.0.0/0
    expose:
      - "9993/tcp"
    ports:
      - "9993:9993/udp"
  zerotier-web:
    image: dec0dos/zero-ui:latest
    container_name: zerotier-web
    build:
      context: .
      dockerfile: ./docker/zero-ui/Dockerfile
    restart: unless-stopped
    depends_on:
      - zerotier
    volumes:
      - ./zerotier-one:/var/lib/zerotier-one
      - ./data:/app/backend/data
    environment:
      - ZU_CONTROLLER_ENDPOINT=http://zerotier:9993/
      - ZU_SECURE_HEADERS=true
      - ZU_DEFAULT_USERNAME=admin
      - ZU_DEFAULT_PASSWORD=zerotier@admin
    expose:
      - "4000"
    ports:
      - "4000:4000"
```


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***



***References***


