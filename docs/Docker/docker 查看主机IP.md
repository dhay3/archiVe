# docker 查看主机IP

我们使用docker是在宿主机上只会显示容器的MAC地址。为了查看容器的IP我们可以使用`docker inspect`来查看主机的IP

```
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "be2480f5c4afc71ce556a3eecaf9503dabd42028dde50eb55e00acf29cc79838",
                    "EndpointID": "8304b1b9471d68d79d0518f626c5102b5c63dfc4d02ddae220cb7816b113c3f8",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.2",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:02",
                    "DriverOpts": null
                }

```

使用`-f`参数缩小范围

```
root in /etc/ssh λ docker inspect --format='{{json .NetworkSettings.Networks.bridge.IPAddress }}' t1
"172.17.0.2"     
```

