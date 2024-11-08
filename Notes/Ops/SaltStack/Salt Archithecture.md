# Salt Architectrue
![Snipaste_2021-08-16_16-42-17](https://github.com/dhay3/image-repo/raw/master/20220223/Snipaste_2021-08-16_16-42-17.rd5cjqezh28.webp)
## Salt Accesses
[https://docs.saltproject.io/en/getstarted/system/index.html](https://docs.saltproject.io/en/getstarted/system/index.html)
**REAL-TIME COMMUNICATION**
Salt 是实时交互的，而不是通过查询数据库来获取信息。但是Salt同样还是需要依赖数据库来查询主机，就好像需要查询每个房间有几个人一样
​

**NO FREELOADERS!**
minion像ARP一样，且master不会对minion操作任何东西
_Communication from the Salt master is a lightweight set of instructions that basically says “if you are a minion with these properties: run this command with these arguments.” Salt minions determine if they match the properties when the command is received.Each Salt minion already has all of the commands that it needs stored locally, so the command can be executed and the results quickly returned back to the Salt master._
_​_

**SALT LOVES TO SCALE**
salt有高性能和高扩展的特性，master 和 minion 之间会通过 ZeroMQ 或 raw TCP 来建立长连接，针对一个 master 可以下挂 10000 台 minion 甚至是 35000 台
​

**NORMALIZE EVERYTHING/MANAGE EVERYTHING**
salt 具有高兼容性，可以运行在 Unix，Windows，MacOS，FreeBSD ，只要可以运行python即可。但是也可以通过 proxy minion，对不能运行python的设备进行管理，例如交换机或路由器。这就意味着只要可以运行网络协议的设备都能接入salt
salt command 会发送给proxy minion，然后会将 salt calls 转为 native protocol 并将其发送给设备，response 会被解析然后返回给master
## Terms 
Overview With Pictrues
[https://docs.saltproject.io/en/getstarted/overview.html](https://docs.saltproject.io/en/getstarted/overview.html)
### Salt master
提供管控，需要运行 salt-master 进程
_Central management system._
_This system is used to send commands and configurations to the Salt minion that is running on managed systems._
_The Salt master/minion model only requires inbound connections into the Salt master. Connections are established from the minion and never from the master._
![Snipaste_2021-08-30_19-28-02](https://github.com/dhay3/image-repo/raw/master/20220223/Snipaste_2021-08-30_19-28-02.22ukygp5mae8.webp)
salt 都是从 minion 向 master 发起的 tcp 链接
```
D:\code\snat_proxy>py sx.py cmd -i 11.181.158.16 -c "ss -npt | head -1 && ss -npt | grep salt"
State      Recv-Q Send-Q        Local Address:Port          Peer Address:Port
ESTAB      0      0              10.150.92.68:10511        11.191.129.59:4506   users:(("salt-minion",3610,77))
ESTAB      0      0              10.150.92.68:58380       11.191.130.188:4505   users:(("salt-minion",3610,40),("salt-minion",31696,40))
ESTAB      0      0              10.150.92.68:18630        11.191.129.59:4505   users:(("salt-minion",3610,70),("salt-minion",31696,70))
SYN-SENT   0      1              10.150.92.68:55433       11.191.133.124:4506   users:(("salt-minion",31696,67))
ESTAB      0      0              11.177.67.54:53399        11.191.129.60:4505   users:(("salt-minion",3610,78),("salt-minion",31696,78))
ESTAB      0      0              11.177.63.50:16642       11.191.130.187:4505   users:(("salt-minion",3610,68),("salt-minion",31696,68))
SYN-SENT   0      1              10.150.92.68:1912         11.191.132.60:4506   users:(("salt-minion",31696,66))
```
master 同时对 minion 提供两种服务
4505 - Event Publisher/Subscriber port (publish jobs/events)
Constant inquiring connection
4506 - Data payloads and minion returns (file services/return data)
Connects only to deliver data
![Snipaste_2021-11-03_15-27-49](https://github.com/dhay3/image-repo/raw/master/20220223/Snipaste_2021-11-03_15-27-49.6w6qkzgdio00.webp)
#### conf
master 默认读取`/etc/salt/master`下的文件，但是也可以使用`-c`来指定配置文件
```
salt-master -h
Usage: salt-master [options]

The Salt Master, used to control the Salt Minions

Options:
  --version             show program's version number and exit
  -V, --versions-report
                        Show program's dependencies version number and exit.
  -h, --help            show this help message and exit
  --saltfile=SALTFILE   Specify the path to a Saltfile. If not passed, one
                        will be searched for in the current working directory.
  -c CONFIG_DIR, --config-dir=CONFIG_DIR
                        Pass in an alternative configuration directory.
                        Default: '/etc/salt'.
```
master 默认会绑定所有可用的interface，然后监听4505和4506端口
```
ss -npt state listening 'sport = :4505' && ss -npt state listening 'sport = :4506'
Recv-Q Send-Q             Local Address:Port               Peer Address:Port
0      1000                           *:4505                          *:*      users:(("salt-master",40610,15))
Recv-Q Send-Q             Local Address:Port               Peer Address:Port
0      1000                           *:4506                          *:*      users:(("salt-master",40925,141))
```
### Salt minion
由 slat master 管控的机器，但是 salt minion 也可以以 stand-alone 模式运行，无需 salt master
_Managed system. This system runs the Salt minion which receives commands and configuration from the Salt master_
![Snipaste_2021-12-08_09-38-21](https://github.com/dhay3/image-repo/raw/master/20220223/Snipaste_2021-12-08_09-38-21.5p4f7s041rs0.webp)
#### conf
minion 默认读取`/etc/slat/minion`下的配置文件，和master一样可以通过`-c`来指定配置文件
```
ps -ef | grep -v grep | grep minion
root      6533     1  0 Jan28 ?        00:40:10  /usr/bin/salt-minion -c /etc/salt -d
root      6535  6533  0 Jan28 ?        00:00:00  /usr/bin/salt-minion -c /etc/salt -d

ls -lh /etc/salt
total 20K
-rw-r----- 1 root root  531 Oct 26  2017 minion
drwxr-xr-x 2 root root 4.0K Aug 14  2017 minion.d
-rw-r--r-- 1 root root   17 Aug 14  2017 minion_id
-rw-r----- 1 root root  476 Oct 23  2017 minion.rpmsave
drwxr-xr-x 3 root root 4.0K Oct 23  2017 pki
```

- minion/minion.d：minion配置
- minion_id：master 通过 minion_id 来识别 minion
## Subsystem
ref:
[https://docs.saltproject.io/en/getstarted/system/plugins.html](https://docs.saltproject.io/en/getstarted/system/plugins.html)
subsystem其实是python module

| Subsystem |  |
| --- | --- |
| Authentication | Authorizes a user before running a job. |
| File server | Distributes files. |
| Secure data store | Makes user-defined variables and other data securely available. |
| State representation | Describes your infrastructure and system configurations. |
| Return formatter | Formats job results into a normalized data structure. |
| Result cache | Sends job results to long-term storage. |
| Remote execution | Runs a wide variety of tasks to install software, distribute files, and other things you need to do to manage systems. |
| Configuration | Configures targeted systems to match a desired state. |

下图用于阐述subsystem
![Snipaste_2022-03-02_12-06-27](https://github.com/dhay3/image-repo/raw/master/20220223/Snipaste_2022-03-02_12-06-27.4eh0qpp0qyw0.webp)
当一个salt任务在运行时会关联多个subsystems
![Snipaste_2022-03-02_12-06-42](https://github.com/dhay3/image-repo/raw/master/20220223/Snipaste_2022-03-02_12-06-42.xbwcjexh0s0.webp)
_At each step, the subsystem delegates its work to the configured plug-in. For example, the job returner plug-in in step 7 might be one of 30 plug-ins including MySQL, redis, or not be configured at all (the job returner plug-in can also run directly on the managed system after step 4)._
_At each step, there are many plug-ins available to perform a task, resulting in hundreds of possible Salt configurations and workflows._
## Communication model
salt 使用 publish-subscribe 模式来通信（生产消费）。链接由minion主动建立，减少了被攻击的风险。master使用 4505 和 4506 端口
![Snipaste_2022-03-02_12-06-53](https://github.com/dhay3/image-repo/raw/master/20220223/Snipaste_2022-03-02_12-06-53.3gk5ye1foim0.webp)
### Publisher
master 向 minion 发送命令
(port 4505) All Salt minions establish a persistent connection to the publisher port where they listen for messages._ Commands are sent asynchronously to all connections over this port_, which enables commands to be executed over large numbers of systems simultaniously.
### Request Server
minion 向 master 回送命令结果
(port 4506) Salt minions connect to the request server as needed to send results to the Salt master, and to securely request files and minion-specific data values (called Salt pillar). Connections to this port are 1:1 between the Salt master and Salt minion (not asynchronous).
### salt minion authentication
当minion第一次启动时，会默认以salt domain name 作为master（但是也可以使用配置指定），如果有找到就会将public key 发送到master（不会存在`/root/.ssh/authorized_keys`），可以使用 salt-key 来查看存储的pk。minion 的 pk 必须被 master 认可，否则minion不能执行任何命令
![Snipaste_2022-03-02_12-07-03](https://github.com/dhay3/image-repo/raw/master/20220223/Snipaste_2022-03-02_12-07-03.1u24a0j3tt0g.webp)
salt 使用RSA非对称加密用于鉴权，AES对称加密用于信息加密传输
