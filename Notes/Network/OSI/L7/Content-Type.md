# Content-Type

[TOC]

==GET请求的方式没有Content-Type==, 只有Post的请求有Content-Type ,通过表单的enctype属性来修改

其中常见的 Content-Type有:

##### ① `application/x-www-form-urlencoded`

- 原生的form表单, 默认采用该Content-Type, 所有字符都会进行编码

```
POST http://192.168.2.12/index HTTP/1.1 
Content-Type: application/x-www-form-urlencoded;charset=utf-8 
title=test&sub%5B%5D=1&sub%5B%5D=2&sub%5B%5D=3 
```

提交的数据按照 key1=val1&key2=val2 的方式进行编码，key 和 val 都进行了 URL 转码

- `Spring`中用@RequestParam来处理`Content-Type`为`application/x-www-form-urlencoded`数据。

修饰的对象可以是基本数据类型和自定义对象。

##### ②`multipart/form-data`

使用表单上传文件时，必须将 的 `enctype`设为 `multipart/form-data`

```
POST http://192.168.2.12/index HTTP/1.1 
Content-Type:multipart/form-data;

boundary=--WebKitFormBoundaryrGKCBY7qhFd3TrwA 
---WebKitFormBoundaryrGKCBY7qhFd3TrwA 
Content-Disposition: form-data; name="text" 
title 
---WebKitFormBoundaryrGKCBY7qhFd3TrwA 
Content-Disposition:form-data;name="file"; filename="chrome.png" 
Content-Type: image/png 
PNG ... content of chrome.png ... 
---WebKitFormBoundaryrGKCBY7qhFd3TrwA
```

- `spring`使用@RequestParam处理接收到的文件

  ```
  @RequestMapping("uploadFile")
      public JsonResult uploadFile(@RequestParam("file") MultipartFile file, @RequestParam String bucket){
          String fileUrl = aliossService.uploadFile(file, bucket);
          Map<String,String> result = new HashMap<>();
          result.put("fileUrl",fileUrl);
          return success(result);
        }
  ```

##### ③ `application/json`

- `application/json`作为请求头，用来告诉服务端**消息主体是序列化的JSON字符串**

  例如ajax将json串传到后端

```
JSvar data = "{'title':'test', 'sub' : [1,2,3]}";
$http.post(url, data).success(function(result) {
...
}); 
最终发送的请求是：
POST http://www.example.com HTTP/1.1 
Content-Type: application/json;charset=utf-8
{"title":"test","sub":[1,2,3]}
```

- `spring`中通过@RequstBody来接收

##### ④ `text/xml`

典型的 XML-RPC 请求是这样的：

```
POST http://www.example.com HTTP/1.1 
Content-Type: text/xml

<?xml version="1.0"?>
<methodCall>
    <methodName>examples.getStateName</methodName>
    <params>
        <param>
            <value><i4>41</i4></value>
        </param>
    </params>
</methodCall>
```

参考转载自

https://blog.csdn.net/baichoufei90/article/details/84030479

https://blog.csdn.net/xuanwugang/article/details/79661672
