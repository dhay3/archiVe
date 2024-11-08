# mybatis-plus 使用

[TOC]

官网: https://mp.baomidou.com/guide/

## #引入依赖

```xml
        <dependency>
            <groupId>com.alibaba</groupId>
            <artifactId>druid-spring-boot-starter</artifactId>
            <version>1.1.22</version>
        </dependency>
        <dependency>
            <groupId>com.baomidou</groupId>
            <artifactId>mybatis-plus-boot-starter</artifactId>
            <version>3.3.1</version>
        </dependency>
```

## #控制台打印sql

方法一

```yaml
mybatis-plus:
  configuration:
    log-impl: org.apache.ibatis.logging.stdout.StdOutImpl #开启log日志
```

方法二

```yaml
logging:
  level:
    com.chz.mapper: debug
```

## #使用

### 查询 select

- 通过查询, 查询所有

```java
    @Test 
    public void tests1() {
        List<Employee> employees = employeeDao.selectList(null);
        employees.forEach(System.out::println);

    }
```

- 按id查询

```java
    @Test 
    public void tests2() {
        Employee employee = employeeDao.selectById(13);
        System.out.println(employee);
    }
```

- 查询一个

```java
    @Test 
    public void tests3() {
        QueryWrapper<Employee> queryWrapper = new QueryWrapper<>();
        queryWrapper.eq("name","zs")
        Employee employee = employeeDao.selectOne(queryWrapper);
        System.out.println(employee);
    }
```

- 使用map作为条件查询

```java
    @Test 
    public void tests4() {
        HashMap<String, Object> map = new HashMap<>();
        //key对应字段名, value对应字段值
        map.put("last_name", "z3f");
        List<Employee> employees = employeeDao.selectByMap(map);
        System.out.println(employees);
    }
```

- 一条记录对应一个map

```java
    @Test 
    public void tests5() {
        QueryWrapper<Employee> queryWrapper = new QueryWrapper<>();
        queryWrapper.eq("gender", 1);
        List<Map<String, Object>> maps = employeeDao.selectMaps(queryWrapper);
        maps.forEach(System.out::println);
    }

```

- 查询多个id

```java
    @Test
    public void test5() {
        List<Integer> list = Arrays.asList(1, 2, 3);
        List<Employee> employees = employeeDao.selectBatchIds(list);
        employees.forEach(System.out::println);
    }
```

- 分页查询

  必须配置分页插件

```java
	@Test 
    public void test6() {
        Page<Employee> page = new Page<>(0, 2);
        //queryWrapper可以是null
        IPage<Employee> employeeIPage = employeeDao.selectPage(page, null);
        //要通过getRecords()拿到结果
        System.out.println(employeeIPage.getRecords());
    }
```

​	分页插件

```java
    @Bean
    public PaginationInterceptor paginationInterceptor() {
        PaginationInterceptor paginationInterceptor = new PaginationInterceptor();
        return paginationInterceptor;
    }
```

- 条加查询

```JAVA
    @Test
    public void test7() {
        QueryWrapper<Employee> wrapper = new QueryWrapper<Employee>().select("age", "gender").eq("last_name", "z3f");
        System.out.println(employeeDao.selectOne(wrapper));
    }
```

- 排序

```java
 	@Test
    public void test8() {
        List<Employee> age = employeeDao.selectList(new QueryWrapper<Employee>().
                orderBy(true, true, "age"));
        System.out.println(age);
    }
```

- 查询列数

```java
    @Test
    public void test9() {
        //查询带条件的count
        Integer integer = employeeDao.selectCount(new QueryWrapper<Employee>().eq("gender", 1));
        //查询所有的count
        Integer integer1 = employeeDao.selectCount(null);
    }
```

### 插入 insert

```java
    @Test
    public void test10() {
        Employee employee = new Employee();
        employee.setLastName("Oka").setEmail("oka@").setGender(2).setAge(3).setSalary(1000D);
        employee.setLastName("z3").setGender(1);
        //insert显示sql当前赋过值的字段
        int num = employeeDao.insert(employee);
        //mp直接会回显主键
        //mybatis需要配置useGeneratedKeys,keyProperty才能回显主键
        Integer id = employee.getId();
        System.out.println("生效行数:" + num + "\t主键值:" + id);
    }
```

### 更新 upadate

```java
    @Test
    public void test11() {
        Employee employee = new Employee();
        //如果不想修改某个字段,不赋值即可
        employee.setId(13).setLastName("z3f").setAge(33);
        int i = employeeDao.updateById(employee);
        System.out.println("生效行数:" + i);
    }
```

支持`lambda`表达式

```java
@Test
public void tests11() {
    UpdateWrapper<Employee> updateWrapper = new UpdateWrapper<>();
    //where条件
    updateWrapper.lambda()
            .eq(Employee::getLastName, "Black")
            .eq(Employee::getGender, 1);
    Employee employee = new Employee();
    //想要修改的值
    employee.setEmail("black123@qq");
    employeeDao.update(employee, updateWrapper);
}
```

### 删除 delete

```java
    @Test 
    public void test12() {
        int i = employeeDao.deleteById(6);
        System.out.println("生效行数:" + i);
    }
```

```java
    @Test //deleteByMap
    public void test13() {
        HashMap<String, Object> map = new HashMap<>();
        //封装条件
        map.put("age", 22);
        employeeDao.deleteByMap(map);
    }
```

## #条件



- gt

  大于

- le

  小于等于

- lt

  小于

- isNull

  空值

- isNotNull

  不为空

- eq

  相等

| 条件        | 描述                                 |
| ----------- | ------------------------------------ |
| ge          | great equal 大于等于                 |
| gt          | great than 大于                      |
| le          | less equal小于等于                   |
| lt          | less than小于                        |
| isNull      | 空值                                 |
| isNotNull   | 不为空值                             |
| eq          | equal 等于                           |
| ne          | not equal 不等于                     |
| between     | 包含                                 |
| notBetween  | 不在区间内                           |
| allEq       | 所有相同                             |
| like        | 模糊查询                             |
| notLike     |                                      |
| in          | 包含                                 |
| notIn       | 不包含                               |
| exists      | 存在                                 |
| notExists   | 不存在                               |
| or          | 不调用or默认使用and                  |
| and         | 默认使用and拼接                      |
| orderBy     | 按指定方式排序                       |
| orderByDesc | 降序                                 |
| orderByAsc  | 升序                                 |
| last        | 直接拼接到sql的最后, 有sql注入的风险 |

