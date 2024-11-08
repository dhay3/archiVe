# examples

## 0x001 在行尾添加

```
sed -n '/nameserver/p' /etc/resolv.conf  | sed 's/$/ ok/'
nameserver 10.195.29.1 ok
nameserver 10.195.29.17 ok
nameserver 10.195.29.33 ok
```

```
cat /etc/resolv.conf | sed '/nameserver/s/$/ ok/'
search tbsite.net aliyun.com
options timeout:2 attempts:2
nameserver 10.195.29.1 ok
nameserver 10.195.29.17 ok
nameserver 10.195.29.33 ok
```

## 0x002 行首添加

快速注释

```
➜  ISO cat /tmp/a
this a comment
➜  ISO sed '/this/s/^/#/' /tmp/a
#this a comment
```

## 0x003 组合行

N会添加一个换行符，使用regex将换行符替换成空字符串

```
seq 4 | sed  '2{N;s/\n//}'
1
23
4
```
