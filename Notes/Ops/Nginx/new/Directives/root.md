# root

## 0x00 Overview

```
Syntax: 	root path;
Default: 	

root html;

Context: 	http, server, location, if in location
```

指定请求使用的 root directory (直接可以理解成存放静态文件的目录)

path 中也可以包含 variables，但是不能包含 `$document_root` 以及 `$realpath_root`

## 0x01 Example

```
server {
	http{
		location /get/ {
   	 	root /data/w3;
		}
	}
}
```

例如当访问 `localhost`