# GPG cli

## Digest

syntax

```
gpg [--homedir dir] [--options file] [options] command [args]
```

## Optional args

- `--sign | -s message`

  可以和 `--encrypt` 一起使用，表示对 message 签名加密

  也可以和 `--symmetric` 一起使用，表示使用对称加密 message

  `--encrypt` 也可以和 `--symmetric` 一起使用

  ==如果没有使用 `--local-user` 或者 `--default-key` gpg 会自动选择签名使用 signing key==

-  `--clear-sign | --clearsign`

  签名以明文的形式显示，不需要使用专有的软件读取

- `--detach-sing | -b`

  生成独立的签名

- `--encrypt | -e`

  使用一个或者多个 public key 加密 message，可以和 `--sign` 或者 `--symmetric` 一起使用

- `--symmetric | -c`

  使用对称加密，默认使用 `AES-128`。可以和 `--sign` 或者 `--encrypt` 一起使用。gpg 默认会 cache symmetric passphrase，可以使用 `--no-symkey-cache` 取消该特性

- `--decrypt | -d file`

  decrypt 文件，可以使用 `--ouput` 输出到指定文件，默认输出到 stdout。如文件有签名，签名也会被一起校验

- `--verify `

  校验文件的签名。如果只有一个 argument，默认为带签名的文件，如果有多个 argument，第一个为 detached signature

- `--export [name]`

  从 keyring 中导出指定的 key，如果没有指定，默认导出所有到 stdout。可以和 `--output` 或 `--armor` 一起使用

- `--export-secret-keys | --export-secret-subkeys [name]`

  和 `--export` 类似，但是导出 secret keys

- `--export-ssh-key name`

  以 OpenSSH 格式导出 public key

- `--import | --fast-import`

  导入 keys

### manage keys options

#### generation

- `--quick-generate-key | --quick-gen-key user_id [algo [usage [expire]]]`

  快速生成 gpg key，无须设定邮箱

- `--generate-key | --gen-key`

  生成 gpg key，需要设定邮箱

- `--full-generate-key | --full-gen-key`

  对 `--generate-key` 扩展

- `--generate-revocation | --gen-revoke user_id`

  生成 gpg key revocation certificate

  如果需要 revoke gpg key，可使用 `--import` 导入生成的 revocation certificate，同时通过 `--send-key` 撤回 keyserver 上对应的 key

#### query

- `--list-keys | --list-public-keys | -k [user_id]`

  显示指定的 public keys，如果没有指定，默认显示所有的

- `--list-secret-keys | -K [user_id] `

  显示指定的 secret key，如果没有指定，默认显示所有的

- `--fingerprint [user_id]`

  显示指定 key fingerprint，如果没有指定，默认显示所有的

#### update

- `--quick-add-uid user-id new-user-id`

  添加指定的 userid 到 key

- `--quick-set-expire fpr expire`

  修改指定 fingerprint  key 的有效时间

- `--change-passphrase | --passwd user-id`

  修改指定 key 的 passphrase

- `--edit-key key`

  修改指定的 key 的参数，具体参考 manual page

  修改完后需要使用 `save` 保存，否则不会生效

#### delete

- `--delete-keys name`

  从 public keyring 中移除指定 key，需要先删除 secret key 才能删除 public key

- `--delete-secret-keys name`

  从 secret keyring 中移除指定 key

- `--delete-secret-and-public-key name`

  同时从 public keyring 和 secret keyring 中同时删除指定的 key

### keyserver options

- `--send-keys keyIDs`

  和 `--export` 类似，但是会将 key 发送到 keyserver，一旦发送到了 keyserver，就无法被删除了

- `--receive-keys | --recv-keys keyIDs`

  从 keyserver 导入指定的 key

- `--refresh-keys`

  从 keyserver 同步 local keyring

-  `--search-keys names`

  从 keyserver 搜索指定的 key，为了搜索精度 name 可以是邮件地址

- `--fetch-keys URIs`

  从指定 URIs 中 retrieve keys

### input and output options

- `--armor | -a`

  以 ASCII 格式输出需求的内容，默认以 binary 格式输出

- `--output | -o file`

  将内容输出到 file

- `--with-colons`

### misellaneous options

- `--gen-random 0|1|2 count`

  随机生成指定 count type 大小的字节数