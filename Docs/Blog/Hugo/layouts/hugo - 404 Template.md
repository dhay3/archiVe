refernece

1. [https://gohugo.io/templates/404/](https://gohugo.io/templates/404/)

hugo 也支持为 404 页面提供 template，模板内容需要存储在 `layouts/404.html`
例如 `layouts/404.html`
```
{{ define "main" }}
  <main id="main">
    <div>
      <h1 id="title"><a href="{{ "" | relURL }}">Go Home</a></h1>
    </div>
  </main>
{{ end }}
```
