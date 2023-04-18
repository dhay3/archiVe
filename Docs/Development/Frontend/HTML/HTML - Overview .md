# HTML - Overview

HTML( HyperText Markup Language ) is one of  markup languages to mark up text for web page

HTML 是 Markup 一种语言

## HTML element syntax

> 带有 closing tag 的 element 也被称为 block element

![](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/HTML_basics/grumpy-cat-small.png)

HTML element 由如下几部分组成

1. the opening tag

   开始标签

2. the closing tag

   结尾标签

3. the content

   标签中的内容

element 也可以包含额外的信息用于描述 element, 被称为 attributes

![](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/HTML_basics/grumpy-cat-attribute-small.png)

attribute 需要满足如下条件

1. a space between it and the elemetn name or the previous attribute, if the element already has one or more attributes

   以空格分隔

2. the attribute name followed by an equal sign

3. the attribute value wrapped by opening and closing quatation marks

### Nesting elements

elements 可以被内嵌（nesting），例如

```html
<p>My cat is <strong>very</strong> grumpy.</p>
```

### Void elements

一些 elements 可以没有 content， 也不需要 closing tag，这些 elements 被称为 void elements，例如`<img>`

```html
<img src="images/firefox-icon.png" alt="My test image" />
```

## Anatomy of an HTML Doc

```html
<!DOCTYPE html>
<html lang="en-US">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width" />
    <title>My test page</title>
  </head>
  <body>
    <img src="images/firefox-icon.png" alt="My test image" />
  </body>
</html>
```

- `<!DOCTYPE html>`

  preamble( 前言 )，实际没有效果（历史上和 XML 的头文件功能类似）

- `<html></html>`

  root element, wraps all the content on the entire page

  包含一个 `lang` attribute 表明主选 language of the document

- `<head></head>`

  一组集合用于描述 html doc 的属性，会被 SEO 检索

- `<meta name="viewport" content="width=device-width">`

  用于

- `<meta charset="utf-8">`

  标明文档使用字符集

- `<title></title>`

  the title of  the page

- `<body></body>`

  contains all the content rhat you want to show

## Special characters

In HTML, the characters `<`, `>`,`"`,`'` and `&` are special characters

转义表参考

| Literal character | Character reference equivalent |
| ----------------- | ------------------------------ |
| <                 | `<`                            |
| >                 | `>`                            |
| "                 | `"`                            |
| '                 | `'`                            |
| &                 | `&`                            |



**references**

1. https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/HTML_basics
2. https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Getting_started