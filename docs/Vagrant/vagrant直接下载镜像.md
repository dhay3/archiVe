# vagrant直接下载镜像

转载：

https://my.oschina.net/cxgphper/blog/1940644

一、进入vagrant官网 https://www.vagrantup.com/

二、点击findbox【寻找box】，进入有很多box的列表 https://app.vagrantup.com/boxes/search 

<img src="..\..\imgs\_vagrant\b0690ee78b3a4f1d5e9d566993417fbe730.jpg"/>

三、比如我要下载centos7，那么我点击它，选择进入了它的详情页 https://app.vagrantup.com/centos/boxes/7 

<img src="..\..\imgs\_vagrant\506d56937551ab5bf95f8048f2a38bb2aa1.jpg"/>

四、点击上图进入版本详情页 https://app.vagrantup.com/centos/boxes/7/versions/1804.02

五、拼URL【百度、谷歌各种搜，没一人知道】 'https://app.vagrantup.com/centos/boxes/7/versions/1804.02'+'/providers/'+'供应商名字/'+'.box'

**公式：下载链接 = 产品版本链接 + 供应商英文意思 + 要下载的供应商名称（如virtualbox）+'.box'**

如上：生成的下载链接为：https://app.vagrantup.com/centos/boxes/7/versions/1804.02/providers/virtualbox.box

其实国外很多开源网站都喜欢这么玩，更愿意让学者多动脑想，发人深思。
