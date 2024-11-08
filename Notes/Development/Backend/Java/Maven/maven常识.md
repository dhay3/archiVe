# maven常识

[TOC]

### maven 常用命令

执行maven命令必须进入pom.xml所在的位置

`mvn clean` 清空产生的项目(target中)

`mvn compile` 编译源代码,生成字节码文件

`mvn package` 打包

`mvn install` 在本地repository中安装jar(包含mvc comile,mvc package,然后上传到本地仓库)

一个特殊的例子:

```java
mvn install:install-file
-DgroupId=com.aliyun 
-DartifactId=aliyun-sdk-vod-upload 
-Dversion=1.4.11 
-Dpackaging=jar 
-Dfile=aliyun-java-vod-upload-1.4.11.jar
```

将指定`GAV`坐标和绝对路径的包, 安装到本地仓库中, 前提是要在当前jar包的目录下运行, 

即`aliyun-java-vod-upload-1.4.11.jar`

<img src="../../imgs/_maven\2.png" style="zoom:60%;" />

`mvn deploy` 上传到私服(包含mvn install,然后上传到私服)

`mvn test` 运行测试

`mvn site` 产生site

`mvn test-compile` 编译测试代码

`mvn -Dtest package` 只打包不测试

`mvn jar:jar` 只打jar包

`mvn source.jar` 源码打包

`mvn dependency:tree` 以树型结构展示依赖

==mvn clean package 组合使用 clean + package==

### POM (Project Object Model 项目对象模型)

gav坐标, 在仓库中**唯一**定位一个maven模块  

[1] groupid: 公司或组织域名倒序+项目名

```
 <groupId>com.alibaba</groupId>
```

[2] artifactid: 模块名

```
<artifactId>druid-spring-boot-starter</artifactId>
```

[3]version: 版本号

```
<version>0.0.1</version>
```

maven模块的坐标与仓库中路径的对应关系

```
<groupId>com.alibaba</groupId>
<artifactId>druid-spring-boot-starter</artifactId>
<version>1.1.17</version>
D:\maven\repository\com\alibaba\druid-spring-boot-starter\1.1.17
```

### 仓库

​		[1]本地仓库:当前电脑上部署的仓库目录,为当前电脑上所有Maven模块服务

​		[2]远程仓库

​							|-私服: 搭建在局域网环境 nexus

​							|-中央仓库: 架设在internet上,为所有maven模块服务

​							|-中央仓库镜像: 分担中央仓库流量

​		

### 依赖

##### 依赖排除

也可以通过pom.xml的Diagrams排除依赖

```
 	  <dependency>
            <groupId>com.baomidou</groupId>
            <artifactId>mybatis-plus</artifactId>
            <exclusions>
                <exclusion>
                    <groupId>com.baomidou</groupId>
                    <artifactId>mybatis-plus-annotation</artifactId>
                </exclusion>
            </exclusions>
            <version>3.3.0</version>
        </dependency>
```

#####  依赖的范围

`<scope> </scope>`

默认scope compile

|                    | compile | test   | provided |
| :----------------- | ------- | ------ | -------- |
| 对主程序是否有效   | 有效    | 无效   | 有效     |
| 对测试程序是否有效 | 有效    | 有效   | 有效     |
| 是否参于打包       | 参与    | 不参与 | 不参与   |

- complie

  是默认值，表示在build,test,runtime阶段的classpath下都有依赖关系。

- test

  表示只在test阶段有依赖关系，例如junit

- provided

  表示在build,test阶段都有依赖，在runtime时并不输出依赖关系而是由容器提供，例如web war包都不包括servlet-api.jar，而是由tomcat等容器来提供, 用于像servlet-api 开发环境要用, 但是tomcat中存在servlet-api, 所以不用打包以免冲突

- runtime
  
  表示在构建编译阶段不需要(参于打包)，只在test和runtime需要。这种主要是指代码里并没有直接引用而是根据配置在运行时动态加载并实例化的情况。虽然用runtime的地方改成compile也不会出大问题，但是runtime的好处是可以避免在程序里意外地直接引用到原本应该动态加载的包。例如JDBC连接池

### optional

C继承B, B继承A

当B声明optional为true时, C如果不显示声明继承A, 就不会继承A的依赖

如果为false统统都会被继承

### 生命周期 LifeCycle

 各个构建环节的顺序

[1] 清理: 将以前编译得到的旧的class字节码文件删除, 为下一次编译做准备

[2] 编译: 将Java源程序编成class字节码文件

[3]测试: 自动测试, 自动调用junit程序

[4]报告: 测试程序执行的结果

[5]打包: 动态web模块打war包, java模块打jar包

[6]安装: maven特定的概念----将打包得到的文件复制到仓库中指定位置

[7]部署: 将动态web模块生成的war包复制到servlet容器的指定目录下,使其可以运行

<img src=".\img\1.PNG" style="zoom:67%;" />

### 统一管理依赖的版本

在properties标签使用自定义标签来声明版本号

在需要统一版本的位置,使用${自定义标签名}来引用版本号

