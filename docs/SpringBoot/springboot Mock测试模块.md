# springboot Mock测试模块

[TOC]

如果测试方法较多, 推荐使用静态导包

##### #方法一:

可用全局的controller

```java
    private MockMvc mockMvc;
    @BeforeEach
    public void setup(WebApplicationContext wac) {
        this.mockMvc = MockMvcBuilders.webAppContextSetup(wac).build();
    }

```

##### #方法二:

细粒度更高的, 指定测试controller

==**这里要注意的是 一定要注入controller**==

```java
    private MockMvc mockMvc;
    //controller中调用了Service, 必须要注入controller,否则会出现空指针
    @Autowired
    private EmployeeController employeeController;
    @BeforeEach
    public void setup() {
        //设置全局
        mockMvc = MockMvcBuilders.standaloneSetup(
                employeeController)
                .alwaysExpect(MockMvcResultMatchers.status().isOk())
                .alwaysDo(MockMvcResultHandlers.print())
                .build();
    }
```

##### #方法三:

```java
//默认使用80端口,需要修改为RANDOM_PORT否则会报空指针
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
@AutoConfigureMockMvc
public class MockTest3 {
    @Autowired
    MockMvc mockMvc;

    @Test
    public void get() throws Exception {
        mockMvc.perform(MockMvcRequestBuilders.post("/employee/get").param("name", "Black"))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andDo(MockMvcResultHandlers.print());
    }
}
```



##### #注意点:

1. `mock`返回的`responseBody`存在中文乱码, 需要通过`produces`来指定`charset`,如下

   ```java
       @BeforeEach
       public void setup() {
           mockMvc = standaloneSetup(teacherController)
                   //alwaysExpect只针对response
                   .alwaysExpect(status().isOk())
                   //mock返回的数据存在中文乱码, 要通过produces指定charset
                   .alwaysExpect(content().contentType("application/json;charset=utf-8"))
                   .alwaysDo(print())
                   .build();
       }
   ```

2. 如果前端发送的请求头的`Content-Type` 不是 默认的请求头要在测试的request中指明 如:

- 模拟接收json:

```java
		ObjectMapper mapper = new ObjectMapper();
        Employee employee = new Employee();
        employee.setLastName("战五渣").setAge(23).setGender("2").setEmail("333");
        String value = mapper.writeValueAsString(employee);
        mockMvc.perform( MockMvcRequestBuilders
                .post("/employee/save")
                //如果接收的Json数据要指定contentType
                .content(value.getBytes()).contentType(MediaType.APPLICATION_JSON));
  
```

可以在测试方法上加上@Transactional 默认rollback = true , 即使方法执行成功也会回滚,

这样数据库中的数据就不会被污染了

- 模拟文件上传:

```java
  mockMvc.perform(MockMvcRequestBuilders.
                multipart("/employee/upload")
                .file("file", "文件内容".getBytes())
                .contentType(MediaType.MULTIPART_FORM_DATA))
                .andExpect(MockMvcResultMatchers.status().isOk())
                .andDo(MockMvcResultHandlers.print());
```

- 模拟事务:

```java
	@Test
    @Rollback//可以加上@Rollback以清楚显示事务回滚, 默认true
    @Transactional
    public void test() throws Exception {
        MultiValueMap<String, String> map = new LinkedMultiValueMap<>();
        map.add("formId", "2");
        map.add("toId", "3");
        map.add("age", "10");
        mockMvc.perform(MockMvcRequestBuilders
                .get("/employee/tran")
                .params(map));
    }
```

- 检查返回的json值

```java
 mockMvc.perform(MockMvcRequestBuilders
                .post("/employee/get?name={name}","Black"))
                //指定返回的json值, 这里的$表示json的根节点
     			//即返回的body中的lastName必须是Black
                .andExpect(MockMvcResultMatchers.jsonPath("$.lastName")
                        .value("Black"));
```

- 检查返回的视图层

```java
mockMvc.perform(MockMvcRequestBuilders
                .get("/employee/test"))
                .andExpect(MockMvcResultMatchers
                        //检查返回的视图层, 这里的viewName不能与请求的uri相同否则会抛出异常
                        .view().name("view"));

```

 重定向 请求转发

```java
     mockMvc.perform(MockMvcRequestBuilders
                .get("/employee/test"))
                .andExpect(MockMvcResultMatchers
//                        请求转发语重定向
//                       .forwardedUrl("/index.html")
                        .redirectedUrl("/view.html"));
```

- 检查model

```java
        mockMvc.perform(MockMvcRequestBuilders
                .get("/employee/model"))
                .andExpect(MockMvcResultMatchers.model().size(1))
                .andExpect(MockMvcResultMatchers.model().attributeExists("key"))
                .andExpect(MockMvcResultMatchers.model().attribute("key","value"));

```

