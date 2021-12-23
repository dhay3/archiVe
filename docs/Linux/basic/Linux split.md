# Linux split

## disgest

> 注意文件是会覆盖的

syntax：`split [option] [file[prefix]]`

用于将文件分割成多个小文件，默认1000以行分割，以‘x’ 作为prefix。例如分割`/etc/services`

```
➜  Desktop split services.bak 
➜  Desktop ls
 xaa   xae   xai
 xab   xaf   xaj  
 services.bak    xac   xag   xak  
 xad   xah   xal
```

文件会以`x{[a-z][a-z]}`的格式显示，如果没有指定文件会从stdin中读取，` cat /etc/services | split -`

## positonal args

## optional args

- `-b | --bytes`

  以指定byte分割文件，可以使用K,M,G来表示

- `-d | -x`

  suffix使用数字 | hex 替代字母，例如：`xaa`会变成`x01`

- `-l | --lines=NUMBER`

  每 number 行分割一个文件

- `-n | --number`

  生成chuncks file，chunks具体可以查看man page

- `-t | --sparator=SEP`

  区别行与行之间不使用newline，使用SEP

- `--verbose`

- `-e | --elide-empty-files`

  不生成空文件，用于指定chunks数时

## chunks

- N

  将文件分割成N个

  ```
  ➜  test split -n 2 services.bak 
  creating file 'xaa'
  creating file 'xab'
  ```

- K/N

  将分割的第K/N个文件的内容输出到stdout

  ```
  ➜  test split -n 1/2 services.bak | more
  # Full data: /usr/share/iana-etc/port-numbers.iana
  
  tcpmux              1/tcp
  tcpmux              1/udp
  ```

## Example

```
➜  test split -dn 3  services.bak services.split
creating file 'services.split00'
creating file 'services.split01'
creating file 'services.split02'
```

