### #为什么要使用gradle

不同开发人员, 使用不同的依赖库(maven version), 就有可能造成版本的不兼容. 所以要引入gradle, 来统一管理

版本库

### #gradle 概念

- distribution

  gradle的版本

  -  `bin`只包括二进制文件
  - `src`只包含源代码, 不能编译你的项目
  - `all`完整文件, 包含bin和src

- wrapper 

  用于解决不同项目使用不同gradle版本的问题, 如果本地有对应版本的gradle, 就不会从gradle官网下载,

  而是用本地仓库中的graddle

  例如:

  一个项目使用的gradle版本是5.2.1, 但是本地的gradle低于或高于此版本库. 那么就可以通过

  `distributionUrl`来指定版本

  `gradlew` 用于linux和mcos

  `gradlew.bat`用于windows

  通过wrapper下载的gradle会默认放在`~/.gradle/wrapper/dists/`

<img src="..\..\..\..\java资料\我的笔记\gradle\1.png"/>

  通过修改idea自动生成的`gradle-wrapper.properites`

  - **distributionUrl** 对应gradle下载的版本,
  - **distributionBase**, **distributionPath**下载并解压后的gradel存放的目录和路径
  - **zipStorePath**,**zipStoreBase**压缩文件存放的gradel存放的目录和路径

- gradle user home

  对应`~/.gradle`, 可以通过idea修改gradle user home

  `caches/modules-2`对应maven的`.m2`文件用来存储下载好的jar包

  ==**注意**==如果使用`mavenLocal()`(maven的本地仓库, 只需要修改maven中的setting.xml即可)并不会直接复用, 而是复制一份到`~/.gradle/caches/modules-2/files-2.1/`目录下

- daemon

  能有效加快编译, 3.0 后自动开启

### #gradle添加aliyun运程仓库

1. 单个项目配置

   ```gradle
   repositories {
       jcenter()
       //先从本地仓库寻找,找不到然后上aliyun,然后到maven
       //本地仓库
       mavenLocal()
       //配置阿里云仓库
       maven {
           url 'http://maven.aliyun.com/nexus/content/groups/public/'
       }
       mavenCentral()
   }
   ```

2. 全局配置

   在gradle安装目录`init.d`下新建`init.gradle`内容如下

   ```
   allprojects{
       repositories {
           def REPOSITORY_URL = 'http://maven.aliyun.com/nexus/content/groups/public/'
           all { ArtifactRepository repo ->
               if(repo instanceof MavenArtifactRepository){
                   def url = repo.url.toString()
                   if (url.startsWith('https://repo1.maven.org/maven2') || url.startsWith('https://jcenter.bintray.com/')) {
                       project.logger.lifecycle "Repository ${repo.url} replaced by $REPOSITORY_URL."
                       remove repo
                   }
               }
           }
           maven {
               url REPOSITORY_URL
           }
       }
   }
   
   ```

### #settings.gradle

```groovy
rootProject.name = 'gradle-test' //项目名
include 'gradle-model1' //子模块的name,即artifact
include 'gradle-model1'
```

### #build.gradle

1.  implementation 等同于compile中添加 optional = true

   maven

   ```xml
   <dependencies>
       <dependency>
           <groupId>log4j</groupId>
           <artifactId>log4j</artifactId>
           <version>1.2.12</version>
       </dependency>
   </dependencies>
   ```

   gradle

   ```groovy
   dependencies {
       implementation 'log4j:log4j:1.2.12'  
   }
   ```

   

### #build.gradle

- Build script structure

  | Block                | Description                                                  |
  | -------------------- | ------------------------------------------------------------ |
  | `allprojects { }`    | Configures this project and each of its sub-projects.        |
  | `artifacts { }`      | Configures the published artifacts for this project.         |
  | `buildscript { }`    | Configures the build script classpath for this project.      |
  | `configurations { }` | Configures the dependency configurations for this project.   |
  | `dependencies { }`   | Configures the dependencies for this project.                |
  | `repositories { }`   | Configures the repositories for this project.                |
  | `sourceSets { }`     | Configures the source sets of this project.                  |
  | `subprojects { }`    | Configures the sub-projects of this project.                 |
  | `publishing { }`     | Configures the [`PublishingExtension`](https://docs.gradle.org/current/dsl/org.gradle.api.publish.PublishingExtension.html) added by the publishing plugin. |

对应maven的`pom.xml`

```groovy
//指定编译语言环境
plugins {
    id 'java'
}
//项目gav坐标
group 'com.chz'
version '1.0-SNAPSHOT'
//jdk版本
sourceCompatibility = 1.8

repositories {
    jcenter()
    //先从本地仓库寻找,找不到然后上aliyun,然后到maven
    //maven本地仓库
    mavenLocal()
    //单项目,配置阿里云仓库
    maven {
        url 'http://maven.aliyun.com/nexus/content/groups/public/'
    }
    mavenCentral()
}

dependencies {
    testCompile group: 'junit', name: 'junit', version: '4.12'
    // https://mvnrepository.com/artifact/org.springframework/spring-context
    compile group: 'org.springframework', name: 'spring-context', version: '5.2.6.RELEASE'
}
```

- compile: 所有的子模块可以方法, 相当于maven默认的scope compile
- api:  等同于compile
- providedCompile: 相当于provided
- compileOnly:  只编译, 不会打包, 相当于maven的provided
- implementation: 相当于denpendencies managerment
- testCompile: 相当于scope test
- testImentation: 详单与scope test 只不过定义在dependencyManagement

| 新配置                         | 已弃用配置 | 描述                             |
| ------------------------------ | ---------- | -------------------------------- |
| api                            | compile    | 相当于maven默认的scope = compile |
| compileOnly                    | provided   |                                  |
| testCompile/testImplementation |            | 相当于maven的scope= test         |
|                                |            |                                  |
|                                |            |                                  |
|                                |            |                                  |
|                                |            |                                  |
|                                |            |                                  |
|                                |            |                                  |

