# Linux man 

> 如果想要精准匹配，使用正则表达式

具体的command manual page 在Linux上按章节存储，一般在`man command`后显示在左上角，例如`sudo(8)`表似第8章

1.  Executable programs or shell commands
2. System calls (functions provided by the kernel)
3. Library calls (functions within program libraries)
4. Special files (usually found in /dev)
5. File formats and conventions, e.g. /etc/passwd
6. Games
7. Miscellaneous (including  macro  packages  and  conventions),  e.g.
          man(7), groff(7)
8. System administration commands (usually only for root)
9. Kernel routines [Non standard]

## 案例

1. 如果在不同的章节有相同的命令，使用`man  man.7`可以从指定的章节中调用manual page

2. 使用`man -k keyword`从manual page中搜索含有关键字的命令

   ```
   root in / λ man -k awk
   awk (1)              - pattern scanning and processing language
   filefuncs (3am)      - provide some file related functionality to gawk
   gawk (1)             - pattern scanning and processing language
   mawk (1)             - pattern scanning and text processing language
   nawk (1)             - pattern scanning and processing language
   readdir (3am)        - directory input parser for gawk
   rwarray (3am)        - write and read gawk arrays to/from files
   time (3am)           - time functions for gawk            
   ```

3. `/\s+-s`匹配精准匹配`-s`参数



