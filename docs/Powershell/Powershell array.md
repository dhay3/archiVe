# Powershell array

> powershell中大小写不敏感

参考：

https://docs.microsoft.com/en-us/powershell/scripting/learn/deep-dives/everything-about-arrays?view=powershell-7.1

syntax：`$arrayname=@(v1,v2,v4...)`

==可以使用`help array`查看manual page==

- 创建数组，都等价

  ```
  PS> $data = @(
      'Zero'
      'One'
      'Two'
      'Three'
  )
  
  PS> $data = @('Zero','One','Two','Three')
  
  PS> $data = 'Zero','One','Two','Three'
  ```

- 获取数组长度

  ```
  PS C:\Users\82341> $data=@(3,4,5)
  PS C:\Users\82341> $data.count
  3
  ```

- 遍历数组

  ```
  PS C:\Users\82341> $data
  3
  4
  5
  ```

- 获取下标对应的值

  ```
  PS C:\Users\82341> $data[1]
  One
  
  PS C:\Users\82341> $data[1,2,3]
  One
  Two
  Three
  
  PS C:\Users\82341> $data[1..3]
  One
  Two
  Three
  
  #reverse
  PS> $data[3..1]
  Three
  Two
  One
  
  #negtive
  PS C:\Users\82341> $data[-1]
  Three
  ```

- 数组越界

  ```
  PS C:\Users\82341> $null -eq $data[3]
  False
  PS C:\Users\82341> $null -eq $data[4]
  True
  ```

## forEach

和其他编程语言一样，foreach item的值可以是任意的

```
PS C:\Users\82341> foreach ( $node in $data )
>> {
>>     "Item: [$node]"
>> }
Item: [1]
Item: [2]
Item: [3]



#还可以结合管道符，但是item的值必须为psitem
PS C:\Users\82341> $data | foreach-Object {"Item: [$PSItem]"}
Item: [1]
Item: [2]
Item: [3]
#类似于jquery，$data.foreach{"Item [$PSItem]"}还可以去掉curly bracket
PS C:\Users\82341> $data.foreach({"Item [$psitem]"})
Item [1]
Item [2]
Item [3]

```

## 更新值

```
PS C:\Users\82341> for ( $index = 0; $index -lt $data.count; $index++ )
>> {
>>     $data[$index] = "Item: [{0}]" -f $data[$index]
>> }
PS C:\Users\82341> $data
Item: [Item: [Zero]]
Item: [Item: [One]]
Item: [Item: [Two]]
Item: [Item: [Three]]
```

## 赋值

powershell命令扩展和unix中不同，可以直接使用cmdlet即可

```
PS C:\Users\82341> $data=get-process
PS C:\Users\82341> $data[0].name
ApplicationFrameHost
```

获取成员变量

```
PS C:\Users\82341> $data=get-process|select-object -first 4
PS C:\Users\82341> $data

Handles  NPM(K)    PM(K)      WS(K)     CPU(s)     Id  SI ProcessName
-------  ------    -----      -----     ------     --  -- -----------
    572      32    33416      35072       0.66   9960   3 ApplicationFrameHost
    175      11     6488      12060       0.11  15444   0 audiodg
    140       9     1684       8864       0.05  11672   3 browser_broker
    601      31    39824       2124       0.55   7556   3 Calculator
```

## -contain and -in

校验数组中是否有值，返回true或false

```
PS D:\asset\note\docs\Powershell> $data = @('red','green','blue')
PS D:\asset\note\docs\Powershell> $data -contains 'green'
True
```

和contain的顺序相反

```
PS D:\asset\note\docs\Powershell> $data = @('red','green','blue')
PS D:\asset\note\docs\Powershell> $data -contains 'green'
True
```

## -eq and -ne

和shell中的用法不同

- eq遍历数组并比较值，如果相同就输出值。
- ne变量数组并比较值，如果不相同就输出值。

```
PS> $data = @('red','green','blue')
PS> $data -eq 'green'
green

PS> $data = @('red','green','blue')
PS> $data -ne 'green'
red
blue
```

和`if`一起使用时只要有一个值符合就返回true，并执行方法体

```
PS D:\asset\note\docs\Powershell> if ( $data -eq 'green' )
>> {
>>     'Green was found'
>> }
Green was found

PS D:\asset\note\docs\Powershell> if ( $data -eq 'lol' )
>> {
>>     'Green was found'
>> }
PS D:\asset\note\docs\Powershell>
```

## -match

会匹配在这个集合中的每一个item

```
PS D:\asset\note\docs\Powershell> $servers = @(
>>     'LAX-SQL-01'
>>     'LAX-API-01'
>>     'ATX-SQL-01'
>>     'ATX-API-01'
>> )
PS D:\asset\note\docs\Powershell> $servers -match 'SQL'
LAX-SQL-01
ATX-SQL-01
PS D:\asset\note\docs\Powershell>
```

## -join

和java中的`String.join`一样，如果数组只有一个值不生效

```
PS D:\asset\note\docs\Powershell> $data -join '-'
1-2-3-4
PS D:\asset\note\docs\Powershell> $data -join '-'
1
```

可以使用`$null`为取消分隔符

```
PS D:\asset\note\docs\Powershell> $data -join $null
1234
```

## -replace and -split

```
PS D:\asset\note\docs\Powershell> $data = @('ATX-SQL-01','ATX-SQL-02','ATX-SQL-03')
PS D:\asset\note\docs\Powershell> $data -replace 'ATX','LAX'
LAX-SQL-01
LAX-SQL-02
LAX-SQL-03

PS D:\asset\note\docs\Powershell> $data -split '-' -join $null
ATXSQL01ATXSQL02ATXSQL03
```























