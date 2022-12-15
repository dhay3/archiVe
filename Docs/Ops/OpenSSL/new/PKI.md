# PKI

ref

https://www.keyfactor.com/resources/what-is-pki/

https://www.cryptomathic.com/news-events/blog/symmetric-key-encryption-why-where-and-how-its-used-in-banking

https://www.ssh.com/academy/pki

https://www.trentonsystems.com/blog/symmetric-vs-asymmetric-encryption

https://zhuanlan.zhihu.com/p/30136885

https://juejin.cn/post/6844903977025273863

https://www.barretlee.com/blog/2016/04/24/detail-about-ca-and-certs/

https://zhuanlan.zhihu.com/p/423506052

## Digest

publick key infrastructure ( PKI ) 公钥基础设施，是一套用于身份认证以及信息加密的规则系统。PKI 应用的一个例子是 TLS( SSL ) certificates 

在了解 PKI 之前需要了解 密钥、对称加密 和 非对称加密。具体查看 [Symmetric Encryption & Asymmetric Encryption.md]()

## Why PKI comes

Symmetric encryption 和 Asymmetric encryption 都有一个比较严重的问题。就是怎么确认需要分发密钥(对称加密的密钥，非对称加密的公钥 )的有效性和准确性，也就是防止 MITMA (man in the middle attack, 具体参考 [MITMA.md]())

PKI 就是在这种场景下产生的，下面一段内容是摘自

