# packer 函数

> 这里的``表示字符常量

packer使用go的模板语法，有内置的一些函数

1. env 获取环境变量

   ```
   {{env `HOME`}}
   ```

2. user 调用variables块中指定的变量

   ```
   "access_key":"{{user `access_key`}}"
   ```

3. build  从provisioners和post-processors中获取值

   [具体可取的值][https://www.packer.io/docs/templates/legacy_json_templates/engine#build]

   ```
   "environment_vars": ["TESTVAR={{ build `PackerRunUUID`}}"]
   ```

   

