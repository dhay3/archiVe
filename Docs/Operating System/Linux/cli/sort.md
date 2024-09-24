# Linux sort

## Digest

syntax：`sort [options] [file]`

==按照ascii表比较==

==从小到大==

每一行为一个值

如果没有指定 file 或者 file 的位置是 `-`，表示从 stdin 读取

## Optional args

### input options

- `-b | --ignore-leading-blanks`

  忽略开头的空字串

- `-f | --ignore-case`

  忽略大小写

- ==`-t | --field-separator=<SEP>`==

  https://stackoverflow.com/questions/17478511/what-does-the-k-parameter-do-in-the-sort-function-linux-bash-scripting

  行与行之间的分隔符，默认`\t`，可以通过该参数来修改

  ```
  #sort -t ' ' a
  e b
  gg f
  k i
  z a
  
  #cat a
  z a
  gg f
  e b
  k i
  ```

  先按照第一列来排序

- `-k | --key=<keydef>`

  通过指定的key来排序(即为column)，默认为第一列通过`-t`指定的参数来分隔

  ```
  #sort -t ' ' -k 2 a
  z a
  e b
  gg f
  k i
  
  #cat a
  z a
  gg f
  e b
  k i
  ```

  上面的例子表示按照第二列来排序，列与列之间使用空格分隔

- `-d | --dictionary-order`

  只对字母按照ascii表排序

- `-g | --genernal-numeric-sort`

  按照数值比较，支持科学计数

- `-n | --numeric-sort`

  按照数值比较，不支持科学计数

- `-M | --month-sort`

  根据month来对比 `compare (unknown) < 'JAN' < ... < 'DEC'`

  ```
  sort -M a
  JAN
  FEB
  MAR
  ```

- `-h | --humna-numeric-sort`

  根据`K`，`M`，`G`来排序

  ```
  sort -h a
  10K
  1M
  20G
  ```

- `-r | --reverse`

  对输出的内容逆序 (即默认从小到大排序，就会变成从大到小)

### other options

- `--parallel=N`

  并行处理sort

  ```
  #sort --parallel=6  -h a
  10K
  1M
  20G
  100G
  ```

- `-u | --unique`

  等价于`sort <file> | unique` 不输出重复的值

  值是否判断为重复又 key(column) 来决定

- `-z | --zero-terminated`

  以 0 byte 来分隔，而不是以newline来分隔

**references**

[^1]:https://stackoverflow.com/questions/1255782/whats-the-difference-between-general-numeric-sort-and-numeric-sort-options
