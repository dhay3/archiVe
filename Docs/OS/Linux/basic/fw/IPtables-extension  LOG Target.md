# IPtables-extension  LOG Target

ref

https://unix.stackexchange.com/questions/423778/log-iptables-events-on-centos-7

## Digest

LOG 模块可以将匹配的报文输出到 kernel log，这样就可以通过 `dmesg` 来查看了。但是 LOG 是 non-terminating target，所以当规则匹配到报文时不会停止，而是会继续匹配其他的规则。如果想要将匹配的报文丢弃就需要配置两条规则一条 LOG 一条 DROP

## Optional args

- `--log-level level`

  记录 logging 的等级，可以是如下的值

  1. emerg
  2. alert
  3. crit
  4. error 缺省值
  5. warning
  6. notice
  7. info
  8. debug

- `--log-prefix prefix`

  在日志前加特定的 prefix

- `--log-tcp-sequence`

  记录 tcp sequence

- `--log-tcp-options`

  记录 tcp options

- `--log-ip-options`

  记录 IP/IPv6 header

- `--log-uid`

  记录 uid

- `--log-macdecode`

  记录 MAC address

## Examples

```
iptables -t filter -A INPUT -j LOG --log-level debug
```

