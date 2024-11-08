# HTML - HEAD

head 部分(`<head>...</head>`) 是 HTML 中的比较特殊的部分，不会被直接显示在 browser。主要被用来

1. link CSS or Javascript
2. link custom icons
3. set the title of the page
4. set primary language of the document
5. set metadata

```
<head>
  <meta charset="utf-8" />
  <link rel="stylesheet" href="my-css-file.css" />
  <script src="my-js-file.js" defer></script>
  <title>My test page</title>
</head>
```

## Link CSS or Javascript

在 HTML 中可以将 Css 和 Javascript 嵌入

```
<link rel="stylesheet" href="my-css-file.css" />
```

`rel="stylesheet"`, which indicates that it is the document's stylesheet, and `href`, which contains the path to the stylesheet file

```
<script src="my-js-file.js" defer></script>
```

> `<script>` element 不能是一个 empty element, 必须含有 closing tag

a `src` attribute containing the path to the JavaScript you want to load, and `defer`, ==which basically instructs the browser to load the JavaScript after the page has finished parsing the HTML==

## Set the title of page

```
<title>What's in the head? Metadata in HTML</title>
```

## Link Custom icons

A favicon can be added to your page by:

1. Saving it in the same directory as the site's index page, saved in `.ico` format (most browsers will support favicons in more common formats like `.gif` or `.png`, but using the ICO format will ensure it works as far back as Internet Explorer 6.)

2. Adding the following line into your HTML's  `<head>` block to reference it:    

   ```
   <link rel="icon" href="favicon.ico" type="image/x-icon" />
   ```

## Set primary language of the document

设置首选语言

```
<html lang="en-US">
  …
</html>

```

## Set Metadata

在计算机语言中的 metadata 指的是描述数据的数据，在 HTML 中通过 `<meta>` 标签来引入

### charset

用于标识文档使用的字符编码集

```
<meta charset="utf-8" />
```

### author description

用于标识 author

```
<meta name="author" content="Chris Mills" />
```

description 部分有助于 SEO 检索

```
<meta
  name="description"
  content="The MDN Web Docs site
  provides information about Open Web technologies
  including HTML, CSS, and APIs for both Web sites and
  progressive web apps." />
```

在 SEO 的 Digest 部分会显示 description 部分

![A Yahoo search result for "Mozilla Developer Network"](https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/The_head_metadata_in_HTML/mdn-search-result.png)

sub titles 部分是由  [Google's webmaster tools](https://search.google.com/search-console/about?hl=en)  配置的

**referneces**

1. https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/The_head_metadata_in_HTML	