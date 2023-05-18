# git - mergetool

## Digest

syntax

```
git mergetool [--tool=<tool>]
```

使用 mergetool 来对比并修改冲突，一般在 `git merge` 后使用

## Optional args

- `-t | --tool=<tool>`

  指定使用的 mergetool，常用的有 `vimdiff`，`kdiff3`

  如果没有指定，默认使用配置文件的 `merge.tool`

## kdiff3

如果需要使用 kdiff3 作为 mergetool, 使用如下配置

```
git config --global diff.tool kdiff3
git config --global difftool.prompt false
git config --global difftool.keepBackup false
git config --global difftool.trustExitCode false
git config --global merge.tool kdiff3
git config --global mergetool.prompt false
git config --global mergetool.keepBackup false
git config --global mergetool.keepTemporaries false
```

需要注意的一点是，如果 kdiff3 没有正常退出，同样会生成 backup 以及 temporary files



**references**

1. https://docs.kde.org/trunk5/en/kdiff3/kdiff3/git.html
2. https://stackoverflow.com/questions/161813/how-do-i-resolve-merge-conflicts-in-a-git-repository