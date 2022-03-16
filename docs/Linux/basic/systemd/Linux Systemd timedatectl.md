# Linux Systemd timedatectl

## Digest

syntax：`timedatectl [options] {command}`

timedatectl may be used to query and change the system clock and its settings, and enable or disable time synchronization services

## Postional args

- `status`

  显示当前的 system clock and RTC(a real-time clock , take a glance in wiki)

  ```
  cpl in /sharing/conf λ timedatectl       
                 Local time: Wed 2022-03-16 21:17:28 HKT
             Universal time: Wed 2022-03-16 13:17:28 UTC
                   RTC time: Wed 2022-03-16 21:16:52
                  Time zone: Asia/Hong_Kong (HKT, +0800)
  System clock synchronized: yes
                NTP service: active
            RTC in local TZ: yes
  ```

- `set-time [TIME]`

  设置系统的时间，TIME 的格式采用`2012-10-30 18:17:16`

- `set-timezone [TIMEZONE]`

  设置timezone

- `list-timezones`

  展示所有可用的timezone

- `set-ntp [BOOL]`

  布尔值，ntp 是否启用

### ntp

- `timesync-status`

  显示ntp的一些信息

  ```
  pl in /sharing/conf λ timedatectl timesync-status          
         Server: 139.199.215.251 (2.manjaro.pool.ntp.org)
  Poll interval: 34min 8s (min: 32s; max 34min 8s)
           Leap: normal
        Version: 4
        Stratum: 2
      Reference: 647A24C4
      Precision: 1us (-23)
  Root distance: 35.178ms (max: 5s)
         Offset: +15.245ms
          Delay: 34.410ms
         Jitter: 6.204ms
   Packet count: 7
  ```

- `ntp-servers INTERFACE SERVER`

     指定 iface 使用的 ntp server