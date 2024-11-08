# HACKTHEBOX邀请码

进入hackthebox就需要你通过hacking获取邀请码, 如果没有头绪让我们看一下want some help, 他会提示`you could check the console`

<img src="..\..\imgs\_CTF\Snipaste_2020-09-12_10-06-55.png"/>

他这里提示让你查看以下js, 我们这里可以通过请求的地址查看js代码

<img src="..\..\imgs\_CTF\Snipaste_2020-09-12_10-10-21.png"/>

发现`inviteapi.min.js`的代码明显有语法上的问题

<img src="..\..\imgs\_CTF\Snipaste_2020-09-12_10-12-03.png"/>

这里以管道符拆分会得到一个函数, makeInviteCode, 让我们调用以下这个函数

<img src="..\..\imgs\_CTF\Snipaste_2020-09-12_10-14-07.png"/>

这里得到一串ROT13加密的字符串, 我们去解密.

<img src="..\..\imgs\_CTF\Snipaste_2020-09-12_10-15-26.png"/>

这里提示我们发送POST请求到`/api/invite/generate`,  我们这里用Postman发送请求

<img src="..\..\imgs\_CTF\Snipaste_2020-09-12_10-16-28.png"/>

明显的得到了一串base64编码的字符串, 用notepad++解码, 就能获取到邀请码了

<img src="..\..\imgs\_CTF\Snipaste_2020-09-12_09-56-59.png"/>



