references

1. [https://gohugo.io/templates/shortcode-templates/](https://gohugo.io/templates/shortcode-templates/)

shortcode 是可以重复利用的 snippets ，可以理解是迷你 template，存储在 `layouts/shortcodes` 目录下
## Example
假设有一个 `layouts/shortcodes/myshortcode.html`
```
this is a shortcode test
```
shortcode 不需要在 `{{define ""}}...{{end}}` 内，因为继承 base template 无意义，那么我们就可以在文章中或者 template 中通过如下的方式来调用 shortcode
```
{{<myshortcode>}}
```
## Lookup order
shortcode 会按照如下顺序搜索

1. `/layouts/shortcodes/<SHORTCODE>.html`
2. `/themes/<THEME>/layouts/shortcodes/<SHORTCODE>.html`
## Parameters
### .Get
shortcode 和 template 一样，有一种 `{{<shortcodeName params>}}` 的格式，将 params 应用到 shortcode。即可以将 shortcode 抽想成函数，params 为 shortcode 对应的函数的入参
例如有一个 `content/post1.md` 中包含如下内容
```
{{<myshortcode color="blue">}}
```
`myshortcode` 就是一个函数，`color="blue"` 就是类似 python 中的 kwargs
`layouts/shortcodes/myshortcode.html` 中的内容如下
```
<p color="{{.Get `color`}}">this is a shortcode test</p>
```
那么 `content/post1.md` 对应部分实际就会渲染成
```
<p color="blue">this is a shortcode test</p>
```
这里的 `{{.Get `color`}}` 是获取到 shortcode 传进来的参数的
### .Inner
shortcode 还有一种特殊的格式，假设有一个 `content/post1.md` 中包含如下内容
```
{{<myshortcode>}}
	this is the content of post1
{{</ myshortcode>}}
```
同时 `layouts/shortcodes/myshortcode.html` 中的内容如下
```
shortcode wrapper {{.inner}} shortcode wrapper 
```
那么最后渲染的内容如下
```
shortcode wrapper this is the content of post1 shortcode wrapper 
```
假如 `{{<myshortcode>}}...{{</myshortcode>}}` 中有类似 Markdown 的语法，是不会被渲染，需要通过 `{{% myshortcode %}}...{{% /myshortcode %}}` 方式
