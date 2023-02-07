# OpenSSL Subcommands req

ref
[https://www.feistyduck.com/library/openssl-cookbook/online/openssl-command-line/creating-certificate-signing-requests.html](https://www.feistyduck.com/library/openssl-cookbook/online/openssl-command-line/creating-certificate-signing-requests.html)

[https://www.trustasia.com/doc/how-to-generate-csr-file-by-using-openssl](https://www.trustasia.com/doc/how-to-generate-csr-file-by-using-openssl)

https://www.ibm.com/docs/en/i/7.2?topic=concepts-distinguished-names-dns

https://www.ibm.com/docs/en/i/7.1?topic=concepts-distinguished-name

https://knowledge.digicert.com/generalinformation/INFO1745.html

https://openwrt.org/docs/guide-user/luci/getting_rid_of_luci_https_certificate_warnings

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
req 用于生成 PKCS#10 CSR 或者是 self-signed certificates

## DN

Distinguished Name 也叫 DN, 通常在生成 CSR 时需要提供用于描述证书的信息, 由多组键值对组成

- countryName abbrive C

  国家，例如 US

- stateOrProvinceName abbrive ST

  州或者省

- localityName abbrive L

  城市

- organizationName abbrive O

  集团名

- organizationalUnitName abbrive OU

  公司名

- commonName abbrive CN

  the fully qualified domain name such as `www.baidu.com`

- emailAddress

  作为额外的可选项，可以需要指定

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

- `-passin|passout arg`

  the input file passphrase

- `-noenc | -nodes`

  使用 req 生成 private key 默认都需要指定 passphrase，如果不想指定 passphrase 需要使用该参数。否则会报错

- `-out filename`

  the output filename to write

- `-text`

  prints out the CSR in text format

  CSR 的详细信息

- `-noout`

  不输出密钥但是输出密钥中的详细信息

- `-pubkey`

  prints out the public key

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

  和`-x509`一起使用，指定证书的有效时间，openssl 不支持永久的证书，但是可以设置 100 years 或者更长的时间。默认 30 天

- `-utf8`

  指定使用 uft8 作为编码而不是 ASCII

- `-config filename`

  指定 `openssl-req` 使用的配置文件，主要用于使用配置文件中预设的值

## Configuration

### req section

The configuration options are specified in the `req` section of the configuration file

- distinguished_name

  表示 value 对应的 section 包含生成 CSR 所需的 distinguish names

- x509_extensions

  表示 value 对应的 section 包含当 `-x509` 被使用时所需的 certificate extensions

- prompt

  如果当值为 no 时，生成 CSR 时不会 prompt

- default_bits

  指定生成 private key 的 bit

- utf8

  如果当值为 yes 时，所有内容会被解析成 utf-8 字符编码

- string_mask

  设置匹配的字符编码，一般为 utf8only

- default_md

  指定需要使用的 digest algorithm

- encrypt_key

  如果当值为 no 时，等价与使用 `-noenc` ，即不指定 private key passphrase

- default_keyfile

  指定 private key 存储的位置，会被 `-keyout` 复写

- input_password, output_password

  指定 private key 的 passphrase

### DN and attribute section

DN 通过只包含如下几个值

commonName, countryName, localityName, organizationName, organizationalUnitName, stateOrProvinceName, emailAddress

DN 值具体查看 `man openssl-req DISTINGUISHED NAME AND ATTIRBUTE SECTION FORMAT`

例如下面是 openWRT 中的用于生成 LuCI 自签证书的对应内容，这里可以使用字段缩写也可以使用全称

```
C                   = US
ST                  = VA
L                   = SomeCity
O                   = OpenWrt
OU                  = Home Router
CN                  = luci.openwrt
```

### Configuration Exmaple

```
[ req ]
prompt		           = no
default_bits           = 2048
default_keyfile        = privkey.pem
distinguished_name     = req_distinguished_name
attributes             = req_attributes
req_extensions         = v3_ca

[ req_distinguished_name ]
countryName                    = CN
localityName                   = Zhejiang
organizationalUnitName         = Tbone
commonName                     = cyberpelican.com
emailAddress                   = cyberpelican@hotfix.com

[ req_attributes ]
challengePassword              = A challenge password
challengePassword_min          = 4
challengePassword_max          = 20

[ v3_ca ]
subjectKeyIdentifier=hash
authorityKeyIdentifier=keyid:always,issuer:always
basicConstraints = critical, CA:true
```

## Examples

```
#使用 key.pem 私钥生成 req.pem csr
openssl req -new -key key.pem -out req.pem

#生成一个新的 key.pem 私钥，使用 key.pem 生成 req.pem csr
openssl req -newkey rsa:2048 -keyout key.pem -out req.pem

#生成 self-signed root certificate
openssl req -x509 -newkey rsa:2048 -keyout fd.pem -out fd.crt

#以非交互式生成 req.pem csr
openssl req -newkey rsa:2048 -keyout key.pem -out fd.crt -nodes -subj "/C=GB/L=London/O=Feisty Duck Ltd/CN=www.feistyduck.com"

#以非交互式生成 fd.crt 自签证书
openssl req -x509 -newkey rsa:2048 -keyout key.pem -out fd.crt -nodes -subj "/C=GB/L=London/O=Feisty Duck Ltd/CN=www.feistyduck.com"

#以非交互式使用 myconfig 生成 fd.crt 自签证书, 这里的 myconfig 使用 Configuration Exmaple 中的内容
openssl req -x509 -newkey rsa:2048 -keyout mycert.key -out mycert.crt -nodes -config myconfig
```

