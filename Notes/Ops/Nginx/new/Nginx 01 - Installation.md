# Nginx 01 - Installation

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

  编译后 Nginx 使用的默认配置文件

- `src`

  按照模块对应源码

执行配置指令(为编译做准备)，可以使用如下模版，全模块编译(只包含 Nginx 内置的模块)

> 如果想要自定义安装的位置，修改 `--prefix` 的值即可(可以设置一个临时环境变量)

```
./configure \
--prefix=/opt/nginx \
--builddir=./builddir \
--user=nobody \
--group=nobody \
--conf-path=/opt/nginx/nginx.conf \
--sbin-path=/opt/nginx/sbin/nginx \
--http-log-path=/opt/nginx/logs/access.log \
--error-log-path=/opt/nginx/logs/error.log \
--pid-path=/opt/nginx/run/nginx.pid \
--lock-path=/opt/nginx/run/nginx.lock \
--http-client-body-temp-path=/opt/nginx/client-body-temp \
--http-proxy-temp-path=/opt/nginx/proxy-temp \
--http-fastcgi-temp-path=/opt/nginx/fastcgi-temp \
--http-scgi-temp-path=/opt/nginx/scgi-temp \
--http-uwsgi-temp-path=/opt/nginx/uwsgi-temp \
--with-threads \
--with-file-aio \
--with-pcre \
--with-pcre-jit \
--with-mail \
--with-mail_ssl_module \
--with-libatomic \
--with-stream \
--with-stream_ssl_module \
--with-stream_realip_module \
--with-http_perl_module \
--with-http_ssl_module \
--with-http_addition_module \
--with-http_xslt_module \
--with-http_v2_module \
--with-http_realip_module \
--with-http_image_filter_module \
--with-http_geoip_module \
--with-http_sub_module \
--with-http_dav_module \
--with-http_flv_module  \
--with-http_mp4_module  \
--with-http_gunzip_module \
--with-http_gzip_static_module \
--with-http_auth_request_module \
--with-http_random_index_module \
--with-http_secure_link_module \
--with-http_slice_module \
--with-http_stub_status_module
```

如果只想用作 HTTP/HTTPS 服务器使用如下指令即可

```
./configure \
--prefix=/opt/nginx \
--builddir=./builddir \
--user=nobody \
--group=nobody \
--conf-path=/opt/nginx/nginx.conf \
--sbin-path=/opt/nginx/sbin/nginx \
--http-log-path=/opt/nginx/logs/access.log \
--error-log-path=/opt/nginx/logs/error.log \
--pid-path=/opt/nginx/run/nginx.pid \
--lock-path=/opt/nginx/run/nginx.lock \
--http-client-body-temp-path=/opt/nginx/client-body-temp \
--http-proxy-temp-path=/opt/nginx/proxy-temp \
--http-fastcgi-temp-path=/opt/nginx/fastcgi-temp \
--http-scgi-temp-path=/opt/nginx/scgi-temp \
--http-uwsgi-temp-path=/opt/nginx/uwsgi-temp \
--with-threads \
--with-file-aio \
--with-pcre-jit \
--with-libatomic \
--with-http_ssl_module \
--with-http_gunzip_module \
--with-http_gzip_static_module \
--with-http_stub_status_module \
--with-pcre=/opt/pcre2-10.43 \
--with-openssl=/opt/openssl-3.2.1 \
--with-zlib=/opt/zlib-1.3.1 \
--with-libatomic=/opt/libatomic_ops-7.8.2
```

如果报错，那是因为对应的模块有依赖，需要手动安装，才可以正常编译

1. 如果出现 `./configure: error: C compiler cc is not found.` 那就需要安装 `gcc` 用于编译 C 文件[^3]，以 rhel 为例，可以执行 `yum install gcc`

   如果没法使用包管理，可以下载 gcc 源码编译安装(逻辑上是更新)

   https://gcc.gnu.org/

   ```
   curl -OL http://ftp.tsukuba.wide.ad.jp/software/gcc/releases/gcc-13.2.0/gcc-13.2.0.tar.gz
   tar -xzvf gcc-13.2.0.tar.gz && cd gcc-13.2.0
   mkdir build && cd build
   ../configure --prefix /opt/gcc
   ```

   执行编译 gcc 需要主机上有 `cc` 或者是 `gcc` (完全就是 catch-22)[^7]

   如果没有只能从同环境且包含 `cc` 或者 `gcc` 的主机上移植一个 binary 文件

