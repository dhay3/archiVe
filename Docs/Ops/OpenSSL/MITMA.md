# MITMA

ref

https://en.wikipedia.org/wiki/Man-in-the-middle_attack

## Digest

二战期间盟军间相互通信需要通过电报，假设法国给德国发电报需要在 1 月 1 号一起进攻苏俄。但是电报早就已经被苏俄的情报网渗透了，德法两国都不知道还以为电报会直接发送给对端。苏俄在法国发送电报后就将其电报截获并将其解密后修改信息成在 2 月 1 号一起进攻苏俄加密并发给德国。德国收到电报后解密，就误以为真的是在 2 月 1 号进攻苏俄。结果可想而知。

这也被称为 eavesdropping，在现代人的生活在屡见不鲜，例如 GFW ，Prism

而这些也是 MITMA ( man in the middle attack ) 的生活中的实例

## What  is MITMA

MITMA 逻辑上的流程具体如下：

假设现在 A 和 B 需要通信

![Snipaste_2020-08-25_17-57-00](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20221214/Snipaste_2020-08-25_17-57-00.1rrb6vapcvsw.webp)

- 首先 A 会发送一串信息给 B 表明自己的身份

- C 拦截到 A 的消息，将消息转发给 B

- B 收到消息后，选择一个不重数 RB，发送给 A，来确认A的身份。但被==C 截获 RB==，C 将 RB 转发给 A

- A 用自己的私钥 SKA 加密 RB 发送给 B。C 拦截到消息并将信息丢弃，用自己的私钥 SKC 冒充是 A 的私钥，对 RB 加密，并发送给 B。

- B 收到消息请求 A 的公钥。但是被 C 拦截，C 将消息转发给 A

- A 把自己的公钥 PKA发送给B，但是被 C 拦截（==获取到 A 的公钥==），C 用自己的公钥 PKC发送给B。

- ==B 用 C 的公钥解密，这样 B 就误以为 C 是A==

- 然后 B 用 C 的公钥发送消息给 A，C 用 A 的公钥转发给 A