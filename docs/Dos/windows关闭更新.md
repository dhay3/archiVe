# windows关闭更新服务

`w+r`输入`services.msc`开启本地服务列表, 找到windows update服务

<img src="..\..\imgs\_Dos\Snipaste_2020-08-24_05-32-05.png" style="zoom:50%;" />

将启动类型置为手动,  第一次失败置为无操作

<img src="..\..\imgs\_Dos\Snipaste_2020-08-24_05-34-50.png" style="zoom:50%;" />

<img src="..\..\\img\Snipaste_2020-08-24_05-35-09.png" style="zoom:50%;" />

`win+r`输入`gpedit.msc`开启本地组策略编辑器, 找到windows更新

<img src="..\..\imgs\_Dos\Snipaste_2020-08-24_05-39-32.png" style="zoom:50%;" />

找到配置自动更新置为已禁用

<img src="..\..\imgs\_Dos\Snipaste_2020-08-24_05-40-51.png" style="zoom:50%;" />
