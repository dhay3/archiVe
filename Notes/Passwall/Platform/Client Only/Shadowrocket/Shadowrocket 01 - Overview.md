---
createTime: 2024-08-02 15:14
tags:
  - "#hash1"
  - "#hash2"
---

# Shadowrocket 01 - Overview

## 0x01 Overview

> [!important] 
> Shadowrocket 目前没有官网 [x.com](https://x.com/ShadowrocketApp)

Shadowrocket 是 IOS 上的一个代理工具(也支持 Mac，但是更推荐使用 Clash)，因为没有官网也就没有对应的配置文件说明。好在 Shadowrocket 默认会提供一个 default.cnf 配置文件，也有类似 [Johnshall/Shadowrocket-ADBlock-Rules-Forever](https://github.com/Johnshall/Shadowrocket-ADBlock-Rules-Forever/blob/release/lazy.conf) 中提供了模板

基础配置可以参考

[confs/shadowrocket/shadowray.conf at main · dhay3/confs · GitHub](https://github.com/dhay3/confs/blob/main/shadowrocket/shadowray.conf)

## 0x02 Modules

上面的配置并不能去除类似 开屏广告以及 Youtube 中插入的广告，如果想要实现更多的功能，需要借助 Shadowrocket 的 Script 以及 Rewrite 功能，可以通过模块导入

> [!NOTE]
> 推荐使用 [GitHub - deezertidal/shadowrocket-rules: 小火箭 shadowrocket 配置文件 模块 脚本 module sgmodule 图文教程 规则 分流 破解 解锁](https://github.com/deezertidal/shadowrocket-rules) 中定义的模块，如果使用模块看 README
> Github 上只记录了一部分模块，所有可用的模块的参考
> https://whatshub.top/shadowrocket

常用的模块

1. https://whatshub.top/module/adultraplus.module
	去开屏广告
2. https://whatshub.top/module/bili.module
	bilibili 解锁大会员清晰度
3. https://whatshub.top/module/YouTubeAd.sgmodule
	youtube 去广告，画中画
4. https://whatshub.top/module/spotifyVIP.module
	spotify 解锁 premium

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

