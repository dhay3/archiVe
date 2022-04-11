# Github SSH keys登入

参考：

https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/connecting-to-github-with-ssh

https://git-scm.com/book/zh/v2/%E6%9C%8D%E5%8A%A1%E5%99%A8%E4%B8%8A%E7%9A%84-Git-%E7%94%9F%E6%88%90-SSH-%E5%85%AC%E9%92%A5

> 注意如果使用ssh登入，需要使用SSH URL。==如果使用2FA必须使用ssh或personal access token登入==

![Snipaste_2020-12-02_11-11-27](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2020-12-02_11-11-27.i9gkbxs3uvs.png)

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
   eval "ssh-agent"
   ```

3. 将密钥添加到ssh-agent中，==否则不会成功==

   ```
   ssh-add ~/.ssh/id_dsa
   ```

4. 添加ssh key 到github

   复制`id_rsa.pub`生成的公钥，将其复制到setting → SSH and GPG keys。

   ![Snipaste_2020-09-25_09-54-21](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-09-25_09-54-21.5hf1rxssqi40.png)

5. 测试 SSH 连接，T表示不分配终端

   ```
   $ ssh -T git@github.com
   ```

   ![Snipaste_2020-09-25_09-56-48](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-09-25_09-56-48.64pv3ziw4b80.png)

## ssh-agent

https://docs.github.com/en/github/authenticating-to-github/working-with-ssh-key-passphrases#auto-launching-ssh-agent-on-git-for-windows

在windows git上ssh-agent不会自动启动，所以每次使用git时都需要手动输入。我们可以使用脚本让它自动启动。在`~/.profile`或者`~/.bashrc`中添加如下代码段

```
env=~/.ssh/agent.env
#如果之前存在文件,先将之前的文件执行
agent_load_env(){
  test -f "$env" && source "$env" >| /dev/null
}
#文件600,目录700
agent_start(){
  umask 077
  ssh-agent >| "$env"
  source "$env"
}
agent_load_env

#0:agent running has keys, 1:agent running has 0 key, 2:agent is not running
ssh_run_state=$(ssh-add -l >| /dev/null 2>&1;echo $?)

if [ ! "$SSH_AUTH_SOCK" ] || [ $ssh_run_state = 2 ]; then
    agent_start
    #通过source将变量赋值,而不是通过eval
    ssh-add
elif [  "$SSH_AUTH_SOCK" ] || [ $ssh_run_state = 1 ]; then
  ssh-add
fi
#取消变量
unset env

```







