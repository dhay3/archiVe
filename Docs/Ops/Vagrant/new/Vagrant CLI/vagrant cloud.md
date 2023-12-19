# vagrant cloud

## Overview

用于和 vagrant cloud 交互的命令，只要用到一个命令

```
vagrant cloud search hashicorp --limit 5
| NAME                    | VERSION | DOWNLOADS | PROVIDERS                       |
+-------------------------+---------+-----------+---------------------------------+
| hashicorp/precise64     | 1.1.0   | 6,675,725 | virtualbox,vmware_fusion,hyperv |
| hashicorp/precise32     | 1.0.0   | 2,261,377 | virtualbox                      |
| hashicorp/boot2docker   | 1.7.8   |    59,284 | vmware_desktop,virtualbox       |
| hashicorp/connect-vm    | 0.1.0   |     6,912 | vmware_desktop,virtualbox       |
| hashicorp/vagrant-share | 0.1.0   |     3,488 | vmware_desktop,virtualbox       |
+-------------------------+---------+-----------+---------------------------------+
```

**references**

[^1]:https://developer.hashicorp.com/vagrant/docs/cli/cloud