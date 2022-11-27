# Linux socat

## Digest

syntax：`socat [options] <address> <address>`

socket cat 可以看成是netcat+socket的升级版本，socat 有 4 个阶段，init，open，transfer，closing

the first address is used by socat for reading data, and the second address for writing data

第一个address用于读，第二个address用于写(其实是双向通信的，可用读也可以写)

## Terms

### ADDRESS

ADDRESS := ADDRESS TYPE + ADDRESS OPTIONS

address 通常由如下3个元素组成 

1. address type keyword，例如 TCP4, OPEN, EXEC。keyword 不区分大小写
2. 零个或多个由`：`分隔的 address parameters，可选值由 address type 决定
3. 零个或多个由`，`分隔的address options

### ADDRESS TYPES

ADDRESS TYPES := keyword + parameters

#### FILE TYPE

- `CREATE:<filename>`

  使用create()函数打开文件，用于写IO. filename 必须一个实际存在的路径

  Option grous: FD, REG, NAMED



- `EXEC:<command-line>`

  



- `IP-SENDTO:<host>:<protocol>`

  打开raw IP socket

## Optional args

- `-h | -hh | -hhh`

  打印帮助信息，h越多越详细

- `-d`

  debug，d越多debug信息越多

- 