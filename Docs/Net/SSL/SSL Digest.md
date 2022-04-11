# SSL Digest

ref:

https://en.wikipedia.org/wiki/Certificate_authority

https://www.ssl.com/faqs/what-is-a-certificate-authority/

https://blog.csdn.net/liuchunming033/article/details/48470575

https://juejin.cn/post/6844903977025273863

https://help.aliyun.com/document_detail/109827.html?spm=a2c4g.11186623.6.625.40b24ff1WNaTFy

https://www.barretlee.com/blog/2016/04/24/detail-about-ca-and-certs/

https://www.jianshu.com/p/29e0ba31fb8d

## why need SSL

在没有使用证书的情况下(无法验证服务端的有效和真实性)，如果请求一个站点，在中间路由的过程中，可能会被mitma。为了避免这种情况需要一种认证机制，也是就是SSL

## Terms

1. CSR(certificate signing reqeust)

   证书请求，是一个加密的信息(通过申请人的私钥加密)，其中包含申请人的公钥和申请人提交的信息（common name，organization，organization unit，city，state，country）

2. CA(certificate authority)

   是一个颁发数字证书(digital certificate)的权威机构。是PKI(publick key infrastructure)中的核心，负责签发证书、认证证书。

   使用CA的私钥对证书签名然后颁发给申请人。任何个体组织都可以扮演CA，但是不能得到浏览器的默认信任(==需要将其加入到系统的受信任的根证书颁发机构中==)。知名的CA有Symantec，Comodo，Godaddy，GolbalSign，Digicert，Let’s Encrypt（免费的）

2. digital certificate 
3. CA certificate
4. 