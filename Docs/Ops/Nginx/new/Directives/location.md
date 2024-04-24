# location

## 0x00 Overview

```
Syntax: 	location [ = | ~ | ~* | ^~ ] uri { ... }
location @name { ... }
Default: 	—
Context: 	server, location
```

location 指令用于匹配 URI (如果 URI 做了 encoding，例如 URI 中有中文, Nginx 会自动 decoding 然后再做 URI 匹配) 执行对应代码块中的内容

==注意 Nginx 中的 URI 和传统意义上的 URI 不同(这点可以从 variable `$uri` 中得出)==

```

```

这里的 Context 指明 location 同样也能出现在 location 指令块中，例如

```
location ^~ /images/ {
    ...
    location ~ /*.jpg/ {
    	...
    }
    location ~ /*.png/ {
    	...
    }
}
```

## 0x01 [= | ~ | ~* | ^~]

可以将这些符号统一称为 modifiers，可以分为两类

1. prefix modifiers 
2. regex modifiers

- `=`

  exact matching, the search terminate immediately

  uri 必须精确匹配才会执行，匹配后不会匹配其他的 locatoin

  属于 prefix modifiers

  ```
  location = /pics/ {
      [ configuration ]
  }
  ```

- `^~`

  uri 匹配当前规则后，不会再去匹配 regex modifiers 的规则

  属于 prefix modifiers

  ```
  location ^~ /images/ {
      [ configuration ]
  }
  ```

- none

  即没有符号

  属于 prefix modifiers

  ```
  location /documents/ {
      [ configuration ]
  }
  ```

- `~*`

  case-insensitive matching

  uri 使用正则匹配，忽略大小写

  属于 regex modifiers

  ```
  location ~* /MAP/ {
      [ configuration ]
  }
  ```

- `~`

  case-sensitive matching

  uri 使用正则匹配，大小写敏感

  属于 regex modifiers

  ```
  location ~ /Map/ {
      [ configuration ]
  }
  ```

  针对大小写不敏感的系统(MacOS/Windows)等价与 `~*`

- `@`

## 0x02 Matching rules

> *To find location matching a given request, nginx first checks locations defined using the prefix strings (prefix locations). Among them, the location with the longest matching prefix is selected and remembered. Then regular expressions are checked, in the order of their appearance in the configuration file. The search of regular expressions terminates on the first match, and the corresponding configuration is used. If no match with a regular expression is found then the configuration of the prefix location remembered earlier is used.*

简单的说就是

1. Nginx 首先会使用 URI 去匹配 logest prefix modifiers (这点和路由的逻辑很像)
2. 然后使用 URI 去匹配 regular modifiers

伪代码逻辑如下

```
def match()
	read rules from top to bottom

def config()
	while rules && match(= uri) then
		apply(= config)
		return
  shift rule
  while rules && match(^~ uri) then
      apply(^= config)
      return
    shift rule
  while rules && (match(~ uri) or match(~* uri)) then
      apply(~ uri) or apply(~*)
      return
    shift rule
  while rules && match(uri)
    if match(uri) then
      apply(uri)
      return
    shift rule
```

1. 先查看是否有 `locatoin = uri` 的，如果有匹配，则使用该配置
2. 如果没有，查看 `location ^~ uri` 的，如果有匹配，则使用该配置
3. 如果没有，查看 `location ~ uri` 和 `location ~* uri` 的，如果有匹配，则使用该配置
4. 如果没有，查看 `location uri` 的,如果有匹配，则使用该配置，如果没有就报错

可以归纳得出 modifiers 优先级如下

1. `=`
2. `^~`
3. `~` `~*`
4. none

## 0x03 @

Nginx

## 0x03 Exmaples



**references**

[^1]:http://nginx.org/en/docs/http/ngx_http_core_module.html#location
[^2]:https://www.digitalocean.com/community/tutorials/understanding-nginx-server-and-location-block-selection-algorithms
[^3]:https://serverfault.com/questions/674425/what-does-location-mean-in-an-nginx-location-block
[^4]:https://serverfault.com/questions/738452/what-does-the-at-sign-mean-in-nginx-location-blocks



