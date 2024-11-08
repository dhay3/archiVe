# sed command

## comment

在sed中，使用`#`表示注释。如果使用了comment必须和换行结合，不能使用`;`来分隔

```
$seq 4 | sed '# print 2 4
{1d;3d}'
2
4
```

## basic command

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

  将当前目录下的指定文件，添加到匹配行的下一行

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

- b label

  无条件切换到label，`:x`表示label（和dos 中的 goto 类似）

  ```
  seq 6 | sed -n 'bx;5d;:x;p'
  1
  2
  3
  4
  5
  6
  ```

- N

  pattern space 添加换行符，然后将下一行拼接到pattern space

  ```
  $seq 4 | sed  -n '2{N;p}'
  2
  3
  ```

## special command

- =

  打印出行号

  ```
  $sed = /etc/resolv.conf
  1
  search tbsite.net aliyun.com
  2
  options timeout:2 attempts:2
  3
  nameserver 10.195.29.1
  4
  nameserver 10.195.29.17
  5
  nameserver 10.195.29.33
  ```

- n

  将下一行作为pattern space，即每隔指定行操作。例如如下命令表示每隔两行操作一次

  ```
  $ seq 6 | sed 'n;n;s/./x/'
  1
  2
  x
  4
  5
  x
  
  sed -n 'n;p' /etc/resolv.conf
  options timeout:2 attempts:2
  nameserver 10.195.29.17
  
  $cat /etc/resolv.conf
  search tbsite.net aliyun.com
  options timeout:2 attempts:2
  nameserver 10.195.29.1
  nameserver 10.195.29.17
  nameserver 10.195.29.33
  ```

- { command }

  对address做聚合操作

  ```
  $ seq 3 | sed -n '2{s/2/X/ ; p}'
  X
  ```

- !

  取反，即对匹配的内容取反。例如如下命令表示非4的行替换成40

  ```
  $seq 4 | sed '/4/!s/./40/'
  40
  40
  40
  4
  ```

## multi command

https://www.gnu.org/savannah-checkouts/gnu/sed/manual/sed.html#Multiple-commands-syntax

如果要对一个文件执行多个命令，可以使用以下几种方法

1. 换行

   `a`，`c`，`i`对文本做增量的都必须使用换行符或`-e`来分隔

   ```
   [admin@gonda033059000078.na175 /home/admin]
   $seq 6 | sed '1d
   > 2d'
   3
   4
   5
   6
   ```

2. `-e`参数

   ```
   $ seq 6 | sed -e 1d -e 3d -e 5d
   2
   4
   6
   ```

3. 使用`;`

   ```
   $ seq 6 | sed '1d;3d;5d'
   2
   4
   6
   ```

4. curly bracket聚合

   ```
   $seq 6 | sed '{1d;3d}'
   2
   4
   5
   6
   ```

