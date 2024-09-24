# Shell cd 

1. shell script会在subshell中运行，所以使用`cd`时subshell会进入到那个目录，脚本执行完毕后退出

   ```
   cpl in /tmp λ cat a.sh 
   cd ~
   pwd
   echo PID=$$;ps -ef | grep $$
   cpl in /tmp λ bash a.sh 
   /home/cyberpelican
   PID=15924
   cpl        15924    3580  0 22:16 pts/0    00:00:00 bash a.sh
   cpl        15925   15924  0 22:16 pts/0    00:00:00 ps -ef
   cpl        15926   15924  0 22:16 pts/0    00:00:00 grep 15924
   ```

2. 和正常使用`cd`命令一样，不能将带有模式扩展的路径放(`~/dev`  or `/dev/sn?/`)在引号内。

   ```
   cpl in /opt λ cd "~";cd "/dev/sn?"
   cd: no such file or directory: ~
   cd: no such file or directory: /dev/sn?
   #如果路径不需要模式扩展可以放在引号内fid
   cpl in /opt λ cd "/opt"
   ```

