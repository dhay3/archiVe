# Linux type

type用来查看命令是内建的还是外部指令。别和windows中的type记反了

```
chz@cyberpelican:/root$ type echo
echo is a shell builtin
chz@cyberpelican:/root$ type ls
ls is aliased to `ls --color=auto'
```

使用`-a`可以查看命令的路径

```
chz@cyberpelican:/root$ type -a ls
ls is aliased to `ls --color=auto'
ls is /usr/bin/ls
ls is /bin/ls
```

