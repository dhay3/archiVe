# Windows修改字体

1. 按Win+R，然后在运行中输入regedit，打开注册表编辑器，如图所示：

   ![win10系统默认字体怎么改？教你修改win10默认字体的操作方法](http://www.kkx.net/uploadfile/2019/1207/20191207044758318.jpg)

2. 在注册表左侧依次定位到：HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Fonts；

![Win10怎么改字体？Win10改系统默认字体二个方法](http://www.kkx.net/uploadfile/2019/1207/20191207044122921.jpg)

3. 在右边找到Microsoft YaHei & Microsoft YaHei UI (TrueType)把后面的msyh.ttc改为你想要的字体文件名（右键需要改的字体即可看到名字），左键双击打开，把“数值数据”更换成前面复制的值，重启电脑，系统默认字体就更换好了。字体的话大家可以通过打开C:\Windows\Fonts 进行查看！

![Win10怎么改字体？Win10改系统默认字体二个方法](http://www.kkx.net/uploadfile/2019/1207/20191207044136868.jpg)

修改完成后重启一次win10系统即可看到效果，当然如果一些不是很完善的字体可能会导致一些文字显示异常！