```
    <properties>
        <spring.version>4.3.9.RELEASE</spring.version>
        <mybatis.version>3.4.5</mybatis.version>
        <mysql.version>8.0.17</mysql.version>
    </properties>
    
    <dependency>
            <groupId>org.springframework</groupId>
            <artifactId>spring-core</artifactId>
            <version>${spring.version}</version>
     </dependency>
      <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
            <version>${mysql.version}</version>
       </dependency>
```

### 继承

- 统一管理版本号

 通过parent标签来引入模块的pom文件, maven和java一样, 单继承, 一个子model中只能出现一个parent标签

一般创建一个parent项目打包类型为**==pom==**,parent中不放任何代码, 子模块通过parent标签来引入父类模块的依赖,

子模块能调用父模块的所有依赖

```
   <parent>
        <artifactId>m03</artifactId>
        <groupId>org.example</groupId>
        <version>1.0-SNAPSHOT</version>
        <!--以当前项目为基准,父模块所在的位置-->
        <relativePath>../m03/pom.xml</relativePath>
    </parent>
```

- 如果项目中有子模块只需要父类依赖中的部分,父模块可以通过`dependencyManagement`标签来统一管理, 子模块要==显示的==声明`dependencyManagemengt`中依的赖才能引入, 否则不会
- 子模块中不需要声明`version`和`group id`,和父模块的相同; 如果子模块中声明了版本, 那就使用子模块就使用  自己的版本号

```
	<!--父模块-->
	<!--父模块的打包方式为pom-->
	<packaging>pom</packaging>
	<--!modules是引用用了当前模块的子模块-->
	<modules>
        <module>../m04</module>
    </modules>
	<dependencyManagement>
        <dependencies>
            <dependency>
                <groupId>org.slf4j</groupId>
                <artifactId>slf4j-log4j12</artifactId>
                <version>1.7.30</version>
                <scope>test</scope>
            </dependency>
        </dependencies>
    </dependencyManagement>
    <!--子模块-->
    <!--groupid与version父模块相同不用指明-->
    <artifactId>m04</artifactId>
    <packaging>jar</packaging>
    <dependency>
    	<groupId>org.slf4j</groupId>
    	<artifactId>slf4j-log4j12</artifactId>
    </dependency>
```

==配置继承后执行install(安装命令)要先安装父模块==

### 可被继承的POM元素

- groupId

- version

- description 

  项目大的描述信息

- organization 

  项目的组织信息

- url 

  项目的url地址

- developers 

  项目开发者信息

- properties

  自定义的maven属性

- dependencies

  项目的依赖配置

- dependencyManagement

  项目的依赖管理配置

- repositories

  项目的仓库配置

- build

  包括项目的源码目录配置, 输出目录配置, 插件配置, 插件管理配置

### relativePath

表示父项目相对当前项目的依赖位置, 默认 `../pom.xml`

### dependencyManagement的特殊点, `<scope>import</scope>`

- 假如说，我不想继承所有，或者想继承多个，怎么做？

- 如果10个、20个甚至更多模块继承自同一个模块，那么按照我们之前的做法，这个父模块的dependencyManagement会包含大量的依赖(==且不推荐这么做,只引入当前子模块所需的依赖即可,不引入多余的依赖==)。如果你想把这些依赖分类以更清晰的管理，那就不可能了，`<import>scope</scope>`依赖能解决这个问题。你可以把dependencyManagement放到单独的专门用来管理依赖的pom中，然后在需要使用依赖的==子模块中的dependcyManagement通过`<scope>import</scope>`==，就可以不使用parent标签引入==父模块中的dependencyManagement==。

- 例如可以写这样一个用于依赖管理的pom：

```
<project>
    <modelVersion>4.0.0</modelVersion>
    <groupId>com.test.sample</groupId>
    <artifactId>base-parent1</artifactId>
    <packaging>pom</packaging>
    <version>1.0.0-SNAPSHOT</version>
    <dependencyManagement>
        <dependencies>
            <dependency>
                <groupId>junit</groupId>
                <artifactid>junit</artifactId>
                <version>4.8.2</version>
            </dependency>
            <dependency>
                <groupId>log4j</groupId>
                <artifactid>log4j</artifactId>
                <version>1.2.16</version>
            </dependency>
        </dependencies>
    </dependencyManagement>
</project>
```

- 然后我就可以通过非继承的方式来引入这段依赖管理配置

```
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>com.test.sample</groupId>
            <artifactid>base-parent1</artifactId>
            <!--版本号要与父模块中的一样-->
            <version>1.0.0-SNAPSHOT</version>
            <type>pom</type>
            <scope>import</scope>
        </dependency>
    </dependencies>
</dependencyManagement>

<dependency>
    <groupId>junit</groupId>
    <artifactid>junit</artifactId>
</dependency>
<dependency>
    <groupId>log4j</groupId>
    <artifactid>log4j</artifactId>
</dependency>

```

==**注意**:  `<scope>import</scope>`只能在dependencyManagement中, 且type=pom, version与父模块中定义的版本号一致==

### 聚合

一键安装或打包多个模块

在总聚合模块中, 配置各个参于聚合的模块

```
 	<!--各个模块的相对路径-->
 	<modules>
 		<module>../m03</module>
        <module>../m04</module>
    </modules>
```

参考:

https://blog.csdn.net/mn960mn/article/details/50894022

https://www.cnblogs.com/huahua035/p/7680607.html

https://segmentfault.com/a/1190000013145264
