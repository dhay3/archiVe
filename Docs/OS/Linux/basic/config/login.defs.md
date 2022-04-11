# login.defs

## 概述

`/etc/login.defs`配置shadow password suite(`/etc/shadow`)的参数。==需要结合pam==

## Config file

### boolean

boolean类型值为no或yes

- FAILLOG_ENAB：是否在`/var/log/faillog`中记录失败登入
- DEFAULT_HOME：如果不能进入登入用户的HOME，是否允许登入
- LOG_OK_LOGIN：记录登入成功的日志
- LOG_UNKFAIL_ENAB：记录系统未知的username
- SYSLOG_SU_ENAB：记录su的日志
- SYSLOG_SG_ENAB：记录sg的日志
- PASS_ALWAYS_WARN：弱密码提示，如果当前用户是root无效
- USERGROUPS_ENAB：删除用户是删除组，如果组是空的话

### number

number使用decimal，hex，octol

- FAIL_DELAY：权鉴失败间隔时间
- LOGIN_RETRIES：最多失败尝试
- LOGIN_TIMEOUT：在输入密码界面的最长时间
- GID_MAX | GID_MIN：gid分配的范围，默认从1000-60000
- UID_MAX | UID_MIN：uid分配的范围，默认从1000-60000
- SYS_GID_MAX | SYS_GID_MIN：系统用户gid分配的范围
- SYS_UID_MAX | SYS_UID_MIN：系统用户uid分配的范围

> tip. PASS_MIN_DAYS，PASS_MAX_DAYS，PASS_WARN_AGE只对新生成的账号有效，不会对已有的账号产生影响

- PASS_MIN_DAYS：必须在多少天之后才能修改密码，-1表示没有限制。没有指定值就位-1
- PASS_MAX_DAYS：密码最长有效时间
- PASS_WARN_AGE：在密码失效几天前提示，0表示在密码失效当天提示，-1表示不提示。没有指定值就

- PASS_MAX_LEN | PASS_MIN_LEN：密码的最长和最小长度，默认最长8，最好不要修改PASS_MAX_LEN考虑到加密算法
- SHA_CRYPT_MIN_ROUNDS | SHA_CRYPT_MAX_ROUNDS：sha加密的次数，默认使用5000
- PASS_CHANGE_TRIES：修改密码时允许的错误次数

### string

- ENCRYPT_METHOD：系统默认密码加密的方式，支持des，md5，sha256

- ENV_PATH：PATH的值

- ENV_SUPATH：root PATH的值

- CONSOLE：root允许登入的设备也可是一个文件，不能带有`/dev/`前缀

  ```
  CONSOLE			/etc/securetty
  CONSOLE			tty1:tty2
  ```

- FAKE_SHELL：使用指定的shell替代`/etc/passwd`中指定的
- HUSHLOGIN_FILE：输入不常用的字符会被记录，指定存储的位置
- LOGIN_STRING：登入时的prompt，默认`Password:`，`%s`表示用户名
- LOGIN_TIMEOUT：在登入界面的最长时间

- MOTD_FILE：message of the day files

