# vagrant destroy

## 0x0 Overview

syntax

```
vagrant destroy [name|id]
```

用于停止 vagrant 运行的虚拟机并销毁对应的虚拟资源 (注意并不会删除对应的 box, 如果需要删除 box 请使用 `vagrant remove`)

需要先试用 `vagrant global-status` 查看对应的 name 和 id

```
(base) 0x00 in ~/Hypervisor/vagrant-machines/centos7 λ vagrant global-status   
id       name    provider   state    directory                                      
------------------------------------------------------------------------------------
24e356c  default virtualbox aborted  /home/0x00/Hypervisor/vagrant-machines/centos7 
bfb9175  default virtualbox poweroff /home/0x00/Hypervisor/vagrant-machines/kylin   
 
The above shows information about all known Vagrant environments
on this machine. This data is cached and may not be completely
up-to-date (use "vagrant global-status --prune" to prune invalid
entries). To interact with any of the machines, you can go to that
directory and run Vagrant, or you can use the ID directly with
Vagrant commands from any directory. For example:
"vagrant destroy 1a2b3c4d"
                                                                                                                                                                                  
(base) 0x00 in ~/Hypervisor/vagrant-machines/centos7 λ vagrant destroy bfb9175        
==> default: VM not created. Moving on...
```