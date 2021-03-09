# openssl 证书

参考：

https://blog.csdn.net/liuchunming033/article/details/48470575

https://juejin.cn/post/6844903977025273863

https://help.aliyun.com/document_detail/109827.html?spm=a2c4g.11186623.6.625.40b24ff1WNaTFy

https://www.barretlee.com/blog/2016/04/24/detail-about-ca-and-certs/

## 术语

- X.509：最普遍的认证系统，通常也被叫做PKI(public key infrastructure)

- 证书请求：certificate signing request(==csr==)通常包含申请人的公钥和申请人提交的信息。

- CA：Certificate Authority，证书的权威机构，使用CA的私钥对证书签名然后颁发给申请人。客户端信任这些CA，就会在本地保存这些CA的根证书(root certificate)，包含CA的==公钥==用于校验CA颁发的证书的有效性。

  在windows上可以通过`certmgr.msc`服务查看存储的根证书

  > 任何个体组织都可以扮演CA，但是不能得到浏览器的默认信任(==需要将其加入到系统的受信任的根证书颁发机构中==)。知名的CA有Symantec，Comodo，Godaddy，GolbalSign，Digicert

  ![](D:\asset\note\imgs\_openssl\Snipaste_2021-03-03_10-03-17.png)

- 证书：由CA颁发的certificate(==crt |  cer==)通常包含版本，签名哈希算法，签名算法，证书申请人的公钥，证书序列号，有效期，颁发者。==需要由CA的根证书对证书签名(非对称加密)==。

  ![](D:\asset\note\imgs\_openssl\Snipaste_2021-03-03_11-36-32.png)

  > 签名哈希算法(对称加密)，签名算法(非对称加密+cipher)。都是CA对证书的签名和摘要。用户可以通过哈希算法对明文部分进行对称解密，然后通过CA的公钥来解密，对比是否相同来校验证书的有效性。

  下面是证书的一个示例：

  ```
  -----BEGIN CERTIFICATE-----
  MIIEqjCCA5KgAwIBAgIQAnmsRYvBskWr+YBTzSybsTANBgkqhkiG9w0BAQsFADBh
  MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
  d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD
  QTAeFw0xNzExMjcxMjQ2MHRpb24gRXZlcnl3aGVyZSBEViBUTFMgQ0EgLSBHCCAS
  oHqt3jRIxW5MDvf9QyiOR7VfFwK656es0UFiIb74N9pRntzF1UgYzDGu3ppZVMdo
  lbxhm6dWS9OK/lFehKNT0OYI9aqk6F+U7cA6jxSC+iDBPXwdF4rs3KRyp3aQn6pj
  pp1yr7IB6Y4zv72Ee/PlZ/6rYR7n9iDuPe1E4IxUMBH/T33+3hAU8wggFLMB0GA1
  BgNVHSMEGDAWgBQD3lA1VtFMu2bwo+IbG8OXsj3RVTAOBgNVHQ8BAf8EBAMCAYYw
  HQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMBIGA1UdEwEB/wQIMAYBAf8C
  AQAwNAYIKwYBBQUHAQEEKDAmMCQGCCsGAQUFBzABhhhodHRwOi8vb2NzcC5kaWdp
  Y2VydC5jb20wQgYDVR0fBDswOTA3oDWgM4YxaHR0cDovL2NybDMuZGlnaWNlcnQu
  gNVHSAERTBDMDcGCWCGSAGGhxodHRwczovL3d3dy5kaWdpY2VydC5jb20vQ1BTAg
  M0r5LUvStcr82QDWYNPaUy4taCQmyaJ+VB+6wxHstSigOlSNF2a6vg4rgexixeiV
  4YSB03Yqp2t3TeZHM9ESfkus74nQyW7pRGezj+TC44xCagCQQOzzNmzEAP2SnCrJ
  sNE2DpRVMnL8J6xBRdjmOsC3N6cQuKuRXbzByVBjCqAA8t1L0I+9wXJerLPyErjy
  rMKWaBFLmfK/AHNF4ZihwPGOc7w6UHczBZXH5RFzJNnww+WnKuTPI0HfnVH8lg==
  -----END CERTIFICATE-----
  ```

  ==可以使用`openssl x509 -in <crt> -noout -text`来查看证书中的具体内容==

