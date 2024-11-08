# Apache 安装SSL证书

参考：

apache2：https://help.aliyun.com/document_detail/102450.html?spm=a2c4g.11186623.2.18.4b3a3cabqG8Z6j#concept-cfn-yf2-kgb

apache：https://help.aliyun.com/document_detail/98727.html?spm=a2c4g.11186623.2.14.50604578Uy6nAz#concept-zsp-d1x-yfb

1. 购买SSL证书并下载，将文件放到服务器

   ```
   root in /etc/apache2/cyberpelican_space_cert λ ll
   .rw-r--r-- root root 1.6 KB Sun Jan  3 19:24:57 2021  5002550_cyberpelican.link.key
   .rw-r--r-- root root 1.6 KB Sun Jan  3 19:24:57 2021  5002550_cyberpelican.link_chain.crt
   .rw-r--r-- root root   2 KB Sun Jan  3 19:24:57 2021  5002550_cyberpelican.link_public.crt                    
   ```

- 证书文件：以`.crt`为后缀或文件类型。
- 证书链文件：以`chain.crt`为后缀或文件类型。
- 密钥文件：以`.key`为后缀或文件类型。

1. 修改配置文件

   ```
   
   ```

   