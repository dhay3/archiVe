ref
[https://git-scm.com/docs/git-log](https://git-scm.com/docs/git-log)
[https://stackoverflow.com/questions/1057564/pretty-git-branch-graphs](https://stackoverflow.com/questions/1057564/pretty-git-branch-graphs)
## Digest
syntax
```
git log [options] [commits]
```
git log 用于查看 commit 记录，==默认值只显示当前分支的 log==
例如
```
root@v2:/home/ubuntu/graphengine/.git# git log 
commit 5c7fe0ef385aa555e55503a44d3b11ba95418e5b (HEAD -> master, origin/master, origin/HEAD)
Merge: 2cbe91c4 2cc4b3eb
Author: yanghaoran <yanghaoran2@huawei.com>
Date:   Thu Mar 23 01:54:51 2023 +0000

    !2154 upgrade ascend software package 22 mar 23
    Merge pull request !2154 from yanghaoran/master

commit 2cc4b3eb183a682f6b0fcec6cf96e20875795d8f
Author: yanghaoran <yanghaoran2@huawei.com>
Date:   Wed Mar 22 16:31:12 2023 +0800

    upgrade ascend software package 22 mar 23

commit 2cbe91c450c0f2472a2bcb1a166873dcef397377
Merge: f6dc0a36 72d99dbc
Author: yanghaoran <yanghaoran2@huawei.com>
Date:   Wed Mar 15 09:35:14 2023 +0000

    !2153 upgrade Ascend software package 15 Mar 23
    Merge pull request !2153 from yanghaoran/master
```
`(HEAD -> master, origin/master, origin/HEAD)` 表示 local HEAD 指针指向 master 分支，同时指向 remote `origin/master, origin/HEAD` 
现在需要看从开始到 `upgrade ascend software package 22 mar 23` 的日志，可以使用对应的 commit hash 值
```
root@v2:/home/ubuntu/test# git log 2cc4b3eb183a682f6b0fcec6cf96e20875795d8f
commit 2cc4b3eb183a682f6b0fcec6cf96e20875795d8f
Author: yanghaoran <yanghaoran2@huawei.com>
Date:   Wed Mar 22 16:31:12 2023 +0800

    upgrade ascend software package 22 mar 23

commit 2cbe91c450c0f2472a2bcb1a166873dcef397377
Merge: f6dc0a36 72d99dbc
Author: yanghaoran <yanghaoran2@huawei.com>
Date:   Wed Mar 15 09:35:14 2023 +0000

    !2153 upgrade Ascend software package 15 Mar 23
    Merge pull request !2153 from yanghaoran/master
```
## Optional args

- `--reflog`

显示所有的引用记录 (==包括 `restore`或者 `reset` 的记录==)

- `--all`

显示 log 时显示所有分支

- `--abbrev-commit`

不显示完整的 commit hash，只显示唯一的最长子串

- `-p | -u | --patch`

查看 commit 日志的同时查看 diff 内容
```
root@v2:/home/ubuntu/gitlab# git log -p
commit 7b7902de3e0e1ca4d944932074b2f5cbd20527b2 (HEAD -> master)
Author: John Doe <qq@com>
Date:   Thu Mar 30 11:51:48 2023 +0800

    test

diff --git a/test b/test
index 493021b..7946e2c 100644
--- a/test
+++ b/test
@@ -1 +1,2 @@
 this is a test file
+3333
```

- `--pretty | --format [format]`

pretty-print the commit logs
format 的值可以是  oneline, short, medium, full, fuller, reference, email, raw 或者自定义的 format:<string>

- `-<number> | -n <number>`

查看从当前指针指向的 commit log 前 number 条日志

- `--skip=<number>`

从当前指针指向的 commit log 开始计算，跳过 number 条

- `--since=<date>`

`--after=<date>`
`--until=<date>`
`--before=<date>`
按照逻辑查看指定日期的 commit log

- `--grep=<patter>`

类似 grep 逻辑，过滤 commit 日志

- `-i | --regexp-ignore-case`

一般和 `--grep` 一起使用，忽略大小写

- `--merges`

只显示 merge 的 commit log

- `--no-merges`

commit log 中不包含 merge 的记录

- `--graph`

在左侧新增一列显示当前分支的 commit log 图表，如果需要显示所有分支的 log 需要和 `--all` 一起使用
```
root@v2:/home/ubuntu/test# git log --graph
* commit a7edc9f013fe8def0b9f575397ca13560447f37c (iss53)
| Author: root <johndoe@exmaple.com>
| Date:   Wed Mar 29 22:00:21 2023 +0800
| 
|     test
|   
*   commit cfd2a15fb92d0a8b06d678d1f4291a65cbebaacc (origin/master, origin/HEAD)
|\  Merge: cda6a6c 6488367
| | Author: cyberPelican <62749885+dhay3@users.noreply.github.com>
| | Date:   Tue Mar 28 21:20:38 2023 +0800
| | 
| |     Merge branch 'iss53'
| | 
| * commit 6488367b66852874ed1f7b15f45eae6a3a6ed94d (origin/iss53)
| | Author: cyberPelican <62749885+dhay3@users.noreply.github.com>
| | Date:   Tue Mar 28 21:00:01 2023 +0800
| | 
| |     iss53
| | 
* | commit cda6a6c43f90968dd884a2a7606c63b88b03f560
|/  Author: cyberPelican <62749885+dhay3@users.noreply.github.com>
|   Date:   Tue Mar 28 21:02:48 2023 +0800
|   
|       hotfix
| 
* commit 56b196e2a4a09ed5d4c7d086c82b304e76ffaf2a (tag: v1.2)
| Author: Scott Chacon <schacon@gmail.com>
| Date:   Mon Mar 17 21:52:11 2008 -0700
```
## format:<string>
format:<string> 被用在 `--format | --pretty` 按照 string 的格式输出 commit 日志。string 可用的值有 
**_%H_**
commit hash
**_%h_**
abbreviated commit hash
**_%T_**
tree hash
**_%t_**
abbreviated tree hash
**_%P_**
parent hashes
**_%p_**
abbreviated parent hashes
**_%an_**
author name
**_%aN_**
author name (respecting .mailmap, see [git-shortlog[1]](https://git-scm.com/docs/git-shortlog) or [git-blame[1]](https://git-scm.com/docs/git-blame))
**_%ae_**
author email
**_%aE_**
author email (respecting .mailmap, see [git-shortlog[1]](https://git-scm.com/docs/git-shortlog) or [git-blame[1]](https://git-scm.com/docs/git-blame))
**_%al_**
author email local-part (the part before the _**@**_ sign)
**_%aL_**
author local-part (see _**%al**_) respecting .mailmap, see [git-shortlog[1]](https://git-scm.com/docs/git-shortlog) or [git-blame[1]](https://git-scm.com/docs/git-blame))
**_%ad_**
author date (format respects --date= option)
**_%aD_**
author date, RFC2822 style
**_%ar_**
author date, relative
**_%at_**
author date, UNIX timestamp
**_%ai_**
author date, ISO 8601-like format
**_%aI_**
author date, strict ISO 8601 format
**_%as_**
author date, short format (YYYY-MM-DD)
**_%ah_**
author date, human style (like the --date=human option of [git-rev-list[1]](https://git-scm.com/docs/git-rev-list))
**_%cn_**
committer name
**_%cN_**
committer name (respecting .mailmap, see [git-shortlog[1]](https://git-scm.com/docs/git-shortlog) or [git-blame[1]](https://git-scm.com/docs/git-blame))
**_%ce_**
committer email
**_%cE_**
committer email (respecting .mailmap, see [git-shortlog[1]](https://git-scm.com/docs/git-shortlog) or [git-blame[1]](https://git-scm.com/docs/git-blame))
**_%cl_**
committer email local-part (the part before the _**@**_ sign)
**_%cL_**
committer local-part (see _**%cl**_) respecting .mailmap, see [git-shortlog[1]](https://git-scm.com/docs/git-shortlog) or [git-blame[1]](https://git-scm.com/docs/git-blame))
**_%cd_**
committer date (format respects --date= option)
**_%cD_**
committer date, RFC2822 style
**_%cr_**
committer date, relative
**_%ct_**
committer date, UNIX timestamp
**_%ci_**
committer date, ISO 8601-like format
**_%cI_**
committer date, strict ISO 8601 format
**_%cs_**
committer date, short format (YYYY-MM-DD)
**_%ch_**
committer date, human style (like the --date=human option of [git-rev-list[1]](https://git-scm.com/docs/git-rev-list))
**_%d_**
ref names, like the --decorate option of [git-log[1]](https://git-scm.com/docs/git-log)
**_%D_**
ref names without the " (", ")" wrapping.

例如
```
root@v2:/home/ubuntu/gitlab# git log --format="format:%h %ae %cD %d"
7b7902d qq@com Thu, 30 Mar 2023 11:52:11 +0800  (HEAD -> master)
a7edc9f johndoe@exmaple.com Wed, 29 Mar 2023 22:00:21 +0800  (iss53)
cfd2a15 62749885+dhay3@users.noreply.github.com Tue, 28 Mar 2023 21:20:38 +0800  (origin/master, origin/HEAD)
cda6a6c 62749885+dhay3@users.noreply.github.com Tue, 28 Mar 2023 21:02:48 +0800 
6488367 62749885+dhay3@users.noreply.github.com Tue, 28 Mar 2023 21:00:01 +0800  (origin/iss53)
56b196e schacon@gmail.com Tue, 28 Mar 2023 17:24:33 +0800  (tag: v1.2)
085bb3b schacon@gmail.com Fri, 17 Apr 2009 21:55:53 -0700 
a11bef0 schacon@gmail.com Sat, 15 Mar 2008 10:31:28 -0700 
```
可以参考 stackoverflow 上的配置
```
[alias]
lg1 = log --graph --abbrev-commit --decorate --format=format:'%C(bold blue)%h%C(reset) - %C(bold green)(%ar)%C(reset) %C(white)%s%C(reset) %C(dim white)- %an%C(reset)%C(auto)%d%C(reset)' --all
lg2 = log --graph --abbrev-commit --decorate --format=format:'%C(bold blue)%h%C(reset) - %C(bold cyan)%aD%C(reset) %C(bold green)(%ar)%C(reset)%C(auto)%d%C(reset)%n''          %C(white)%s%C(reset) %C(dim white)- %an%C(reset)'
lg = lg1
```
