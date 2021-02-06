# postman使用指南

[TOC]

### #url中含有中文

<img src="..\..\imgs\_Postman\10.png" style="zoom:80%;" />

postman发送中文请求时, 必须要编码, 否则会报错

### #get请求

<img src="..\..\imgs\_Postman\1.png" alt="1" style="zoom:80%;" />

---

### #post请求

> 同理put和delete请求

<img src="..\..\imgs\_Postman\2.png" alt="2" style="zoom:80%;" />

---

### #post 发送Json串

这里无需再设置Headers中的Content-Type

<img src="..\..\imgs\_Postman\3.png" style="zoom:80%;" />

---

### #上传文件

<img src="..\..\imgs\_Postman\4.png" alt="4" style="zoom:80%;" />

----

<img src="..\..\imgs\_Postman\5.png" alt="5" style="zoom:80%;" />---

---

### #postman 发送post ajax

<img src="..\..\imgs\_Postman\8.png" style="zoom:80%;" />

---

<img src="..\..\imgs\_Postman\9.png" style="zoom:80%;" />

### #postman 发送shiro权限认证

方法一:  可以通过postman登入, postman会保存cookie

方法二: 使用浏览器登入, 将登入后的cookie复制

<img src="..\..\imgs\_Postman\6.png"/>

粘贴到postman中

<img src="..\..\imgs\_Postman\7.png" style="zoom:80%;" />
