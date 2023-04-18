# CSS - Selectors

CSS selector 中文也被称为选择器，用于选中指定的 HTML 元素，然后对选择的元素做属性赋值，即修改层叠样式

## Basic selectors
CSS 有多种基础的选择器

- type selector
- universal selector
- id selector
- class selector
- attribute selector
- pseudo-class selector

以如下内容为例
```
<p>this is p1</p>
<p class="c1">this is p2</p>
<p class="c2 c1">this is p3</p>
<p id="id1">ths is p4</p>
<div>this a d1</div>
<span>this a s1</span>
<p attr="my-attr1">this is pa1</p>
<p attr="my-attr2 my-attr1">this pa2</p>
<span attr="my-attr3">this is sa1</span>
<div attr="my-attr-value">this is da1</div>
```
为了方便描述，如下内容等价

- `p1 == <p>this is p1</p>`
- `p2 == <p class="c1">this is p2</p>`
- `p3 == <p class="c2 c1">this is p3</p>`
- `p4 == <p id="id1">ths is p4</p>`
- `d1 == <div>this a d1</div>`
- `s1 == <span>this a s1</span>`
- `pa1 == <p attr="my-attr1">this is pa1</p>`
- `pa2 == <p attr="my-attr2 my-attr1">this pa2</p>`
- `sa1 == <span attr="my-attr3">this is sa1</span>`
- `da1 == <div attr="my-attr-value">this is da1</div>`
### Element selector
元素选择器，也被称为 type selector (类型选择器)。按照 HTML 元素(标签) 匹配

`element` 匹配 `<element>`
例如

```
p {
	color: red;
}
```
会选中所有的 `p` 标签
### Universal selector
通配符选择器，选择所有的元素
例如
```
* {
	color: red;
}
```
会选中所有的 `p` 和 `div` 以及 `span` 标签
### ID selector
ID 选择器，按照标签的 `id` 属性匹配。ID 全局唯一，否则错误

`#idName` 选中 `<element id="idName">`

例如

```
#id1 {
	color: red;
}
```
会选中 `p4`
### Class selector
类选择器，按照标签的 `class` 属性匹配。`class` 属性的值可以是单值，也可以是列表 (以空格分隔)

`.className` 会选中`<element class="className">` 或者 `<element class="className ...">`，即只要 `class` 属性的值包含 `className` 即可
例如

```
.c1 {
	color: red;
}
```
会选中 `p2` 和 `p3`

### Atrribute selector

> `*` 表示 wildcard
>
> `... `表述属性的值是变参

属性选择器，按照标签的属性匹配

#### presence match

- `[attr]` 

  匹配任意含有 `attr` 属性的 HTML 元素。会选中 `<element attr="*">`

  例如

  ```
  [attr] {
  	color: red;
  }
  ```

  会选中 `pa1`,  `pa2`，`sa1` 以及 `da1`

- `[attr=attrValue]`

  匹配含有 `attr` 属性，并且值为 `attrValue` 的 HTML  元素。 会选中 `<element attr="attrValue">`

  例如

  ```
  [attr="my-attr1"] {
  	color: red;
  }
  ```

  会选中 `pa1`，这里不会选中 `pa2`，因为需要值一致

- `[attr~=attrValue]` 

  匹配含有 `attr` 属性，并且值( 也可以是以空格分隔的列表  )包含 `attrValue` 的 HTML 元素。会选中 `<element attr="attrValue">` 或者 `<element attr="... attrValue ...">` 

  例如

  ```
  [attr~="my-attr1"] {
  	color: red;
  }
  ```

  会选中 `pa1` 和 `pa2`

- `[attr|=attrValue]` 

  匹配含有 `attr` 属性，并且值( 值也可以是以空格分隔的列表  )以 `attrValue`  开头， 后面必须是 `hyphen` 的 HTML 元素。会选中 `<element attr="attrValue">` 或者 `<element attr="... attrValue-* ...">` 

  例如

  ```
  [attr|="my-attr"] {
  	color: red;
  }
  ```

  会选中 `da1`

#### substring match

- `[attr^=attrValue]` 

  匹配含有 `attr` 属性，并且值( 值也可以是以空格分隔的列表 )以 `attrValue` 开头的 HTML 元素。会选中 `<element attr="attrValue">` 或者 `<element attr="... attrValue* ...">` 。理解成正则表达式中的 `^`

  例如

  ```
  [attr^="my-attr"] {
  	color: red;
  }
  ```

  会选中  `pa1`, `pa2`, `sa1 `以及 `da1`

- `[attr$=attrValue]`

  匹配含有 `attr` 属性，并且值( 值也可以是以空格分隔的列表 )以 `attrValue` 结尾的 HTML 元素。会选中 `<element attr="attrValue">` 或者 `<element attr="... *attrValue ...">` 。理解成正则表达式中的 `$`

  例如

  ```
  [attr$="value"] {
  	color: red;
  }
  ```

  会选中 `da1`

- `[attr*=attrValue]`

  匹配含有 `attr` 属性，并且值( 值也可以是以空格分隔的列表 )包含 `attrValue` 的 HTML 元素。会选中 `<element attr="attrValue">` 或者 `<element attr="... *attrValue* ...">` 。理解成正则表达式中的 `*`

  ```
  [attr*="attr1"] {
  	color: red;
  }
  ```

  会选中 `pa1` 和 `pa2`

### Pseudo-class selector

> 具体可以参考 MDN Pseudo-classes reference 部分

pseudo-class 也被称为伪类，用于表示处于特殊状态的 HTML 元素

伪类选择器，即当选中的 HTML 元素处于指定状态时生效，通过 `:pseudo-class` 方式来指定伪类

例如

