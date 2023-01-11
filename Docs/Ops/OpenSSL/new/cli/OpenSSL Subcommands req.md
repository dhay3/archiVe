# OpenSSL Subcommands req

ref
[https://www.feistyduck.com/library/openssl-cookbook/online/openssl-command-line/creating-certificate-signing-requests.html](https://www.feistyduck.com/library/openssl-cookbook/online/openssl-command-line/creating-certificate-signing-requests.html)

[https://www.trustasia.com/doc/how-to-generate-csr-file-by-using-openssl](https://www.trustasia.com/doc/how-to-generate-csr-file-by-using-openssl)

https://www.ibm.com/docs/en/i/7.2?topic=concepts-distinguished-names-dns

https://www.ibm.com/docs/en/i/7.1?topic=concepts-distinguished-name

https://knowledge.digicert.com/generalinformation/INFO1745.html

## Digest
syntax
```
openssl req [-help] [-inform PEM|DER] [-outform PEM|DER] [-in filename]
[-passin arg] [-out filename] [-passout arg] [-text] [-pubkey] [-noout]
[-verify] [-modulus] [-new] [-rand file...]  [-writerand file] [-newkey
rsa:bits] [-newkey alg:file] [-nodes] [-key filename] [-keyform PEM|DER]
[-keyout filename] [-keygen_engine id] [-digest] [-config filename]
[-multivalue-rdn] [-x509] [-days n] [-set_serial n] [-newhdr] [-addext
ext] [-extensions section] [-reqexts section] [-precert] [-utf8]
[-nameopt] [-reqopt] [-subject] [-subj arg] [-sigopt nm:v] [-batch]
[-verbose] [-engine id]
```
req 用于生成 PKCS#10 CSR

## DN

Distinguished Name 也叫 DN, 通常在生成 CSR 时需要提供用于描述证书的信息, 由多组键值对组成

- Country alias C

  2 character country code such as US

- Locality/City alias L

- State alias S

  must be spelled out completely such as New York

- Organization alias O

  legal company name

- Organizational Unit alias OU

  company department 

- Common Name alias CN

  the fully qualified domain name such as `www.baidu.com`

使用 req 生成 CSR 时，如果想留空需要使用 single dot `.` 而不应该直接回车（使用回车会使用 openssl 预设的值）

## Optional args

- `-inform DER|PEM`

  input format 是 DER 还是 PEM，默认 PEM ( DER format base64 encoded with additional header and footer lines)

- `-outform DER|PEM`

  output format 是 DER 还是 PEM，默认 PEM

- `-in filename`

  读取指定的 CSR，只有在`-new` 和 `-newkey` 没有被指定的情况下生效

- `-key filename`

  读取指定的 private key

- `keyform PEM|DER`

  private key 的格式，默认 PEM

- `-passin arg`

  the input file passphrase

- `-noenc | -nodes`

  使用 req 生成 private key 默认都需要指定 passphrase，如果不想指定 passphrase 需要使用该参数

- `-out filename`

  the output filename to write

- `-text`

  prints out the CSR in text format

- `-noout`

  不输出密钥但是输出密钥中的详细信息

- `-pubkey`

  outputs the public key

- `-verify`

  verifies the signature on the request

- `-new`

  生成新的 CSR，如果没有指定`-key`会生成新的 RSA private key 用于新的 CSR

- `-newkey arg`

  生成新的 CSR 和 private key。arg 的值可以是 `algorithm:params` 例如`rsa:2048`生成 rsa 2048 bit 的私钥

- `-keyout filename`

  写入新生成 private key 的文件名

- `-[digest]`

  指定生成 CSR 的签名算法

- `-subj arg`

  设置 CSR subjectname, 必须以`/type0=value0/type1=value1/type2=....`格式

  ```
  -subj "/C=GB/L=London/O=Feisty Duck Ltd/CN=www.feistyduck.com"
  ```

- `-x509`

  输出 self signed certificates 而不是 CSR

- `-days n`

  和`-x509`一起使用，指定证书的有效时间默认 30 天

- `-utf8`

  指定使用 uft8 作为编码而不是 ASCII

## Examples

```
#使用 key.pem 私钥生成 req.pem csr
openssl req -new -key key.pem -out req.pem

#生成一个新的 key.pem 私钥，使用 key.pem 生成 req.pem csr
openssl req -newkey rsa:2048 -keyout key.pem -out req.pem

#生成 self-signed root certificate
openssl req -x509 -newkey rsa:2048 -keyout key.pem -out req.pem

#以非交互式生成 req.pem csr
openssl req -newkey rsa:2048 -keyout key.pem -out req.pem -nodes -subj "/C=GB/L=London/O=Feisty Duck Ltd/CN=www.feistyduck.com"

#以非交互式生成 fd.crt 自签证书
openssl req -x509 -newkey rsa:2048 -keyout key.pem -out fd.crt -nodes -subj "/C=GB/L=London/O=Feisty Duck Ltd/CN=www.feistyduck.com"
```

