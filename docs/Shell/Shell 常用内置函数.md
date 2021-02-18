# Shell 常用内置函数

1. dirname

   https://zh.wikipedia.org/wiki/Dirname
   
   一般用于获取脚本的目录的绝对路径
   
   ```
   root in /usr/local/\/shell_test λ cat dirname.sh 
     File: dirname.sh
     #!/usr/bin/env bash
     echo $(dirname $0)
     bin_path=$(cd $(dirname $0);pwd)
     echo $bin_path                                                       
                                             
   root in /usr/local/\/shell_test λ sh dirname.sh  #如果字符串中没有/,就会输出.
   .
   /usr/local/\/shell_test                    
   ```

2. printf

   与c中的printf一样

   ```
   printf "%-10s %-8s %-4s\n" 姓名 性别 体重kg  
   printf "%-10s %-8s %-4.2f\n" 郭靖 男 66.1234
   printf "%-10s %-8s %-4.2f\n" 杨过 男 48.6543
   printf "%-10s %-8s %-4.2f\n" 郭芙 女 47.9876
   ```

   