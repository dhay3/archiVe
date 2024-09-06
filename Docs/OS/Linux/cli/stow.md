---
createTime: 2024-09-03 10:35
tags:
  - "#hash1"
  - "#hash2"
---

# stow

## 0x01 Preface

> GNU Stow is a symlink farm manager which takes distinct sets of software and/or data located in separate directories on the filesystem, and makes them all appear to be installed in a single directory tree.

简而言之 stow 就是一个配置管理器，通过 symlink 可以让不同位置的配置集中在一个目录进行管理

## 0x02 Terms

### 0x02a package

> A *package* is a related collection of files and directories that you wish to administer as a unit — e.g., Perl or Emacs — and that needs to be installed in a particular directory structure — e.g., with bin, lib, and man subdirectories.

简单的说 package 就是你想要管理的内容(**package 必须是目录**，不能是文件否则会出现 `stow: ERROR: The stow directory xxx does not contain package xxx` )，package 必须在 stow directory 内

### 0x02b target directory

> A *target directory* is the root of a tree in which one or more packages wish to appear to be installed. /usr/local is a common choice for this, but by no means the only such location. Another common choice is ~ (i.e. the user’s `$HOME` directory) in the case where Stow is being used to manage the user’s configuration (“dotfiles”) and other files in their `$HOME`. The examples in this manual will use /usr/local as the target directory.

简单的说 target directory 就是 symlink 的存放的地址，可以通过 `-t | --target` 来设置，默认为 stow directory 父级目录

> A *stow directory* is the root of a tree containing separate packages in private subtrees. When Stow runs, it uses the current directory as the default stow directory. The examples in this manual will use /usr/local/stow as the stow directory, so that individual packages will be, for example, /usr/local/stow/perl and /usr/local/stow/emacs.

简单的说 stow direcotory 就是包含 packages 的目录。如果没有指定(通过 `-d | --dir`)，默认会使用 current directory 

## 0x03 Usage

```
stow [options] [action flag] package …
```

> [!NOTE]
> stow 的过程就是将 stow directory 中的 packages 在 target directory 中生成对应 packages 的 symlink

### 0x03a Optional args

- `-d dir | --dir=dir`

	指定 stow directory，如果没有指定该参数，且有设置过 STOW_DIR 环境变量，会优先使用该值。如果没有使用该参数且没有设置过 STOW_DIR ，默认会使用当前目录

- `-t dir | --target=dir`

	指定 target directory，如果没有指定该参数，默认会使用 stow directory 的父级目录

- `-n | --no | --simulate`

	以 dryrun 的方式运行 stow，通常和 `-v` 一起使用输出详细信息

- `-v | --verbose=[n]`

	输出详细信息，0 - 5 表示详细程度

- `--dotfiles`

	如果 stow directory 中的 package 文件名以 `dot-` 开头，会在 target directory 中生成以 `.` 开头的 symlink

- `-D | --delete`

	从 target directory 删除对应 stow directory package 中的 symlink

- `-R | --restow`

	先做 `-D` 的操作然后，再做 stow

- `--ignore=regex`

	在 stow directory 中不对符合 regex 的文件做 stow 的操作
	regex 如 `\.git*`

## 0x04 Examples

例如有如下层级的目录，现在想要将 package 通过 stow 放到  t 中

```
$ lta
 .
├──  d
│   └──  package
│       ├──  .c
│       ├──  a
│       ├──  b
│       └──  dot-d
└──  t
```

就可以使用如下指令

```
$ stow --verbose=2 -d d -t t --dotfiles package
   stow dir is /tmp/test/d
   stow dir path relative to target /tmp/test/t is ../d
   Planning stow of: package ...
   Planning stow of package package...
       level of .c is 0
   LINK: .c => ../d/package/.c
       level of a is 0
   LINK: a => ../d/package/a
       level of b is 0
   LINK: b => ../d/package/b
       level of dot-d is 0
   LINK: .d => ../d/package/dot-d
   Planning stow of package package... done
   Processing tasks...
   Processing tasks... done
```

那么效果就如下

```
$ lta
 .
   ├──  d
   │   └──  package
   │       ├──  .c
   │       ├──  a
   │       ├──  b
   │       └──  dot-d
   └──  t
       ├──  .c ⇒ ../d/package/.c
       ├──  .d ⇒ ../d/package/dot-d
       ├──  a ⇒ ../d/package/a
       └──  b ⇒ ../d/package/b
```

## 0x05 Github

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***References***

- [Stow](https://www.gnu.org/software/stow/manual/stow.html)


***FootNotes***

