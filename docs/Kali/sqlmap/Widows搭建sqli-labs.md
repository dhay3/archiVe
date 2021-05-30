# Widows搭建sqli-labs

> 从github上下载的sqli-labs, 在本地安装可能由于php版本的不一样, 会无法运行，所以这里我们采用phpstudy集成环境。

具体使用参考：

https://www.xp.cn/wenda/388.html

将下载好的sqli-labs解压放入`D:\phpStudy\PHPTutorial\WWW`。修改sqli-labs配置文件

```php
<?php
//give your mysql connection username n password
$dbuser ='root';
$dbpass ='root';
$dbname ="security";
$host = 'localhost';
$dbname1 = "challenges";
?>
```

将用户和密码都设置为root，具体参考：https://www.xp.cn/a.php/186.html

访问http://localhost/sqli/生成数据库即可
