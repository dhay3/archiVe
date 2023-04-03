# noreplay email

参考：

https://docs.github.com/en/github/setting-up-and-managing-your-github-user-account/managing-email-preferences/setting-your-commit-email-address

> 1. for web-based git operations, set your commit email address on github
> 2. for commits you push from the cli, set your commit email address in git

使用git commit时，会将git设置的邮箱与commits关联。为了保护邮箱的安全可以使用github提供的==noreplay email==(`ID+username@users.noreply.github.com`)。可以通过如下链接获取ID

`https://api.github.com/users/<username>`

为了在commits中使用noreply email，需要在git中设置关联邮箱为noreply email。

```
git config --global 642885+oray@users.noreply.github.com
```

### block cli pushes that expose my email

当选中该选项时，github会校验git commit关联的email。如果选中了keep my email addresses private，并且git关联的email是github上设置的email时，github就会拒绝commit

## gpg

参考：

https://docs.github.com/en/github/authenticating-to-github/managing-commit-signature-verification/generating-a-new-gpg-key

还可以使用no replay email做gpg签名

```
cpl in ~ λ cat .gitconfig 
[user]
        name = cyberPelican
        email = 62749885+dhay3@users.noreply.github.com
        signingKey = F3A82ABD5E016AC9
[filter "lfs"]
        smudge = git-lfs smudge -- %f
        process = git-lfs filter-process
        required = true
        clean = git-lfs clean -- %f
[http]
        postBuffer = 157286400
        proxy = socks5://127.0.0.1:1089
[lfs]
        locksverify = false
[commit]
        gpgSign = no
[core]
        editor = vim
[alias]
        ci = commit -S -m 
[advice]
        addIgnoreFile = true
```

