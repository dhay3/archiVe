# go 指令

参考：

https://golang.org/doc/cmd

### 通用参数

- -n 

  打印将要运行的命令，但是不执行

- -x

  打印将要运行的命令并执行

### go build

编译包和依赖，只有在目录中有main包才会编译

syntax：`go build [-o output] [build flags] [packages]`

如果没有指定packages默认以当前目录编译，在当前目录生成以目录为名的二进制文件

### go clean

清除编译的二进制文件或压缩文件

syntax：`go clean [clean flags] [build flags] [packages]`

### go run

编译并运行文件(==但是不会保留编译的文件==)

syntax：`go run [build flags] package`

==注意，go run 必须指明package或main入口函数==

```
D:\workspace_for_go\learn\src\2.编译>go run 2.编译
hello world

D:\workspace_for_go\learn\src\2.编译>go run TestCompile.go
hello world

D:\workspace_for_go\learn\src\2.编译>type TestCompile.go
package main
import "fmt"
func main() {
        fmt.Print("hello world")
}
```

### go generate

go generate 扫描文件，找到`//go:generate command argument`并执行command

```
//go:generate echo 你好
//go:generate go run
//go:generate echo $GOPATH
func main() {
	fmt.Print("hello world")
}

输出
你好
hello world
```

### go install

编译并安装包，文件会被安装在`$GOPATH/bin`中

syntax：`go install [build flags] [packages]`

当没有带packages，go install 默认对当前包编译并安装，使用如下命令来查看具体调用的命令。通过`echo %GOPATH%`查看当前项目的GOPATH

```
go install -n
```

### go get

添加依赖到当前模块并安装

syntax：`go get [-d] [-t] [-u] [-v] [-insecure] [build flags] [packages]`

To add a dependency for a package or upgrade it to its latest version:

```
go get example.com/pkg
```

To upgrade or downgrade a package to a specific version:

```
go get example.com/pkg@v1.2.3
```

To remove a dependency on a module and downgrade modules that require it:

```
go get example.com/mod@none
```

### go vet

用来检验代码里的错误

```
D:\workspace_for_go\learn\src\2.编译>type TestCompile.go
package main
import "fmt"
func main() {
        str := "hello world!"
        fmt.Printf("%d\n", str)
}

D:\workspace_for_go\learn\src\2.编译>go vet
# 2.编译
.\TestCompile.go:7:2: Printf format %d has arg str of wrong type string
```

### go test

运行测试代码并校验

```
D:\workspace_for_go\learn\src\test>type main.go
package main

func Method(text string) string {
        return text
}

D:\workspace_for_go\learn\src\test>type main_test.go
package main
import "testing"

func TestMethod(t *testing.T)  {
        if "hello" != Method("hello") {
                panic("error")
        }
}

D:\workspace_for_go\learn\src\test>go test
PASS
ok      test    0.502s
```

### go doc

获取指定包，方法或属性的描述

```
go doc fmt
go doc fmt.print
go doc fmt.
```

### go env

获取当前GO的环境变量，或通过`-w`设置变量值

```
go env -w GO111MODULE=on
```

### go fix

更新指定包的api

syntax：`go fix [packages]`