# Shell <<<

reference:

https://unix.stackexchange.com/questions/80362/what-does-mean

`command <<< “input”` 等价与 `echo input | command`

例如：

```
awk '{print(0+$1);if(0+$1 > 3.5){print 1;}else{print 2;}}' <<< "200.5"
200.5
```

