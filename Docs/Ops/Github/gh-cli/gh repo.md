ref
[https://cli.github.com/manual/gh_repo](https://cli.github.com/manual/gh_repo)
管理 repository 的一系列子命令
## Query Subcommands
### List
syntax
```
gh repo list [<owner>] [flags]
```
用于查询 repository，可以通过指定 owner 来查看指定的 owner 的 repository
```
λ / gh repo list torvalds
```
#### Optional args

- `--fork`

只显示 fork 的 repository

- `-l | --language <string>`

只显示使用特定 language 写的 repository

- `--source`

显示不是 fork 的 repository
### View
syntax
```
gh repo view [<repository>] [flags]
```
用于查看 repository  README 信息，默认查看当前目录对应的 repository，也可以手动指定 repository
```
λ ~/ gh repo view torvalds/linux
```
#### Optional args

- `-b | --branch <string>`

查看指定分支

- `-w | --web`

使用 browser 打开
## Update Subcommands
### Create
syntax
```
gh repo create [flags] [<name>] 
```
用于创建 repository
```
λ ~/ gh repo create helloworld --public 
✓ Created repository dhay3/helloworld on GitHub
```
#### Optional args

- `--private | public`

创建的 repository 是否是公开或者私有的

- `--add-readme`

创建 repository 时自动创建 README

- `-c | clone`

克隆 repository 等价于 `git clone`

- `--disable-issues`

新创建的 repository 不包含 issues 模块

- `--disable-wiki`

新创建的 repository 不包含 wiki 模块
### Edit
syntax
```
gh repo edit [<repository>] [flags]
```
修改当前或者指定 repository 的设置，`--enable-xxx` 的参数都可以通过 `--enable-xxx=false` 方式来取反
#### Optional args

- `--add-topic <string>`

为 repository 添加话题

- `--allow-forking`

允许 fork

- `--default-branch <name>`

设置 repository 初始的 branch name

- `--delete-branch-on-merge`

pull request merge 后删除分支

- `--enable-discussions`

开启 discussions

- `--enable-issues`

开启 issues

- `--enable-wiki`

开启 wiki

- `--visibility <string>`

修改仓库的属性到 public/private/internal
### Delete
syntax
```
gh repo delete [<repository>] [flags]
```
删除指定的 repository，如果没有指定 repository 默认删除当前的所处的 repository
```
λ ~/ gh repo delete helloworld --yes
```
#### Optional args

- `--yes`

不提示 prompt 直接删除
### Rename
syntax
```
gh repo rename [<new-name>] [flags]
```
用于重命名当前的 repository
#### Optional args

- `-R | --repo <[HOST/]OWNER/REPO>`

使用指定的 repository 而不是默认的当前 repository
```
λ / gh repo rename -R dhay3/gh-test ghlab
✓ Renamed repository dhay3/ghlab
```
### Archive/Unarchive
syntax
```
gh repo archive [<repository>] [flags]
gh repo unarchive [<repository>] [flags]
```
将指定或者默认当前 repository 存档 (archive, read-only and indicate that it's no longer actively maintained )
```
λ ~/ gh repo archive helloworld
? Archive dhay3/helloworld? Yes
✓ Archived repository dhay3/helloworld
```
可以使用 `unarchive` 取消存档
```
λ ~/ gh repo unarchive dhay3/helloworld
? Unarchive dhay3/helloworld? Yes
✓ Unarchived repository dhay3/helloworld
```
#### Optional args

- `--yes`

skip the confirmation prompt
### Clone
syntax
```
gh repo clone <repository> [<directory>] [-- <gitflags>...]
```
用于克隆 repository，如果需要指定额外的 `git clone` 的参数，需要在 `--` 之后
```
λ ~/ gh repo clone dhay3/gitlab lab -- --depth=1
Cloning into 'lab'...
remote: Enumerating objects: 6, done.
remote: Counting objects: 100% (6/6), done.
remote: Compressing objects: 100% (5/5), done.
remote: Total 6 (delta 0), reused 2 (delta 0), pack-reused 0
Unpacking objects: 100% (6/6), done.
```
#### Optional args

- `-u | --upstream-remote-name <string>`

指定 remote repository 的名字
### Fork
syntax
```
gh repo fork [<repository>] [-- <gitflags>...] [flags]
```
用于 fork a repository 如果没有指定 repository，默认使用当前目录对应的 repository
```
λ ~/ gh repo fork cli/cli
✓ Created fork dhay3/cli
? Would you like to clone the fork? No
```
#### Optional args

- `--fork-name <string>`

rename the forked repository

- `--clone`

fork 完成后直接 clone

- `--default-branch-only`

只克隆主分支，即 master 或者 main
#### Examples
```
# Skipping clone prompts using flags
~/Projects$ gh repo fork cli/cli --clone=false
- Forking cli/cli...
✓ Created fork user/cli
~/Projects$
```
