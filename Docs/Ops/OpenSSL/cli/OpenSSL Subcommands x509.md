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

- `-out filename`

  the output filename

- `-x509toreq`

  根据 x509 certificate 生成 CSR，必须和 `-key|signkey` 一起使用

### Certificate print options

- `-nocert`

  默认会输出证书，使用该参数不输出证书，只输出参数指定的内容

- `-noout`

  只输出参数指定的内容

  ```
  openssl x509 -serial -in fd.crt -nocert
  serial=7A43B2C2FA5409CA7EE1A0033557539A2A0BA008
  ```

- `-text`

  输出证书的详细信息

  ```
  openssl x509 -in fd.crt -nocert -text
  ```

- `-figerprint`

  输出证书的指纹信息

- `-serial`

  输出证书的序列号

- `-startdate | enddate | dates`

  输出证书的有效的开始和结束时间，dates = startdate + enddate

  ```
  openssl x509  -in fd.crt -noout -startdate -enddate
  notBefore=Jan 11 13:11:33 2023 GMT
  notAfter=Feb 10 13:11:33 2023 GMT
  
  openssl x509  -in fd.crt -noout -dates  notBefore=Jan 11 13:11:33 2023 GMT
  notAfter=Feb 10 13:11:33 2023 GMT
  ```

- `-subject`

  输出证书 DN 信息 

- `-issuer`

  输出证书的颁发者信息

- `-email`

  prints the email address if any

- `-pubkey`

  输出证书的公钥

### Certificate checking options

- `-checkend arg`

  校验证书是否会在 arg 秒内过期

- `-checkhost host`

  校验证书是否匹配 host

  ```
  openssl x509  -in fd.crt -noout -checkhost  www.feistyduck.com
  Hostname www.feistyduck.com does match certificate
  ```

### Micro-CA options

- `-CA filename|uri`

  指明用于签名的 CA certificate

- `-CAform DER|PEM|P12`

  CA certifcate 的格式

- `-CAKey filename|uri`

  指明用于签名的 CA private key 必须匹配 `-CA` 标识的 certificate 

## Examples

```
#输出 fd.crt 证书的详细信息
openssl x509 -in fd.crt -noout -text

#将 fd.crt 证书从 PEM 格式转成 DER 格式
openssl x509 -in fd.crt -inform PEM -out fd.der -outform DER

#按照 fd.crt 证书生成 fd.csr CSR
openssl x509 -x509toreq -in fd.crt -out fd.csr

#按照 fd.csr 使用 key.pem 私钥 生成自签 fd.crt 证书
openssl x509 -req -in fd.csr  -key key.pem -out fd.crt
```

