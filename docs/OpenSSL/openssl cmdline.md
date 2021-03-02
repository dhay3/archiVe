# openssl cmdline

## 概述

openssl是一个调用openssl libcrypto api的命令行工具，有主要如下几个功能

- Creation and management of private keys, public keys and parameters
- Public key cryptographic operations
- Creation of X.509 certificates, CSRs and CRLs
- Calculation of Message Digests and Message Authentication Codes
- Encryption and Decryption with Ciphers
- SSL/TLS Client and Server Tests
- Handling of S/MIME signed or encrypted mail
- Timestamp requests, generation and verification

默认的配置文件在`/etc/ssl/openssl.cnf`

## passphrase-options

用于指定私钥的passpharse才会被用到

- `pass:password`

  私钥的密码。不安全因为是密文的，所以可以被`ps`查看到

- `env:var`

  私钥的密码是环境变量

- `file:pathname`

  文件中的第一行是密码

- `fd:number`

  指定私钥的密钥是文件句柄

- `stdin`

  从stdin中输入密码

## subCommands

> 使用`openssl list`可以查看openssl中自带的命令，所有的subCommands都自带`-help`

- `ca`

  Certificate Authority (CA) Management.

- `ciphers`

- `cms`

- `genpkey`

  生成私钥

- `rsa`

  rsa密钥管理中心

- `md5`

  md5相关模块

- `sha1 | sha224 | sha256 | sha512`

### genrsa

syntax：`openssl genrsa [options] [numberbit]`

用于生成rsa私钥。没有任何参数，默认将生成的私钥输出到stdout。numberbit表示生成密钥的长度默认2048不能小于512

- `-out <outfile>`

  生成一个rsa私钥到指定文件

  ```
  openssl genrsa -out outfile
  ```

- `-writerand <outfile>`

  往文件中写入随机的数据并将私钥输出到stdout

- `-passout <val>`

  指定生成私钥的passpharse，val的值需要重paspharse-options取

  ```
  root in /usr/local/\ λ openssl genrsa -passout pass:toor
  ```

- `-cipher`

  指定生成密钥的加密算法

  ```
  root in /usr/local/\/ssl λ openssl genrsa -des 
  ```

### rsa

用于对私钥的格式化和生成公钥

syntax：`openssl rsa [options]`

- `-in <filename>`

  输出密钥，如果密钥有passpharse会询问passpharse。如果没有指定会从stdin中读入

- `-pubin`

  默认读取私钥，使用该标志表示读取公钥

- `-out <filename>`

  输出密钥，如果没有会输出到stdout

- `-pubout`

  输出公钥

- `-passin arg`

  指定输入私钥的passpharse

- `-passout`

  指定输出私钥的paspharse

- `-cipher`

  使用指定的cipher对输出加密，具体支持的参数查看`man opensll-rsa`。需要指定PEM passphrase，passphrase过短会报错

  ```
  root in /usr/local/\/ssl λ openssl rsa -in rsa -des
  writing RSA key
  Enter PEM pass phrase:
  Verifying - Enter PEM pass phrase:
  -----BEGIN RSA PRIVATE KEY-----
  Proc-Type: 4,ENCRYPTED
  DEK-Info: DES-CBC,41A6A9CF81E97DD0
  
  K/fV8K2z0mBfFdzZUZSvO/iBRwp7roo5kNqSWX2FO7lLYuFQrmaIhJKEINhz7gbK
  GOcZulkdhAr/58/5SPgN6zIt19rTDMuvJjZ9sCZ2fSXqENTNqboM2Pl9VrcwcDrN
  ZF0Cw7zDrjK4lsJTBmvzfJGYRm9rE+YKc/rMW9t81cJ1lTQPlpuleUJv3EHSUv+E
  +mNBKi/9Hb33pk2pgdt8YyAOgmpMmul7FMkvnN8fgG6gLnqZH+LCo6VKiUNao8ks
  0bqXszDCabBxmARkV39bwYExYmDi2PkS+HTc4R0SQcoIWFz8fQ5025SusXYPu9y9
  pQ3n0clh070fd5wQn8eOEqAAuVA2tRh3Gtnp/XTap1GIjm9qibF02ZnQ5p1ubnND
  e8uTTL8hxYDmGpw9u5uw4JMZS/rCCHONIM37fVyHYCs=
  -----END RSA PRIVATE KEY-----
  ```

- `-check`

  检查密钥是否符合rsa的规则

- `-text`

  以明文格式输出公钥的和加密的密钥

