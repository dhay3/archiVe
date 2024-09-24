# Linux passwd

`passwd`用于修改用户的密码，`/etc/login.defs`会影响该工具的使用(例如：ENCRYPT_METHOD，OBSCURE_CHECKS_ENAB)，具体[参考](../config/login.defs.md)

## options

- -e | --expire

  让密码立即失效

  ```
  cpl in / λ sudo passwd -e ttt
  passwd: password expiry information changed.
  cpl in / λ sudo passwd -S ttt
  ttt P 01/01/1970 0 99999 7 -1
  ```

- -d | --delete

  删除用户的密码，置位空

- -l | --lock 

  将用户的密码上锁，可以使用`-u`解锁。用户任然可以通过ssh登入，但是不能通过`su`等方式登入

  ```
  cpl in / λ sudo passwd -l ttt
  passwd: password expiry information changed.
  cpl in / λ su ttt            
  Password: 
  su: Authentication failure
  cpl in / λ sudo passwd -u ttt
  passwd: password expiry information changed.
  cpl in / λ su ttt            
  Password: 
  [ttt@cyberpelican /]$ 
  ```

- -S | --status

  展示当前用户账户的信息

  ```
  cpl in / λ passwd -S
  cpl P 05/28/2021 0 99999 7 -1
  ```

  1. 用户名
  2. 密码是否上锁，L(locked)，NP(no password)，P(usable password)
  3. 密码变更的时间
  4. minimum age
  5. maximum age
  6. warning period
  7. inactivity period

- -n | --mindays
- -x | --maxdays
- -w | --wardays