

# Zabbix host

## 概述

> A "host" in Zabbix is quite flexible. It can be a physical server, a network switch, a virtual machine or some application. ==All access permissions are assigned to host groups, not individual hosts. That is why a host must belong to at least one group.==

- 当前host没有item被监控

<img src="https://www.zabbix.com/documentation/5.0/_media/manual/quickstart/icon_zbx_grey.png"/>

- 当前host有item被监控

<img src="https://www.zabbix.com/documentation/5.0/_media/manual/quickstart/icon_zbx_green.png"/>

- 当前host不可用

<img src="https://www.zabbix.com/documentation/5.0/_media/manual/quickstart/icon_zbx_red.png"/>

## Hosts/Host

*a networked device that you want to monitor, with IP/DNS.*

==注意，单独主机无任何权限，需要指定host group==

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-06_11-12-32.png"/>

- host name：主机名字
- visible host：显示在host dashboard中的名字
- groups：指定host groups，即权限
- interfaces：指定套接字或DNS
- monitored by proxy：是否有proxy 代理
- enabled：是否启用当前host

## Host/Items

*a particular piece of data that you want to receive off of a host, a metric of data.*

新添加的主机无任何监控项，需要添加item

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-06_10-59-19.png"/>

- name：item的名字

- key：指定监控的value

- type of information：数据类型

- units：数据类型的单位

- updates interval：数据更新的间隔

- custom intervals：自定义数据更新interval

- history storage period：历史记录（updates后）保存的时间

- trend storage period：trend保存的时间

  趋势数据就是每个小时收集到的历史数据中的最大值、最小值，平均值以及每个小时收集到的历史数据的数量，所以，趋势数据每小时收集一次，数据量不会特别大，一般情况下，历史数据的保留时间都比趋势数据的保留时间短很多，因为历史数据比较多，如果我们监控的主机非常多，而且监控的频率特别频繁，那么数据库的压力则会变得非常大。

- show value：改变数据展示的方式

- new application：用来标识当前item属于哪一个应用集，可以在该栏中添加新的应用集

- application：选中已有的应用集

- populates host inventory field：

配置完成后我们可以使用`zabbix-get`在服务端来校验

> 注意需要开启客户端的agent（默认10050）端口
>
> -s 指定host，-k 指定item的key

```
[root@chz zabbix]# zabbix_get -s 192.168.80.100 -k system.cpu.switches
4889961
```

**解决zabbix_get [10158]: Check access restrictions in Zabbix agent configuration**

需要在客户端指定服务端的IP

```
[root@chz ~]# vim /etc/zabbix/zabbix_agentd.conf 

...
Server=server-host-ip
...

[root@chz ~]# systemctl restart zabbix-agent
```

### Example

添加根分区磁盘使用率监控项

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-07_10-45-56.png"/>

> `vfs.fs.size[fs.mode]`使用具体参考官方文档https://www.zabbix.com/documentation/5.0/manual/config/items/itemtypes/zabbix_agent
>
> ==添加item后需要重启zabbix-agent==

校验

```
zabbix-server
[root@chz ~]# zabbix_get -s 192.168.80.100 -k vfs.fs.size[/,pused]
31.197989

zabbix-agnet
[root@chz ~]# df -h
Filesystem               Size  Used Avail Use% Mounted on
devtmpfs                 470M     0  470M   0% /dev
tmpfs                    487M     0  487M   0% /dev/shm
tmpfs                    487M   15M  472M   3% /run
tmpfs                    487M     0  487M   0% /sys/fs/cgroup
/dev/mapper/centos-root   17G  5.3G   12G  32% /
/dev/sda1               1014M  213M  802M  21% /boot
/dev/sr0                 4.4G  4.4G     0 100% /run/media/root/CentOS 7 x86_64
tmpfs                     98M   28K   98M   1% /run/user/0

```

在mointoring sidebar可以查看图形化指标

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-07_10-57-06.png"/>

## Hosts/triggers

> trigger被触发处于problem，如果没有被触发处于ok

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-07_10-06-44.png"/>

- Name：显示在列表中的trigger的名字

- operational data：

- serverity：该trigger的严重性，颜色越深越严重

- expression：该trigger表达式（到达该阈值触发trigger）

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-07_10-12-24.png"/>

  - item：该host配置的监控项，该trigger是定义哪个监控项的阈值
  - function：该trigger需要调用的函数

  > time的单位为sec

  - last of：T value的值，会根据function改变
  - time shift：如果T value的值是时间，需要在该项中指明时长
  - result：N的值

  ```
  {zb_01:system.cpu.switches.avg(10)}=300
  ```

  上图所示表达式表示，如果10s内system.cpu.switches的平均值等于300就触发trigger

- expression constructor：可以构造表达式之间的逻辑关系

- ok event generation：trigger处于ok的表达式
- problem event generation mode：如果选择multiple，如果trigger处于problem，则重复生成对应的problem event
- ok event closes：如果当前trigger处于ok，关闭problem
- allow manual close：是否允许手动关闭
- url：与trigger关联的url

### expression

参考：https://www.zabbix.com/documentation/5.0/manual/config/triggers/expression

pattern：`{hostname:item.function(T)}>N`

```
{zb_01:vfs.fs.size[/,pused].last(#2)}>5
```

- hostname：该trigger用于哪一个host，==注意是hostname，而不是visible name==

- item：该trigger用于哪一个item
- function：调用的函数

> 注意：如果参数前面带有”#“做为前缀，则表示次数，反之则表示时间。expression支持加减乘除运算

### Example 

当`vfs.fs.size`监控项的值大于5%时触发trigger

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-07_11-16-52.png"/>

检查trigger

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-07_11-19-48.png"/>

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-07_11-21-36.png"/>

<img src="..\..\imgs\_Zabbix\Snipaste_2020-11-07_11-23-13.png"/>

