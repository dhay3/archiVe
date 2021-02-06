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
   
   