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

**references**

1. https://developer.mozilla.org/en-US/docs/Learn/CSS/First_steps/How_CSS_is_structured

