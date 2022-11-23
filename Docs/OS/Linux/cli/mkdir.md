# Linux mkdir

## Digest

syntax

```
mkdir [OPTION]... DIRECTORY...
```

用于创建目录

## Optional args

- `-m | --mode=MODE`

  创建目录时指定文件的 mode permission（使用 `chmod` 相同格式）

  ```
  [vagrant@localhost ~]$ mkdir test
  [vagrant@localhost ~]$ ll
  total 0
  drwxrwxr-x. 2 vagrant vagrant 6 Nov 22 14:01 test
  [vagrant@localhost ~]$ mkdir -m 777 test2
  [vagrant@localhost ~]$ ll
  total 0
  drwxrwxr-x. 2 vagrant vagrant 6 Nov 22 14:01 test
  drwxrwxrwx. 2 vagrant vagrant 6 Nov 22 14:01 test2
  ```

- `-p | --parents`

  创建目录时，如果父目录不存在会报错，使用该参数时如果父目录不存在自动创建
