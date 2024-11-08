# GPG - dirmngr

在 GPG 2.1 之后使用 dirmngr 进程来管理 GPG 和 keyserver 交互，默认会读取 `~/.gnupg/dirmngr.conf` 中的内容，可用的参数和 dirmngr 的 options 一样。常用配置文件如下

```
#指定默认使用的 keyserver
keyserver hkps://keyserver.ubuntu.com
```

只要将 dirmngr 进程 kill 掉，下次 GPG 需要和 keyserver 交互时 ( `--search-keys`, `--send-keys`, `--recv-keys` ) 会自动启动进程并读取配置 

```
root@v2:~/.gnupg# ps -ef | grep -v grep | grep dirmng
root     1070848       1  0 19:48 ?        00:00:00 dirmngr --daemon --homedir /root/.gnupg
root@v2:~/.gnupg# kill -9 1070848
root@v2:~/.gnupg# gpg -v --search-keys hostlockdown@gmail.com
gpg: no running Dirmngr - starting '/usr/bin/dirmngr'
gpg: waiting for the dirmngr to come up ... (5s)
gpg: connection to dirmngr established
gpg: data source: http://162.213.33.8:11371
(1)     cyberpelican <hostlockdown@gmail.com>
          3072 bit RSA key DFCFDF52F1627354, created: 2020-12-10
(2)       3072 bit RSA key 905A3BEBFB96C1D5, created: 2020-12-09
Keys 1-2 of 2 for "hostlockdown@gmail.com".  Enter number(s), N)ext, or Q)uit >
```

