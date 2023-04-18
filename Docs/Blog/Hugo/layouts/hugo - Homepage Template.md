# hugo - Homepage Template

homepage 作为站点特殊的页面，hugo 也支持使用单独的 template
## Example
例如 `layouts/index.html` 中的内容如下
```
{{ define "main" }}
	<h1>this is the homepage.</h1>
{{ end }}
```
当访问 `example.org` 就会显示对应 `layouts/index.html` 渲染后的内容
## Lookup oder
对应 homepage，会按照如下的顺序选择 template

1. layouts/index.html.html
2. layouts/home.html.html
3. layouts/list.html.html
4. layouts/index.html
5. layouts/home.html
6. layouts/list.html
7. layouts/_default/index.html.html
8. layouts/_default/home.html.html
9. layouts/_default/list.html.html
10. layouts/_default/index.html
11. layouts/_default/home.html
12. layouts/_default/list.html
## _index.md
homepage 和 list pages 类似，也可以从 `content/_index.md` 中获取 front matter，并应用在 homepage template 中
例如 `content/_index.md` 中的内容如下
```
---
title: "how to build a site by hugo"
---
1. download hugo
2. run hugo server
```
那么在 template 中就可以通过 `{{.title}}` 和 `{{.Content}}` 获取 `_index.md` 中对应的内容
```
{{ define "main" }}
  <main aria-role="main">
    <header class="homepage-header">
      <h1>{{ .Title }}</h1>
      {{ with .Params.subtitle }}
      <span class="subtitle">{{ . }}</span>
      {{ end }}
    </header>
    <div class="homepage-content">
      <!-- Note that the content for index.html, as a sort of list page, will pull from content/_index.md -->
      {{ .Content }}
    </div>
    <div>
      {{ range first 10 .Site.RegularPages }}
          {{ .Render "summary" }}
      {{ end }}
    </div>
  </main>
{{ end }}
```



**references**

1. [https://gohugo.io/templates/homepage/](https://gohugo.io/templates/homepage/)

