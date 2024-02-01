# virsh

参考：

https://medium.com/@zyshang/%E4%BD%BF%E7%94%A8-virsh-%E7%AE%A1%E7%90%86%E8%99%9A%E6%8B%9F%E6%9C%BA-458a1a866a54

https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/virtualization_deployment_and_administration_guide/sect-Domain_Commands-Creating_a_virtual_machine_XML_dump_configuration_file#example-virsh-dumpxml

`virsh`是`libvirt`对外提供的命令行接口(CLI)，运行用户在创建虚拟机、删除虚拟机、创建网络、删除网络等。和virt-manager一样需要启动`libvirtd`

syntax：`virsh [option] <command> <domain>`

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

```xml
<domain id="1" type="kvm">
  <!-- general metadata -->
  <name>guest-host</name>
  <!-- 如果没有指定uuid会自动生成 -->
  <uuid>285a6535-39df-410a-bf78-76b6e4d00402</uuid>
  <!-- a short description -->
  <title>It's a guest host</title>
  <!-- description -->
  <description>Some human readable description</description>
  <!-- OS booting -->
  <os fireware="efi">
    <!-- 虚拟机OS镜像的类型 
hvm 表示在镜像裸机上运行，宿主机需要完全虚拟化
arch 宿主机cpu架构，machine
-->
    <type arch="x86_64" machine=""/>
    <!-- boot order -->
    <boot dev="hd"/>
    <boot dev="cdrom"/>
    <boot dev="fd"/>
    <bootmenu enable="yes" timeout="3000"/>
  </os>
  <!-- cpu核数 -->
  <vcpu>2</vcpu>
  <iothreads>2</iothreads>
  <!-- 引导阶段的最大RAM -->
  <memory uint="G">4</memory>
  <!-- 运行阶段的最大RAM -->
  <maxMemory uint="G">4</maxMemory>
  <!-- virsh reboot 和 virsh shutdown触发的事件 -->
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>restart</on_crash>
  <!-- cpu和虚拟机的一些特性 -->
  <features>
    <pae/>
    <acpi/>
    <apic/>
    <hap/>
  </features>
  <!-- 使用hardware来确定时间 -->
  <clock offset="localtime"/>
  <devices>
    <emulator>/usr/bin/qemu-system-x86_64</emulator>
    <!-- 
  type虚拟磁盘的来源可选file,block,dir,network,nvme等
  device虚拟磁盘的具体用途可选floppy，disk，cdrom，lun,defaulting
 -->
    <!-- 安装光盘的位置 -->
    <disk device="cdrom" type="file">
      <source file="/sharing/vm/ISO/Win10_21H1_Chinese(Simplified)_x64.iso"/>
    </disk>
    <disk device="disk" type="file">
      <!-- 虚拟磁盘的具体位置，由disk的type决定 -->
      <source file="/win/win-disk"/>
      <!-- 硬盘的详细信息，该配置用于指定主硬盘 -->
      <driver name="qemu" type="qcow2"/>
      <!-- host machine暴露给guest machine的磁盘 
      模拟磁盘的类型-->
      <target bus="hdb" dev="ide"/>
      <!-- 宿主可以通过访问本地目录访问虚拟机
passthrough 宿主机可以访问虚拟机上的用户，缺省值
mapped 宿主机可以访问虚拟机上 
 -->
    </disk>
    <filesystem type="mount">
      <source name=""/>
      <target dir=""/>
    </filesystem>
  </devices>
</domain>
```



