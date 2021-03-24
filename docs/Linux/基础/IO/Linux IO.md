# Linux I/O模式

参考：

https://blog.csdn.net/jyxmust/article/details/88354485

https://www.jianshu.com/p/486b0965c296

https://zhuanlan.zhihu.com/p/96391501

https://www.junmajinlong.com/os/kernelBuffer_ioBuffer/

<img src="..\..\..\..\imgs\_Linux\Snipaste_2021-03-15_14-02-13.png"/>

## bufferd I/O

缓存IO又叫做标准IO。大多数的文件系统都是缓存IO。

==数据先会被拷贝到page cache中==。

调用`read()`时，如果page cache(操作系统页缓存)有数据就读取出数据并直接返回给应用程序。如果没有就从磁盘读取数据拷贝到page cache

调用`write()`时，数据将应用程序的地址空间拷贝到page cache，然后根据linux的延迟写机制(可以使用`sync`命令直接将数据刷入到硬盘)，定期将page cache刷到磁盘上

**优点**

1. 使用page cache保护了磁盘
2. 减少IO读盘次数，提高了IO读写速度

**缺点**

1. 由于page cache处于内核空间，不能直接被应用程序直接寻址。读操作还需要将页缓存数据拷贝到内存中对应的用户空间。这样需要两次数据拷贝。写操作也是一样的。这样会花费大量的内存。

但并不会因此降低性能。最常见的一个优化时预读，它表示在读数据时，会比所请求要读取的数据量多读一点放入到page cache。

**mmap**

因此，Buffer I/O 中引入一类特别的操作叫做内存映射文件（mmap）。

使用mmap函数的时候，会在当前进程的用户地址空间中开辟一块内存，==这块内存与系统的文件进行映射==。对其的读取和写入，会转化为对相应文件的操作。 并且，在进程退出的时候，会将变化的内容（脏页）自动回写到对应的文件里面。

## direct I/O

直接IO，使用O_DIRECT标记。数据直接在用户地址空间的缓冲区和磁盘直接传输，中间少了page cache。

## sync I/O

同步IO和Java中同步相同，发起IO请求后阻塞(blocking)直到完成。buffered I/O和 direct I/O都术语同步IO

## Async I/O

异步IO和Java中的异步形同，发起IO请求不阻塞。通常用内核的Libaio提供





















