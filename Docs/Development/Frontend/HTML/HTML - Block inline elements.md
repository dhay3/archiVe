# HTML - Block inline elements

在 HTML 中的所有元素，都可以分为两类

1. Block elements
2. Inline elements，中文也被分为内联元素

## Block elements

中文也被称为块级元素，有如下几点特性

1. Any content that follows a block elements also appears on a new line (即块级元素后默认有换行符)
2. A block element wouldn’t be nested inside an inline element (块级元素不能定义在内联元素内)

例如 `div`, `p` 等都是块级元素

## Inline elements

中文也被分为内联元素，有如下几点特性

1. Inline elements are contained within block-level elements (内联元素被定义在块级元素内)
2. An inlne element will not cause a new line to appear in the document (即内联元素不会换行)

例如 `a`, `span` 都是内联元素

**references**

1. https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Getting_started#block_versus_inline_elements