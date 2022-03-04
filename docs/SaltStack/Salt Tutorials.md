ref:
[https://docs.saltproject.io/en/getstarted/fundamentals/index.html](https://docs.saltproject.io/en/getstarted/fundamentals/index.html)
[https://docs.saltproject.io/en/getstarted/system/index.html](https://docs.saltproject.io/en/getstarted/system/index.html)
[https://docs.saltproject.io/en/getstarted/config/index.html](https://docs.saltproject.io/en/getstarted/config/index.html)
[https://docs.saltproject.io/salt/user-guide/en/latest/](https://docs.saltproject.io/salt/user-guide/en/latest/)
[https://docs.saltproject.io/en/getstarted/](https://docs.saltproject.io/en/getstarted/)
[https://docs.saltproject.io/en/latest/](https://docs.saltproject.io/en/latest/)


documentation -> tutorials
## Install
参考
[https://docs.saltproject.io/en/getstarted/fundamentals/install.html](https://docs.saltproject.io/en/getstarted/fundamentals/install.html)
![Snipaste_2021-08-16_16-42-17](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220223/Snipaste_2021-08-16_16-42-17.43ff85s9osy0.webp)
vagrant 镜像启动后，会启动 3 台机器，1 台 master，2 台 minion
```
D:\salt\salt-vagrant-demo-master\salt-vagrant-demo-master>echo %cd%
D:\salt\salt-vagrant-demo-master\salt-vagrant-demo-master

#password vagrant
vagrant ssh master
sudo su
```
## Accept Connections
在master上可以通过`salt-key --list-all`来查看当前接入在master上的minion id
```
root@saltmaster:/home/vagrant# salt-key --list-all
Accepted Keys:
minion1
minion2
Denied Keys:
Unaccepted Keys:
Rejected Keys:
```
可以通过`salt-key --accept=<key>`或`salt-key --accept-all`将minion public key 接入到 master
## Command execution
在master上通过execution module 来向 minion 发送命令，需要读取master的配置，如果不是默认配置可能会报错，使用`-c`来指定配置文件的路径
```
root@saltmaster:/home/vagrant# salt * test.ping
minion1:
    True
minion2:
    True
root@saltmaster:/home/vagrant# salt * cmd.run hostname
minion2:
    minion2
minion1:
    minion1

D:\code\snat_proxy>py sx.py cmd -i 11.191.130.187 -c "salt -c /home/admin/netops-channel/conf 'LVS-APP-G8-1.ET2' cmd.run 'ip r'"
LVS-APP-G8-1.ET2:
    10.86.10.64/30 dev t1  proto kernel  scope link  src 10.86.10.66
    10.86.14.64/30 dev t2  proto kernel  scope link  src 10.86.14.66
    blackhole 10.198.248.0/24  proto zebra
    default via 10.86.10.65 dev t1  proto zebra  metric 10001

```
如果错误会返回
```
root@saltmaster:/home/vagrant# salt * test.ping
minion1:
    Minion did not return. [No response]
    The minions may not have all finished running and any remaining minions will return upon completion. To look up the return data for this job later, run the following command:
    
    salt-run jobs.lookup_jid 20220223021622638639
minion2:
    Minion did not return. [No response]
    The minions may not have all finished running and any remaining minions will return upon completion. To look up the return data for this job later, run the following command:
    
    salt-run jobs.lookup_jid 20220223021622638639
```
例如执行
`salt '*' test.rand_sleep 120`

1. master 通过 publisher port 下发任务，4505
1. 每个minion都会校验自己是不是执行改命令的主体目标
1. 匹配的minion会执行命令，并将reponse 返回给master

salt 在执行命令时会通过execution model，拆分成两个 worker thread
minion在收到命令后会去找 test module（source code in slat/moudles），调用rand_sleep function，提供 120 作为参数
![Snipaste_2021-08-30_19-28-02](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220223/Snipaste_2021-08-30_19-28-02.29mrtnf0tvok.webp)
### Target
[https://docs.saltproject.io/en/getstarted/fundamentals/targeting.html](https://docs.saltproject.io/en/getstarted/fundamentals/targeting.html)
具体可以参考man page target options
target指的是调用command的minion，一般就是 minion id
#### globbing
target支持wildcard通配符，regexp，list
```
salt -E 'minion[0-9]' test.ping

salt -L 'minion1,minion2' test.ping

salt -C 'G@os:Ubuntu and minion* or S@192.168.50.*' test.ping
```
## Grains
[https://docs.saltproject.io/en/latest/topics/grains/index.html](https://docs.saltproject.io/en/latest/topics/grains/index.html)
Grains 是 salt 的一个接口用于获取底层 OS 信息，通常包含 OS type，memory，IP address，kernel versoin 等信息，相对是静止
可以使用`salt '*' grains.ls`查看所有可以使用的grains
```
root@saltmaster:/srv/pillar# salt 'minion1' grains.ls
minion1:
    - biosreleasedate
    - biosversion
    - cpu_flags
    - cpu_model
    - cpuarch
    ...
```
## State
> A reusable declaration that configures a specific part of a system. Each state is defined using a state declaration.

_Remote execution is a big time saver, but it has some shortcomings. Most tasks you perform are a combination of many commands, tests, and operations, each with their own nuances and points-of-failure. Often an attempt is made to combine all of these steps into a central shell script, but these quickly get unwieldy and introduce their own headaches._
_To solve this, SaltStack configuration management lets you create a re-usable configuration template, called a state, that describes everything required to put a system component or application into a known configuration._
### State File
> A file with an SLS extension that contains one or more state declarations.

State file 通常使用 YAML 文件来描述，多组或一组 command 组成，可以看成docker 中的 service
![Snipaste_2021-11-03_15-27-49](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220223/Snipaste_2021-11-03_15-27-49.1mgvq2irtxeo.webp)
可以在宿主机vagrant的镜像文件中编辑(`/salt-vagrant-demo-master/slatstack/salt`)，或者在镜像的`/srv/salt`中编辑
```
install_network_packages:
  pkg.installed:
    - pkgs:
      - rsync
      - lftp
      - curl
```
将上述内容以`nettools.sls`保存在`/srv/salt`中
`install_network_packages`表示 state declaration 也为ID
`pkg.installed`表示调用的State Function(调用state 中的pkg 中的installed 函数，具体可以看source code)
具体可以调用的state function，start with `salt.states.*`
[https://docs.saltproject.io/en/latest/py-modindex.html](https://docs.saltproject.io/en/latest/py-modindex.html)
`pkgs`表示参数，如果参数是list使用yaml list的格式，如果参数是单个使用yaml单个的格式
​

在master使 state file 生效`salt minion1 state.apply nettools`
nettools 对应state file 文件名，state.apply 为调用state module 中的apply 函数
```
root@saltmaster:/srv/salt# salt minion1 state.apply nettools

minion1:
----------
          ID: install_network_packages
    Function: pkg.installed
      Result: True
     Comment: The following packages were installed/updated: lftp
              The following packages were already installed: rsync, curl
     Started: 02:42:33.770032
    Duration: 44064.356 ms
     Changes:   
              ----------
              lftp:
                  ----------
                  new:
                      4.8.4-2build3
                  old:

Summary for minion1
------------
Succeeded: 1 (changed=1)
Failed:    0
------------
Total states run:     1
Total run time:  44.064 s
```
如果指定apply的minion数量很多，可以使用`--batch-size`来指定一次执行多少台
```
salt --batch-size 10 '*' state.apply
```
### TOP file
top file(top.sls) 会被应用于所有的minion，被用于将多个state files 应用于minions
例如：
```
base:
	'*':
  	- vim
    - scripts
  '*web*':
  	- apache
    - python
  '*db*':
  	- mysql
```

- `*`

表示任意minion都会调用 vim，scripts state file

- `*web*`

minion-id 中正则匹配 web 都会调用 apache，python state file

- `*db*`

同上
在`/srv/salt`目录下已经存在了一个top.sls文件，可以修改成如下
```
root@saltmaster:/srv/salt# cat top.sls 
base:
  '*':
    - common
  'minion2':
    - nettools
```
在master 运行`slat '*' state.apply`就会自动调用top.sls，minion1 只会调用common state file，minion 2 会调用 common 和 nettools state file
### Init file
init.sls，会在调用`state.apply`或 TOP file 被调用时使用
### Exmaple
具体使用可以查看源码
安装包
```
install vim:
  pkg.installed:
    - name: vim
```
删除包
```
remove vim:
  pkg.removed:
    - name: vim
```
生成目录文件
```
create my_new_directory:
 file.directory:
   - name: /opt/my_new_directory
   - user: root
   - group: root
   - mode: 755
```
运行服务
```
Make sure the mysql service is running:
  service.running:
    - name: mysql
```
安装服务并运行
```
Install mysql and make sure the mysql service is running:
  pkg.installed:
    - name: mysql
  service.running:
    - name: mysql
```
## Pillar
pillar 用于数据存储，类似于SpEL 表达式，使用 jinja2 语法(`{{}}`)
### TOP file
和 state 类似的 pillar 也有 top file(`/srv/pillar`)
```
root@saltmaster:/srv/pillar# cat top.sls 
base:
  '*':
    - default
root@saltmaster:/srv/pillar# cat default.sls 
# Default pillar values
command: hostname
```
所有minion都会调用pillar default.sls，使用`salt '*' saltutil.refresh_pillar`来刷新minion上的pillar(如果只改了value的值可以不使用，但是如果改了key的值就一定要使用)
```
root@saltmaster:/srv/pillar# salt '*' saltutil.refresh_pillar
minion2:
    True
minion1:
    True
```
这样就可以在sate中通过jinja的语法来调用设定的pillar
```
root@saltmaster:/srv/salt# cat test_pillar.sls 
test:
  cmd.run:
    - name: {{ pillar['command'] }}
```
```

root@saltmaster:/srv/salt# salt 'minion1' state.apply test_pillar
minion1:
----------
          ID: test
    Function: cmd.run
        Name: hostname
      Result: True
     Comment: Command "hostname" run
     Started: 06:10:51.726355
    Duration: 6.46 ms
     Changes:   
              ----------
              pid:
                  3210
              retcode:
                  0
              stderr:
              stdout:
                  minion1

Summary for minion1
------------
Succeeded: 1 (changed=1)
Failed:    0
------------
Total states run:     1
Total run time:   6.460 ms
```
如果只是为了测试或临时使用，可以通过如下方式传入pillar
```
root@saltmaster:/srv/pillar# salt 'minion1' state.apply test_pillar pillar='{"command":"whoami"}'
minion1:
----------
          ID: test
    Function: cmd.run
        Name: whoami
      Result: True
     Comment: Command "whoami" run
     Started: 06:15:42.687060
    Duration: 8.375 ms
     Changes:   
              ----------
              pid:
                  3284
              retcode:
                  0
              stderr:
              stdout:
                  root

Summary for minion1
------------
Succeeded: 1 (changed=1)
Failed:    0
------------
Total states run:     1
Total run time:   8.375 ms
```
可以使用`salt '*' pillar.items`，来看所有调用的pillar
```
root@saltmaster:/srv/pillar# salt '*' pillar.items
minion1:
    ----------
    command:
        df -hT
minion2:
    ----------
    command:
        df -hT
```
## Include
include 可以让 salt state 复用，和 nginx 中的 include direction 一样。但是需要在 state file 的置顶位置声明
```
include:
  - sls1 
  - sls2
```
这里的 sls1、sls2 表示 state file without .sls extension，如果 state file 在目录下需要使用如下格式
```
include:
  - dir.sls1 
```
可以使用double dots 的格式
```
include:
  - ..etc.check
```
同样的也可以把 include 写入 state top file
```
base:
  'web*':
    - sls1
    - sls2
```
### Example
lftp.sls
```
install lftp:
  pkg.installed:
    - name: lftp
```
dir-sync.sls
```
include:
  - lftp

sync directory using lftp:
  cmd.run:
    - name: lftp -c "open -u {{ pillar['ftpusername'] }},{{ pillar['ftppassword'] }}
           -p 22 sftp://example.com;mirror -c -R /local /remote"
```
## Ordering
在salt中是按照文件中记录的顺序执行命令的
![Snipaste_2021-12-08_09-38-21](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20220223/Snipaste_2021-12-08_09-38-21.13dp8f4o8e40.webp)
先执行formula1(先执行id1，再执行id2)，再执行formula2(先执行id3，在执行id4)
可以使用`salt 'minion1' state.show_sls examples`来看state file中执行的具体顺序，example 表示state file
## JINJA
语法有点像 servlet + spEL，用于 pillar 取逻辑变量
### conditionals
如果minion由不同OS组成可以使用grains来获取对应的信息，
这里的`grians[os_famliy] == salt '*' grains.item os_famliy`
```
root@saltmaster:/srv/pillar# cat default.sls 
# Default pillar values
{% if grains['os_family'] == 'RedHat' %}
apache: httpd
git: git
{% elif grains['os_family'] == 'Debian' %}
apache: apache2
git: git-core
{% endif %}
```
查看每个minion使用的pillar
```
root@saltmaster:/srv/pillar# salt '*' pillar.items
minion1:
    ----------
    apache:
        apache2
    git:
        git-core
minion2:
    ----------
    apache:
        apache2
    git:
        git-core
```
这样就可以在state file 中自动按照OS来选择安装的包
```
install apache:
  pkg.installed:
    - name: {{ pillar['apache'] }}
```
### loops
```
{% for usr in ['moe','larry','curly'] %}
{{ usr }}:
  user.present
{% endfor %}
```
```
{% for DIR in ['/dir1','/dir2','/dir3'] %}
{{ DIR }}:
  file.directory:
    - user: root
    - group: root
    - mode: 774
{% endfor %}
```
## Manage file
## salt://
类似于java classpath，表示`srv/salt`下的相对路径
```
deploy the http.conf file:
  file.managed:
    - name: /etc/http/conf/http.conf
    - source: salt://apache/http.conf
```
这里的source 等价于`/srv/salt/apache/http.conf`，类似的也可以使用`http`、`https`
