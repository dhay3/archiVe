# sed args

文本处理工具，如果没有指定`-e`，`-f`那么第一个不是参数的选项做为sed script

syntax：`sed [options] [script] [input-file] `

## Optional args

- -n 

  sed会输出原有的内容，可以使用该参数不输出

  ```
  [root@k8snode01 opt]# sed '/^label/P' Dockerfile
  FROM busybox
  label:hello world
  label:hello world
  pwd
  
  [root@k8snode01 opt]# sed -n '/^label/P' Dockerfile
  label:hello world
  ```

- -e script

  指定多个script，cli形式

  ```
  [root@k8snode01 opt]# sed  -e 'i\3' -e 'a\4' Dockerfile
  3
  FROM busybox
  4
  3
  RUN ["sh","-c","echo ------------building-------------"]
  4
  3
  WORkDIR /
  4
  3
  CMD ["sh","-i"]
  4
  ```

- -f script-file

  指定一个文件含有多个script

- -i [suffix]

  修改文件，如果有指定suffix就会生成备份文件。`-i`后不能带有空格.。==如果和`-n`一起使用回导致原来的文件内容为空==

  ```
  [root@k8snode01 opt]# sed  -i.bak 'i\3' Dockerfile
  [root@k8snode01 opt]# ls
  calico.yaml  cni  Dockerfile  Dockerfile.bak  index.html  nginx.conf  rh
  [root@k8snode01 opt]# diff Dockerfile Dockerfile.bak
  1d0
  < 3
  3d1
  < 3
  5d2
  < 3
  7d3
  < 3
  ```

- --posix

  关闭GUN扩展，同awk

- -r

  使用extended regular expression

- {}

  子命令

  ```
  [root@k8snode01 opt]# seq 10 | sed  '4,10{d;p}'
  1
  2
  3
  ```

  



