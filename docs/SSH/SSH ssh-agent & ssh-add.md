# SSH ssh-agent & ssh-add

> 公钥认证时，需要输入设置私钥时的passpharse。使用ssh-agent就不用每次在使用私钥时输入passpharse
>
> windows中ssh-agent对应服务OpenSSH Authentication Agent，不会自动启动。linux一般设置为自动启动

## ssh-agent

密钥管理器，然后进行公钥认证。

`-s`缺省参数，默认会议Bourne shell启动

```
root in ~ λ ssh-agent
SSH_AUTH_SOCK=/tmp/ssh-YXuiYfOL71Pm/agent.8172; export SSH_AUTH_SOCK;
SSH_AGENT_PID=8173; export SSH_AGENT_PID;
echo Agent pid 8173;
root in ~ λ eval $(ssh-agent)
Agent pid 8178
```

## ssh-add

将私钥加入到ssh-agent，用于ssh认证。==如果没有参数，默认将所有的密钥加入到ssh-agent==

```
82341@bash MINGW64 /d/asset/note/docs/SSH (master)
$ ssh-add
Identity added: /c/Users/82341/.ssh/id_rsa (82341@bash)
Identity added: /c/Users/82341/.ssh/id_dsa (82341@bash)
Identity added: /c/Users/82341/.ssh/id_ed25519 (82341@bash)
```

如果添加的不是默认生成的私钥，`ssh-add`命令需要显示指定私钥

```
$ ssh-add my-other-key-file
```

**参数**

- `-D`

  删除加入到ssh-agent中的所有私钥，不会对本地的私钥产生影响

- `-l`

  展示当前agent中的私钥指纹

  ```
  82341@bash MINGW64 /d/asset/note/docs/SSH (master)
  $ ssh-add -l
  3072 SHA256:SU95FYIyOdtG0zAkfAFZnFFPxJ4TxpvyakTlw5u9lWo 82341@bash (RSA)
  1024 SHA256:pGpEWiCy9gIhQ4LlIvO6K1oplMU5EyMf9SBgSz3Kuuw 82341@bash (DSA)
  256 SHA256:3Pk8/tJFx6nnplbfffsHJhl7dmb7ExOMhHxGYEvX73A 82341@bash (ED25519)
  ```

- `-x`

  对agent加密

- `-X`

  对agent解密



