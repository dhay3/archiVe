# packer 入门

> packer从1.7后就开始推荐使用HCL2来做为配置文件，替代JSON格式
>
> 如果packer在有特殊字符的地方被调用，会出错
>
> json格式需要以`.json`结尾，如果想要将HCL2格式转换为JSON，需要以`.pkr.json`结尾
>
> https://github.com/hashicorp/packer/issues/9112
>
> hcl格式需要以`.pkr.hcl`结尾

## 概述

packer是一个自动化的工具，我们可以依靠packer来快速生成镜像

## terms

- builder

  packer用于生成镜像的组件

- artifact

  builder生成代表镜像的文件

- build

  生成镜像的单个任务

- command

  packer的子命令

- data source

  packer用于获取数据的组件

- post-processor

  https://www.packer.io/docs/templates/hcl_templates/blocks/build/post-processor

  packer用于处理build完成后的组件

- provisioner

  https://www.packer.io/docs/templates/hcl_templates/blocks/build/provisioner

  packer用于处理artifact启动后执行的动作

- template

  生成镜像的配置文件



​	