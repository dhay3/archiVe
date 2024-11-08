# Linux Systemd analyze

## digest

用于分析和debug系统，通常用于 boot-up performance debugging，如果没有指定任何positional

arg 默认使用`systemd-analyze time`

## positional args

- `time`

  默认项，查看boot花费的时间

  ```
  cpl in ~/note/docs/Linux/basic/systemd on master λ systemd-analyze     
  Startup finished in 5.131s (firmware) + 4.059s (loader) + 701ms (kernel) + 3.285s (userspace) = 13.178s 
  graphical.target reached after 2.294s in userspace
  ```

- `blame`

  this command prints a list of all running units, orderd by the time they took to initialize

- `critical-chain [unit]`

  以树形显示 time-critical chain for units

  ```
  [root@cyberpelican ~]# systemd-analyze critical-chain httpd.service
  The time after the unit is active or started is printed after the "@" character.
  The time the unit takes to start is printed after the "+" character.
  
  httpd.service +25.646s
  └─remote-fs.target @8.921s
    └─remote-fs-pre.target @8.919s
      └─iscsi-shutdown.service @8.855s +57ms
        └─network.target @8.835s
          └─wpa_supplicant.service @13.631s +67ms
            └─basic.target @4.799s
              └─sockets.target @4.798s
                └─dbus.socket @4.798s
                  └─sysinit.target @4.784s
                    └─sys-fs-fuse-connections.mount @30.645s +24ms
                      └─system.slice
                        └─-.slice
  ```

- `verify FILE`

  校验unit file 是否争取

  ```
  $ cat ./user.slice
  [Unit]
  WhatIsThis=11
  Documentation=man:nosuchfile(1)
  Requires=different.service
  
  [Service]
  Description=x
  
  $ systemd-analyze verify ./user.slice
  [./user.slice:9] Unknown lvalue 'WhatIsThis' in section 'Unit'
  [./user.slice:13] Unknown section 'Service'. Ignoring.
  Error: org.freedesktop.systemd1.LoadFailed:
  Unit different.service failed to load:
  No such file or directory.
  Failed to create user.slice/start: Invalid argument
  user.slice: man nosuchfile(1) command failed with code 16
  ```

  