[ssh 官网](https://www.ssh.com/academy/pki) 的内容

The users and devices that have keys are often just called entities. In general, anything can be associated with a key that it can use as its identity. Besides a user or a devicce, it could be a program, process, manufacturer, component, or something else the purpose of a PKI is to securely associate a key with an entity

其实核心思想就是验证 cryptographic key 的有效性和准确性

## Roles of PKI

在 PKI 系统有这几个 核心 角色

- certificate signing request

  CSR, 中文叫做 证书请求, 申请 certificate 需要提交给 CA 的文件，包含申请人的公钥以及相关信息

- certificate

  由 CA 签发的证书

- certificate authority

  CA, the trusted party siging the document associating the key with the device is called a certificate authority

   对关联设备的文件(方便记忆可以直接理解成证书)进行颁布和签名，即权威机构。任何主机或者这个人都可以扮演 CA，但是因为不被公开范围承认所以得不到信任( ==需要将其加入到系统的受信任的根证书颁发机构中== )

  知名的CA有Symantec，Comodo，Godaddy，GolbalSign，Digicert，let’s Encrypt（免费）

- root certificate

  中文叫做 根证书，可以直接理解成 CA 的公钥( 实际中还有其他一些信息，成员和普通证书一样 )，存储在本地。主要用于验证 CA 颁发的证书的有效性

  在windows上可以通过`certmgr.msc`服务查看存储的根证书

  ![Snipaste_2021-03-03_10-03-17](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-03-03_10-03-17.56yitsbh8aw0.png)

  linux 存储默认在 `/etc/ssl/certs/ca-bundle.crt` , 可以通过`openssl x509 -text -in certificate` 来查看

## Digital Certificates

digital certificates (数字证书) 通常也被称为 X.509 certificates( 最常见的格式就是 X.509 ) 或者 PKI certificates 或者是 TLS/SSL 证书

X.509 是数字证书的一种 ISO 格式，其中 X.509V3 是现在最广为流传的版本。通常为了方便直接称为 X.509 ( 具体可以查看 RFC 3280 )

### Content of Certificates

![Snipaste_2021-03-03_11-36-32](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-03-03_11-36-32.2re74gpl9wo0.png)

以 windows 上的为例，主要包含如下几个字段

1. 证书持有人的公钥，即申请人的公钥
2. 签名以及哈希算法，如上表示证书是先通过 CA 的私钥加密(签名)，然后 sha1 哈希
3. 证书序列号 UUID
4. 有效期
5. 证书版本
6. 证书分发机构
7. 证书使用机构

### Formats of Certificates

> 在 Linux 并不以文件后缀作文标识，所以证书文件，可能叫`xxx.pem` 或者 `xxx.crt` 或这 `xxx.cer` 等等都行

证书通常有两种格式

- DER

  istinguished encoding rules

  ASN.1 数据加密后的二进制格式。证书一般不会以这种格式显示

- PEM

  privacy enhanced mail 

  DER base64编码后的格式。证书(==实际上不仅仅是证书，还包含 证书请求、公私钥==)一般以这种形式显示

  ```
  ----BEGIN PRIVATE KEY-----
  MAwGCCqGSIb3DQIJBQAwFAYIKoZIhvcNAwcECPMFLXmqQ/0uBIIEyJOJv78Izytj
  -----END PRIVATE KEY-----
  ```

### Certificate types

证书根据场景和支持的功能主要分如下几种

- DV（Domain Validation）

  面向个体用户，安全体系相对较弱，验证方式就是向 whois 信息中的邮箱发送邮件，按照邮件内容进行验证即可通过；

- OV（Organization Validation）

  面向企业用户，证书在 DV 证书验证的基础上，还需要公司的授权，CA 通过拨打信息库中公司的电话来确认；

- EV（Extended Validation）

  打开 Github 的网页，你会看到 URL 地址栏展示了注册公司的信息，这会让用户产生更大的信任，这类证书的申请除了以上两个确认外，还需要公司提供金融机构的开户许可证，要求十分严格。

同时证书支持的范围也不一样

|      | 单域名            | 多域名                                                   | 泛域名          | 多泛域名                                                |
| ---- | ----------------- | -------------------------------------------------------- | --------------- | ------------------------------------------------------- |
| DV   | 支持              | 不支持                                                   |                 |                                                         |
| OV   | 支持              |                                                          |                 |                                                         |
| EV   | 支持              | 不支持                                                   |                 |                                                         |
| 举例 | www.barretlee.com | www.barretlee.com<br>www.xiaohuzige.com<br>www.barret.cc | *.barretlee.com | \*.barretlee.com <br>\*.xiaohuzige.com <br>\*.barret.cc |

## How Certificates come

1. 首先申请人需要在本地生成一对 对称加密 密钥，或者复用现有的
2. 使用申请人的私钥生成 CSR（签名），CSR 中包含申请人的公钥
3. CA 收到 CSR 后，先用申请人的公钥解密。核对无误后，使用 CA 私钥对 申请人的 公钥以及其他信息签名并哈希
4. CA 将证书返回给申请人

## How to verify Certificates

1. 客户端请求服务端
2. 服务端将证书返回给客户端( 证书中包含申请证书是的公钥 )
3. 客户端获取到证书后，对证书的明文信息通过指定 哈希算法哈希，然后使用客户端上对应证书颁发机构( Issuer )的 CA 证书( 含有对证书签名时的公钥 )对密文部分做非对解密得到一串 哈希
4. 对比两者是否有差异，如果有差异即表示校验失败

## Certificate chains

和传统的分销模式一样，证书机构也有分销，这在 PKI 中被称为 认证链 （chain of trust )

![Snipaste_2021-03-03_11-12-07](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-03-03_11-12-07.5pmxm181oiw0.png)

假设现在 A 为 root CA，把分发 certificates 的权限下发给了 B，这样 B 就可以通过自己的私钥来签发证书了。现在 C 向 B 申请了证书。那么

1. 客户在请求 C 站点时，先会使用 B 的 CA 证书校验 C 的证书
2. 然后在通过 A 的 CA 证书校验 B 的证书

如果中间添加了其他的“分销”节点，在证书校验的过程就会消耗很多时间和资源

同时 B 分发给 C 的证书也会含有 B 的证书信息 (为了校验)

![2022-12-16_02-41](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221216/2022-12-16_02-41.3hpf70x75p34.webp)

## pfx

pfx 是 PKCS #12 ( public key cruptography standards #12 )在 windows 上的说法，是一种用于保存 X.509 证书文件的格式

通过它我们就可以导出和导入证书

![Snipaste_2021-03-03_10-49-46](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-03-03_10-49-46.62upmes092s0.png)