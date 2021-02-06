# JSON

参考:

https://www.jianshu.com/p/46afe382b5dd?from=singlemessage

[TOC]

## 定义

JSON(JavaScript Object Notation)是一种轻量级的数据交换格式

## 语法

很多人搞不清除JSON和JS对象的关系,其实可以这么理解:

JSON是JS对象的字符串表示,它使用文本表示一个JS对象的信息,本质是一个字符串

```javascript
var obj = {a: 'Hello', b: 'World'}; //这是一个对象，注意键名也是可以使用引号包裹的
var json = '{"a": "Hello", "b": "World"}'; //这是一个 JSON 字符串，本质是一个字符串
```

- JSON对象

```json
{ "firstName":"John" , "lastName":"Doe" }

/*~~~~~~~~~~~~~~~~~~*/
//前后端分离，推荐后端返回给前端数据格式
{
"status" : 0 ,          //执行状态码
"msg"    : "SUCCESS",   //说明文字信息，没有为NULL
"data"   :[{            //对象中嵌套数组，数组是返回的数据，
"id"    : 1 ,
"name"  : "xiaohong"
},{
"id"    : 2,
"name"  : "xiaoming"
}]
}
```

- JSON数组

```json
[{
"id" : 1 ,
"name" : "xiaoming"
},{
"id" : 2 , 
"name" : "xiaohong"
}]
```

- 反例

```js
{"id" : ox16 } //不合法，数值需要是十进制

{"name" : underfined } //不合法，没有该值

[{
"name" : NUll,
"school" : function() {
console.log("该写法是错误的")
}
}]//json中不能使用自定义函数，或系统内置函数
```

## JSON和JS对象互转

要实现从JSON字符串转换为JS对象，使用`JSON.parse()`方法:

```js
var obj = JSON.parse('{"a": "Hello", "b": "World"}'); //结果是 {a: 'Hello', b: 'World'}
```

要实现从JS对象转换为JSON字符串，使用 `JSON.stringify()` 方法：

```js
var json = JSON.stringify({a: 'Hello', b: 'World'}); //结果是 '{"a": "Hello", "b": "World"}'
```

