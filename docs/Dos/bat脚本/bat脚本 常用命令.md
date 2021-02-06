# bat脚本

## 常用命令

> dos中支持使用通配符`*`

| 操作      | 描述                                                         |
| --------- | ------------------------------------------------------------ |
| cls       | 清屏                                                         |
| md        | 建立子目录                                                   |
| rd        | 删除目录                                                     |
| echo      | 标准输出                                                     |
| dir       | 列出文件名                                                   |
| type      | 查看文件内容                                                 |
| del       | 删除文件，支持`*`通配符                                      |
| attrib    | 修改文件属性，`attrib+H ./*`隐藏当前目录下的所有文件         |
| copy      | 复制文件，`copy con test.txt`将控制台中的内容保存            |
| move      | 移动文件，同样可用于重命名                                   |
| shutdown  | 关机                                                         |
| tasklist  | 列出当前进程                                                 |
| taskkill  | 杀死进程                                                     |
| ren       | 重命名                                                       |
| cd        | 切换目录，盘符直接切换到该目录                               |
| assoc     | 修改后缀关联文件，`assoc .txt=exefile`将后缀为txt文件变成可执行文件 |
| net user  | 查看用户                                                     |
| net share | 查看共享文件                                                 |
| fsutil    | `fsutil file createnew a.txt 102400`创建一个100KB的文件      |
|           |                                                              |
|           |                                                              |