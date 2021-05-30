# Select-Object

select-object 类似于Unix中的`grep + tail + head + unique` 的结合

```
PS C:\Users\82341> $data=1,2,3,4,5
```

## First

类head

```
PS C:\Users\82341> $data | Select-Object -First 2
1
2
```

## Last

类tail

```
PS C:\Users\82341> $data | Select-Object -Last 2
4
5
```

## property

如果对象是一个hashtable，可以获取指定的属性值

```
PS C:\Users\82341> $data=@{
>> name='zs'
>> age=23
>> }

#具体可以获取什么property，通过 $data | Get-Member
PS C:\Users\82341> $data | Select-Object -Property Keys

Keys
----
{age, name}
```









