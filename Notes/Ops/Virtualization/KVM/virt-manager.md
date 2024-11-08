# virt-manager

virt-manager是一款通过libvirt来管理虚拟机的UI工具。前提是开启`libvirtd`

```
systemctl start libvrtd && systemctl enable libvrtd
```

添加用户组

```
cpl in /mnt/win λ sudo usermod -aG libvirt $USER
```

