# Linux date

## Digest

syntax：`date [OPTION]...[+FORMAT]`

用来展示 OS 当前的时间

## Optional args

- `-r | --reference=FILE`

  输出指定文件的 mtime

  ```
  ➜  /tmp date -r /etc/resolv.conf
  Thu Jul 28 02:55:01 PM HKT 2022
  ➜  /tmp ll /etc/resolv.conf
  .rw-r--r-- root root 71 B Thu Jul 28 14:55:01 2022  /etc/resolv.conf
  ➜  /tmp stat /etc/resolv.conf
    File: /etc/resolv.conf
    Size: 71              Blocks: 8          IO Block: 4096   regular file
  Device: 259,7   Inode: 12848283    Links: 1
  Access: (0644/-rw-r--r--)  Uid: (    0/    root)   Gid: (    0/    root)
  Access: 2022-07-28 14:55:01.498020811 +0800
  Modify: 2022-07-28 14:55:01.498020811 +0800
  Change: 2022-07-28 14:55:01.498020811 +0800
   Birth: 2022-07-28 14:55:01.498020811 +0800
  ```

- `-s | --set=STRING`

  以 string 的格式设置当前的时间

### output format

- `-I | --iso-8601`

  以 ISO 8601 的格式输出时间，FMT 的值按照输出的精度可以是 date（缺省值）, hours, mintues, seconds or ns (最精确)

  ```
  ➜  /tmp date --iso-8601=ns
  2022-07-28T15:25:54,879368887+08:00
  ```

- `--rfc-3339=FMT`

  以 RFC 3339 的格式输出，FMT 的值按照输出的精度可以是 date（缺省值）, hours, mintues, seconds or ns (最精确)

  ```
  ➜  /tmp date --rfc-3339=ns     
  2022-07-28 15:25:32.010522194+08:00
  ```

  ## FORMAT

- `%a`

  abbrv weekday

  ```
  #date +%a
  Thu
  ```

- `%b`

  abbrv month

  ```
  #date +%b
  Jan
  ```

- `%d`

  day of month

- `%D`

  date

  ```
  #date +%D
  01/27/22
  ```

- `%H`，`%I`

  24小时制，12小时制

- `%Y`、`%m`

  month

- `%M`

  minute

- `%R`、`%T`

  24-hour hour and minute; same as %H:%M

  time; same as %H:%M:%S

  ```
  #date +%R
  11:17
  
  #date +%T
  11:18:22
  ```

## Examples

```
#date +"%F %T"
2022-01-27 11:38:08
```

