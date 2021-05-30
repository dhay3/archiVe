# Git LFS

参考：

https://git-lfs.github.com/

https://blog.csdn.net/Tyro_java/article/details/53440666

## 概述

GIT large files storage 用于解决git上传单个大文件（100MB的限制）到remote resposirory。详细可见

https://docs.github.com/en/free-pro-team@latest/github/managing-large-files/working-with-large-files

## 操作

> 如果出现
>
> `Remote "origin" does not support the LFS locking API. Consider disabling it with:  `
>
> 参考：https://github.com/git-lfs/git-lfs/issues/3400#
>
> 1. 尝试升级GIT
> 2. `git config lfs.locksverify false`
>
> 如果出现
>
> `batch response: Post "https://lfs.github.com/dhay3/archive/objects/batch": proxyconnect tcp: dial tcp: lookup socksh: no such host`
>
> 参考：
>
> https://github.com/git-lfs/git-lfs/issues/1424
>
> https://stackoverflow.com/questions/55067898/git-lfs-pull-in-repository-produces-an-error-about-dial-tcp
>
> ```
> git config --global --unset http.proxy
> git config --global --unset https.proxy
> ```

```
82341@bash MINGW64 /d/asset/note (master)
$ git lfs install
Updated git hooks.
Git LFS initialized.

82341@bash MINGW64 /d/asset/note (master)
$ git lfs track "*.pdf"
Tracking "k"
Tracking "*.pdf"

```

