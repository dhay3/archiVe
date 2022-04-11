# robots.txt

参考：

https://zh.wikipedia.org/wiki/Robots.txt

https://developers.google.com/search/docs/advanced/robots/intro?hl=zh-cn

https://ahrefs.com/blog/zh/robots-txt/

## 概述

robots.txt用告诉搜索引擎改怎么抓取你的站点内容(对爬虫有一定的反扒效果)。大多数搜索引擎都会遵循规则，所以robot.txt不一定完全有效。==robots.txt文件应该位于站点的根目录==

```
Sitemap: [URL location of sitemap]

User-agent: [bot identifier]
[directive 1]
[directive 2]
[directive ...]

User-agent: [another bot identifier]
[directive 1]
[directive 2]
[directive ...]
```

1. 文件名必须为`robots.txt`
2. 一个站点只能有一个`robots.txt`文件
3. 必须放在根目录
4. robots.txt对子域名同样生效
5. `#`表示注释

robots.txt文件由一个或多个组组成，每个组由多条指令组成。

## group

- user-agent

  必不可少，即网页抓取工具软件的名称。可以使用`*`通配符。如果想要防止非浏览器工具的爬取，需要明确指定user-agent

- Disallow

  用户代理不能抓取目录或网页。如果要指定目录，需要标记以`/`结尾。支持通配符

- Allow

  与Disallow相反

- Sitemap

  网站的站点地图

## example

> 如果disallow与allow指令冲突，使用指令长度较长的，下面示例使用`Disallow:/blog/`
>
> 可以使用`$`正则匹配

```
User-agent: *
Disallow: /blog/
Allow: /blog

User-agent: Googlebot
Disallow: /*.gif$

User-agent: Googlebot
Allow: /
```







