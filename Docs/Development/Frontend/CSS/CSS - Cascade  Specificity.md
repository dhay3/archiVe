# CSS - Cascade Specificity

在 CSS 中不同的 selector 可能会选中在同一个 HTML 元素，但是 selector 对应的 ruleset 内容不同

例如，以下内容会选中同一个 `p` 元素

```
<style>
.special {
  color: red;
}

p {
  color: blue;
}
</style>
<p class="special">What color am I?</p>
```

这时 CSS 会使用那个 ruleset ？ 红色还是蓝色呢？

在 CSS 中会有 4 个因素决定，优先级从低到高

1. Source code order
2. Specificity
3. Inline style
4. !important

## Source code order

> When two rules from the same cascade layer apply and both have equal  specificity, the one that is defined last in the stylesheet is the one  that will be used.

按照 ruleset 声明的先后顺序( 即在源码中的位置 )决定优先级，最后声明的 ruleset 优先级最高

假设现在有 两个 p element selector 同时声明使用红色和蓝色

```
<style>
p {
  color: red;
}

p {
  color: blue;
}
</style>
<p class="special">What color am I?</p>
```

因为两个 p element selector 具有一样的 specificity，那么就会使用最后声明的 rulset，所以最后渲染的结果就是蓝色

伪代码逻辑如下

```
if ruleset1.specificity == ruleset2.specificity then
    if ruleset1.order < ruleset2.order then
       	  return ruleset2
     else
     	return ruleset1
```

## Specificity

假设 现在一个使用 element selector ，另外一个使用 class selector，结果有会怎么样呢？

```
<style>
.special {
  color: red;
}

p {
  color: blue;
}
</style>
<p class="special">What color am I?</p>
```

CSS 会优先采用 `.special` 对应的 ruleset。因为 element selector 优先级最低，而 class selector 优先级比较高，所以最后渲染的颜色为红色

这里的优先级就是 Specificity，指的是描述的详细程度，是 CSS 预设的逻辑规则。按照不同的 selector 赋予不同的权重，权重值越高，优先级越高

可以分为 3 类

> 这里为了方便描述给权重一个具体的值，但是需要说明的是 `Identifiers == classes * 10 == elements * 100` 关系式并不等价,  实际应为 `Identifiers > classes * 10 > elements * 100`

- Identifiers

  对应 id selector，权重 100

- Classes

  对应 class selector, attribute selector, pseudo-class selector，权重 10

- Elements

  对应 element(type) selector, pseudo-element selector，权重 1

需要注意的两点是

1. 组合 selector 的符号 (例如 `+`, `>` 等等)不会影响权重
2. `:not()`, `:has()`, `:is()` 等伪类选择器，不会影响权重。但是对应的入参会影响

例如

| Selector                                  | Identifiers | Classes | Elements | Total specificity | Comment                                                      |
| ----------------------------------------- | ----------- | ------- | -------- | ----------------- | ------------------------------------------------------------ |
| `h1`                                      | 0           | 0       | 1        | 0-0-1             | `h1` 是一个元素选择器，属于 elements。所以最后的权重得分为 1 |
| `h1 + p::first-letter`                    | 0           | 0       | 3        | 0-0-3             | `h1 ` , `p`, 是元素选择器，`::first-letter` 是伪元素选择器，属于 elements。所以 3 者相加最后的权重得分为 3 |
| `li > a[href*="en-US"] > .inline-warning` | 0           | 2       | 2        | 0-2-2             | `li`  是元素选择器，`a` 是元素选择器，2 者都属于 elements；  `[href*="eb-US"]` 是属性选择器，`.inline-warning` 是类选择器，均属于 classes。4 者相加最后的权重得分为 22 |
| `#identifier`                             | 1           | 0       | 0        | 1-0-0             | `#identifier` 是 id 选择器，属于 identifer。所以最后的权重得分为 100 |
| `button:not(#mainBtn, .cta`)              | 1           | 1       | 1        | 1-1-1             | `button` 是元素选择器，属于 elements；`#mainBtn` 是 id 选择器，属于 identifer；`.cta` 是类选择器，属于类选择器。所以最后的权重得分为 111 |

伪代码逻辑如下

```
if ruleset1.specificity < ruleset2.specificity then
       	  return ruleset2
elif ruleset1.specificity > ruleset2.specificity then
     	return ruleset1
else
	  return compare(rulset1.order, ruleset2.order)
```

## Inline style

不管是什么选择器，最后的权重值是多少，Inline style (通过 `style` 属性声明的 )的优先级比这些选择器都高

例如

```
#id {
    color: green;
}
---
<p id="id" style="color: red">this is a paragraph</p>
```

颜色会使用红色

伪代码如下

```
if rulset.property in element.inline_style then
	rulset.property = element.inline_style.property
else
	return compare(rulset1, rulset2)
```

## !important

在 CSS 中还有一个特殊的用法 -- `!important`，比 inline style 优先级高，是 CSS 中优先级最高的

例如，针对同一个元素使用 class selector 和 id selector 并声明使用不同的颜色

```
.class {
    color: red !important;
}
#id {
    color: green;
}
---
<p class="class" id="id">this is a paragraph</p>
```

按照逻辑，id selector 优先级比 class selector 高，应该会渲染成绿色，实际显示为红色。因为使用了 `!important` 即使优先级低，也会强制使用声明了 `!important` 对应的 property 的 value

如果我们要想 override 声明了 `!important` 的 properties，我们可以使用 2 种方法

1. 同 specificity 但是在源码中的声明的位置靠后，例如

   ```
   .class {
       color: red !important;
   }
   .class {
       color: green !important;
   }
   ```

   会变成绿色

2. 使用 sepcificity 优先级高的 selector

   ```
   .class {
       color: red !important;
   }
   #id {
       color: green !important;
   }
   ```

   会变成绿色

伪代码如下

```
if ruleset1.property has important and rulset2.property has important then
    if rulset1.specificity != rulset2.specificity then
        return compare(rulset1.specificity, rulset2.specificity)
    else
        returrn compare(ruleset1.order, rulset2.order)
elif rulset1.property has important and ruleset2.property has not importnat then
	ruleset1
elif rulset1.property has not important and rulset2.property has important then
	ruleset2
```

虽然 `!important` 优先级最高，但是会导致 CSS debug 复杂，应该尽量避免使用该用法

## Summarize

用伪代码的逻辑总结一下

```
if ruleset1.property has important or ruleset2.property has important then
    if ruleset1.property has important and rulset2.property has important then
        if rulset1.specificity != rulset2.specificity then
            return compare(rulset1.specificity, rulset2.specificity)
        else
            returrn compare(ruleset1.order, rulset2.order)
    elif rulset1.property has important and ruleset2.property has not importnat then
        ruleset1
    elif rulset1.property has not important and rulset2.property has important then
        ruleset2
elif rulset.property in element.inline_style then
	rulset.property = element.inline_style.property
elif rulset1.specificity != rulset2.specificity then
	return compare(rulset1.specificity, rulset2.specificity)
else
	returrn compare(ruleset1.order, rulset2.order)
```

**references**

1. https://developer.mozilla.org/en-US/docs/Learn/CSS/First_steps/How_CSS_is_structured
2. https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Cascade_and_inheritance