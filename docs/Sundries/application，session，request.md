# application，session，request

## application

作用范围：

从web服务器开始执行，到服务器关闭。所有用户的共享

## session

作用范围：

1、设置会话的时间到期。

2、关闭浏览器（关闭会话窗口不会造成session失效）。区别于Cookie，Cookie如果没有设置过期时间关闭浏览器后就会失效；如果设置了过期时间，Cookie只在达到过最大存活时间才会失效。

## request

作用范围：

单次请求