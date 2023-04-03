# git pull

ref

https://git-scm.com/docs/git-pull

## Digest

syntax

```
git pull [options] [repository] [remote_branch...]
```

“Incorporates changes from a remote repository into the current branch”

`git pull` 是用于将 remote repository 中的内容合并到 current branch

如果当前分支落后于 remote 分支，且内容不冲突就会使用 Fasst-Forward 合并分支

简单的说就是 `git pull`  包含两部分 `git fetch` 和 `git merge` 或者 `git rebase`

假设当前的状态如下

```
	  A---B---C master on origin
	 /
    D---E---F---G master
	^
	origin/master in your repository
```

如果这时候使用 `git pull`，就需要将 master on origin 和 master 合并

```
	  A---B---C origin/master
	 /         	     \
    D---E---F---G---H master
```

## Optional args

具体参数可以参考 `git fetch` 和 `git merge` 或 `git rebase`

