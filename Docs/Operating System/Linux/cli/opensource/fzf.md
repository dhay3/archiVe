# fzf

ref

https://github.com/junegunn/fzf#usage

https://www.freecodecamp.org/news/fzf-a-command-line-fuzzy-finder-missing-demo-a7de312403ff/

https://www.baeldung.com/linux/fzf-command

## Digest

fzf a command-line fuzzy finder

`fzf` 是一个非常好用的查询工具，完全可以替代 `find` 且支持的功能更多 

1. 可以直接从 Stdin 中读取内容查询
2. interactive 交互式

## Supported search syntax

| Token     | Match type                 | Description                          |
| --------- | -------------------------- | ------------------------------------ |
| `sbtrkt`  | fuzzy-match                | Items that match `sbtrkt`            |
| `'wild`   | exact-match (quoted)       | Items that include `wild`            |
| `^music`  | prefix-exact-match         | Items that start with `music`        |
| `.mp3$`   | suffix-exact-match         | Items that end with `.mp3`           |
| `!fire`   | inverse-exact-match        | Items that do not include `fire`     |
| `!^music` | inverse-prefix-exact-match | Items that do not start with `music` |
| `!.mp3$`  | inverse-suffix-exact-match | Items that do not end with `.mp3`    |

## Optional args

- `-x | --extended`

  extended-search mode，this is enable by default

  默认模糊查询

- `-e | --excat`

  enable exact-match

- `-i`

  case-insensitive match( default: smart-case match )

- `+i`

  case-sensitive match

- `--reverse`

  display from the top of the screen

- `-f | --filte=STR`

  show matches from StR  without the interactive finder

  一般用于脚本

## Cautions

### search position

==fzf 默认从当前目录开始查询==

例如 如果需要搜索 `libcrypto.so`，假设在当前用户的 HOME 目录下，默认 fzf 查不到，需要到根目录才能查询到

### search hidden files

https://github.com/junegunn/fzf/issues/337

fzf 默认不会查询 hidden files，如果需要查询 hidden files

```
#手动配置查询 hidden files
find . | fzf

#默认设置查询 hidden files
export FZF_DEFAULT_COMMAND='find .'
```

### fzf in zsh

> 这里需要注意的是`<TAB>` 只能在末尾才会生效

fzf 支持将 `**<TAB>` 扩展文件，但是需要在 zsh 中安装对应的 fzf 插件

```
plugins=(fzf)
```

## Examples

匹配所有以`.log` 结尾，且不包含 centos 字符的内容

```
.log$ !centos
```

查询所有匹配`/etc/h`的文件

```
ls /etc/h**<TAB>
```

查询文件并用 `vim` 打开

```
vim $(fzf)
```

