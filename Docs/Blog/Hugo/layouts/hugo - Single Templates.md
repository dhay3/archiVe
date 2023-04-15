references

1. [https://gohugo.io/templates/single-page-templates/](https://gohugo.io/templates/single-page-templates/)

在 hugo 中 `content` 目录下的文章在渲染时，会使用 single template 作为模板
![](https://intranetproxy.alipay.com/skylark/lark/0/2023/svg/23156369/1681205263432-baf10060-d014-4cd6-8d2f-c56f9175ce17.svg#clientId=u5c634856-46ca-4&from=paste&id=u1d502317&originHeight=446&originWidth=1730&originalType=url&ratio=1&rotation=0&showTitle=false&status=done&style=none&taskId=ub7b1dc32-d874-4846-8fd3-5b71e77d73f&title=)
例如上图中的，绿色部分就会使用 single template 来渲染对应的内容。同时也是 base template 的补充
## Example
single template 是可以通过 page variables 和 site variables 获取 `content` 目录下的文章对应的内容的
例如 `layouts/posts/single.html` 中的内容如下
```
{{ define "main" }}

<section id="main">
  <h1 id="title">{{ .Title }}</h1>
  <div>
    <article id="content">
      {{ .Content }}
    </article>
  </div>
</section>
<aside id="meta">
  <div>
  <section>
    <h4 id="date"> {{ .Date.Format "Mon Jan 2, 2006" }} </h4>
    <h5 id="wordcount"> {{ .WordCount }} Words </h5>
  </section>
    {{ with .GetTerms "topics" }}
      <ul id="topics">
        {{ range . }}
          <li><a href="{{ .RelPermalink }}">{{ .LinkTitle }}</a></li>
        {{ end }}
      </ul>
    {{ end }}
    {{ with .GetTerms "tags" }}
      <ul id="tags">
        {{ range . }}
          <li><a href="{{ .RelPermalink }}">{{ .LinkTitle }}</a></li>
        {{ end }}
      </ul>
    {{ end }}
  </div>
  <div>
    {{ with .PrevInSection }}
      <a class="previous" href="{{ .Permalink }}"> {{ .Title }}</a>
    {{ end }}
    {{ with .NextInSection }}
      <a class="next" href="{{ .Permalink }}"> {{ .Title }}</a>
    {{ end }}
  </div>
</aside>
{{ end }}
```
需要注意的一点是 template 需要在 `{{define}}...{{end}}` 内，否则不能渲染
## Lookup order
例如 `content/posts` 中的文章就会按照如下的顺序选择 template

1. layouts/posts/single.html.html
2. layouts/posts/single.html
3. layouts/_default/single.html.html
4. layouts/_default/single.html
