# getopts

解析复杂的脚本命令行参数，通常与`while`循环一起使用，取出脚本所有的带有前置连词线（`-`）的参数。

pattern：`getopts optstring name`

- optstring：表示脚本所有的连词线参数。比如，某个脚本可以有三个配置项参数`-l`、`-h`、`-a`，其中只有`-a`可以带有参数值，而`-l`和`-h`是开关参数，那么`getopts`的第一个参数写成`lha:`，==顺序不重要==。注意，==`a`后面有一个冒号，表示该参数带有参数值==，`getopts`规定带有参数值的配置项参数，后面必须带有一个冒号（`:`）。

- name：是一个变量名，用来保存当前取到的配置项参数，即`l`、`h`或`a`。

## exmaple

注意，只要遇到不带连词线的参数，`getopts`就会执行失败，从而退出`while`循环。比如，`getopts`可以解析`command -l foo`，但不可以解析`command foo -l`。另外，多个连词线参数写在一起的形式，比如`command -lh`，`getopts`也可以正确处理。

```
[root@cyberpelican opt]# cat test.sh 
while getopts 'lha:' OPTION; do
  case "$OPTION" in
    l)
      echo "linuxconfig"
      ;;

    h)
      echo "h stands for h"
      ;;

    a)
      #avalue="$OPTARG"
      echo "The value provided is $OPTARG"
      ;;
    ?)
      echo "script usage: $(basename $0) [-l] [-h] [-a somevalue]" >&2
      exit 1
      ;;
  esac
done
[root@cyberpelican opt]# ./test.sh -la testValue
linuxconfig
The value provided is testValue
```

> $OPTAGR保存带有参数的连词线参数的值