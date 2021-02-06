# Linux 删除virbr0 NIC

参考：

https://blog.51cto.com/xjsunjie/1914963

```
virsh net-list
virsh net-destroy default
virsh net-undefine default
service libvirtd restart
```

