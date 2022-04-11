# 查看主机公网IP

1. ifconfig.me，海外节点

   https://ifconfig.me/

   ```
   C:\Users\82341>curl ifconfig.me/all
   ip_addr: 115.233.222.34
   remote_host: unavailable
   user_agent: curl/7.55.1
   port: 53146
   language:
   referer:
   connection:
   keep_alive:
   method: GET
   encoding:
   mime: */*
   charset:
   via: 1.1 google
   forwarded: 115.233.222.34, 216.239.34.21
   ```

   可以直接`curl ifconig.me`直接获取ip，无需使用文本工具截取

2. cip.cc

   http://www.cip.cc/

   ```
   C:\Users\82341>curl cip.cc
   IP      : 61.175.192.50
   地址    : 中国  浙江  杭州
   运营商  : 电信
   
   数据二  : 浙江省杭州市 | 电信
   
   数据三  : 中国浙江杭州 | 电信
   
   URL     : http://www.cip.cc/61.175.192.50
   ```

   