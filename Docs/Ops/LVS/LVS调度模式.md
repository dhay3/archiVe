# LVS调度模式

## LVS-NAT

<img src="D:\asset\note\imgs\_LVS\Snipaste_2020-11-22_13-50-25.png"/>

- 请求：源地址CIP，目标地址VIP，LVS将目标地址转换为RIP，具体请求那一台主机更具配置的算法
- 响应：源地址RIP，目标地址CIP，LVS将源地址转为换为VIP

