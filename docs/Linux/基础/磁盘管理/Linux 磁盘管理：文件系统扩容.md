# Linux 磁盘管理：扩容

> 只针对非LVM系统使用(LVM扩容非常简单)，可以按照aliyun文档的方法扩容
>
> https://help.aliyun.com/document_detail/113316.html?spm=a2c4g.11186623.6.874.24a368a29SfAVQ
>
> https://help.aliyun.com/document_detail/111738.html?spm=a2c4g.11186623.2.26.37e84eb7nD8yMv
>
> 无需修改`/ect/fstab`

1. ```
   umount /dev/sdb1
   ```

2. 删除原有的分区（在分区上原来的数据不会消失）

   ```
   fdisk /dev/sdb1 
   d
   w
   ```

3. 新建分区（新建分区必须与原分区起始块位置相同，终止位置大于原有的）

   ```
   fdisk /dev/sdb
   n
   w
   ```

4. 扩容文件系统
   - ext* `resize2fs /dev/sdb1`
   - xsf `xfs_growfs /mnt`

5. 挂载

   `mount /dev/vdb1 /mnt`

