# CSS - Inheritance override

CSS 也支持继承和重载的特性，父级元素匹配的 ruleset 中的一些 property 会被子级元素继承，子级元素可以在匹配的 ruleset 中声明对应的 property 进行重载

例如，`body` HTML 元素声明使用蓝色，而 `span` HTML 元素声明使用黑色

```
body {
    color: blue;
}

span {
    color: black;
}
 ---  
<body>
<p>As the body has been set to have a color of blue this is inherited through the descendants.</p>
<p>We can change the color by targeting the element with a selector, such as this <span>span</span>.</p>
</body>
```

因为 `body` 是 `p` 和 `span` 的父级元素，`p` 和 `span` 就会继承  `body` 对应的 ruleset，所以应该显示蓝色。但是 `span`  对应的 ruleset 设置了自己的 `color`， 所以对从 `body` 继承的 `color` 做了 override，最后显示黑色

## Controlling inheritance

CSS 为每个 HTML 元素 property，提供了 5 个内置的 properties value 用于控制继承

- `inherit`

  对应的 property 的值从父级元素继承，等价于 turns on inheritance

- `initial`

  对应的 property 的值使用原始的样式

- `revert`

  和 `unset` 类似，直接理解成 `unset` (实际不同，具体参考 MDN 文档)

- `revert-layer`

- `unset`

  如果对应的 property 不是从父级元素继承的，使用默认的样式。如果对应的 property 是从父级元素继承的，使用父级元素的样式

### All property

CSS 中有一个特殊的 property -- `all` 指代所有的从父级元素继承的 properties，可以通过该值快速有效的管理继承的效果

```
div {
    color: red;
    font-size: 40px
}
p {
    all: initial;
}
---
<div>
    div
    <p>p</p>
</div>
```

## Caution

特别需要注意的是和其他编程语言在的继承不一样，CSS 不会继承父级元素所有的 property。例如 子级元素不会继承父级元素中的 `width:50%`，设想一下如果可以继承，那么 CSS 将会非常复杂

除 `width` 外，诸如  `margin`,`padding`,`border` , etc 都不会被继承



**references**

1. https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Cascade_and_inheritance
2. https://stackoverflow.com/questions/33834049/what-is-the-difference-between-the-initial-and-unset-values