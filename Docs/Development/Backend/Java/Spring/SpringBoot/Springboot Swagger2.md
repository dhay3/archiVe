# Swagger2

### #常见注解

- @Api : 修饰整个类, 描述controller的作用

- @ApiOperation : 描述一个类的一个方法, 或者说一个接口
- @ApiParam:  单个参数描述
- @ApiModel:  用在放回对象上，描述对象
- @ApiProperty: 用对象来接收参数时, 描述对象的一个字段
- @ApiResponse: Http响应其中1个描述
- @ApiIgnore: 使用该注解忽略这个API
- @ApiError: 发生错误返回的信息

- @ApiImplicitParam: 一个请求参数
- @ApiImplicitParmas: 多个请求参数

### #依赖

```xml
        <dependency>
            <groupId>io.springfox</groupId>
            <artifactId>springfox-swagger2</artifactId>
            <version>2.9.2</version>
        </dependency>
        <!--可视化界面-->
        <dependency>
            <groupId>io.springfox</groupId>
            <artifactId>springfox-swagger-ui</artifactId>
            <version>2.9.2</version>
        </dependency>
		<!--二选一-->
		<dependency>
            <groupId>com.github.xiaoymin</groupId>
            <artifactId>swagger-bootstrap-ui</artifactId>
            <version>1.9.6</version>
        </dependency>
```

springfox-swagger-ui 地址 `http://localhost:8080/swagger-ui.html`

swagger-bootstrap-ui 地址 `http://localhost:8080/doc.html`

### #配置类

```java
/**
 * swagger2 配置类
 * 开启swagger2
 */
@Configuration
@EnableSwagger2
public class SwaggerConf {
    //版本
    public static final String VERSION = "1.1";
    /*
    这里需要注入bean
     */
    @Bean
    public Docket createRestApi() {
        return new Docket(DocumentationType.SWAGGER_2)
                .apiInfo(apiInfo())
            	//会自动扫描当前模块下使用了Swagger2注解的类
                .select()
                //这些都是springboot中默认的一些url, 需要忽略
                .paths(Predicates.not(PathSelectors.regex("/admin/.*")))
                .paths(Predicates.not(PathSelectors.regex("/error.*")))
                .build();
    }

    /*
     设置api的信息 ,这些信息会展示在文档中
     */
    private ApiInfo apiInfo() {
        return new ApiInfoBuilder()
                //设置文档标题
                .title("swagger first doc")
                //设置文档描述
                .description("api 接口文档")
                .contact(new Contact("chz", "https://www.chz.com", "kikochz@163.com"))
                .version(VERSION)
                .build();
    }
```

### #entity

```java
@ApiModel(value = "员工")
@Data
@TableName("tbl_employee")
public class Employee implements Serializable {

    private static final long serialVersionUID = 1L;
    @ApiModelProperty(value = "id")
    @TableId(value = "id", type = IdType.AUTO)
    private Integer id;
    @ApiModelProperty(value = "姓名")
    private String lastName;
    @ApiModelProperty(value = "邮箱")
    private String email;
    @ApiModelProperty(value = "性别")
    private String gender;
    @ApiModelProperty(value = "年龄")
    private Integer age;
    @ApiModelProperty(value = "乐观锁")
    @Version
    private Integer version;
    @ApiModelProperty(value = "逻辑删除")
    @TableLogic
    private Integer deleted;
    @ApiModelProperty(value = "日期")
    @DateTimeFormat(pattern = "yyyy-MM-dd")
    @JsonFormat(pattern = "yyyy-MM-dd", timezone = "GTM+8")
    private LocalDateTime date;
}
```

### #controller

```java
@Api("employee controller")
@RestController
@RequestMapping("/employee")
public class EmployeeController {
    @Autowired
    private IEmployeeService employeeService;

    @ApiIgnore//忽略该api
    @RequestMapping("/hello")
    public String hello() {
        return "hello world";
    }

    /*
    value是对方法的描述, notes是注意点
     */
    @ApiOperation(value = "查询所有", notes = "注意实体类有乐观锁")
    @GetMapping("/list")
    public List<Employee> list() {
        return employeeService.list();
    }

    @ApiOperation(value = "查询用户", notes = "通过id查询")
    @GetMapping("/list/{id}")
    public Employee get(@PathVariable("id") Integer id) {
        return employeeService.getById(id);
    }

    @ApiOperation(value = "添加用户", notes = "传过来的参数可以不一样")
    @PostMapping("/add")
    public boolean add(@RequestBody Employee employee) {
        return employeeService.save(employee);
    }

    /*
    name 参数名
    value 参数的简要描述
    required 参数是否必须
    dataTypeClass 参数的数据类型
    parameterType http参数类型
     */
    @ApiImplicitParam(name = "id", value = "用户id", required = true, dataType = "Integer", paramType = "path")
    @DeleteMapping("/{id}")
    public boolean delete(@PathVariable("id") Integer id) {
        return employeeService.removeById(id);
    }

    @ApiOperation(value = "更新用户", notes = "根据用户id更新用户")
    @ApiImplicitParams({
            @ApiImplicitParam(name = "id", value = "用户id", required = true, dataTypeClass = Integer.class, paramType = "path"),
            @ApiImplicitParam(name = "employee", value = "更新的参数", dataTypeClass = Employee.class)

    })
    @PutMapping("/{id}")
    public boolean update(@PathVariable("id") Integer id, @RequestBody Employee employee) {
        return employeeService.update(employee,
                new UpdateWrapper<Employee>().eq("id", id));
    }
```

### #最后效果

<img src="D:\java资料\我的笔记\springboot\img\3.PNG" alt="3" style="zoom:60%;" /><img src="D:\java资料\我的笔记\springboot\img\5.PNG" alt="5" style="zoom:60%;" /><img src="D:\java资料\我的笔记\springboot\img\6.PNG" alt="6" style="zoom:60%;" />
