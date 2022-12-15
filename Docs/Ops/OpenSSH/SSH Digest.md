# SSH Digest

ref

https://en.wikipedia.org/wiki/Secure_Shell

https://www.openssh.com/

> 本系列文章以 OpenSSH 为基础

Secure Shell protocol ( SSH ) 是一个加密的网络协议，和 HTTP 一样都是以 C/S 模式工作

SSH 支持如下功能

- For login to a shell on a remote host (replacing [Telnet](https://en.wikipedia.org/wiki/Telnet) and [rlogin](https://en.wikipedia.org/wiki/Rlogin))
- For executing a single command on a remote host (replacing [rsh](https://en.wikipedia.org/wiki/Remote_shell))
- For setting up automatic (passwordless) login to a remote server (for example, using [OpenSSH](https://en.wikipedia.org/wiki/OpenSSH)[[26\]](https://en.wikipedia.org/wiki/Secure_Shell#cite_note-26))
- In combination with [rsync](https://en.wikipedia.org/wiki/Rsync) to back up, copy and mirror files efficiently and securely
- For [forwarding](https://en.wikipedia.org/wiki/Port_forwarding) a port
- For [tunneling](https://en.wikipedia.org/wiki/Tunneling_protocol) (not to be confused with a [VPN](https://en.wikipedia.org/wiki/VPN), which [routes](https://en.wikipedia.org/wiki/VPN#Routing) packets between different networks, or [bridges](https://en.wikipedia.org/wiki/VPN#OSI_Layer_1_services) two [broadcast domains](https://en.wikipedia.org/wiki/Broadcast_domain) into one).
- For using as a full-fledged encrypted VPN. Note that only [OpenSSH](https://en.wikipedia.org/wiki/OpenSSH) server and client supports this feature.
- For forwarding [X](https://en.wikipedia.org/wiki/X_Window_System) from a remote [host](https://en.wikipedia.org/wiki/Host_(network)) (possible through multiple intermediate hosts)
- For browsing the web through an encrypted proxy connection with SSH clients that support the [SOCKS protocol](https://en.wikipedia.org/wiki/SOCKS).
- For securely mounting a directory on a remote server as a [filesystem](https://en.wikipedia.org/wiki/File_system) on a local computer using [SSHFS](https://en.wikipedia.org/wiki/SSHFS).
- For automated remote monitoring and management of servers through one or more of the mechanisms discussed above.
- For development on a mobile or embedded device that supports SSH.
- For securing file transfer protocols.

SSH 支持两种变体协议 SSH1 和 SSH2 两者互不兼容 。SSH 1 在 OpenSSH version 7.6 后被移除 