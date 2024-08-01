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

Mihomo GEOIP GEOSITE 的逻辑继承自 V2ray-core[^2]

例如 

v2ray core 配置中指定 `geoip:` 或者 `geosite:` 指令，会分别查询 geoip.dat 和 geosite.dat。当有因子命中时，会打上指定的 Tag，然后按照 Tag 执行指定的出站策略

```json
{
    "domainMatcher": "mph",
    "type": "field",
    "domains": [
        "geosite:cn"
    ],
    "ip": [
        "geoip:cn",
    ],
    "port": "53,443,1000-2000",
    "sourcePort": "53,443,1000-2000",
    "network": "tcp",
    "user": [
        "love@v2ray.com"
    ],
    "inboundTag": [
        "tag-vmess"
    ],
    "protocol": [
        "http",
        "tls",
        "bittorrent"
    ],
    "attrs": "attrs[':method'] == 'GET'",
    "outboundTag": "direct",
    "balancerTag": "balancer"
}
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

而在 [Mihomo 09 - Configuration Procedure](Mihomo%2009%20-%20Configuration%20Procedure.md) 中分析了 Clash 解析配置的过程，显然可以得出默认会使用如下几个链接获取 GEOIP GEOSITE 相关的信息

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

但是在 GUI client 中同步方式可能不一样。但在 Clash Verge Rev 会同样使用 MetaCubeX/meta-rules-dat 中的 geoip.dat 和 geosite.dat(具体代码看 [clash-verge-rev/scripts/check.mjs at main · clash-verge-rev/clash-verge-rev · GitHub](https://github.com/clash-verge-rev/clash-verge-rev/blob/main/scripts/check.mjs))

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

### 0x03c File Formats[^5]

> [!NOTE] 
> 以 geoip 为切入点，geosite 逻辑上相同

MetaCubeX/meta-rules-dat 中关于 geoip geosite 有 3 种格式的文件，但是文档里并没有介绍几者之间的区别。如果我们要想搞清楚区别，就需要从 Mihomo 源码入手

#### db metadb

geoip.db 和 geoip.metadb 会在 [func (p \*path) MMDB()](https://github.com/MetaCubeX/mihomo/blob/Meta/constant/path.go#L95) 中被引用

```go
func (p *path) MMDB() string {
	files, err := os.ReadDir(p.homeDir)
	if err != nil {
		return ""
	}
	for _, fi := range files {
		if fi.IsDir() {
			// 目录则直接跳过
			continue
		} else {
			if strings.EqualFold(fi.Name(), "Country.mmdb") ||
				strings.EqualFold(fi.Name(), "geoip.db") ||
				strings.EqualFold(fi.Name(), "geoip.metadb") {
				GeoipName = fi.Name()
				return P.Join(p.homeDir, fi.Name())
			}
		}
	}
	return P.Join(p.homeDir, "geoip.metadb")
}
```

然后会在 [IPInstance()](https://github.com/MetaCubeX/mihomo/blob/Meta/component/mmdb/mmdb.go#L60) 被调用

```go
func IPInstance() IPReader {
	IPonce.Do(func() {
		mmdbPath := C.Path.MMDB()
		log.Infoln("Load MMDB file: %s", mmdbPath)
		mmdb, err := maxminddb.Open(mmdbPath)
		if err != nil {
			log.Fatalln("Can't load MMDB: %s", err.Error())
		}
		IPreader = IPReader{Reader: mmdb}
		switch mmdb.Metadata.DatabaseType {
		case "sing-geoip":
			IPreader.databaseType = typeSing
		case "Meta-geoip0":
			IPreader.databaseType = typeMetaV0
		default:
			IPreader.databaseType = typeMaxmind
		}
	})

	return IPreader
}
```

这里可以看到会根据 `maxminddb.Open(mmdbPath)` 封装的 `mmdb.Metadata.DatabaseType` 类型来生成 IPreader。 

而 IPreader 会在 [Match(metadata \*C.Metadata)](https://github.com/MetaCubeX/mihomo/blob/Meta/rules/common/geoip.go#L34) 中用于对比数据包中目的 IP 的 GEO CODE 和 IPreader 中载入的 geoip.db 或者是 geoip.metadb (这里只讨论目的 IP)

```go
func (g *GEOIP) Match(metadata *C.Metadata) (bool, string) {
	ip := metadata.DstIP
	...
	// C.GeodataMode 取值的逻辑在 https://github.com/MetaCubeX/mihomo/blob/Meta/main.go
	// flag.BoolVar(&geodataMode, "m", false, "set geodata mode")
	// if geodataMode {
	//	C.GeodataMode = true
	//}
	// geodataMode 需要通过 -m 参数传递，默认为 false
	if !C.GeodataMode {
		...
		metadata.DstGeoIP = mmdb.IPInstance().LookupCode(ip.AsSlice())
		for _, code := range metadata.DstGeoIP {
			if g.country == code {
				return true, g.adapter
			}
		}
		...
	}
	...
}
```

虽然 `IPInstance()` 会根据 `mmdb.Metadata.DatabaseType` 值返回不同的 IPreader，但是最后只有一个 code 会被使用到，所以 geoip.db 和 geoip.meta 在使用上其实没有什么不同(也可以得出 country.mmdb 在使用上也相同)

到目前为止并不能发现有什么特殊的区别。那就需要知道 geoip.db 和 geoip.metadb 是怎么生成的，`sing-geoip` 和 `Meta-geoip0` 又是什么东西

geoip.db 和 geoip.metadb 会在 [meta-rules-dat/.github/workflows/run.yml at master · MetaCubeX/meta-rules-dat · GitHub](https://github.com/MetaCubeX/meta-rules-dat/blob/master/.github/workflows/run.yml) 中生成

```yaml
      - name: Build db and metadb file
        env:
          NO_SKIP: true
        # geo convert ip -i <input_type> -o <output_type> -f [output_filename] input_filename
        run: |
          go install -trimpath -ldflags="-s -w -buildid=" github.com/metacubex/geo/cmd/geo@master
          geo convert site -i v2ray -o sing -f geosite.db ./custom/publish/geosite.dat
          geo convert site -i v2ray -o sing -f geosite-lite.db ./community/geosite-lite.dat
          geo convert ip -i v2ray -o sing -f geoip.db ./geoip.dat
          geo convert ip -i v2ray -o meta -f geoip.metadb ./geoip.dat
          geo convert ip -i v2ray -o sing -f geoip-lite.db ./geoip-lite.dat
          geo convert ip -i v2ray -o meta -f geoip-lite.metadb ./geoip-lite.dat
