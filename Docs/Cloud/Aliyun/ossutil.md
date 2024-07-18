# ossutil

> 如果想要查看某一命令的具体用法，可以使用`ossutil hepl [command]`
>
> 还可以通过`-L ch`以中文显示帮助信息
>
> ossutil支持模糊匹配，包括文件、目录和bucket

## 下载安装

https://help.aliyun.com/document_detail/120075.html?spm=a2c4g.11186623.2.7.78bb2e43wz0SL1

> tip
>
> 创建链接或`alias`使用更通畅，如果需要修改shell的配置文件，==注意windows和linux上的crlf不同==

下载后的配置的选项将会做为ossutil的全局变量，如果ossutil需要更新可以使用`ossutil update`

## 常用命令

### config

https://help.aliyun.com/document_detail/120072.html?spm=a2c4g.11186623.6.866.3702ca0dRza2yK

config命令用于创建配置文件来存储oss访问信息。

可以指定文件做为配置

### select

#### ls

- `ls [option] [url]`

  查看bucket，或objects，url可以带有某一个object的关键字，ossuitil会自动模糊查询该关键字指定的object或object。

  可以使用`--limited-num`控制展示的行数，`--marker`用于指定起始位置跳过几行。使用`-m`指定bucket中为完成的分片上传的任务，`-a`会列举出所有上传成功和为上传成功的objecct

  ```
  #列出所有的bucket
  root in /usr/local/etc λ oss ls
  CreationTime                                 Region    StorageClass    BucketName
  2020-06-19 15:39:17 +0800 HKT        oss-cn-beijing              IA    oss://gulied-program
  Bucket Number is: 1
  
  0.251266(s) elapsed      
  
  ---
  
  #列出指定bucket(oss://bucket-name)下的object
  root in ~ λ oss ls oss://gulied-program
  LastModifiedTime                   Size(B)  StorageClass   ETAG                                  ObjectName
  2020-06-20 17:57:49 +0800 HKT        56297            IA   6364821010871169E60E9F5187D5700F      oss://gulied-program/2020/06/20/558b4622e1f24a59bc8e2017608af8883.jpg
  2020-06-21 18:00:32 +0800 HKT        26775            IA   FE80C7644EFCFB2769413881F2831740      oss://gulied-program/2020/06/21/162f868b10274ad2bad5ce08446f0a65file.png
  2020-06-21 17:49:28 +0800 HKT       108269            IA   1D6D20E4D444D22EEBA92C11E26D6E9F      oss://gulied-program/2020/06/21/1a299fc3585e43d5871e16bc8efe1fdefile.png
  
  ---
  
  root in ~ λ oss ls oss://gulied-program -m
  UploadID Number is: 0
  
  0.128831(s) elapsed     
  ```

  ​	使用`-s`参数，展示object在bucket中的路径

  ```
  root in ~ λ oss ls oss://gulied-program -s
  oss://gulied-program/2020/06/20/558b4622e1f24a59bc8e2017608af8883.jpg
  oss://gulied-program/2020/06/21/162f868b10274ad2bad5ce08446f0a65file.png
  oss://gulied-program/2020/06/21/1a299fc3585e43d5871e16bc8efe1fdefile.png
  ```

  使用`-d`参数展示指定目录下的文件和文件夹，不会递归遍历

  ```
  root in ~ λ oss ls oss://gulied-program -d
  oss://gulied-program/avatar-boy.gif
  oss://gulied-program/default-tea-img.gif
  oss://gulied-program/template.xlsx
  oss://gulied-program/v-play-bg.jpg
  oss://gulied-program/课题一_V1.docx
  oss://gulied-program/2020/
  oss://gulied-program/banner/
  oss://gulied-program/course/
  oss://gulied-program/teacher/
  Object and Directory Number is: 9
  ```

  使用`--all-version`会查询object的version

  ```
  ./ossutil ls oss://bucket1 --all-versions
  LastModifiedTime                   Size(B)  StorageClass  ETAG                                  VERSIONID                                                           IS-LATEST   DELETE-MARKER  ObjectName
  2019-06-11 10:54:51 +0800 CST            0                                                      CAEQARiBgICUsOuR2hYiIDI3NWVjNmEyYmM0NTRkZWNiMTkxY2VjMDMwZjFlMDA3    true        true           oss://bucket1/test1.jpg
  2019-06-11 11:03:37 +0800 CST            0                                                      CAEQARiBgIDZtvuR2hYiIDNhYjRkN2M5NTA5OTRlN2Q4YTYzODQwMzQ4NDYwZDdm    true        true           oss://bucket1/test.jpg
  
  
  Object Number is: 4
  
  0.692000(s) elapsed
  ```

