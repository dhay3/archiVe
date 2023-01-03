## Digest
syntax
```
openssl pkey [-help] [-inform PEM|DER] [-outform PEM|DER] [-in filename]
[-passin arg] [-out filename] [-passout arg] [-traditional] [-cipher]
[-text] [-text_pub] [-noout] [-pubin] [-pubout] [-engine id] [-check]
[-pubcheck]
```
genpkey 用于生成 private key
## Optional args

- `-inform DER|PEM`

input format 是 DER 还是 PEM，默认 PEM

- `-outform DER|PEM`

output format 是 DER 还是 PEM，默认 PEM

- `-in filename`

the input filename to read a key

- `-pubin`

默认输入 private key，如果指定该参数输入 public key

- `-passin arg`

the input file passphrase

- `-[cipher]`

指定 private key 使用的 cipher

- `-text`

以格式化的形式输出 private key 或者 public key 使用的参数和 PEM 文件

- `-text_pub`

以格式化的形式只输出 public key 部分内容

- `-noout`

do not output the encoded version of the key

- `-out filename`

the output filename to write a key。==这里需要注意的是 output filename 不能和 input filename 一样，否则秘钥会被清空==

- `-pubout`

默认输出 private key，如果指定该参数输出 public key。如果输入的是 public key 默认会使用该参数
## Exmaples
```
#输出私钥对应的公钥
openssl pkey -in key.pem -pubout -out pubkey.pem

#取消私钥使用的 passphrase
openssl pkey -in key.pem -out keyout.pem

#将私钥从 PEM 转成 DER 格式
openssl pkey -in key.pem -outform DER -out keyout.der
```
