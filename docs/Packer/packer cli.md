# packer cli

> 使用前可以使用`packer -autocomplete-install`可以安装自动补全的插件，运行效率不是特别高，而且不智能(不推荐使用)https://github.com/hashicorp/terraform/issues/18649

当键入packer默认会打印所有的子命令

```
PS C:\WINDOWS\system32> packer
Usage: packer [--version] [--help] <command> [<args>]

Available commands are:
    build           build image(s) from template
    console         creates a console for testing variable interpolation
    fix             fixes templates from old versions of packer
    fmt             Rewrites HCL2 config files to canonical format
    hcl2_upgrade    transform a JSON template into an HCL2 configuration
    init            Install missing plugins or upgrade plugins
    inspect         see components of a template
    validate        check that a template is valid
    version         Prints the Packer version
```

使用`packer subcommand -h`可以查看具体子命令的用法

```
PS C:\WINDOWS\system32> packer init -h
Usage: packer init [options] [config.pkr.hcl|folder/]

  Install all the missing plugins required in a Packer config. Note that Packer
  does not have a state.

  This is the first command that should be executed when working with a new
  or existing template.

  This command is always safe to run multiple times. Though subsequent runs may
  give errors, this command will never delete anything.

Options:
  -upgrade                     On top of installing missing plugins, update
                               installed plugins to the latest available
                               version, if there is a new higher one. Note that
                               this still takes into consideration the version
                               constraint of the config.
```

## init

==init命令对传统的JSON不生效，需要使用HCL格式==

这个命令通常用于为一个新的template下载packer的二进制插件。这个命令可以被安全的调用多次。

使用`required_plugin`来导入插件

```
packer {
  required_plugins {
    happycloud = {
      version = ">= 2.7.0"
      source = "github.com/azr/happycloud"
    }
  }
}

---

root in /usr/local/packer_test λ packer init t.pkr.hcl
no release version found for the github.com/azr/happycloud plugin matching the constraint(s): ">= 2.7.0"
```

## build

用于生成artifacts(镜像)

- `-var `

  设置在template中使用的变量，可以被重复使用

  ```
  packer build \
      -var "ami_name=packer-tutorial" \
      example.pkr.hcl
  ```

- `-var-file`

  设置变量从文件中读取，一行

- `-force`

  强制生成artifacts，即使存在

- `-on-error=cleanup(default) | abort | ask | run-cleanup-provisioner`

  只在provisioner run(aritfact boot阶段)生效，不在post-processor run(build 完成后)生效

  1. cleanup：失败后回滚，删除所有生成的文件
  2. abort：退出进程但是不删除文件，下次生成文件需要使用`-force`
  3. ask：询问使用策略
  4. run-cleanup-provisioner

- `-parallel-builds=N`

  限制builds个数(构建镜像时的线程)，默认0表示无限制

## console

用于测试配置文件的变量是否设置正确

- `-var`
- `-var-file`

## fix

> fix 命令当前对HCL2格式的配置文件不生效

可以将配置文件中不兼容的部分升级到可以被最新版本的packer识别的内容，默认在stdout中输出。

## hcl2_upgrade

将JSON格式的配置文件更新为HCL2格式的配置文件

## validate

检测配置文件的语法是否正确

- `-var`
- `-var-file`

如果文件有错会输出错误，如果正确不会输出任何信息

```
root in /usr/local/packer_test λ packer validate t.json
Error: Failed to prepare build: "alicloud-ecs"

1 error(s) occurred:

* ALICLOUD_ACCESS_KEY and ALICLOUD_SECRET_KEY must be set in template file or
environment variables.

```

## fmt

格式化HCL2配置文件













