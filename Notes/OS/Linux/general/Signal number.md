# Signal number

> 如果是zsh的默认不会显示number，需要使用`bash -c`

可以通过`kill -l`或`trap -l`来获取所有常见的posix signal number

也可以通过`kill -l <SIG>`来获取单一的sig

```
cpl in ~/note/docs/Linux/basic on master λ kill -l SIGHUP
1
```

