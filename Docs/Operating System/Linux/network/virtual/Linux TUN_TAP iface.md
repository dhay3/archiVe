ref
[https://docs.kernel.org/networking/tuntap.html](https://docs.kernel.org/networking/tuntap.html)
## Digest
TUN/TAP seen as simple Point-to-Point or Ethernet device, which instead of receiving packets from pyhsical media, receives them from user space program and instead of sending packets via physiacl media writes them to the user space program
TUN/TAP是点对点通信的，通常被用在用户进程直接通信的。为了使用TUN/TAP，进程必须打开`/dev/net/tun`，设备会以tunXX 或者 tapXX 显示

