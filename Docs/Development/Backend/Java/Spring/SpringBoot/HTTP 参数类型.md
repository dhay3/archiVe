# HTTP 参数类型

### #请求头参数 (head)

请求头参数顾名思义，是存放在请求头中发送给服务器的参数，服务器通过解析请求头获取参数内容。通常会存放本次请求的基本设置，以帮助服务器理解并解析本次请求的body体。

参数形式如：

```apl
Accept-Language : zh-CN,zh;q=0.8,zh-TW;q=0.5,en-US;q=0.3,en;q=0.2
```

其中 `Accept-Language` 为参数名称 `:`号后面紧跟请求的值。

> ps:如上则表示本次请求的语言为中文和英文。 q用来标识语种权重 值在 0-1之间
> 服务器根据会这个请求头选择返回的语言

### #路径参数 (path)

路径参数顾名思义，本次请求的资源路径,紧跟请求域名的后面，服务器通过解析路径参数获取资源位置。路径参数是用`/`分隔的一系列字符串，其意义在于对服务器资源进行统一定位，如：

```apl
http://www.madmk.com/office/zhangsan
```

其中 /office/zhangsan 便是路径参数，其含义可解释为 找办公室里的张三，大多数情况下路径参数会代表服务器真实的文件路径。

> REST API的兴起赋予了路径参数更为广阔的含义，有兴趣的朋友可以找一下有关 REST API 设计的文章。
>

> ps:其中参数如有中文 或特殊符号如/ ,:,?,#,+,=等需要进行转义处理

### #查询参数 (query)

```api
userId=1
```

其中 `userId` 表示参数名称 `1`表示参数的值。参数名称为可重复的。
请求地址与参数之间用`?`进行分隔 多个参数之间用 `&`进行分隔，完整请求如下：

```apl
http://www.madmk.com/a/b/c?userId=1&userId=1&age=18&sex=男
```

### #请求体参数 (body)

请求体参数顾名思义，是存放在请求体中发送给服务器的参数。请求体参数格式复杂多变，服务器会先根据请求头中的 `Content-Type` 获取其格式，然后再根据其格式进行解析，常见的格式如下：

| Content-Type                      | 内容格式                                             | 示例                            |
| --------------------------------- | ---------------------------------------------------- | ------------------------------- |
| application/x-www-form-urlencoded | 表单传值，也是默认的解析形式,服务器会对表单进行解析  | userId=1&userId=1&age=18&sex=男 |
| text/plain                        | 文本值，服务器会将本次请求的请求体当作普通字符串看待 | Hello world                     |
| application/json                  | json,服务器会将请求体进行json解析，获取参数          | {“userId”:1,“sex”:“男”}         |
| application/xml                   | xml,服务器会将请求体进行xml解析，获取参数            | 参见 xml 标准格式               |
| text/html                         | html,服务器会将请求体进行html解析，获取参数          | 参见 html 标准格式              |

