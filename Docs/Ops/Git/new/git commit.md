ref
[https://git-scm.com/docs/git-commit](https://git-scm.com/docs/git-commit)
[https://git-scm.com/book/en/v2/Git-Basics-Undoing-Things](https://git-scm.com/book/en/v2/Git-Basics-Undoing-Things)
## Digest
syntax
```
git commit [options]
```
将 staging area 的内容提交到 repository
## Optional args

- `-m | --message <msg>`

指定需要使用的 commit message

- `--amend`

将提交到 repository 中的内容回撤，然后已当前 staging area 的内容提交
```
root@v2:/home/ubuntu/test# git status
On branch master

No commits yet

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        a
        b

nothing added to commit but untracked files present (use "git add" to track)
root@v2:/home/ubuntu/test# git add .
root@v2:/home/ubuntu/test# git commit -m "1"
[master (root-commit) a7af196] 1
 2 files changed, 2 insertions(+)
 create mode 100644 a
 create mode 100644 b
root@v2:/home/ubuntu/test# echo 333 >> a
root@v2:/home/ubuntu/test# git add .
root@v2:/home/ubuntu/test# git commit --amend -m "2"
[master eb0ec74] 2
 Date: Thu Mar 30 15:57:59 2023 +0800
 2 files changed, 3 insertions(+)
 create mode 100644 a
 create mode 100644 b
```
上述例子回退了 message 为 1 的 commit，而使用了 message 为 2 的 commit

- `--dry-run`

不实际提交文件到 repository，只测试

- `-S | --gpg-sign [keyid]`

使用 GPG key 签名 commit
## Master branch before commit
```
root@v2:/home/ubuntu/test# ls
file1  file2  file3
root@v2:/home/ubuntu/test# git init
Initialized empty Git repository in /home/ubuntu/test/.git/
root@v2:/home/ubuntu/test# git br -v
root@v2:/home/ubuntu/test# 
```
当使用 `git init`时并不会直接生成 master branch，需要 commit 后才会出现
```
root@v2:/home/ubuntu/test# git add .
root@v2:/home/ubuntu/test# git commit -m "init"
[master (root-commit) 1aa0bad] init
 3 files changed, 3 insertions(+)
 create mode 100644 file1
 create mode 100644 file2
 create mode 100644 file3
root@v2:/home/ubuntu/test# git br -v
* master 1aa0bad init
```
