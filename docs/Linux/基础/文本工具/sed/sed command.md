# sed command

- a\text  | a text

  append text after a line

  ```
  [root@k8snode01 opt]# sed '1a3' Dockerfile
  FROM busybox
  3
  RUN ["sh","-c","echo ------------building-------------"]
  WORkDIR /
  CMD ["sh","-i"]
  label:hello world
  ```

- i \text | i text

  insert text before a line

- c\text | c text

  repalce change lines with text

  ```
  [root@k8snode01 opt]# sed '2c3' Dockerfile
  FROM busybox
  3
  WORkDIR /
  CMD ["sh","-i"]
  label:hello world
  ```

- d

  delete the pattern space

  ```
  [root@k8snode01 opt]# sed '1 d' Dockerfile
  RUN ["sh","-c","echo ------------building-------------"]
  WORkDIR /
  CMD ["sh","-i"]
  label:hello world
  [root@k8snode01 opt]# sed '1d' Dockerfile
  RUN ["sh","-c","echo ------------building-------------"]
  WORkDIR /
  CMD ["sh","-i"]
  label:hello world
  ```

- e 

  执行指定位置的命令，并替换

  ```
  [root@k8snode01 opt]# sed '3e' Dockerfile
  FROM busybox
  label:hello world
  /opt
  ```

- e command

  将command执行的结果替换到指定的位置，==执行的结果被认为一行==

  ```
  [root@k8snode01 opt]# sed '1,2e ls' Dockerfile
  calico.yaml
  cni
  Dockerfile
  index.html
  nginx.conf
  rh
  FROM busybox
  calico.yaml
  cni
  Dockerfile
  index.html
  nginx.conf
  rh
  label:hello world
  pwd
  ```

- F

  打印文件的名字

  ```
  [root@k8snode01 opt]# sed '3F' Dockerfile
  FROM busybox
  label:hello world
  Dockerfile
  pwd
  ```

- g

  删除并添加CR

  ```
  [root@k8snode01 opt]# sed '/^label/g' Dockerfile
  FROM busybox
  
  pwd
  ```

- G

  append a newline to the contents  of the pattern space

  ```
  [root@k8snode01 opt]# sed '/^label/G' Dockerfile
  FROM busybox
  label:hello world
  
  pwd
  ```

- p

  打印内容

  ```
  [root@k8snode01 opt]# sed -n  '/^label/P' Dockerfile
  label:hello world
  ```

- q[exit-code] | Q[exit-code]

  已制定的exit-code退出sed，Q不会打印匹配的内容。如果没有指定exit-code以0退出

  ```
  [root@k8snode01 opt]# sed   '/^label/q127' Dockerfile;echo $?
  FROM busybox
  label:hello world
  127
  [root@k8snode01 opt]# sed   '/^label/Q127' Dockerfile;echo $?
  FROM busybox
  127
  [root@k8snode01 opt]# sed   '/^label/q' Dockerfile;echo $?
  FROM busybox
  label:hello world
  0
  ```

- r filename

  对去当前目录下的文件，添加到匹配行的下一行

  ```
  [root@k8snode01 opt]# cat test;echo ------;sed '/^label/r test' Dockerfile
  123
  ------
  FROM busybox
  label:hello world
  123
  pwd
  ```

- s/regexp/replacement/[flags]

  替换

  ```
  [root@k8snode01 opt]# sed 's/^label/vocal:hello world/' Dockerfile
  FROM busybox
  vocal:hello world:hello world
  pwd
  ```

## 