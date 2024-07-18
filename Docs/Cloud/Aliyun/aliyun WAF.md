# aliyun WAF

## 概述

web application firewall (WAF)

## Term

- 源站

  指提供服务的后端服务器

- 回源IP

  指WAF与源站建立网络连接的IP地址

  ```mermaid
  graph LR
  a(用户)-->|双向|b(DNS服务器)
  a -->|经过WAF 回源IP|c(web app) 
  ```

  

- Web应用

  指用户通过browser访问的应用程序

- 四层代理

  只分析请求报文中的目的地址和端口信息，结合服务器选择，直接访问请求转发到源站服务器

- 七层代理

  除了分析四层代理，还分析报文中应用层内容。根据报文中特定字段以及服务器选择规则，进行请求转发