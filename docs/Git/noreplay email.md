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

