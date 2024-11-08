# DoH vs DoT

> DNS queries and reponses are sent in plain text via UDP

这意味着DNS返回的信息可以被ISP(DNS posion)，network，wifi-admin读取，即使是HTTPS

## DoH

DNS over HTTPS，通过HTTP 或者 HTTP/2 而不是通过UDP直接和DNS交互，7 layer，使用443端口(==意味这可以和https流量一起混淆，且不容易被firewall拦截==)

- firefox : https://support.mozilla.org/en-US/kb/firefox-dns-over-https

## DoT

DNS over TLS，通过TLS和DNS交互，4 layer，使用853端口(==意味这更容易被firewall拦截==) 

## cloudflared

https://developers.cloudflare.com/1.1.1.1/dns-over-https/cloudflared-proxy

浏览器可以通过简单的设置让dns quries走DoH，如果想让OS的DNS queries走DoH，可以使用cloudflared(cloudflare开源的一个GO项目)

### install

https://developers.cloudflare.com/cloudflare-one/connections/connect-apps/install-and-setup/installation#linux

可以使用build from source，方便之后升级。编译之后会在当前目录出现一个可执行文件`cloudflared`。

为了方便也可以使用docker来安装

```
cpl in ~/Downloads λ coproc docker run --rm --name cldf  cloudflare/cloudflared proxy-dns --port 5553 
[1] 92297
```

