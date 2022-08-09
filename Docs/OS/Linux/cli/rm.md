# Linux rm

## Digest

syntax：`rm [optin]... [file]...`

remove echo specified file. By default, it does not remove directories

## Optional args

- `-f | --force`

  强制删除

- `-i`

  删除文件前提示

- `-I`

  当删除文件个数大于3个或递归删除时提示

- `-r | -R | --recursive`

  rm 默认不删除目录，使用该参数递归删除

- `-d`

  删除空文件

- `-v`

  verbose

## Caution

1. `rm -f * `默认不会删除隐藏文件，删除隐藏文件使用`rm -f .*`