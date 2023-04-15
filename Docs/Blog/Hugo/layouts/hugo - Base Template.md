reference

1. [https://gohugo.io/templates/base/](https://gohugo.io/templates/base/)

hugo 除了支持 list template 和 single template 外，还支持一种特殊的 template - base template，为所有的 template 提供 template。可以将其抽象成接口
例如定义了一个 base template `_default/baseof.html` 内容如下
```
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>{{ block "title" . }}
      <!-- Blocks may include default content. -->
      {{ .Site.Title }}
    {{ end }}</title>
  </head>
  <body>
  	<h1>part before block main</h1>
    {{ block "main" . }}
    	<h1>part in block main</h1>
    {{ end }}
  	<h1>part after block main</h1>
  	<h1>part before block footer</h1>
    {{ block "footer" . }}
		<h1>part in block footer</h1>    
		{{ end }}
		<h1>part after block footer</h1>
  </body>
</html>

```
那么其他的 template 就可以通过声明 `{{define "main"}}...{{end}}` 的方式来 "继承" base template，假设 `layouts/index.html` 中内容如下
```
{{ define "main" }}
  <h1>Posts</h1>
  {{ range .Pages }}
    <article>
      <h2>{{ .Title }}</h2>
      {{ .Content }}
    </article>
  {{ end }}
{{ end }}
```
可以想象成如下内容，但是需要知道的一点是这种语法是错误的，在 hugo 中只允许将模板内容定义在 `{{define}}...{{end}}` 内
```
<h1>part before block main</h1>
{{ define "main" }}
  <h1>Posts</h1>
  {{ range .Pages }}
    <article>
      <h2>{{ .Title }}</h2>
      {{ .Content }}
    </article>
  {{ end }}
{{ end }}
<h1>part after block main</h1>
<h1>part before block footer</h1>
{{ block "footer" . }}
<h1>part in block footer</h1>    
{{ end }}
<h1>part after block footer</h1>
```
渲染内容如下
```
  <body>
  	<h1>part before block main</h1>
  	<h1>Posts</h1>
  	<h1>part after block main</h1>
  	<h1>part before block footer</h1>
		<h1>part in block footer</h1>    
		<h1>part after block footer</h1>
  </body>
```
这里并没有输出 `<h1>part in block main</h1>`，因为 `_default/baseof.html` 中的`{{block "main"}}...{{end}}` 中的内容，会被 `{{define "main"}}...{{end}}` 中的内容 override
