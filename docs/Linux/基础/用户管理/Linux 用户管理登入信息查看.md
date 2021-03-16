# 用户管理登入信息查看

- last

  查看最近登入成功的用户，调用wtmp日志

  ```
  root in /dev/pts λ last | head -5
  root     tty7         :0               Fri Mar 12 10:59   still logged in
  reboot   system boot  5.7.0-kali1-amd6 Fri Mar 12 10:57   still running
  root     tty7         :0               Thu Mar 11 16:48 - 18:42  (01:53)
  reboot   system boot  5.7.0-kali1-amd6 Thu Mar 11 16:48 - 18:42  (01:53)
  root     tty7         :0               Thu Mar 11 14:29 - 16:48  (02:18)
  ```

- lastb

  查看最近登入失败的用户，调用btmp日志，查到之后可以通过有效的方式来限制恶意的IP

  ```
  root in ~ λ lastb
  admin    ssh:notty    47.31.16.105     Fri Mar 12 01:23 - 01:23  (00:00)
  admin    ssh:notty    47.31.16.105     Fri Mar 12 01:23 - 01:23  (00:00)
  admin    ssh:notty    61.7.138.31      Thu Mar 11 23:37 - 23:37  (00:00)
  admin    ssh:notty    61.7.138.31      Thu Mar 11 23:37 - 23:37  (00:00)
  admin    ssh:notty    61.7.138.31      Thu Mar 11 23:37 - 23:37  (00:00)
  admin    ssh:notty    187.109.212.218  Thu Mar 11 21:59 - 21:59  (00:00)
  admin    ssh:notty    187.109.212.218  Thu Mar 11 21:59 - 21:59  (00:00)
  admin    ssh:notty    78.190.195.222   Thu Mar 11 18:57 - 18:57  (00:00)
  admin    ssh:notty    103.81.182.190   Thu Mar 11 18:37 - 18:37  (00:00)
  default  ssh:notty    111.94.47.218    Wed Mar  3 15:19 - 15:19  (00:00)
  ```

- w

  查看当前登入的用户，调用utmp日志。这里的pts表示persudotty

  ```
  root in ~ λ w
   13:36:17 up 19:36,  4 users,  load average: 0.00, 0.00, 0.00
  USER     TTY      FROM             LOGIN@   IDLE   JCPU   PCPU WHAT
  root     pts/0    115.233.222.34   12:51   44:28   0.09s  0.09s -zsh
  root     pts/1    115.233.222.34   11:09    2:06m  0.38s  0.38s -zsh
  root     pts/2    115.233.222.34   12:37   37:08   0.28s  0.28s -zsh
  root     pts/3    115.233.222.34   13:33    0.00s  0.18s  0.00s w
  ```

  