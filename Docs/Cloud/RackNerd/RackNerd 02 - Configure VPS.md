---
createTime: 2024-07-09 14:12
tags:
  - "#RackNerd"
  - "#CloudOps"
---

# RackNerd 02 - Configure VPS

## 0x01 Overview

购买 VPS 后有几件事是必须做的

## 0x02 Configure SSH

```
Include /etc/ssh/sshd_config.d/*.conf
Port 65522
PermitRootLogin no
ChallengeResponseAuthentication no
UsePAM yes
X11Forwarding yes
PrintMotd no
AcceptEnv LANG LC_*
Subsystem       sftp    /usr/lib/openssh/sftp-server
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

