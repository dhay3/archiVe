# SSL & TSL

参考：

https://segmentfault.com/a/1190000014740303

https://cshihong.github.io/2019/05/09/SSL%E5%8D%8F%E8%AE%AE%E8%AF%A6%E8%A7%A3/

## 概述

SSL (Secure Socket Layer)安全套接层。主要私密性，信息完整性和身份认证。SSL 2.0和SSL 3.0被公开发布和使用。

TLS (Transport Layer  Securtiy)安全传输层。SSL是他的前生。TLS 1.0对应这SSL 3.0。TLS后来又有了1.1版本和1.2版本，1.3版本目前还在草案中。

### 知名协议

- HTTP over SSL

  简写https，端口使用443

  ![Snipaste_2020-08-25_17-24-19](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2020-08-25_17-24-19.4429l90bojo0.png)

- Email over SSL

  SMTP，POP3，IMAP也支持SSL

## 主要协议

1. 握手协议

   是客户端和服务段用于SSL连接通信的第一个协议。让通讯双方进行身份认证，协商加密算法，交换密钥和保护SSL记录中的信息。

2. 记录协议

   在客户端和服务端握手成功后使用记录协议。实现握手协议中定义的加密算发，数据压缩和封装。

3. 报警协议

   客户端和服务端发现错误后，向对方发送一个警报信息。如果是致命错误，则算法立即关闭SSL连接，双方还会先删除相关的会话号，秘密和密钥。

## 通信过程

Handshake phase(握手阶段)

握手阶段一共需要发送13个数据包，*代表不是必须

![Snipaste_2021-02-24_10-48-29](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-02-24_10-48-29.5p4pubr0c6s0.png)

![Snipaste_2021-02-24_10-40-39](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-02-24_10-40-39.6soepwrd6o00.png)

1. ClientHello

   客户端发包。包含SSL版本，用于认证的随机数(Random1)，数据压缩算法，加密套件(cipher 客户端支持的加密算法)，Session id等。类似于TCP SYN

   ![Snipaste_2021-02-24_10-59-54](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-02-24_10-59-54.4upp6r1lh3s0.png)

2. ServerHello

   服务端收到ClientHello必须发送ServerHello。服务器会检查指定诸如TLS版本和算法的客户端问候的条件，如果服务器接受并支持所有条件，它将发送其证书以及其他详细信息，否则，服务器将发送握手失败消息。

   ![Snipaste_2021-02-24_11-07-51](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-02-24_11-07-51.1zylqplvn2ps.png)

   如果接收服务端会选择ClientHello中的一个cipher做为plian text的加密算法，同时生成一个随机数(Random2)。Random1和Random2最后会影响生成的密钥

3. Certificate

   服务端发送验证证书（证书中包含验证申请证书时的公钥）

   ![Snipaste_2021-02-24_11-20-48](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-02-24_11-20-48.6cvabvbszqk0.png)

4. Server key Exchange

   决定服务端和客户端密钥交换的算法(比如DH)

   ![Snipaste_2021-02-24_11-23-49](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-02-24_11-23-49.6tx40jytw9w0.png)

   这里使用Diffie-Hellman算法交换密钥，算法如下

   - Alice与Bob确定两个大素数n和g，这两个数不用保密
   - Alice选择另一个大随机数x，并计算A如下：$A=g^{x} \bmod  n$
   - Alice将A发给Bob

   - Bob  选择另一个大随机数y，并计算B如下：$B=g^{y} \bmod n$

   - Bob将B发给Alice
   - 计算秘密密钥K1如下：$K1=B^{x} \bmod n$
   - 计算秘密密钥K2如下：$K2=A^{2}\bmod n$
   -  K1=K2，因此Alice和Bob可以用其进行加解密

5. Server Hello Done

   表明服务器的发送信息完毕

6. Client Key exchange

   根据从服务端收到的随机数，按照不同的密钥交换算法。按照不同的密钥交换算法，算出一个pre-master，发送给服务器，服务器端收到pre-master算出main master。而客户端当然也能自己通过pre-master算出main master。如此以来双方就算出了对称密钥。

   如果是RSA算法，会生成一个48字节的随机数，然后用server的公钥加密后再放入报文中。如果是DH算法，这是发送的就是客户端的DH参数，之后服务器和客户端根据DH算法，各自计算出相同的pre-master secret.

   ![Snipaste_2021-02-24_19-15-26](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-02-24_19-15-26.2e99ws409p1c.png)

7. Change Cipher Spec 

   表示随后的信息都将用双方商定的加密方法和密钥发送（ChangeCipherSpec是一个独立的协议，体现在数据包中就是一个字节的数据，用于告知服务端，客户端已经切换到之前协商好的加密套件（Cipher Suite）的状态，准备使用之前协商好的加密套件加密数据并传输了）。

   ![Snipaste_2021-02-24_19-19-09](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-02-24_19-19-09.6p5ai8mvkes0.png)

8. Encrypted Handshake Message

   加密握手信息

9. Change Cipher Spec

   高数客户端，服务端已经切换到之前协商好的加密套件（Cipher Suite）的状态，准备使用之前协商好的加密套件加密数据并传输了

   ![Snipaste_2021-02-24_19-22-56](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-02-24_19-22-56.1hf0nblfvbpc.png)

10. Encrypted Handshake Message

    加密握手信息

11. Application Data 

    客户端传输加密数据包		

    ![Snipaste_2021-02-24_19-24-41](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210601/Snipaste_2021-02-24_19-24-41.323z9x0wvku0.png)

12. Application Data

    服务端传输加密数据包