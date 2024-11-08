# postman使用指南

[TOC]

### #url中含有中文

![10](https://github.com/dhay3/image-repo/raw/master/20210601/10.19xbpaiaz074.png)

postman发送中文请求时, 必须要编码, 否则会报错

### #get请求

![1](https://github.com/dhay3/image-repo/raw/master/20210601/1.65ippixy5o00.png)

---

### #post请求

> 同理put和delete请求

![2](https://github.com/dhay3/image-repo/raw/master/20210601/2.2bw7aa40rgu8.png)

---

### #post 发送Json串

这里无需再设置Headers中的Content-Type

![3](https://github.com/dhay3/image-repo/raw/master/20210601/3.7iyueydeot40.png)

---

### #上传文件

![4](https://github.com/dhay3/image-repo/raw/master/20210601/4.385hiy4w7h40.png)

----

![5](https://github.com/dhay3/image-repo/raw/master/20210601/5.2y4808tnxu80.png)---

---

### #postman 发送post ajax

![8](https://github.com/dhay3/image-repo/raw/master/20210601/8.rc2cjkoczz4.png)

---

![9](https://github.com/dhay3/image-repo/raw/master/20210601/9.4kn21htq9la0.png)

### #postman 发送shiro权限认证

方法一:  可以通过postman登入, postman会保存cookie

方法二: 使用浏览器登入, 将登入后的cookie复制

![6](https://github.com/dhay3/image-repo/raw/master/20210601/6.4a4gnjobu6y0.png)

粘贴到postman中

![7](https://github.com/dhay3/image-repo/raw/master/20210601/7.646uzdrjnos0.png)

