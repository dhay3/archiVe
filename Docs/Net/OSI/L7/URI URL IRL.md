# URI URL  IRL

## URI

A URI(Uniform Resource Identifier) is an unique sequence identifer that identifies an resource of but it might won’t tell you how to access it or where it’s located

简单的说 URI 用于标识资源，==但是用户**不一定**能通过 URI 标识来直接获取资源==。例如 书库编号 `ISBN 0-486-27557-4` 可以标识编号对应的书籍，但是我们不能通过这个编号直接获取资源

在网络中下列均为 URI

```
#定位 index.html
https://developer.mozilla.org
https://developer.mozilla.org/en-US/search?q=URL
ftp://ds.internic.net/internet-drafts/draft-ietf-uri-irl-fun-req-02.txt
#Nginx 中的 location directive 中就是使用 URI 来匹配规则的
/image/
```

## URL

A URL(Uniform Resource Locator) is a specific type of identifier that not only identifies the resource but tell you how to acess it or where it’s located

简单的说 URL 用于定位资源，==用户可以通过 URL 定位来获取资源==

在网络中下列均为 URL

```
#获取 index.html
http://developer.mozilla.org
https://developer.mozilla.org/en-US/search?q=URL
ftp://ds.internic.net/internet-drafts/draft-ietf-uri-irl-fun-req-02.txt
jdbc:mysql://10.0.3.34:3306/mysql?characterEncoding=utf8&zeroDateTimeBehavior=convertToNull&useSSL=false&useJDBCCompliantTimezoneShift=true&useLegacyDatetimeCode=false&serverTimezone=GMT%2B8&allowMultiQueries=true&allowPublicKeyRetrieval=true
```

通常 URL 会以如下格式表示，必须包含 scheme (个人持相反观点，并非必须包含)

```
URL = scheme ":" ["//" authority] path ["?" query] ["#" fragment]
```

## IRI

IRI(Internationalized Resource Identifier) is a type of URL that can contain non-ASCII characters

早期 URL 中只支持含有 ASCII 码中的字符，而 IRI 是 URL 中的一种，允许 URL 中包含非 ASCII 码中的字符(例如 中文，日文, 特殊字符 等)

在网络下列均为 IRI

```
https://www.google.com/search?q=波士顿动力
https://c plus plus
```

但是有些程序或者是应用不支持 IRI，通常为了兼容这些程序或者是应用会将 IRI 以 UTF-8 格式编码(也被称为 URLencode)。例如上述 IRI 就会被编码成下列格式

```
https://www.google.com/search?q=%E6%B3%A2%E5%A3%AB%E9%A1%BF%E5%8A%A8%E5%8A%9B
```

## Differences

从上面的定义中很容易得出如下关系，也是最重要的一点不同

$IRI \in URL \in URI$

URI 可以表示 URL 也可以表示 IRI，是统称。IRI 是 URL 中的一种格式，可以包含非 ASCII 码中的字符

举几个例子

```
#URI
ISBN 0-486-27557-4
#URI or URL
https://developer.mozilla.org
#IRI
https://c plus plus
```

**references**

[^1]:https://auth0.com/blog/url-uri-urn-differences/
[^2]:https://fusion.cs.uni-jena.de/fusion/2016/11/18/iri-uri-url-urn-and-their-differences/
[^3]:https://developer.mozilla.org/en-US/docs/Learn/Common_questions/Web_mechanics/What_is_a_URL
[^4]:https://en.wikipedia.org/wiki/Uniform_Resource_Identifier
