# Git 设置代理

参考：

https://gist.github.com/laispace/666dd7b27e9116faece6

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
ProxyCommand "D:/git/Git/mingw64/bin/connect.exe" -S 127.0.0.1:10808  %h %p
```

可能出现问题

参考:

 https://www.idzd.top/archives/2536/

https://zhuanlan.zhihu.com/p/126117538

## windows 自动启动ssh-agent

==git中ssh-agent无法自动启动，使用ssh前需要执行如下命令==

```
eval `ssh-agent -s`
ssh-add ~/.ssh/srectkey
```

可以通过如下方式配置ssh-agent自动启动

https://docs.github.com/en/github/authenticating-to-github/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent

## 参考配置文件

```
Host *
	TCPKeepAlive yes
	ServerAliveInterval 120
	
Host github.com 
	User git
	Hostname github.com
	Port 22
	IdentityFile "‪C:/Users/82341/.ssh/id_ed25519"
	#-S参数表示使用Socks5代理, 如果是HTTP代理则为-H
	ProxyCommand "D:/git/Git/mingw64/bin/connect.exe" -S 127.0.0.1:1080  %h %p

```



