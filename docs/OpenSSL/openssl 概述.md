# openssl 概述

> 使用`openssl list`可以查看openssl中自带的命令，所有的subCommands都自带`-help`

bash completion 

https://github.com/scop/bash-completion/blob/master/completions/openssl

## 概述

openssl是一个调用openssl libcrypto api的命令行工具，有主要如下几个功能

- Creation and management of private keys, public keys and parameters
- Public key cryptographic operations
- Creation of  ==X.509== certificates, CSRs and CRLs
- Calculation of Message Digests and Message Authentication Codes
- Encryption and Decryption with Ciphers
- SSL/TLS Client and Server Tests
- Handling of S/MIME signed or encrypted mail
- Timestamp requests, generation and verification

默认的配置文件在`/etc/ssl/openssl.cnf`