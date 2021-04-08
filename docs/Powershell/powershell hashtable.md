# powershell hashtable

参考：

https://docs.microsoft.com/en-us/powershell/scripting/learn/deep-dives/everything-about-hashtable?view=powershell-7.1

以key/value形式存储，和redis中的hash一样

`help hashtable`

syntax：

```
PS C:\Users\82341> $hashtable=@{k1='v1'
>> k2 = 'v2'
>>
```

## 创建

==不需要手动 添加` 换行符，powershell会自动识别语法然后换行==

```
$environments = @{
    Prod = 'SrvProd05'
    QA   = 'SrvQA02'
    Dev  = 'SrvDev12'
}
```

## 取值

```
PS C:\Users\82341> $environments

Name                           Value
----                           -----
Prod                           SrvProd05
Dev                            SrvDev12
QA                             SrvQA02

PS C:\Users\82341> $environments['QA']
SrvQA02

PS C:\Users\82341> $environments[@('QA','Dev')]
SrvQA02
SrvDev12

PS C:\Users\82341> $environments[('QA','Dev')]
SrvQA02
SrvDev12
```

## 获取长度

```
PS C:\Users\82341> $environments.count
3
```

## 遍历

遍历keys，`$_`表示当前对象

```
PS C:\Users\82341> $environments.foreach({$_.keys})
Prod
Dev
QA
```

遍历values

```
PS C:\Users\82341> $environments.foreach({$_.values})
SrvProd05
SrvDev12
SrvQA02
```

## 添加kv对

```
PS C:\Users\82341> $person=@{ name='chz'
>> age=33}
PS C:\Users\82341> $person

Name                           Value
----                           -----
age                            33
name                           chz


PS C:\Users\82341> $person.city = 'jh'
PS C:\Users\82341> $person

Name                           Value
----                           -----
name                           chz
age                            33
city                           jh


PS C:\Users\82341> $person.add('gender','male')
PS C:\Users\82341> $person

Name                           Value
----                           -----
name                           chz
age                            33
gender                         male
city                           jh
```

## 修改kv对

```
PS C:\Users\82341> $person['age']=99
PS C:\Users\82341> $person

Name                           Value
----                           -----
name                           chz
age                            99
gender                         male
city                           jh
```

## 校验

```
if( $person.age ){...}
if( $person.age -ne $null ){...}
if( $person.ContainsKey('age') ){...}
```

## 删除kv对

```
$person.remove('age')
$person = @{}
$person.clear()
```

