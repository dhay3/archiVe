# gh completion

ref

https://cli.github.com/manual/gh_completion

https://unix.stackexchange.com/questions/239528/dynamic-zsh-autocomplete-for-custom-commands

这里只介绍如何在 oh-my-zsh 中使用 zsh completion 

先查看 oh-my-zsh 补全脚本的存储位置

```
print -rl -- $fpath
```

将补全脚本写入存储的位置即可

```
gh completion -s zsh > ~/.oh-my-zsh/completions/_gh
```

