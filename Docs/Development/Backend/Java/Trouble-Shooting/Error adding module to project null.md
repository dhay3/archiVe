# Error adding module to project: null

如果 JDK 版本高于 Maven 适配的版本，在创建 Maven 模块时就可能出现 `Error adding module to project: null` 的错误信息

![img](https://img-blog.csdnimg.cn/83b4547df2f24609bff3db77cfbf210f.png)

将 Project 版本 JDK downgrade 到 Maven 对应的版本即可 

```
mvn --version
Apache Maven 3.9.3 (21122926829f1ead511c958d89bd2f672198ae9f)
Maven home: /sharing/env/maven/apache-maven-3.9.3
Java version: 19.0.1, vendor: Oracle Corporation, runtime: /sharing/env/jdk/jdk-19.0.1
Default locale: en_US, platform encoding: UTF-8
OS name: "linux", version: "6.1.38-1-manjaro", arch: "amd64", family: "unix"
```

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230802/2023-08-02_11-16.200vakrqldeo.webp)

**references**

1. [^https://blog.csdn.net/weixin_45993334/article/details/122454009]