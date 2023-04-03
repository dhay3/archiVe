# git switch

ref

https://git-scm.com/docs/git-switch

## Digest

syntax

```
git switch [options] [branch]
```

用于切换到指定 branch，语义上比 checkout 更加明确

和 shell  `cd -` 一样，switch 也可以使用 `-` 指代前一个 branch

```
(base) cpl in /tmp/test on main λ git switch topic 
Switched to branch 'topic'
                                                                                                                                                         
(base) cpl in /tmp/test on topic λ git switch -
Switched to branch 'main'
                                                                                                                                                         
(base) cpl in /tmp/test on main λ 
```

## Optional args

- `-c | --create <new-branch>`

  等价于 `git checkout -b new-branch`