#### cat

https://help.aliyun.com/document_detail/120070.html?spm=a2c4g.11186623.6.865.7bf1509frhcTQC

cat和linux中的cat相同，用于查看文件中的具体内容

pattern；`cat [oss_url]`

```
root in /usr/local/\ λ oss cat oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/test.sh
#!/bin/bash
trap 'exit 0' SIGTERM
while true; do; done;
```

查看某一具体版本的文件使用`--version-id`

### Insert

#### mb

- `mb [option] [bucket-name]`

  用于创建bucket，bucketname是唯一标识符

  通过`--acl`指定bucket的访问权限

  1. private 缺省值
  2. public-read
  3. public-read-write

  ```
  root in ~ λ oss mb oss://$(uuidgen) --acl public-read
  
  0.531465(s) elapsed       
  ```

  通过`--storage-class`指定bucket的默认存储类型

  1. Standard：标准存储，缺省值
  2. IA：低频访问
  3. Archive：归档存储
  4. ColdArchive：冷归档存储

  ```
  root in ~ λ oss mb --storage-class IA oss://$(uuidgen)
  
  0.560694(s) elapsed                                                                                                                         /0.6s
  root in ~ λ oss ls
  CreationTime                                 Region    StorageClass    BucketName
  2021-02-09 10:57:08 +0800 HKT        oss-cn-beijing              IA    oss://af989cd8-85ae-4e1c-8614-c4371720112e
  
  ```

  通过`--redundancy-type`指定bucket的冗余存储类型

  1. LRS：本地冗余存储，缺省值
  2. ZRS：同城冗余存储

  ```
  root in ~ λ oss mb --redundancy-type ZRS  oss://$(uuidgen)
  
  0.464704(s) elapsed                                                                                                                         /0.5s
  root in ~ λ oss ls
  CreationTime                                 Region    StorageClass    BucketName
  2021-02-09 10:57:08 +0800 HKT        oss-cn-beijing              IA    oss://af989cd8-85ae-4e1c-8614-c4371720112e
  2021-02-09 10:58:35 +0800 HKT        oss-cn-beijing        Standard    oss://b5908e8d-e2b5-43ec-9dcd-cf73e22d2dd7
  
  ```

  通过`-e`指定创建bucket的region，如果没有指定默认使用配置文件的endpoint来指定bucket的region，这里使用region不是region id 而是 endpoint

  ```
  root in ~ λ oss mb -e oss-cn-hangzhou.aliyuncs.com oss://$(uuidgen)
  
  0.201905(s) elapsed                                                                                                                         /0.2s
  root in ~ λ oss ls
  CreationTime                                 Region    StorageClass    BucketName
  2021-02-09 11:07:17 +0800 HKT       oss-cn-hangzhou        Standard    oss://14938619-9e36-40f6-865a-cd013d2fb916
  ```

#### cp

> 这里只做简单的cp介绍，详细看CP block

- `cp [option] [local_url | filename] [oss_url | filename]`

   使用cp可以将本地文件上传至oss，需要指明文件名。如果想要上传文件夹使用`-r`参数。cp还可以用作下载。

  ```
  root in /usr/local/\ λ oss cp resizeApi.png oss://gulied-program/test.png
  Succeed: Total num: 1, size: 6,492. OK num: 1(upload 1 files).
  
  average speed 33000(byte/s)
  
  0.194656(s) elapsed   
  root in /usr/local/\ λ oss ls  oss://gulied-program/test.png
  LastModifiedTime                   Size(B)  StorageClass   ETAG                                  ObjectName
  2021-02-09 11:18:45 +0800 HKT         6492            IA   E75D4A1BDBD0B2F5EA79F073A68DA6A4      oss://gulied-program/test.png
  Object Number is: 1
  
  ```

