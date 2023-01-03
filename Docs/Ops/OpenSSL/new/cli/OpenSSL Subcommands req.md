ref
[https://www.feistyduck.com/library/openssl-cookbook/online/openssl-command-line/creating-certificate-signing-requests.html](https://www.feistyduck.com/library/openssl-cookbook/online/openssl-command-line/creating-certificate-signing-requests.html)
[https://www.trustasia.com/doc/how-to-generate-csr-file-by-using-openssl](https://www.trustasia.com/doc/how-to-generate-csr-file-by-using-openssl)
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

- `-out filename`

the output filename to write

- `-text`

prints out the CSR in text format

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

-  `-[digest]`

指定生成 CSR 的签名算法

- `-subj arg`

设置 CSR subjectname, 必须以`/type0=value0/type1=value1/type2=....`格式

- `-x509`

输出 self signed certificates 而不是 CSR

- `-days n`

和`-x509`一起使用，指定证书的有效时间默认 30 天

- `-utf8`

指定使用 uft8 作为编码而不是 ASCII


