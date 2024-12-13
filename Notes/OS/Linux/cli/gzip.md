# gzip

## Digest

GUN/gzip用于压缩和解压`.gz`文件。==会替换原有的文件==

> gunzip用于==解压==`.tgz`，`taz`是`.tar.gz`缩写，同样也能解压`gz`
>
> zcat等价于`gunzip -c`，按照编码来识别压缩文件，无需`.gz`也可，输出stdout

## Optional args

- `-d`

  decompose

  解压

  ```
  [root@chz Desktop]# gzip -d minikube.gz 
  [root@chz Desktop]# ls
  minikube  test
  ```

- -f

  强制压缩或解压(如果生成的文件已尽存在)

- -v

  显示详细信息

  ```
  [root@chz Desktop]# gzip -v minikube 
  minikube:	 56.7% -- replaced with minikube.gz
  ```

- -r

  将文件夹中的文件压缩，==文件夹不会改变==

  ```
  [root@chz Desktop]# gzip -r test
  [root@chz Desktop]# ls
  minikube.gz  test
  [root@chz Desktop]# cd test
  [root@chz test]# ls
  tset1.gz  tset2.gz
  ```

  test目录下的所有文件都以`.gz`为扩展名

- -c

  将文件压缩后的内容输出，不会改变文件内容

  ```
  [root@chz test]# gzip -c test1 test2
  ��_test1�H���W(�/�I�-�
                           ��_test2+).1��ç�
  ```


## Examples

- 压缩文件

  不带任何参数不能压缩目录，只能压缩文件

  ```
  cpl in /tmp λ gzip test
  gzip: test is a directory -- ignored
  
  [root@chz Desktop]# gzip minikube 
  [root@chz Desktop]# ls
  minikube.gz  test
  ```

- 文件内容合并压缩

  ```
  [root@chz test]# gzip -c test1 test2 > test.gz
  [root@chz test]# ls
  test1  test2  test.gz
  #使用cat方式压缩效率更高
  [root@chz test]# cat test1 test2| gzip > test.gz
  [root@chz test]# ls
  test1  test2  test.gz
  ```

  