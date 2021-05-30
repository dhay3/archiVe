# MySQL UNION

现有一张表person1

<img src="..\..\..\imgs\_MySQL\Snipaste_2020-09-14_15-42-18.png"/>

- 使用UNION要求字段数相同

<img src="..\..\..\imgs\_MySQL\Snipaste_2020-09-14_15-44-21.png"/>

- 使用相同字段数，不同字段

  > ==UNION会自动去重，会将结果拼接到第一张表查询的结果之后，使第一张表的字段做为字段名==
  >
  > 如果想要保留重复的字段使用UNION ALL

<img src="..\..\..\imgs\_MySQL\Snipaste_2020-09-14_15-47-05.png"/>

- UNION使用ORDER BY

> 内层的排序不会生效，最外层的才会生效

<img src="..\..\..\imgs\_MySQL\Snipaste_2020-09-14_16-05-19.png"/>

<img src="..\..\..\imgs\_MySQL\Snipaste_2020-09-14_16-06-23.png"/>
