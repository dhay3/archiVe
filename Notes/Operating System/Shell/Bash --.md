# Bash --

A `--` signals the end of options and disables further option processing. Any arguments after the `--` are treated as filenames and arguments.

使用`--`表示所有的opotions结束

```
root in /opt λ ls
alibabacloud  b-v  containerd  Dockerfile  lsd-0.18.0-x86_64-unknown-linux-gnu  tput_t.sh
root in /opt λ ll | grep -- -v
-rw-r--r-- 1 root root    0 Feb 18 18:20 b-v
```

**references**

1. [^1]:https://www.gnu.org/software/bash/manual/bash.html#Invoking%20Bash

2. [^2]:https://unix.stackexchange.com/questions/11376/what-does-double-dash-mean
