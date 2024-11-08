# hugo - List Templates

list templates 用于渲染 `content` 下目录对应的页面，或者`content`目录本身( 对应的 homepage )，即包含站点文章的目录
![](https://intranetproxy.alipay.com/skylark/lark/0/2023/svg/23156369/1681205263432-baf10060-d014-4cd6-8d2f-c56f9175ce17.svg#clientId=u5c634856-46ca-4&from=paste&id=u1d502317&originHeight=446&originWidth=1730&originalType=url&ratio=1&rotation=0&showTitle=false&status=done&style=none&taskId=ub7b1dc32-d874-4846-8fd3-5b71e77d73f&title=)
例如上图中蓝色和紫色的部分，都会使用 list templates 来渲染对应的内容。同时也是 base template 的补充
有几个位置用于存储 list templates

1. `layouts/_default/list.html`
2. `theme/<THEME>/layouts/_default/list.html` 需要配置 theme
## Example
例如 `layouts/section/posts.html` 内容如下
```
{{ partial "header.html" . }}
{{ partial "subheader.html" . }}
<main>
  <div>
   <h1>{{ .Title }}</h1>
        <ul>
        <!-- Renders the li.html content view for each content/posts/*.md -->
            {{ range .Pages }}
                {{ .Render "li" }}
            {{ end }}
        </ul>
  </div>
</main>
{{ partial "footer.html" . }}
```
## Lookup order
对应 `content`目录下目录
例如 `content/posts` section，会按照如下的顺序选择 template

1. layouts/posts/posts.html.html
2. layouts/posts/section.html.html
3. layouts/posts/list.html.html
4. layouts/posts/posts.html
5. layouts/posts/section.html
6. layouts/posts/list.html
7. layouts/section/posts.html.html
8. layouts/section/section.html.html
9. layouts/section/list.html.html
10. layouts/section/posts.html
11. layouts/section/section.html
12. layouts/section/list.html
13. layouts/_default/posts.html.html
14. layouts/_default/section.html.html
15. layouts/_default/list.html.html
16. layouts/_default/posts.html
17. layouts/_default/section.html
18. layouts/_default/list.html
## _index.md
list templates 会关联 `_index.md` 中的 front matter，这样我们就可以在 list templates 中获取到对应 front matter 中元数据的值。同样的如果需要获取自定义的值，需要使用 `.Params`
例如文件目录如下
```
.
...
├── content
|   ├── posts
|   |   ├── _index.md
|   |   ├── post-01.md
|   |   └── post-02.md
|   └── quote
|   |   ├── quote-01.md
|   |   └── quote-02.md
...
```
`content/posts/_index.md` 中的内容如下
```
---
title: My Go Journey
date: 2017-03-23
publishdate: 2017-03-24
---

I decided to start learning Go in March 2017.

Follow my journey through this new blog.
```
那么就可以在 list templates 中获取到 `_index.md` 的内容，例如 `layouts/_default/list.html` 
```
{{ define "main" }}
<main>
  <article>
    <header>
      <h1>{{ .Title} }</h1>
    </header>
    <!-- "{{ .Content} }" pulls from the markdown content of the corresponding _index.md -->
    {{ .Content }}
  </article>
  <ul>
    <!-- Ranges through content/posts/*.md -->
    {{ range .Pages }}
      <li>
        <a href="{{ .Permalink }}">{{ .Date.Format "2006-01-02" }} | {{ .Title }}</a>
      </li>
    {{ end }}
  </ul>
</main>
{{ end }}
```
那么就会渲染对应的内容如下，需要注意的是 template 需要在 `{{define "main"}}...{{end}}`内，否则不能渲染
```
<!--top of your baseof code-->
<main>
    <article>
        <header>
            <h1>My Go Journey</h1>
        </header>
        <p>I decided to start learning Go in March 2017.</p>
        <p>Follow my journey through this new blog.</p>
    </article>
    <ul>
        <li><a href="/posts/post-01/">Post 1</a></li>
        <li><a href="/posts/post-02/">Post 2</a></li>
    </ul>
</main>
<!--bottom of your baseof-->
```
上面的例子还可以得出一点的是，如果通过 go tamplate 获取到的 front matter 中的元数据，没有在 html 标签内，自动会渲染成 `<p>` 和 HTML 中的规则一样
如果没有 `_index.md`，那么在 list templates 中的 `{{.Content}}` 的值就会空，最后也不会渲染对应的内容。例如 `content/quote/` 渲染后对应的 `/quote/index.html`内容如下
```
<!--baseof-->
<main>
    <article>
        <header>
        <!-- Hugo assumes that .Title is the name of the section since there is no _index.md content file from which to pull a "title:" field -->
            <h1>Quotes</h1>
        </header>
    </article>
    <ul>
        <li><a href="https://example.com/quote/quotes-01/">Quote 1</a></li>
        <li><a href="https://example.com/quote/quotes-02/">Quote 2</a></li>
    </ul>
</main>
<!--baseof-->
```
## Order content
站点中的文章，可以通过 front matter 指定的字段按照指定的顺序来显示。顺序为 Weight>Date>LinkTitle>FilePath。这里只介绍默认的顺序，具体参考官网
即当 front matter 中有 Weight 字段，优先采用 Weight，反之使用 Date。以此类推
```
<ul>
    {{ range .Pages }}
        <li>
            <h1><a href="{{ .Permalink }}">{{ .Title }}</a></h1>
            <time>{{ .Date.Format "Mon, Jan 2, 2006" }}</time>
        </li>
    {{ end }}
</ul>
```
上述是默认按照 weight 字段来排序
假设，现在有一个 `content/post1.md`
```
---
title: "My Important post"
date: 2020-09-15T11:30:03+00:00
weight: 1
---
```
还有一个 `content/post2.md`
```
---
title: "My 2nd Important post"
date: 2020-09-15T11:30:03+00:00
weight: 2
---
```
那么对应的 list pages 中就会先显示 post1，再显示 post2
也可以理解成 post1 置顶

**references**

1. [https://gohugo.io/templates/lists/](https://gohugo.io/templates/lists/)
2. [https://github.com/adityatelange/hugo-PaperMod/wiki/FAQs](https://github.com/adityatelange/hugo-PaperMod/wiki/FAQs)
