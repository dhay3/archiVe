# openssl ca

小型CA认证授权

## flags

- `-selfsign`

- `-notext`

  不将证书输出在stdout

- `-batch`

  batch mode，不会询问直接自动生成证书

- `-create_serial`

  如果配置文件的`${CA_default::serial}`读取失败，就会创建一个随机的serial。如果配置文件中有就用配置文件中的

## options

- `-config <filename>`

  指定读取的配置文件，默认`/etc/ssl/openssl.cnf`

- `-in <filename>`

  读取的csr文件

- `-infiles [file...]`

  同时对多个csr文件签名，必须在最后使用

  ```
  openssl ca -infiles req1.pem req2.pem req3.pem
  ```

- `-out <filename>`

  生成证书写入的文件，默认以PEM格式，可以使用`-spkac`输出DER格式。

- `-cert <CA root cer>`

  指定用于签名的CA根证书

- `-keyfile <private_key>`

  用于签名的CA私钥，否则默认会读取配置文件中指定的私钥

- `-keyform PEM|DER`

  指定CA私钥的格式，默认PEM

- `-cert <ca_cert>`

  指定用于签发证书的ca根证书

- `-startdate | -enddate`

  指定证书有效的日期，以YYYYMMDDHHMMSSZ格式

- `-days`

  指定有效的天数，从签发起开始算

- `-md <dgst>`

  指定摘要的加密算法，可以通过`openssl list --digest-algorithms`查看

- 