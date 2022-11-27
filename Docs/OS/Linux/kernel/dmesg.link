# dmesg

## ring buffer

https://en.wikipedia.org/wiki/Circular_buffer

ring buffer是一个固定大小的空间，可以想象成一个环。如果收到新的消息会将旧的消息删除(FIFO)。TCP的接受窗口就是一个例子。==linux kernel ring buffer==记录了涉及到kernel相关的操作

## dmesg

https://www.thegeekstuff.com/2010/10/dmesg-command-examples/

https://www.computerhope.com/unix/dmesg.htm

dmesg用于打印出linux kernel ring buffer。一般用于trouble shoot bootup 过程涉及到kernel的部分，==还可以查看GRUB使用的CMD_LINE_LINUX==

- -T | --ctime

  dmesg默认以计算机的时间来显示数据，可以使用该参数以可读的状态显示

  ```
  cpl in /etc λ sudo dmesg -T | head -1
  [Mon Jul 19 11:20:54 2021] amdgpu 0000:03:00.0: amdgpu: SMU is resumed successfully!
  ```

- -W | --follow-new

  以follow的形式显示kernel buffer ring 中的新内容

- -C | --clear

  清空ring buffer中的内容



















