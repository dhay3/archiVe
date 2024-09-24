# Shell select

`select`结构主要用来生成简单的菜单。

pattern：

```
select name
[in list]
do
  commands
done
```

1. `select`生成一个菜单，内容是列表`list`的每一项，并且每一项前面还有一个数字编号。
2. Bash 提示用户选择一项，输入它的编号。
3. 用户输入以后，Bash 会将该项的内容存在变量`name`，该项的编号存入环境变量`REPLY`。如果用户没有输入，就按回车键，Bash 会重新输出菜单，让用户选择。
4. 执行命令体`commands`
5. 执行结束后，回到第一步，重复这个过程。

## Example01

```
[root@cyberpelican ~]# cat test.sh 
select brand in Samsung Sony iphone symphony Walton
do
  echo "You have chosen $brand"
done
[root@cyberpelican ~]# ./test.sh 
1) Samsung
2) Sony
3) iphone
4) symphony
5) Walton
#? 2
You have chosen Sony
#? ^C[root@cyberpelican ~]#
```

如果用户没有输入编号，直接按回车键。Bash 就会重新输出一遍这个菜单，直到用户按下`Ctrl + c`，退出执行。

## Example02

`select`可以与`case`结合，针对不同项，执行不同的命令。

```
[root@cyberpelican ~]# ./test.sh 
./test.sh: line 1: o: command not found
1) Ubuntu
2) LinuxMint
3) Windows8
4) Windows10
5) WindowsXP
#? 3
Why don't you try Linux?
#? 2
I also use LinuxMint.
#? 

```

