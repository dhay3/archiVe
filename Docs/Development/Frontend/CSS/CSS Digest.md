# CSS Digest

ref

https://developer.mozilla.org/en-US/docs/Learn/CSS/First_steps/What_is_CSS

https://developer.mozilla.org/en-US/docs/Learn/CSS/First_steps/How_CSS_is_structured

Cascading Style Sheets( CSS )  可以理解成是一种语言用于 marking up HTML

## Syntax

CSS is a rule-based  language - you define the rules by specifying groups of styles that should be applied to particular elements or groups of elements on your web page

```
h1 {
  color: red;
  font-size: 5em;
}
```

rule 由以下几部分组成

1. selector

   which selects the HTML element that we are going to style, in this case `h1` is the selector

2. curly barces

3. declarations

   a set property and value parirs inside the curly braces, in this case 

   ` color: red;
   font-size: 5em;`

   is the declarations

   and color and font-size are proerties, red and 5em are values

根据 properties 不同，value 的值也不同，在一个 CSS stylesheet 中会包含许多 rules

## CSS specifications

直接理解成 Network RFC，是行业内的权威机构（W3C,ECMA,WHATWG）发布一些条文规定 CSS behavior

## Proerties and values

- properties

  these are human-readable identifiers that indicate which stylistic features you want to modify. For example, color, width, font-size

- values

  each property is assigned a value. This value indicates how to style the proerty

需要注意的一点是如果 property is unknown or invalid , 该 property 的 value 就会被忽略

## Moudles

CSS 是一个庞大的系统语言，为了方便管理将功能相近或者类似的 properties  划分为 seperated moudles

具体参考 MDN reference

https://developer.mozilla.org/en-US/docs/Web/CSS

## Shorthands

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



## Adding CSS

有 3 种方式来添加 CSS

- External stylesheet

  ```
  <!DOCTYPE html>
  <html lang="en-GB">
    <head>
      <link rel="stylesheet" href="styles.css" />
    </head>
  </html>
  ```

  在`<head>` element 中通过 link href 的方式导入 CSS 文件

- Internal CSS

  ```
  <!DOCTYPE html>
  <html lang="en-GB">
    <head>
      <style>
        p {
          color: red;
        }
      </style>
    </head>
    <body>
      <p>This is my first CSS example</p>
    </body>
  </html>
  ```

  直接在本文件中添加 rule set

  如果是 chunk of sites 且 sites 需要使用 uniform CSS，这时候 internal CSS 效率就会很低，所以推荐使用 Internal CSS

- inline styles

  ```
  <p style="color:red;">this is p</p>
  ```

  直接在元素上添加CSS

  尽量避免使用这种方式，这是一种效率最低的方式

## Adding a class

adding a class to HTML elements is a way to apply CSS for the bunch of elements

例如下面的 HTML snippet

```
<ul>
  <li>Item one</li>
  <li class="special">Item two</li>
  <li>Item <em>three</em></li>
</ul>
```

这样就可以使用 CSS class selector 选中 elements 并赋予 CSS style

```
.special {
  color: orange;
  font-weight: bold;
}
```

