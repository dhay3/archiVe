# Macchanger

转自:

https://blog.csdn.net/qq_26090065/article/details/80500137

###### MAC地址

> - 物理地址、硬件地址
> - 定义网络设备的位置
> - 在OSI模型中,第二层数据链路层则负责 MAC地址
> - MAC地址是网卡决定的，是固定的。

```
ifconfig eth0 #或是使用 ip l 显示
# HWaddr即为MAC地址，前三项是 `设备制造商` 信息12
```

MAC地址在`同一内网`中是可见的

# `macchanger` —– 修改mac地址的工具

> 只能修改让别人看见的mac地址，`实际`真正的mac地址`不能改变`

```
macchanger -s eth0   #查看eth0的MAC地址
macchanger -r eth0   #随机生成并修改eth0的MAC地址

其他参数：

  -h,  --help                   帮助信息
  -V,  --version                输出版本信息并退出
  -s,  --show                   输出MAC地址并退出
  -e,  --ending                 不改变有关设备商信息的字节
  -a,  --another                设置相同供应商的随机MAC地址
  -A                            设置任意类型的供应商MAC
  -p,  --permanent              重置为原始、永久的硬件MAC
  -r,  --random                 设置完全随机MAC地址
  -l,  --list[=keyword]         输出已知的供应商
  -b,  --bia                    冒充地址
  -m,  --mac=XX:XX:XX:XX:XX:XX
       --mac XX:XX:XX:XX:XX:XX  设置 MAC XX:XX:XX:XX:XX:XX
123456789101112131415161718
```

# 开机自动修改 MAC地址

```
crontab -e
```

> crontab命令 用于设置周期性被执行的指令

选择一个编辑器，回车默认选择nano打开
最后一行添加

```
@reboot macchanger -r eth01
```

Ctrl+x 退出
回车 确认文件名