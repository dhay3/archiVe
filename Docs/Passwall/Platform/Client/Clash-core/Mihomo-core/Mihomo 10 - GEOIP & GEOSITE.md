---
createTime: 2024-07-29 12:00
tags:
  - "#Passwall"
  - "#Clash"
---

# Mihomo 10 - GEOIP & GEOSITE

## 0x01 Overview

> [!important]
> 针对 Mihomo core 的逻辑同样适用于 Clash Core

Internet geolocation [^1]是一种机制，可以通过一些因子(IP,TLD,Domain)来判断设备接入互联网的 geographci location

GEOIP 和 GEOSITE 分别是 Internet geolocation 的 2 种实现

- GEOIP 根据 IP 来判断 geographic location
- GEOSITE 根据 TLD/Domain 来判断 geographic location

在 Mihomo 中 GEOIP 和 GEOSITE 被用于 rule mode 分流

例如

```
- GEOIP,CN,DIRECT
- GEOSITE,CN,DIRECT 
```

- 当因子是 IP 时，Mihomo 会查询 geoip.dat 来判断 geographic location，即 GEOIP
- 当因子是 TLD/Domain 时，Mihomo 会查询 geosite.dat 来判断 geographic location，即 GEOSITE

## 0x02 V2ray GEOIP/GEOSITE

> [!NOTE] 
> 这里不分析 V2ray-core 的代码

Mihomo 查询 geoip.dat 和 geosite.dat 的这种机制继承自 V2ray-core[^2]

例如 

v2ray core 配置中的 `geoip:` 和 `geosite:` 指令分别查询 geoip.dat 和 geosite.dat

```json
      "ip": [
           "0.0.0.0/8",
           "10.0.0.0/8",
           "fc00::/7",
           "fe80::/10",
           "geoip:cn",
           "geoip:!cn",
           "ext:customizedGeoIPFile.dat:cn",
           "ext:customizedGeoIPFile.dat:!cn",
           "ext-ip:customizedGeoIPFile.dat:cn",
           "ext-ip:customizedGeoIPFile.dat:!cn"
       ],
       "domains": [
           "baidu.com",
           "qq.com",
           "geosite:cn",
           "ext:customizedGeoSiteFile.dat:cn"
       ],
   ```

### 0x02a geoip.dat

