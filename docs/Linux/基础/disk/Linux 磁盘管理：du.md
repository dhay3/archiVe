# Linux 磁盘管理：du

`du`命令用于查看文件的大小，==`du`默认会遍历当前所有目录==

- -a

  du默认只输目录的大小，可以使用该参数输出文件

  ```bash
  cpl in ~/.ssh λ du 
  20	.
  cpl in ~/.ssh λ du -a
  4	./known_hosts
  4	./id_rsa.pub
  4	./id_rsa
  4	./config
  20	.
  ```

- -h | --human-readable

  以K，M，G为单位显示

  ```
  cpl in ~/note on master ● ● λ du -hd 1
  352M	./.git
  49M	./imgs
  407M	./docs
  807M	.
  ```

- -B

  以指定单位显示，==但是文件大小不足时以最小单位显示==

  ```bash
  cpl in ~/note on master ● ● λ du -BM -d 1
  352M	./.git
  49M	./imgs
  407M	./docs
  807M	.
  cpl in ~/note on master ● ● λ du -BG -d 1
  1G	./.git
  1G	./imgs
  1G	./docs
  1G	.
  ```

- -d | --max--depth=N

  输出指定目录深度的内容，当前目录从0开始

  ```bash
  cpl in ~/note on master ● ● λ du -d 0
  825360	.
  cpl in ~/note on master ● ● λ du -d 1
  359776	./.git
  49300	./imgs
  416264	./docs
  825360	.
  ```

- --exclude=pattern

  使用posix regex，过滤指定的文件

  ```
  du --exclude='*.o'
  ```

  