- 证书链：chain of trust(_chain.crt)，类似于DNS查询的方式按照`_chain.crt`文件递归校验

  ![](D:\asset\note\imgs\_openssl\Snipaste_2021-03-03_11-12-07.png)

- pfx：以PKCS#12格式包含X.509证书和私钥(可选)受密码保护二进制文件，可以被导入到系统中，需要导出时的密码

  <img src="..\..\imgs\_openssl\Snipaste_2021-03-03_10-49-46.png" style="zoom:80%;" />

- PKCS#*：最主要的有PKCS#12(标准加密格式用于证书)，PKCS#10(证书请求的加密格式)，PKCS#7(通用加密格式)

  参考：https://en.wikipedia.org/wiki/PKCS

## 证书分类

https://developer.aliyun.com/article/238048

1. DV：简易型SSL证书，起到加密传输的作用，但无法向用户证明网站的真实身份。支持单域名，多域名；不支持泛域名，泛多域名。免费
2. OV：能提供加密传输，真实身份校验。支持单域名，多域名，泛域名，泛多域名。收费
3. EV：最安全SSL证书，同时提供身份校验。支持单域名，多域名；不支持泛域名，泛多域名。收费

## 证书认证过程

- 请求证书

  1. 申请人生成csr
  2. CA校验申请人的信息，如果信息通过颁发证书

  3. 申请人将证书部署在服务器上

- 认证

  ==具体抓包[TLS](../Net/SSL & TSL)分析==

  1. 客户端请求服务器
  2. 服务端将证书返回给客户端(证书中包含申请证书时的公钥)
  3. 客户端获取到证书后==对证书中的明文信息通过指定的哈希算法对称解密，然后使用客户端上指定CA的公钥对密文部分非对称解密==。比对两者是否有差异。

## 证书申请

### req

按照PKCS#10生成证书请求或CA根证书

- `-config <filename>`

  指定读取的配置文件，默认`/etc/ssl/openssl.cnf`

- `-new`

  生成证书请求，如果没有通过`-key`指定priate key，会自动生成一个加密的privkey.pem用作申请证书请求的私钥，需要妥善保管。==注意这里如果使用enter不是留空值，如果想要留空值使用`.`==

  ```
  root in /usr/local/\/ssl λ openssl req -new 
  Generating a RSA private key
  ..............................................+++++
  ...........+++++
  writing new private key to 'privkey.pem'
  Enter PEM pass phrase:
  Verifying - Enter PEM pass phrase:
  -----
  You are about to be asked to enter information that will be incorporated
  into your certificate request.
  What you are about to enter is what is called a Distinguished Name or a DN.
  There are quite a few fields but you can leave some blank
  For some fields there will be a default value,
  If you enter '.', the field will be left blank.
  -----
  Country Name (2 letter code) [AU]:
  State or Province Name (full name) [Some-State]:
  Locality Name (eg, city) []:
  Organization Name (eg, company) [Internet Widgits Pty Ltd]:
  Organizational Unit Name (eg, section) []:
  Common Name (e.g. server FQDN or YOUR name) []:
  Email Address []:
  
  Please enter the following 'extra' attributes
  to be sent with your certificate request
  A challenge password []:
  An optional company name []:
  -----BEGIN CERTIFICATE REQUEST-----
  MIICijCCAXICAQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUx
  ITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDCCASIwDQYJKoZIhvcN
  AQEBBQADggEPADCCAQoCggEBAMDs80tWijhX8VG6/M3k1cwLd04uwWL+1+d1h2IM
  GYWVuwtUNXRm1D9CBYV6P4LFUkuodn9ObuMf1VqFWf1ieGMupgqSNCWCze+YW8Yr
  5OTz4Kr8x13wYiPiV4Lq4zOJjadX9XQeQ3dZV1lSMbzSTjoGhRiIH4fHlmqo1KDs
  ...
  ```

