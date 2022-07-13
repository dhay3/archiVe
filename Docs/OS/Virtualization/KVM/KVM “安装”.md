ref
[https://wiki.archlinux.org/title/KVM](https://wiki.archlinux.org/title/KVM)
## 前置条件
### 硬件

1. A VT capable Intel processor, or an SVM capable AMD processor

如果是intel CPU virtualization 的值必须是 VT-x，如果是amd CPU virtualization 的值必须是AMD-V。可以通过如下的方式校验
```
cpl in ~ λ lscpu| grep -i virtual
Address sizes:                   48 bits physical, 48 bits virtual
Virtualization:                  AMD-V

#or

grep -E --color=auto 'vmx|svm|0xc0f' /proc/cpuinfo
```

