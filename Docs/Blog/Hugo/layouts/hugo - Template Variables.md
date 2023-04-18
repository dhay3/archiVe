# hugo - Template Variables

### Page Variables
hugo 中的每一个 template 都会传入一个 `Page` 对象，即 `Page` 对象就是 `{{.}}`，这样我们就可以通过 go template 的语法来获取对应的值

- `.Title`

对应文章名

- `.Content`

对应 front matter 下面的部分，如果是 list template 那么就是文章的内容

- `.WordCount`

content 字数

- `.Date`

对应 front matter 中的 `date`

- `.PublishDate`

对应 front matter 中的 `publishDate`

- `.Description`

the description of  the page

- `.Draft`

对应 front matter 中的 `draft`

- `.IsHome`

是否是 homepage

- `.IsPage`

是否是文章

- `.Pages`

如果

- `.Permalink`

文章对应的 permanent link

- `.RelPermalink`

文章对应的 relative permanent link

- `.Next`

当前文章的，后一篇文章

- `.Prev`

当前文章的，前一篇文章

## Site Variables

hugo 中的全局变量，一部分是 built-in 的，一部分是被定义在 site configuration (默认为 `config.toml` ) 中的变量

- `.Site.AllPages`

  array of all pages

- `.Site.BaseURL`

  对应 site configuration 中的 baseURL

- `.Site.BuildDrafts`

  对应 site configuratoin 中的 buildDrafts

- `.Site.Data`





**reference**

1. [https://gohugo.io/templates/introduction/](https://gohugo.io/templates/introduction/)
2. [https://gohugo.io/variables/page/](https://gohugo.io/variables/page/)
3. https://gohugo.io/variables/site/
