# packer打包阿里镜像

https://help.aliyun.com/document_detail/118730.html

## required

- `access_key(string)`，可以设置ALICLOUD_ACCESS_KEY环境变量来导入
- `secret_key(string)`，可以设置ALICLOUD_SECRECT_KEY环境变量来导入
- `region(string)`，可以设置[ALICLOUD_REGION][https://help.aliyun.com/document_detail/50086.html]环境变量来导入

前三个值也可以设置在profile中

- `instance_type(string)`，实例的规格
- `source_image(string)`，基础镜像的id，从[阿里云公共镜像列](https://help.aliyun.com/document_detail/100410.html?spm=a2c4g.11186623.6.768.5ad36e815TQN2m)中获得
- `image_name(string)`，自定义生成镜像的名字

## optional

- description(string)

- internet_charge_type(string)，网络付费类型，有以下两个值
  1. PayByBandwidth，按带宽付费，缺省值
  2. PayByTraffic，按量付费
- io_optimized(boolean)，IO是否优化，默认按照`instance_type`决定
- tags(maps[string]string)，为生成的镜像应用阿里云的标签
- security_group_id(string)，将生成的镜像加入到指定的安全组id。缺省值加入到默认安全组。
- vpc_id(string)，将生成的镜像计入到指定的vpc id
- vpc_name(string)，缺省值空
- vpc_cidr_block(string)，Value options: 192.168.0.0/16 and 172.16.0.0/16. When not specified, the default value is 172.16.0.0/16.
- sercurity_group_name(string)，将生成的镜像加入到指定的安全组name。缺省值空
- system_disk_mapping

## 例子

这里使用go template，`env`表示从环境变量中获取值(AK设置为环境变量)，`user`表示从配置文件获取值

```
{
  "variables": {
    "access_key": "{{env `ALICLOUD_ACCESS_KEY`}}",
    "secret_key": "{{env `ALICLOUD_SECRET_KEY`}}"
  },
  "builders": [{
    "type":"alicloud-ecs",
    "access_key":"{{user `access_key`}}",
    "secret_key":"{{user `secret_key`}}",
    "region":"cn-shenzhen",
    "image_name":"cus_image",
    "source_image":"ubuntu_18_04_x64_20G_alibase_20210128.vhd",
    "ssh_username":"root",
    "instance_type":"ecs.t5-c1m2.large",
    "internet_charge_type":"PayByTraffic",
    "io_optimized":"true",
    "tags": {
      "tag_key": "tag_value"
    }
  }]
}
```

创建镜像

```
root in /usr/local/packer_test λ packer build t.json
alicloud-ecs: output will be in this color.

==> alicloud-ecs: Prevalidating source region and copied regions...
==> alicloud-ecs: Prevalidating image name...
    alicloud-ecs: Found image ID: ubuntu_18_04_x64_20G_alibase_20210128.vhd
==> alicloud-ecs: Creating temporary keypair: packer_6041a252-f7e0-5ba7-1c12-c6b1d22c1dce
==> alicloud-ecs: Creating vpc...
    alicloud-ecs: Created vpc: vpc-wz971ynrvbdx81w6jc8rd
==> alicloud-ecs: Creating vswitch...
    alicloud-ecs: Created vswitch: vsw-wz9bbdrni4t2ab0c3puuv
==> alicloud-ecs: Creating security group...
    alicloud-ecs: Created security group: sg-wz9h5wjsrwdko7t6y213
==> alicloud-ecs: Creating instance...
==> alicloud-ecs: Error creating instance: SDK.ServerError
==> alicloud-ecs: ErrorCode: InvalidAccountStatus.NotEnoughBalance
==> alicloud-ecs: Recommend: https://error-center.aliyun.com/status/search?Keyword=InvalidAccountStatus.NotEnoughBalance&source=PopGw
==> alicloud-ecs: RequestId: C88C584E-BA14-4A1E-B302-F81532313629
==> alicloud-ecs: Message: Your account does not have enough balance.
==> alicloud-ecs: Deleting security group because of cancellation or error...
==> alicloud-ecs: Deleting vSwitch because of cancellation or error...
==> alicloud-ecs: Deleting VPC because of cancellation or error...
==> alicloud-ecs: Deleting temporary keypair...
```

