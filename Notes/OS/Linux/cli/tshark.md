# tshark

## Digest

tshark 是 wireshark 下衍生的 TUI CLI 抓包工具，对标 tcpdump

```
tshark -i any -O tls  -Y 'tcp.port eq 443' --hexdump all -S '-------------------------'
```