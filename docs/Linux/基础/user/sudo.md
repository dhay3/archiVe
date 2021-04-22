# sudo

`sudo`允许用户以superuser或another user 来执行命令。sudo默认的安全策略是sudoers，通过`/etc/sudoers`或通过LDAP来配置。==sudo默认只会打开stdin，stdout，stderr三个流。==

sudoers policy默认会在15mins内缓冲认证(15mins内再次调用sudo无需输入用户密码)，可以使用`-v`来直接更新缓存。

syntax：`sudo [options] [-u user] command `

## 参数

- -i | --login

  用loginshell

- -u [user]

  使用特定的用户替代默认的root用户，可以是uid或是用户名

  ```
  cowrie@win2k:/home/ubuntu$ sudo -u ubuntu -- ls /root/
  ls: cannot open directory '/root/': Permission denied
  cowrie@win2k:/home/ubuntu$ sudo  -- ls /root/
  admin_reserve_kbytes~
  ```

- -l | --list

  如果没有command，展示当前用户可以是以使用的命令

  ```
  cowrie@win2k:/home/ubuntu$ sudo -l
  [sudo] password for cowrie:
  Sorry, user cowrie may not run sudo on win2k.
  
  cowrie@win2k:/home/ubuntu$ sudo -l
  [sudo] password for cowrie:
  Matching Defaults entries for cowrie on win2k:
      env_reset, mail_badpass, secure_path=/usr/local/sbin\:/usr/local/bin\:/usr/sbin\:/usr/bin\:/sbin\:/bin\:/snap/bin
  
  User cowrie may run the following commands on win2k:
      (ALL : ALL) ALL
  ```

- -k 

  ```
  #清空crendentials，下次调用sudo需要用户密码
  sudo -k
  ```

- -A

  使用图形化的界面输入密码。只有在`sudo.conf(5)`中包含`Path askpass <helper>`才会生效，否则报错。

- -b | --background

  以后台的形式运行命令

- -E | --preserver-env

  保留当前shell中的环境变量

- -H | --set-home

  要求sudoer将设置环境变量HOME为用户的家目录

- -s | --shell

  使用当前的SHELL环境变量做为运行的shell

- -U 

  查看指定用户的权限，只能root用户或是有ALL privilege权限的用户

  ```
  cowrie@win2k:/home/ubuntu$ sudo -l -U ubuntu
  Matching Defaults entries for ubuntu on win2k:
      env_reset, mail_badpass, secure_path=/usr/local/sbin\:/usr/local/bin\:/usr/sbin\:/usr/bin\:/sbin\:/bin\:/snap/bin
  
  User ubuntu may run the following commands on win2k:
      (ALL : ALL) ALL
      (ALL) NOPASSWD: ALL
      (ALL) NOPASSWD: ALL
      (ALL : ALL) NOPASSWD: ALL
  ```

- -T

  为命令设置，需要在sudoer中设置