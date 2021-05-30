# genrsa

syntax：`openssl genrsa [options] [numbits]`

==只能用于生成私钥==，numbits标识生成密钥的长度默认2048，不能小于512。默认会将密钥输出在stdout

## 参数

- -rand file

  将文件做为生成密钥的随机参数

- -rand file

  使用指定文件做为seed，可以有多个文件

- -primes num

  指定运算的互质数，必须是1-16之间的数字

- `-<cipher>`

  指定对密钥加密的cipher，==默认不对密钥加密==，如果提供该参数并且没有使用-passout需要提供passpharse

  ```
  root@ubuntu18:/opt/ssl# openssl genrsa -des -rand /etc/resolv.conf rsa.pk 512
  Generating RSA private key, 512 bit long modulus (2 primes)
  ...+++++++++++++++++++++++++++
  .....................+++++++++++++++++++++++++++
  e is 65537 (0x010001)
  Enter pass phrase:
  Verifying - Enter pass phrase:
  -----BEGIN RSA PRIVATE KEY-----
  Proc-Type: 4,ENCRYPTED
  DEK-Info: DES-CBC,3AFED6566E4B4A3A
  
  tYpVFGZ507v0ZcrEexwnc30Bmyu6nO2jSF3rjhoa16hh4RzKEH/ERGhxy9eEwURK
  wqeLE/ba0FhtvPGSpirqM5/kxtX+EqaZKPDhomSuk/AuflTKBnOE3/1EHA/0i2Th
  XqL4iSwlNb6+mmbG1grlpEGFp4vL3HbVW2G2Q2RSWevtB3pdO+ACpBgEmBXB7b9p
  WDXFdnwUGqqN9gtoXz6bwpTg8xgrhFfgRC8YZ4QxiWP6x+5SBfzEOJQ3OoDQrNtM
  X65yIbIMRwvQzFf3G1weLkdE2q6b0q44vLH97dFXyfHEshP9ANVbTbQjLc+tRC7s
  U0XG1w0FfK530SKAaU6NE/oU9NkeCL+XVno+VWrJw2cODqc2bAkNGICrcJo02YRS
  t6hhgvImtkUAwYo4E8qI++XSrrgWAiZP3yYR8QeOWIU=
  -----END RSA PRIVATE KEY-----
  ```

- -passout arg

  指定私钥的密码，格式参照[passpharse](../openssl passpharse)。这个参数实际上是不安全的

## 例子

prompt passwd

```
root@ubuntu18:/opt/ssl# openssl genrsa -des -primes 2 -rand /etc/resolv.conf rsa.pk 512
Generating RSA private key, 512 bit long modulus (2 primes)
..+++++++++++++++++++++++++++
........+++++++++++++++++++++++++++
e is 65537 (0x010001)
Enter pass phrase:
Verifying - Enter pass phrase:
-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: DES-CBC,8F6C0F7D8F6B7150

/ZkTHIoeiBGRNddKLV2uHWuj4ZXJTQS/qd1sTpipwZkZHB4A1JXw0yVHGfHZlR97
7jfvpd+ohvlStExDKKnMx7cJKqRYTYqypFd8sCVmtjNFBXA3JuduTvTnjKFKPPkt
JS19pYQISgSmYGi8f9VT9alT7PeQ8kZLtHtfPZKLVWhFvFHdTKiq4wzZk5n0FZB2
6wtwfP1gK6PQ4HIes1fNAEsTbrxTZFu5ljk/DO4rUNw9Nbb+Ps+pZDT3R7nIxmuR
TGHTggDmjSppRODbxJxDFMiFG4LI3muk0xZoWc5wmGCv90658V5O5jQvo3pqYNiP
g8KA9BZMtgPZFo8zNrXCgbcPmxgRN8Bq/o0zTy43tnjwt97c7HY9TsX5xRqiq/y6
tnMOPFXoacylMk+GdQ52hPBmm2zelX6v8j5BlwJ3js4ShNaBoToiIw==
-----END RSA PRIVATE KEY-----
```

fixed password

```
root@ubuntu18:/opt/ssl# openssl genrsa -des -passout pass:1234 -primes 2 -rand /etc/resolv.conf wrsa.pk 512
Generating RSA private key, 512 bit long modulus (2 primes)
........+++++++++++++++++++++++++++
...............+++++++++++++++++++++++++++
e is 65537 (0x010001)
-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: DES-CBC,4A6E76A5312BCCAA

EldQOGLOIbOZ5P+mu/l5dG0z7NMeaZLYSQjaiibjPPtP5r2g/4mKJ1ZOahEjjSLE
G0UyxdAZh9Wr3ZPbHAwKtWFzrKQoSyuYjW/nuervsGb2Se0I+9wIos7QABZyXfPf
Te+YNXT1I2PlUVkemMyyQm2unBD1BNEz0rhF4Zev2e33awyeiKj/qby0Cugz9/n2
AbJhiIlNM8q+JvzFzykO96/zm7jsKkAMD7tGSoW6bD53xEHGT16Ogx1Y1SQxkqO7
wbMeYmdKlp9YripDqmfDsEuzeuEc3dNjDk3fWmh+OODK8xe5cC19YyPTwSXV51Rc
OSCpR479L/3iFzf57LTGejDB48B7yyLXFNZDlOQa+wXftpG6WKH5HAWKgSQvPIXi
kAiHQHcdOBKxiu5s28ugy9s4vy3U5RSToN5kBHMIdimVgSbcYadZNA==
-----END RSA PRIVATE KEY-----
```

















