---
createTime: 2024-07-29 12:00
tags:
  - "#Passwall"
  - "#Clash"
---

# Mihomo 09 - GEOIP GEOSITE

## 0x01 Overview

## 0x02 GEOIP

GeoIP 是一种机制，可以根据 IP address 来判断 Geographic location[^1](其实本体就是一个数据库，通过 IP 来查询对应的信息)

在 Clash 中被用于 rule mode 分流，通过 GEOIP rule 实现

例如

> GEOIP rule 只支持一个参数，即所属国家代码

如果域名解析出来的 IP 是属于 geoip.dat 中 CN 的，就会使用 DIRECT 策略

```
- GEOIP,CN,DIRECT
```

### 0x02a geoip.dat

Mihomo core 默认会使用 [GitHub - MetaCubeX/meta-rules-dat: rules-dat for mihomo](https://github.com/MetaCubeX/meta-rules-dat?tab=readme-ov-file) 中的 geoip.dat 做为数据源(具体代码看 [mihomo/config/config.go at Meta · MetaCubeX/mihomo · GitHub](https://github.com/MetaCubeX/mihomo/blob/Meta/config/config.go))

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

而在 clash verge rev 中的表现为(具体代码看 [clash-verge-rev/scripts/check.mjs at main · clash-verge-rev/clash-verge-rev · GitHub](https://github.com/clash-verge-rev/clash-verge-rev/blob/main/scripts/check.mjs))

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

Mihomo core 用的 geoip.dat 和 [Loyalsoldier/v2ray-rules-dat](https://github.com/Loyalsoldier/v2ray-rules-dat) 相同。实际上会从 [Loyalsoldier/geoip](https://github.com/Loyalsoldier/geoip) 同步(具体可以看 [v2ray-rules-dat/.github/workflows/run.yml at master · Loyalsoldier/v2ray-rules-dat · GitHub](https://github.com/Loyalsoldier/v2ray-rules-dat/blob/master/.github/workflows/run.yml))

```yaml
		...
      - name: Get geoip.dat relative files
        run: |
          wget https://github.com/Loyalsoldier/geoip/raw/release/geoip.dat
          wget https://github.com/Loyalsoldier/geoip/raw/release/geoip.dat.sha256sum
		...
```

所有可用的 GEOIP 参数可以参考 [geoip/text at release · Loyalsoldier/geoip · GitHub](https://github.com/Loyalsoldier/geoip/tree/release/text)(如果有对 IP 地址有疑惑的，可以使用 [GeoIP databases demo | MaxMind](https://www.maxmind.com/en/geoip-demo) 查询)

常用的 GEOIP 参数有(大小写均可)
- private - 私网地址
- cn - 中国的地址段
- us - 美国的地址段

除此外还有几个特殊的地址段
- cloudflare
- cloudfront
- facebook
- fastly
- google
- netflix
- telegram
- twitter

在 Mihomo core 的配置文件中表现为

```
- GEOIP,CN,DIRECT
- GEOIP,US,PROXY
- GEOIP,PRIVATE,DIRECT
```

## 0x03 GEOSITE

GEOSITE 是另外一种机制，根据 Domain 来判断服务

在 Clash 中被用于 rule mode 分流，通过 GEOSITE rule 实现

例如

> GEOSITE rule 只支持一个参数


```
- GEOSITE,category-ads-all,REJECT
```

### 0x02a geosite.dat

Mihomo core 用的 geosite.dat 同样也来自 [GitHub - MetaCubeX/meta-rules-dat: rules-dat for mihomo](https://github.com/MetaCubeX/meta-rules-dat?tab=readme-ov-file)

github workflow 会从其他仓库中同步规则，根据 `./community/data` 中的内容(也是可用参数)生成 geosite.data

```yaml
      - name: Build geosite.dat file
        run: |
          cd custom || exit 1
          go run ./ --datapath=../community/data
```

可用参数如下(大小写均可)

1. cn
2. geolocation
3. category-ads-all
4. china-list
5. google-cn
6. apple-cn
7. gfw
8. greatfire
9. win-spy
10. win-update
11. win-extra

具体可用的参数可以看 workflow

[GitHub - Loyalsoldier/v2ray-rules-dat at release](https://github.com/Loyalsoldier/v2ray-rules-dat/tree/release) 中的 txt 文件名

在 Mihomo core 的配置文件中表现为

```
- GEOSITE,CN,DIRECT
- GEOSITE,CATEGORY-ADS-ALL,REJECT
```

## 0x04 Configuration

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

