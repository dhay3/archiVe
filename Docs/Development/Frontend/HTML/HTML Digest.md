# HTML Digest

ref:

https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/HTML_basics

https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Getting_started

HTML( HyperText Markup Language ) is one of  markup languages to mark up text for web page

## Anatomy of an HTML element

> 带有 closing tag 的 element 也被称为 block element

![](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/HTML_basics/grumpy-cat-small.png)

HTML element 由如下几部分组成

1. the opening tag

   wrapped in opening and closing angle brackets

2. the closing tag

   same as the opening tag, except that is includes a forward slash before the element name

3. the content

4. the eleme nt

   the opening tag, the closing tag

element 也可以包含  attributes

![](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/HTML_basics/grumpy-cat-attribute-small.png)

attribute 需要满足如下条件

1. a space between it and the elemetn name or the previous attribute, if the element already has one or more attributes
2. the attribute name followed by an equal sign
3. the attribute value wrapped by opening and closing quatation marks

### Nesting elements

elements 可以被内嵌（nesting），例如

```html
<p>My cat is <strong>very</strong> grumpy.</p>
```

### Empty elements

一些 elements 可以没有 content， 这些 elements 被称为 empty elements，例如`<img>`

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

  preamble，实际没有效果（历史上和 XML 的头文件功能类似）

- `<html></html>`

  root element, wraps all the content on the entire page

  包含一个 `lang` attribute 表明主选 language of the document

- `<head></head>`

  一组集合用于描述 html doc 的属性

- `<meta charset="utf-8">`

  标明文档使用字符集

- `<meta charset="utf-8">`

  the title of  the page

- `<body></body>`

  contains all the content rhat you want to show

## boolean attributes

sometmes you will see attributes written wihout values. This is entirely acceptable. These are called boolean attributes

例如

disable attribute

```
<input type="text" disabled="disabled" />
```

as shorthand 也可以写成

```
<!-- using the disabled attribute prevents the end user from entering text into the input box -->
<input type="text" disabled />

<!-- text input is allowed, as it doesn't contain the disabled attribute -->
<input type="text" />
```

## Quotes

1. attributes 可以使用 quotes 标明也可以不使用，但是在不使用 quotes 的情况下可能会造成异常
2. single or double quotes 没有实际的影响

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