# rsautil

syntax：`openssl rsautil [options]`

sign, verify, encrypt and decrypt data using the RSA algorithm

## options

- -in file

  用于加密、解密、签名或是签名校验的文件。如果没有该参数从stdin中读入

- out file

  输出到文件

- -inkey file

  指定读入的key，默认rsa 私钥

- -pubin

  读入的是rsa 公钥

- -sign

  对数据签名，需要私钥。==默认会对in file进行hash计算==

- -verify

  校验签名并恢复数据

- -encrypt

  用公钥加密

- -decrypt

  用私钥解密

- -hexdump

  以hex格式输出

## 例子

对数据签名

```
root@ubuntu18:/opt/ssl# openssl rsautl -in data -inkey rsa.pk -out data.sign -sign
Enter pass phrase for rsa.pk:
root@ubuntu18:/opt/ssl# cat data.sign
▒▒wo$▒0▒▒h▒▒▒3t▒~i▒彤-▒▒▒▒Ͱ^2",C▒K▒▒A▒Zq
                                        ▒▒▒▒1▒
```

校验签名(私钥和公钥都可以校验)

```
root@ubuntu18:/opt/ssl# openssl rsautl -in data.sign -pubin -inkey rsa.pub  -verify
hello world
root@ubuntu18:/opt/ssl# openssl rsautl -in data.sign -inkey rsa.pk  -verify
hello world
```

公钥对数据加密(也可以私钥加密私钥解密)

```
root@ubuntu18:/opt/ssl# openssl rsautl -in data -pubin -inkey rsa.pub -encrypt -out rsa.enc1
root@ubuntu18:/opt/ssl# cat rsa.enc1
▒CX▒d▒▒N▒▒m▒m▒▒Dlq▒▒▒▒[▒▒▒l ▒▒=▒▒ʔ▒\|q▒▒▒▒啼ٲm9Ye▒T
```

私钥对数据解密

```
root@ubuntu18:/opt/ssl# openssl rsautl -in rsa.enc1 -decrypt -inkey rsa.pk
Enter pass phrase for rsa.pk:
hello world
```























