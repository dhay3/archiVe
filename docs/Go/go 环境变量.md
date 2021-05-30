# Go 环境变量

https://golang.org/cmd/go/#hdr-Environment_variables

- GO111MODULE

  是否使用[module-aware mode](https://golang.org/ref/mod#mod-commands)来管理go的包。如果使用module-aware mode，go command通过go.mod文件去使用依赖。如果使用GOPATH mode，go command通过`%GOPATH%`去使用依赖。可选off，on，auto

- GOARCH

  编译代码的架构可选amd64，386，arm，ppc64

- GOBIN

  `go install`安装的目录

- GOCACHE

  go command存储缓存信息的目录

- GOENV

  Go 环境变量配置文件存储的位置

- GOOS

  编译代码的平台，可以事linux，darwin，windows

- GOPATH

  GOPATH用来解析import。如果没有设置GOPATH环境变量，默认使用`~/go`。GOPATH下要包含如下三个文件

  1. src：src存储源码，例如GOPATH为DIR，在DIR/src/foo/bar有源文件，那么import的将是foo/bar。==所以在src下的源文件是不能被import的==
  2. pkg：`go install`安装的存档文件
  3. bin：编译后可执行的二进制文件。例如GOPATH为DIR，在DIR/src/foo/bar有源文件，那么编译后是DIR/bin/bar，而不是DIR/bin/foo/bar。默认编译在GOBIN，如果设置了GOBIN需要使用绝对路径

  ```
   GOPATH=/home/user/go
  
      /home/user/go/
          src/
              foo/
                  bar/               (go code in package bar)
                      x.go
                  quux/              (go code in package main)
                      y.go
          bin/
              quux                   (installed command)
          pkg/
              linux_amd64/
                  foo/
                      bar.a          (installed package object)
  ```

  Go会从所有的GOPATH中找源码，==但是只会将下载的内容放在GOPATH中第一个目录中==

  **moudel-aware mode VS GOPATH**

  如果使用modules，GOPATH不再被用来解析imports。==但是还是被用来下载源码和编译。==

  **Internal**

  如果一个包名为internal，这个包中的内容只能在internal父级的包中使用(不包括父级的父级和同级)

  ```
      /home/user/go/
          src/
              crash/
                  bang/              (go code in package bang)
                      b.go
              foo/                   (go code in package foo)
                  f.go
                  bar/               (go code in package bar)
                      x.go
                  internal/
                      baz/           (go code in package baz)
                          z.go
                  quux/              (go code in package main)
                      y.go
  ```

  例如这里的`bar/z`只能在foo的这级包中使用

- GOPROXY

  配置代理参考：https://goproxy.io/zh/docs/getting-started.html

  module proxy(==只有在module mode下才会被使用==)，如果有多个proxy urls通过commas或pipeline隔开。

  ```
  GOPROXY=https://proxy.golang.org,direct
  
  off: disallows downloading modules from any source.
  direct: download directly from version control repositories instead of using a module proxy.
  ```

- GOROOT

  go的根目录(安装目录)













