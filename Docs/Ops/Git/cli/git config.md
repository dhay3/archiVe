# git config

ref

https://git-scm.com/docs/git-config

https://git-scm.com/book/en/v2/Customizing-Git-Git-Configuration

## Digest

Git 可以使用 `git config` 来为不同域设置不同的配置，按照域( scope )会被存储在不同的位置

1. `/etc/gitconfig`

   ontains values applied to every user on the system and all their repositories

   全局配置，可以使用 `git --system config` 来设置，优先级低

2. `~/.gitconfig | ~/.config/git/config`

   Values specific personally to you, the user

   用户配置，可以使用 `git --global config` 来设置，优先级中

3. `.git/config`

   in the Git directory (that is, `.git/config`) of whatever repository you’re currently using: Specific to that single repository.

   仓库配置，可以使用 `git --local config` 来设置，优先级高

如果需要给 key 赋值可以使用 `git config core.editor typora` 类似的方式，key 会按照配置文件的优先级被 override

## Optional args

- `--system/--global/--local`

  for writing options

  只在写操作是生成，表示各个域

- `-l | --list`

  查看当前应用的配置

- `--show-origin`

  查看应用的配置，从那个文件来

  ```
  git config --list --show-origin
  file:/root/.gitconfig   init.defaultbranch=main
  file:/root/.gitconfig   usre.name=John Doe
  file:config     core.repositoryformatversion=0
  file:config     core.filemode=true
  file:config     core.bare=false
  ```

- `--get <key>`

  查询指定 key 的值

  ```
  git config --get user.name
  user.name John Doe
  ```

  也可以不指定 `--get`

  ```
  git config user.name
  user.name John Doe
  ```

- `--get-regexp <pattern>`

  按照正则查询配置

  ```
  git config --get-regexp name
  user.name John Doe
  ```

- `--unset <key>`

  删除指定的配置

  ```
  git config --unset diff.renames
  ```

- `--remove-section <section>`

  删除指定的 section 配置

- `-e`

  打开 editor 编辑配置文件，需要和 `--global`, `--system`, `--local` 一起使用，默认 `--local`

## Configuration file

> 具体查看 man page

### Boolean

如果配置的 key 使用的类型是布尔值

- `yes`, `on`, `true`, `1` 均表示 true
- `no`, `off`,`true`,  `0` 均表示 false

### Section

在 Git 中，配置文件按照 section 来划分功能。section 还可以划分成 subsection，但是需要按照如下格式

```
[section "subsection"]
```

subsection name 必须在 double quote 内

### Includes

和 `nginx` 配置一样，Git 也有 `Include` 指令 (不能通过 `git config key=value` 来设置)

```
[include]
	path=.gitconfig~
```

例如上述配置就会将相对路径 `.gitconfig~` 中的配置内容导入当配置文件中。同样的也可以使用绝对路径

```
[include]
    path = /path/to/foo.inc ; include by absolute path
    path = foo.inc ; find "foo.inc" relative to the current file
    path = ~/foo.inc ; find "foo.inc" in your `$HOME` directory
```

还有一个 `includeIf` 具体可以参考 man page

## Example

### 常用配置

```
git config --global user.name cyberPelican
git config --global user.email 62749885+dhay3@users.noreply.github.com
git config --global core.editor vim
git config --global core.autocrlf input
git config --global core.ignoreCase true
git config --global color.ui always
git config --global init.defaultBranch main

#gpg
git config --global commit.gpgSign 1
git config --global push.gpgSign 1
git config --global tag.gpgSign 1
git config --global user.signingKey DFEBFAF653ED6ACCEEA7ADF1F3A82ABD5E016AC9 



#别名
git config --global alias.br branch
git config --global alias.co checkout
git config --global alias.sw switch
git config --global alias.st status
git config --global alias.ci commit
git config --global alias.lg1  "log --all --graph --format='format:%C(auto)%h%C(reset) - %C(auto)(%ar)%C(reset) %C(bold white)%s%C(reset) - %C(auto)%an%C(reset) %C(auto)%d%C(reset)'"
git config --global alias.lg2  "log --all --graph --format='format:%C(auto)%h%C(reset) - %C(auto)(%aD)%C(reset) %C(bold white)%s%C(reset) - %C(auto)%an%C(reset) %C(auto)%d%C(reset)'"

#分页
git config --global core.pager less
git config --global pager.show false
git config --global pager.branch false
git config --global pager.tag false
git config --global pager.log false
git config --global pager.reflog false
git config --global pager.diff false

#密码保存策略
git config --global credential.helpr  'cache --timeout 3600'
git config --global credential.helpr store 
```

### 额外配置

```
git config --global diff.tool kdiff3
git config --global difftool.prompt false
git config --global difftool.keepBackup false
git config --global difftool.trustExitCode false
git config --global merge.tool kdiff3
git config --global mergetool.prompt false
git config --global mergetool.keepBackup false
git config --global mergetool.trustExitCode false
```

### 选择配置

```
#代理
git config --global  core.gitProxy
git config --global core.sshCommand

#日志
git config --global log.showSignature 1
```
