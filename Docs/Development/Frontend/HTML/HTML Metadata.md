# HTML Metadata

ref

https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/The_head_metadata_in_HTML

the head of an HTML document is the part that is not displayed in the web browser when the page is loaded

head 部分用于描述 HTML doc metadata

## Author description

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

## Custom icons

A favicon can be added to your page by:

1. Saving it in the same directory as the site's index page, saved in `.ico` format (most browsers will support favicons in more common formats like `.gif` or `.png`, but using the ICO format will ensure it works as far back as Internet Explorer 6.)

2. Adding the following line into your HTML's 

   `<head>`

    block to reference it:    

   ```
   <link rel="icon" href="favicon.ico" type="image/x-icon" />
   ```

## Applying Css and Javascript

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