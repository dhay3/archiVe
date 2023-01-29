# OpenSSL Subcommands ca

ref

https://www.feistyduck.com/library/openssl-cookbook/online/openssl-command-line/private-ca-creating-root.html

## Digest

syntax

```
openssl ca [-help] [-verbose] [-config filename] [-name section] [-section section] [-gencrl] [-revoke file] [-valid file] [-status serial]
[-updatedb] [-crl_reason reason] [-crl_hold instruction] [-crl_compromise time] [-crl_CA_compromise time] [-crl_lastupdate date] [-crl_nextupdate
date] [-crldays days] [-crlhours hours] [-crlsec seconds] [-crlexts section] [-startdate date] [-enddate date] [-days arg] [-md arg] [-policy arg]
[-keyfile filename|uri] [-keyform DER|PEM|P12|ENGINE] [-key arg] [-passin arg] [-cert file] [-certform DER|PEM|P12] [-selfsign] [-in file]
[-inform DER|<PEM>] [-out file] [-notext] [-dateopt] [-outdir dir] [-infiles] [-spkac file] [-ss_cert file] [-preserveDN] [-noemailDN] [-batch]
[-msie_hack] [-extensions section] [-extfile section] [-subj arg] [-utf8] [-sigopt nm:v] [-vfyopt nm:v] [-create_serial] [-rand_serial]
[-multivalue-rdn] [-rand files] [-writerand file] [-engine id] [-provider name] [-provider-path path] [-propquery propq] [certreq...]
```

`ca` 是 openssl 中用来模拟 CA 的工具 ，可以对 CSR 签名生成证书(也可以用 req 和 x509 来生成)以及生成 CRLs

## Optional args

- `-config filename`

  指定使用的配置的文件，如果没有指定默认会使用`openssl version -a` 中 `OPENSSLDIR` 对应的 `openssl.cnf` 配置文件

- `-name section`

- `-in filename`

  an input filename containing a CSR to be signed by the CA

- `-out filename`

  the out file to output certifcates to

- `-inform DER|PEM`

  输入 CSR 的格式

- `-cert filename`

  CA 使用的证书, 必须和 `-keyfile` 匹配

- `-keyfile filename|uri`

  CA 用于签名的 private key, 必须和 `-cert` 匹配

- `-keyform DER|PEM|P12|ENGINE`

  CA private key 格式

- `-key password`

  CA private key 的密码, 会被显示在 `ps -ef` 中, 最好使用 `-passin` 替代

- `-passin arg`

  CA private key 的密码

- `-selfsign`

  标识对 certificates 签名的 private 和 对 CSR 签名的 certificates 相同 

- `-ss_cert filename`

  a single self-signed certificate to be signed by the CA

- `-startdate | enddate date`

  证书起始和失效时间，默认以`YYMMDDHHMMSSZ ASN1 UTCTime` 格式表示

- `-days arg`

  the number of days to certify the certifcate for

- `-utf8`

  将 DN 以 uft8 的格式解析，默认 ASCII

## Exmaples

```
#对 CSR 签名生成 CRT
openssl ca -in fd.csr -keyfile fd.prv -out fd.crt
```

