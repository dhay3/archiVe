# docker-network

管理容器的网络，与`docker run`的`--network`参数类似，但是多用于管理bridge网络。

## docker network create

用于创建一个网络

pattern：`docker network create [options] <network-name>`

### 通用参数

- `--driver | -d`

  指定生成网络的类型，缺省值bridge。

  ```
  root in ~ λ docker network create -d bridge my-bridge-network
  8227855b6c13d2a3ce154efb8cba58e9be8ced2a276ae7f61bcf36c7e2790ff7
  ```

- `--subnet`指定网络所处的子网

  `--ip-range`指定网络范围

  `--gateway`指定网关，如果忽略默认会从ip pool中选一个

  ```
  $ docker network create \
    --driver=bridge \
    --subnet=172.28.0.0/16 \
    --ip-range=172.28.5.0/24 \
    --gateway=172.28.5.254 \
    br0
  ```

  

## docker network connet

用于将容器连接到指定的网络。可以使用docker network inspect来查看。会在容器原有网络的基础上新建iface

pattern：`docker network connect [options] <networkname> <container> `

```
root in ~ λ docker network connect my-bridge t1
root in ~ λ docker network connect my-bridge t2

root in ~ λ docker network inspect my-bridge -f '{{json .Containers}}'
{"9cf3328ccd3b8a4cd062227868be6f039b98b78f8cff247a9e9c7403da31e6a9":{"Name":"t2","EndpointID":"077888f68e08e51f02f1faa2afd1bd751ed5fe48b1605f87f38890516ac29fa5","MacAddress":"02:42:ac:15:00:02","IPv4Address":"172.21.0.2/16","IPv6Address":""},"c3cc16c9484450f64210f0ae0e0e9ac1bf1f57bd54cf96f4c56af283a0f33b6c":{"Name":"t1","EndpointID":"e6a3b5f3b4a147f59dc84934f8a53bce4acde9468220c4e22868d520edc99cd1","MacAddress":"02:42:ac:15:00:03","IPv4Address":"172.21.0.3/16","IPv6Address":""}}
```

## docker network disconnect

将容器与网络断开

```
root in ~ λ docker network disconnect help  t2
```

## docker network ls

展示docker deamon记录在案的所有network，参数与docker inspect相同。可以使用go template

```
root in ~ λ docker network  ls
NETWORK ID     NAME                DRIVER    SCOPE
80b9b205cd4c   bridge              bridge    local
9dfaae564286   help                bridge    local
81078d4d17f9   host                host      local
8227855b6c13   my-bridge-network   bridge    local
6cc6d934173d   none                null      local
```

## docker network inspect

查看指定网络的具体信息

```
root in ~ λ docker network inspect help
[
    {
        "Name": "help",
        "Id": "9dfaae5642869ed43741a1d8a895e4a88efbe734060430b6b9e930b64e867c87",
        "Created": "2021-02-23T11:51:31.421126036+08:00",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "172.20.0.0/16",
                    "Gateway": "172.20.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {},
        "Options": {},
        "Labels": {}
    }
]
```

## docker network prune

删除所有没有被使用的网络

```
root in ~ λ docker network prune
WARNING! This will remove all custom networks not used by at least one container.
Are you sure you want to continue? [y/N] y
Deleted Networks:
my-bridge-network
help
```

## docker network rm

删除指定网络