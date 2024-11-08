# HTML structure of website

ref

https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Document_and_website_structure

HTML websites 通常按照一定的架构编写，主要包含如下几个部分

- header

  Usually a big strip across the top with a big heading, logo, and perhaps a tagline. This usually stays the same from one webpage to another.

- navigation bar

  Links to the site's main sections; usually represented by menu buttons,  links, or tabs. Like the header, this content usually remains consistent from one webpage to another — having inconsistent navigation on your  website will just lead to confused, frustrated users. Many web designers consider the navigation bar to be part of the header rather than an  individual component, but that's not a requirement; in fact, some also  argue that having the two separate is better for [accessibility](https://developer.mozilla.org/en-US/docs/Learn/Accessibility), as screen readers can read the two features better if they are separate.

- main content

  A big area in the center that contains most of the unique content of a  given webpage, for example, the video you want to watch, or the main  story you're reading, or the map you want to view, or the news  headlines, etc. This is the one part of the website that definitely will vary from page to page!

- sidebar

  Some peripheral info, links, quotes, ads, etc. Usually, this is  contextual to what is contained in the main content (for example on a  news article page, the sidebar might contain the author's bio, or links  to related articles) but there are also cases where you'll find some  recurring elements like a secondary navigation system.

- footer

  A strip across the bottom of the page that generally contains fine  print, copyright notices, or contact info. It's a place to put common  information (like the header) but usually, that information is not  critical or secondary to the website itself. The footer is also  sometimes used for [SEO](https://developer.mozilla.org/en-US/docs/Glossary/SEO) purposes, by providing links for quick access to popular content.

## Non-semantic wrappers

在 HTML 中有一些 element 是没有语义的例如：`<div>`, `<span>`

if you might want to just group a set of elements together to affect them all as a sigle entity with some CSS or Javascript, `<div>` or `<span>` should be chosen

`<span>` is an inline non-semantic element

`<div>` is a block level non-semantic element