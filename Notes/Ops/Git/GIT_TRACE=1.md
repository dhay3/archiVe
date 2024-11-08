# GIT_TRACE=1

参考：

https://devopshero.wordpress.com/2013/06/19/git-trick-git_trace-1/

https://git-scm.com/book/en/v2/Git-Internals-Environment-Variables

在使用的命令之前使用`GIT_TRACE=1`可以显示详细信息

```
82341@bash MINGW64 /f/test (master)
$ GIT_TRACE=1 git commit -m"2"
22:42:20.955499 exec-cmd.c:237          trace: resolved executable dir: D:/git/Git/mingw64/bin
22:42:20.955499 git.c:439               trace: built-in: git commit -m2
22:42:21.399644 run-command.c:663       trace: run_command: git gc --auto
22:42:21.416497 exec-cmd.c:237          trace: resolved executable dir: D:/git/Git/mingw64/libexec/git-core
22:42:21.416497 git.c:439               trace: built-in: git gc --auto
[master 154c3b1] 2
 1 file changed, 1 insertion(+)
 create mode 100644 2.txt
```