- `-key <private_key>`

  申请证书请求的私钥或是CA的私钥，同样也接收PKCS#8格式的pem文件

- `-in <csr>`

  读取证书请求

- `-x509`

  ==使用该参数生成自用的证书，而不是证书请求==。通常用于生成测试的证书或自签CA根证书。需要与`-in`一起使用才会生成自签的证书，否则还是生成证书请求

  ```
  root in /usr/local/\/ssl λ openssl req -in req -x509 -key privkey.pem 
  Enter pass phrase for privkey.pem:
  -----BEGIN CERTIFICATE-----
  MIIDazCCAlOgAwIBAgIUUQVz0X3JHzU7pw+X6HxTekk06fkwDQYJKoZIhvcNAQEL
  BQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
  GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yMTAzMDMwNTM5NDNaFw0yMTA0
  MDIwNTM5NDNaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
  HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEB
  fkYCjdE336TU8ZneNsWu97/OrzNhQBGeiGmxHMr1yPP/D5a1JPQSkZnrMm9v4KZD
  AJq1DpdooCqpT0VZwxKmISosmbiwinaO3yFeqtKYG0vVkARwoF7i/dfoOJEpgS4/
  vktxg12mFVPtrDK+w+Z4taiDK915fjVOK/Kld/r9DphwBfrDobsoo9qij5FQXBy9
  cbcQ0DIi0RtU3MdGvmD/6rVctMr7ZM2PtfjHe5J1ar4XJApStyqz/tiurQ9QRu/c
  soshMQSS8pgswMwWaJocgDTdse7WvtF8grL6MIbsAXH4ULdISj+Dkg4psGKBMqAx
  XPbpNRS630t297yn/EgR2I9m6WVWLtYoN4JjdARlV0Xua34Xz8wXVlLK2zO+DpU3
  V6HKLbk82KJVgTNHJ10+
  -----END CERTIFICATE-----
  ```

  这里的`-key`被用作CA的私钥

- `-days <n>`

  需要使用`-x509`一起使用，否则没有意义。表示证书的有效期，默认30天

- `-out <file>`

  将生成的证书或证书请求输出到指定文件

- `-keyout <file>`

  指定`-newkey`生成私钥写入的文件

  ```
  root in /usr/local/\/ssl λ openssl req -newkey rsa:512 -keyout rsa.pk -out rsa_req.pem
  ```

  这里的`-out`生成的时csr文件

- `-verify`

  校验csr文件是否正确

  ```
  root in /usr/local/\/ssl λ openssl req -in req -verify 
  verify OK
  -----BEGIN CERTIFICATE REQUEST-----
  MIICijCCAXICAQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUx
  ITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDCCASIwDQYJKoZIhvcN
  AQEBBQADggEPADCCAQoCggEBAK9Dm47C1id7zWtIcM10vlLLQYDKw7C6TdtKrpLo
  yGfHaFzCJuQ+EIvUduxpxZv/IWFSQegdkYRXkh7pfAJ39eyiu8zhfavYf++LV0vo
  jfx3cbmT2hlBu1neE/rIY4y0og4M4cI/HWLSUn9yE72ROrOQV6/tnIn0K0ux6coA
  jBxpgZr019dFFtMU+j5ebUmuV7sMCmNUbD997CJ3gB4ntq2+EpWVPW7JypZVTl2e
  47svQefsHtA6psWyVeA/lHXx8FbduHqxLL9ux3/rZkkbI1rwoRLG+rY8+L2dh4wJ
  rieV8pn/52S6tqHT72uT2sdZK3qoI5LII8logi8k
  -----END CERTIFICATE REQUEST-----
  ```

- `-newkey <algorithm:nbit>`

  生成证书请求和私钥，使用`rsa:512`格式表示使用的加密算法和位数

### ca

小型ca

- `-config <filename>`

  指定读取的配置文件，默认`/etc/ssl/openssl.cnf`

