# Github操作

[TOC]

### #菜单栏

- watch:

  是否持续受到项目的动态

- fork:

  复制某个项目到自己的Github仓库中

- star:

  可以理解为点赞

- clone:

  将项目下载至本地

- follow:

  关注你感兴趣的作者, 会收到他们的动态

### #关键字

!**注意**冒号后面没有空格

- xxx in:name

  查询名字叫xxx的仓库(**默认按照star降序**)

  可以组合使用

  As Fllow

  ```
  seckill in:name,readme
  ```

  在仓库名和readme中有seckill关键字

- xxx stars:> 5000

  查询指定star范围内, 带有关键字xxx的仓库

- xxx forks:>5000

  查询指定fork范围内, 带有关键字xxx的仓库

 **组合使用fork和star**

As Fllow

```
springboot forks:100..200 stars:80..100
```

fork数在100到200, stars数80到100之间, 名字中包含springboot的仓库

### #特殊用法

- #L 

  L代表line

  高亮显示指定源码中指定行数代码

  As Fllow

  ```
  https://github.com/JeffLi1993/springboot-learning-example/blob/master/chapter-3-spring-boot-web/src/main/java/demo/springboot/web/BookController.java#L45
  ```

  高亮显示第45行

  As Follow

  ```
  https://github.com/JeffLi1993/springboot-learning-example/blob/master/chapter-3-spring-boot-web/src/main/java/demo/springboot/web/BookController.java#L45-#L60
  ```

  高亮显示第45行到60行
