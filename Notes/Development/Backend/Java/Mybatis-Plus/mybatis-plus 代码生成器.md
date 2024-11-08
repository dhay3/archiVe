# mybatis-plus 代码生成器

```java
    @Test
    public void getGenerator() {
        //项目路径
        System.out.println(System.getProperty("user.dir"));

        //代码生成器
        AutoGenerator autoGenerator = new AutoGenerator();

        //1.全局配置 调用generator.config下的
        GlobalConfig gc = new GlobalConfig();
        //获取当前项目的路径
        String path = System.getProperty("user.dir");
        //设置是否开启AR
        gc.setAuthor("chz")
                //文件输出路径
                .setOutputDir(path + "/src/main/java")
                //生成时是否打开文件
                .setOpen(false)
                //是否覆盖文件
                .setFileOverride(true)
                //设置主键自增策略
                .setIdType(IdType.ASSIGN_UUID)
                //DO中日期类的类型, 缺省值为TIME_PACK
                .setDateType(DateType.TIME_PACK)
                //是否开启resultMap,默认false
                .setBaseResultMap(true)
                //是否开启sql片段,默认false
                .setBaseColumnList(true)
                //去掉service接口首字母的I, 如DO为User则叫UserService
                .setServiceName("%sService")
                //开启Swagger2
                .setSwagger2(true);


        //2.数据源配置
        DataSourceConfig dataSourceConfig = new DataSourceConfig();
        //设置数据源类型
        dataSourceConfig.setDbType(DbType.MYSQL)
                .setDriverName("com.mysql.cj.jdbc.Driver")
                .setUrl("jdbc:mysql://localhost:3306/gedu?useUnicode=true&characterEncoding=utf-8&serverTimezone=Asia/Shanghai")
                .setUsername("root")
                .setPassword("12345");

        //3.策略配置
        StrategyConfig strategyConfig = new StrategyConfig();
        //是否开启大写命名,默认不开启
        strategyConfig.setCapitalMode(false)
                //生成的DO自动实现Serializable接口, 默认true
                .setEntitySerialVersionUID(true)
                //数据库 表 映射到实体类命名策略
                .setNaming(NamingStrategy.underline_to_camel)
                //数据库表 字段 映射到实体类的命名策略
                .setColumnNaming(NamingStrategy.underline_to_camel)
                //设置想要生成的表
                .setInclude("edu_teacher")
                //生成的dao,service,entity不再带tbl_前缀
                .setTablePrefix("edu_")
                //设置lombok, @Accessor(chain = true),@Data等
                .setEntityLombokModel(true)
                //controller使用@RestController
                .setRestControllerStyle(true)
                //Mapping驼峰转连字
                .setControllerMappingHyphenStyle(true)
                //自动填充字段
                .setTableFillList(Arrays
                        .asList(new TableFill("gmt_create", FieldFill.INSERT),
                                new TableFill("gmt_modified",FieldFill.INSERT_UPDATE)))
//                .setVersionFieldName("")//乐观锁属性名
                //表中字段为is_deleted,生成的DO中去掉is前缀
                .setEntityBooleanColumnRemoveIsPrefix(true)
              .setLogicDeleteFieldName("deleted");//逻辑删除属性名


        //4.包配置
        PackageConfig packageConfig = new PackageConfig();
        //setParent设置统一的包路径
        //设置模块名,对应controller中使用servicedu作为url, 如@RequestMapping("/servicedu/teacher"), 所有生成的都会在以该模块名为的包下
        packageConfig.setModuleName("servicedu")
                .setParent("com.chz")
                .setMapper("mapper")
                .setService("service")
                .setController("controller")
                .setEntity("entity")
                .setXml("mapper");

        //整合配置
        autoGenerator.setPackageInfo(packageConfig)
                .setDataSource(dataSourceConfig)
                .setGlobalConfig(gc)
                .setStrategy(strategyConfig);
        //执行
        autoGenerator.execute();
    }
```

生成的项目结构如下 

<img src="D:\java资料\我的笔记\mp\1.png" style="zoom:80%;" />

controller

```java
@Controller
@RequestMapping("/servicedu/teacher")
public class TeacherController {

}
```

