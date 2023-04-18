CSS - Box model

在 CSS 中将所有的 HTML 元素抽象成 Boxes，根据 Boxes 默认的显示特性，分为两类

1. block boxes
2. inline boxes
## Display
在了解 Box model 前，需要了解一下 Boxes 的显示特性。每一个 Box 都有默认的显示特性，特性不同，Box 的显示方式不同。可以通过元素的 `display` 来手动设定
根据针对的对象不同，分为两类

1. Outer display - 针对元素本身的显示方式，有
   - `block` 对应 block boxes 
   - `inline` 对应 inline boxes
   - `inline-block`
2. Inner display - 针对元素内部元素的显示方式，有
   - `flex`
### Outer display
指明元素本身以什么方式显示
#### block
元素默认的 `display` 值为 `block`，或者元素通过层叠样式手动设定 `display` 值为 `block`，即为 block boxes
有如下特性

1. the box will break onto a new line (可以理解成对应的元素会在结尾使用换行符)
2. the `width` and `height` properties are respected
3. padding, margin, border will cause other elements to be pushed away from the box
4. if `width` is not specified, the box will extend in the inline direction to fill the space avaliable in its container. In most cases, the box will become as wide as its container, filling up 100% of the space available (没有标识 `width` 值，默认为 100%)

例如 `<h1>`, `<p>`, `<div>` 等, 都默认使用 `block` 作为 `display` 的值
#### inline
元素默认的 `display` 值为 `inline`，或者元素通过层叠样式手动设定 `display` 值为 `inline`，即为 inline boxes
有如下特性

1. the box will not break onto a new line
2. the `width` and `height` properties will not apply (即使设置了`width` 和 `height` 也不会生效)
3. top and bottom padding, margins, and borders will apply but will not cause other inline boxes to move away from the box (对元素本身同样生效，如果其他的 boxes 是 block boxes，同样会生效)
4. left and right padding, margins, and borders will apply and will cause other inline boxes to move away from the box (对元素本身同样生效，如果其他的 boxes 是 inline boxes，默认只有左右向的 padding, margin, border 会生效)

例如 `<a>`, `<span>`, `<em>` 等, 都默认使用 `inline`作为 `display` 的值
#### inline-block
介于 `inline` 和 `block` 之间，只有通过手动设置 `display` 为 `inline-block` 实现
有如下特性

1. the box will not break onto a new line
2. the `width` and `height` properties are respected
3. padding, margin, border will cause other elements to be pushed away from the box (不管是 inline boxes 或者是 block boxes，任意方向的 padding, margin, border 都会生效)
### Inner display
指明元素内部元素以什么方式显示
#### flex
所有的 HTML 元素默认都会以 `block` 或者 `inline` 的方式显示，如果需要使用 inner display 需要手动指定 `display` 的值为 `flex`。所有当前元素下的==子元素==都会变成 flex items，但是并不改变 items 本身的 display 值
## Box model
> 在 CSS 中只有 block boxes 具有完整的 Box model, inline boxes 只有 Box model 中的部分

