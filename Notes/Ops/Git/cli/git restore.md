# git restore

ref
[https://git-scm.com/docs/git-restore](https://git-scm.com/docs/git-restore)
[https://git-scm.com/docs/git](https://git-scm.com/docs/git)

## Digest
syntax
```
git restore [options] [files_path]
```
和 `git reset` 一样用于回退，但是粒度更细腻用于回退指定文件
## Optional args

- `-s | --source tree`

  指定回退使用那个分支

- `-W | --worktree`

  指定需要 restore 文件的来源是 working direcotry，默认 

- `-S | --staged`

  指定需要 restore 文件的来源是 staging area

## Exmaples
假设不小心删除了 file1，这时我们就可以使用 `git restore` 来恢复 file1
```
root@v2:/home/ubuntu/test# rm -f file1
root@v2:/home/ubuntu/test# git status
On branch master
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        deleted:    file1

no changes added to commit (use "git add" and/or "git commit -a")
root@v2:/home/ubuntu/test# git restore file1
root@v2:/home/ubuntu/test# git status
On branch master
nothing to commit, working tree clean
root@v2:/home/ubuntu/test# ls
file1  file2  file3
```
