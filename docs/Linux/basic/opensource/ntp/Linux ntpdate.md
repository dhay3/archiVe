# Linux ntpdate

参考：

https://linux.die.net/man/8/ntpdate

syntax：`ntpdate [options] <server>`

ntpdate用于从ntp server上同步时间信息，可以手动设置也可以在boot阶段由ntpd设置，==ntpdate会拒绝设置时间，如果ntpd在同一个主机上==



