# packer  builder docker

https://www.packer.io/docs/builders/docker

## 常用属性

- `pull`：true会检查当前是否已有镜像，如果没有就会从docker hub上拉取，如果有就以当前镜像为基础。默认false，每次都会从docker hub上拉取

## 例子

将拉取的镜像打包成tar文件

```
{
  "builders": [
    {
      "type": "docker",
      "image": "busybox",
      "export_path":"./docker.tar"
    }
  ]
}
```

从docker hub拉取指定镜像创建完镜像后执行docker commit，生成的镜像没有tag

```
{
  "builders": [
    {
      "type": "docker",
      "image": "busybox",
      "commit": "true"
    }
  ]
}
```