```

而 `sing-geoip` 和 `Meta-geoip0` 会分别在 `V2RayIPToSing(geoipList []*v2raygeo.GeoIP, output io.Writer)` 和 `V2RayIPToMetaV0(geoipList []*v2raygeo.GeoIP, output io.Writer)` 中被引用。而相应的 `V2RayIPToSing(geoipList []*v2raygeo.GeoIP, output io.Writer)` 和 `V2RayIPToMetaV0(geoipList []*v2raygeo.GeoIP, output io.Writer)` 会在 [ip(cmd \*cobra.Command, args []string)](https://github.com/MetaCubeX/geo/blob/master/cmd/geo/internal/convert/ip.go) 中被调用

```go
...
func init() {
	CommandIP.PersistentFlags().StringVarP(&fromType, "from-type", "i", "", "specify input database type")
	CommandIP.PersistentFlags().StringVarP(&toType, "to-type", "o", "meta", "set output database type")
	CommandIP.PersistentFlags().StringVarP(&output, "output-name", "f", "", "specify output filename")
}
...
func ip(cmd *cobra.Command, args []string) error {
	var (
		buffer   bytes.Buffer
		...
	)
	...
	fileContent, err := os.ReadFile(args[0])
	...
	buffer.Grow(8 * 1024 * 1024) // 8 MiB
	...
	switch strings.ToLower(fromType) {
	...
	case "v2ray":
		var geoipList []*v2raygeo.GeoIP
		geoipList, err = v2raygeo.LoadIP(fileContent)
		if err != nil {
			return err
		}
		switch strings.ToLower(toType) {
		case "sing", "sing-geoip":
			err = convert.V2RayIPToSing(geoipList, &buffer)
			if err != nil {
				return err
			}
			filename += ".db"

		case "meta", "meta0", "meta-geoip0":
			err = convert.V2RayIPToMetaV0(geoipList, &buffer)
			if err != nil {
				return err
			}
			filename += ".metadb"

		default:
			return E.New("unsupported output GeoIP database type: ", toType)
		}
	...
	err = os.WriteFile(filename, buffer.Bytes(), 0o666)
	...
}

```

当 cmd 为 `geo convert ip -i v2ray -o sing -f geoip.db ./geoip.dat`，就会调用 `convert.V2RayIPToSing(geoipList, &buffer)`；当 cmd 为 `geo convert ip -i v2ray -o meta -f geoip.metadb ./geoip.dat`，就会调用 `convert.V2RayIPToMetaV0(geoipList, &buffer)`

对比 `V2RayIPToSing` 和 `V2RayIPToMetaV0` 这两个函数，就可以发现
- geoip.db 只包含 CIDR 和 GEO CODE
- geoip.metadata 会根据 geoip.dat 中的内容，来判断生成的内容
```go
func V2RayIPToSing(geoipList []*v2raygeo.GeoIP, output io.Writer) error {
	writer, err := mmdbwriter.New(mmdbwriter.Options{
		DatabaseType: "sing-geoip",
		Languages: common.Map(geoipList, func(it *v2raygeo.GeoIP) string {
			return strings.ToLower(it.CountryCode)
		}),
		IPVersion:               6,
		RecordSize:              24,
		Inserter:                inserter.ReplaceWith,
		DisableIPv4Aliasing:     true,
		IncludeReservedNetworks: true,
	})
	...
	for _, geoipEntry := range geoipList {
		for _, cidrEntry := range geoipEntry.Cidr {
			ipAddress := net.IP(cidrEntry.Ip)
			...
			ipNet := net.IPNet{
				IP:   ipAddress,
				Mask: net.CIDRMask(int(cidrEntry.Prefix), len(ipAddress)*8),
			}
			err = writer.Insert(&ipNet, mmdbtype.String(strings.ToLower(geoipEntry.CountryCode)))
			...
		}
	}
	...
}

func V2RayIPToMetaV0(geoipList []*v2raygeo.GeoIP, output io.Writer) error {
	writer, err := mmdbwriter.New(mmdbwriter.Options{
		DatabaseType:            "Meta-geoip0",
		IPVersion:               6,
		RecordSize:              24,
		Inserter:                inserter.ReplaceWith,
		DisableIPv4Aliasing:     true,
		IncludeReservedNetworks: true,
	})
	...
	included := make([]netip.Prefix, 0, 4*len(geoipList))
	codeMap := make(map[netip.Prefix][]string, 4*len(geoipList))
	for _, geoipEntry := range geoipList {
		code := strings.ToLower(geoipEntry.CountryCode)
		for _, cidrEntry := range geoipEntry.Cidr {
			addr, ok := netip.AddrFromSlice(cidrEntry.Ip)
			...
			prefix := netip.PrefixFrom(addr, int(cidrEntry.Prefix))
			included = append(included, prefix)
			codeMap[prefix] = append(codeMap[prefix], code)
		}
	}
	included = common.Uniq(included)
	...
	for _, prefix := range included {
		ipAddress := net.IP(prefix.Addr().AsSlice())
		ipNet := net.IPNet{
			IP:   ipAddress,
			Mask: net.CIDRMask(prefix.Bits(), len(ipAddress)*8),
		}
		codes := codeMap[prefix]
		_, record := writer.Get(ipAddress)
		switch typedRecord := record.(type) {
		case nil:
			if len(codes) == 1 {
				record = mmdbtype.String(codes[0])
			} else {
				record = mmdbtype.Slice(common.Map(codes, func(it string) mmdbtype.DataType {
					return mmdbtype.String(it)
				}))
			}
		case mmdbtype.String:
			recordSlice := make(mmdbtype.Slice, 0, 1+len(codes))
			recordSlice = append(recordSlice, typedRecord)
			for _, code := range codes {
				recordSlice = append(recordSlice, mmdbtype.String(code))
			}
			recordSlice = common.Uniq(recordSlice)
			if len(recordSlice) == 1 {
				record = recordSlice[0]
			} else {
				record = recordSlice
			}
		case mmdbtype.Slice:
			recordSlice := typedRecord
			for _, code := range codes {
				recordSlice = append(recordSlice, mmdbtype.String(code))
			}
			recordSlice = common.Uniq(recordSlice)
			record = recordSlice
		default:
			panic("bad record type")
		}
		err = writer.Insert(&ipNet, record)
		...
	}
	...
}
```

所以结论就是 geoip.db 和 geoip.metadb 在使用其实没有什么区别(也包括 country.mmdb)，但是 geoip.metadb 包含的信息可能比 geoip.db 更多。同理 geosite

#### dat

在 [0x02a geoip.dat](#0x02a%20geoip.dat) workflows 部分已经分析了 Mihomo 使用的 geoip.dat 和 Loyalsoldier/geoip 中的 geoip.dat 一致，但是入口和 db 以及 metadb 并不相同

```go
func (p *path) GeoIP() string {
	files, err := os.ReadDir(p.homeDir)
	if err != nil {
		return ""
	}
	for _, fi := range files {
		if fi.IsDir() {
			// 目录则直接跳过
			continue
		} else {
			if strings.EqualFold(fi.Name(), "GeoIP.dat") {
				GeoipName = fi.Name()
				return P.Join(p.homeDir, fi.Name())
			}
		}
	}
	return P.Join(p.homeDir, "GeoIP.dat")
}
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