# HTML Process Validation

ref

https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Debugging_HTML

在 HTML 中主要有 2 种错误

- syntax errors

- logic errors

但是 HTML 并不会因为 syntax error 而终止运行，相反的 browser 会将错误的内容解析出来

之所以 HTML is parsed permissvely 是因为 HTML 是不断更新的，但是 browser 可能并没有适配相关的 syntax 这样就会导致网站不能被正常浏览

## HTML validation

由于 HTML 的特性, 导致 debugging 会比较困难。可以通过 w3c 提供的在线服务来校验 HTML code

https://validator.w3.org/