# vagrant snapshot

## 0x0 Overview

管理 snapshot 的一系列接口

- push
- pop
- save
- restore
- list
- delete

### list

显示当前的 snapshot

```
(base) 0x00 in ~/Hypervisor/vagrant-machines/centos7 λ vagrant snapshot list
==> default: 
centos
                 
```

### save

保存 snapshot

### restore

回退到 snapshot

### delete

删除 snapshot

**references**

[^1]:https://developer.hashicorp.com/vagrant/docs/cli/snapshot