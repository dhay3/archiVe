# Spring定时任务/Cron

参考:

https://www.jb51.net/article/138900.htm

https://www.cnblogs.com/summertime-wu/p/7641589.html

https://www.cnblogs.com/mmzs/p/10161936.html

### Cron 表达式

- 6位长度    秒  分  时  日  月  星期
- 7位长度    秒  分  时  日  月  星期  年

一般采用6位长度

### 字段允许出现的值

| 字段                   | 允许值                                 | 允许的特殊字符             |
| ---------------------- | -------------------------------------- | -------------------------- |
| 秒（Seconds）          | 0~59的整数                             | , - * /   四个字符         |
| 分（Minutes）          | 0~59的整数                             | , - * /   四个字符         |
| 小时（Hours）          | 0~23的整数                             | , - * /   四个字符         |
| 日期（DayofMonth）     | 1~31的整数（但是你需要考虑你月的天数） | ,- * ? / L W C   八个字符  |
| 月份（Month）          | 1~12的整数或者 JAN-DEC                 | , - * /   四个字符         |
| 星期（DayofWeek）      | 1~7的整数或者 SUN-SAT （1=SUN）        | , - * ? / L C #   八个字符 |
| 年(可选，留空)（Year） | 1970~2099                              | , - * /   四个字符         |

- *

  表示匹配该域的任意值。假如在Minutes域使用*, 即表示每分钟都会触发事件。

- ?

  只能用在DayofMonth和DayofWeek两个域。它也匹配域的任意值，但实际不会。因为DayofMonth和DayofWeek会相互影响。例如想在每月的20日触发调度，不管20日到底是星期几，则只能使用如下写法： 13 13 15 20 * ?, 其中最后一位只能用？，而不能使用*，如果使用*表示不管星期几都会触发，实际上并不是这样。

- -

  表示范围。例如在Minutes域使用5-20，表示从5分到20分钟每分钟触发一次

- ,

  表示列出枚举值。例如：在Minutes域使用5,20，则意味着在5和20分每分钟触发一次。

- L

  表示最后，只能出现在DayofWeek和DayofMonth域。如果在DayofWeek域使用5L,意味着在最后的一个星期四触发。

- W

  表示有效工作日(周一到周五),只能出现在DayofMonth域，系统将在离指定日期的最近的有效工作日触发事件。例如：在 DayofMonth使用5W，如果5日是星期六，则将在最近的工作日：星期五，即4日触发。如果5日是星期天，则在6日(周一)触发；如果5日在星期一到星期五中的一天，则就在5日触发。另外一点，W的最近寻找不会跨过月份 。

- LW

  这两个字符可以连用，表示在某个月最后一个工作日，即最后一个星期五。

- #

  用于确定每个月第几个星期几，只能出现在DayofMonth域。例如在4#2，表示某月的第二个星期三。

### 常用例子

- `0 0 2 1 * ? *`

  表示在每月的1日的凌晨2点调整任务

- `0 15 10 * * ? `

  每天上午10:15触发

- `0 * 14 * * ?`

  在每天下午2点到下午2:59期间的每1分钟触发

- `0 0/5 14 * * ? `

  在每天下午2点到下午2:55期间的每5分钟触发

- `0 0/5 14,18 * * ? `

  在每天下午2点到2:55期间和下午6点到6:55期间的每5分钟触发

- `0 0-5 14 * * ? `

  在每天下午2点到下午2:05期间的每1分钟触发

- `0 15 10 ? * MON-FRI`

  周一至周五的上午10:15触发

- `0 15 10 L * ? `

   每月最后一日的上午10:15触发

### 整合SpringBoot

方法一

```java
@Configuration      
@EnableScheduling   
public class SaticScheduleTask {
  
    @Scheduled(cron = "0/5 * * * * ?")
    //@Scheduled(fixedRate=5000)
    private void configureTasks() {
        System.err.println("执行静态定时任务时间: " + LocalDateTime.now());
    }
}
```

方法二

```java
@Component
public class ScheduleTask implements SchedulingConfigurer {

    @Override
    public void configureTasks(ScheduledTaskRegistrar taskRegistrar) {
        taskRegistrar.addCronTask(new CronTask(()-> System.out.println("hello world"),
                "* * * 2 8 ?"));
    }
}
```

