# OpenSSL Subcommands x509

## Digest

syntax

```
openssl x509 [-help] [-in filename|uri] [-passin arg] [-new] [-x509toreq] [-req] [-copy_extensions arg] [-inform DER|PEM] [-vfyopt nm:v] [-key
filename|uri] [-keyform DER|PEM|P12|ENGINE] [-signkey filename|uri] [-out filename] [-outform DER|PEM] [-nocert] [-noout] [-dateopt] [-text]
[-certopt option] [-fingerprint] [-alias] [-serial] [-startdate] [-enddate] [-dates] [-subject] [-issuer] [-nameopt option] [-email] [-hash]
[-subject_hash] [-subject_hash_old] [-issuer_hash] [-issuer_hash_old] [-ext extensions] [-ocspid] [-ocsp_uri] [-purpose] [-pubkey] [-modulus]
[-checkend num] [-checkhost host] [-checkemail host] [-checkip ipaddr] [-set_serial n] [-next_serial] [-days arg] [-preserve_dates] [-subj arg]
[-force_pubkey filename] [-clrext] [-extfile filename] [-extensions section] [-sigopt nm:v] [-badsig] [-digest] [-CA filename|uri] [-CAform
DER|PEM|P12] [-CAkey filename|uri] [-CAkeyform DER|PEM|P12|ENGINE] [-CAserial filename] [-CAcreateserial] [-trustout] [-setalias arg] [-clrtrust]
[-addtrust arg] [-clrreject] [-addreject arg] [-rand files] [-writerand file] [-engine id] [-provider name] [-provider-path path] [-propquery
propq]
```

x509 用于生成 x509 格式的证书

## Optional args

### Input, output, and general purpose options

- `-in filename|url`

  读取 filename private key，如果和 `-req` 一起使用表示读取 CSR

- `-passin arg`

  指定 private key passphrase

- `-x509toreq`

  输出证书对应的 x509 CSR，必须指定 `-key` 标识自签使用的私钥

- `-req`

  标识 `-in` 输入的文件是 CSR

- `-inform DER|PEM`

  the input file format

- `-key|signkey filename|url`

  the private key for signing a new certificate or certificate request

- `-keyform DER|PEM|P12|ENGINE`

  指定私钥使用的格式

- `-outform DER|PEM`

  the output file format, 默认 PEM

- 