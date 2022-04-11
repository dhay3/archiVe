# Linux tr

tr用于从stdin中读取内容删除字符或替换字符，如果没有指定option默认使用translate

syntax：`tr [option] <set1> [set2]`

## sets

sets具体可使用的值查看manual page

## 例子

- 小写变大写

  ```
  cpl in ~ λ echo TAOBAO.COM | tr 'A-Z' 'a-z'
  taobao.com
  
  cpl in ~ λ tr 'A-Z' 'a-z' < TAOBAO.com
  ```

- 替换字符

  但是只有当

  ```
  cpl in ~ λ echo {taobao.com} | tr '{}' '()'
  (taobao.com)
  ```

- 删除字符

  ```
  cpl in ~ λ echo "aaaabbb" | tr -d 'a'                  
  bbb
  ```

- 删除多余字符

  ```
  cpl in ~ λ echo "tao  bao  .          com" | tr -s ' ' 
  tao bao . com
  ```

- 只保留set1内容

  ```
  cpl in ~ λ echo 'aabbcc' | tr -cd 'aa\n'
  aa
  ```

- 用于字符，这里使用`\000`（null）替换`\n`换行符

  ```
  root in ~ λ cat /proc/version 
    File: /proc/version
    Linux version 5.7.0-kali1-amd64 (devel@kali.org) (gcc version 9.3.0 (Debian 9.3.0-14), GNU ld (GNU Binutils for Debian) 2.34) #1 SMP Debian 5.7.6-1kali2 (2020-07-01)
  root in ~ λ cat /proc/version | tr '\000' '\n'
  Linux version 5.7.0-kali1-amd64 (devel@kali.org) (gcc version 9.3.0 (Debian 9.3.0-14), GNU ld (GNU Binutils for Debian) 2.34) #1 SMP Debian 5.7.6-1kali2 (2020-07-01)
  ```

  

