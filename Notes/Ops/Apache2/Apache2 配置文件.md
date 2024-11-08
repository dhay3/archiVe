# Apache2 配置文件

https://httpd.apache.org/docs/2.4/mod/quickreference.html

> 如果指令有多个值，换行
>
> `/etc/apach22/apache2.conf`为全局配置，如果匹配到了指令块，优先使用指令块中的配置

- ServerAdmin

  错误信息沟通的邮箱地址，该参数被弃用

- DocumentRoot

  文件存放的根目录（默认`/var/www/html`）

  ```
  DocumentRoot "/usr/web"
  ```

  当访问`http://my.example.com/index.html`时，读取`/usr/web/index.html`内容

- Listen

  指定Apache监听的端口，一般可以配置在全局，==必须要配置该项否则无法启动Apache==

- ServerName

  常用做于Virtrual Host指令块中辨别主机的标识符。

  如果使用了SSL证书需要指明scheme为https和端口443。

  如果未在全局正确设置ServerName，就会提示

  https://www.digitalocean.com/community/tutorials/apache-configuration-error-ah00558-could-not-reliably-determine-the-server-s-fully-qualified-domain-name
  
  ```
httpd: Could not reliably determine the server's fully qualified domain name, using rocinante.local for ServerName
  ```

- ServerRoot

  apache服务器安装的根路径

- TimeOut

  请求失败后，服务器等待的最长时间

- KeepAlive

  是否启用长连接

- MaxkeepAliveRequests 

  单次连接最大允许请求量

- KeepAliveTimeout

  断开连接等待的最长时间

- ErrorLog

  错误日志存放路径



