# OID

参考：

https://en.wikipedia.org/wiki/Object_identifier

https://stackoverflow.com/questions/14623335/how-to-specify-the-syntax-for-values-of-private-oids-while-configuring-in-openss

OID(object Identifier)，对象的唯一标识符。从根节点到祖先节点，由整数组成`.`号分隔。例如：

`1.3.6.1.4.1.343`。通常被用于X.509认证

- 1后面的子节点是由ISO(international organization for standardization)分配
- 1.3.6 后面的子节点由US Department of Defense分配
- 1.3.6.1.4.1后面的子节点由IANA分配
- 1.3.6.1.4.1.343后面的子节点由Intel Corporation分配