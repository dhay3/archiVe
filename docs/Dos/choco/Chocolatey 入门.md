# Chocolatey 入门

https://docs.chocolatey.org/en-us/

choco是windows上的一个包管理器，类似于snap。和apt或是yum一样，choco也有源。

```
#添加一个源，名为bob
PS C:\WINDOWS\system32> choco source add -n=bob -s="https://somewhere/out/there/api/v2/"
Chocolatey v0.10.15
Added bob - https://somewhere/out/there/api/v2/ (Priority 0)

#显示当前使用的源
PS C:\WINDOWS\system32> choco source list
Chocolatey v0.10.15
chocolatey - https://chocolatey.org/api/v2/ | Priority 0|Bypass Proxy - False|Self-Service - False|Admin Only - False.
bob - https://somewhere/out/there/api/v2/ | Priority 0|Bypass Proxy - False|Self-Service - False|Admin Only - False.

choco source disable -n=bob
choco source enable -n=bob
choco source remove -n=bob
```

未来方便使用choco，可以设置代理

```
choco config list
choco config set proxy 	<key> <value>
choco config get proxy 
choco config unset proxy
```

搜索指定的包或查看详细信息

```
PS C:\WINDOWS\system32> choco.exe list git | more
PS C:\WINDOWS\system32> choco info git
```

安装，更新，卸载

```
choco install git
choco upgrade git
choco uninstall git
```

