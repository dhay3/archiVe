# URI&URL&IRL

ref

https://auth0.com/blog/url-uri-urn-differences/

https://fusion.cs.uni-jena.de/fusion/2016/11/18/iri-uri-url-urn-and-their-differences/

## URL

uniform resource locator

you can think of a URL like your home address, it contians all the information to find your home. Similarly your can define a URL as a string that denotes the location of a given resource on the Internet

## URI

uniform resource identifier

it is a string that identifies a resource. From a syntical point of view, a URI string mostly follows the same format as the URLs. However, while URLs allow you to locate a resource, a URI simply identifies a resource

常见的例子就是XML头文件

## IRI

internationalized resource identifier

the superset of URIs to allow internationalization like non-Latin symbols and more

## URL vs URI

说白的就是家庭住址和身份证的区别，定位一个资源是否是必要的逻辑条件。通过URL一定能找到，但是URI不一定。同时URI包括URL

## URI vs IRI

区别就是特殊字符是否有编译，如果编译了就是IRI没有就是URI，例如`https://c plus plus`就是URI，而`https://c%20plus%20plus`就是IRI

