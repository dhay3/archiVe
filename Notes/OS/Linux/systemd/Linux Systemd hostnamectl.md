# Linux Systemd hostnamectl

## Digest

Syntax：`hostnamectl [options] {command}`

control the system hostname

query and change the system hostname and related settings



## Postional args

- `status`

  等价于`hostnamectl`，对比`uname -r`和`cat /proc/version`

  ```
  [root@cyberpelican systemd]# hostnamectl 
     Static hostname: cyberpelican
           Icon name: computer-vm
             Chassis: vm
          Machine ID: fbb74b6620184684961580de92e236c2
             Boot ID: d466e1cc725d4316875231666fd7de59
      Virtualization: vmware
    Operating System: CentOS Linux 7 (Core)
         CPE OS Name: cpe:/o:centos:centos:7
              Kernel: Linux 3.10.0-1062.el7.x86_64
        Architecture: x86-64
        
  [root@cyberpelican systemd]# cat /proc/version 
  Linux version 3.10.0-1062.el7.x86_64 (mockbuild@kbuilder.bsys.centos.org) (gcc version 4.8.5 20150623 (Red Hat 4.8.5-36) (GCC) ) #1 SMP Wed Aug 7 18:08:02 UTC 2019
  [root@cyberpelican systemd]# uname -r
  3.10.0-1062.el7.x86_64
  
  
  ```

- `set-hostname `

  设置主机名

  ```
  [root@cyberpelican systemd]# hostnamectl set-hostname  cyberpelican
  ```
