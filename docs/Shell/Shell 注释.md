# Shell 注释

## 单行注释

```bash
[root@cyberpelican opt]# cat test.sh
#!/bin/bash
#this is a test
echo hello 

---

[root@cyberpelican opt]# /bin/sh test.sh 
hello
```

## 多行注释

- `<< COMMENT ... COMMENT`

  ```
  [root@cyberpelican opt]# cat test.sh 
  #!/bin/bash
  echo hello 
  << a
  this is a test
  this is a test
  a
  ```

  COMMENT可以自定义

- `:' ... '`

  ```
  [root@cyberpelican opt]# cat test.sh 
  #!/bin/bash
  echo hello 
  : '
  this is a test
  this is a test
  '
  ```

## 
