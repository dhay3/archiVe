# openssl 文件格式

https://www.openssl.org/docs/manmaster/man1/openssl-format-options.html

https://www.shangyang.me/2017/05/24/encrypt-rsa-keyformat/

https://blog.51cto.com/wushank/1915795

https://segmentfault.com/a/1190000016249967

文件不一定以某种格式结尾，按照文件内容区别

- DER/distinguished encoding rules

  ASN.1数据加密后的二进制文件，密钥原始的格式但是一般都会以pem格式展示

- PEM/privacy enhanced mail 

  base64编码后的文件。被编码的文件只能是证书，密钥或证书请求

  ```
  ----BEGIN PRIVATE KEY-----
    MIIFHDBOBgkqhkiG9w0BBQ0wQTApBgkqhkiG9w0BBQwwHAQIBuNYFjKVCZ0CAggA
    MAwGCCqGSIb3DQIJBQAwFAYIKoZIhvcNAwcECPMFLXmqQ/0uBIIEyJOJv78Izytj
    qy7qwSPn7pEEoeDMKvPujlWM5mHR+SME175+Q7S+210VydikE8ZA3aevoXN9aDv2
  -----END PRIVATE KEY-----
  ```

- csr

  Certificate Signing Request，证书签名请求。向权威证书颁发机构获得签名证书的申请，核心是一个公钥。具体查看`req`

  ```
  -----BEGIN CERTIFICATE REQUEST-----
  ...
  -----END CERTIFICATE REQUEST-----
  ```

- cer | crt

  certificate，证书。可能是PEM或DER格式

  ```
  -----BEGIN CERTIFICATE-----
  ...
  -----END CERTIFICATE-----
  ```

- key

  密钥文件

  可以通过`openssl list -public-key-algorithms`来查看支持的非对称加密

  ```
  ----BEGIN RSA PRIVATE KEY-----
  ...
  -----END RSA PRIVATE KEY-----
  ```

