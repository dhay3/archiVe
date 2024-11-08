# springboot Jackson

@ResponseBody

方法返回的是字符串, 返回到前端的就是字符串

方法返回的是对象, 返回到前端就是JSON

> 序列化

- **什么是序列化**: 是将对象转为字节数组的过程, 反序列化是将字节数组转为对象. 对象序列化后可以在网络上传输, 或者保存在硬盘

- **Json在前端的序列化:**  `JSON.stringify(json对象)` ,将json对象转为json串

- **Json在前端的反序列化:** `JSON.parseJSON(str)`, 将json串转为json对象

- **Json后端的序列化:** 将pojo转为JSON
- **Json后端的反序列化:** 将JSON转为pojo

> 为什么使用Content-Type = application/json

==告诉服务端发送的数据类型是Json串==

`ajax` 请求时不设置任何`contentType`，默认将使用`contentType: "application/json”application/x-www-form-urlencoded`，这种格式的特点就是，name/value 成为一组

每组之间用 & 联接，而 name与value 则是使用 = 连接。如`www.baidu.com/query?user=username&pass=password` 这是get请求, 而 post 请求则是使用请求体，参数不在 url 中，在请求体中的参数表现形式也是: `user=username&pass=password`的形式。

使用这种contentType时，对于简单的json对象类型，如：`{“a”:1,"b":2,"c":3}` 这种，将也会被转成`user=username&pass=password` 这种形式发送到服务端。而服务端接收时就按照正常从from表单中接收参数那样接收即可，不需设置@RequestBody之类的注解。但对于复杂的json 结构数据，这种方式处理起来就相对要困难，服务端解析时也难以解析，所以，就有了`application/json `这种类型，这是一种数据格式的申明，明确告诉服务端接收什么格式的数据，服务端只需要根据这种格式的特点来解析数据即可。

> ajax  dataType , Content-Type

- dataType 表示ajax预接收的数据类型

  `"json"`表示接收Json对象

  **@ResponseBody如果返回的是字符串. 那么转为的就是字符串, 如果返回的是对象那么转为Json对象,**

  **所以返回字符串时不能使用dataType: "json" 来接收**

- Content-Type 表示传入到服务端的参数类型

​       `"applicaiton/json"`表示传入到服务端的参数类型是Json串

### #pojo

```java
public class User implements Serializable {
    private static final long serialVersionUID = 262962110318827040L;
    private String name;
    private Integer age;
    private Integer password;
    private Date date;
}
```

- @JsonFormat

  将时间对象序列化为Json串, pattern指定时间格式, timzone指定时间

  一般采用如下格式

  ```java
  @JsonFormat(pattern = "yyyy-MM-dd HH:mm:ss",timezone = "GMT+8")
  private Date date;
  ```

  未经格式化Json

  ```json
  {
      "name": "李文良",
      "age": 12,
      "password": 123,
      "date": "2020-04-26T02:53:11.127+0000"
  }
  ```

  格式化Json

  ```json
  {
      "name": "李文良",
      "age": 12,
      "password": 123,
      "date": "2020-04-26 02:53:11"
  }
  ```

  

- @JsonProperty

  为属性的Json key 指定一个别名

  ```java
  @JsonProperty("bth")
  private String name;
  ```

  输出的json

  ```json
  {
      "age": 12,
      "password": 123,
      "date": "2020-04-26 02",
      "username": "李文良"
  }
  ```

- @JsonIgnore

  忽略属性Json key(==注意不是属性==)序列化与反序列化,

  ```java
  @JsonIgnore
  private Integer password;
  ```

  输出的Json

  ```json
  {
      "name": "李文良",
      "age": 12,
      "date": "2020-04-26 02"
  }
  ```

- @JsonIgnoreProperties

  忽略忽略一组json key

  ```java
  @JsonIgnoreProperties({"name","age"})
  public class User implements Serializable {
  ```

  输出的Json

  ```json
  {
      "password": 123,
      "date": "2020-04-26 02"
  }
  ```

  因为是忽略Json key,  但是如果属性取了别名

  ```java
  @JsonIgnoreProperties({"name","age"})
  public class User implements Serializable {
      @JsonProperty("username")
      private String name;
  ```

  输出的Json就会是

  ```json
  {
      "password": 123,
      "date": "2020-04-26 02",
      "username": "李文良"
  }
  ```

- @JsonNaming

  指定Json key的生成策略

  ```java
  @JsonNaming(PropertyNamingStrategy.UpperCamelCaseStrategy.class)
  public class User implements Serializable {
      private static final long serialVersionUID = 262962110318827040L;
      private String UserName;
      private Integer user_age;
      private Integer password;
      @JsonFormat(pattern = "yyyy-MM-dd hh")
      private Date date;
  }
  ```

  `PropertyNamingStrategy.UpperCamelCaseStrategy`

  将默认的驼峰命名规则json key 改为首字母大写

  `PropertyNamingStrategy.SnakeCaseStrategy`

  将`UserName`拆分成`user_name`作为json key

- @JsonSerialize

  指定一个类来实现属性或类的序列化, 指定的类必须实现Serializalier接口

  将对象序列化为json对象, 输出到前端

  ```java
  @JsonSerialize(using = LocalDateTimeSerializer.class)
  private LocalDateTime localDateTime;
  ```

- @JsonDeserialize

  指定一个类来实现属性或类的反序列化, 指定的类必须实现Serializalier接口

  将前端传过来的json串转为对象

  ```java
  @JsonDeserialize(using = LocalDateTimeDeserializer.class)
  private LocalDateTime localDateTime;
  ```

  ==这里的LocalDateTime系列的类 必须采用该方法, 否则 会抛出异常==

- JsonView

  属性只有和controller层的@JsonView指定的Class对应,才会打印属性

  ```java
  public class User implements Serializable {
      private static final long serialVersionUID = 262962110318827040L;
      public interface UserNameFieldView{
  
      }
      public interface AllUserFieldView extends UserNameFieldView{
  
      }
      @JsonView(UserNameFieldView.class)
      private String UserName;
      @JsonView(AllUserFieldView.class)
      private Integer user_age;
      @JsonView(AllUserFieldView.class)
      private Integer password;
      @JsonFormat(pattern = "yyyy-MM-dd hh")
      @JsonView(AllUserFieldView.class)
      private Date date;
  }
  
  ```

  ```java
  @JsonView(User.UserNameFieldView.class)
      @RequestMapping("/view")
      public User JsonView() {
          User user = new User();
          user.setUserName("李文良").setUser_age(12).setDate(new Date()).setPassword(123);
          return user;
      }
  ```

  输出的json

  ```json
  {
      "name": "李文良"
  }
  ```

  将controller中的@JsonView替换为Usr.AllUserFieldView.class

  ```json
  {
      "name": "李文良",
      "age": 12,
      "password": 123,
      "date": "2020-04-26 03"
  }
  ```

  由于AllUserFieldView继承了UserNameFieldView, 所以也会显示name

