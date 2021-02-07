# Git 设置代理

参考：https://gist.github.com/laispace/666dd7b27e9116faece6

## 设置http，https协议

使用socks5

```
git config --global http.proxy 'socks5://127.0.0.1:1080'
git config --global https.proxy 'socks5://127.0.0.1:1080'
```

针对`github.com`单独配置

```
#只对github.com
git config --global http.https://github.com.proxy socks5://127.0.0.1:1080

#取消代理
git config --global --unset http.https://github.com.proxy
```

## 设置git协议

==git协议使用ssh==

在`~/.ssh/config` 文件后面添加几行，没有可以新建一个。==注意这里不要使用ncat，可能会导致宿主机与虚拟机无法通信==

```
Host github.com
Username git
Hostname github.com
ProxyCommand nc -X 5 -x 127.0.0.1:1080 %h %p #linux

Host github.com
Username git
Hostname github.com
ProxyCommand "D:/git/Git/mingw64/connect.exe" -S 127.0.0.1:10808  %h %p
```

可能出现问题

参考: https://www.idzd.top/archives/2536/

```
$ git push origin master
/usr/bin/bash: line 0: exec: nc: not found
kex_exchange_identification: Connection closed by remote host
fatal: Could not read from remote repository.

Please make sure you have the correct access rights
and the repository exists.
```

解决

```
Host github.com
#-S参数表示使用Socks5代理, 如果是HTTP代理则为-H
ProxyCommand "D:/git/Git/mingw64/connect.exe" -S 127.0.0.1:10808  %h %p
```

## 参考配置文件

```
Host github.com 
User git
Hostname github.com
Port 22
IdentityFile ‪~/.ssh/id_ed25519
ProxyCommand "D:/git/Git/mingw64/connect.exe" -S 127.0.0.1:10808  %h %p
ServerAliveInterval 120
```