在 CSS 中每一个 HTML 元素都被抽象成 Box，用 Box model 来规范 Box 以及 Box 之间显示的方式
![image.png](https://intranetproxy.alipay.com/skylark/lark/0/2023/png/23156369/1681804486550-919794ad-fa4b-4b7b-a24d-cdd5f244314d.png#clientId=ue2c2e752-ccff-4&from=paste&id=ud14f9d4c&name=image.png&originHeight=300&originWidth=544&originalType=url&ratio=1&rotation=0&showTitle=false&size=11822&status=done&style=none&taskId=u560191a9-43c4-4041-9ad6-a3d8f534272&title=)
一个完整的 Box model 由几个部分构成

1. Content

the area where you content is displayed
关联 `inline-size`, `block-size`, `width` 和 `height` 等属性

2. Padding

the space between the content and border
关联 `padding` 等属性

3. Border

the solid line that is just outside the padding
关联 `border` 等属性

4. Margin

the space aroud the outside of the border
关联 `margin` 等属性
### Example
例如 一个 block box 的层叠样式如下
```
.box {
  width: 350px;
  height: 150px;
  margin: 10px;
  padding: 25px;
  border: 5px solid black;
}
```
![Snipaste_2023-04-18_16-49-44.png](https://intranetproxy.alipay.com/skylark/lark/0/2023/png/23156369/1681807865014-6b0e72b1-3596-4bd3-ad5c-b34f08d4c96c.png#clientId=ue2c2e752-ccff-4&from=paste&height=191&id=u3de6f6fa&name=Snipaste_2023-04-18_16-49-44.png&originHeight=191&originWidth=430&originalType=binary&ratio=1&rotation=0&showTitle=false&size=20799&status=done&style=none&taskId=ub73355f0-a834-487a-b9c0-d69e90b5cdc&title=&width=430)
那么整个  box 大小为 
![](https://intranetproxy.alipay.com/skylark/lark/__latex/0b49796dd20b0ebfc4f3f99e2ce2eeff.svg#card=math&code=wide%20%3D%20410px%3D%20width%28350px%29%20%2B%20padding%2825px%29%20%5Ctimes%202%20%2B%20border%285px%29%20%5Ctimes%202&id=FaIOz)
![](https://intranetproxy.alipay.com/skylark/lark/__latex/680a66841f9640a33d2c474fb6b3665d.svg#card=math&code=high%20%3D%20210%20px%20%3D%20height%28150px%29%20%2B%20padding%2825px%29%20%5Ctimes%202%20%2B%20border%285px%29%20%5Ctimes%202&id=Qv0j3)
这里并没有将 `margin` 计算在内，虽然 `margin` 会影响 box model 从而影响前端显示，但是在 box 外
### border-box
`box-sizing` 属性中有一个特殊的值 `border-box`。使用该值，high 即 `height`，wide 即 `width`
借用上一个例子
```
.box {
  width: 350px;
  height: 150px;
  margin: 10px;
  padding: 25px;
	box-sizing: border-box;
  border: 5px solid black;
}
```
那么整个  box 大小为 
![image.png](https://intranetproxy.alipay.com/skylark/lark/0/2023/png/23156369/1681807965910-bf78eb98-e102-455b-b0e9-8fa735716df5.png#clientId=ue2c2e752-ccff-4&from=paste&height=194&id=ud4392cc0&name=image.png&originHeight=194&originWidth=421&originalType=binary&ratio=1&rotation=0&showTitle=false&size=22827&status=done&style=none&taskId=u5b993977-b3bd-4c3c-9812-54437920c25&title=&width=421)
![](https://intranetproxy.alipay.com/skylark/lark/__latex/cb500d9dadfe0c0417afb62e402d9e13.svg#card=math&code=wide%20%3D%20width%28350px%29%20%3D%20padding%2825px%29%20%5Ctimes%202%20%2B%20border%285px%29%20%5Ctimes%202%20%2B%20auto%5C_width&id=SzM0p)
![](https://intranetproxy.alipay.com/skylark/lark/__latex/20fcaa1a938460cb082af97720a0a577.svg#card=math&code=high%20%3D%20height%28150px%29%20%3D%20padding%2825px%29%20%5Ctimes%202%20%2B%20border%285px%29%20%5Ctimes%202%20%2B%20auto%5C_height&id=dIdVG)
可以想象成 `border` 和 `padding` 都在 `height` 和 `width` 内
### Inline box model
因为 inline box 的 `width` 和 `height` 属性并不会生效 ，且 `padding`，`border`，`margin` 在 inline box 间两两互不影响，完整的 box model 并不适用于 inline box
例如
```
span {
  margin: 200px;
  padding: 20px;
  width: 80px;
  height: 50px;
  background-color: lightblue;
  border: 2px solid blue;
}
----
<p>
    I am a paragraph and this is a <span>span</span> inside that paragraph. A span is an inline element and so does not respect width and height.
</p>
```
上述修改 `maring` 只有 `margin-left`, `margin-right` 为 `200px`；修改 `padding` 只有 `padding-left`, `padding-right` 为 `200px`；修改 `width` 和 `height` 不会生效；修改 `border` 会生效

**references**

1. [https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/The_box_model](https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/The_box_model)
2. [https://developer.mozilla.org/en-US/docs/Learn/CSS/CSS_layout/Flexbox](https://developer.mozilla.org/en-US/docs/Learn/CSS/CSS_layout/Flexbox)
3. [https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics#css_all_about_boxes](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics#css_all_about_boxes)
