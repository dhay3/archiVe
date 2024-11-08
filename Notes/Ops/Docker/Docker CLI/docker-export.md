# docker-export

用于将容器的fs以压缩文件的形式导出，也可以使用`--output`参数

```
root in /opt/t λ docker export t1 > t1.tar
root in /opt/t λ ls
t1.tar
root in /opt/t λ tar -xf t1.tar
root in /opt/t λ ls
bin  dev  etc  home  opt  proc  root  sys  t1.tar  tmp  usr  var
```

