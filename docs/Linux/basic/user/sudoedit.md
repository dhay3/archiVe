# sudoedit

`sudoedit`用于编辑文本而不是执行命令，等价于`sudo -e`

```
cowrie@win2k:/home/ubuntu$ sudoedit /root/.bashrc
sudoedit: /root/.bashrc unchanged
```

如果没有设置 SUDO_EDITOR,VISUAL and EDITOR 环境变量，默认使用`/etc/sudoers`中的editor。如果指定的文件不存在，会被创建。