- `-in <filename>`

  读取的csr文件，==会读取默认配置文件中CA的私钥==。

- `-infiles [file...]`

  同时对多个csr文件签名

  ```
  openssl ca -infiles req1.pem req2.pem req3.pem
  ```

- `-out <filename>`

  生成证书写入的文件，默认以PEM格式，可以使用`-spkac`输出DER格式。

  ```
  root in /usr/local/\/ssl λ openssl ca -in req  -out req.crt
  ```

- `-ss_cert <filename>`

- `-keyfile <private_key>`

  用于签名的私钥，否则默认会读取配置文件中指定的私钥

### x509

用于对证书的格式化，和`ca`一样为证书请求签发证书

- `-in <filename>`

  读取指定的证书文件

- `-outform <DER|PEM>`

  指定输出证书的格式

  ```
  root in /usr/local/\/ssl λopenssl x509 -in req.crt -outform DER
  0��0�*�H���P�����T\ʽ,U��0
  0E1
     0	UAU10U
  
  Some-State1!0U
  
  210402063348Z0E1 Pty Ltd0
                  0	UAU10U
  
  Some-State1!0U
  
  �0�net*�H�� Pty Ltd0�"0
  ��C����'{�kHp�t�R�A��ð�M�J����g�h\�&�>��v�iś�!aRA���W��|w�좻��}���WK�a�8�O�
  ���xsa��Ӿ�3O�	��\�ȿ� ��s�Q-I�F&�h�R�Pp���($\�qn�Щ�._3kt�jFu�:�2�_����zS����K~F��7ߤ���6Ů��ί3a@��i��������$����2oo4w��y *�H�� �S0Q0U$�,��r$DS��� ���0U#0�$�,��r$DS��� ���
  ```

- `-text`

  输出所有的明文信息

  ```
  root in /usr/local/\/ssl λ openssl x509 -in crt -text -noout                          
  Certificate:
      Data:
          Version: 3 (0x2)
          Serial Number:
              34:c9:8c:41:cd:63:76:14:f8:bc:30:e3:03:a4:28:47:e7:c4:a1:8e
          Signature Algorithm: sha256WithRSAEncryption
          Issuer: C = AU, ST = Some-State, O = Internet Widgits Pty Ltd
          Validity
              Not Before: Mar  3 06:44:50 2021 GMT
              Not After : Apr  2 06:44:50 2021 GMT
          Subject: C = AU, ST = Some-State, O = Internet Widgits Pty Ltd
          Subject Public Key Info:
              Public Key Algorithm: rsaEncryption
                  RSA Public-Key: (512 bit)
                  Modulus:
                      00:9a:93:26:74:f2:ba:4e:eb:da:4c:45:1d:f1:17:
                      b9:88:f5:a7:cb:df:42:b9:ed:e6:07:79:04:c4:11:
                      91:0f:99:92:b8:af:50:20:06:bc:de:f5:ef:66:0c:
                      4d:2c:b3:2d:09:f1:b7:9d:35:53:16:f3:4d:3c:d2:
                      bb:58:a4:31:23
                  Exponent: 65537 (0x10001)
          X509v3 extensions:
              X509v3 Subject Key Identifier: 
                  8D:E9:B0:41:61:53:A6:D6:EE:E5:F4:9C:34:3F:C3:98:1C:F5:FD:A9
              X509v3 Authority Key Identifier: 
                  keyid:8D:E9:B0:41:61:53:A6:D6:EE:E5:F4:9C:34:3F:C3:98:1C:F5:FD:A9
  
              X509v3 Basic Constraints: critical
                  CA:TRUE
      Signature Algorithm: sha256WithRSAEncryption
           11:84:d1:38:2a:9e:d2:83:11:30:9e:bc:38:2b:23:4c:a7:fa:
           e9:fa:d9:2a:76:ea:14:5d:52:7c:3d:d4:d2:80:d7:8f:2a:8b:
           82:a1:3e:79:c0:23:ff:fb:43:c8:5c:80:22:e6:96:7d:79:06:
           a9:16:98:41:55:74:1f:ca:82:a1
  ```

