# Nginx location

参考：

https://www.cnblogs.com/jpfss/p/10418150.html

https://nginx.org/en/docs/http/ngx_http_core_module.html#location

```
Syntax:	location [ = | ~ | ~* | ^~ ] uri { ... }
location @name { ... }
Default:	—
Context:	server, location
```

location用于匹配request  URI，==可以是字符串，也可以是正则表达式==，采用最佳匹配

```
location = / {
   #规则A
}
location = /login {
   #规则B
}
location ^~ /static/ {
   #规则C
}
location ~ \.(gif|jpg|png|js|css)$ {
   #规则D，注意：是根据括号内的大小写进行匹配。括号内全是小写，只匹配小写
}
location ~* \.png$ {
   #规则E
}
location !~ \.xhtml$ {
   #规则F
}
location !~* \.xhtml$ {
   #规则G
}
location / {
   #规则H
}
```

- `=`：表示精确匹配
- `~*`：表示忽略大小写匹配
- `~`：表示大小写敏感匹配
- `^~`：表示不检查正则表达式（即只匹配字符）

## 特殊案例

**0x001**

如果请求没有匹配到其他的location，则使用该规则

```
location / {
   #规则H
}
```

**0x002**

遵循正则表达式，匹配指定类型的文件

```
location ~* \.(gif|jpg|jpeg)$ {
    #规则H
}
```

