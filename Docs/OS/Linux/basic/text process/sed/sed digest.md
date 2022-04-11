# sed digest



https://stackoverflow.com/questions/12833714/the-concept-of-hold-space-and-pattern-space-in-sed

https://www.gnu.org/savannah-checkouts/gnu/sed/manual/sed.html

sed 有两种 data buffer，pattern space 和 hold space，默认初始均为空

sed 循环读取每一行，每读一行，会去掉后导换行符，然后将内容放入 pattern space，再执行命令。然后再将换行符放回，读取下一行。

## pattern space

When sed reads a file line by line, *the line that has been currently read is inserted into the pattern buffer (pattern space)*. Pattern buffer is like the temporary buffer, the scratchpad where the current information is stored. When you tell sed to print, it prints the pattern buffer.

## hold space

Hold buffer / hold space is like a long-term storage, *such that you can catch something, store it and reuse it later when sed is processing another line*. You do not directly process the hold space, instead, you need to copy it or append to the pattern space if you want to do something with it. For example, the print command `p` prints the pattern space only. Likewise, `s` operates on the pattern space.

## sed script

参考：

https://www.gnu.org/software/sed/manual/sed.html#sed-scripts

如果没有指定`-f`或`-e`参数，sed默认第一个非参数选项做为sed script，由三部分组成：

```
[address]command[options]
```

- address：If address is specified, the command X will be executed only on the matched lines. 

  can be a ==single line number, a regular expression, or a range of lines==

- command：对匹配行的操作，有POSIX和GUN extensions两种。
- options：command的选项

sed scripts同样有注解，以`#`开头

```
[root@k8snode01 opt]# cat script.sed
#!/bin/sed -f

#this is a comment
/^label/a 3
[root@k8snode01 opt]# cat Dockerfile |  ./script.sed
FROM busybox
label:hello world
3
pwd
```















