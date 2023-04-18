# CSS - Unit

## Absolute length units

绝对长度

| Unit | Name                | Equivalent to            |
| ---- | ------------------- | ------------------------ |
| `cm` | Centimeters         | 1cm = 37.8px = 25.2/64in |
| `mm` | Millimeters         | 1mm = 1/10th of 1cm      |
| `Q`  | Quarter-millimeters | 1Q = 1/40th of 1cm       |
| `in` | Inches              | 1in = 2.54cm = 96px      |
| `pc` | Picas               | 1pc = 1/6th of 1in       |
| `pt` | Points              | 1pt = 1/72nd of 1in      |
| `px` | Pixels              | 1px = 1/96th of 1in      |

常用的一般有 `px`

## Relative length units

CSS 有一些相对长度的单位

| Unit       | Relative to                                                  |
| ---------- | ------------------------------------------------------------ |
| `em`       | Font size of ==the parent==, in the case of typographical properties like        `font-size`, and font size of the element itself, in the case of other properties        like `width`. |
| `ex`       | x-height of the element's font.                              |
| `ch`       | The advance measure (width) of the glyph "0" of the element's font. |
| `rem`      | Font size of the root element.                               |
| `lh`       | Line height of the element.                                  |
| `rlh`      | Line height of the root element. When used on the `font-size` or `line-height` properties of the root element, it refers to the properties' initial value. |
| `vw`       | 1% of the viewport's width.                                  |
| `vh`       | 1% of the viewport's height.                                 |
| `vmin`     | 1% of the viewport's smaller dimension.                      |
| `vmax`     | 1% of the viewport's larger dimension.                       |
| `vb`       | 1% of the size of the initial containing block in the direction of the root element's [block axis](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Logical_Properties#block_vs._inline). |
| `vi`       | 1% of the size of the initial containing block in the direction of the root element's [inline axis](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Logical_Properties#block_vs._inline). |
| `svw, svh` | 1% of the [small viewport](https://developer.mozilla.org/en-US/docs/Web/CSS/length#relative_length_units_based_on_viewport)'s width and height, respectively. |
| `lvw, lvh` | 1% of the [large viewport](https://developer.mozilla.org/en-US/docs/Web/CSS/length#relative_length_units_based_on_viewport)'s width and height, respectively. |
| `dvw, dvh` | 1% of the [dynamic viewport](https://developer.mozilla.org/en-US/docs/Web/CSS/length#relative_length_units_based_on_viewport)'s width and height, respectively. |

常用的一般有 `em`, `rem`, `vw`, `vh`

其中 viewpoint 含义为

> *which is the visible area of your page in the browser you are using to view a site*

## Percentages

length 一般还支持 percentage 的方式，表示相对父元素的值

例如

```
.outer{
    font-size: 100px;
}
.inner{
    font-size: 50%;
}
---
<div class="outer">
    outer text
    <div class="inner">inner text</div>
</div>
```

上述表示 `inner` 对应的字体大小是 `outer` 的一般，即 `50px`

**references**

1. https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Values_and_units
2. https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Sizing_items_in_CSS#viewport_units