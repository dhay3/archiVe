转载：

https://zhuanlan.zhihu.com/p/53139267

1. **FTP服务**

   FTP是文件传输协议（File Transfer Protocol）的简称，该协议属于应用层协议（端口号通常为21），用于Internet上的双向文件传输（即文件的上传和下载）。在网络上有很多服务器提供FTP服务，用来存放大量的文件供他人下载。FTP服务的主要作用是让远程用户可以连接上来，查看服务器上有哪些文件，然后下载，当然用户也可以将本地文件上传到服务器，共享给其他人以便下载。FTP服务提供上传下载服务的同时，管理员也可以设置不同用户对不同文件夹的访问权限（读、写）。

   **在Win7的IIS上搭建FTP服务**

   初学Web开发的人，一般的情况下都认为IIS只能提供Web服务。其实IIS不仅可以提供Web服务，还可以提供其他服务，如网络新闻服务（NNTP）、简单邮件传输服务（SMTP），当然还有文件传输服务（FTP）。只是在IIS上，用得最多的是Web服务。默认安装的情况下，IIS也只会安装Web服务，FTP服务只在特定选择的情况下才会被安装到IIS环境中。下面，我们来看看在Win7的IIS上安装FTP服务的步骤：

   1、打开“控制面板”的“程序和功能”：

   

<img src="https://pic4.zhimg.com/80/v2-7464baa06646a0f52e2cb78051c649a9_720w.jpg"/>

   

   2、进入到“程序和功能”的界面，选择左侧的“打开或关闭Windows功能”，打开“Windows”功能对话框：

   

<img src="https://pic2.zhimg.com/80/v2-e8337196810ae2732633f0fa8ed35445_720w.jpg"/>

   

   3、展开“Internet信息服务”节点，发现默认情况下“FTP服务器”子节点并没有被选中，选中该子节点，点击“确定”按钮，Windows开始安装FTP服务（由于Win7的所有组件安装文件其实已经在Win7的安装过程中被拷贝到系统中，故不想WinXP下安装组件需要系统盘）：

   

<img src="https://pic1.zhimg.com/80/v2-afaa8f58abfe1c66e8641ce74e4c6a39_720w.jpg"/>

   

   4、安装完成后，你会在“服务”管理工具中看到“MIcrosoft FTP Service”服务：

   

<img src="https://pic3.zhimg.com/80/v2-d884975f1dbfc33ed17ba4cd7f9410ed_720w.jpg"/>

   

   **在IIS中添加FTP站点**

   1、在“管理工具”的“Internet信息服务（IIS）管理器”中，选中服务器，点击右键菜单中的“添加FTP站点…”子菜单项：

   

<img src="https://pic3.zhimg.com/80/v2-6c6202fe76b706f1e081fb58db42f61d_720w.jpg"/>

   

   2、在“添加FTP站点”对话框的“FTP站点名称”中输入你所期望的站点名称，并选择你期望用来存放文件的“内容目录”：

   

<img src="https://pic3.zhimg.com/80/v2-f5af1c572d35782111f0c4fd32ffdf40_720w.jpg"/>

   

   3、点击“下一步”按钮，进入IP地址绑定和SSL设置，由于我们这里不使用SSL，所以SSL选项选择“无”，至于什么是“SSL”，本人将在后续的文章中讲解：

   

<img src="https://pic3.zhimg.com/80/v2-65f7c2f5363f65c780854bfd9d715141_720w.jpg"/>

   

   4、继续点击“下一步”，进入到“身份验证和授权信息”环节，“身份验证”选择“匿名”和“基本”，“授权”的“允许访问”选择“所有用户”，“权限”选择“读取”：

   

<img src="https://pic2.zhimg.com/80/v2-982ae6bfd0643a482b8cab458d25000b_720w.jpg"/>

   

   5、点击“完成”按钮，完成FTP站点的添加过程。这时会发现IIS管理器“网站”节点下多了一项刚才添加的FTP站点：

   

<img src="https://pic1.zhimg.com/80/v2-dc724ae92cc360a3f5ff9af4cd732997_720w.jpg"/>

   

   **FTP站点的配置及授权**

   1、选中要配置的FTP站点，点击“内容视图”，发现内容为空：

   

<img src="https://pic4.zhimg.com/80/v2-a18151accac69bc6a1a337fdc121bcc7_720w.jpg"/>

   

   这是因为到目前为止，我们还没有在与FTP关联的文件夹中添加任何内容，在Windows资源管理中打开对应的文件夹，并添加几个测试目录：

   

