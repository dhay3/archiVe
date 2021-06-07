# Linux 创建隐藏文件

参考：

https://blog.csdn.net/lswzw/article/details/101457387

- ```
  root in /opt λ mkdir .hide                                                   /0.0s
  root in /opt λ lla
  drwxr-xr-x root root     4 KB Wed Jan 27 07:20:04 2021  .
  drwxr-xr-x root root     4 KB Tue Dec 22 23:47:37 2020  ..
  drwxr-xr-x root root     4 KB Wed Jan 27 07:20:04 2021  .hide
  .rw-r--r-- root root   1.3 MB Wed Sep  9 00:12:37 2020  bak.xml
  drwxrwxrwx root root     4 KB Sun Sep 13 23:45:24 2020  burpsuite pro
  drwx--x--x root root     4 KB Sat Sep 19 22:45:25 2020  containerd
  .rw-r--r-- root root   1.3 MB Wed Sep  9 00:12:37 2020  ibak.xml
  ```

- 转义高级用法

  在linux中`\`表示转义，这里实际创建了两个目录`\`和`hide`，第一个`\`做为转译符。如果`cd \`是进不去的应为在命令行中表示拼接命令。 

  ```
  root in /opt λ mkdir \\/hide -p                                              /0.0s
  root in /opt λ ll
  drwxr-xr-x root root     4 KB Wed Jan 27 07:21:12 2021  \
  .rw-r--r-- root root   1.3 MB Wed Sep  9 00:12:37 2020  bak.xml
  drwxrwxrwx root root     4 KB Sun Sep 13 23:45:24 2020  burpsuite pro
  drwx--x--x root root     4 KB Sat Sep 19 22:45:25 2020  containerd
  .rw-r--r-- root root   1.3 MB Wed Sep  9 00:12:37 2020  ibak.xml
  drwxr-xr-x root root     4 KB Thu Sep 10 07:28:47 2020  jdk-14.0.2
  .rwxr--r-- root root 182.1 MB Thu Sep 10 07:28:10 2020  jdk-14.0.2_linux-x64_bin.tar.gz
  .rw-r--r-- root root 661.5 KB Sat Dec 12 09:07:29 2020  lsd_0.18.0_amd64.deb
  .rw-r--r-- root root  1000 GB Wed Jan 27 07:08:12 2021  t1
  .rw-r--r-- root root    35 B  Thu Jan 21 06:43:05 2021  test.sh             /0.0s
  
  root in /opt λ cd \\    
  
  root in /opt/\ λ ll
  drwxr-xr-x root root 4 KB Wed Jan 27 07:21:12 2021  hide  
  
  root in /opt/\ λ cd ..                                                       /0.0s
  root in /opt λ cd \
  > 
  ```

  
