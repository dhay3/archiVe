# SSL&TLS Digest

https://www.ssl.com/faqs/faq-what-is-ssl/#

https://en.wikipedia.org/wiki/Transport_Layer_Security#SSL_1.0,_2.0,_and_3.0

https://segmentfault.com/a/1190000014740303

https://www.ibm.com/docs/en/ibm-mq/7.5?topic=ssl-overview-tls-handshake

https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange

https://en.wikipedia.org/wiki/Cipher_suite

https://blog.csdn.net/qq_31442743/article/details/116199453

https://security.stackexchange.com/questions/89383/why-does-the-ssl-tls-handshake-have-a-client-and-server-random

https://security.stackexchange.com/questions/218491/why-using-the-premaster-secret-directly-would-be-vulnerable-to-replay-attack

https://medium.com/@clu1022/%E9%82%A3%E4%BA%9B%E9%97%9C%E6%96%BCssl-tls%E7%9A%84%E4%BA%8C%E4%B8%89%E4%BA%8B-%E4%B9%9D-ssl-communication-31a2a8a888a6

## Digest

Secure sockets Layer (SSL) is protocol for establishing authenticated and encrypted links between networked computers

![Snipaste_2020-08-25_17-24-19](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2020-08-25_17-24-19.4429l90bojo0.png)

SSL 是一个被广泛应用的加密及身份验证的表现层( presentation )协议，中文直译为 安全套接层

- SMTP
- POP3
- HTTPS

这些都是比较常见的基于 SSL 的协议

## TLS

Transport Layer security ( TLS ) 是 IETF 标准，在 1999 后替代了 SSL，其核心功能和 SSL 一样。但是一般也会将 TLS 统称为 SSL

按照版本可以分成 TLS 1.0 /TLS 1.1/ ... 目前最新的是 TLS 1.3

## SSL certificate 

在 SSL/TLS 体系中有一个核心角色 -- SSL certificate 也被称为 PKI certificates

这里简单的介绍一下，具体可以查看  [PKI.md]() 

SSL certifacte 一般都是 X.509 certificates，在 SSL 中主要起校验 server 的有效性和准确性

## Pre-master Secrect/Master Secrt

