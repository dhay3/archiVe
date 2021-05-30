# Maven中的Build标签

### [1] 基本元素

```
<build>

  		<finalName>maven-build-demo</finalName>
        <defaultGoal>install</defaultGoal>
        <directory>${basedir}/target</directory>
        <finalName>${artifactId}-${version}</finalName>
        <filters>
                <filter>filters/filter1.properties</filter>
        </filters>
         ...
</build>
```

- **finalName**

  指定打包文件名称(可用于出去jar文件版本号)

- defaultGoal

  执行build任务时, 如果没有指定目标, 将使用默认值

  如上配置: 在命令行中执行mvn, 则相当于执行mvn install

- directory

  build目标文件的存放目录, 默认在${basedire}/target , 即项目根目录下的target

- fileName

  build目标文件的名称, 默认情况为${artifactId}-${version}

- **filters**

  给出对资源文件进行过滤的属性文件的路径，

  默认位于${basedir}/src/main/filters/目录下。

  属性文件中定义若干了键值对，用于在构建过程中将资源文件中出现的变量（键）替换为对应的值。

  例如 properties文件中

  有name=value , 则resources文件中定义的${name}值就是value

  

### [2] Resource 配置

```
<build>
        ...
       <resources>
                  <resource>
                        <targetPath>META-INF/plexus</targetPath>
                        <filtering>false</filtering>
                        <directory>${basedir}/src/main/plexus</directory>
						<includes>
						<include>configuration.xml</include>
						</includes>
						<excludes>
						<exclude>**/*.properties</exclude>
						</excludes>
		 		</resource>
	</resources>
	<testResources>
		...
	</testResources>
	...
</build>
```

- resources

  对应项目的resource文件,可以配置多个项目资源

- targetPath

  指定build后的resource存放的文件夹, 默认是basedir. 

  通常被打包在jar中的resources的目标路径是META-INF

- **filtering**

  true/false, 表示该pom配置的filter是否激活

- directory

  定义resource文件所在的文件夹, 默认为${basedir}/src/main/resource

- includes

  包含内容(编译时仅复制包含的内容)

  只包括`${basedir}/src/main/plexus`目录下的`configuration.xml`

- excludes

  排除内容(编译时不复制指定排除的内容)

  这里表示排除`${basedir}/src/main/plexus`目录下所有匹配`**/*.properties`的文件

- testResourecs

  定义和resource类似, 只不过在test时使用

### [3]plugins配置

```
<build>
    ...
	<plugins>
		<plugin>
			<groupId>org.apache.maven.plugins</groupId>
			<artifactId>maven-jar-plugin</artifactId>
			<version>2.0</version>
			<extensions>false</extensions>
			<inherited>true</inherited>
			<configuration>
				<classifier>test</classifier>
			</configuration>
			<dependencies>...</dependencies>
			<executions>...</executions>
		</plugin>
    </plugins>
</build>
```

- GAV

  指定插件的标准坐标

- extensions

  是否加载plugin的extensions, 默认为false

- inherited

  true/false, 这个plugin是否应用到该pom的子pom, 默认为true

- configuration

  配置该plugin期望得到的properties

- dependencies

   作为plugin的依赖

### [4]占位符

configuration 

​		|-delimiters

​				|-delimiter

maven占位符默认的是${},也可以自己指定，如下：

```
<plugin>
    <groupId>org.apache.maven.plugins</groupId>
    <artifactId>maven-resources-plugin</artifactId>
    <version>2.5</version>
    <configuration>
        <useDefaultDelimiters>false</useDefaultDelimiters>
        <delimiters>
        <delimiter>$[*]</delimiter>
        </delimiters>
        <encoding>UTF-8</encoding>
    </configuration>
</plugin>
```

配置delimiter后, resources文件下的配置文件, 不再通过默认的${...}来取值,

而通过$[...]来取值

