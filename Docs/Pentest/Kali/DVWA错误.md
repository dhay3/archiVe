# DVWA错误

搭建环境Kali



## The PHP function allow_url_include is not enabled.

```shell
locate php.ini
...
vim /etc/php/7.4/php.ini
...
allow_url_include=on
systemctl restart apache2
```

## Incorrect folder permissions: /var/www/html/dvwa/hackable/uploads/ Folder is

```shell
chmode 777 /var/www/html/dvwa/hackable/uploads
```

