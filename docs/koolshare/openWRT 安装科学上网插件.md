# openWRT 安装科学上网插件

从github下载[离线包](https://github.com/hq450/fancyss_history_package)(lede)

<img src="..\..\..\imgs\_koolshare\Snipaste_2020-10-09_14-54-59.png"/>

> 软件中心可能会屏蔽科学上网工具包的离线安装
>
> 参考:
>
> https://github.com/hq450/fancyss/issues/929
>
> https://github.com/hq450/fancyss

进入服务-->终端输入`sed -i 's/\tdetect_package/\t# detect_package/g' /koolshare/scripts/ks_tar_install.sh`

<img src="..\..\..\imgs\_koolshare\Snipaste_2020-10-09_15-11-47.png"/>

然后刷新后台管理界面，进入酷软-->离线安装，选择安装包，就可以安装插件

<img src="..\..\..\imgs\_koolshare\Snipaste_2020-10-09_15-13-09.png"/>
