# Linux cp & install

## cp

> 由于是当前用户cp创建的文件，所以owner和owner-group都是当前用户和对应的组

syntax：`cp [options] <src...> <dest>`

`cp`也可以生成symbolic links或hard links

- -n | --no-clobber

  如果目标文件存在对文件不覆盖

  ```
  ➜  test cat 1
  ➜  test cat 2
  4
  ➜  test cp -n 2 1
  ➜  test cat 1
  ➜  test 
  ```

- -p | --preserver

  ==复制文件是保留mode，owner-ship，timestamps==

  ```
  ➜  test cp 1 3
  ➜  test ll
  .rwxrwxr-- root root 0 B Thu Jun 10 11:57:13 2021  1
  .rwxr-xr-- root root 0 B Thu Jun 10 12:00:37 2021  3
  ➜  test cp -p 1 4  
  ➜  test ll
  .rwxrwxr-- root root 0 B Thu Jun 10 11:57:13 2021  1
  .rwxr-xr-- root root 0 B Thu Jun 10 12:00:37 2021  3
  .rwxrwxr-- root root 0 B Thu Jun 10 11:57:13 2021  4
  ```

- -a

  等价于`-dR --preserve`

- -d | --no-dereference

  复制时复制link

- -L | --dereference

  复制时复制link指向的文件而不是link，==默认==

## install

install 命令用于拷贝文件的同时设置权限和属性

syntax；`install [options] <src> <dest>`

`-m`拷贝的同时设置权限，默认755

```
install -o root -g root -m 755 kubectl  /usr/local/bin/different encryption methods and formattion
```
