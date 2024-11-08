# Linux Systemd Loginctl

## Digest

Control the systemd login manager

## Positional args

### Session Commands

- `list-sessions`

  等价与`loginctl`，显示当前的sessions（tty）

  ```
  [root@cyberpelican systemd]# loginctl 
     SESSION        UID USER             SEAT            
          c1         42 gdm              seat0           
           1          0 root             seat0           
          22          0 root                             
  
  3 sessions listed.
  ```

- `session-status [ID...]`

  显示指定 session id 关联的 units

- `show-session [ID...]`

  显示指定 session id 的 properties

- `active [ID...]`

  this brings a session into the foreground if another session is currently in the foreground on the respective seat

- `lock-session [ID...], unlock-session`

  对指定 session id 锁屏

- `terminate-session ID`

  kill 掉 session id 关联的所有进程，等价与关机

- `kill-session ID`

  kill 掉 session id 关联的所有进程，退回到登录界面

### Usesr Commands

- `list-users`

  展示当前登录的用户

- `show-user [USER]`

  显示当前 USER 的关联属性，对比`id`

  ```
  cpl in ~ λ loginctl show-user cpl
  UID=1000
  GID=1000
  Name=cpl
  Timestamp=Wed 2022-03-16 19:56:21 HKT
  TimestampMonotonic=55370159
  RuntimePath=/run/user/1000
  Service=user@1000.service
  Slice=user-1000.slice
  Display=5
  State=active
  Sessions=5
  IdleHint=no
  IdleSinceHint=0
  IdleSinceHintMonotonic=0
  Linger=no
  ```

- `terminate-user [USER]`

  kill 掉 关联 user 的所有进程