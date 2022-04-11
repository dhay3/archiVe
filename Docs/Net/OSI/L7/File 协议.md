# File 协议

参考：https://en.wikipedia.org/wiki/File_URI_scheme

file协议通常用于列出本机文件目录或文件内容(==不适用于远程主机==)

pattern：`file://host/path`

如果host省略，默认使用localhost，==但是slash不可以省略==

- Linux访问

  ```
  file://localhost/etc/fstab
  file:///etc/fstab
  ```

- windows访问

  ```
  file:///C:/
  file://localhost/c$/ #如果使用美元符需要带上host
  ```

  

