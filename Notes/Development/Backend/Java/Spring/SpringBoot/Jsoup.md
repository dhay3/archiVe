# Jsoup

##### #字符串->dom

```java
String html = "<html><head><title>First parse</title></head>"
  + "<body><p>Parsed HTML into a doc.</p></body></html>";
Document doc = Jsoup.parse(html);
```

返回的Document,其中包含(至少) 一个head和一个body元素。

##### #指定uri获取dom

```java
     //发送请求到url获取dom元素,可以是get,post
        Document document = Jsoup.connect("http://www.baidu.com")
                .cookie("auth", "token")
                .timeout(3000)
                .get();
        //获取dom元素中的指定块
        String title = document.title();
        System.out.println(title);
```

##### #本地文件获取dom

  如果设置了baseUri为`http://www.qq.com/`
  则图片的实际路径为`http://www.qq.com/images/logo.jpg`

```
Document doc = Jsoup.parse(new File("path"), "utf-8", "http://www.baidu.com");
```

##### #dom元素的具体操作

传送门 https://www.open-open.com/jsoup/dom-navigation.htm

**注意一点**:

```
String relHref = link.attr("href"); //   "/"
String absHref = link.attr("abs:href"); // "http://www.open-open.com/"
```

##### #jsoup防止xss攻击

- WhiteList

| 白名单对象      | 白名单标签                                                   | 说明                                              |
| :-------------- | :----------------------------------------------------------- | :------------------------------------------------ |
| none            | 无                                                           | 只保留标签内文本内容                              |
| simpleText      | b,em,i,strong,u                                              | 简单的文本标签                                    |
| basic           | a,b,blockquote,br,cite,code,dd, dl,dt,em,i,li,ol,p,pre,q,small,span, strike,strong,sub,sup,u,ul | 基本使用的标签                                    |
| basicWithImages | basic 的基础上添加了 img 标签 及 img 标签的 src,align,alt,height,width,title 属性 | 基本使用的加上 img 标签                           |
| relaxed         | a,b,blockquote,br,caption,cite, code,col,colgroup,dd,div,dl,dt, em,h1,h2,h3,h4,h5,h6,i,img,li, ol,p,pre,q,small,span,strike,strong, sub,sup,table,tbody,td,tfoot,th,thead,tr,u,ul | 在 basicWithImages 的基础上又增加了一部分部分标签 |

使用:

```java
private static final Whitelist  whitelist=Whitelist.basicWithImages();
```

同时还可以指定添加或删除白名单标签

```java
Whitelist.basic().removeTags("a").addTags("img");
```

允许标签添加属性

这里的 :all 表示所有标签

```java
Whitelist.basic().addAttributes(":all","style").removeAttributes(":all","style");
```

