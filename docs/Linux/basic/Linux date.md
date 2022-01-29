# Linux date

## digest

用于暂时当前系统的时间

## optional

- `-s`

  设置时间，推荐使用ntp而非手动设置

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

  ```
  #date +%R
  11:17
  
  #date +%T
  11:18:22
  ```

## example

```
#date +"%F %T"
2022-01-27 11:38:08
```

