# rsa

syntax：`openssl rsa [options]`

用于生成公钥和格式化私钥

## options

- -inform DER|PEM

  指定输入文件的格式

- -outform DER|PEM

  指定输出文件的格式

- -out filename

  输出到指定文件

- -in filename

  指定输入的文件，默认标识私钥

- `-<cipher>`

  使用指定的cipher对私钥加密

- -pubin

  指定输入公钥，默认输入私钥

- -pubout

  指定输出公钥，默认输出私钥

- -passin

  输入文件密码

- -check

  输出前，检查rsa私钥

- ==-noout==

  不在stdout上输出

## 例子

根据私钥生成公钥

```
root@ubuntu18:/opt/ssl# openssl rsa -in rsa.pk -pubout
Enter pass phrase for rsa.pk:
writing RSA key
-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBALmiJw8hCmjvrW8y9LTjw7f15XalDjEV
myT6E6+LXL8peyT2t4jQU2d5C6vYJzR3Q+b63IFHzkyJy1LxWCkmGqsCAwEAAQ==
-----END PUBLIC KEY-----


root@ubuntu18:/opt/ssl# openssl rsa -in rsa.pk -passin pass:1234 -pubout
writing RSA key
-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBALmiJw8hCmjvrW8y9LTjw7f15XalDjEV
myT6E6+LXL8peyT2t4jQU2d5C6vYJzR3Q+b63IFHzkyJy1LxWCkmGqsCAwEAAQ==
-----END PUBLIC KEY-----
```

格式化私钥，同样也能格式化公钥

```
root@ubuntu18:/opt/ssl# openssl rsa -in rsa.pk -outform DER -out rsa.der
Enter pass phrase for rsa.pk:
writing RSA key
0▒<A▒▒'!
h▒o2▒▒÷▒▒v▒1▒$▒▒▒\▒){$▒▒Sgy
```

















