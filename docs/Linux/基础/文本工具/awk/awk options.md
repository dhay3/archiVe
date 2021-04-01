# awk options

- -f

  指定awk脚本从文件中来，而不是命令行。

  ```
  root in /usr/local/\ λ cat t.awk
  {print $0}
  
  root in /usr/local/\ λ df -hT | awk -f t.awk
  Filesystem     Type      Size  Used Avail Use% Mounted on
  udev           devtmpfs  2.0G     0  2.0G   0% /dev
  tmpfs          tmpfs     395M  6.0M  389M   2% /run
  /dev/vda1      ext4       40G  5.9G   32G  16% /
  tmpfs          tmpfs     2.0G     0  2.0G   0% /dev/shm
  tmpfs          tmpfs     5.0M     0  5.0M   0% /run/lock
  tmpfs          tmpfs     2.0G     0  2.0G   0% /sys/fs/cgroup
  tmpfs          tmpfs     395M     0  395M   0% /run/user/0
  ```

  ==常用在shebang中使用==

  ```
  root in /usr/local/\ λ cat t.awk
  #!/usr/bin/mawk -f
  {print $0}
  
  root in /usr/local/\ λ df -hT | ./t.awk
  Filesystem     Type      Size  Used Avail Use% Mounted on
  udev           devtmpfs  2.0G     0  2.0G   0% /dev
  tmpfs          tmpfs     395M  6.0M  389M   2% /run
  /dev/vda1      ext4       40G  5.9G   32G  16% /
  tmpfs          tmpfs     2.0G     0  2.0G   0% /dev/shm
  tmpfs          tmpfs     5.0M     0  5.0M   0% /run/lock
  tmpfs          tmpfs     2.0G     0  2.0G   0% /sys/fs/cgroup
  tmpfs          tmpfs     395M     0  395M   0% /run/user/0
  ```

- -F

  指定分隔符

- -v var=val

  分配一个环境变量，可以覆盖原有的

  ```
  [root@chz t]# awk -v shell=$SHELL 'BEGIN{print shell}'
  /bin/bash
  ```

- --posix

  使用posix规则

- -d

  将awk调用的环境变量存储在`awkvars.out`

  ```
  [root@chz t]# ip a | awk -d '/192\.168\.[0-9]{1,3}\.[0-9]{1,3}/{print$0}'
      inet 192.168.80.140/24 brd 192.168.80.255 scope global noprefixroute dynamic ens33
      inet 192.168.122.1/24 brd 192.168.122.255 scope global virbr0
  [root@chz t]# ls
  1  2  3  awkvars.out
  [root@chz t]# cat awkvars.out 
  ARGC: 1
  ARGIND: 0
  ARGV: array, 1 elements
  BINMODE: 0
  CONVFMT: "%.6g"
  ERRNO: ""
  FIELDWIDTHS: ""
  FILENAME: "-"
  FNR: 20
  FPAT: "[^[:space:]]+"
  FS: " "
  IGNORECASE: 0
  LINT: 0
  NF: 4
  NR: 20
  OFMT: "%.6g"
  OFS: " "
  ORS: "\n"
  RLENGTH: 0
  RS: "\n"
  RSTART: 0
  RT: "\n"
  SUBSEP: "\034"
  TEXTDOMAIN: "messages"
  ```

  
