## strftime

strftime是c# time.h中的一个函数，python中也有，经常结合tcpdump一起使用

- %y

  the year as a decimal number without a century 00 -99

- %Y

  the year as a decimal number including the century

- %m

  month as a decimal number 01 - 12

- %d

  day of month as a decimal number 01 - 31

- %H

  hour as a decimal number 00 - 23

- %M

  minute as a decimal number 00 -59

- %S

  second as a decimal number 00 - 59

例如：

```
root in /home/ubuntu λ tcpdump -i eth0 -G 10 -W 3 -w %y%m%d%H%M%S

root in /home/ubuntu λ ls
210928145109  210928145131  210928145120  
```