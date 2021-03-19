# Shell Shebang

参考：

https://zh.wikipedia.org/wiki/Shebang

https://tldp.org/LDP/abs/html/sha-bang.html

`#!`(shebang也被称为Hashbang)。在文件存在Shebang的情况下，==Unix操作系统的程序加载器会分析Shebang后的内容，将这些内容做为解释器指令，并调用该指令。==

例如，以指令`#!/bin/sh`开头的文件在执行时会实际调用`/bin/sh`程序（通常是Bourne shell或兼容的shell，例如bash、dash等）来执行。如果是以BSD为基础了需要在shebang后加空格，例如void。

常见的一些shebang

- `#!/usr/bin/perl -w`
- `#!/usr/bin/python -O`
- `#!/usr/bin/make -f`
- `#!/bin/awk -f`
- `#!/bin/sed -f`
- `#!/usr/bin/env bash` 兼容

## note

- Shebang只能在第一行才能被解释
- To avoid this possibility, a script may begin with a `#!/bin/env bash` *sha-bang* line. This may be useful on UNIX machines where *bash* is not located in `/bin`

- 不同的shell有不同的语法和特性，例如zsh的rehash但是bash没有。所以如果想要以特定的shell执行脚本就需要指定shebang

  ```
  root@ubuntu18:/usr/local/\# cat shebang.sh
  rehash
  
  root@ubuntu18:/usr/local/\# echo $0
  bash
  root@ubuntu18:/usr/local/\# ./shebang.sh
  ./shebang.sh: line 1: rehash: command not found
  
  root@ubuntu18:/usr/local/\# cat shebang.sh
  #!/bin/zsh
  rehash
  root@ubuntu18:/usr/local/\# ./shebang.sh
  ```

- `./shebang.sh`和`sh shebang.sh`是不同的，alse check [the defferent of source and bash](./Shell source)

  ```
  root in /usr/local/\ λ cat shebang.sh
  #!/bin/cat
  df -hT
  
  #这里没有手动指定解释器，默认使用shebang指定的解释器。caution! 这里同样会输出shebang所有的行
  root in /usr/local/\ λ ./shebang.sh
  #!/bin/cat
  df -hT
  
  #这里使用zsh做为解释器
  root in /usr/local/\ λ zsh shebang.sh
  Filesystem     Type      Size  Used Avail Use% Mounted on
  udev           devtmpfs  2.0G     0  2.0G   0% /dev
  tmpfs          tmpfs     395M  6.0M  389M   2% /run
  /dev/vda1      ext4       40G  5.9G   32G  16% /
  tmpfs          tmpfs     2.0G     0  2.0G   0% /dev/shm
  tmpfs          tmpfs     5.0M     0  5.0M   0% /run/lock
  tmpfs          tmpfs     2.0G     0  2.0G   0% /sys/fs/cgroup
  tmpfs          tmpfs     395M     0  395M   0% /run/user/0
  ```

  ==但都会开一个sub shell 去执行脚本==，[check the means of  bash $$](https://www.gnu.org/software/bash/manual/bash.html)

  ```
  root in /usr/local/\ λ cat shebang.sh
  echo "pid=$$"
  root in /usr/local/\ λ echo $$
  8679
  root in /usr/local/\ λ ./shebang.sh
  pid=9166
  root in /usr/local/\ λ sh shebang.sh
  pid=9169
  ```

- starting a `README` file with a `#!/bin/more`, and making it executable. The result is a self-listing documentation file

## tricks

- 结合awk

  ```
  root in /usr/local/\ λ cat t.awk
  #!/usr/bin/mawk -f
  /tmpfs/{print $0}
  root in /usr/local/\ λ df -hT | ./t.awk
  udev           devtmpfs  2.0G     0  2.0G   0% /dev
  tmpfs          tmpfs     395M  6.0M  389M   2% /run
  tmpfs          tmpfs     2.0G     0  2.0G   0% /dev/shm
  tmpfs          tmpfs     5.0M     0  5.0M   0% /run/lock
  tmpfs          tmpfs     2.0G     0  2.0G   0% /sys/fs/cgroup
  tmpfs          tmpfs     395M     0  395M   0% /run/user/0
  ```

- 结合makefile，等价make file1

  ```
  root in /usr/local/\ λ cat t.make
  #!/usr/bin/make -f
  
  file1:
   echo "file1" >> file1
  root in /usr/local/\ λ ./t.make
  echo "file1" >> file1
  root in /usr/local/\ λ cat file1
  file1
  ```

  

