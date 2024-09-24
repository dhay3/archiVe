HTTP报文

[TOC]

可以参考相关链接

https://blog.csdn.net/lyn_00/article/details/84953763

https://www.cnblogs.com/kikochz/p/12751405.html

https://www.cnblogs.com/ldq2016/p/9055933.html

报文 = 报文行 + 报文头 + 报文体

## 请求报文

<img src="..\..\..\..\imgs\_Net\20160921092902554.png"/>

> 请求行

<img src="..\..\..\..\..imgs\_Net\20160921092902556.jpg" style="zoom:80%;" />

- 请求方式

  GET：==没有请求体==，参数通过键值对传递，形如

  `www.baidu.com/query?user=username&pass=password`

  如果数据是英文字母/数字，原样发送，如果是空格，转换为+，如果是中文/其他字符，则直接把字符串用BASE64加密

  OPTIONS：预检请求，不会携带Cookie和参数，非简单请求之前会调用，如Content-Type为application/json，或是PUT，DELETE请求

  参考链接: https://www.jianshu.com/p/5cf82f092201?tdsourcetag=s_pctim_aiomsg

  CONNECT：CONNECT的作用是将服务器作为代理，让服务器代替用户访问其他网页，之后将数据返回给用户。如果使用抓包工具或是vps就会先通过代理转发请求

  参考连接：

  https://www.jianshu.com/p/54357cdd4736
  
  [https://blog.csdn.net/a464057216/article/details/52732501#%E8%BD%AC%E5%8F%91%E4%BB%A3%E7%90%86%E6%9C%8D%E5%8A%A1%E5%99%A8](https://blog.csdn.net/a464057216/article/details/52732501#转发代理服务器)
  
  HEAD：与GET类似，==但是返回的响应中没有响应体==，一般用于判断某个资源是否存在
  
   
  
  ```java
    @ResponseBody
      @RequestMapping(value = "/test", method = RequestMethod.HEAD)
    public String test() {
          return "hello world";
    }
  ```
  
  <img src="..\..\..\..\imgs\_Net\Snipaste_2020-08-21_20-46-18.png" style="zoom:80%;" />
  
  POST: 一般请求行种没有参数，有请求体，具体数据存在请求体中,==注意这里用POSTMAN或是fiddler==模拟的请求不能模拟出浏览器的效果
  
- PUT：更新请求
  - DELETE：删除请求
  
  

> 请求头

- User-Agent：表示首部包含了一个特征字符串，用来让网络协议的对端来识别发起请求的用户代理软件的应用类型、操作系统、软件开发商以及版本号

- Connection：如keep-alive ， 表示tcp连接不关闭，不会永久保持连接，服务器可设置，缺省值

- Accept：表示请求接受响应报文必须是该MIME类型的

  ```
  accept: Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
  ```

  这里的 `*/*` 表示接收所有MIME 类型 , 所以前面表示接收的顺序
  `application/json` 也是MIME类型

- Accept-Charset：告诉服务器能发送哪些字符集

- Accept-Encoding：告诉服务器能发送哪些编码方式

- Authorization：授权信息

- Referer：请求来自哪里，以URL的方式显示

- Host：接受请求的主机名和端口，80端口不用显示

- Cookie：客户端用它像服务器发送数据

- Content-Type：请求体内容的类型，==get没有请求体所以没有该属性==

  具体参考:https://www.cnblogs.com/kikochz/p/13546104.html

- Content-Length：请求体的长度

- DNT（do not track）：表明用户对于网站的追踪偏好

  0代表愿意被追踪

  1代表不愿意被追踪

> 空行

告诉服务端请求头部到此为止

> 请求体

- GET没有请求体

- POST有请求体

<img src="..\..\..\..\imgs\_Net\Snipaste_2020-08-22_15-26-12.png"/>

## 应答报文

<img src="..\..\..\..\imgs\_Net\20160921092902557.jpg" style="zoom:80%;" />

> 响应行
>
> 常见状态码参考：https://www.cnblogs.com/kikochz/p/12840652.html

- 状态码

> 响应头
>
> 参考：https://www.cnblogs.com/ldq2016/p/9055933.html

- Connection：如keep-alive ， 表示tcp连接不关闭，缺省值。使用同一个TCP连接来接收多个请求和响应。
- Content-Length：响应体的长度
- Content-Type：响应体内容的类型
- Content-Disposition：告诉浏览器以下载方式打开数据
- Last-Modified：在浏览器第一次请求某一个URL时，服务器端的返回状态200，内容是你请求的资源，同时有一个Last-Modified的属性标记（响应头）此文件在服务器端最后被修改的时间，格式类似这样：Last-Modified：Tue，24 Feb 2009
   客户端第二次请求此URL时，根据HTTP协议的规定，浏览器会向服务器传送If-Modified-Since报头（请求头），询问该事件之后文件是否有被修改过，如果服务器端的资源没有变化，则自动返回HTTP：304状态码，内容为空，这样就节省了传输数据量。当服务器端代码发生改变或者重启服务器时，则重新发出资源，返回和第一次请求时类似。从而保证不向客户端重复发出资源，也保证当服务器有变化时，客户端能够得到最新的资源。

- Set-Cookie：服务端向客户端发送cookie，前端通过`document.cookie`获取

  例如：

  响应头中有Set-Cookie: username=JasonChi，那么浏览器会在==当前页面所在域名（cookie不能跨域）==设置cookie字符串。

  当浏览器再次发送请求时，浏览器默认会自动将cookie中的字符串放在请求头中的Cookie项中发送给Web服务器。

- Expires：缓存的有效时间，如果和Cache-Control中的max-age同时存在，Expires优先级高

- Cache-Control：告诉浏览器如何控制响应内容的缓存

  具体参考https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Cache-Control

  - private：内容只缓存到私有缓存中(仅客户端可以缓存，代理服务器不可缓存)
  - public：所有内容都将被缓存(客户端和代理服务器都可缓存)
  - max-age= xxx ：设置缓存存储的最大周期，超过这个时间缓存被认为过期(单位秒)。与`Expires`相反，时间是相对于请求的时间。
  - no-cache： 在发布缓存副本之前，强制要求缓存把请求提交给原始服务器进行验证(协商缓存验证)。
  - no-store：所有内容都不会被缓存到缓存或 Internet 临时文件中

- Location：标明请求重定向后的URL，仅在与3xx（重定向）或201（已创建）状态响应一起使用时才提供含义。

> 响应体

<img src="..\..\..\..\imgs\_Net\Snipaste_2020-08-22_16-07-52.png" style="zoom:80%;" />
