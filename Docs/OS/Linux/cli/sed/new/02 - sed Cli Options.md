# 02 - sed Cli Options

## 0x01 Syntax

`sed` 通常以如下格式使用

```
sed OPTIONS... [SCRIPT] [INPUTFILE...]
```

`INPUTFILE...` 是一个变参

1. 当 `INPUTFILE...` 为空，或者是 `-` 时，`sed` 会从 stdin 中读取内容，用于执行 Action
2. 当有多个 `INPUTFILE` 时，分别对每一个文件执行 Action

## 0x02 Options

- `-n | --quiet | silent`

  `sed` 默认会输出所有的处理过程(除了使用 `d` modified/deleted 的内容)，使用改参数 `sed` 不会输出处理的过程
  
  ```
  
  ```

- `-e <script> | --expression=<script>`

  指定使用的 script，通常在需要多个 script 时使用

  ```
  
  ```

- `-f <script-file> | --file=<script-file>`

  指定一个含有 script 的文件用于

- `--posix`

- `--follow-symlinks`

  `sed` 默认不会对 symbolic links 执行 scripts, 使用该参数可以修改 symbolic links 指向的时间文件中匹配的内容

  ```
  
  ```

- `-E | -r | --regexp-extended`

  使用 extended regular expressions 而不是 basic regular expressions (同 `grep` 中的逻辑)

- `--debug`

  用于 debug，会输出 `sed` 的执行过程

  ```
  (base) cc in /tmp λ echo 1 | sed --debug '\%1%s21232'
  SED PROGRAM:
    /1/ s/1/3/
  INPUT:   'STDIN' line 1
  PATTERN: 1
  COMMAND: /1/ s/1/3/
  MATCHED REGEX REGISTERS
    regex[0] = 0-1 '1'
  PATTERN: 3
  END-OF-CYCLE:
  3
  ```

- `--version`

  输出版本信息

**references**

[^1]:https://www.gnu.org/savannah-checkouts/gnu/sed/manual/sed.html#Command-Line%20Options