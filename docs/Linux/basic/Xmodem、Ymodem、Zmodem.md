# Xmodem、Ymodem、Zmodem

参考:

https://blog.csdn.net/tianlesoftware/article/details/7746005

文件传输协议：

文件传输是数据交换的主要形式。在进行文件传输时，为使文件能被正确识别和传送，我们需要在两台计算机之间建立统一的传输协议。这个协议包括了文件的识别、传送的起止时间、错误的判断与纠正等内容。

在SecureCRT下的传输协议有ASCII、Xmodem、Ymodem、Zmodem4种。

（1）ASCII：这是最快的传输协议，但只能传送文本文件。 

（2）Xmodem：这种古老的传输协议速度较慢，但由于使用了CRC错误侦测方法，传输的准确率可高达99.6%。 

XModem协议介绍：
XModem是一种在串口通信中广泛使用的异步文件传输协议，分为XModem和1k-XModem协议两种，前者使用128字节的数据块，后者使用1024字节即1k字节的数据块。

（3）Ymodem：这是Xmodem的改良版，使用了1024位区段传送，速度比Xmodem要快。 

（4）Zmodem：Zmodem采用了串流式（streaming）传输方式，传输速度较快，而且还具有自动改变区段大小和断点续传、快速错误侦测等功能。这是目前最流行的文件传输协议。 

options->session options ->Terminal->Xmodem/Zmodem 下

<img src="..\..\imgs\_Linux\Snipaste_2020-09-09_20-19-37.png"/>

然后就可以使用X/Y/Zmodem传输数据了。

Zmodem传输数据会使用到2个命令：

   sz：将选定的文件发送（send）到本地机器

   rz：运行该命令会弹出一个文件选择窗口，从本地选择文件上传到服务器(receive)

 
