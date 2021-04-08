# SSH 密钥校验

> ssh 默认有登入校验的优先级，推荐使用publickey authentication，这里也使用publickey authentication来做为登入校验。（只有所有认证都失效了才会采用密码校验）
>
> 使用公钥认证无需密码
>
> ```
> gssapi-with-mic,hostbased,publickey,
> keyboard-interactive,password
> ```
>
> ==针对当前上传了公钥的用户，和sshd上指定的用户，否则就会造成公钥认证失败==

## 方法一：手动上传密钥

> 也可以手动将公钥赋值到sshd服务器上

1. 生成密钥`ssh-keygen`

2. `PS C:\Users\82341\.ssh> cat .\id_rsa.pub |  ssh root@192.168.80.143 "cd ~/.ssh  && cat >> ~/.ssh/authorized_keys "`

3. 设置文件权限`chmod 644 authorized_keys`，==注意这里的权限必须是必须限制其他用户的写权限==，否则服务器会拒绝读取该文件

4. 登入校验，这里没有设置私钥的passpharse，如果设置了passpharse需要passpharse才能登入

   ```
   PS C:\Users\82341\.ssh> ssh root@192.168.80.143
   Last login: Sat Dec 19 11:20:49 2020 from 192.168.80.1
    __             _
   /   \/|_  _  __|_) _  |  o  _  _ __
   \__ / |_)(/_ | |  (/_ |  | (_ (_|| |
   
   Sat Dec 19 11:20:57 CST 2020
   ```

## 方法而：自动上传密钥

参考：[ssh-copy-id](./SSH ssh-copy-id)