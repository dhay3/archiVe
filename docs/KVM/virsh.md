# virsh

参考：

https://medium.com/@zyshang/%E4%BD%BF%E7%94%A8-virsh-%E7%AE%A1%E7%90%86%E8%99%9A%E6%8B%9F%E6%9C%BA-458a1a866a54

https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/virtualization_deployment_and_administration_guide/sect-Domain_Commands-Creating_a_virtual_machine_XML_dump_configuration_file#example-virsh-dumpxml

`virsh`是`libvirt`对外提供的命令行接口(CLI)，运行用户在创建虚拟机、删除虚拟机、创建网络、删除网络等。和virt-manager一样需要启动`libvirtd`

syntax：`virsh [option] <command><domain>`

domain也可以理解为virtual machine

- virsh define | create

  使用xml定义一台虚拟机

  ```
  
  ```

  可以通过`--validate`来校验

  ```
  
  ```

  可以通过`virsh dumpxml <domain>`导出domain xml文件

  ```
  # virsh dumpxml guest1 > guest1.xml
  # cat guest1.xml
  <domain type='kvm'>
    <name>guest1-rhel6-64</name>
    <uuid>b8d7388a-bbf2-db3a-e962-b97ca6e514bd</uuid>
    <memory>2097152</memory>
    <currentMemory>2097152</currentMemory>
    <vcpu>2</vcpu>
    <os>
      <type arch='x86_64' machine='rhel6.2.0'>hvm</type>
      <boot dev='hd'/>
    </os>
  [...]
  ```

  

- virsh list

  查看本机连接的虚拟机

  ```
  
  ```

## domain xml

参考：

https://libvirt.org/formatdomain.html

根元素domain，需要有两个属性type(hypervisor虚拟机的管理器)可以是xen，kvm，qemu和lxc。id用于标识guest machine

```
<domain type='kvm' id='1'>...</domain>
```