<img src="https://pic2.zhimg.com/80/v2-a4c7a53df221c559b0c6b205950f3fb7_720w.jpg"/>

   

   在IIS的中刷新“内容视图”，这时我们看到，新增的文件夹出现在“内容视图”中了：

   

<img src="https://picb.zhimg.com/80/v2-af7bbe5cf6446002188a396bb1e63370_720w.jpg"/>

   

   2、我们打开Windows资源管理器或者浏览器，在地址栏中输入ftp://127.0.0.1（如果是远程访问，请输入服务器的IP地址），我们可以看到FTP服务器上的目录：

   

<img src="https://pic1.zhimg.com/80/v2-0422e0b63736d6b7ee00375bb519faab_720w.jpg"/>

   

   3、双击“开发文档”文件夹，进入该文件，我们视图将本地文件复制到该文件夹中时，因为权限不够（之前只设定“读取”的权限），系统提示错误：

   

<img src="https://pic3.zhimg.com/80/v2-efbebb24a0d4d424eed6d6c5c65f1731_720w.jpg"/>

   

   4、这时便需要我们来设置FTP站点的权限。为了方便，我们在Windows系统中添加一个名为“FTPUser”的用户，在“开始菜单”中选择“计算机”，点击右键菜单中的“管理”子菜单：

   

<img src="https://pic4.zhimg.com/80/v2-26cf061b89274f00da7b82c781f373aa_720w.jpg"/>

   

   打开“本地用户和组”的“用户”节点，

   

<img src="https://picb.zhimg.com/80/v2-4fd6fdf2d1880061cfa23004d4bec0b2_720w.jpg"/>

   

   添加用户名为“ftpuser”的新用户：

   

<img src="https://pic2.zhimg.com/80/v2-8fd73ad2abd8657c40c83ff16a5d1020_720w.jpg"/>

   

   点击“创建”按钮为Windows系统创建新的用户。

   3、在IIS管理器的FTP站点中，选中你要授权的文件夹，并切换到“功能视图”，选中“FTP授权规则”，

   

<img src="https://pic1.zhimg.com/80/v2-8808fcebc8cab3ffb8705342a25bd54c_720w.jpg"/>

   

   双击“FTP授权规则”，进入“授权规则”管理界面，点击右键菜单的“添加允许规则”，弹出对话框，选择“指定的用户”并输入“ftpuser”，设置其权限为“读取”和“写入”：

   

<img src="https://pic3.zhimg.com/80/v2-d243c1d1bef3aadeca769000d966533f_720w.jpg"/>

   

   点击“确定”完成“授权规则”的添加。

   4、回到Windows资源管理器，进入“ftp://127.0.0.1/开发文档”文件夹，点击右键菜单的“登录”子菜单，弹出“登录身份”对话框，输入用户名ftpuser和对应的密码，点击“登录”按钮登录ftp的文件夹：

   

<img src="https://picb.zhimg.com/80/v2-23f84e69a6338ad8ee39776cfb64bf04_720w.jpg"/>

   

   5、这时，在试图将文件或文件夹拷贝到ftp目录中，依然弹出“权限不足”的错误提示，这是为什么呢？原来ftp的权限是在Windows用户权限的基础上的，所以我们要在资源管理器中，为ftp对应的文件夹为特定的用户添加对应的权限。在资源管理器中，选定相关的文件夹，点击右键菜单中的“属性”子菜单，弹出“属性”对话框，却换到“安全”tab页：

   

<img src="https://pic3.zhimg.com/80/v2-227bd604d9cfe793ad141065aa7556dd_720w.jpg"/>

   

   点击“编辑”按钮，弹出权限编辑对话框，输入ftpuser，并“检查名称”：

   

<img src="https://pic1.zhimg.com/80/v2-681de018f951dbba2c737f04b2412098_720w.jpg"/>

   

   点击“确定”按钮，ftpuser被添加到用户列表中，在“ftpuser的权限”列表中，选中“修改”权限，点击“确定”按钮，完成Windows文件夹授权：

   

<img src="https://pic1.zhimg.com/80/v2-cb47fd7ac3762be28669b41d1565686f_720w.jpg"/>

   

   这时，我们再次使用ftpuser登录到“ftp://127.0.0.1/开发文档”中，就可以完成“新建文件夹”及将文件拷贝到该文件夹中的操作了。

   

<img src="https://picb.zhimg.com/80/v2-b2d53e857e1a7285d49bd3574d644c54_720w.jpg"/>
