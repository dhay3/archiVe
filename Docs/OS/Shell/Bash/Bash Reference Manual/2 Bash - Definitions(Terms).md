# 2 - Definitions(Terms)

- POSIX

  应用于 Unix OS 上的一系列标准

  ```
  POSIX = "A family of open system standards"
  ```

- blank

  <kbd>space</kbd> 或者是 <kbd>tab</kbd>

  ```
  blank = "A space or table character"
  ```

- builtin

  Bash 内建的脚本或者是程序

  ```
  builtin = "A command that is implemented internally by the shell itself, rather than by an executable program somewhere in the file system. "
  ```

- control operator

  控制运算符，例如

  `||`, `&&`, `&`, `;`, `;;`, `;;&`, `|`, `|&`, `(` , `)`

  ```
  control operator = "||" | "&&" | "&" | ";" | ";;" | ";;&" | "|" | "|&" | "(" | ")"
  ```

- exit status(return status)

  The value returned by a command to its caller

  ```
  exit status = 1 - 255
  return status = "exit status"
  ```

  最大只有 8 bit 所以十进制最大值 255

- field

  shell 扩展后的结果
  
  ```
  field = "A unit of text that is the result of one of the shell expansions"
  ```

- filename

  文件名

  ```
  filename = "A string of characters used to identify a file"
  ```

- job

  一组进程

  ```
  job = "A unit of text that is the result of one of the shell expansions"
  ```

- job control

  ```
  job control = "A mechanism by which users can selectively stop (suspend) and restart (resume) execution of processes."
  ```

- metacharacter

  ```
  metacharacter = "space" | "tab" | "newline" | "|" | "&" | ";" | "(" | ")" | "<" | ">"
  ```

  metacharacter 不在引号内

- name

  用于定义变量和函数名

  ```
  name = "A word consisting solely of letters, numbers, and underscores, and beginning with a letter or underscore"
  ```

- operator

  由 control operator 和 redirection operator 组成，至少包含一个 metacharacter

  ```
  operator = "control operator" | "redirection operator"
  ```

- process group

- process group ID

- reserved word

  shell 保留的词

  ```
  reserved word = "for" | "while" | ...
  ```

- signal

  ```
  singal = "A mechanism by which a process may be notified by the kernel of an event occurring in the system"
  ```

- special builtin

  ```
  special builtion = "A shell builtin command that has been classified as special by the POSIX standard."
  ```

- token

  word 和 operator 的统称

  ```
  token = word | operator
  ```

- word

  词，不包含 metacharacter 中的字符

  ```
  word = "A sequence of characters treated as a unit by the shell. Words may not include unquoted metacharacters."
  ```

**references**

1. [^1]:https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#Definitions