#### appendfromfile

用于向oss中已存在的文件追加写入内容

```

```



### Delete

#### rm

- `rm [option] [bkname|objname]`

  用于删除object

  ```
  #删除空bucket，必须要添加-b参数
  root in ~ λ oss rm oss://9b6fa6d0-f073-42cf-879f-34a5871ac60b -b
  Do you really mean to remove the Bucket: 9b6fa6d0-f073-42cf-879f-34a5871ac60b(y or N)? y
  Removed Bucket: 9b6fa6d0-f073-42cf-879f-34a5871ac60b
  
  3.134278(s) elapsed  
  
  ---
  #删除bucket中的文件和bucket，必须要添加-bar参数
  root in ~ λ oss rm oss://9dd02231-cfff-4563-929a-a13268f40f16 -bar
  Do you really mean to remove recursively objects and multipart uploadIds of oss://9dd02231-cfff-4563-929a-a13268f40f16(y or N)? y
  Do you really mean to remove the Bucket: 9dd02231-cfff-4563-929a-a13268f40f16(y or N)? y
  Succeed: Total 0 objects, 0 uploadIds. Removed 0 objects, 0 uploadIds.
  Removed Bucket: 9dd02231-cfff-4563-929a-a13268f40f16
  
  7.342973(s) elapsed    
  
  ---
  #删除单个文件
  ./ossutil rm oss://bucket1/path/object
  ```

- `-r`

  删除指定前缀的object

  ```
  root in ~ λ oss ls oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c -d
  oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/test.sh
  oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/dt/
  oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/t1/
  oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/t2/
  Object and Directory Number is: 4
  
  0.104494(s) elapsed
  root in ~ λ oss rm -r oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/test
  Do you really mean to remove recursively objects of oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/test(y or N)? y
  Succeed: Total 1 objects. Removed 1 objects.
  
  1.982891(s) elapsed
  root in ~ λ oss ls oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c -d
  oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/dt/
  oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/t1/
  oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/t2/
  Object and Directory Number is: 3
  
  0.092097(s) elapsed
  ```

- `-m`

  选中指定未完成上传的object

- `-a`

  选中未完成上传和已完成上传的object

- `--include  | --exclude`

## CP

https://help.aliyun.com/document_detail/179388.html?spm=a2c4g.11186623.6.871.330b1f52CPomdz

cp还可以用作重命名文件(备份)，和linux中的cp类似，不过使用oss_url替换

### upload

> 上传文件时可以通过`--meta`参数对文件的默认acl做覆盖，默认acl继承bucket

普通上传

pattern：`cp [option] [local_url | filename] [oss_url | filename]`

```
root in /usr/local/\ λ oss cp resizeApi.png oss://gulied-program/test.png
Succeed: Total num: 1, size: 6,492. OK num: 1(upload 1 files).

average speed 33000(byte/s)

0.194656(s) elapsed   
root in /usr/local/\ λ oss ls  oss://gulied-program/test.png
LastModifiedTime                   Size(B)  StorageClass   ETAG                                  ObjectName
2021-02-09 11:18:45 +0800 HKT         6492            IA   E75D4A1BDBD0B2F5EA79F073A68DA6A4      oss://gulied-program/test.png
Object Number is: 1
```

- `-r`

  上传文件夹，如果oss中目录不存储，会自动创建

  ```
  root in /usr/local/\ λ oss cp -r docker_test oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c
  Succeed: Total num: 7, size: 300. OK num: 7(upload 5 files, 2 directories).
  
  average speed 0(byte/s)
  
  0.383812(s) elapsed
  
  root in /usr/local/\ λ oss ls -d oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c
  oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/test.sh
  oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/t1/
  oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/t2/
  Object and Directory Number is: 4
  
  0.114918(s) elapsed
  
  ```

- `-u`

  如果批量上传中的文件已经在oss中，在做上传时就会报错，可以使用`-u`参数，表示增量上传。

- `--meta`

  指定上传文件的meta信息，具体查看`set-meta`。通过该参数，可以指定文件的存储类型，acl，加密方式

