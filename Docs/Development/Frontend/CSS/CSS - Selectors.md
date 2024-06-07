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
ID 选择器，按照标签的 `id` 属性匹配。==ID 全局唯一，否则错误(假命题)[^9]==

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

`.className` 会选中 `<element class="className">` 或者 `<element class="className ...">`，即只要 `class` 属性的值包含 `className` 即可
例如

```
.c1 {
	color: red;
}
```
会选中 `p2` 和 `p3`

有一种特殊的情况，假设 class name 中包含空格

例如

```
<p class=" test-class">this is a test message</p>
```

在 CSS 中会自动去掉 class name 的前导和后导空格，所以只需要使用 `.test-class` 就可以选中上述元素

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

pseudo-class 也被称为伪类，用于表示处于特殊状态的 HTML 元素，可以将其想象成新加一个 class

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

- `:has(selector)`

  选中元素包含指定的 selector

- `:not(selector)`

  选中元素不包含指定的 selector

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

pseudo-element 也被称为伪元素，功能和 pseudo-class 类似，但是 pseudo-element 就好像是选中元素内的一个新 HTML 元素 ( 意味着有独立的 box model )

伪元素选择器，即为选中的 HTML 元素内( 这也表示了 `<input>` 不能很好的使用这些伪元素选择器，因为 `<input>` 是一个 void element )添加一个伪元素，为该元素添加层叠样式。通过 `::pseudo-element` 方式来指定选中的 HTML 元素 (早的时候也可以用 `:pseudo-element` 的方式来指定)

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

- `.class1.class2.class3`
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
    <span class="s4 s5">s45</span>
    <div a1="v1" a2="v2">d2</div>
    <div class="d3">d3</div>
</div>
<div class="d2">d2</div>
<div>d3</div>
```
为了方便表述，以下内容等价

- `d1 == <div class="d1">d1</div>`
- `s1 == <span class="s1">s1</span>`
- `s2 == <span class="s2">s2</span>`
- `s3 == <span class="s3">s3</span>`
- `s45 == <span class="s4 s5">s45</span>`
- `d2 == <div a1="v1" a2="v2">d2</div>`
- `d3 == <div class="d3">d3</div>`

这里只两两组合实际可以任意个数的组合

### classes selector

`.className1.className2`

选中同时包含 `.className1` 和 `.className2` 的元素，逻辑与

例如 如下表示选中同时包含 `s4` class 和 `s5` class 的元素，即 `s45`

```
.s4.s5 {
	color: red;
}
```

### attribute selectors selector

`[attr1=value1][attr2=value2]`，逻辑与

选中同时含有 `attr1=value1` 和 `attr2=value2` 的元素

例如 如下表示选中同时包含 `a1=v1` 和 `a2=v2` 的元素，即 `d2`

```
[attr1="value1"][attr2="value2"] {
	color: red;
}
```

### descendant selector

`selector1 selector2`

选中匹配 selector1 元素下，所有匹配 selector2 的 selector1 ==子孙元素==
例如 如下表示选中 `div` 下子孙级的 `span`，即 `s1`, `s2`, `s3`, `s45`

```
div span {
	color: red;
}
```
### list selector

`selector1, selector2`

选中匹配 selector1 元素，或匹配 selector2 的元素，逻辑或

例如 如下表示选中包含 `d3` class 或 包含 `s1` class 的元素，即 `d3` 和 `s1`

```
.d3,.s1 {
	color: red;
}
```
但是上述 `s2` 的样式也会修改。因为是 `s1` 的子孙元素，在 CSS 中子孙元素的 CSS 样式会继承祖先元素的(这个逻辑不一定对，具体需要看元素和样式)

### child selector

`selector1>selector2`

选中匹配 selector1 元素下匹配 selector2 且是 selector1 的==子元素==
例如 如下表示选中 `div` 下所有子 `span` 元素，即 `s1`, `s3`, `s45`。这里不会匹配 `s2` 是因为不是直接的子元素，是子孙元素

```
div>span {
	color: red;
}
```
但是上述 `s2` 的样式也会修改。因为是 `s1` 的子孙元素，在 CSS 中子孙 CSS 样式会继承祖先元素的

### adjacent sibling selector

`selector1 + selector2`

选中匹配 selector1 元素之后，匹配 selector2 所有元素( sibling )的第一个元素

例如 如下表示选中包含 `s1` class 的元素之后的第一个 `span` ，即 `s3`

```
.s1 + span {
	color: red;
}
```
### general sibling selector

`selector1 ~ selector2`

选中匹配 selector1 元素之后，匹配 selector2 所有元素( siblings )

例如 如下表示选中  `div` 之后的所有 `div`, 即 `d2` 和 `d3`

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
8. https://stackoverflow.com/questions/2587669/can-i-use-a-before-or-after-pseudo-element-on-an-input-field
9. [^9]:https://softwareengineering.stackexchange.com/questions/127178/two-html-elements-with-same-id-attribute-how-bad-is-it-really
