# CSS - Overview

Cascading Style Sheets 简称 CSS，中文也被称为层叠样式。可以通过 CSS 对 HTML 元素选择性的添加样式，诸如

- 字体颜色，大小
- 背景图片，背景颜色
- etc

## CSS ruleset syntax

CSS 是 rule-based 语言，语法如下

![image](https://intranetproxy.alipay.com/skylark/lark/0/2023/png/23156369/1681463869053-ce81eaec-1d3a-4c97-a0f5-ba21ec336909.png#clientId=u11e18562-e6d8-4&from=paste&id=ua91f08e5&originHeight=480&originWidth=850&originalType=url&ratio=1&rotation=0&showTitle=false&status=done&style=none&taskId=u612a1b31-df21-4e4d-9b1b-d2a4dcb8fae&title=)

上图的结构被称为 CSS ruleset，由 2 部分组成

1. Selector

   属性选择器

2. Declaration

   由两部分组成

      - Property

        属性选择器 选中的 HTML element 对应的属性

      - Property Value

        对应属性的值
        每个 Declaration 需要以 `;` 结尾，同时 Property 和 Property Value 需要使用 `:` 分隔

property 不同，value 的可取值也不同。在一个 CSS stylesheet 中会包含许多 rule sets。如果 property is unknown or invalid , 该 property 的 value 就会被忽略，但是对应的 rule set 中其他的 property 还是有效的

### shorthands

在 CSS properties 有一些 shorthand properties, 例如 `<font>`, `<background>`, `<border>`, `margin` 等，也可以拆分层 longhand properties

```
/* in the order top, right, bottom, left */
padding: 10px 15px 15px 5px;
padding-top: 10px;
padding-right: 15px;
padding-bottom: 15px;
padding-left: 5px;

background: red url(bg-graphic.png) 10px 10px repeat-x fixed;
background-color: red;
background-image: url(bg-graphic.png);
background-position: 10px 10px;
background-repeat: repeat-x;
background-attachment: fixed;
```

## CSS specifications

直接理解成类似 Network RFC，是行业内的权威机构（W3C,ECMA,WHATWG）发布一些条文规定 CSS behavior

## CSS moudles

CSS 是一个庞大的系统语言，为了方便管理将功能相近或者类似的 properties  划分为 seperated moudles

例如 `background-color` 和 `border-color` 就是一类属于 color

具体分类参考 MDN reference

https://developer.mozilla.org/en-US/docs/Web/CSS



**references**

1. https://developer.mozilla.org/en-US/docs/Learn/CSS/First_steps/What_is_CSS
1. [https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics)
