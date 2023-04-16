# CSS - Applying CSS

可以通过 3 种方式来使用 CSS

1. external stylesheet
2. internal stylesheet
3. inline styles

## External stylesheet

external stylesheet 指的是含有 CSS 层叠样式，后缀为 `.css ` 的文件。我们可以通过如下方式导入

```
// styles.css
h1 {
  color: blue;
  background-color: yellow;
  border: 1px solid black;
}

p {
  color: red;
}
---
<link rel="stylesheet" href="styles.css" />
```

- `rel`

  声明是导入文件中的内容是 CSS stylesheet

- `href`

  `.css` 文件的路径，支持绝对或者是相对路径

## Internal stylesheet

通过  `<style></style>` 标签，将 CSS 层叠样式声明在同一个 HTML 文件中

```
<!DOCTYPE html>
<html lang="en-GB">
  <head>
    <meta charset="utf-8" />
    <title>My CSS experiment</title>
    <style>
      h1 {
        color: blue;
        background-color: yellow;
        border: 1px solid black;
      }

      p {
        color: red;
      }
    </style>
  </head>
  <body>
    <h1>Hello World!</h1>
    <p>This is my first CSS example</p>
  </body>
</html>
```

## Inline style

通过标签的`style` 属性，将 CSS 层叠样式直接声明在 HTML 元素上

```
<!DOCTYPE html>
<html lang="en-GB">
  <head>
    <meta charset="utf-8" />
    <title>My CSS experiment</title>
  </head>
  <body>
    <h1 style="color: blue;background-color: yellow;border: 1px solid black;">
      Hello World!
    </h1>
    <p style="color:red;">This is my first CSS example</p>
  </body>
</html>
```



**references**

1. https://developer.mozilla.org/en-US/docs/Learn/CSS/First_steps/How_CSS_is_structured