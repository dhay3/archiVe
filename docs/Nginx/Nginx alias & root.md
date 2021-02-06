# Nginx alias & root

> 我们因该使用root替换alias

## alias

```
location /i/ {
    alias /data/w3/images/;
}
```

访问`/i/top.gif`，将会请求`/data/w3/images/top.gif`，==请求不会带上location==

## root

```
location /i/ {
    root /data/w3;
}
```

访问`/i/top.gif`，将会请求`/data/w3/i/top.gif`，==请求带上location==

