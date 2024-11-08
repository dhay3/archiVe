---
createTime: 2024-07-31 17:06
tags:
  - "#hash1"
  - "#hash2"
---

# Mihomo 09 - Configuration Procedure

为了更好的理解 Mihomo 配置的逻辑，现在针对源码做拆分

不用说入口就是 [main.go](https://github.com/MetaCubeX/mihomo/blob/Meta/main.go)

```go
import (
	...
	C "github.com/metacubex/mihomo/constant"
	"github.com/metacubex/mihomo/config"
	"github.com/metacubex/mihomo/hub/executor"
	...
)
func main() {
	...
	if homeDir != "" {
		if !filepath.IsAbs(homeDir) {
			currentDir, _ := os.Getwd()
			homeDir = filepath.Join(currentDir, homeDir)
		}
		C.SetHomeDir(homeDir)
	}

	if configFile != "" {
		if !filepath.IsAbs(configFile) {
			currentDir, _ := os.Getwd()
			configFile = filepath.Join(currentDir, configFile)
		}
	} else {
		configFile = filepath.Join(C.Path.HomeDir(), C.Path.Config())
	}
	...
	if err := config.Init(C.Path.HomeDir()); err != nil {
		log.Fatalln("Initial configuration directory error: %s", err.Error())
	}
	...
	for {
		select {
		case <-termSign:
			return
		case <-hupSign:
			if cfg, err := executor.ParseWithPath(C.Path.Config()); err == nil {
				executor.ApplyConfig(cfg, true)
			} else {
				log.Errorln("Parse config error: %s", err.Error())
			}
		}
	}

```

假设没有设置或者指定 `configFile`，Mihomo 会使用 [config.Init(C.Path.HomeDir()](https://github.com/MetaCubeX/mihomo/blob/Meta/config/initial.go#L12) 函数按照 `C.Path.HomeDir()/C.Path.Config` 生成一个配置，默认只含有 `mixed-port: 7890`

```go
func Init(dir string) error {
	...
	// initial config.yaml
	if _, err := os.Stat(C.Path.Config()); os.IsNotExist(err) {
		log.Infoln("Can't find config, create a initial config file")
		f, err := os.OpenFile(C.Path.Config(), os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			return fmt.Errorf("can't create file %s: %s", C.Path.Config(), err.Error())
		}
		f.Write([]byte(`mixed-port: 7890`))
		f.Close()
	}

	return nil
}
```

然后会使用 [executor.ParseWithPath(C.Path.Config())](https://github.com/MetaCubeX/mihomo/blob/Meta/hub/executor/executor.go#L65) 去解析配置

```go
func ParseWithPath(path string) (*config.Config, error) {
	buf, err := readConfig(path)
	if err != nil {
		return nil, err
	}

	return ParseWithBytes(buf)
}
```

在通过 `ParseWithBytes(buf []byte)` 调用 [config.Parse(buf)](https://github.com/MetaCubeX/mihomo/blob/Meta/config/config.go#L403)

```go
func ParseWithBytes(buf []byte) (*config.Config, error) {
	return config.Parse(buf)
}
```

再通过 `Parse(buf []byte)` 调用 [UnmarshalRawConfig(buf []byte)](https://github.com/MetaCubeX/mihomo/blob/Meta/config/config.go#L412)

```go
func Parse(buf []byte) (*Config, error) {
	rawCfg, err := UnmarshalRawConfig(buf)
	if err != nil {
		return nil, err
	}
	return ParseRawConfig(rawCfg)
}
```

会生成一个逻辑上缺省的基础配置(**即在配置文件中虽然没有声明，但是实际会使用这些配置项**)

```go
func UnmarshalRawConfig(buf []byte) (*RawConfig, error) {
	// config with default value
	rawCfg := &RawConfig{
		AllowLan:          false,
		BindAddress:       "*",
		LanAllowedIPs:     []netip.Prefix{netip.MustParsePrefix("0.0.0.0/0"), netip.MustParsePrefix("::/0")},
		IPv6:              true,
		Mode:              T.Rule,
		GeoAutoUpdate:     false,
		GeoUpdateInterval: 24,
		GeodataMode:       C.GeodataMode,
		GeodataLoader:     "memconservative",
		UnifiedDelay:      false,
		Authentication:    []string{},
		LogLevel:          log.INFO,
		Hosts:             map[string]any{},
		Rule:              []string{},
		Proxy:             []map[string]any{},
		ProxyGroup:        []map[string]any{},
		TCPConcurrent:     false,
		FindProcessMode:   P.FindProcessStrict,
		GlobalUA:          "clash.meta/" + C.Version,
		Tun: RawTun{
			Enable:              false,
			Device:              "",
			Stack:               C.TunGvisor,
			DNSHijack:           []string{"0.0.0.0:53"}, // default hijack all dns query
			AutoRoute:           true,
			AutoDetectInterface: true,
			Inet6Address:        []netip.Prefix{netip.MustParsePrefix("fdfe:dcba:9876::1/126")},
		},
		TuicServer: RawTuicServer{
			Enable:                false,
			Token:                 nil,
			Users:                 nil,
			Certificate:           "",
			PrivateKey:            "",
			Listen:                "",
			CongestionController:  "",
			MaxIdleTime:           15000,
			AuthenticationTimeout: 1000,
			ALPN:                  []string{"h3"},
			MaxUdpRelayPacketSize: 1500,
		},
		EBpf: EBpf{
			RedirectToTun: []string{},
			AutoRedir:     []string{},
		},
		IPTables: IPTables{
			Enable:           false,
			InboundInterface: "lo",
			Bypass:           []string{},
			DnsRedirect:      true,
		},
		NTP: RawNTP{
			Enable:        false,
			WriteToSystem: false,
			Server:        "time.apple.com",
			ServerPort:    123,
			Interval:      30,
		},
		DNS: RawDNS{
			Enable:         false,
			IPv6:           false,
			UseHosts:       true,
			UseSystemHosts: true,
			IPv6Timeout:    100,
			EnhancedMode:   C.DNSMapping,
			FakeIPRange:    "198.18.0.1/16",
			FallbackFilter: RawFallbackFilter{
				GeoIP:     true,
				GeoIPCode: "CN",
				IPCIDR:    []string{},
				GeoSite:   []string{},
			},
			DefaultNameserver: []string{
				"114.114.114.114",
				"223.5.5.5",
				"8.8.8.8",
				"1.0.0.1",
			},
			NameServer: []string{
				"https://doh.pub/dns-query",
				"tls://223.5.5.5:853",
			},
			FakeIPFilter: []string{
				"dns.msftnsci.com",
				"www.msftnsci.com",
				"www.msftconnecttest.com",
			},
		},
		Experimental: Experimental{
			// https://github.com/quic-go/quic-go/issues/4178
			// Quic-go currently cannot automatically fall back on platforms that do not support ecn, so this feature is turned off by default.
			QUICGoDisableECN: true,
		},
		Sniffer: RawSniffer{
			Enable:          false,
			Sniff:           map[string]RawSniffingConfig{},
			ForceDomain:     []string{},
			SkipDomain:      []string{},
			Ports:           []string{},
			ForceDnsMapping: true,
			ParsePureIp:     true,
			OverrideDest:    true,
		},
		Profile: Profile{
			StoreSelected: true,
		},
		GeoXUrl: GeoXUrl{
			Mmdb:    "https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/geoip.metadb",
			ASN:     "https://github.com/xishang0128/geoip/releases/download/latest/GeoLite2-ASN.mmdb",
			GeoIp:   "https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/geoip.dat",
			GeoSite: "https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/geosite.dat",
		},
		ExternalUIURL: "https://github.com/MetaCubeX/metacubexd/archive/refs/heads/gh-pages.zip",
	}

	if err := yaml.Unmarshal(buf, rawCfg); err != nil {
		return nil, err
	}

	return rawCfg, nil
}
```

回到 `Parse(buf []byte)` 的逻辑，Parse 会在调用 `UnmarshalRawConfig(buf []byte)` 后调用 [ParseRawConfig(rawCfg \*RawConfig)](https://github.com/MetaCubeX/mihomo/blob/Meta/config/config.go#L537)

```go
func ParseRawConfig(rawCfg *RawConfig) (*Config, error) {
	config := &Config{}
	...
	general, err := parseGeneral(rawCfg)
	...
	config.General = general
	...
	proxies, providers, err := parseProxies(rawCfg)
	...
	config.Proxies = proxies
	config.Providers = providers
	...
	listener, err := parseListeners(rawCfg)
	...
	config.Listeners = listener
	...
	ruleProviders, err := parseRuleProviders(rawCfg)
	...
	config.RuleProviders = ruleProviders
	...
	subRules, err := parseSubRules(rawCfg, proxies, ruleProviders)
	...
	etc.
```

在各个 `parseXxx(cfg *RawConfig)` 中对 config 的封装类或者是对 const 赋值(同时也会对`UnmarshalRawConfig(buf []byte) ` 中生产逻辑上缺省的基础配置进行覆写)，例如 [parseGeneral(cfg \*RawConfig)](https://github.com/MetaCubeX/mihomo/blob/Meta/config/config.go#L638)

```go
func parseGeneral(cfg *RawConfig) (*General, error) {
	geodata.SetGeodataMode(cfg.GeodataMode)
	geodata.SetGeoAutoUpdate(cfg.GeoAutoUpdate)
	geodata.SetGeoUpdateInterval(cfg.GeoUpdateInterval)
	geodata.SetLoader(cfg.GeodataLoader)
	geodata.SetSiteMatcher(cfg.GeositeMatcher)
	C.GeoAutoUpdate = cfg.GeoAutoUpdate
	C.GeoUpdateInterval = cfg.GeoUpdateInterval
	C.GeoIpUrl = cfg.GeoXUrl.GeoIp
	C.GeoSiteUrl = cfg.GeoXUrl.GeoSite
	C.MmdbUrl = cfg.GeoXUrl.Mmdb
	C.ASNUrl = cfg.GeoXUrl.ASN
	C.GeodataMode = cfg.GeodataMode
	C.UA = cfg.GlobalUA
	...
	return &General{
		Inbound: Inbound{
			Port:              cfg.Port,
			SocksPort:         cfg.SocksPort,
			RedirPort:         cfg.RedirPort,
			TProxyPort:        cfg.TProxyPort,
			MixedPort:         cfg.MixedPort,
			ShadowSocksConfig: cfg.ShadowSocksConfig,
			VmessConfig:       cfg.VmessConfig,
			AllowLan:          cfg.AllowLan,
			SkipAuthPrefixes:  cfg.SkipAuthPrefixes,
			LanAllowedIPs:     cfg.LanAllowedIPs,
			LanDisAllowedIPs:  cfg.LanDisAllowedIPs,
			BindAddress:       cfg.BindAddress,
			InboundTfo:        cfg.InboundTfo,
			InboundMPTCP:      cfg.InboundMPTCP,
		},
		Controller: Controller{
			ExternalController:     cfg.ExternalController,
			ExternalUI:             cfg.ExternalUI,
			Secret:                 cfg.Secret,
			ExternalControllerUnix: cfg.ExternalControllerUnix,
			ExternalControllerTLS:  cfg.ExternalControllerTLS,
			ExternalDohServer:      cfg.ExternalDohServer,
		},
		UnifiedDelay:            cfg.UnifiedDelay,
		Mode:                    cfg.Mode,
		LogLevel:                cfg.LogLevel,
		IPv6:                    cfg.IPv6,
		Interface:               cfg.Interface,
		RoutingMark:             cfg.RoutingMark,
		GeoXUrl:                 cfg.GeoXUrl,
		GeoAutoUpdate:           cfg.GeoAutoUpdate,
		GeoUpdateInterval:       cfg.GeoUpdateInterval,
		GeodataMode:             cfg.GeodataMode,
		GeodataLoader:           cfg.GeodataLoader,
		TCPConcurrent:           cfg.TCPConcurrent,
		FindProcessMode:         cfg.FindProcessMode,
		EBpf:                    cfg.EBpf,
		GlobalClientFingerprint: cfg.GlobalClientFingerprint,
		GlobalUA:                cfg.GlobalUA,
	}, nil
```

然后在其他类中通过引入 config 中的结构体或者是对应的 const 来获取配置对应的值，例如 [updateGeoDatabases()](https://github.com/MetaCubeX/mihomo/blob/Meta/component/updater/update_geo.go#L24)

```go
func updateGeoDatabases() error {
	...
	if C.GeodataMode {
		data, err := downloadForBytes(C.GeoIpUrl)
		if err != nil {
			return fmt.Errorf("can't download GeoIP database file: %w", err)
		}
	...
}
```

自此 Mihomo 配置解析的过程就完成了

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

