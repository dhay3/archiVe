# Nginx Installation

## 0x01 Package manager

安装 Nginx 最简单的方式就是通过包管理器安装

http://nginx.org/en/linux_packages.html#RHEL

## 0x02 Building from Sources

### 0x021 Prepare

当然 Nginx 也支持使用源码编译，首先下载源码包，可以在 [下载页面](http://nginx.org/en/download.html) 选择想要安装的版本

```
curl -O http://nginx.org/download/nginx-1.24.0.tar.gz
```

检查签名

```
#下载签名文件
[root@d2f7bc37acff ~]# curl -O http://nginx.org/download/nginx-1.24.0.tar.gz.asc

#检查签名使用的公钥指纹
[root@d2f7bc37acff ~]# gpg --verify nginx-1.24.0.tar.gz.asc             
gpg: assuming signed data in 'nginx-1.24.0.tar.gz'
gpg: Signature made Tue 11 Apr 2023 04:04:31 PM UTC
gpg:                using RSA key 13C82A63B603576156E30A4EA0EA981B66B0D967
gpg:                issuer "k.pavlov@f5.com"
gpg: Can't check signature: No public key

#导入公钥
[root@d2f7bc37acff ~]# gpg --keyserver hkps://keyserver.ubuntu.com --search-keys 13C82A63B603576156E30A4EA0EA981B66B0D967
gpg: data source: https://keyserver.ubuntu.com:443
(1)     Konstantin Pavlov <k.pavlov@f5.com>
        Konstantin Pavlov <thresh@nginx.com>
          3072 bit RSA key A0EA981B66B0D967, created: 2018-05-07
Keys 1-1 of 1 for "13C82A63B603576156E30A4EA0EA981B66B0D967".  Enter number(s), N)ext, or Q)uit > 1
gpg: key A0EA981B66B0D967: public key "Konstantin Pavlov <thresh@nginx.com>" imported
gpg: Total number processed: 1
gpg:               imported: 1

#校验指纹
[root@d2f7bc37acff ~]# gpg --verify nginx-1.24.0.tar.gz.asc 
gpg: assuming signed data in 'nginx-1.24.0.tar.gz'
gpg: Signature made Tue 11 Apr 2023 04:04:31 PM UTC
gpg:                using RSA key 13C82A63B603576156E30A4EA0EA981B66B0D967
gpg:                issuer "k.pavlov@f5.com"
gpg: Good signature from "Konstantin Pavlov <thresh@nginx.com>" [unknown]
gpg:                 aka "Konstantin Pavlov <k.pavlov@f5.com>" [unknown]
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: 13C8 2A63 B603 5761 56E3  0A4E A0EA 981B 66B0 D967
```

### 0x022 Configure

确认签名是正常的后，解压并进入对应的目录

```
tar -xvzf nginx-1.24.0.tar.gz && cd nginx-1.24.0
```

查看一下解压后有那些文件

```
[root@d2f7bc37acff nginx-1.24.0]# tree -L 2
.
├── auto
│   ├── cc
│   ├── define
│   ├── endianness
│   ├── feature
│   ├── have
│   ├── have_headers
│   ├── headers
│   ├── include
│   ├── init
│   ├── install
│   ├── lib
│   ├── make
│   ├── module
│   ├── modules
│   ├── nohave
│   ├── options
│   ├── os
│   ├── sources
│   ├── stubs
│   ├── summary
│   ├── threads
│   ├── types
│   └── unix
├── CHANGES
├── CHANGES.ru
├── conf
│   ├── fastcgi.conf
│   ├── fastcgi_params
│   ├── koi-utf
│   ├── koi-win
│   ├── mime.types
│   ├── nginx.conf
│   ├── scgi_params
│   ├── uwsgi_params
│   └── win-utf
├── configure
├── contrib
│   ├── geo2nginx.pl
│   ├── README
│   ├── unicode2nginx
│   └── vim
├── html
│   ├── 50x.html
│   └── index.html
├── LICENSE
├── man
│   └── nginx.8
├── README
└── src
    ├── core
    ├── event
    ├── http
    ├── mail
    ├── misc
    ├── os
    └── stream

19 directories, 38 files
```

- `configure`

  是一个 Bash 脚本，会按照 auto 下的目录执行，用于生成 Makefile

- `conf`

  编译后默认的配置文件

- `src`

  按照模块对应源码

执行配置指令(为编译做准备)，可以使用如下模版，全模块编译(只包含 Nginx 内置的模块)

```
./configure \
--prefix=/etc/nginx \
--conf-path=/etc/nginx/nginx.conf \
--sbin-path=/usr/bin/nginx \
--pid-path=/run/nginx.pid \
--lock-path=/run/lock/nginx.lock \
--user=nginx \
--group=nginx \
--http-log-path=/var/log/nginx/access.log \
--error-log-path=/var/log/nginx/error.log \
--http-client-body-temp-path=/var/lib/nginx/client-body \
--http-proxy-temp-path=/var/lib/nginx/proxy \
--http-fastcgi-temp-path=/var/lib/nginx/fastcgi \
--http-scgi-temp-path=/var/lib/nginx/scgi \
--http-uwsgi-temp-path=/var/lib/nginx/uwsgi \
--with-compat \
--with-debug \
--with-file-aio \
--with-http_addition_module \
--with-http_auth_request_module \
--with-http_dav_module \
--with-http_flv_module \
--with-http_geoip_module \
--with-http_gunzip_module \
--with-http_gzip_static_module \
--with-http_mp4_module \
--with-http_random_index_module \
--with-http_realip_module \
--with-http_secure_link_module \
--with-http_slice_module \
--with-http_ssl_module \
--with-http_stub_status_module \
--with-http_sub_module \
--with-http_v2_module \
--with-mail \
--with-mail_ssl_module \
--with-pcre-jit \
--with-stream \
--with-stream_geoip_module \
--with-stream_realip_module \
--with-stream_ssl_module \
--with-stream_ssl_preread_module \
--with-threads
```

如果报错，那是因为对应的模块有依赖，需要手动安装，才可以正常编译

1. 如果出现 `./configure: error: C compiler cc is not found.` 那就需要安装 `gcc` 用于编译 C 文件[^3]，以 rhel 为例，可以执行 `yum install gcc`

   如果没法使用包管理，可以下载 gcc 源码编译安装

   https://gcc.gnu.org/

2. 如果出现 `./configure: error: the HTTP rewrite module requires the PCRE library.` 对应 `--with-pcre` 那就需要安装 `pcre-devel`，以 rhel 为例，可以执行 `yum install pcre-devel`

   如果没办法使用包管理，可以下载 pcre 源码编译安装

   https://github.com/PCRE2Project/pcre2

3. 如果出现 `./configure: error: the HTTP gzip module requires the zlib library.` 对应 `--with-http_gunzip_module` 那就需要安装 `zlib-devel`, 以 rhel 为例，可以执行 `yum install zlib-devel`

   如果没办法使用包管理，可以下载 zlib 源码编译安装

   https://github.com/madler/zlib

4. 如果出现 `./configure: error: SSL modules require the OpenSSL library.`. 对应 `--with-http_ssl_module` 那就需要安装 `openssl-devel`，以 rhel 为例，可以执行 `yum install openssl-devel`

   如果没办法使用包管理，可以下载 openssl 源码编译安装

   https://www.openssl.org/

5. 如果出现 `./configure: error: the GeoIP module requires the GeoIP library.` 对应 `--with-http_geoip_module` 那就需要安装 `geoip-devel`, 以 rhel 为例，可以执行 `yum install geoip-devel`

   如果没有 `geoip-devel` 那就先执行 `yum install epel-release ` 安装额外的 repo [^4]

6. 如果出现`./configure: error: invalid option "--with-http_v3_module"` 那是因为 Nginx 将对应的指令删除了[^5], 具体也可以使用 `./configure --help` 来佐证

配置完成后提示编译后使用的目录(如果没有按照上述模板配置会使用默认值)

> 需要注意的一点是 `html` 目录是默认在 prefix 下的，不会像通过 package manager 安装的方式在 `/usr/share/nginx/` 下
>
> 如果为了和标准的相同可以手动创建目录，修改配置文件的 `root` 指令块值即可 

```
  nginx path prefix: "/etc/nginx"
  nginx binary file: "/usr/bin/nginx"
  nginx modules path: "/etc/nginx/modules"
  nginx configuration prefix: "/etc/nginx"
  nginx configuration file: "/etc/nginx/nginx.conf"
  nginx pid file: "/run/nginx.pid"
  nginx error log file: "/var/log/nginx/error.log"
  nginx http access log file: "/var/log/nginx/access.log"
  nginx http client request body temporary files: "/var/lib/nginx/client-body"
  nginx http proxy temporary files: "/var/lib/nginx/proxy"
  nginx http fastcgi temporary files: "/var/lib/nginx/fastcgi"
  nginx http uwsgi temporary files: "/var/lib/nginx/uwsgi"
  nginx http scgi temporary files: "/var/lib/nginx/scgi"
```

同时会生成一个 `Makefile` 用于编译 Nginx (其实还有一个 `objs` 目录，存储实际使用的 Makefile 和 编译后的文件，例如 nginx)

```
├── Makefile
...
├── objs
│   ├── autoconf.err
│   ├── Makefile
│   ├── ngx_auto_config.h
│   ├── ngx_auto_headers.h
│   ├── ngx_modules.c
│   └── src
```

> 如果提示 `bash: make: command not found` 安装 `make` 即可

1. 如果想要编译，执行 `make build`，会将编译出来的文件存储在 `objs`

   ```
   ├── objs
   │   ├── autoconf.err
   │   ├── Makefile
   │   ├── nginx
   │   ├── nginx.8
   │   ├── ngx_auto_config.h
   │   ├── ngx_auto_headers.h
   │   ├── ngx_modules.c
   │   ├── ngx_modules.o
   │   └── src
   ```

   如果想要重新编译可以使用 `make clean` 来清空 `./configure` 产生的 `objs` 和 `Makefile`

2. 如果想要编译并安装到对应的目录，执行 `make install` (正常一般没啥问题直接执行这个就可以了)

Nginx 安装后可以通过 `nginx -V` 来校验并查看对应的编译指令

```
[root@d2f7bc37acff nginx]# nginx -V
nginx version: nginx/1.24.0
built by gcc 8.5.0 20210514 (Red Hat 8.5.0-4) (GCC) 
built with OpenSSL 1.1.1k  FIPS 25 Mar 2021
TLS SNI support enabled
configure arguments: --prefix=/etc/nginx --conf-path=/etc/nginx/nginx.conf --sbin-path=/usr/bin/nginx --pid-path=/run/nginx.pid --lock-path=/run/lock/nginx.lock --user=nginx --group=nginx --http-log-path=/var/log/nginx/access.log --error-log-path=/var/log/nginx/error.log --http-client-body-temp-path=/var/lib/nginx/client-body --http-proxy-temp-path=/var/lib/nginx/proxy --http-fastcgi-temp-path=/var/lib/nginx/fastcgi --http-scgi-temp-path=/var/lib/nginx/scgi --http-uwsgi-temp-path=/var/lib/nginx/uwsgi --with-compat --with-debug --with-file-aio --with-http_addition_module --with-http_auth_request_module --with-http_dav_module --with-http_flv_module --with-http_geoip_module --with-http_gunzip_module --with-http_gzip_static_module --with-http_mp4_module --with-http_random_index_module --with-http_realip_module --with-http_secure_link_module --with-http_slice_module --with-http_ssl_module --with-http_stub_status_module --with-http_sub_module --with-http_v2_module --with-mail --with-mail_ssl_module --with-pcre-jit --with-stream --with-stream_geoip_module --with-stream_realip_module --with-stream_ssl_module --with-stream_ssl_preread_module --with-threads
```

通过 `nginx -t` 校验 Nginx 配置是否正常

1. 如果出现 `nginx: [emerg] getpwnam("nginx") failed` 那是因为，编译的时候指定了 `--user=nginx` 和 `--group=nginx`,但是当前环境没有对应的 user 和 usergroup。可以通过 2 种方式解决

   1. 添加 nginx user nginx usergroup

      ```
      #创建用户的同时会创建用户组
      [root@d2f7bc37acff home]# useradd nginx
      ```

      这里推荐使用下面的命令，不给 nginx 用户分配 login shell 也不创建 home 目录

      ```
      [root@d2f7bc37acff home]# useradd --system --no-create-home --shell /bin/false nginx
      ```

   2. 修改配置文件的 `user` 指令块到当前系统中存在的用户

      ```
      user  nobody;
      ```

2. 如果出现 

   `nginx: [emerg] mkdir() "/var/lib/nginx/client-body" failed (2: No such file or directory)` 

   `nginx: [emerg] mkdir() "/var/lib/nginx/proxy" failed (2: No such file or directory)` 

   `nginx: [emerg] mkdir() "/var/lib/nginx/fast-cgi" failed (2: No such file or directory)` 

   ...

   创建对应的目录即可(Nginx 并不会在安装时创建目录，只会使用 `cp` 命令将编译出来的文件或者是原目录下的文件移动到指定位置，而 client-body 等文件，默认并不会生成，所以交由用户来创建才是合理的)

   ```
   mkdir -p /var/lib/nginx
   ```

到此可以使用 `nginx` 启动 Nginx，然后使用 `curl` 来校验 Nginx 是否正常(默认配置会监听 0.0.0.0:80)

```
[root@d2f7bc37acff ~]# curl -I localhost
HTTP/1.1 200 OK
Server: nginx/1.24.0
Date: Thu, 23 Nov 2023 18:37:13 GMT
Content-Type: text/html
Content-Length: 615
Last-Modified: Thu, 23 Nov 2023 18:11:47 GMT
Connection: keep-alive
ETag: "655f95e3-267"
Accept-Ranges: bytes
```

#### configure

记录常见的一些构建参数，具体详情查看 [configure](http://nginx.org/en/docs/stream/ngx_stream_ssl_module.html)

- `--prefix=path`

  编译后 Nginx 所有文件存储的路径，默认 `/usr/local/nginx`

- `–-sbin-path=path`

  编译后 binary 文件存放的目录，默认 `prefix/sbin/nginx`

- `--conf-path=path`

  编译后 Nginx 默认配置文件的路径，默认使用 `prefix/conf/nginx.conf`

- `--error-log-path=path`

  编译后 Nginx 默认错误日志存储的路径,默认 `prefix/logs/error.log`

- `--pid-path=path`

  启动 Nginx 后 pid 存储的路径，默认 `prefix/logs/nginx.pid`

- `--lock-path=path`

  启动 Nginx 后 lock 文件(类似 pacman.lock 的原理)存储的路径，默认 `prefix/logs/nginx.lock`

- `--user=name`

  设置可以调用 Nginx 进程的用户名，默认 `nobody`

- `–-group=name`

  设置可以调用 Nginx 进程的用户组，默认和 `--user=name` 的值相同

- `--with-select_module`

  允许 nginx 调用 `select()`，默认构建

- `--with-poll_module`

  允许 nginx 调用 `poll()`，默认构建

- `--with-threads`
- `--with-file-aio`

- `--with-http_ssl_module`

  增加 [HTTPS protocol support](http://nginx.org/en/docs/http/ngx_http_ssl_module.html) 指令块的模组，使用 HTTPS 需要指定，默认不构建

- `with-http_v2_module`

  增加 [HTTP/2](http://nginx.org/en/docs/http/ngx_http_v2_module.html) 指令块的模组，使用 HTTP2 需要指定，默认不构建

- `with-http_v3_module`

  增加 [HTTP/3](http://nginx.org/en/docs/http/ngx_http_v3_module.html) 指令块的模组，使用 HTTP3 需要指定，默认不构建

- `--with-http_realip_module`

  增加 [ngx_http_realip_module](http://nginx.org/en/docs/http/ngx_http_realip_module.html) 指令块的模组，用于修改 Nginx 收到报文中的 client address and port，可以通过该指令模组获取真实的 IP 地址，默认不构建

- `--with-http_addition_module`

  增加 [ngx_http_addition_module](http://nginx.org/en/docs/http/ngx_http_addition_module.html) 指令块模组，收到请求后可以将中返回响应前后添加其他的响应，默认不构建

- `--with-http_xslt_module`

  `--with-http_xslt_module=dynamic`

  增加 [ngx_http_xslt_module](http://nginx.org/en/docs/http/ngx_http_xslt_module.html) 指令块模组，支持将 XML 转换为 XSLT，默认不构建

- `--with-http_image_filter_module`

  `--with-http_image_filter_module=dynamic`

  增加 [ngx_http_image_filter_module](http://nginx.org/en/docs/http/ngx_http_image_filter_module.html) 指令块模组，支持修改图片，默认不构建

- `--with-http_geoip_module`

  `--with-http_geoip_module=dynamic`

  增加 [ngx_http_geoip_module](http://nginx.org/en/docs/http/ngx_http_geoip_module.html) 指令块模组，可以按照 geoIP 分流，默认不构建

- `--with-http_sub_module`

  增加 [ngx_http_sub_module](http://nginx.org/en/docs/http/ngx_http_sub_module.html) (subsitute)指令块，允许对返回的 response 做修改，默认不构建

- `--with-http_dav_module`

  增加 [ngx_http_dav_module](http://nginx.org/en/docs/http/ngx_http_dav_module.html) 指令块，如果需要使用 `PUT`, `DELETE` 就必须要要指定，默认不构建

- `--with-http_mp4_module`

  增加 [ngx_http_mp4_module](http://nginx.org/en/docs/http/ngx_http_mp4_module.html) 指令块，默认不构建

- `--with-http_gunzip_module`

  增加 [ngx_http_gunzip_module](http://nginx.org/en/docs/http/ngx_http_gunzip_module.html) 指令块，允许对 response Hearder 返回 `Content-Encoding: gzip` 的 body 解压，默认不构建

- `--with-http_gzip_static_module`

  增加 [ngx_http_gzip_static_module](http://nginx.org/en/docs/http/ngx_http_gzip_static_module.html) 指令块，允许发送以 `.gz` 格式的 precompressed 文件，默认不构建

- `--with-http_auth_request_module`

  增加 [ngx_http_auth_request_module](http://nginx.org/en/docs/http/ngx_http_auth_request_module.html) 指令块，可以发送 subrequest 来做鉴权，默认不构建

- `--with-http_random_index_module`

  增加 [ngx_http_random_index_module](http://nginx.org/en/docs/http/ngx_http_random_index_module.html) 指令块，针对以 slash character 的结尾的请求，会从目录中随机选择一个文件作为 index file，默认不构建

- `--with-http_slice_module`

  增加 [ngx_http_slice_module](http://nginx.org/en/docs/http/ngx_http_slice_module.html) 指令块，可以拆分 request 成 subrequests，针对大包请求有效，默认不构建

- `--with-http_stub_status_module`

  增加 [ngx_http_stub_status_module](http://nginx.org/en/docs/http/ngx_http_stub_status_module.html) 指令块，可以生成一个简单的 status 页面

- `--with-perl_module`

  增加 [embedded Perl module](http://nginx.org/en/docs/http/ngx_http_perl_module.html) 功能，使用内嵌脚本需要指定,默认不构建

- `--http-log-path=path`

  编译后 Nginx access.log 存储的路径，默认 `prefix/logs/access.log`

- `--http-client-body-temp-path=path`

  存储 request body 的临时文件路径，默认 `prefix/client_body`

- `--http-proxy-temp-path=path`

  存储从 proxied 服务器收到临时文件的路径，默认 `prefix/proxy_temp`

- `--http-fastcgi-temp-path=path`

  存储从 FastCGI 服务器收到临时文件的路径，默认 `prefix/fastcgi_temp`

- `--http-uwsgi-temp-path=path`

  存储从 uwsgi 服务器收到临时文件的路径，默认 `prefix/uwsgi_temp`

- `--http-scgi-temp-path=path`

  存储从 scgi 服务器收到临时文件的路径，默认 `prefix/scgi_temp`

- `--with-mail`

  `--with-mail=dynamic`

  增加 [mail proxy](http://nginx.org/en/docs/mail/ngx_mail_core_module.html) 指令块，允许构建 POP3/IMAP4/SMTP 服务，默认不构建

- `--with-mail_ssl_module`

  增加 [ngx_mail_ssl](http://nginx.org/en/docs/mail/ngx_mail_ssl_module.html) 指令块，默认不构建

- `--with-stream`

  增加 [stream module](http://nginx.org/en/docs/stream/ngx_stream_core_module.html) 指令块，使用 `upstream` 指令块需要使用，默认不构建

- `--with-stream_ssl_module`

  增加 [ngx_stream_ssl](http://nginx.org/en/docs/stream/ngx_stream_ssl_module.html)，配置 TLS 需要指定，默认不构建

- `--with-stream_realip_module`

  增加 [ngx_stream_realip_module](http://nginx.org/en/docs/stream/ngx_stream_realip_module.html) 指令块，默认不构建

- `--with-stream_geoip_module`

  `--with-stream_geoip_module=dynamic`

  增加 [ngx_stream_ssl_preread_module](http://nginx.org/en/docs/stream/ngx_stream_ssl_preread_module.html) 指令块，允许按照 geoIP 进行分流，默认不构建

- `--with-stream_ssl_preread_module`

  增加 [ngx_stream_ssl_preread_module](http://nginx.org/en/docs/stream/ngx_stream_ssl_preread_module.html) 指令块，允许按照 TLS ClientHello 中的信息(SNI,protocol version)进行分流，默认不构建

- `--with-compat`

  启用模块兼容

- `--with-pcre`

  使用 [location](http://nginx.org/en/docs/http/ngx_http_core_module.html#location) 和 [ngx_http_rewrite_module](http://nginx.org/en/docs/http/ngx_http_rewrite_module.html) 指令块必须指定的

- `--with-debug`

  允许记录 debug log

**references**

1. [^1]:http://nginx.org/en/download.html

2. [^2]:http://nginx.org/en/docs/configure.html

3. [^3]:https://www.google.com/search?q=.%2Fconfigure%3A+error%3A+C+compiler+cc+is+not+found&oq=.%2Fconfigure%3A+error%3A+C+compiler+cc+is+not+found&gs_lcrp=EgZjaHJvbWUyBggAEEUYOdIBBzIyMGowajSoAgCwAgA&sourceid=chrome&ie=UTF-8

4. [^4]:https://serverfault.com/questions/372978/installing-geoip-on-centos

5. [^5]:https://hg.nginx.org/nginx/rev/113e2438dbd4