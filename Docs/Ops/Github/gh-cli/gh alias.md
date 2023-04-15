# gh alias

ref

https://cli.github.com/manual/gh_alias_set

和 git 一样 gh 也可以设置 alias

## list

syntax

```
gh alias list
```

用于查看当前设置的 alias

```
gh alias list               
co: pr checkout
ro: repo
se: search
```

## uset

syntax

```
gh alias set <alias> <expansion> [flags]
```

`<expansion> [flags]` 部分为 alias 实际执行的部分。如果 alias 部分包含 positional placeholder 例如 `$1`，alias 后面的第一个参数就会被映射到 `$1` 。也可以使用 `-` 表示从 stdin 输入

目前不太推荐使用该功能，因为会导致不能使用 `gh completion` 的功能

### Optional args

- `-s | --shell`

  `<expansion> [flags]` 部分直接由 shell 解析

### Examples

```
gh alias set ro 'repo'
gh alias set iss 'issue'
gh alias set se 'search'
gh alias set login 'auth login'
gh alias set logout 'auth logout'
gh alias set st 'auth status'
```

## delete

syntax

```
gh alias delete <alias>
```

用于删除 alias

### Examples

删除指定的 alias

```
gh alias delete co
✓ Deleted alias co; was pr checkout
```

删除所有的 alias

```
(base) cpl in ~/hugo/themes λ gh alias list | awk -F : '{print($1)}' | xargs -i gh alias delete {}
✓ Deleted alias iss; was issue
✓ Deleted alias login; was auth login
✓ Deleted alias logout; was auth logout
✓ Deleted alias ro; was repo
✓ Deleted alias rov; was repo view
✓ Deleted alias se; was search
✓ Deleted alias st; was auth status
```