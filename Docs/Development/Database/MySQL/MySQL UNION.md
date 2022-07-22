# MySQL UNION

现有一张表person1

![Snipaste_2020-09-14_15-42-18](https://git.poker/dhay3/image-repo/blob/master/20220722/Snipaste_2020-09-14_15-42-18.6sii1sa8gpog.webp?raw=true)

- 使用UNION要求字段数相同

![Snipaste_2020-09-14_15-44-21](https://git.poker/dhay3/image-repo/blob/master/20220722/Snipaste_2020-09-14_15-44-21.32pjksusdku8.webp?raw=true)

- 使用相同字段数，不同字段

  > ==UNION会自动去重，会将结果拼接到第一张表查询的结果之后，使第一张表的字段做为字段名==
  >
  > 如果想要保留重复的字段使用UNION ALL

![Snipaste_2020-09-14_15-47-05](https://git.poker/dhay3/image-repo/blob/master/20220722/Snipaste_2020-09-14_15-47-05.3ouwlba2t88w.webp?raw=true)

- UNION使用ORDER BY

> 内层的排序不会生效，最外层的才会生效

![Snipaste_2020-09-14_16-05-19](https://git.poker/dhay3/image-repo/blob/master/20220722/Snipaste_2020-09-14_16-05-19.2aqjvzg9m5a8.webp?raw=true)



![Snipaste_2020-09-14_16-06-23](https://git.poker/dhay3/image-repo/blob/master/20220722/Snipaste_2020-09-14_16-06-23.6zyptnj0mpkw.webp?raw=true)
