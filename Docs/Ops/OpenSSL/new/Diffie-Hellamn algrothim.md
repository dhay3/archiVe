# Diffie-Hellamn algrothim

ref

https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange

https://en.wikipedia.org/wiki/Elliptic-curve_Diffie%E2%80%93Hellman

https://wiki.openssl.org/index.php/Diffie_Hellman

https://en.wikipedia.org/wiki/Elliptic-curve_Diffie%E2%80%93Hellman

https://www.cryptologie.net/article/340/tls-pre-master-secrets-and-master-secrets/

https://crypto.stackexchange.com/questions/27131/differences-between-the-terms-pre-master-secret-master-secret-private-key

https://blog.csdn.net/superprintf/article/details/123072390

https://blog.csdn.net/qq_31442743/article/details/116199453

https://www.cryptologie.net/article/340/tls-pre-master-secrets-and-master-secrets/

## Digest

Diffie-Hellman key exchange is a mathematical method of securely  exchanging cryptographic keys over a public channel 

Diffie-Hellman key exchange 通常也被称为 DHE，其算法主要用于安全交换密钥。最广泛的实现是 SSL/TSL

## Cryptographic explanation

以 Alice 与 Bob 之前需要交换密钥为例子

1. 首先 Alice 与 Bob 会公开确认使用两个素数 p 和 g

   p 和 g 两个素数会通过 exchange 宣告给对端

2. Alice 选择一个 secret integer a（这里直接直接理解成 Alice 的私钥，只有 Alice 自己知道）, 并计算 A 如下：A = $g^a \bmod p$ 

3. Alice 将 A （Alice 的公钥） 发送给 Bob

4. Bob 选择一个 secret integer b （Bob 的私钥，只有 Bob 自己知道）, 并计算 B 如下：B = $g^b \bmod p$

5. Bob 将 B （Bob 的公钥）发送给 Alice

6. Alice 计算密钥 K = $B^a \bmod p$ = $g^{ab} \bmod p$

7. Bob 计算密钥 K = $A^b \bmod p$ = $g^{ab} \bmod p$

8. 因此 Alice 和 Bob 可以用其进行加对称解密（==但是K1 K2 是通过非对称加密计算出来的==），计算所得的 K 不会通过网络传输所以是相对安全的


如果在 Alice 与 Bob 之前出现一个监听者 Eve，Eve 是不能知道 K 的值的，具体可以参考下面这张图

![2022-12-21_21-30](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221221/2022-12-21_21-30.1aln9fdvrgu8.webp)

## Vibrants

在 SSL/TLS 中有 几种 DH 算法的变体

- Anoymous Diffie-Hellman

  exchange key 没有被两端校验，存在被 MITMA，安全性最弱。目前已被摒弃

- Fixed Diffie-Hellman

  DH public key 需要的参数内嵌在 certificate 中，参数不在保持不变

- Ephemeral Diffie-Hellman

  使用 ephemeral publick key，目前 SSL/TLS 中大多采用这种方案

- Elliptic-curve Diffie-Hellman

  使用椭圆曲线计算 shared secret。目前使用的最广

### ECDHE

Elliptic-curve Diffie-Hellman exchange (ECDHE)  使用 elliptic-curve public-private key 生成 shared secret 来加密通信。和 DHE 在逻辑上差不多，但是算法优于 DHE 所以 TLS 都采用 ECDHE 替代了传统的 DHE

假设现在 Alice 想要和 Bob 通信

1. 首先 Alice 与 Bob 会公开确认使用椭圆曲线和曲线的基点 G

   椭圆权限 和 G 会通过 exchange 宣告给对端

2. Alice 选择一个 secret integer dA（这里直接直接理解成 Alice 的私钥，只有 Alice 自己知道）, 并计算 A 如下：QA = $dA.G$ 
3. Alice 将 A （Alice 的公钥） 发送给 Bob
4. Bob 选择一个 secret integer dB （Bob 的私钥，只有 Bob 自己知道）, 并计算 B 如下：QB = $dB.G$
5. Bob 将 B （Bob 的公钥）发送给 Alice
6. Alice 计算 point 密钥 (xk,yk) = $dA.QB$ = $dA.dB.G$
7. Bob 计算 point 密钥 (xk,yk) = $dB.QA$ = $dB.dA.G$
8. point 相同，因此 Alice 和 Bob 可以用其进行加对称解密。计算所得的 (xk,yk) 不会通过网络传输所以是相对安全的
