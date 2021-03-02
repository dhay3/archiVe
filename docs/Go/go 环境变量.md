# go 环境变量

go中所有的环境变量可以使用`go env`来获取

- GOOROOT

  该变量的值是Go的当前安装目录

  ```
  root in /usr/local/go_code/src λ go env | grep GOROOT
  GOROOT="/usr/local/src/go"
  ```

- GOPATH

  该变量的值为Go的工作集合(意味这可以有很多个，和`${PATH}`相似用`:`隔开)。

  一般会有三个目录，src，pkg，bin

  ```
  /home/halfrost/gorepo
  ├── bin
  ├── pkg
  └── src
  ```

  1. pkg用来存储通过`go install` 安装后的代码包归档文件，以`.a`结尾

  2. bin用来保存编译后可执行的二进制文件

  3. src用来保存Go源码文件

     源码文件有分为三种

     - 命令源文件

       Go程序的入口，main函数所在的文件

     - 库源码文件

       无法被主动执行的文件，被用作库被调用

     - 测试源码文件

       `_test.go`为后缀的代码文件，函数名为TestXxx

- GOOS和GOARCH

  操作系统和架构，包中默认值。

- GOBIN

  `go install`安装后存储可执行二进制文件的路径