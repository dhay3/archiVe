# aliyun VPN

## VPC 2 IDC

![Snipaste_2021-04-21_10-54-27](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2021-04-21_10-54-27.1hq8qrkhlev4.png)

1. VPC与IDC的网络地址不能冲突，ICD的VPN网关必须配置一个静态公网IP
2. VPN网关是在VPC内的，VPC内ECS可以通过VPN网关向外通信。
3. 创建完成后，可以在VPC内任意无公网IP的主机通过ping命令，ping IDC中私网IP进行校验

## VPC 2 VPC

1. 需要互通的VPC私网IP地址不重叠
2. 需要为两个VPC都创建VPN，且VPC与VPN地域相同
3. 需要开启SSL-VPN功能