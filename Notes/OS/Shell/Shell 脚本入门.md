# Shell 脚本入门

https://www.gnu.org/software/bash/manual/bash.html

## 执行权限和路径

> 如果想通过`./script.sh`方式执行脚本使用shebang解释，需要给予`x`权限，且在正确的目录下。
>
> 但通过`sh script.sh`方式无需有`x`权限，且用sh做为解释器

- 权限

  ```
  # 给所有用户执行权限
  $ chmod +x script.sh
  
  # 给所有用户读权限和执行权限
  $ chmod +rx script.sh
  # 或者
  $ chmod 755 script.sh
  
  # 只给脚本拥有者读权限和执行权限
  $ chmod u+rx script.sh
  ```

- 路径

  除了执行权限，脚本调用时，一般需要指定脚本的路径（比如`path/script.sh`）。==如果将脚本放在环境变量`$PATH`指定的目录中，就不需要指定路径了。==因为 Bash 会自动到这些目录中，寻找是否存在同名的可执行文件。

  > 建议在主目录新建一个`~/bin`子目录，专门存放可执行脚本，然后把`~/bin`加入`$PATH`。
  >
  > ```
  > export PATH=$PATH:~/bin
  > ```
  >
  > 上面命令改变环境变量`$PATH`，将`~/bin`添加到`$PATH`的末尾。==可以将这一行加到`~/.bashrc`文件里面，然后重新加载一次`.bashrc`，这个配置就可以生效了。==
  >
  > 以后不管在什么目录，直接输入脚本文件名，脚本就会执行。

- source 命令

## 脚本参数

调用脚本时，脚本文件名或可以带有参数

```
$ script.sh word1 word2 word3
```

上面例子中，`script.sh`是一个脚本文件，`word1`、`word2`和`word3`是三个参数。

脚本文件内部，可以使用特殊变量，引用这些参数。

- `$0`：脚本文件名，即`script.sh`。
- `$1`~`$9`：对应脚本的第一个参数到第九个参数。
- `$#`：参数的总数。
- `$@`：全部的参数，参数之间使用空格分隔。==如果参数在双引号内被视为一个参数==
- `$*`：全部的参数，参数之间使用变量`$IFS`值的第一个字符分隔，默认为空格，但是可以自定义。

如果脚本的参数多于9个，那么第10个参数可以用`${10}`的形式引用，以此类推。

注意，如果命令是`command -o foo bar`，那么`-o`是`$1`，`foo`是`$2`，`bar`是`$3`。

```
[root@cyberpelican opt]# cat test.sh 
#!/bin/bash
echo '$1 = ' $1
echo '$2 = ' $2
echo '$3 = ' $3
echo '$4 = ' $4
[root@cyberpelican opt]# ./test.sh a b c d
$1 =  a
$2 =  b
$3 =  c
$4 =  d
```

