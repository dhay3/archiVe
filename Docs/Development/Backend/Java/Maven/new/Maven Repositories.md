# Maven Repositories

Repository 是一个存储第三方 artifacts(jar 包) 的集合，主要分为两种

1. local

   本地仓库

2. remote

   远程仓库

## 本地仓库

本地仓库就像一个缓存，当你从远程仓库中下载依赖时，会拷贝一份到本地仓库。如果之后需要调用对应的依赖，你不必再次从远程仓库下载依赖

默认存储在 `~/.m2/repository`，可以通过 `<localRepository/>` 标签来修改默认的存储位置

```
    <settings>
      ...
      <localRepository>/path/to/local/repo/</localRepository>
      ...
    </settings>
```

## 远程仓库

不是架设在本地的仓库，通过网络下载 artifacts

例如 https://mvnrepository.com 就是一个远程仓库

```
    <settings>
      ...
      <mirrors>
        <mirror>
          <id>other-mirror</id>
          <name>Other Mirror Repository</name>
          <url>https://other-mirror.repo.other-company.com/maven2</url>
          <mirrorOf>central</mirrorOf>
        </mirror>
      </mirrors>
      ...
    </settings>
```

## 查找顺序

如果需要使用 artifact，maven 首先会从本地仓库寻找，如果存在就使用本地仓库中的 artifact，如果不存在就会从远程仓库寻找并将其拷贝到本地仓库

```
if artifact in local repository:
	search local repository
else
	search remote repository
	copy to local repository
```

```mermaid
flowchart LR
Maven项目-->本地仓库 --> 远程仓库
```

**referneces**

1. [^https://maven.apache.org/guides/introduction/introduction-to-repositories.html]