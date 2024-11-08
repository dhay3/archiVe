# Powershell overview

> `Update-Help`用于更新Help内容
>
> 1. powershell大小写不敏感，但是有数据类型
>
>    ```
>    PS C:\Users\82341> $env=1,2,3
>    PS C:\Users\82341> $env='Prod','QA'
>    PS C:\Users\82341> $env=Prod,QA
>    所在位置 行:1 字符: 10
>    + $env=Prod,QA
>    +          ~
>    参数列表中缺少参量。
>        + CategoryInfo          : ParserError: (:) [], ParentContainsErrorRecordException
>        + FullyQualifiedErrorId : MissingArgument
>    ```
>
> 2. 调用函数与jquery类似
>
> 3. ==powershell 中 ` 表示手动行换但是不推荐使用可能会报错，powershell会自动识别命令然后换行。==
>
>    https://stackoverflow.com/questions/3235850/how-to-enter-a-multi-line-command

powershell支持如下几种特性：

- Robust command-line [history](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_history)
- Tab completion and command prediction (See [about_PSReadLine](https://docs.microsoft.com/en-us/powershell/module/psreadline/about/about_psreadline))
- Supports command and parameter [aliases](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_aliases)
- [Pipeline](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_pipelines) for chaining commands
- ==In-console [help](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/get-help) system, similar to Unix `man` pages==

## get-command

==powershell中的命令被称为cmdlet，是一个类==

==可以通过`Get-Command | more`获取所有可以使用的命令==，支持通配符

```
PS C:\Users\82341> Get-Command -Name *Disk

CommandType     Name                                               Version    Source
-----------     ----                                               -------    ------
Function        Add-PhysicalDisk                                   2.0.0.0    Storage
Function        Clear-Disk                                         2.0.0.0    Storage
Function        Clear-StorageBusDisk                               1.0.0.0    StorageBusCache
```

具体可以查看：https://docs.microsoft.com/en-us/powershell/scripting/learn/tutorials/01-discover-powershell?view=powershell-7.1

```
#类似于unix 中的 head
PS C:\Users\82341> Get-Command | Select-Object -First 3

CommandType     Name                                               Version    Source
-----------     ----                                               -------    ------
Alias           Add-AppPackage                                     2.0.1.0    Appx
Alias           Add-AppPackageVolume                               2.0.1.0    Appx
Alias           Add-AppProvisionedPackage                          3.0        Dism

#进程名为p*
PS C:\Users\82341> Get-Process | Where-Object {$_.ProcessName -Like "p*"}

Handles  NPM(K)    PM(K)      WS(K)     CPU(s)     Id  SI ProcessName
-------  ------    -----      -----     ------     --  -- -----------
    124      16    13904      20692              1140   0 pacjsworker
    852      35   143636     166036       6.00  14068   3 powershell
```

## get-member

可以获取cmdlet==对应数据类型==的成员变量

```
PS C:\Users\82341> get-process|get-member


   TypeName:System.Diagnostics.Process

Name                       MemberType     Definition
----                       ----------     ----------
Handles                    AliasProperty  Handles = Handlecount
Name                       AliasProperty  Name = ProcessName
NPM                        AliasProperty  NPM = NonpagedSystemMemorySize64
PM                         AliasProperty  PM = PagedMemorySize64
SI                         AliasProperty  SI = SessionId
VM                         AliasProperty  VM = VirtualMemorySize64
WS                         AliasProperty  WS = WorkingSet64
Disposed                   Event          System.EventHandler Disposed(System.Object, System.EventArgs)
ErrorDataReceived          Event          System.Diagnostics.DataReceivedEventHandler ErrorDataReceived(System.Obje...
Exited                     Event          System.EventHandler Exited(System.Object, System.EventArgs)
OutputDataReceived         Event          System.Diagnostics.DataReceivedEventHandler OutputDataReceived(System.Obj...
BeginErrorReadLine         Method         void BeginErrorReadLine()
BeginOutputReadLine        Method         void BeginOutputReadLine()
CancelErrorRead            Method         void CancelErrorRead()
CancelOutputRead           Method         void CancelOutputRead()
Close                      Method         void Close()
CloseMainWindow            Method         bool CloseMainWindow()
```

按照名字获取cmdlet的成员变量

```
PS C:\Users\82341> get-process|get-member  -name *name


   TypeName:System.Diagnostics.Process

Name        MemberType    Definition
----        ----------    ----------
Name        AliasProperty Name = ProcessName
__NounName  NoteProperty  string __NounName=Process
MachineName Property      string MachineName {get;}
ProcessName Property      string ProcessName {get;}
```

按照成员变量的类型获取cmdlet的成员变量

```
PS C:\Users\82341> get-process|get-member -membertype property -name *name


   TypeName:System.Diagnostics.Process

Name        MemberType Definition
----        ---------- ----------
MachineName Property   string MachineName {get;}
ProcessName Property   string ProcessName {get;}
```







