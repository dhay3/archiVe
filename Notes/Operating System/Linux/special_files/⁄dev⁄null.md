# Linux  /dev/null

`/dev/null`是Unix系统中的一个特殊设备文件，它丢弃一切写入其中的数据

```
root in ~ λ cat ~/.zshrc > /dev/null                                                           /0.0s
root in ~ λ cat /dev/null    

root in /opt λ echo $(test) 1 > /dev/null 2>&1      #输入的同时将stderr输入到stdout中，但是默认忽略1
```

可以使用`/dev/null`文件快速清空文件内容

```
root in /opt λ cat test
hello                                                                            /0.0s
root in /opt λ cat /dev/null > test                                              /0.0s
root in /opt λ cat test            
```

可以将配置文件使用链接指向`/dev/null`来让系统忽略配置

```
ln /etc/systemd/resolved.conf /dev/null
```

