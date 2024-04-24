# Google Searches

## 0x01 Overview

Google SEO 提供了 2 种方式过滤以获取详细的信息

1. Advanced search
2. Operators search

## 0x02 Advanced search

高级搜索顾名思义，这里不过多介绍

![](https://github.com/dhay3/picx-images-hosting/raw/master/20240410/2024-04-10_08-54.4jnvw2s308.webp)

## 0x03 Operators Search

在不使用 operators 的情况下，如果你搜索 meme cat 会得到很多条结果，你可能并不能获取到自己想要的结果。而在使用 operators 的情况下可以，你可以更准确的过滤这些信息，例如

```
image meme cat before:2015 intitle:grumpy
```

Operators 可以是如下 3 种

1. properties operators

   格式如下

   ```
   properties-operators:<keyword>
   ```

2. logical operators

3. special words operators

   格式如下

   ```
   special-words-operators <keyword>
   ```

在 Google operators search 还需要注意如下几点（并不一定适用其他 SEO，例如duckduckgo 中 operator 和 your search 之前有空格和没有空格结果相同）

1. If punctuation(标点符号) is not part of the search operator, Google will usually ignore it. 

   ==但是除 `site:` 外==

   例如

   `intitle:mom s spaghetti` 和 `intitle:mom's spaghetti` 结果相同，因为符号会被忽略

2. Do not put spaces between the operator and your search.

   例如

   `site:nytimes.com` 可以搜索到 nytime.com 站点内的内容

   `site: nytimes.com` 只能搜索到 nytimes.com 相关的内容

3. Put the keyword before the operators if not sure where to use or where to search

### Properties operators

- `filetype:<extension>`

  将搜索结果限制为特定的文件类型，如：PDF、PPT 等。也可以使用“ext:”，作用是相同的。

  ```
  kali filetype:torrent
  ```

- `book:<keyword>`

  查询书籍

  ```
  kali book:hacker
  ```

- `site:<site>`

  限定搜索结果在指定的网站内

  ```
  site:google.com
  ```

  site 会自动扩展成子域名，无需使用 `*`

- `inurl:<keyword>`

  限定搜索结果的 URL 包含指定 keyword

  ```
  aur inurl:archlinux.org
  ```

  注意和 `site:` 做区别，同时标点符号会被忽略

- `intitle:<keyword>`

  限定搜索的网页标题(可以点击的链接)包含指定 keyword

  ```
  intitle:linux kernel
  ```

- `intext:<keyword>`

  限定搜索的网页内容(在网页标题下的文字)包含指定 keyword

  ```
  intext:admin|password
  ```

- `cache:<URL>`

  从 Google cache (一些被删除的页面并不会立即消失，而会被 google 缓存)中查询指定的网页

  ```
  cache:https://support.google.com/looker-studio/answer/10468382?hl=en
  ```

- `related:<keyword>`

  限定搜索结果和对应站点相关

  ```
  related:perplexity.com
  ```

- `inanchor:<keyword>`

  限定搜索结果和对应锚点相关，通常其他 properites operators 一起使用，例如 `site:`

  ```
  site:support.google.com inanchor:Operators
  ```

- `number1 .. number2`

  限定搜索结果在指定数字范围内，通常和年份一起使用，但是结果不一定直接和年份相关

  ```
  site:www.kernel.org  intext:release  2023..2024
  ```
  
- `before:<date>`

  `after:<date>`

  搜索特定时间段，比 `number1..number2` 更加准确，支持日期

  ```
  ipad after:2010 before:2015
  iphone after:2010-04-01 before:2015-04-01
  ```

- `@`

  限定搜索结果在特定社交媒体

  ```
  @twitter elon musk
  ```

- `#`

  限定搜索结果在相关话题内

  ```
  @twitter #ukraine war
  ```

### Logical operators

- `"keyword1 keyword2"`

  强制进行完全匹配

  ```
  "steve jobs"
  ```

- `keyword1 AND keyword2`

  搜索结果为 keyword1 和 keyword2

  ```
  arch AND manjaro
  ```

- `keyword1 OR keyword1`

  搜索结果为 keyword1 或者 keyword2，也可以使用符号 `|` 替代 `OR`

  ```
  mapple OR majito
  ```

- `-keyword`

  `-keyword` 表示排除搜索词 keyword，即返回的搜索结果排除与 keyword 有关的。

  ```
  apple -iphone
  ```

  也可以和 properties operators 一起使用

  ```
  apple -intext:iphone
  ```

- `*`

  通配符, 匹配任何关键字。

  ```
  steve *
  ```

- `()`

  通过 `()` 对 operators 进行分组，以控制搜索的执行方式。

  ```
  intext:(git OR svn) commit
  ```

### Special words operators

一些特定的关键字会触发 google 提供的内置功能

- `<calculator|calc> [...]`

  会触发计算器，例如 `calc 128x128/10`

- `translate [...]`

  会触发翻译，例如 `translate professor`

- `define [...]`

  会触发字典,例如 `define meme`

- `unit converter [...]`

  会触发单位转换器，例如 `unit converter  1024byte`

- `color piker`

  触发拾色器，可以做 RGB/HEX 等转换

- `currency converter [...]`

  会触发货币转换，例如 `currency converter 20RMB`

  还可以使用 `to/in` 关键字做特定的货币转换(推荐)，例如 `100RMB to USD`

- `weather`

  天气预报，例如 `weather hangzhou`

- `map`

  地图，例如 `map hangzhou`

- `movie`

  电影，例如 `movie "buster keaton"`

  关键字可以是影片，导演或者是演员

- `stocks`

  股票，例如 `stocks apple`

- `image`

  图片，例如 `image meme cat`

## 0x04 Combining Operators

Operators 之间可以互相组合

### Common Used

1. 域名遍历，但是不包括主机值为 `www` 的域名

   ```
   site:google.com -www
   ```

2. 在特定站点，查询特定内容

   ```
   site:github.com intext:aic8800
   ```

3. 在特定时间段，查询特定内容

   ```
   chian 64 after:1988 before:2000
   ```


4. 查找未使用 https 的页面

   ```
   site:google.com -inurl:https
   ```

### Hacking Used

可以在 exploit-db 找到各种用于 hacking 的 operators

https://www.exploit-db.com/google-hacking-database

例如

1. 找域名可爆破的后台

   ```
   site:github.com inurl:(login|admin|manage|member|admin_login|login_admin|system|login|user|main|cms)
   ```

2. 找 github 上泄漏的账号密码

   ```
   site:github.com inurl:src intext:password OR accessKeySecret OR accessKeyId -filetype:md
   ```

3. 找密码相关的 SQL

   ```
   *.sql password intitle:"index of"
   ```

4. 找渗透相关的文档

   ```
   pentest intitle:"index of" OR filetype:pdf
   ```

**references**

[^1]:https://support.google.com/websearch/answer/2466433?hl=en
[^2]:https://searchengineland.com/advanced-google-search-operators-388355
[^3]:https://support.google.com/websearch/answer/3284611?hl=en&ref_topic=3081620&sjid=2614773873898016221-AP
[^4]:https://blog.google/products/search/how-were-improving-search-results-when-you-use-quotes/
[^5]:https://web.archive.org/web/20221004180254/https://support.google.com/websearch/answer/2466433?hl=en
[^6]:https://web.archive.org/web/20161110213529/https://support.google.com/websearch/answer/2466433?hl=en
[^7]:https://moz.com/learn/seo/search-operators
[^8]:https://www.semrush.com/blog/google-search-operators/