- `--include | --exclude`

  批量上传时通过匹配条件匹配该文件是否需要上传，支持通配符`*`，`?`，`[..]`，`[!..]`。==按照参数给定的顺序匹配过滤==

  ```
  ./ossutil cp localfolder/ oss://examplebucket/desfolder/ --include "*.txt" -r
  ```

- `--only-current-dir`

  上传文件夹时，使用该参数指明表示只上传当前目录，而不上传子目录

  ```
  ./ossutil cp localfolder/ oss://examplebucket/desfolder/ --only-current-dir -r
  ```

- `--disable-dir-object`

  上传文件不为目录生成object

  ```
  ./ossutil cp localfolder/ oss://examplebucket/desfolder/ --disable-dir-object -r
  ```

- `--disable-all-symlink`

  上传时忽略链接

- `--snapshot-path <path>`‘

  上传文件的同时生成快照，并将快照放入到snapshot-path指定的路径

### download

普通下载，与上传类似，都有`-r`，`-u`，`snapshot-path`，`--include`，`--exlude`。如果想要下载到当前目录必须指明当前目录的绝对路径或`.`

pattern；`cp [oss_url | name] [local_url | filename] `

```
root in /usr/local/\ λ oss cp  oss://gulied-program/avatar-boy.gif ab.gif
Succeed: Total num: 1, size: 3,756. OK num: 1(download 1 objects).

average speed 16000(byte/s)

0.229897(s) elapsed                                                                                 
root in /usr/local/\ λ ls
ab.gif  complete.sh  docker_test  pid  resizeApi.png        
```

- `--range`

  指定下载文件中指定字节数的范围

  ```
  root in /usr/local/\ λ oss cp oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/test.sh  . --range=2-5
  Succeed: Total num: 1, size: 4. OK num: 1(download 1 objects).
  
  average speed 0(byte/s)
  
  0.132068(s) elapsed
  
  ```

- `--version-id`

  如果bucket开启版本控制，针对数据覆盖和删除操作会以历史记录保留，可以同通过该参数下载指定版本的文件

  ```
  ./ossutil cp oss://my-bucket/test.jpg localfolder/ --version-id  CAEQARiBgID8rumR2hYiIGUyOTAyZGY2MzU5MjQ5ZjlhYzQzZjNlYTAyZDE3MDRk
  ```

### copy

cp被用作上传和下载的同时，还支持不同bucket之间拷贝文件。都有`-r`，`-u`，`snapshot-path`，`--include`，`--exlude`。

pattern：`cp [options] [src_oss_url] [dest_oss_url]`	

```
root in /usr/local/\ λ oss cp oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/test.sh oss://gulied-program/test.sh
Succeed: Total num: 1, size: 56. OK num: 1(copy 1 objects).

average speed 0(byte/s)

0.338178(s) elapsed
```

### recover

同时ossutil还支持开启版本控制的文件恢复，和覆盖

```
./ossutil cp oss://examplebucket1/examplefile.txt oss://examplebucket2/ --version-id  CAEQARiBgID8rumR2hYiIGUyOTAyZGY2MzU5MjQ5ZjlhYzQzZjNlYTAyZDE3MDRk
```

## stat

查看object的信息，一般是meta-date，如果想要查看大小使用`oss du`命令

- 查看bucket信息

  ```
  root in ~ λ oss stat oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c
  Name              : 785e7ff7-42cb-43ad-b33f-54e0021fb88c
  Location          : oss-cn-beijing
  CreationDate      : 2021-02-09 12:47:54 +0800 HKT
  ExtranetEndpoint  : oss-cn-beijing.aliyuncs.com
  IntranetEndpoint  : oss-cn-beijing-internal.aliyuncs.com
  ACL               : private
  Owner             : 1787466735923101
  StorageClass      : Standard
  RedundancyType    : LRS
  ```

- 查看object信息

  ```
  root in ~ λ oss stat oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/dt/test.sh
  ACL                   : default
  Accept-Ranges         : bytes
  Cache-Control         : no-cache
  Content-Length        : 56
  Content-Md5           : jTCNQnl+z6ZNMegCFzF/jQ==
  Content-Type          : text/x-sh; charset=utf-8
  Etag                  : 8D308D42797ECFA64D31E80217317F8D
  Last-Modified         : 2021-02-09 15:43:04 +0800 HKT
  Owner                 : 1787466735923101
  X-Oss-Hash-Crc64ecma  : 5399636048623593122
  X-Oss-Object-Type     : Normal
  X-Oss-Storage-Class   : Standard
  ```

