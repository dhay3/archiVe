# CSS @rules



## Digest

和 JAVA 的注解一样，用于描述一种功能告诉 CSS 需要怎么操作

## Syntax

```
/* General structure */
@identifier (RULE);

/* Example: tells browser to use UTF-8 character set */
@charset "utf-8";
```

## Common @rule

- `@charset`

  defines the character set used by the style sheet

- `@import`

  include an external stylesheet

- `@font-face`

  describe an external font to be download

- `@keyframes`

  describe the a custom CSS animation

- `@media`

  a conditional rule that will apply its content if the device meets the criteria of the condition

## Nested

```
@identifier (RULE) {
}
```

当条件匹配时使用 declaration sets

例如：

```
body {
  background-color: pink;
}

@media (min-width: 30em) {
  body {
    background-color: blue;
  }
}
```

if the browser viewport is wider than 30em, blue will be  apply to the body’s background-color



**referneces**

[^1]:https://developer.mozilla.org/en-US/docs/Learn/CSS/First_steps/How_CSS_is_structured
[^2]: https://developer.mozilla.org/en-US/docs/Web/CSS/At-rule