- `-subject | -issuer | -email | -enddate | -serial | ...`

  输出证书上的明文信息，具体查看命令

  ```
  root in /usr/local/\/ssl λ openssl x509 -enddate -in crt -noout
  notAfter=Apr  2 06:44:50 2021 GMT
  ```

- `-digest`

  以指定摘要算法显示信息，可以通过`openssl list -digest-algorithms`来查看

  ```
  root in /usr/local/\/ssl λ openssl x509 -sha1 -in crt -noout -fingerprint  
  SHA1 Fingerprint=F9:8F:E3:B4:33:09:19:72:CF:46:89:04:47:39:30:E5:92:EC:90:1B
  ```

- `-req` 

  `-in <csr>`输入的是一个csr文件

  `-signkey <key>`指定签名的私钥

  这三个参数一起使用表示使用key对csr生成证书

  ```
  root in /usr/local/\/ssl λ openssl x509 -req -days 30 -in tmp.csr -out tmp.new -signkey etter.ssl.crt
  Signature ok
  subject=C = AU, ST = Some-State, O = Internet Widgits Pty Ltd
  Getting Private key
  root in /usr/local/\/ssl λ ls
   etter.ssl.crt   tmp.csr   tmp.new
  root in /usr/local/\/ssl λ cat tmp.new
    File: tmp.new
    -----BEGIN CERTIFICATE-----
    MIIBhzCCATECFCPCOcTel51a9hyPylZQMxVAqra5MA0GCSqGSIb3DQEBCwUAMEUx
    CzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRl
    cm5ldCBXaWRnaXRzIFB0eSBMdGQwHhcNMjEwMzA0MDIxOTEyWhcNMjEwNDAzMDIx
    OTEyWjBFMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UE
    CgwYSW50ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMFwwDQYJKoZIhvcNAQEBBQADSwAw
    SAJBAMY8ozUVsAO6Z6qjnu19VsticVO+fdMBzw+wUqbO1KqJ4ByE0Lmw4OdxoqrJ
    zfJKUcIkXwerraO/a3CsELI5eGsCAwEAATANBgkqhkiG9w0BAQsFAANBALeeozIL
    QiN2ZS05olbVgRB1xWHKp15WLsq1JcgxmkoCURX+WFyc8lBq7tqEn8mUSot4uBrr
    mTbEo0FRzcLSlgA=
    -----END CERTIFICATE-----
  ```

- `-CA <filename>`

  指定用于签名的CA证书(通常CA证书中有私钥)

- `-CAkey <filename>`

  指定用于签名的CA私钥，如果没有指定就会从CA的证书中找私钥

- `-CAcreateserial`

  如果没有CA的序列号文件，会生成

  ```
  root in /usr/local/\/ssl λ openssl x509 -req -in req -CA crt -CAkey pk -CAcreateserial
  Signature ok
  subject=C = AU, ST = Some-State, O = Internet Widgits Pty Ltd
  Getting CA Private Key
  Enter pass phrase for pk:
  -----BEGIN CERTIFICATE-----
  MIIBhzCCATECFERAJNJjIAD6ArXvDCt2uttMFImBMA0GCSqGSIb3DQEBCwUAMEUx
  CzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRl
  cm5ldCBXaWRnaXRzIFB0eSBMdGQwHhcNMjEwMzAzMDcwNjE5WhcNMjEwNDAyMDcw
  SAJBAJqTJnTyuk7r2kxFHfEXuYj1p8vfQrnt5gd5BMQRkQ+ZkrivUCAGvN7172YM
  TSyzLQnxt501UxbzTTzSu1ikMSMCAwEAATANBgkqhkiG9w0BAQsFAANBABNajDDa
  vBh2BqOG0ljWXQea7hnxv9ei5aYDbZNYS+ngEqIQBXvbO27bCsAXBhFCtI/BwCG2
  lpopAr1vGLM3XFE=
  -----END CERTIFICATE-----
  ```

  















