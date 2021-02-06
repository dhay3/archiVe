# wget

## 概述

默认将输出的内容保存成文件

wget用于从web站点下载文件，wget也可以拿来做爬虫，但是需要遵循robots.txt文件的规则

```
root in /etc/cron.d λ wget -q -O- https://pastebin.com/raw/e8XzcU2Q
echo "data from pastebin"
```

- `-q`表示关闭wget输出
- `-O-`表示将输出显示在stdout