# @Validated 和 @Valid

@Valid有jsr303提供,而@Validated有spring提供
 效果一样但是,后者比前者功能更强大, 提供分组(类似于@JsonView)

正常在 @Validated 注解不加分组参数的情况下，@Valid 注解和 @Validated 注解会随机校验实体类中不加 groups 属性的字段，不分先后顺序

### #pojo

```java
@Data
public class User implements Serializable {

    private static final long serialVersionUID = -8677072538094918597L;

    public interface IUserName {
    }
    //与@JsonView效果差不多
    public interface IAllFiled extends IUserName{
    }

    @NotBlank(groups = IUserName.class, message = "username can't be null")
    private String username;
    @NotNull(groups = IAllFiled.class, message = "age cant'be null")
    private Integer age;
    @NotNull(message = "性别只能为0或1")
    @Min(value = 0, message = "性别只能为0或1")
    @Max(value = 1, message = "性别只能为0或1")
    private Integer gender;
}
```

### #Controller

没有指定groups

```java

    @ResponseBody
    @RequestMapping("/test3")
    public String test3(@Validated User user, BindingResult result) {
       if (result.hasErrors()){
           result.getFieldErrors().stream()
                   .map(DefaultMessageSourceResolvable::getDefaultMessage)
                   .forEach(System.out::println);
       }
        System.out.println(user);
        return "test3";
    }

    
    @ResponseBody
    @RequestMapping("/test4")
    public String test4(@Valid User user,BindingResult result) {
       if (result.hasErrors()){
           result.getFieldErrors().stream()
                   .map(DefaultMessageSourceResolvable::getDefaultMessage)
                   .forEach(System.out::println);
       }
        System.out.println(user);
        return "test4";
    }
```

发送请求 `localhost:8080/test3?username&age&gender=12`

结果:

<img src="..\..\..\..\java资料\我的笔记\springboot\img\27.png"/>

**结论**: 如果属性加了groups使用@Validated或是@Valid不指定groups,那么带有groups的属性验证不会生效

---

指定gruop , @Valide不支持指定group

这里的`IAllField`继承了`IUserName`所以会校验age也会校验username

```java
    @ResponseBody
    @RequestMapping("/test3")
    public String test3(@Validated(User.IAllFiled.class) User user, BindingResult result) {
        result.getFieldErrors().stream()
                .map(DefaultMessageSourceResolvable::getDefaultMessage)
                .forEach(System.out::println);
        System.out.println(user);
        return "test3";
    }
```

结果:

<img src="..\..\..\..\java资料\我的笔记\springboot\img\28.PNG"/>

---

所以推荐使用@validated
