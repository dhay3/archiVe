# OpenWrt LuCI

ref

https://openwrt.org/docs/guide-user/luci/luci.essentials#installation

LuCI 是 openWrt 的一个 WebUI 管理工具，OpenWrt generic 默认会安装 LuCI

```
➜  /etc opkg list-installed luci
luci - git-20.074.84698-ead5e81
```

## TLS over LuCI

1. 证书自签

   https://openwrt.org/docs/guide-user/luci/getting_rid_of_luci_https_certificate_warnings#create_install

   文档中第 7 步上传 TLS 证书在 Administration HTTP(S) Access 部分，也可以修改 `/etc/config/uhttpd` 配置中的 `option cert` 和 `option key` 到前面生成的 digital certificate 和 private key  

2. 自签证书导入系统根证书

   可以按照文档中的操作（ 部分 chrome 版本在 privacy tab 中找有关 certificate 管理的选项），也可以直接点击自签证书安装到 受信任的根证书颁发机构

3. 校验

   修改本地的 hosts 文件或者 DNS 服务器把 192.168.2.100 指向 luci.openwrt

## LuCI themes

opkg 主题

https://openwrt.org/docs/guide-user/luci/luci.themes

开源主题 github 关键字 luci-theme, 如果不需要修改样式直接安装即可，反之需要编译

- Argon

  https://github.com/jerrykuku/luci-theme-argon

- Neobird

  https://github.com/thinktip/luci-theme-neobird

- Infinity Freedom

  https://github.com/xiaoqingfengATGH/luci-theme-infinityfreedom

## LuCI statistics

https://openwrt.org/docs/guide-user/luci/luci_app_statistics

