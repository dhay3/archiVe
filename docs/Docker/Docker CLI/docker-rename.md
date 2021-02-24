# docker rename

docker rename 用于重命名容器

pattern：`docker rename <old-container> <new-container>`

```
root in ~ λ docker rename t2 t3
root in ~ λ docker ps
CONTAINER ID   IMAGE     COMMAND            CREATED          STATUS          PORTS     NAMES
da045bd29d53   cbb       "/bin/sh -c top"   4 minutes ago    Up 4 minutes              t3
c082c5068def   busybox   "sh"               25 minutes ago   Up 25 minutes             t1
```

