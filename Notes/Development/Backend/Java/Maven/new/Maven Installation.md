# Maven Installation

Maven 是一个基于 Java 的 jar 包管理工具，所以在安装 Maven 之前需要先安装 Java

> 推荐使用 JDK 8

```
(base) cpl in ~ λ java -version
java version "19.0.1" 2022-10-18
Java(TM) SE Runtime Environment (build 19.0.1+10-21)
Java HotSpot(TM) 64-Bit Server VM (build 19.0.1+10-21, mixed mode, sharing)
```

有 2 种可选的方式

![](https://github.com/dhay3/image-repo/raw/master/20230801/2023-08-01_14-55.hduqlofjohc.webp)

1. Binary archive

   包含可以直接运行的文件

2. Source archive

   不包含可以直接运行的文件，需要自己编译

这里选择 Binary 建议下载完检查一下 Hash 和 Signature

```
#检查 Hash
(base) cpl in /sharing/env λ sha512sum apache-maven-3.9.3-bin.tar.gz

#校验 Signature
(base) cpl in /sharing/env λ gpg --verify apache-maven-3.9.3-bin.tar.gz.asc apache-maven-3.9.3-bin.tar.gz
gpg: Signature made Fri 23 Jun 2023 09:05:31 PM HKT
gpg:                using RSA key 29BEA2A645F2D6CED7FB12E02B172E3E156466E8
gpg: Can't check signature: No public key

#导入对应的公钥
(base) cpl in /sharing/env λ gpg --keyserver keys.openpgp.org --recv-keys 29BEA2A645F2D6CED7FB12E02B172E3E156466E8

#校验签名
(base) cpl in /sharing/env λ gpg --verify apache-maven-3.9.3-bin.tar.gz.asc apache-maven-3.9.3-bin.tar.gz         
gpg: Signature made Fri 23 Jun 2023 09:05:31 PM HKT
gpg:                using RSA key 29BEA2A645F2D6CED7FB12E02B172E3E156466E8
gpg: Good signature from "Tamas Cservenak (ASF) (Release key) <cstamas@apache.org>" [unknown]
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: 29BE A2A6 45F2 D6CE D7FB  12E0 2B17 2E3E 1564 66E8
```

Hash 和 Signature 都校验通过后，解压文件

```
(base) cpl in /sharing/env λ tar -zxvf apache-maven-3.9.3-bin.tar.gz
```

然后创建 soft link 到 executable file

```
(base) cpl in /sharing/env/apache-maven-3.9.3/bin λ sudo ln -s /sharing/env/apache-maven-3.9.3/bin/mvn /usr/local/bin/mvn
```

校验是否成功

```
(base) cpl in ~ λ mvn -v
Apache Maven 3.9.3 (21122926829f1ead511c958d89bd2f672198ae9f)
Maven home: /sharing/env/apache-maven-3.9.3
Java version: 19.0.1, vendor: Oracle Corporation, runtime: /sharing/env/jdk/jdk-19.0.1
Default locale: en_US, platform encoding: UTF-8
OS name: "linux", version: "6.1.38-1-manjaro", arch: "amd64", family: "unix"
```

**references**

1. [^https://maven.apache.org/guides/getting-started/maven-in-five-minutes.html]
2. [^https://maven.apache.org/ref/3.9.3/maven-settings/settings.html]