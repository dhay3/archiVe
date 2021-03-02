# OpenSSL arch

参考：

https://www.openssl.org/docs/OpenSSLStrategicArchitecture.html

## Architecture

OpenSSL主要有4个组件构成

### libcrypto

提供加密的模块基础API，一般被Engine调用

### Engine

可被加载的模块，通过该模块调用libcrypto

### libssl

依赖libcrypto，实现TLS和DTLS协议

### application

提供cmdline tool的集合



application依赖于TLS和crypto，TLS依赖于crypto，Engine调用crypto

![架构图](D:\asset\note\imgs\_openssl\Snipaste_2021-03-01_10-34-10.png)

![](D:\asset\note\imgs\_openssl\Snipaste_2021-03-01_10-37-55.png)

