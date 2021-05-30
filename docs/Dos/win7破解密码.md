# 5shift破解win7密码

在开机输入密码时，连续连按5次shift，出现如下界面

==通过这个图标找到对应系统盘中的文件，将其改名，然后将cmd.exe复制一份，将其命名为该文件==

> 由于管理员的运行的文件在C:\Windows\System32\cmd.exe和该图标对应的文件在同一个目录下

<img src="..\..\imgs\_Dos\Snipaste_2020-08-29_18-52-08.png" style="zoom:80%;" />

然后重启，在出现windows图标时断电，进入修复模式

<img src="..\..\imgs\_Dos\Snipaste_2020-08-29_18-54-13.png" style="zoom:80%;" />

进入界面后选择取消还原计算机

<img src="..\..\imgs\_Dos\Snipaste_2020-08-29_18-55-34.png" style="zoom:80%;" />

等待系统尝试修复，当出现如下界面，选择查看问题详细信息

<img src="..\..\imgs\_Dos\Snipaste_2020-08-29_19-17-02.png" style="zoom:80%;" />

点击最后一个超链接

<img src="..\..\imgs\_Dos\Snipaste_2020-08-29_19-18-16.png" style="zoom:80%;" />

进入记事本点击打开，找到图标对应的文件，将其重命名，并将cmd.exe复制一份重命名

<img src="..\..\imgs\_Dos\Snipaste_2020-08-29_19-27-37.png" style="zoom:80%;" />

然后正常重启计算机，同样以5shift弹出窗口，就会发现原来的界面变成cmd窗口了

<img src="..\..\imgs\_Dos\Snipaste_2020-08-29_19-32-23.png" style="zoom:80%;" />

`net user John ""`

or

`net user hacker 123 /add && net localgroup administrators hacker /add `

如果只是将当前用户密码置为空，直接登入即可。进入后可以切换用户，记得删除创建的用户

`net user hacker /del` 



