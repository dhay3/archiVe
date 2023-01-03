## Digest
syntax: `openssl command [ command_opts ] [ command_args ]`
OpenSSL 是一个实现了 TLS/SSL 的加密工具，可以用于

1. Creation and management of private keys, public keys and parameters
2. Public key cryptographic operations
3. Creation of X.509 certificates, CSRs and CRLs
4. Calculation of Message Digests
5. Encryption and Decryption with Ciphers
6. SSL/TLS Client and Server Tests
7. Handling of S/MIME signed or encrypted mail
8. Time Stamp requests, generation and verification
## Version
openssl 版本不同支持的加密协议也不同，例如 1.1.0 支持 TLS 1.2 但是不支持 TLS 1.3。同时 OPENSSLDIR 对应会记录 openssl 使用的配置文件目录以及 CA 根证书
```
$openssl version -a
OpenSSL 1.0.2k-fips  26 Jan 2017
built on: reproducible build, date unspecified
platform: linux-x86_64
options:  bn(64,64) md2(int) rc4(16x,int) des(idx,cisc,16,int) idea(int) blowfish(idx) 
compiler: gcc -I. -I.. -I../include  -fPIC -DOPENSSL_PIC -DZLIB -DOPENSSL_THREADS -D_REENTRANT -DDSO_DLFCN -DHAVE_DLFCN_H -DKRB5_MIT -m64 -DL_ENDIAN -Wall -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector-strong --param=ssp-buffer-size=4 -grecord-gcc-switches   -m64 -mtune=generic -Wa,--noexecstack -DPURIFY -DOPENSSL_IA32_SSE2 -DOPENSSL_BN_ASM_MONT -DOPENSSL_BN_ASM_MONT5 -DOPENSSL_BN_ASM_GF2m -DRC4_ASM -DSHA1_ASM -DSHA256_ASM -DSHA512_ASM -DMD5_ASM -DAES_ASM -DVPAES_ASM -DBSAES_ASM -DWHIRLPOOL_ASM -DGHASH_ASM -DECP_NISTZ256_ASM
OPENSSLDIR: "/etc/pki/tls"
engines:  rdrand dynamic
```
## Subcommands
openssl 提供了 help subcommands 用于查看所有支持的 subcommands，常用的 subcommands 有

- ca
## Key and certificate management
大部分用户使用 openssl 主要用于部分 SSL 到服务器

1. generate a private key
2. create a certificate signing request (CSR) and send it to a CA
3. install the CA provided certificate in your web server
### Key generation
在服务器上部署 SSL, 首先需要生成 private key，而 private key 通常由几个要素决定
**key algorithm**
私钥使用的加密算法。openssl 支持 RSA, DSA (obsoleted), ECDSA 和 EdDSA。通常使用 RSA 或者 ECDSA
**key size**
私钥的加密比特。RSA 默认 512 bits，为了更安全应该使用 2048 bits
**passphrase**
私钥的密码。可选，每次使用私钥时需要提供

