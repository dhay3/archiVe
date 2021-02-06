# SSH ssh-copy-id

该命令将客户端的==当前用户的密钥(不同用户登入会失败==)拷贝到服务器的==指定用户==的`~/.ssh/authorized_keys`中，如果文件不存在，就会新建文件并写入。

使用`-i`指定公钥位置，默认为`~/.ssh/id*.pub`，如果文件名不带`.pub`，`ssh-copy-id`会自动带上

```
┌─────( root)─────(~/.ssh) 
 └> $ ssh-copy-id root@192.168.80.143
/usr/bin/ssh-copy-id: INFO: Source of key(s) to be installed: "/root/.ssh/id_rsa.pub"
/usr/bin/ssh-copy-id: INFO: attempting to log in with the new key(s), to filter out any that are already installed
/usr/bin/ssh-copy-id: INFO: 1 key(s) remain to be installed -- if you are prompted now it is to install the new keys
Host key fingerprint is SHA256:S6TRfoi/8wkrM74w95gjaTnPZApIKB3W2xJ9Pbghlyk
+---[ECDSA 256]---+
|                 |
|   . . . =       |
|  o o E X o      |
|.o . + @ + .     |
|o.. o + S .      |
|o.   . o o       |
|. . oo+ +        |
|   .*O=+.+ .     |
|   ..=B*+oo      |
+----[SHA256]-----+
root@192.168.80.143's password: 

Number of key(s) added: 1

Now try logging into the machine, with:   "ssh 'root@192.168.80.143'"
and check to make sure that only the key(s) you wanted were added.

```

使用`ssh username@hostname`来校验

```
 ┌─────( root)─────(~/.ssh) 
 └> $ ssh root@192.168.80.143
Host key fingerprint is SHA256:S6TRfoi/8wkrM74w95gjaTnPZApIKB3W2xJ9Pbghlyk
+---[ECDSA 256]---+
|                 |
|   . . . =       |
|  o o E X o      |
|.o . + @ + .     |
|o.. o + S .      |
|o.   . o o       |
|. . oo+ +        |
|   .*O=+.+ .     |
|   ..=B*+oo      |
+----[SHA256]-----+
Last login: Sat Dec 19 11:43:20 2020 from 192.168.80.200

 __             _                   
/   \/|_  _  __|_) _  |  o  _  _ __ 
\__ / |_)(/_ | |  (/_ |  | (_ (_|| |


Sat Dec 19 11:46:04 CST 2020
```

