# MySQL 万能密码

> 使用mybatis的预先编辑即可很好的解决sql注入问题

-  admin'#

```
SELECT * FROM sys_user WHERE name='admin'# AND password='123'
```

- admin' AND 1=1 -- 

```
SELECT * FROM sys_user WHERE name='admin' AND 1=1 -- AND password =123
```