- `-noout`

  不输出加密的密钥，一般与`-text`一起使用

  ```
  root in /usr/local/\/ssl λ openssl rsa -in rsa -noout -text
  RSA Private-Key: (512 bit, 2 primes)
  modulus:
      00:c0:29:18:f7:ce:2b:3c:4d:7e:67:d1:4a:b0:c6:
      12:de:3c:a1:25:3b:cd:93:4e:a1:a1:c1:6c:81:98:
      7d:6d:31:31:18:7f:31:b3:ce:7b:c8:55:0c:f7:67:
      8d:50:cf:22:a6:61:fe:f8:b4:75:4c:25:79:95:b0:
      22:02:01:7a:23
  publicExponent: 65537 (0x10001)
  privateExponent:
      63:12:8d:51:ee:14:f2:81:4d:c3:be:ef:50:56:bf:
      11:9f:96:c1:b8:a4:93:e7:3d:84:45:52:69:3a:b2:
      a8:21:88:06:9d:cd:e9:10:3b:e1:ae:4c:b5:5a:ca:
      f8:61:49:4b:1b:96:b5:cb:74:f2:97:d9:ff:08:0c:
      8f:b7:aa:81
  prime1:
      00:f2:12:7c:99:3e:07:e6:64:b1:a2:11:17:16:56:
      1b:22:c1:85:9d:2d:de:62:6b:75:16:8e:ff:03:39:
      b7:9b:63
  .....
  ```

  利用私钥生成公钥

  ```
  root in /usr/local/\/ssl λ cat rsa
    File: rsa
    -----BEGIN RSA PRIVATE KEY-----
    MIIBOwIBAAJBAMApGPfOKzxNfmfRSrDGEt48oSU7zZNOoaHBbIGYfW0xMRh/MbPO
    e8hVDPdnjVDPIqZh/vi0dUwleZWwIgIBeiMCAwEAAQJAYxKNUe4U8oFNw77vUFa/
    EZ+Wwbikk+c9hEVSaTqyqCGIBp3N6RA74a5MtVrK+GFJSxuWtct08pfZ/wgMj7eq
    gQIhAPISfJk+B+ZksaIRFxZWGyLBhZ0t3mJrdRaO/wM5t5tjAiEAyzd0iJLDFWqX
    lntehD6MURR8uS783+VAZSfv2T12wkECIA0ZBvjbrF3A8QON3SvuOMWmpu4cPz4g
    BlPUJOQtyUt1AiEAx+ukj3+iwHz+6KIyF/PY4yM+mIgrarEEqv+hLJ0VKoECIQCU
    hCig1R5aipSUWADVbxjyUCiKRmmJqzosrzIMjrWa1Q==
    -----END RSA PRIVATE KEY-----
  root in /usr/local/\/ssl λopenssl rsa -in rsa -pubout -out rsa.pub
  writing RSA key
  root in /usr/local/\/ssl λ cat rsa.pub 
    File: rsa.pub
    -----BEGIN PUBLIC KEY-----
    MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAMApGPfOKzxNfmfRSrDGEt48oSU7zZNO
    oaHBbIGYfW0xMRh/MbPOe8hVDPdnjVDPIqZh/vi0dUwleZWwIgIBeiMCAwEAAQ==
    -----END PUBLIC KEY-----
  ```

### rsautl

rsa utility，可以对数据通过rsa的方式加密和解密

sytanx：`openssl rsautl [options]`

- `-in <file>`

  指定读入的文件，如果没有会从stdin中读入

- `-out <file>`

  指定输出的文件，如果没有会输出到stdout

- `-inkey <file>`

  指定读入的key，默认是rsa 私钥

- `-keyform PEM|DER|ENGINE`

  指定key的格式

- `-pubin`

  输入的文件是rsa 公钥

- `-certin`

  输入的一个证书文件包含rsa 公钥

- `-sign`

  对输入的文件签名，然后输出签名，需要私钥

- `-encrypt`

  对输入的

### enc

通过指定的cipher对输入的内容加密和解密(==对称加密==)

sytanx：`openssl enc -cipher [options]`

```
root in /usr/local/\/ssl λopenssl enc -des -pbkdf2 -iter 3 -in /etc/resolv.conf 
enter des-cbc encryption password:
Verifying - enter des-cbc encryption password:
Salted__f�FĀ�
               y��A�s�Q�dQI�b��
                                       2Z&�rr��I!��jK�;Y.�JG�6l�"L���5 m%n�����7���r��Sw��`�/-f��������dc�x/9�#root in /usr/local/\/ssl λ 

