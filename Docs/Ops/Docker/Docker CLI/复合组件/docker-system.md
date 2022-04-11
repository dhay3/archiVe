# docker-system

用于管理docker

## docker system df

用于展示docker deamon占用磁盘的空间

```
root in / λ docker system df -v
Images space usage:

REPOSITORY   TAG       IMAGE ID       CREATED        SIZE      SHARED SIZE   UNIQUE SIZE   CONTAINERS
busybox      latest    491198851f0c   3 days ago     1.232MB   0B            1.232MB       1
nginx        latest    f6d0b4767a6c   5 weeks ago    133MB     0B            133MB         0
centos       latest    300e315adb2f   2 months ago   209.3MB   0B            209.3MB       0

Containers space usage:

CONTAINER ID   IMAGE     COMMAND   LOCAL VOLUMES   SIZE      CREATED          STATUS          NAMES
94a996549284   busybox   "sh"      1               0B        18 minutes ago   Up 18 minutes   t1

Local Volumes space usage:

VOLUME NAME   LINKS     SIZE
hello         1         921B

Build cache usage: 0B

CACHE ID   CACHE TYPE   SIZE      CREATED   LAST USED   USAGE     SHARED
```

## docker system info

用于展示docker system-wide的信息

```

root in / λ docker info
Client:
 Context:    default
 Debug Mode: false
 Plugins:
  app: Docker App (Docker Inc., v0.9.1-beta3)
  buildx: Build with BuildKit (Docker Inc., v0.5.1-docker)

Server:
 Containers: 1
  Running: 1
  Paused: 0
  Stopped: 0
 Images: 3
 Server Version: 20.10.2
 Storage Driver: overlay2
  Backing Filesystem: extfs
  Supports d_type: true
  Native Overlay Diff: true
 Logging Driver: json-file
 Cgroup Driver: cgroupfs
 Cgroup Version: 1
 Plugins:
  Volume: local
  Network: bridge host ipvlan macvlan null overlay
  Log: awslogs fluentd gcplogs gelf journald json-file local logentries splunk syslog
 Swarm: inactive
 Runtimes: io.containerd.runc.v2 io.containerd.runtime.v1.linux runc
 Default Runtime: runc
 Init Binary: docker-init
 containerd version: 269548fa27e0089a8b8278fc4fc781d7f65a939b
 runc version: ff819c7e9184c13b7c2607fe6c30ae19403a7aff
 init version: de40ad0
 Security Options:
  apparmor
  seccomp
   Profile: default
 Kernel Version: 4.15.0-132-generic
 Operating System: Ubuntu 18.04.5 LTS
 OSType: linux
 Architecture: x86_64
 CPUs: 2
 Total Memory: 3.852GiB
 Name: ubuntu18.04
 ID: PLGP:UTKH:WJJS:COB4:ZS4T:LCEF:G4GI:F452:FO7K:MDHV:4CSI:3PIK
 Docker Root Dir: /var/lib/docker
 Debug Mode: false
 Registry: https://index.docker.io/v1/
 Labels:
 Experimental: false
 Insecure Registries:
  127.0.0.0/8
 Live Restore Enabled: false

WARNING: No swap limit support
```

## docker system prune

- `-a`

  同时会删除没有在使用的镜像

- `--volumes`

  删除容器时同时删除volumes，默认不会删除

```
root in / λ docker system prune
WARNING! This will remove:
  - all stopped containers
  - all networks not used by at least one container
  - all dangling images #repositry和tag为none的镜像
  - all dangling build cache

Are you sure you want to continue? [y/N] y
Total reclaimed space: 0B
root in / λ docker images
REPOSITORY   TAG       IMAGE ID       CREATED        SIZE
busybox      latest    491198851f0c   3 days ago     1.23MB
nginx        latest    f6d0b4767a6c   5 weeks ago    133MB
centos       latest    300e315adb2f   2 months ago   209MB
```

