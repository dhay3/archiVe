# login.defs

## 0x01 Overview

`login.defs` 是 `shadow-utils` 中的一部分（`rpm -qf /etc/login.defs`），用于配置 shadow password 的策略(例如 `login`,`useradd`, `userdel`,`usermod`, `su` 等指令都会读取该配置)。该配置文件并不会涉及 PAM 模块，所以例如使用 `passwd` 时并不会使用该配置文件（如果想要修改 `passwd` 的策略，具体查看 `/etc/pam.d/system-authq`）

## 0x02 Configuration Directives

指令块的值支持 3 种类型

1. string
2. boolean (yes or no)
3. numbers

> 具体指令块查看 man page，这里只列举常用的指令块
>
> 一般在等保中只需要修改
>
> PASS_MAX_DAYS
>
> PASS_WARN_AGE

- CONSOLE (string)

  只允许 root 在指定的 tty 登入

- FAIL_DELAY (number)

  Delay in seconds before being allowed another attempt after a login failure.

  例如增加如下配置

  ```
  FAIL_DELAY 30
  ```

  如果密码不对，就会卡 Password 30 秒后，才可以登录指定用户

  ```
  [vagrant@localhost ~]$ time su tester
  Password:
  su: Authentication failure
  
  real    0m33.708s
  user    0m0.003s
  sys     0m0.005s
  ```

- ISSUE_FILE (string)

  If defined, this file will be displayed before each login prompt.

- LOGIN_TIMEOUT

- PASS_ALWAYS_WARN

  Warn about weak passwords (but still allow them) if you are root.

- PASS_CHANGE_TRIES

  Maximum number of attempts to change password if rejected (too easy).

- PASS_MAX_DAYS

  The maximum number of days a password may be used. If the password is older than this, a password change will be forced.

  ==只在账号创建的时候生效，不会对现有的账户产生影响==

- PASS_MIN_DAYS

  The minimum number of days allowed between password changes. Any password changes attempted sooner than this will be rejected.q

  ==只在账号创建的时候生效，不会对现有的账户产生影响==

- PASS_WARN_AGE

  The number of days warning given before a password expires. A zero means warning is given only upon the day of expiration, a negative value means no warning is given. If not specified, no warning will be provided.

  只有正数才会提前通知

  ==只在账号创建的时候生效，不会对现有的账户产生影响==

- PASS_MAX_LEN

  PASS_MIN_LEN

  Number of significant characters in the password

  ==和 `passwd` 没有关联（`passwd` 使用 PAM 模块）==

**references**

[^1]:`man login.defs`