## du

查看bucket或是object的大小

```
root in ~ λ oss du oss://gulied-program
storage class   object count            sum size(byte)
----------------------------------------------------------
IA              46                      2458517
----------------------------------------------------------
total object count: 46                          total object sum size: 2458517
total part count:   0                           total part sum size:   0

total du size(byte):2458517

0.235545(s) elapsed

```

## set-acl

设置object（包括bucket和文件）的acl，有private，public-read，public-read-write

pattern：`./ossutil set-acl oss://bucket[/prefix] [acl] [-r] [-b] [-f] [-c file]`

- `-b`

  表示object为bucket

  ```
  ./ossutil set-acl oss://bucket1 private -b       
  ```

- `-r`

  递归遍历设置

## set-meta

用于设置已上传object的metadata(元信息，一般是修改请求时的请求头)。headers-key不区分大小写，但是headers-value区分大小写。

pattern：`oss set-meta [--update | --delete ] <oss_url>  [header:value#header:value...]`

当前可选的header

```
Headers:
      Expires(time.RFC3339:2006-01-02T15:04:05Z07:00)
      X-Oss-Object-Acl
      Origin
      X-Oss-Storage-Class
      Content-Encoding
      Cache-Control
      Content-Disposition
      Accept-Encoding
      X-Oss-Server-Side-Encryption
      Content-Type
      以及以X-Oss-Meta-开头的header（用户自定义的）
```

如果没有指定`--update`（表示只对某几个header-key进行更新，value可以为空）或`--delete`（删除指定header-key，value必须为空），默认更新header中所有的值，除了不可删除的header-key以外，其他的header-key都会被删除。可以通过`-r`参数指定对一个文件夹中所有的文件设置

```
root in ~ λ oss set-meta oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c/dt/test.sh cache-control:no-cache
Warning: --update option means update the specified header, --delete option means delete the specified header, miss both options means update the whole meta info, continue to update the whole meta info(y or N)? y
```

## referer

referer用于设置bucket防盗链(与http中的referer字段相同一般是一个host)，与restful请求类似。

- put

  pattern：`oss referer --method <method> <bucket_url> referer-value <rv>`

  用于添加和修改防盗链

  ```
  root in ~ λ oss referer --method put oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c referer-value cyberpelican [--disable-empty-referer]
  
  0.155093(s) elapsed
  ```

  1. referer-value：允许访问bucket的白名单，支持wildcard，`*`，`?`。多个域名可以用空格隔开
  2. `--disable-empty-referer`：表示防盗链不允许为空

- get

  pattern：`./ossutil referer --method get oss://bucket  [local_xml_file]`

  查看防盗链，可以将stdout保存成本地文件local_xml_file

  ```
  root in ~ λ oss referer --method get oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c
  <?xml version="1.0" encoding="UTF-8"?>
    <RefererConfiguration>
        <AllowEmptyReferer>true</AllowEmptyReferer>
        <RefererList>
            <Referer>kikochz</Referer>
            <Referer>referer-value</Referer>
        </RefererList>
    </RefererConfiguration>
  ```

- delete

  pattern：`./ossutil referer --method delete oss://bucket`

  删除防盗链

  ```
  root in ~ λ oss referer --method delete oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c
  
  0.153992(s) elapsed
  root in ~ λ oss referer --method get oss://785e7ff7-42cb-43ad-b33f-54e0021fb88c
  <?xml version="1.0" encoding="UTF-8"?>
    <RefererConfiguration>
        <AllowEmptyReferer>true</AllowEmptyReferer>
        <RefererList></RefererList>
    </RefererConfiguration>
  0.108986(s) elapsed
  ```

## lifecycle

https://help.aliyun.com/document_detail/122574.html?spm=a2c4g.11186623.6.880.23f13d68dP7orW

以xml的格式设置object的生命周期

## probe

https://help.aliyun.com/document_detail/120061.html?spm=a2c4g.11186623.6.887.3c91509fDh23IO

