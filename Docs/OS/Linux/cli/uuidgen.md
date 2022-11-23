# Linux uuidgen

## Digest

> 也可以通过系统文件生成 uuid
>
> `cat /proc/sys/kernel/random/uuid`

`uuid` 用于随机生成一个 UUID (universally unique identifier)，可以是以下 3 种类型的 UUID

1. time-based UUID
2. random-based UUID ( default )
3. hash-based UUID 部分版本不支持

## Optional args

- `-t | --time`

  生成 time-based UUID

- `-r | --random`

  生成 random-based UUID

- `-m | --md5`

  `-s | --sha1`

  hash-based UUID

  使用 md5 或者 sha1 加密生成 UUID