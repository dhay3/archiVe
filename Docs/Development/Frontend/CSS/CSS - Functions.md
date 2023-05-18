# CSS - Functions

> 具体 function 查看 MDN references 部分

## var

用于调用自定义的变量

syntax

```
var(<custom-property-name>,[fallback-value])
```

- custom-property-name

  自定义的变量名

- fallback-value

  如果自定义变量的值无效，使用该值

例如

```
:root {
  --main-bg-color: pink;
}

body {
  background-color: var(--main-bg-color);
}
```

上述即 `background-color` 的值就是 `--main-bg-color` 的值

## url

用于调用 url，一般用在 `background` 或者 `background-image`

syntax

```
url(<URI>)
```

例如

```
background-image: url("star.gif");
```

## attr

用于获取选中元素对应的属性

syntax

```
attr(<attr-name>)
```

例如

```
[data-foo]::before {
  content: attr(data-foo) " ";
}
---
<p data-foo="hello">world</p>
```

## repeat



## calc

使用该函数可以进行数学计算

syntax

```
calc(<arg1> operator <arg2>)
```

例如

```
.outer {
  border: 5px solid black;
}

.box {
  padding: 10px;
  width: calc(90% - 30px);
  background-color: rebeccapurple;
  color: white;
}
---
<div class="outer"><div class="box">The inner box is 90% - 30px.</div></div>
```

例如上述 `.inner` 的 `width` 的值就是 `width(.outer)*90%-30px`

## has

一般用在 selectore 中，选中 parent elements 或者是 sibling elements

syntax

```css
:has(<relative-selector-list>) {
  /* ... */
}
```

例如

```css
/* Selects an h1 heading with a
paragraph element that immediately follows
the h1 and applies the style to h1 */
h1:has(+ p) {
  margin-bottom: 0;
}
```

会选中一个 `h1` element, 其子元素的第一个元素是 `p`

## not

一般用于 selector 中，表示选择的 elements 不能包含 not 中的内容

syntax

```css
:not(<complex-selector-list>) {
  /* ... */
}
```

例如

```
div:not(.div2){
	color: red;
}
---
<div class="div1">div1</div>
<div class="div2">div2</div>
```

只在选中不包含 `class="div2"` 的元素，即 `<div class="div1">div1</div>`

**references**

1. https://developer.mozilla.org/en-US/docs/Learn/CSS/First_steps/How_CSS_is_structured
