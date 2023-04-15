references

1. [https://hugo-mini-course.netlify.app/sections/styling/custom/](https://hugo-mini-course.netlify.app/sections/styling/custom/)

## Static Directory
`static` 是 hugo 中特殊的一个目录，存储在改目录下的内容。可以通过 `/` 的方式来获取
例如，`static` 目录有下一个 duck.png 图片
那么就可以在文章中通过 `![pic](/duck.png)` 的方式来调用图片
同样的在 template 中，可以通过 `<img src="/duck.png">` 的方式来调用
## Example 
假设当前站点结构如下
```
my-blog
├── archetypes
|   └── default.md
├── content
|   └── posts
|       └── my-first-post.md
├── data
├── layouts
|   ├── _default
|   |   ├── summary.html
|   |   ├── single.html
|   |   ├── list.html
|   |   └── baseof.html
|   ├── partials
|   |   ├── footer.html
|   |   └── navbar.html
|   ├── 404.html
|   └── index.html
├── static
|   └── styles.css
├── themes
└── config.toml
```
`static/styles.css` 中内容如下
```
body {
  background: blue;
}
```
`layouts/_default/baseof.html` 中内容如下，`<link rel="stylesheet" href="/styles.css">` 会自动引入 `static` 下的 `styles.css`
```
<!DOCTYPE html>
<html>

<head>
  <meta charset="utf-8">
  <link rel="stylesheet" href="/styles.css">
  <title>{{ block "title" . }}
    {{ .Site.Title }}
    {{ end }}</title>
</head>

<body>
  {{ partial "navbar.html" . }}
  {{ block "main" . }}
  <!-- The part of the page that begins to differ between templates -->
  {{ end }}
  {{ partial "footer.html" . }}
</body>

</html>
```
当然也支持从互联网引入 CSS
```
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.0/css/bulma.min.css">
```
如果在 base template 中引入 CSS，那么也会被其他 template 继承，在其他 template 不需要显示的引入 CSS一样能使用对应的 CSS
