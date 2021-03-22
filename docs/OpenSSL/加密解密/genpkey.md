# genpkey

通用，生成私钥。

syntax：`openssl genpkey [options]`

## options

- -algorithm alg

  指定私钥的加密算法

- -pkeyopt opt:value

  按照指定key generation options生成私钥

- -out filename

- -out form DER | PEM

- -pass arg

  私钥的密码，具体值参考[passpharse](../../openssl passpharse)

- `-<cipher>`

  指定私钥加密的cipher

## key generation options

### rsa

==可以参考rsa相关==

- rsa_keygen_bits:numbers

  默认1024

- rsa_keygen_primes:numprimes

  默认2

### dsa

- dsa_paramgen_bits:numbits
- dsa_paramgen_q_bits:numbits
- dsa_paramgen_md:digest

## 例子

生成rsa私钥

```
root@ubuntu18:/opt/ssl# openssl genpkey -des -algorithm RSA -pkeyopt rsa_keygen_bits:512  -pass pass:1234
.+++++++++++++++++++++++++++
..+++++++++++++++++++++++++++
-----BEGIN ENCRYPTED PRIVATE KEY-----
MIIBsTBLBgkqhkiG9w0BBQ0wPjApBgkqhkiG9w0BBQwwHAQINXHf3hgETBwCAggA
MAwGCCqGSIb3DQIJBQAwEQYFKw4DAgcECFBIfg6b8jGOBIIBYH/uhFYzYq+e6ykc
ZKKWp7UdT4fwQ/Vx9AVu5pa91ocN29njKpMZP4Ivs+6i9VXrawwmUUKCwOkmAkBr
iIcpMGgfR1rq/vPNygAMU/aSn2oL7Y9+qcp5JMZFYqbWNcHYuI1F1sUnHvNpNx4x
4k5J6LaAardVLfoMQ5pCD7E/vJ5onmmBlBv2s7pohEMd5bM3YzTZy9cFaJMGMDCm
afkG6OM8QhN0pmnjliN4Y6N1rthT2LeXpCKWtElR8ILT5PYBjHlII+5yUDPHgZT5
51zgtTf3e0heVd/10h/mqbCq6OwHD91gxjbHWanjVmJUg2c2C066MnCBtbii05i1
DoJmbFqzE9ChgJh+fU+iZbjvdpQk/Z+HKCo9LO5GLkx7vVxp9ZCX+/QE18dMibCD
0s6ae4RSfTNwW2KHf8pdJoxLJjB69reQUms53bgwD2B4LGNAdre4zZ6SbCmWjzIM
7l18jgM=
-----END ENCRYPTED PRIVATE KEY-----

```

生成dsa私钥

```
root@ubuntu18:/opt/ssl# openssl genpkey -des -genparam -algorithm DSA -pkeyopt dsa_paramgen_bits:512  -pass pass:1234
..+........+..+....................+...+..........+.............+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*
.+.......+...........................+.................+............................+..+......+......+.............+...+..+.........+..........................................+........................+.......................+..........+.+.................+..+...+.+..............+.....+........................+..........+.........+..................+...................+......+....+.......+.+.+..................+...............+.....+..............+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*
-----BEGIN DSA PARAMETERS-----
MIGcAkEAug4dwoA1CDLki8gD3SB0kmZanBYRY0s2Wzhk+1PWw8Gj6xfpfrjs6Vcb
efyds578n4b5QAI8eYI2XmFI1zlXiQIVAN3JsODFZvhsTiC5Gij6fPjV3W4lAkBd
KnX03dHs7HqI+Ctk9HDu49DhmddgkAREL+sH6/qMIc5MUCSLK9nKdSPB5MHdMaRH
lt5VGt7LiPdYzTR1EIFW
-----END DSA PARAMETERS-----
```













