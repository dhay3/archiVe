# ICFS关闭共享文件

通过`\\IP`可以访问同一局域网下的另外一台主机的共享文件，`\\IP\共享文件名`可以访问指定的共享文件。’

使用`net share 共享名=磁盘路径`共享文件

- ==远程的一方的权限取安全与共享的交集==

- 如果父级目录共享，在同一父级目录下，所有文件都共享，即使该文件是隐藏共享文件

==windows默认开启磁盘共享==，如果另一方获取到权限就可以远程访问这些共享文件，非常不安全，所以我们需要删除这些共享文件

![Snipaste_2020-08-30_00-15-45](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-08-30_00-15-45.dix4xs3ibao.png)

通过`net share`命令我们可以看到所有的共享文件，包括隐藏共享文件（以`$`结尾）。这里`IPC$`表示空连接,表示可以访问所有资源。

通过`services.msc`将`server`启动类型置为禁用，恢复置为不操作

![Snipaste_2020-08-30_01-07-19](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-08-30_01-07-19.74wu6iab67c0.png)

