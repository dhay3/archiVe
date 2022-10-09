# CSS Cascade Specificity

 ref

https://developer.mozilla.org/en-US/docs/Learn/CSS/First_steps/How_CSS_is_structured

在 CSS 中会经常碰到 different selectors 匹配到 the same HTML element, 这时就需要 CSS specificity 来做判断，例如：

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

这时 CSS 会使用那条 rule ?

答案是由 *cascade* 和 *specificity* 决定

## Cascade

```
p {
  color: red;
}

p {
  color: blue;
}
```

如果 specificity 相同，the later one declaration 会被使用，这就是 cascade

## Specificity