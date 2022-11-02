# IPtables-extensions MASQUERADE Target

## Digest

只在 nat table 中的 POSTROUTING chain 中生效

用于 dynmically assigned IP connections 动态伪装源 IP

 如果需要指定静态的 IP 可以使用 SNAT target

## Optional args

- `-to-ports port[-port]`

  指定伪装的源端口，只能在 tcp, udp, dccp, sctp 模块中被使用

- `--random`

  随机分配源端口

- `--random-fully`

  随机分配源端口

## Examples