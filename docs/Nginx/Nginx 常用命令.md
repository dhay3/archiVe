# Nginx 常用命令

## 离线安装命令

以下命令都在`/usr/local/nginx/sbin`目录下执行

- 启动nginx

  ```
  ./nginx
  ```

- 关闭

  ```
  ./nginx -s stop
  ```

- 重启(linux下会热部署)

  ```
  ./nginx -s reload
  ```

## epel安装命令

- 查看nginx版本

  ```
  [root@cyberpelican ~]# nginx -v
  nginx version: nginx/1.16.1
  ```

- 启动nginx

  ```
  systemctl start nginx
  ```

- 查看nginx状态

  ```
  [root@cyberpelican nginx]# systemctl status nginx
  ● nginx.service - The nginx HTTP and reverse proxy server
     Loaded: loaded (/usr/lib/systemd/system/nginx.service; enabled; vendor preset: disabled)
     Active: active (running) since Wed 2020-11-25 10:17:55 CST; 5s ago
    Process: 4770 ExecReload=/bin/kill -s HUP $MAINPID (code=exited, status=1/FAILURE)
    Process: 4926 ExecStart=/usr/sbin/nginx (code=exited, status=0/SUCCESS)
    Process: 4923 ExecStartPre=/usr/sbin/nginx -t (code=exited, status=0/SUCCESS)
    Process: 4922 ExecStartPre=/usr/bin/rm -f /run/nginx.pid (code=exited, status=0/SUCCESS)
   Main PID: 1571 (code=exited, status=0/SUCCESS)
      Tasks: 2
     CGroup: /system.slice/nginx.service
             ├─4929 nginx: master process /usr/sbin/nginx
             └─4931 nginx: worker process
  
  Nov 25 10:17:55 cyberpelican systemd[1]: Starting The nginx HTTP and reverse proxy server...
  Nov 25 10:17:55 cyberpelican nginx[4923]: nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
  Nov 25 10:17:55 cyberpelican nginx[4923]: nginx: configuration file /etc/nginx/nginx.conf test is successful
  Nov 25 10:17:55 cyberpelican systemd[1]: Failed to parse PID from file /run/nginx.pid: Success
  Nov 25 10:17:55 cyberpelican systemd[1]: Started The nginx HTTP and reverse proxy server.
  ```

  这里可以发现使用的配置文件是`/etc/nginx/nginx.conf`

- 关闭nginx

  ```
  systemctl stop nginx
  ```

  如果使用`nginx -s stop`来关闭nginx可能会出现问题，使用`killall  nginx`来杀死进程

- 自动启动nginx

  ```
  systemctl enable nginx
  ```

- 重启(linux下会热部署)

  ```
  systemctl reload nginx
  ```

- 检查配置文件

  ```
  nginx -t
  ```

  
