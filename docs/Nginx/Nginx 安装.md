# Nginx 安装

https://nginx.org/en/download.html

https://segmentfault.com/a/1190000018109309

> 接下来的所有内容都以epel源为基础
>
> 如果nginx启动不了，查看80端口是否被apache2占用

Nginx有三个版本

1. Mainline version：最新版本，但是有些功能可能没有经过太多的测试
2. Stable version：稳定版本
3. Legacy version：历史版本

## 离线安装

> wget 默认将文件下载到当前目录

1. 下载压缩包

   ```
   [root@cyberpelican opt]#wget https://nginx.org/download/nginx-1.18.0.tar.gz
   [root@cyberpelican opt]#tar -zxvf nginx-1.18.0.tar.gz
   ```

2. 下载Nginx依赖

   ```
   yum install -y gc gcc gcc-c++
   yum install -y openssl openssl-devel  
   yum install -y pcre pcre-devel
   yum install -y zlib zlib-devel
   ```

3. 生成Makefile

   ```
   [root@cyberpelican nginx-1.18.0]#./configure
   [root@cyberpelican nginx-1.18.0]#ls
   auto     CHANGES.ru  configure  html     Makefile  objs    src
   CHANGES  conf        contrib    LICENSE  man       README
   
   ```

4. 执行Makefile，创建文件

   > 默认安装在`/usr/local/nginx`

   ```
   [root@cyberpelican nginx-1.18.0]# make && make install
   ```

5. ==进入 目录 /usr/local/nginx/sbin/nginx 启动服务==

   ```
   [root@cyberpelican nginx-1.18.0]# cd /usr/local/nginx/sbin
   [root@cyberpelican sbin]# ls
   nginx
   [root@cyberpelican sbin]# ./nginx
   [root@cyberpelican sbin]# netstat -lnpt | grep nginx
   tcp        0      0 0.0.0.0:80              0.0.0.0:*               LISTEN      15645/nginx: master 
   ```

## yum安装

<img src="..\..\imgs\_Nginx\Snipaste_2020-11-22_18-13-02.png"/>

这里使用CentOS7

1. 配置nginx yum repo

   > 这里可以使用aliyun 的 epel 源替代

   ```
   [root@cyberpelican ~]#touch /etc/yum.repos.d/nginx.repo 
   [root@cyberpelican yum.repos.d]#cat nginx.repo 
   [nginx-stable]
   name=nginx stable repo
   baseurl=http://nginx.org/packages/centos/$releasever/$basearch/
   gpgcheck=1
   enabled=1
   gpgkey=https://nginx.org/keys/nginx_signing.key
   module_hotfixes=true
   
   [nginx-mainline]
   name=nginx mainline repo
   baseurl=http://nginx.org/packages/mainline/centos/$releasever/$basearch/
   gpgcheck=1
   enabled=0
   gpgkey=https://nginx.org/keys/nginx_signing.key
   module_hotfixes=true
   ```

   安装时默认使用stable

2. 安装nginx

   ```
   yum install nginx
   ```

## 开启防火墙端口

```
[root@cyberpelican nginx]# firewall-cmd --permanent --add-service=http
success
[root@cyberpelican nginx]# firewall-cmd --permanent --add-service=https
success
[root@cyberpelican nginx]# firewall-cmd --reload
success
[root@cyberpelican nginx]# firewall-cmd --list-all
public (active)
  target: default
  icmp-block-inversion: no
  interfaces: ens33
  sources: 
  services: dhcpv6-client http https ssh
  ports: 
  protocols: 
  masquerade: no
  forward-ports: 
  source-ports: 
  icmp-blocks: 
  rich rules: 
```

## 校验

- epel校验

  注意如果是以epel安装的，出现的页面将是CentOS默认页面，具体查看`/usr/share/nginx/html/index.html`。我们可以通过修改index.html来校验，或是通过修改默认配置文件

  ```
   server {
          listen       80 default_server;
          listen       [::]:80 default_server;
          server_name  _;
          root         /usr/share/nginx/html;
  
          # Load configuration files for the default server block.
          include /etc/nginx/default.d/*.conf;
  
          location / {
            proxy_pass http://baidu.com;
          }
  
  ```

  这里让根路径路由到百度

- 离线安装校验

  开启nginx服务并访问`localhost`出现如下界面

<img src="..\..\imgs\_Nginx\Snipaste_2020-11-25_11-05-23.png"/>



## 文件目录

> 使用yum安装的文件结构与离线安装不同，通过`locate nginx`来查看

- `/etc/nginx`配置文件
- `/usr/share/nginx/html`静态文件
