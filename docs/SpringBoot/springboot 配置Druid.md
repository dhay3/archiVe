# springboot配置Druid

```
spring:
  datasource:
    driver-class-name: com.mysql.cj.jdbc.Driver
    url: jdbc:mysql://localhost:3306/mp?useUnicode=true&characterEncoding=utf-8&serverTimezone=Asia/Shanghai
    username: root
    password: 12345
    type: com.alibaba.druid.pool.DruidDataSource
    druid:
      initial-size: 5 #初始连接数
      max-active: 10 #最大活动连接
      max-wait: 60000 #从池中取连接(没有闲置连接)的最大等待时间,-1表示无限等待
      min-idle: 5 #最小闲置数,小于min-idle连接池会主动创建新的连接
      time-between-eviction-runs-millis: 60000 #清理线程启动的间隔时间,当线程池中没有可用的连接启动清理线程
      min-evictable-idle-time-millis: 300000 #清理线程保持闲置最小时间
      validation-query: SELECT 1  #用于校验连接
      test-on-borrow: false #请求连接时是否校验,比较消耗性能,一般设置false
      test-on-return: false #归还连接时是否校验,比较消耗性能,一般设置false
      test-while-idle: true #清理线程通过validation-query来校验连接是否正常,如果不正常将从连接池中移除
      pool-prepared-statements: true #存储相同逻辑的sql到连接池的缓存中
      filters: stat,wall #监控统计web的statement(sql),以及防sql注入的wall
      web-stat-filter:
        enabled: true #开启监控的uri,默认false
        url-pattern: /* #添加过滤规则
        exclusions: "*.js,*.gif,*.jpg,*.png,*.css,*.ico,/druid" #忽略过滤
      stat-view-servlet:
        enabled: true #开启视图管理界面,默认false
        url-pattern: /druid/* #视图管理界面uri
        login-username: druid #账号
        login-password: 12345 #密码
```

