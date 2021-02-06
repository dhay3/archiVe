# Github SSH keys登入

参考：

https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/connecting-to-github-with-ssh

https://git-scm.com/book/zh/v2/%E6%9C%8D%E5%8A%A1%E5%99%A8%E4%B8%8A%E7%9A%84-Git-%E7%94%9F%E6%88%90-SSH-%E5%85%AC%E9%92%A5

> 注意如果使用ssh登入，需要使用SSH URL。==如果使用2FA必须使用ssh或https登入==

<img src="..\..\imgs\_Git\Snipaste_2020-12-02_11-11-27.png"/>

## 概述

使用ssh keys登入github可以免去账户和密码。一般存储在`~/.ssh`中，分为

公钥和私钥。公钥后缀带有`.pub`

## SSH key

具体参考：

Linux sshd

https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent

> 如果windows无法连接参考
>
> https://stackoverflow.com/questions/52113738/starting-ssh-agent-on-windows-10-fails-unable-to-start-ssh-agent-service-erro
>
> https://murphypei.github.io/blog/2020/07/win10-ssh

1. 生成SSH key，这里最好不适用默认的rsa，应为如果配置了代理

   ```
   ssh-keygen -t ed25519 -C "your_email@example.com"
   ```

   使用`-t`参数指定加密算法，默认使用rsa。`-C`表示注释

2. 确保ssh-agent运行

   ```
   eval "ssh-agent s"
   ```

3. 将密钥添加到ssh-agent中，==否则不会成功==

   ```
   ssh-add ~/.ssh/id_dsa
   ```

4. 添加ssh key 到github

   复制`id_rsa.pub`生成的公钥，将其复制到setting → SSH and GPG keys。

   <img src="..\..\imgs\_Git\Snipaste_2020-09-25_09-54-21.png"/>

5. 测试 SSH 连接，T表示不分配终端

   ```
   $ ssh -T git@github.com
   ```

   <img src="..\..\imgs\_Git\Snipaste_2020-09-25_09-56-48.png"/>