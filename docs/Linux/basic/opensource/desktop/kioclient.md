# kioclient

## Digest

syntax：`kioclient [options] command urls...`

kioclient 是 KDE 中的一个 CLI，用于 network-transparent 以及 local-transparent

## Positional args

- `exec URL`

  使用特定的软件(根据L7协议)打开URL指定的文档

  ```
  #使用默认浏览器打开
  kioclient exec https://baidu.com
  #打开回收站
  kioclient exec trash:/
  #打开指定目录
  kioclient exec ~/note
  #也可以在本地目录前添加file协议，效果同上
  kiolienct exec file:/tmp/a
  #打开当前目录
  kioclient .
  ```

- `move|mv SRC DEST`

  将src的内容移动到dest

  ```
  #将文件快速移动到回收站
  kioclient move /tmp/a trash:/
  ```

- `kioclient download URL`

  下载指定URL的内容，基于KDE可以替换wget

- `kioclient copy|cp SRC DEST`

  复制文件

- `kioclient cat URL`

  查看URL中的内容输出到stdout

- `kioclient ls|remove|stat URL`

  等同与linux中的相同的命令，大多数remote URL不支持查看这些内容