root in /usr/local/\/ssl λ openssl enc -des -pbkdf2 -iter 3 -d -in encrypt                
enter des-cbc decryption password:
# Generated by NetworkManager
search localdomain
nameserver 114.114.114.114
nameserver 8.8.8.8
```

- `-e`

  对输入的内容加密，默认值

- `-iter <count>`

  指定加密或解密迭代的次数，迭代次数越高安全性越高

- `-pbkdf2`

  使用pbkdf2加密算法对迭代次数加密和解密

- `-d`

  对输入的内容解密

  ```
  root in /usr/local/\/ssl λ openssl enc -des -e -in rsa -out rsa.sym
  enter des-cbc encryption password:
  Verifying - enter des-cbc encryption password:
  *** WARNING : deprecated key derivation used.
  Using -iter or -pbkdf2 would be better.
  
  
  root in /usr/local/\/ssl λ openssl enc -des -d -in rsa.sym
  enter des-cbc decryption password:
  *** WARNING : deprecated key derivation used.
  Using -iter or -pbkdf2 would be better.
  -----BEGIN RSA PRIVATE KEY-----
  MIIBOwIBAAJBAMApGPfOKzxNfmfRSrDGEt48oSU7zZNOoaHBbIGYfW0xMRh/MbPO
  e8hVDPdnjVDPIqZh/vi0dUwleZWwIgIBeiMCAwEAAQJAYxKNUe4U8oFNw77vUFa/
  EZ+Wwbikk+c9hEVSaTqyqCGIBp3N6RA74a5MtVrK+GFJSxuWtct08pfZ/wgMj7eq
  gQIhAPISfJk+B+ZksaIRFxZWGyLBhZ0t3mJrdRaO/wM5t5tjAiEAyzd0iJLDFWqX
  lntehD6MURR8uS783+VAZSfv2T12wkECIA0ZBvjbrF3A8QON3SvuOMWmpu4cPz4g
  BlPUJOQtyUt1AiEAx+ukj3+iwHz+6KIyF/PY4yM+mIgrarEEqv+hLJ0VKoECIQCU
  hCig1R5aipSUWADVbxjyUCiKRmmJqzosrzIMjrWa1Q==
  ```

- `-a`

  对数据进行base64加密，通过`-d`指定对数据进行base64解密

  ```
  root in /usr/local/\/ssl λ openssl enc  -a -in /etc/resolv.conf
  IyBHZW5lcmF0ZWQgYnkgTmV0d29ya01hbmFnZXIKc2VhcmNoIGxvY2FsZG9tYWlu
  Cm5hbWVzZXJ2ZXIgMTE0LjExNC4xMTQuMTE0Cm5hbWVzZXJ2ZXIgOC44LjguOAo=
  
  root in /usr/local/\/ssl λ cat base64                                  
    File: base64
    IyBHZW5lcmF0ZWQgYnkgTmV0d29ya01hbmFnZXIKc2VhcmNoIGxvY2FsZG9tYWluCm5hbWVzZXJ2
    ZXIgMTE0LjExNC4xMTQuMTE0Cm5hbWVzZXJ2ZXIgOC44LjguOAo=
  
  root in /usr/local/\/ssl λ openssl enc -a -d -in base64
  # Generated by NetworkManager
  search localdomain
  nameserver 114.114.114.114
  nameserver 8.8.8.8
  
  ```

- `-in`

  指定输入的文件，如果没有指定会从stdin中读入

  ```
  root in /usr/local/\/ssl λ openssl enc -des -e -in rsa
  enter des-cbc encryption password:
  Verifying - enter des-cbc encryption password:
  *** WARNING : deprecated key derivation used.
  Using -iter or -pbkdf2 would be better.
  Salted__|4]��9˧��z8~��,��ݏ����杠?J�_�i�2@�#�uC}�Ch�����^	m݄69�/�ɟRC4��0Y%�ai�1��wv��-TWG_|Zw�)���v�JZ��P���5
  ��aЊQL��<Q�Ғ�(�;���N
         �<C1�"��V��,L!R\1@�.��ҁl5Ÿ3�0�?����z�1J�(	�M�o}���|�݊��u
  ```

- `-out`

  指定输出的文件，如果没有会输出到stdout

- `-salt`

  加密时使用salt，默认会生成一个随机的salt

- `-S salt`

  指定salt，必须是十六进制，可以使用`xxd`来生成hex digit

  ```
  root in /usr/local/\/ssl λ openssl enc -des -S '7436' -in /etc/resolv.conf
  enter des-cbc encryption password:
  Verifying - enter des-cbc encryption password:
  hex string is too short, padding with zero bytes to length
  *** WARNING : deprecated key derivation used.
  Using -iter or -pbkdf2 would be better.
  Salted__t6����gٻ�Ncc\��}ߵ�r,���.�A�®�v̺�]���,YJLԄ�<_�z��,�B=N'�EF�b�����Q]�f:�[�E4q�ny�^m#  
  ```

  

















