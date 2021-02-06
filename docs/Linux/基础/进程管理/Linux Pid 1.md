# Linux Pid 1

参考：

https://vagga.readthedocs.io/en/latest/pid1mode.html

在linux 中pid 为 1 的进程由linux kernel启动。类似的第一个启动的进程的pid就是1（==这个进程负责分配其他的进程的PID，即是其他进程的父进程，由fork()函数调用==）

PID为1的进程与普通的其他进程有如下的区别

1. pid为1的进程不管因为什么原因退出，会对其他所有的进程发送SIGKILL(该信号不可屏蔽)

2. 如果一个有子进程的进程退出，该进程的子进程重新挂到pid为1的进程下(它会收留所有的孤儿进程)
3. 许多信号不会对pid为1的进程生效(例如SIGINT，SIGTERM)。所以Docker中