V2ray-core 默认会使用 [GitHub - v2fly/geoip: GeoIP for V2Ray. Generate and release GeoIP files such as geoip.dat and geoip-only-cn-private.dat for V2Ray automatically every Thursday.](https://github.com/v2fly/geoip) 生成的 release 作为 geoip.dat。通过 `geoip:` 指令来查询 geoip.dat

`geoip:` 有且只有一个参数(大小写均可)，通常是国家代码[^3]

例如常用的有[^2]
 - cn - 中国的地址段
 - us - 美国的地址段
 
 除此外还有一个特殊地址段
 - private - 私网地址

`geoip:` 所有可用参数可以看 [geoip/text at release · v2fly/geoip · GitHub](https://github.com/v2fly/geoip/tree/release/text) 即文件名

### 0x02b geosite.dat

V2ray-core 默认会使用 [GitHub - v2fly/domain-list-community: Community managed domain list. Generate geosite.dat for V2Ray.](https://github.com/v2fly/domain-list-community) 生成的 release 作为 geosite.dat。通过  `geosite:` 指令来查询 geosite.dat

`geosite:` 有且只有一个参数(大小写均可)，通常是域名类别[^2]

例如常用的有[^4]
- category-ads - 包含常见的广告域名
- category-ads-all - 包含常见的广告域名，以及广告提供商域名
- tld-cn - 包含了 CNNIC 管理的用于中国大陆的顶级域名，如以 `.cn`、`.中国` 结尾的域名
- tld-!cn - 包含了非中国大陆使用的顶级域名，如以 `.hk`（香港）、`.tw`（台湾）、`.jp`（日本）、`.sg`（新加坡）、`.us`（美国）`.ca`（加拿大）等结尾的域名
- geolocation-cn - 包含了常见的大陆站点域名
- geolocation-!cn - 包含了常见的非大陆站点域名，同时包含了 `tld-!cn`
- cn - 相当于 `geolocation-cn` 和 `tld-cn` 的合集
- apple - 包含了 Apple 旗下绝大部分域名。
- google - 包含了 Google 旗下绝大部分域名。
- microsoft - 包含了 Microsoft 旗下绝大部分域名。
- facebook - 包含了 Facebook 旗下绝大部分域名。
- twitter - 包含了 Twitter 旗下绝大部分域名。
- telegram - 包含了 Telegram 旗下绝大部分域名。

`geosite` 所有可用参数可以看 [domain-list-community/data at master · v2fly/domain-list-community · GitHub](https://github.com/v2fly/domain-list-community/tree/master/data) 即文件名

但是需要注意的一点是，在类似 steam 或者是 google (按照地区提供不同服务的服务商)的规则中可能会有类似  `@cn` 的规则

```
steamchina.com @cn
google.cn @cn
```

在 V2ray 中被称为 attributes，通过 attributes 可以过滤子域名，例如

```
geosite:steam@cn
```

就会过滤 steam 规则中，含有 `@cn` 的所有域名，具体说明可以看 [GitHub - v2fly/domain-list-community: Community managed domain list. Generate geosite.dat for V2Ray.](https://github.com/v2fly/domain-list-community?tab=readme-ov-file#attributes)

## 0x03 Mihomo GEOIP/GEOSITE












Mihomo core 通过 GEOIP 和 GEOSITE 指令分别查询 geoip.dat 和 geosite.dat

```yaml
- GEOIP,PRIVATE,DIRECT
- GEOIP,CN,DIRECT
- GEOIP,US,PROXY
- GEOSITE,category-ads-all,REJECT
```

Mihomo core 和 V2ray core 的 geoip 和 geosite 使用的逻辑上大体相同，但是 2 者的  geoip.dat 和 geosite.dat 绝对不能混用，后面会分析

Mihomo core 默认会从 [GitHub - MetaCubeX/meta-rules-dat: rules-dat for mihomo](https://github.com/MetaCubeX/meta-rules-dat?tab=readme-ov-file) 同步 geoip.dat 和 geosite.dat(具体代码看 [mihomo/config/config.go at Meta · MetaCubeX/mihomo · GitHub](https://github.com/MetaCubeX/mihomo/blob/Meta/config/config.go))

```go
		...
		GeoXUrl: GeoXUrl{
			Mmdb:    "https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/geoip.metadb",
			ASN:     "https://github.com/xishang0128/geoip/releases/download/latest/GeoLite2-ASN.mmdb",
			GeoIp:   "https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/geoip.dat",
			GeoSite: "https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/geosite.dat",
		}
		...
```



但是在 GUI client 中同步方式可能不一样。但是在 Clash Verge Rev 会同样使用 MetaCubeX/meta-rules-dat 中的 geoip.dat 和 geosite.dat(具体代码看 [clash-verge-rev/scripts/check.mjs at main · clash-verge-rev/clash-verge-rev · GitHub](https://github.com/clash-verge-rev/clash-verge-rev/blob/main/scripts/check.mjs))

```ts
const resolveMmdb = () =>
  resolveResource({
    file: "Country.mmdb",
    downloadURL: `https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/country.mmdb`,
  });
const resolveGeosite = () =>
  resolveResource({
    file: "geosite.dat",
    downloadURL: `https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/geosite.dat`,
  });
const resolveGeoIP = () =>
  resolveResource({
    file: "geoip.dat",
    downloadURL: `https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/geoip.dat`,
  });
```

### 0x03a geoip.dat

Mihomo core 使用的 geoip.dat 虽然来自 MetaCubeX/meta-rules-dat 但是实际来自 [Loyalsoldier/v2ray-rules-dat](https://github.com/Loyalsoldier/v2ray-rules-dat) (README 中已经声明) 

具体逻辑可以看 [meta-rules-dat/.github/workflows/run.yml at master · MetaCubeX/meta-rules-dat · GitHub](https://github.com/MetaCubeX/meta-rules-dat/blob/master/.github/workflows/run.yml)

```yaml
	...
      - name: Checkout MetaCubeX/meta-rules-converter
        uses: actions/checkout@v4
        with:
          repository: MetaCubeX/meta-rules-converter
          path: convert
    ...
    # 从 Loyalsoldier/geoip 获取 v2ray geoip.dat
      - name: Get geoip.dat relative files
        run: |
	      ...
          wget https://github.com/Loyalsoldier/geoip/raw/release/geoip.dat
    ...  
	# 将 geoip.data 发送到 publish 等待发送到 github release
      - name: Move and zip files
        run: |
          mkdir -p ./publish/
          ...
          install -Dp ./geoip.dat ./publish/
          ...
    ...   
    # 将 publish 目录下的内容发送到 github release  
      - name: Create and Upload Release
        id: upload_release
        uses: svenstaro/upload-release-action@v2
        with:
          repo_token: ${{ secrets.GITHUB_TOKEN }}
          release_name: Release ${{ env.BUILDTIME }}
          tag: latest
          file_glob: true
          overwrite: true
          file: ./publish/*
    ...
```

而 Loyalsoldier/v2ray-rules-data 实际上会从 [Loyalsoldier/geoip](https://github.com/Loyalsoldier/geoip) 同步(具体可以看 [v2ray-rules-dat/.github/workflows/run.yml at master · Loyalsoldier/v2ray-rules-dat · GitHub](https://github.com/Loyalsoldier/v2ray-rules-dat/blob/master/.github/workflows/run.yml))

```yaml
		...
      - name: Get geoip.dat relative files
        run: |
          wget https://github.com/Loyalsoldier/geoip/raw/release/geoip.dat
          wget https://github.com/Loyalsoldier/geoip/raw/release/geoip.dat.sha256sum
		...
```

Mihomo 中通过 `GEOIP` 指令来查询 geoip.dat

例如

```
- GEOIP,CN,DIRECT
- GEOIP,US,PROXY
- GEOIP,PRIVATE,DIRECT
```

`GEIP` 有且只有一个参数，在 V2ray core [0x02a geoip.dat](#0x02a%20geoip.dat) 的基础上增加了

- cloudflare
- cloudfront
- facebook
- fastly
- google
- netflix
- telegram
- twitter

在上面分析 geoip.dat 生成的过程中，可以知道会将 metadb 推送到 meta 分支，所以所有 GEOIP 可用的参数可以参考 [meta-rules-dat/geo/geoip at meta · MetaCubeX/meta-rules-dat · GitHub](https://github.com/MetaCubeX/meta-rules-dat/tree/meta/geo/geoip) 中的文件名

### 0x03b geosite.dat

同样的 Mihomo core 使用的 geosite.dat 虽然来自 MetaCubeX/meta-rules-dat 但是实际来自由多个源组成(README 中已经声明) 

> - `geosite:category-ads-all` 仅使用域名作为广告拦截用途作用有限，因此不作额外域名添加
> - `geosite:cn` 源替换为 [ios_rule_script/ChinaMax_Domain](https://github.com/blackmatrix7/ios_rule_script/tree/master/rule/Clash/ChinaMax)
> - `geosite:onedrive` 合并 [ios_rule_script/OneDrive](https://github.com/blackmatrix7/ios_rule_script/tree/master/rule/Clash/OneDrive)
> - `geosite:steam@cn` 合并 [ios_rule_script/SteamCN](https://github.com/blackmatrix7/ios_rule_script/tree/master/rule/Clash/SteamCN) 的内数据
> - 新增类别 - `geosite:biliintl` 来源 [biliintl](https://raw.githubusercontent.com/xishang0128/rules/main/biliintl.list) - `geosite:tracker` 来源 [TrackersList](https://trackerslist.com/#/zh)以及[blackmatrix7 / ios_rule_script](https://github.com/blackmatrix7/ios_rule_script/tree/master/rule/Clash/PrivateTracker)

具体可以看 [meta-rules-dat/.github/workflows/run.yml at master · MetaCubeX/meta-rules-dat · GitHub](https://github.com/MetaCubeX/meta-rules-dat/blob/master/.github/workflows/run.yml)

```yaml
	...
      - name: Checkout MetaCubeX/meta-rules-converter
        uses: actions/checkout@v4
        with:
          repository: MetaCubeX/meta-rules-converter
          path: convert
    ...
	# 将 Loyalsoldier/domain-list-custom 的内容放到 custom 目录
      - name: Checkout Loyalsoldier/domain-list-custom
        uses: actions/checkout@v4
        with:
          repository: Loyalsoldier/domain-list-custom
          path: custom
    ...
    # go run ./ 会运行 https://github.com/Loyalsoldier/domain-list-custom/blob/master/main.go 生成 geosite.dat
    # ../community/data 中数据来源多个上游，具体看 workflow
      - name: Build geosite.dat file
        run: |
          cd custom || exit 1
          # echo sentry.io >>  ../community/data/openai
          echo ipleak.net >> ../community/data/geolocation-\!cn && echo browserleaks.org >> ../community/data/geolocation-\!cn
          go run ./ --datapath=../community/data
    ...
	# 将 gepssite.data 发送到 publish 等待发送到 github release
      - name: Move and zip files
        run: |
          mkdir -p ./publish/
          ...
          install -Dp ./custom/publish/geosite.dat ./publish/
          ...
    ...     
    # 将 publish 目录下的内容发送到 github release  
      - name: Create and Upload Release
        id: upload_release
        uses: svenstaro/upload-release-action@v2
        with:
          repo_token: ${{ secrets.GITHUB_TOKEN }}
          release_name: Release ${{ env.BUILDTIME }}
          tag: latest
          file_glob: true
          overwrite: true
          file: ./publish/*
    ...
```

Mihomo 中通过 `GEOSITE` 指令来查询 geosite.dat

例如

```
- GEOSITE,CN,DIRECT
- GEOSITE,CATEGORY-ADS-ALL,REJECT
```

`GEIP` 有且只有一个参数，在 V2ray core [0x03b geosite.dat](#0x03b%20geosite.dat) 的基础上增加了(使用上和 Loyalsoldier/v2ray-rules-dat 相同) 

- win-spy
- win-update
- win-extra

GEOSITE 所有可用的参数可以参考 [meta-rules-dat/geo/geosite at meta · MetaCubeX/meta-rules-dat · GitHub](https://github.com/MetaCubeX/meta-rules-dat/tree/meta/geo/geosite) 中的文件名

### 0x03c dat db metadb[^5]

MetaCubeX/meta-rules-dat 中关于 geoip geosite 有 3 种格式的文件，但是文档里并没有介绍几者之间的区别。如果我们要想搞清楚区别，就需要从 Mihomo 源码入手






除了 geoip.dat 和 geosite.dat 你在 MetaCubeX/meta-rules-dat 上还会看到 geoip.db geoip.metadb geosite.db geosite.meta

实际上 db 结尾的文件是用于 singbox 的，而 meta 结尾的文件是用于 meta 的

这点可以从 workflows 中大概推测出来

```yaml
      - name: Build db and metadb file
        env:
          NO_SKIP: true
        run: |
          go install -trimpath -ldflags="-s -w -buildid=" github.com/metacubex/geo/cmd/geo@master
          geo convert site -i v2ray -o sing -f geosite.db ./custom/publish/geosite.dat
          geo convert site -i v2ray -o sing -f geosite-lite.db ./community/geosite-lite.dat
          geo convert ip -i v2ray -o sing -f geoip.db ./geoip.dat
          geo convert ip -i v2ray -o meta -f geoip.metadb ./geoip.dat
          geo convert ip -i v2ray -o sing -f geoip-lite.db ./geoip-lite.dat
          geo convert ip -i v2ray -o meta -f geoip-lite.metadb ./geoip-lite.dat
```

也可以在官网得到证实 [GeoIP - sing-box](https://sing-box.sagernet.org/configuration/route/geoip/#download_detour)




这里需要注意的一点是 Mihomo geoip.dat 和 V2ray geoip.dat 不能混用，上面的 worklflow 也已经说明了，需要经过 convert 转换才可以

对比 V2ray core geoip metadb 和 Mihomo core geoip metadb 也可以证明

[V2ray core geoip metadb private.txt](https://github.com/v2fly/geoip/blob/release/text/private.txt)

```
0.0.0.0/8
10.0.0.0/8
100.64.0.0/10
127.0.0.0/8
169.254.0.0/16
172.16.0.0/12
192.0.0.0/24
192.0.2.0/24
192.88.99.0/24
192.168.0.0/16
198.18.0.0/15
198.51.100.0/24
203.0.113.0/24
224.0.0.0/3
::/127
fc00::/7
fe80::/10
ff00::/8
```

[Clash core geoip metadb private](https://github.com/MetaCubeX/meta-rules-dat/blob/meta/geo/geoip/private.yaml)

> Mihomo 实际上支持多种格式 txt、mrs、txt

```yaml
payload:
    - 0.0.0.0/8
    - 10.0.0.0/8
    - 100.64.0.0/10
    - 127.0.0.0/8
    - 169.254.0.0/16
    - 172.16.0.0/12
    - 192.0.0.0/24
    - 192.0.2.0/24
    - 192.88.99.0/24
    - 192.168.0.0/16
    - 198.18.0.0/15
    - 198.51.100.0/24
    - 203.0.113.0/24
    - 224.0.0.0/3
    - ::/127
    - fc00::/7
    - fe80::/10
    - ff00::/8
```


```yaml
    # 使用 MetaCubeX/meta-rules-converter 中的 main.go 将 geoip.dat 转为 ../meta-rule/geo/geoip 
    # commandIP.PersistentFlags().StringVarP(&inPath, "file", "f", "", "Input File Path")
	# commandIP.PersistentFlags().StringVarP(&outType, "type", "t", "", "Output Type")
	# commandIP.PersistentFlags().StringVarP(&outDir, "out", "o", "", "Output Path")
      - name: Convert geo to meta-rule-set
        env:
          NO_SKIP: true
        run: |
          mkdir -p meta-rule/geo/geosite && mkdir -p meta-rule/geo/geoip
          cd convert
          ...
          go run ./ geoip -f ../geoip.dat -o ../meta-rule/geo/geoip
	...
```

```yaml
    # 将原 v2ray 格式的 geoip.dat 转为 meta(mihomo) 格式的 metadb
    # geo convert ip -i <input_type> -o <output_type> -f [output_filename] input_filename
      - name: Build db and metadb file
        env:
          NO_SKIP: true
        run: |
          go install -trimpath -ldflags="-s -w -buildid=" github.com/metacubex/geo/cmd/geo@master
          ...
          geo convert ip -i v2ray -o meta -f geoip.metadb ./geoip.dat
          ...
    ...
```

## 0x04 Configuration

geo 的配置很简单只有几个

```yaml
geodata-loader: standard
geo-auto-update: true
geo-update-interval: 168
geox-url:
  geoip: "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@release/geoip.dat"
  geosite: "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@release/geosite.dat"
  mmdb: "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@release/country.mmdb"
  asn: "https://github.com/xishang0128/geoip/releases/download/latest/GeoLite2-ASN.mmdb"
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Internet geolocation - Wikipedia](https://en.wikipedia.org/wiki/Internet_geolocation)
[^2]:[Routing 路由 | V2Fly.org](https://www.v2fly.org/config/routing.html#ruleobject)
[^3]:[國家地區代碼 - 维基百科，自由的百科全书](https://zh.wikipedia.org/wiki/%E5%9C%8B%E5%AE%B6%E5%9C%B0%E5%8D%80%E4%BB%A3%E7%A2%BC)
[^4]:[Routing 路由 | V2Fly.org](https://www.v2fly.org/config/routing.html#%E9%A2%84%E5%AE%9A%E4%B9%89%E5%9F%9F%E5%90%8D%E5%88%97%E8%A1%A8)
[^5]:[对各种geo文件以及格式的疑问 · Issue #44 · MetaCubeX/meta-rules-dat · GitHub](https://github.com/MetaCubeX/meta-rules-dat/issues/44)