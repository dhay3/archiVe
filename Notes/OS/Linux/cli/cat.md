# Linux cat

ref:

https://blog.csdn.net/zongshi1992/article/details/71693045

https://www.gnu.org/software/coreutils/manual/html_node/cat-invocation.html#cat-invocation

## Digest

syntax：`cat [OPTION]... [FILE]...`

Concatenate FILE to stdout （==直到读取到文件的 EOF 时停止==）

with no FILE, or when FILE is -, read stdin

可以使用`-`表示从stdin读取内容

## Optional args

- `-E | --show-ends`

  在每行末尾添加`$`，==用来判断末尾是否有空格(有些格式的配置文件因为空格会导致语法出错例如 yaml)==

- `-n | --number`

  number all output lines，每行显示序号。==对一些文本处理工具非常有用，例如sed==

- `-s | --squeeze-blank`

  如果某行为空格，cat 会忽略它

- `-T | --show-tabs`

  display TAB characters as `^I`

## With Here document

执行脚本时，如果需要往文件中写入 N 行内容，如果使用 echo 追加的方式，效率极低。这时候就可以使用 shell redirection 中的 here document

> here document 中的 word 部分可以替换成任意单词或字符。

下面是在 zsh 的环境中运行的，如果是在 bash 的环境下运行 prompt 部分不会显示 heredoc

```
#这种方式会覆盖testcat
root in /opt λcat > testcat <<EOF
heredoc> catEOF
heredoc> EOF
root in /opt λcat testcat 
  File: testcat
  catEOF

#这种方式会追加
root in /opt λ cat >> testcat << EOF     
heredoc> testEOF
heredoc> EOF
root in /opt λ cat testcat 
  File: testcat
  catEOF
  testEOF
```

EOF的位置可以调换

```
cat << EOF >> testcat 
>...
>EOF
```

## Examples

这里将`br_netfilter`写入到管道里然后通过tee到`/etc/modules-load.d/k8s.conf`

```
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF
```