probe命令用于探测oss与本地的网络状况，探测结束后会生成一个日志文件

### 上传探测

pattern：`ossutil probe {--upload [file_name]} {--bucketname bucket_name} [--object object_name] [--addr domain_name] [--upmode]`

```
root in ~ λ oss probe --upload --bucketname gulied-program
begin parse parameters and prepare file...[√]
begin network detection...[√]
begin upload file(normal)...[√]

*************************  upload result  *************************
upload file:success
upload file size:122880(byte)
upload time consuming:221(ms)
(only the time consumed by probe command)


************************* report log info*************************
report log file:/root/logOssProbe20210209143847.log


```

- `--upload`

  指定探测的方式为上传探测。文件不会被真正的上传至oss

- file_name

  上传至bucket的本地路径，如果没有指定默认生成一个临时文件由于探测

- `--bucketname`

  指定需要探测的bucketname，不能带有`oss://`

- `--addr`

  默认值`www.aliyun.com`

- `--upmode`

  指定文件的上传方式

  1. normal：简单上传，缺省值
  2. append：追加上传
  3. multipart：分片上传

### 下载探测

pattern：`ossutil probe {--download} {--url http_url} [--addr=domain_name] [file_name]`

如果没有指定file_name。默认下载至当前路径，文件名与请求的文件名相同

```
root in ~ λ oss probe --download --url https://785e7ff7-42cb-43ad-b33f-54e0021fb88c.oss-cn-beijing.aliyuncs.com/t1/Dockerfile\?Expires\=1612853806\&OSSAccessKeyId\=TMP.3KjFoVS8YuDZdXpHa7HNFa88L1zh9tLA3LWadHpQeCCHWiBbgP3zoqYTzWScYktaycsQpnpQK4adS9B4WxYkEkSuNKgfqT\&Signature\=WZ1uGpq1UCxrt%2BTcxaPF4X9qaiU%3D
begin parse parameters and prepare object...[√]
begin network detection...[√]
begin download file...[√]

*************************  download result  *************************
download file:success
download file size:51(byte)
download time consuming:283(ms)
(only the time consumed by probe command)

download file is /root/Dockerfile

************************* report log info*************************
report log file:/root/logOssProbe20210209145326.log
```

- `--download`

  指定探测方式为下载探测

- `--url`

  探测的文件，可以从oss控制台中的url查看

### 探测指定项目

pattern：`./ossutil probe {--probe-item item_value} {--bucketname bucket-name} [--object object_name]`

- 探测上传带宽

  表示设备为双核，最大带宽320kbps

  ```
  root in ~ λ oss probe --probe-item upload-speed --bucketname gulied-program
  cpu core count:2
  parallel:2,average speed:126.90(KB/s),current speed:192.00(KB/s),max speed:320.00(KB/s)
  parallel:3,average speed:146.00(KB/s),current speed:96.00(KB/s),max speed:480.00(KB/s))
  ```

- 探测下载带宽

  ```
  root in ~ λ oss probe --probe-item download-speed --bucketname gulied-program --object v-play-bg.jpg
  cpu core count:2
  parallel:2,average speed:10189.00(KB/s),current speed:10391.78(KB/s),max speed:10391.78(KB/s)
  ```

## sign

https://help.aliyun.com/document_detail/120064.html?spm=a2c4g.11186623.6.896.69464812kserB7

用于生成访问object的签名的url，默认超时时间为60sec

pattern：`./ossutil sign oss://bucket/object [--timeout t] [--version-id versionId] [--trafic-limit limitSpeed] [--disable-encode-slash] [--payer requester]`

## getallpartsize

https://help.aliyun.com/document_detail/120068.html?spm=a2c4g.11186623.6.876.37591498uzHyzu

获取指定bucket中上传成功的multipart(会占用bucket空间)

## 补全脚本

参考：

https://github.com/scop/bash-completion/blob/master/completions/_svn#L74

