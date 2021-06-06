# IPSec

参考：

https://www.cloudflare.com/zh-cn/learning/network-layer/what-is-ipsec/

https://www.netadmin.com.tw/netadmin/zh-tw/technology/97CAE3300E9546B0A3A9443B6608A8D8

Internet Protocol Security（IPsec）是一个协议包，对IP的分组进行加密来保护IP协议，==通常使用500端口，处在网络层==。

![](D:\asset\note\imgs\_aliyun\Snipaste_2021-04-21_12-49-20.png)

通过电脑软件和路由器建立IPSec VPN Tunnel

IPsec不是一个协议，而是由一系列协议组成的：

1. 认证头(AH)，为IP数据包提供无连接数据完整性，消息认证以及重放攻击保护
2. 封装安全载荷(ESP)，提供机密性，数据源认证，无连接完整性，防重放和有限的传输流机密性
3. 因特网密钥交互(IKE)，为AH、ESP操作所需的安全关联提供算法、数据包和密钥参数

**IPSec连接流程**

1. Key exchange：密钥交换对主机将的信息加密和解密
2. Packet headers and trailers：为IP分组添加头部包含认证和授权信息
3. Authentication：确保数据包来自受信任的地址而不是攻击者
4. Encryption：对数据包中的payload加密
5. Transmission：加密数据包通过UDP协议传输（不保证可靠连接），允许通过firewall
6. Decryption：对加密的数据报解密

