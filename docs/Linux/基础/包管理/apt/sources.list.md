# sources.list

## 更新源

apt将更新源配置在`/etc/apt/sources.list`和`/etc/apt/sources.list.d/`中。配置完成后需要`apt-get update`才能更新信息。默认按照顺序解析。

## 格式

更新源有两种格式，one-line format和deb822-style format。两种格式的配置文件可以同时存在。

### one-line format

文件需要以`.list`结尾，通用格式所有的apt版本都支持

```
deb [arch=amd64,x86] https://download.docker.com/linux/debian kali-rolling stable
```

- 选项名使用equal sign(=)分隔。例如`arch=...`
- 如果需要提供选项，使用square brackets([])。例如`[arch=...]`
- 如果选项有多个值，需要使用comma(,)。例如`amd64,x86`
- 同时支持`+=`和`-=`，加默认值得同时修改

==one-line format 被用在 deb 和 deb-src types==

```
deb-src http://http.kali.org/kali kali-rolling main contrib non-free
deb [arch=amd64] https://download.docker.com/linux/debian kali-rolling stable
deb-src [arch=amd64] https://download.docker.com/linux/debian kali-rolling stable
deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse

deb [ option1=value1 option2=value2 ] uri suite [component1] [component2] [...]
deb-src [ option1=value1 option2=value2 ] uri suite [component1] [component2] [...]
```

- deb表示binary packages的来源，deb-src表示source packages的来源
- suite通常对应发行版本
- component 表示安装包来自的repo通常为 main，contrib，non-free ...

### deb822-style format

文件需要以`.sources`结尾，并不是通用的(atp 1.1后支持)。注意中间有一个空格。如果一个属性有多个值，通过空格隔开。

```
deb [ option1=value1 option2=value2 ] uri suite [component1] [component2] [...]
deb-src [ option1=value1 option2=value2 ] uri suite [component1] [component2] [...]
#等价
Types: deb deb-src
URIs: uri
Suites: suite
Components: [component1] [component2] [...]
option1: value1
option2: value2
```

## option

用于过滤源中的包

- arch：只有指定架构编译的包的才能被本机接受，默认所有架构
- trusted：yes 表示所有来源的包都被信任即使没有被认证，no表示所有来源的包都被设置为不信任即使通过认证，留空表示由apt自己选择

## kali apt source

```
Types: deb deb-src
URIs: https://mirrors.aliyun.com/kali https://mirrors.tuna.tsinghua.edu.cn/kali/dists/kali-rolling/
Suites: kali-rolling
Components:  main non-free contrib
```









