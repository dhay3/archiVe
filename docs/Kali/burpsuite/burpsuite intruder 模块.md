# burpsuite intruder 模块

![Snipaste_2020-09-06_10-02-10](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-09-06_10-02-10.6fqh6g2gjbo0.png)

## Sniper

在一个position打payload， 没有标记marker的不会被影响

```go
for i:= 0; i < len(payload position); i++ {
    tmp := payload position[i]
    for j:=0; j < len(payloads); j++ {
        payload position [i] = payloads[j]
    }
    payload position[i] = tmp
}
```

![Snipaste_2020-09-06_10-03-28](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-09-06_10-03-28.3r76ugh7i1u0.png)

以上图为例子，从第一个position开始，做1-10遍历，然后下一个position做1-10遍历。

> 以下模式，需要设置payload set

## Battering ram

在不同的payloads position打相同的payload

```go
for i:=0; i < len(payloads);i++{
	for j:=0; j < len(payloads postions) ; j++{
		payload postions[j] = payloads[i]
	}
}
```

https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-09-06_10-15-17.4gnshw3myf00.png

所有的位置都被替换为同一发payload

## Pitchfork

在不同payload position 打不同的payload set(最大上限20个)，有多少个payload position就有多少个payload set。如果payload set长度不一样选取最小的做为request count

```go
for i:=0; i < len([0][]payload set); i++{
	for j:=0 ; j < len(payload position); j++{
		payload postions[j] = payload set[j][i]
	}
}
```

![Snipaste_2020-09-06_10-34-57](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-09-06_10-34-57.14n6vmyz84e8.png)

第一列表示position，右边表示每次发送的payload

## Cluster bomb

全排序，有多少个payload position 就有多少个payload set

```c
for(int i=k;i<=m;i++){
	swap(list[i],list[k]);
	Perm(list,k+1,m);
	swap(list[i] , list[k]);
}
```

![Snipaste_2020-09-06_10-54-48](https://cdn.jsdelivr.net/gh/dhay3/image-repo@master/20210518/Snipaste_2020-09-06_10-54-48.3mwekk69ue60.png)

假设 position[0] = 1-1，position[1] = 2-3, position[2] = 3 - 5

就有一共 1 * 2 * 3种可能

