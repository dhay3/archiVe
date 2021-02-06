# Metasploit 体系结构

参考：

https://www.fujieace.com/metasploit

https://blog.csdn.net/qq_28437139/article/details/84194578

## 模块

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_21-42-40.png"/>

- Auxiliary / 辅助模块

  扫描、发掘漏洞、嗅探

  主要子模块`gather,scanner,server,dos,admin`

- Exploits / 渗透模块

  利用已发现的漏洞对远程目标系统进行攻击，植入并运行攻击载荷，从而控制目标系统

  主要子模块 `andriod,linux,unix,windows,multi,apple_ios`

- Payload /攻击载荷模块

  在渗透攻击触发漏洞后劫持程序执行流程并跳入这段代码。

  主要子模块`andriod,linux,unix,windows,cmd,apple_ios`

  常用于目标主机与攻击机之间的逆向连接`reverse_tcp`

- Nops /空指令模块

  为了避免攻击载荷在执行的过程中出现随机地址和返回地址错误而在执行shellcode之前假如一些控指令，使得在执行shellcaode时有一个交大的安全着落区

- Encoders /编码器模块

  将攻击载荷进行编码（类似于加密），避免操作系统和杀毒软件辨认出，但是会让载荷的体积变大，这个时候需要选择传输器和传输体配对成的攻击载荷下载目标载荷并运行

- Post /后渗透模块

  当获取目标机的控制权限后，可以获取目标机上的信息，也可以通过目标机继续渗透其他主机。

  主要子模块`andriod,linux,unix,windows,cmd,apple_ios`

## 文件目录

参考：

https://blog.csdn.net/whatday/article/details/82918998

通过`where is metasploit-framework`来查询安装的位置

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_22-30-55.png"/>

document 存放一些具体使用的md文档

moudels 存放主要模块

data 存放辅助文件，如wordlist（简易字典）， exploits（cve漏洞利用代码）

## exploit rank



参考:

https://blog.csdn.net/whatday/article/details/102906448

Rank 常量的值可以是下面的表格中的其中之一，按照可靠性降序排列。

| Ranking              | Description                                                  |
| :------------------- | :----------------------------------------------------------- |
| **ExcellentRanking** | 漏洞利用程序绝对不会使目标服务崩溃，就像 SQL 注入，命令执行，远程文件包含，本地文件包含等等。**除非有特殊情况，典型的内存破坏利用程序不可以被评估为该级别**。（[WMF Escape()](https://github.com/rapid7/metasploit-framework/blob/master/modules/exploits/windows/browser/ms06_001_wmf_setabortproc.rb)） |
| **GreatRanking**     | 该漏洞利用程序有一个默认的目标系统，并且可以自动检测适当的目标系统，或者在目标服务的版本检查之后可以返回到一个特定于应用的返回地址。（译者注：有些二进制的漏洞利用成功后，需要特别设置 Shell 退出后的返回地址，否则当 Shell 结束后，目标服务器会崩溃掉。） |
| **GoodRanking**      | The exploit has a default target and it is the "common case" for this type of software (English, Windows 7 for a desktop app, 2012 for server, etc). 该漏洞利用程序有一个默认目标系统，并且是这种类型软件的“常见情况”（英文，桌面应用程序的Windows 7，服务器的2012等）（译者注：这段翻译的不是很懂，因此保留原文） |
| **NormalRanking**    | 该漏洞利用程序是可靠的，但是依赖于特定的版本，并且不能或者不能可靠地自动检测。 |
| **AverageRanking**   | 该漏洞利用程序不可靠或者难以利用。                           |
| **LowRanking**       | 对于通用的平台而言，该漏洞利用程序几乎不能利用（或者低于 50% 的利用成功率） |
| **ManualRanking**    | 该漏洞利用程序不稳定或者难以利用并且基于拒绝服务（DOS）。如果一个模块只有在用户特别配置该模块的时候才会被用到，否则该模块不会被使用到，那么也可以评为该等级。（例如：[exploit/unix/webapp/php_eval](https://github.com/rapid7/metasploit-framework/blob/master/modules/exploits/unix/webapp/php_eval.rb)） |
