# git - References

## Digest

> references 实际不仅仅指向 branches，还包括 tags 还有其他的一些信息

references (也被称为 refs) 是 Git 中重要的一个逻辑概念，类似与指针，指向 branch’s last commit

## Local reference

local reference 是指向对应 local branch’s last commit 的指针，一般以 branch 的名字来命名。例如本地有一个 master branch，那么对应的 local reference 也叫做 master

可以使用 `git branch -v` 来查看，当前所有的 local references

```
(base) cpl in /tmp/gitignore on main λ git branch -v
* main 4488915 Merge pull request #3567 from aaronfranke/godot
```

上述表述 main local references 指向 main branch 4488915 commit (last commit)

### HEAD

HEAD reference 是一个特殊的指针，指向当前所处 branch 对应的 local reference

```
(base) cpl in /tmp/gitignore on main λ git branch -v
* main 4488915 Merge pull request #3567 from aaronfranke/godot
```

使用 `git branch` 并不会显示类似 HEAD 字样，`*` 表示当前所处的 branch ，即 HEAD 指针指向的 local reference

## Remote reference

remote reference 也被称为 remote-tracking branch 是指向 remote branch’s last commit 的==本地指针==，只有在和 remote repository 做 交互时才会移动 (例如 `git fetch`, `git pull`, `git push`)，git 通过它来追踪 remote branches。以 `<remote>/<branch>` 来命名，可以通过 `git branch --remote`  来查看

例如

```
(base) cpl in /tmp/gitignore on main λ git branch --remote
  origin/HEAD -> origin/main
  origin/add-metadata
  origin/annotating-visualstudio
  origin/ghfw
  origin/main
  origin/old-ghfw
  origin/reduce-noise
  origin/rmw-universe
  origin/rmw-universe-2022
```

上述以 origin 开头的都是 remote references

### origin/HEAD

`origin/HEAD` 和 `HEAD` 类似，指向 remote repository 当前所在的 branch 对应的 remote reference

## Examples

### 0x001 one remote repository

例如 你现在通过 `git clone` 克隆了 `git.ourcompany.com` 到本地，那么 git 会自动为 remote repository 创建一个 origin 别名，同时创建一个 master branch，working directory 中内容和 remote repository master branch 中的一致。并创建 remote reference origin/master 指针指向 remote master branch last commit，local reference master 指针指向 local master branch last commit

![Server and local repositories after cloning](https://git-scm.com/book/en/v2/images/remote-branches-1.png)

上半部分是在 remote repository 中的状态，下半部分是在本地通过 `git clone` 克隆到本地的状态

现在针对 master branch 做一些修改并 commit，同时有人往 `git.ourcompany.com` 做了 push master branch 的操作，更新了 remote repository

![Local and remote work can diverge](https://git-scm.com/book/en/v2/images/remote-branches-2.png)

remote reference origin/master 是不会移动的，因为实际是一个本地指针，并不会自动和 remote repository 做同步

如果需要和 remote repository 同步，需要使用 `git fetch <remote>`，在本例子中使用 `git fetch origin`

![`git fetch` updates your remote-tracking branches](https://git-scm.com/book/en/v2/images/remote-branches-3.png)

使用对应命令后，remote reference orgin/master 才会和 remote repository 同步，并移动 orgin/master 指针。当然这时 working direcotory 中的内容并不会改变，需要使用 `git merge` 对 dirgence branch 做合并

### 0x002 multi remote repositories

假设 `git.ourcompany.com` 是预发环境，现在需要添加生产环境 `git.team1.ourcompany.com`

执行如下命令

```
git remote add teamone git://git.team1.ourcompany.com
```

那么状态如下

![Adding another server as a remote](https://git-scm.com/book/en/v2/images/remote-branches-4.png)

可以发现，remote references 是没有移动的，因为 `git remote add` 并不会和 remote repository 做同步，需要使用 `git fetch teamon`

![Remote-tracking branch for `teamone/master`](https://git-scm.com/book/en/v2/images/remote-branches-5.png)

这里就会创建一个 `teamone/master` remote references 指向 teamone master remote branch

**references**

1. https://git-scm.com/book/en/v2/Git-Internals-Git-References
2. https://git-scm.com/book/en/v2/Git-Branching-Remote-Branches