- Pre-master secret 

  中文预主密钥 是通过 Key exchange 获取的，对应 [Diffie-Hellman algrothim#Cryptographic explations]() 的第 3，5 步，对应 TLS 中的 client key exchange 和 server key exchange。可以直接理解成 DH 生成的 publickey

- Master secret 是在 pre-master key 基础上==加两个随机数==（在 clienthello 和 serverhello 中宣告的）生成的对称密钥，==不会在网络上传输==。也被称为 shared key

  ```
  shared key = client_random + server_random + pre-master_secret
  ```

  [Diffie-Hellman algrothim#Cryptographic explations]() 中第 6，7 步，对应 TLS 中的 encrypted handshake message 是用来校验 master key 的

## SSL/TLS handshake

> DH 算法可以参考 [Diffie-Hellamn algrothim]()

An SSL/TLS handshake is a negotiation between two parties on a network - such as browser and web server - to establish the details of their connection. It determines 

1. what version of SSL/TLS will be used
2. which cipher suite will encrypt communication
3. verfies the server( and sometimes also the client )
4. establishes that a secure connection is in place before transferring data

handshake phase 主要包含下列几个流程

![2022-12-17_00-50](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221217/2022-12-17_00-50.5p1rin7gsd1c.webp)

![2022-12-17_03-35](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221221/2022-12-17_03-35.20m8hk4mgwsg.webp)

这里的 retrans 和 dup ack 可以忽略，中间过代理 RTO 重传

### Connect Request

client 发送 4 层 SYN 报文

### Connection Ackownledged

server 回送 4 层 SYN-ACK 报文

### ClientHello

client 回送 ACK 报文，并发送 ClientHello

![2022-12-17_01-54](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221221/2022-12-17_01-54.2ungvtpo8pq8.webp)

向 server 宣告 

- client 使用的 SSL/TLS 版本

- 用于生成 master key 的随机数 A
- 随机的 sessionId
- client 支持的 cipher suites
- client 支持的压缩算法
- client 支持的 签名哈希算法( 用于校验证书 )

### ServerHello

server 收到  ClientHello 是会回送一个 ACK 包

如果 ClientHello 中的信息都没有问题并且 server 都支持相应的配置，就会回送 ServerHello （这里也可以采用 TCP 的捎带回送，只回送 ServerHello）。反之会回送 Server handshake error

![2022-12-17_03-27](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221221/2022-12-17_03-27.5gugf5iw69o.webp)

向 client 宣告

- server 采用 SSL/TLS 的版本
- 用于生成 master key 的随机数 B
- 记录的 sessionId ( client 发送的 sessionId )
- server 从 clienthello 中挑选并采用的 cipher suites
- server 从 clienthello 中挑选并采用的压缩算法

### Certificate

client 发送 ACK 报文，server 收到后发送 server certificate 给 client，同理一样也可以捎带回送

![2022-12-17_03-18](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221221/2022-12-17_03-18.4x4n6esato1s.webp)

注意这里会回送 2 个证书，一个是 server 自己的证书，另一个是用于校验 server 证书的 CA 根证书

### ServerKey exchange

宣告使用 ECDH( Diffle-Hellman ) 算法同时宣告DH 算法生成的 Publickey (这里的公钥不是证书的，是通过 DH 算法计算出来的)

![2022-12-17_02-23](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221221/2022-12-17_02-23.2ij9uenwmtxc.webp)

### ServerHello done

声明 ServerHello 相关的信息已经全部发送

![2022-12-21_18-02](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221221/2022-12-21_18-02.552hgcoeim4g.webp)

### Clientkey exchange

宣告使用 ECDH( Diffle-Hellman ) 算法同时宣告DH 算法生成的 Publickey (这里的公钥不是证书的，是通过 DH 算法计算出来的)

![2022-12-21_18-04](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221221/2022-12-21_18-04.33nbez4hzn0g.webp)

### Change cipher spec(client)

向 server 宣告随后的信息都将用双方商定的加密方法和密钥发送

ChangeCipherSpec是一个独立的协议，体现在数据包中就是一个字节的数据，用于告知对端已经切换到之前协商好的加密套件（Cipher Suite）的状态，准备使用之前协商好的加密套件加密数据并传输了。

![2022-12-21_18-05](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221221/2022-12-21_18-05.3znn36x4yon4.webp)

### Encrypted handshake message 

用于校验 DH 算法中生成的 master key 

![2022-12-21_18-06](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221221/2022-12-21_18-06.2mtykp14lxs0.webp)

### Change cipher spec (server)

向 client 宣告随后的信息都将用双方商定的加密方法和密钥发送

![2022-12-21_18-08](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221221/2022-12-21_18-08.16kid8v02a1s.webp)

### Encrypted handshake message 

用于校验 DH 算法中生成的 master key 

![2022-12-21_18-08_1](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221221/2022-12-21_18-08_1.4i1bi0qkhf5s.webp)

### Application data

加密后的报文，这里就不展示了

## Cipher suite

cipher suite 直译就是密码套件，是一组算法的集合。按照使用类型分成 3 种

1. key exchange algorithm

   用于两端 key exchange 的算法

2. bulk encryption alogrithm

   对数据加密的算法

3. MAC alogrithm

   ensure that the data sent does not change in transit

   确认报文的一致性

chipher suite 会以一串类似 `ECDHE-RSA-AES128-GCM-SHA256` 的字符显示，从左到右解释为

1. ECDHE indicates the ==key exchange algorithm== being used
2. RSA authentication mechanish during the handshake
3. AES session cipher
4. 128 session encryption key size (bits) for cipher
5. GCM type of encryption (cipher-block dependency and additional options)
6. SHA (SHA2)hash function. For a digest of 256 and higher.  Signature mechanism. Indicates the message authentication algorithm which is used to authenticate a message
7. 256 Digest size (bits)

## Why need nonces

ClientHello random number 和 ServerHello random number  也被统称为 nonces

Pre-master secrect 是由私钥计算而来，所以 MITM 是不能破解然后对报文解密的。既然 MITM 不能解密那么为什么 SSL/TLS 需要这两个随机数来生成 Master secrect 并采用 Master secrect 来加密解密呢？

Premaster 本身就是一个随机数，且能在所有主机上生成，那么 Premaster 就有可能被猜出来，所以我们需要另外 2 个随机数来增加随机性，这样生成的密钥就不容易被猜出来了

网上也有说是为了防止重放攻击( replay attack )，==但是个人认为是无法防止重放攻击的。==如果攻击者仅仅是复制一份原始报文，由于报文是合理的，服务端任然能处理

*replay attack is someone who is not authorize to get the response, copy your request and getting the same response!!.*

## No Certificate phase in wireshark

https://osqa-ask.wireshark.org/questions/62514/tls-handshake-without-server-certificate/

![2022-12-17_01-34](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221217/2022-12-17_01-34.5uoi2amka3cw.webp)

wireshark 中有一个 bug 如果证书的认证链较长，那么 SSL 中的报文的内容就非常多，目前 wireshark 无法正常识别。所以就会出现缺少 Certificate phase 的现象