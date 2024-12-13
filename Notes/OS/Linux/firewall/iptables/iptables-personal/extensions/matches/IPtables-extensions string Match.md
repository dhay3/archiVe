# IPtables-extensions string Match

## Digest

> 需要注意的一点是如果用了 TLS，因为报文加密无法过滤 URL_PATH 但是可以过滤域名

iptables 有一个模块用于批准报文的 ascii 内容，一般用于过滤 URL

## Optional args

- `--algo {bm|kmp}`

  选择字符串的匹配模式，必须指定

- `[!] --string pattern`

  matches the given pattern

- `--icase`

  Ignore case when searching

## Examples

例如禁止访问 baidu

```
iptables -t filter -A OUTPUT -m string --string baidu.com --algo bm -j DROP
```

