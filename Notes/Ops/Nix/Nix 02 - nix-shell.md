---
createTime: 2024-09-10 16:26
tags:
  - "#hash1"
  - "#hash2"
---

# Nix 02 - nix-shell

## 0x01 Preface

nix-shell 会根据参数启动一个 subshell(会根据 NIX_BUILD_SHELL 环境变量来决定 shell，默认 bash)

具体内容请查看 nix-shell 文档[^1]

## 0x02 Synopsis

```
nix-shell [--arg name value] [--argstr name value] [{--attr | -A} attrPath] [--command cmd] [--run cmd] [--exclude regexp] [--pure] [--keep name] {{--packages | -p} {packages | expressions} … | [path]}
```

如果没有指定 path，默认会优先使用当前目录的 `shell.nix`，如果 `shell.nix` 也不存在就会使用 `default.nix`

### 0x02a Optional args

- `--command <cmd>`

	在 interactive shell 中运行 cmd

	虽然是在 interactive shell 中运行，但是如果 cmd 运行完成后，Nix 会发送一个 `exit` 的指令推出 interactive command。如果想要 interactive shell 在执行完 cmd 后不退出，需要在 cmd 后添加 `return` 关键字

	例如：

	`nix-shell -p bash --command "top;return"`

- `--run <cmd>`

	以 non-interactive shell 的方式运行 cmd

- `--pure`

	以 pure cleared environment (环境变量还原)运行 shell

- `--keep <names...>`

	和 `--pure` 一起使用，表示指定的环境变量除外

- `--packages | -p <packages...>`

	在 Ad hoc shell 中会安装指定的 packages

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***References***

- [nix-shell - Nix Reference Manual](https://hydra.nixos.org/build/271691842/download/1/manual/command-ref/nix-shell.html)


***FootNotes***

[^1]:[nix-shell - Nix Reference Manual](https://hydra.nixos.org/build/271691842/download/1/manual/command-ref/nix-shell.html)