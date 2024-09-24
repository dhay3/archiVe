# LIMIT

MySQL下标从0开始, 左闭右包

- 查询学生表前5条记录

  `SELECT * FROM tbl_stu LIMIT 5`

- 按照5条记录为一页, 查询第二页

  `SELECT * FROM tbl_stu LIMIT 5, 5`

  (currentPage - 1) * pageSize, pageSize

  这里从第6条记录开始查询5条, 即6到10

  select DISTINCT id, `name`, country, province, city from person;

