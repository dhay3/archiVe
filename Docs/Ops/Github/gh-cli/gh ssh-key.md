# gh ssh-key

https://cli.github.com/manual/gh_ssh-key_add

## add

syntax

```
gh ssh-key add [<key-file>] [flags]
```

用于添加 ssh public key 到 github

### Optional args

- `-t | --title <string>`

  指定 ssh key 使用的 title

## delete

syntax

```
gh ssh-key delete <id> [flags]
```

用于删除指定的而 ssh public key

### Optional args

- `-y | --yes`

## list

syntax

```
gh ssh-key list
```

显示上传的 public key