# IPtables-extensions Modules

ref:

https://www.zsythink.net/archives/1544

syntax

```
ip6tables [-m name [module-options...]]  [-j target-name [target-options...]
iptables  [-m name [module-options...]]  [-j target-name [target-options...]
```

## Digest

iptables 除了使用默认 matches 和 targets 参数外，还可以通过`-m | --match` 来指定使用特定模块，用于扩展 matches 和 targets 的参数

如果使用了`-p`参数指定了协议，iptables 会自动加载对应协议名的 moudle，==这样就可以额外再使用其他的 modules 了==

==module 具体可以使用的参数，可以通过 `iptables -m mode_name -h` 来查看==