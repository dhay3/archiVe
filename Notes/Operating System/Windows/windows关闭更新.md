# windows关闭更新服务

`w+r`输入`services.msc`开启本地服务列表, 找到windows update服务

![Snipaste_2020-08-24_05-32-05](https://github.com/dhay3/image-repo/raw/master/20210518/Snipaste_2020-08-24_05-32-05.4ndtdo7piiy0.png)

将启动类型置为手动,  第一次失败置为无操作

![Snipaste_2020-08-24_05-34-50](https://github.com/dhay3/image-repo/raw/master/20210518/Snipaste_2020-08-24_05-34-50.1fn7ofw8y4u8.png)

![Snipaste_2020-08-24_05-35-09](https://github.com/dhay3/image-repo/raw/master/20210518/Snipaste_2020-08-24_05-35-09.3kj4uv3fu2g0.png)

`win+r`输入`gpedit.msc`开启本地组策略编辑器, 找到windows更新

![Snipaste_2020-08-24_05-39-32](https://github.com/dhay3/image-repo/raw/master/20210518/Snipaste_2020-08-24_05-39-32.7fjgwfu3wlk0.png)

找到配置自动更新置为已禁用

![Snipaste_2020-08-24_05-40-51](https://github.com/dhay3/image-repo/raw/master/20210518/Snipaste_2020-08-24_05-40-51.1wmru3zs9yv4.png)
