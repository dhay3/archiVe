# Maven pom.xml

Project Object Model(POM) 是 Maven 中最基础的 Unit，以 XML 的格式编写，在项目中通常以 `pom.xml` 来命名

一个最基础的 `pom.xml` 需要包含如下几部分

1. `porject`

   根标签

2. `modelVersion` should be set to `4.0.0`

   ```
   <modelVersion>4.0.0</modelVersion>
   ```

3. `groupId` the id of the project’s group

   通常是公司或者组织域名倒序

   ```
   <groupId>com.mycompany.app</groupId>
   ```

4. `artifactId` the id of the artifact(project)

   模块名

   ```
   <artifactId>my-app</artifactId>
   ```

5. `version` the version of the artifact under the specified group

   版本号

   ```
   <version>1</version>
   ```

例如一个最小的 `pom.xml`

```
<project>
	<modelVersion>4.0.0</modelVersion>
  <groupId>com.alibaba</groupId>
  <artifactId>druid-spring-boot-starter</artifactId>
  <version>0.0.1</version>
</project>
```

## GAV

其中 `groupId`, `artifactId`, `version` 也被称为 GAV 坐标，可以构成项目或者模块的唯一标识符 `<groupId>:<artifactId>:<version>`

例如

```
<!-- https://mvnrepository.com/artifact/org.mybatis/mybatis -->
<dependency>
    <groupId>org.mybatis</groupId>
    <artifactId>mybatis</artifactId>
    <version>3.5.13</version>
</dependency>
```

唯一标识符就是 `org.mybatis:mybatis:3.5.13`

## [POM 标签](https://maven.apache.org/ref/3.9.3/maven-model/maven.html)

常见的标签有

- `<packaging/>`

  打包的方式，可以是 

  1. `jar`
  2. `war`
  3. `pom`
  4. `ear`

- `<dependencies/>`
- `<parent>`
- `<dependencyManagement>`
- `<modules>`
- `<properties>`



**references**

1. [^https://maven.apache.org/guides/introduction/introduction-to-the-pom.html]
2. [^https://maven.apache.org/pom.html]
3. [^https://www.runoob.com/maven/maven-pom.html]