```shell
#oss completion
_oss() {
    local cur pre words cword
    # shellcheck disable=SC2206
    words=(${COMP_WORDS[@]})
    cword="${COMP_CWORD}"
    cur="$2"
    pre="${words[cword - 1]}"
    local commands
    [[ "${commands}" ]] || commands="appendfromfile bucket-encryption bucket-policy bucket-tagging bucket-version \
  cat config cors cors-options cp create-symlink du getallpartsize hash help inventory \
  lifecycle listpart logging ls mb mkdir object-tagging probe read-symlink referer restore request-payment \
  revert-versioning rm set-acl set-meta sign stat sync update website worm help "

    local common
    [[ "${common}" ]] || common="-c --config-file -e --endpoint -i --access-key-id -k --access-key-secret -t --sts-token \
    --proxy-host --proxy-user --proxy-pwd --retry-times --loglevel "

    if ((${cword-} == 1)); then
        # shellcheck disable=SC2207
        COMPREPLY=($(compgen -W "$commands" -- "$cur"))
        return
    else
        local options
        case "$pre" in
        --loglevel)
            options="info debug"
            # shellcheck disable=SC2207
            COMPREPLY=($(compgen -W "$options" -- "$cur"))
            return
            ;;
        --acl)
            options="private public-read public-read-write"
            # shellcheck disable=SC2207
            COMPREPLY=($(compgen -W "$options" -- "$cur"))
            return
            ;;
        --storage-class)
            options="Standard IA Archive ColdArchive"
            # shellcheck disable=SC2207
            COMPREPLY=($(compgen -W "$options" -- "$cur"))
            return
            ;;
        -L | --language)
            options="CH EN"
            # shellcheck disable=SC2207
            COMPREPLY=($(compgen -W "$options" -- "$cur"))
            return
            ;;
        --redundancy-type)
            options="LRS ZRS"
            # shellcheck disable=SC2207
            COMPREPLY=($(compgen -W "$options" -- "$cur"))
            return
            ;;
        --encoding-type)
            options="url"
            # shellcheck disable=SC2207
            COMPREPLY=($(compgen -W "$options" -- "$cur"))
            return
            ;;
        --payer)
            options="requester"
            # shellcheck disable=SC2207
            COMPREPLY=($(compgen -W "$options" -- "$cur"))
            return
            ;;
        esac

        local command=${words[1]}
        if [[ "${cur-}" == -* ]]; then
            local options
            options="$common"
            case "$command" in
            cp)
                options+="-r --recursive -f --force -u --update --output-dir --bigfile-threshold --part-size --checkpoint-dir \
        --range --encoding-type --include --exclude --meta --acl -j --jobs --parallel --snapshot-path --disable-crc64 --payer \
        --maxupspeed --partition-download --version-id --local-host --enable-symlink-dir --only-current-dir --disable-dir-object \
        --disable-all-symlink --disable-ignore-error --tagging "
                ;;
            ls)
                options+="--payer -s --short-format -d --directory -m --multipart -a --all-type --limited-num --marker \
                --upload-id-marker --encoding-type --include --exclude --all-version --version-id-marker"
                ;;
            mb)
                options+="--proxy-host --proxy-user -L --language --acl --storage-class --redundancy-type"
                ;;
            rm)
                options+="-r --recursive -b --bucket -f --force -m --multipart -a --all-type --encoding-type --include \
                --exclude --version-di --all-versions --payer"
                ;;
            set-meta)
                options+="-r --recursive -u --update --delete -f --force --encoding-type --include --exclude -j --jobs -L \
        --language --output-dir --version-id"
                ;;
            esac
            # shellcheck disable=SC2207
            COMPREPLY=($(compgen -W "$options" -- "$cur"))
        else
            if [[ "${command-}" == @(help|[h?]) ]]; then
                # shellcheck disable=SC2207
                COMPREPLY=($(compgen -W "$commands" -- "$cur"))
            elif [[ "${command-}" == "set-meta" ]]; then
                metadata="Expires X-Oss-Object-Acl Origin X-Oss-Storage-Class Content-Encoding Cache-Control Content-Disposition Accept-Encoding X-Oss-Server-Side-Encryption Content-Type"
                # shellcheck disable=SC2207
                COMPREPLY=($(compgen -W "$metadata" -- "$cur"))
            else
                # shellcheck disable=SC2207
                COMPREPLY=($(compgen -f -d -- "$cur"))
            fi
        fi
    fi
}
complete -o filenames -o nospace -o bashdefault -F _oss oss

```