```
<a href="https://example.org">link</a>
---
a:hover {
	color: red;
}
```
上述表示当 `a` HTML element 处于 hover 状态时，颜色改成红色

常见的 pseudo-class selectors 有

- `:last-child`

  选中元素的最后一个子元素

- `:first-child`

  选中元素的第一个子元素

- `:hover`

  鼠标位于选中元素之上

- `:focus`

  鼠标点击选中元素

- `:link`

  未被点击的链接

- `:visited`

  被点击后的链接

### Pseudo-element selector

> 具体可以参考 MDN Pseudo-elements reference 部分

pseudo-element 也被称为伪元素，功能和 pseudo-class 类似，但是 pseudo-element 就好像是一个新的 HTML 元素

伪元素选择器，即选中的 HTML 元素内像有一个匹配的伪元素，为该元素添加层叠样式，通过 `::pseudo-element` 方式来指定 (早的时候也可以用 `:pseudo-element` 的方式来指定)

```
article p::first-line {
    color: red;
}   
----
<article>
    <p>Veggies es bonus vobis, proinde vos postulo essum magis kohlrabi welsh onion daikon amaranth tatsoi tomatillo melon azuki bean garlic.</p>
    <p>Gumbo beet greens corn soko endive gumbo gourd. Parsley shallot courgette tatsoi pea sprouts fava bean collard greens dandelion okra wakame tomato. Dandelion cucumber earthnut pea peanut soko zucchini.</p>
</article>   
```

例如 上述会将选中的 `article p` 中的第一行作为伪元素，将颜色改成红色

常见的 pseudo-element selector 有

- `::first-line`

  选中元素的第一行

- `::first-letter`

  选中元素的第一单词

- `::before`

  在选中的元素内，在 `<element>` 后，元素内容之前插入 `content` 内容以及对应的 ruleset

- `::after`

  在选中的元素内，在 `</element>` 前，元素内容之后插入 `content` 内容以及对应的 ruleset

`::before` 和 `::after` 的使用比较特殊，需要在 ruleset 中使用 `content` property

例如

```
<div id="id1">
    div
</div>
---
#id1::before {
    content:"this a paragraph before div";
    color: red;
}
#id1::after {
    content:"this a paragraph after div";
    color: green;
}
```

上述会在，`#id` 中插入对应 content 中的内容

```
<div id="id1">
   ::before
    div
    ::after
</div>
```

`::before` 替换成红色的 `this a paragraph before div`， `::after` 替换成绿色的 `this a paragraph after div`

## Combined Selectors

basic selectors 之间可以互相组合，例如

- `p[attr*=value],div`
- `#id[attr~=value] p`
- `.class[attr]::after`
- `p:first-child::first-line`
- `div>*`

假设有如下元素

```
<div class="d1">
    d1
    <span class="s1">s1<span class="s2">s2</span></span>
    <span class="s3">s3</span>
    <div class="d3">d3</div>
</div>
<div class="d2">d2</div>
<div>d3</div>
```
这里只两两组合实际可以任意个数的组合

### descendant selector

`selector1 selector2`

选中匹配 selector1 元素下，所有匹配 selector2 的 selector1 ==子孙元素==
例如 如下表示选中 `div` 下子孙级的 `span`，即为 `<span>s1<span>s2</span></span>` 和`<span class="s3">s3</span>`

```
div span {
	color: red;
}
```
### list selector

`selector1, selector2`

选中匹配 selector1 元素，和匹配 selector2 的元素
例如 如下表示选择 `<div class="d2">` 和 `<span class="s1">s1</span>`

```
.d2,.s1 {
	color: red;
}
```
但是上述 `<span class="s2">s2</span>` 的样式也会修改。因为是 `<span class="s3">s3</span>` 的子孙元素，在 CSS 中子孙 CSS 样式会继承祖先元素的

### child selector

`selector1>selector2`

选中匹配 selector1 元素下匹配 selector2 且是 selector1 的==子元素==
例如 如下表示选中 `div` 下所有子 `span` 元素，即 `<span class="s1">s1</span>` 和 `<span class="s3">s3</span>`，这里不会匹配 `s2` 是因为不是直接的子元素，是子孙元素

```
div>span {
	color: red;
}
```
但是上述 `<span class="s2">s2</span>` 的样式也会修改。因为是 `<span class="s3">s3</span>` 的子孙元素，在 CSS 中子孙 CSS 样式会继承祖先元素的

### adjacent sibling selector

`selector1 + selector2`

选中匹配 selector1 元素之后，匹配 selector2 所有元素( sibling )的第一个元素

例如 如下表示选中 `<span class="s3">s3</span>`

```
.s1 + span {
	color: red;
}
```
### general sibling selector

`selector1 ~ selector2`

选中匹配 selector1 元素之，匹配 selector2 所有元素( siblings )

例如 如下表示选中 `<div class="d2">d2</div> ` 和 `<div>d3</div>`

```
div ~ div {
	color: red;
}
```

## Cautions

只要有一个 selector 存在语法错误，对应的 ruleset 都会失效。假设 selector 如下
```
..d2,div {
	color: red;
}
```
因为 `..d2` 语法错误，会导致整条 ruleset 失效



**referneces**

1. [https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Selectors](https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Selectors)
2. [https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Selectors](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Selectors)
3. [https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Selectors/Type_Class_and_ID_Selectors](https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Selectors/Type_Class_and_ID_Selectors)
4. [https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Selectors/Attribute_selectors](https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Selectors/Attribute_selectors)
5. https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Selectors/Pseudo-classes_and_pseudo-elements
6. https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Selectors/Combinators
7. [https://www.ruanyifeng.com/blog/2009/03/css_selectors.html](https://www.ruanyifeng.com/blog/2009/03/css_selectors.html)