2. 如果出现 `./configure: error: the HTTP rewrite module requires the PCRE library.` 对应 `--with-pcre` 那就需要安装 `pcre-devel`，以 rhel 为例，可以执行 `yum install pcre-devel`

   如果没办法使用包管理，可以下载 pcre 源码编译安装

   https://github.com/PCRE2Project/pcre2

   ```
   curl -OL https://github.com/PCRE2Project/pcre2/releases/download/pcre2-10.43/pcre2-10.43.tar.gz
   tar -xzvf pcre2-10.43.tar.gz && cd pcre2-10.43
   ./configure && make install
   ```

   或者使用 `--with-pcre=path` 指定 pcre 的源码包

   ```
   ./configure \
   ...
   --with-pcre=/opt/pcre2-10.43 \
   ...
   ```

   或者 configure 的时候不使用 `--with-pcre` (不能使用 [ngx_http_rewrite_module](http://nginx.org/en/docs/http/ngx_http_rewrite_module.html) 内的指令块)

3. 如果出现 `./configure: error: the HTTP gzip module requires the zlib library.` 对应 `--with-http_gunzip_module` 那就需要安装 `zlib-devel`, 以 rhel 为例，可以执行 `yum install zlib-devel`

   如果没办法使用包管理，可以下载 zlib 源码编译安装

   https://github.com/madler/zlib

   ```
   curl -OL https://github.com/madler/zlib/releases/download/v1.3.1/zlib-1.3.1.tar.gz
   tar -xzvf zlib-1.3.1.tar.gz && cd zlib-1.3.1
   ./configure && make install
   ```

   或者使用 `--with-zlib=path` 指定 pcre 的源码包

   ```
   ./configure \
   ...
   --with-zlib=/opt/zlib-1.3.1 \
   ...
   ```

   或者 configure 的时候不使用 `--with-zlib` (不能使用 [ngx_http_gzip_module](http://nginx.org/en/docs/http/ngx_http_gzip_module.html) 内的指令块)

4. 如果出现 `./configure: error: SSL modules require the OpenSSL library.`. 对应 `--with-http_ssl_module` 那就需要安装 `openssl-devel`，以 rhel 为例，可以执行 `yum install openssl-devel`

   如果没办法使用包管理，可以下载 openssl 源码编译安装

   https://www.openssl.org/

   ```
   curl -OL https://www.openssl.org/source/openssl-3.2.1.tar.gz
   sudo tar -xzvf openssl-3.2.1.tar.gz && cd openssl-3.2.1
   ./configure && make install
   ```

   或者使用 `--with-openssl=path` 指定 pcre 的源码包

   ```
   ./configure \
   ...
   --with-openssl=/opt/openssl-3.2.1 \
   ...
   ```

   或者 configure 的时候不使用带 ssl 的参数

5. 如果出现 `./configure: error: the GeoIP module requires the GeoIP library.` 对应 `--with-http_geoip_module` 那就需要安装 `geoip-devel`, 以 rhel 为例，可以执行 `yum install geoip-devel`

   如果没有 `geoip-devel` 那就先执行 `yum install epel-release ` 安装额外的 repo [^4]

   或者不使用 `--with-http_geoip_module` 或者是 `--with-stream_geoip_module` (不能使用 [ngx_http_geoip_module](http://nginx.org/en/docs/http/ngx_http_geoip_module.html) 内的指令块)

6. 如果出现`./configure: error: invalid option "--with-http_v3_module"` 那是因为 Nginx 将对应的指令删除了[^5], 具体也可以使用 `./configure --help` 来佐证

7. 如果出现 `./configure: error: the HTTP XSLT module requires the libxml2/libxslt libraries.` 对应 `--with-http_xslt_module` 那就需要安装 `libxslt-devel`, 以 rhel 为例，可以执行 `yum install libxslt-devel`

   或者不使用 `--with-http_xslt_module` (不能是使用 [ngx_http_xslt_module](http://nginx.org/en/docs/http/ngx_http_xslt_module.html) 内的指令块)

8. 如果出现 `./configure: error: the HTTP image filter module requires the GD library.` 对应 `--with-http_image_filter_module` 那就需要安装 `gd-devel`, 以 rhel 为例，可以执行 `yum install gd-devel`

   或者不使用 `--with-http_image_fileter_module` (不能是使用 [ngx_http_image_filter_module](http://nginx.org/en/docs/http/ngx_http_image_filter_module.html) 内的指令块)

9. 如果出现 `./configure: error: perl 5.8.6 or higher is required`， 那就需要升级 perl，以 rhel 为例，可以执行 `yum install perl`

10. 如果出现 `./configure: error: perl module ExtUtils::Embed is required`，那就需要安装 `perl-ExtUtils-Embed`, 以 rhel 为例，可以执行 `yum install perl-ExtUtils-Embed`

11. 如果出现 `./configure: error: libatomic_ops library was not found.`, 对应 `--with-libatomic`，那就需要安装 `libatomic_ops-devel`， 以 rhel 为例，可以执行 `yum install libatomic_ops-devel`

    如果没办法使用包管理，可以下载 libatomic 源码编译安装

    https://github.com/ivmai/libatomic_ops

    ```
    curl -OL https://github.com/ivmai/libatomic_ops/releases/download/v7.8.2/libatomic_ops-7.8.2.tar.gz
    tar -xzvf libatomic_ops-7.8.2.tar.gz && cd libatomic_ops-7.8.2
    ./configure && make install
    ```

    或者使用 `--with-libatomic=path` 指定 libatomic 的源码包

    ```
    ./configure \
    ...
    --with-libatomic=/opt/libatomic_ops-7.8.2 \
    ...
    ```

配置完成后提示编译后使用的目录(如果没有按照上述模板配置会使用默认值)

> 需要注意的一点是 `html` 目录是默认在 prefix 下的，不会像通过 package manager 安装的方式在 `/usr/share/nginx/` 下
>
> 如果为了和标准的相同可以手动创建目录，修改配置文件的 `root` 指令块值即可 

```
Configuration summary
  + using threads
  + using PCRE2 library: /opt/pcre2-10.43
  + using OpenSSL library: /opt/openssl-3.2.1
  + using zlib library: /opt/zlib-1.3.1
  + using system libatomic_ops library

  nginx path prefix: "/opt/nginx"
  nginx binary file: "/opt/nginx/sbin/nginx"
  nginx modules path: "/opt/nginx/modules"
  nginx configuration prefix: "/opt/nginx"
  nginx configuration file: "/opt/nginx/nginx.conf"
  nginx pid file: "/opt/nginx/run/nginx.pid"
  nginx error log file: "/opt/nginx/logs/error.log"
  nginx http access log file: "/opt/nginx/logs/access.log"
  nginx http client request body temporary files: "/opt/nginx/client-body-temp"
  nginx http proxy temporary files: "/opt/nginx/proxy-temp"
  nginx http fastcgi temporary files: "/opt/nginx/fastcgi-temp"
  nginx http uwsgi temporary files: "/optnginx/uwsgi-temp"
  nginx http scgi temporary files: "/opt/nginx/scgi-temp"
```

### 0x023 make

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

可能出现的错误会有

1. `Can't locate IPC/Cmd.pm in @INC`

   `make[1]: *** [/opt/openssl-3.2.1/.openssl/include/openssl/ssl.h] Error 2`

   安装 `perl-IPC-Cmd` 即可

### 0x024 Post installation

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

Nginx 安装后可以通过 `nginx -V` 来校验并查看对应的编译指令

```
[vagrant@localhost sbin]$ ./nginx -V
nginx version: nginx/1.24.0
built by gcc 4.8.5 20150623 (Red Hat 4.8.5-44) (GCC)
built with OpenSSL 3.2.1 30 Jan 2024
TLS SNI support enabled
configure arguments: --builddir=build --prefix=/opt/nginx --user=nobody --group=nobody --conf-path=/opt/nginx/nginx.conf --sbin-path=/opt/nginx/sbin/nginx --http-log-path=/opt/nginx/logs/access.log --error-log-path=/opt/nginx/logs/error.log --pid-path=/opt/nginx/run/nginx.pid --lock-path=/opt/nginx/run/nginx.lock --http-client-body-temp-path=/opt/nginx/client-body-temp --http-proxy-temp-path=/opt/nginx/proxy-temp --http-fastcgi-temp-path=/opt/nginx/fastcgi-temp --http-scgi-temp-path=/opt/nginx/scgi-temp --http-uwsgi-temp-path=/opt/nginx/uwsgi-temp --with-threads --with-file-aio --with-pcre-jit --with-mail --with-mail_ssl_module --with-libatomic --with-stream --with-stream_ssl_module --with-stream_realip_module --with-http_perl_module --with-http_ssl_module --with-http_addition_module --with-http_xslt_module --with-http_v2_module --with-http_realip_module --with-http_addition_module --with-http_image_filter_module --with-http_geoip_module --with-http_sub_module --with-http_dav_module --with-http_flv_module --with-http_mp4_module --with-http_gunzip_module --with-http_gzip_static_module --with-http_auth_request_module --with-http_random_index_module --with-http_secure_link_module --with-http_slice_module --with-http_stub_status_module --with-pcre=/opt/pcre2-10.43 --with-openssl=/opt/openssl-3.2.1 --with-zlib=/opt/zlib-1.3.1
```

### 0x025 configure cli

记录常见的一些构建参数，具体详情查看 [configure](http://nginx.org/en/docs/stream/ngx_stream_ssl_module.html)

- `--prefix=path`

  编译后 Nginx 所有文件存储的路径，默认 `/usr/local/nginx`

- `–-sbin-path=path`

  编译后 binary 文件名（即启动 nginx 的指令），默认 `prefix/sbin/nginx`


- `--modules-path=path`

  编译后自定义模块的存储路径，默认 `prefix/modules`

- `--conf-path=path`

  编译后 Nginx 默认配置文件的路径，默认使用 `prefix/conf/nginx.conf`

  启动 Nginx 时可以手动使用 `-c` 来指定配置文件


- `--http-log-path=path`

  编译后 Nginx access.log 存储的路径，默认 `prefix/logs/access.log`


- `--error-log-path=path`

  编译后 Nginx 默认错误日志存储的路径,默认 `prefix/logs/error.log`

  可以通过 `error_log` directive 来修改

- `--pid-path=path`

  启动 Nginx 后 pid 存储的路径，默认 `prefix/logs/nginx.pid`

  可以通过 `pid` directive 来修改

- `--lock-path=path`

  启动 Nginx 后 lock 文件(类似 pacman.lock 的原理)存储的路径，默认 `prefix/logs/nginx.lock`

  可以通过 `lock_file` directive 来修改

- `--user=name`

  设置可以调用 Nginx 进程的用户名，默认 `nobody`[^6]

  可以通过 `user` directive 来修改

- `–-group=name`

  设置可以调用 Nginx 进程的用户组，默认和 `--user=name` 的值相同

  可以通过 `user` directive 来修改


- `--builddir=path`

  设置编译时使用的 working directory

- `--with-select_module`

  允许 nginx 调用 `select()`，默认构建

- `--with-poll_module`

  允许 nginx 调用 `poll()`，默认构建

- `--with-threads`

  允许使用 `thread_pool` directives

- `--with-file-aio`

  允许开启 asynchronous file I/O

- `--with-http_ssl_module`

  增加 [HTTPS protocol support](http://nginx.org/en/docs/http/ngx_http_ssl_module.html) 指令块的模组，使用 HTTPS 需要指定，默认不构建(通常需要指定)

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

- `--with-http_flv_module`

  增加 [ngx_http_flv_module](http://nginx.org/en/docs/http/ngx_http_flv_module.html) 指令块，默认不构建

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

- `--with-http_secure_link_module`

  增加 [ngx_http_secure_link_module](http://nginx.org/en/docs/http/ngx_http_secure_link_module.html) 指令块，默认不构建

- `--with-http_degradation_module`

- `--with-http_stub_status_module`

  增加 [ngx_http_stub_status_module](http://nginx.org/en/docs/http/ngx_http_stub_status_module.html) 指令块，可以生成一个简单的 status 页面

- `--with-perl_module`

  `--with-perl_module=dynamic`

  增加 [embedded Perl module](http://nginx.org/en/docs/http/ngx_http_perl_module.html) 功能，使用内嵌脚本需要指定,默认不构建

- `--with-perl_modules_path=path`

  用于构建的 perl module 目录

- `--http-client-body-temp-path=path`

  存储 request body 的临时文件路径，默认 `prefix/client_body`

  可以通过 `client_body_temp_path` directive 来修改

- `--http-proxy-temp-path=path`

  存储从 proxied 服务器收到临时文件的路径，默认 `prefix/proxy_temp`

  可以通过 `proxy_tmemp_path` directive 来修改

- `--http-fastcgi-temp-path=path`

  存储从 FastCGI 服务器收到临时文件的路径，默认 `prefix/fastcgi_temp`

  可以通过 `fastcgi_temp_path` directive 来修改

- `--http-uwsgi-temp-path=path`

  存储从 uwsgi 服务器收到临时文件的路径，默认 `prefix/uwsgi_temp`

  可以通过 `uwsgi_temp_path` directive 来修改

- `--http-scgi-temp-path=path`

  存储从 scgi 服务器收到临时文件的路径，默认 `prefix/scgi_temp`

  可以通过 `scgi_temp_path` directive 来修改

- `--with-mail`

  `--with-mail=dynamic`

  增加 [mail proxy](http://nginx.org/en/docs/mail/ngx_mail_core_module.html) 指令块，允许构建 POP3/IMAP4/SMTP 服务，默认不构建

- `--with-mail_ssl_module`

  增加 [ngx_mail_ssl](http://nginx.org/en/docs/mail/ngx_mail_ssl_module.html) 指令块，默认不构建

- `--with-stream`

  `--with-stream=dynamic`

  增加 [stream module](http://nginx.org/en/docs/stream/ngx_stream_core_module.html) 指令块，默认不构建，和 `upstream` direcitve 无关

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

  启用动态模块兼容

- `--with-pcre`

  使用 [location](http://nginx.org/en/docs/http/ngx_http_core_module.html#location) 和 [ngx_http_rewrite_module](http://nginx.org/en/docs/http/ngx_http_rewrite_module.html) 指令块必须指定的

- `--with-pcre=path`

  指定 pcre source code 的路径(可以无需编译，交由 Nginx 自己处理)

- `--with-pcre-jit`

  允许使用 `pcre_jit` directive 加快正则的处理

- `--with-libatomic`

  强制使用 libatomic_ops libaray


- `--with-zlib=path`

  指定 zlib source code 的路径(可以无需编译，交由 Nginx 自己处理)

- `--with-openssl=path`

  指定 openssl source code 的路径(可以无需编译，交由 Nginx 自己处理)

- `--with-cc=path`

  指定 C comipler cc ==binary== 的路径(需要编译 gcc)

- `--with-cpp=path`

  指定 Cpp comipler binary 的路径

- `--with-debug`

  允许记录 debug log

- `--add-modules=path`

  构建时额外使用的 modules 路径

### 0x026 How to uninstall

针对编译的 Nginx，直接删除对应的目录即可

**references**

1. [^1]:http://nginx.org/en/download.html

2. [^2]:http://nginx.org/en/docs/configure.html

3. [^3]:https://www.google.com/search?q=.%2Fconfigure%3A+error%3A+C+compiler+cc+is+not+found&oq=.%2Fconfigure%3A+error%3A+C+compiler+cc+is+not+found&gs_lcrp=EgZjaHJvbWUyBggAEEUYOdIBBzIyMGowajSoAgCwAgA&sourceid=chrome&ie=UTF-8

4. [^4]:https://serverfault.com/questions/372978/installing-geoip-on-centos

5. [^5]:https://hg.nginx.org/nginx/rev/113e2438dbd4

6. [^6]:https://unix.stackexchange.com/questions/186568/what-is-nobody-user-and-group

7. [^7]:https://unix.stackexchange.com/questions/378060/gcc-error-no-acceptable-c-compiler-found-in-path

8. [^8]:https://docs.nginx.com/nginx/admin-guide/installing-nginx/installing-nginx-open-source/#installing-nginx-dependencies