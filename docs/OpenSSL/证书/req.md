# req

参考：

https://www.trustasia.com/doc/how-to-generate-csr-file-by-using-openssl

按照PKCS#10生成证书请求或CA根证书

- `-config <filename>`

  指定读取的配置文件，默认`/etc/ssl/openssl.cnf`

- -keyform PEM | DER

- -new

  生成证书请求，如果没有通过`-key`指定priate key，会自动生成一个加密的RSA privkey.pem用作申请证书请求的私钥，需要妥善保管。==注意这里如果使用enter不是留空值，如果想要留空值使用`.`==

  ```
  root@ubuntu18:/opt/ssl# openssl req -new -key rsa.pk -passin pass:1234
  You are about to be asked to enter information that will be incorporated
  into your certificate request.
  What you are about to enter is what is called a Distinguished Name or a DN.
  There are quite a few fields but you can leave some blank
  For some fields there will be a default value,
  If you enter '.', the field will be left blank.
  ```

- -newkey

  按照指定规则生成私钥

  ```
  openssl req -newkey rsa:512 -passin pass:1234
  ```

- -keyout  filename

  指定自动生成私钥存储的文件

- -digest

  指定摘要使用的算法

- -subj arg

  设置subjectname

  ```
  -subj "/C=CN/ST=sh/L=sh/O=亚数信息科技/OU=IT/CN=demo.trustasia.com"
  ```

- -verify

  校验签名

## 例子

生成证书请求

```
root@ubuntu18:/opt/ssl# openssl req -new -sha256 -newkey rsa:2048 -nodes -keyout demo.trustasia.com.key -out demo.trustasia.com.csr -subj "/C=CN/ST=sh/L=sh/O=亚数信息科技/OU=IT/CN=demo.trustasia.com"
```

使用特定的digest加密请求

```
root@ubuntu18:/opt/ssl# openssl req -new -sha1 -key rsa.pk -passin pass:1234 -out rsa.csr
```

校验csr文件是否正确

```
root@ubuntu18:/opt/ssl# openssl req -in rsa.csr -verify -noout
verify OK
```



