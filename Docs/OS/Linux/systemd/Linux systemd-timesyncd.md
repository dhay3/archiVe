# Linux systemd-timesyncd

## Digest

> 有些OS也使用 ntpd 或 chrony 来同步

systemd-timesyncd is a system service that may be used to synchronize the local system clock with a remote network time protocol (NTP) server

ntp server 由`/etc/systemd/timesyncd.conf`决定（具体查看`timesyncd.conf(5)`），可以使用`timedatectl set-ntp`来决定是否使用ntp

```
cpl in /usr/lib/systemd λ systemctl status systemd-timesyncd.service 
● systemd-timesyncd.service - Network Time Synchronization
     Loaded: loaded (/usr/lib/systemd/system/systemd-timesyncd.service; enabled; vendor preset: enabled)
     Active: active (running) since Wed 2022-03-16 20:47:49 HKT; 1h 0min ago
       Docs: man:systemd-timesyncd.service(8)
   Main PID: 607 (systemd-timesyn)
     Status: "Initial synchronization to time server 139.199.215.251:123 (2.manjaro.pool.ntp.org)."
      Tasks: 2 (limit: 16595)
     Memory: 2.3M
        CPU: 65ms
     CGroup: /system.slice/systemd-timesyncd.service
             └─607 /usr/lib/systemd/systemd-timesyncd
```

## conf

Initially the main configuration file in `/etc/systemd` contains commented out entries showing the defaults as a guid to the admin

当配置修改后可以使用`systemctl restart systemd-timesyncd`来让配置生效

```
[Time]
#NTP=
#FallbackNTP=0.manjaro.pool.ntp.org 1.manjaro.pool.ntp.org 2.manjaro.pool.ntp.org 3.manjaro.pool.ntp.org
#RootDistanceMaxSec=5
#PollIntervalMinSec=32
#PollIntervalMaxSec=2048
```

- `NTP`

  a space-speraated list of NTP server host names or IP addresses

- `FallbackNTP`

  a space-spearated list of NTP server host names or IP addresses to be used as the fallback NTP servers

- `RootDistanceMaxSec`

  maximum acceptable root distance

- `PollIntervalMinSec, PollIntervalMaxSec`

  the minimum and maximum poll intervals for NTP messages

- `ConnectionRetrySec`

  Specifies the minimum delay before subesequent attemps to contact a new NTP server are made

```
[Time]
NTP=0.cn.pool.ntp.org 1.cn.pool.ntp.org 2.cn.pool.ntp.org 3.cn.pool.ntp.org
FallbackNTP=0.manjaro.pool.ntp.org 1.manjaro.pool.ntp.org 2.manjaro.pool.ntp.org 3.manjaro.pool.ntp.org
RootDistanceMaxSec=3
PollIntervalMinSec=30
PollIntervalMaxSec=2048
```

