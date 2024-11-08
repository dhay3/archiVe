# docker-inspect

以JSON的形式展示container，image，volume等信息

```
root in ~ λ docker ps
CONTAINER ID   IMAGE     COMMAND       CREATED       STATUS       PORTS     NAMES
05950b8c4220   test1     "/bin/bash"   2 hours ago   Up 2 hours             test11                                                    /0.1s
root in ~ λ docker inspect test11
[
    {
        "Id": "05950b8c4220c8328592bec85c2154665207bd48869ff953b7cd7eb58d10b286",
        "Created": "2021-01-22T03:38:41.465708866Z",
        "Path": "/bin/bash",
        "Args": [],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 18851,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2021-01-22T03:38:41.97321086Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
        "Image": "sha256:fb1f2d882cf56b06f6fbe55597fc59fcb67c00ee9e780c9bc67ce44016c8fefa",
        "ResolvConfPath": "/var/lib/docker/containers/05950b8c4220c8328592bec85c2154665207bd48869ff953b7cd7eb58d10b286/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/05950b8c4220c8328592bec85c2154665207bd48869ff953b7cd7eb58d10b286/hostname",
        "HostsPath": "/var/lib/docker/containers/05950b8c4220c8328592bec85c2154665207bd48869ff953b7cd7eb58d10b286/hosts",
        "LogPath": "/var/lib/docker/containers/05950b8c4220c8328592bec85c2154665207bd48869ff953b7cd7eb58d10b286/05950b8c4220c8328592bec85c2154665207bd48869ff953b7cd7eb58d10b286-json.log",
        "Name": "/test11",
        "RestartCount": 0,
        "Driver": "overlay2",
        "Platform": "linux",
        "MountLabel": "",
        "ProcessLabel": "",
        "AppArmorProfile": "docker-default",
        "ExecIDs": null,
        "HostConfig": {
            "Binds": null,
            "ContainerIDFile": "",
            "LogConfig": {
                "Type": "json-file",
                "Config": {}
            },
            "NetworkMode": "default",
            "PortBindings": {},
            "RestartPolicy": {
                "Name": "no",
                "MaximumRetryCount": 0
            },
            "AutoRemove": false,
            "VolumeDriver": "",
            "VolumesFrom": null,
            "CapAdd": null,
            "CapDrop": null,
            "CgroupnsMode": "host",
            "Dns": [],
            "DnsOptions": [],
            "DnsSearch": [],
            "ExtraHosts": null,
            "GroupAdd": null,
            "IpcMode": "private",
            "Cgroup": "",
            "Links": null,
            "OomScoreAdj": 0,
            "PidMode": "",
            "Privileged": false,
            "PublishAllPorts": false,
            "ReadonlyRootfs": false,
            "SecurityOpt": null,
            "UTSMode": "",
            "UsernsMode": "",
            "ShmSize": 67108864,
            "Runtime": "runc",
            "ConsoleSize": [
                0,
                0
            ],
            "Isolation": "",
            "CpuShares": 0,
            "Memory": 0,
            "NanoCpus": 0,
            "CgroupParent": "",
            "BlkioWeight": 0,
            "BlkioWeightDevice": [],
            "BlkioDeviceReadBps": null,
            "BlkioDeviceWriteBps": null,
            "BlkioDeviceReadIOps": null,
            "BlkioDeviceWriteIOps": null,
            "CpuPeriod": 0,
            "CpuQuota": 0,
            "CpuRealtimePeriod": 0,
            "CpuRealtimeRuntime": 0,
            "CpusetCpus": "",
            "CpusetMems": "",
            "Devices": [],
            "DeviceCgroupRules": null,
            "DeviceRequests": null,
            "KernelMemory": 0,
            "KernelMemoryTCP": 0,
            "MemoryReservation": 0,
            "MemorySwap": 0,
            "MemorySwappiness": null,
            "OomKillDisable": false,
            "PidsLimit": null,
            "Ulimits": null,
            "CpuCount": 0,
            "CpuPercent": 0,
            "IOMaximumIOps": 0,
            "IOMaximumBandwidth": 0,
            "MaskedPaths": [
                "/proc/asound",
                "/proc/acpi",
                "/proc/kcore",
                "/proc/keys",
                "/proc/latency_stats",
                "/proc/timer_list",
                "/proc/timer_stats",
                "/proc/sched_debug",
                "/proc/scsi",
                "/sys/firmware"
            ],
            "ReadonlyPaths": [
                "/proc/bus",
                "/proc/fs",
                "/proc/irq",
                "/proc/sys",
                "/proc/sysrq-trigger"
            ]
        },
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/be13e7dcf7a468fd5e6e53bc80b841b3b4e582d1acf7ec8685509c8769a2dece-init/diff:/var/lib/docker/overlay2/fb687346c740870d06c7bd175c2536aa256bf80c5c4a9dfc354eb21a6d3efc6a/diff:/var/lib/docker/overlay2/8b34d519559d6fbce6056a122345a65769d285781a83273111f633cd7f07aa1a/diff:/var/lib/docker/overlay2/6cc36754563028365968486d0c53e8a1667ff344c49695b53e4aa8d4da16cee6/diff",
                "MergedDir": "/var/lib/docker/overlay2/be13e7dcf7a468fd5e6e53bc80b841b3b4e582d1acf7ec8685509c8769a2dece/merged",
                "UpperDir": "/var/lib/docker/overlay2/be13e7dcf7a468fd5e6e53bc80b841b3b4e582d1acf7ec8685509c8769a2dece/diff",
                "WorkDir": "/var/lib/docker/overlay2/be13e7dcf7a468fd5e6e53bc80b841b3b4e582d1acf7ec8685509c8769a2dece/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [
            {
                "Type": "volume",
                "Name": "e266cda9aaef84862674ff49822e6d6da59fd62e2c04a02f9e579a62cc3f4de2",
                "Source": "/var/lib/docker/volumes/e266cda9aaef84862674ff49822e6d6da59fd62e2c04a02f9e579a62cc3f4de2/_data",
                "Destination": "/myData",
                "Driver": "local",
                "Mode": "",
                "RW": true,
                "Propagation": ""
            }
        ],
        "Config": {
            "Hostname": "05950b8c4220",
            "Domainname": "",
            "User": "",
            "AttachStdin": true,
            "AttachStdout": true,
            "AttachStderr": true,
            "Tty": true,
            "OpenStdin": true,
            "StdinOnce": true,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Cmd": [
                "/bin/bash"
            ],
            "Image": "test1",
            "Volumes": {
                "/myData": {}
            },
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": {
                "org.label-schema.build-date": "20201204",
                "org.label-schema.license": "GPLv2",
                "org.label-schema.name": "CentOS Base Image",
                "org.label-schema.schema-version": "1.0",
                "org.label-schema.vendor": "CentOS"
            }
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "ec7b8adcecfa3bdf698516613759b298defbfab4191aa2cc09a01bcad33dc750",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {},
            "SandboxKey": "/var/run/docker/netns/ec7b8adcecfa",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "ff7511f0a0a799dd424673c548317444b5b5d0b239ba2dbbbd7969127bf9a2e1",
            "Gateway": "172.17.0.1",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "172.17.0.2",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "MacAddress": "02:42:ac:11:00:02",
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "4f18c15cf716797ccc3f786a92648a01ccc239d82605727cae520b4b4545da0c",
                    "EndpointID": "ff7511f0a0a799dd424673c548317444b5b5d0b239ba2dbbbd7969127bf9a2e1",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.2",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:02",
                    "DriverOpts": null
                }
            }
        }
    }
]                        
```

- `--format`

  以go template的形式打印指定的模块。这里的json是调用函数，Mounts是参数。

  ```
  root in ~ λ docker inspect t1 --format="{{json .Mounts}}"
  [{"Type":"volume","Name":"bb7b9d2e8c652b2cd30fdde0ab4cc10965967031418c5afe02c03a48046de1a5","Source":"/var/lib/docker/volumes/bb7b9d2e8c652b2cd30fdde0ab4cc10965967031418c5afe02c03a48046de1a5/_data","Destination":"/etc","Driver":"local","Mode":"","RW":true,"Propagation":""}] 
  ```

  可以使用python来格式化JSON输出
  
  ```
    test docker inspect centos --format="{{json .NetworkSettings}}" | python -m json.tool
  {
      "Bridge": "",
      "SandboxID": "080f049186919dd56ab6733211d18c8c3f1d5b0b93d84811960dc7328e49968d",
      "HairpinMode": false,
      "LinkLocalIPv6Address": "",
      "LinkLocalIPv6PrefixLen": 0,
      "Ports": {},
  ...omit...
  ```
  
  