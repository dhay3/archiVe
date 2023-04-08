# GPG cli

## Digest

syntax

```
gpg [--homedir dir] [--options file] [options] command [args]
```

## Options

### Signature options

- `--sign | -s message`

  对文件签名，可以和 `--encrypt` 一起使用，表示对 message 签名加密

  也可以和 `--symmetric` 一起使用，表示使用对称加密 message

  ==如果没有使用 `--local-user` 或者 `--default-key` gpg 会自动选择签名使用的 user-id，即 `-k` 中出现的一个gpg key==

-  `--clear-sign | --clearsign`

  签名以明文的形式显示，不需要使用专有的软件读取

- `--detach-sing | -b`

  生成独立的签名


- `--default-key name`

  指定签名使用的 user-id，如果没有使用该参数，默认使用 `-k` 发现的第一个 GPG key。该参数会被 `-u | --local-user` overrides

- `--local-user | -u name`

  指定签名使用的 user-id

- `--verify `

  校验文件的签名。如果只有一个 argument，默认为带签名的文件，如果有多个 argument，第一个为 detached signature

### Encryption Options

- `--encrypt | -e`

  使用一个或者多个 public key 加密 message，可以和 `--sign` 或者 `--symmetric` 一起使用

- `--symmetric | -c`

  使用对称加密，默认使用 `AES-128`。可以和 `--sign` 或者 `--encrypt` 一起使用。gpg 默认会 cache symmetric passphrase，可以使用 `--no-symkey-cache` 取消该特性

- `--decrypt | -d file`

  decrypt 文件，可以使用 `--ouput` 输出到指定文件，默认输出到 stdout。如文件有签名，签名也会被一起校验

- `--recipient | -r name`

  指定加密使用的 user-id 公钥，即接受加密信息这的公钥

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

- `--edit-key key`

  修改指定的 key 的参数，具体参考 manual page

  修改完后需要使用 `save` 保存，否则不会生效。尽量使用该参数来对 GPG key 做修改的操作

- `--quick-add-key fpr [algo[usage[expire]]]`

  添加 subkey 到指定的 key

- `--quick-add-uid user-id new-user-id`

  添加指定的 userid 到 key

- `--quick-set-expire fpr expire`

  修改指定 fingerprint  key 的有效时间

- `--change-passphrase | --passwd user-id`

  修改指定 key 的 passphrase

#### delete

- `--delete-keys name`

  从 public keyring 中移除指定 key，需要先删除 secret key 才能删除 public key

- `--delete-secret-keys name`

  从 secret keyring 中移除指定 key

- `--delete-secret-and-public-key name`

  同时从 public keyring 和 secret keyring 中同时删除指定的 key

#### import/export

- `--export [name]`

  导出公钥，如果没有指定，默认导出所有到 stdout。可以和 `--output` 或 `--armor` 一起使用。需要注意的一点是在 GPG 中 subkey 的公钥和 primary 的公钥是绑定在一起的，所以会被一起导出

- `--export-secret-keys | --export-secret-subkeys [name]`

  和 `--export` 类似，但是导出私钥

- `--export-ssh-key name`

  以 OpenSSH 格式导出 public key，需要开启 authentication 的功能才可以。可以使用 `--edit-keys` 中的 `change-usage` 打开

- `--import | --fast-import`

  导入 keys, 不分公钥还是私钥

### Keyserver options

- `--send-keys keyIDs`

  和 `--export` 类似，但是会将 key 发送到 keyserver，一旦发送到了 keyserver，几乎就无法被删除了


- `--keyserver name`

  用于指定使用的 keyserver，即 `--send-keys`, `--receive-keys`, `--search-keys` 关联使用的 keyserver

  如果在 dirmngr 配置文件中没有指定明确的 keyserver，默认会使用 https://keyserver.ubuntu.com 作为 keyserver 

  因为 GPG 的版本不同，允许的 scheme 也不同，为了减少出现错误的情况，应该尽量使用 `hkps` scheme

  大多数的 keyservers 会和其他 keyservers 会自动同步 GPG keys，所以不需要上传至多个 GPG keys

  在 GPG 2.1 之后应该需要使用 `dirmngr.conf` 中的 keyserver 来替代该 option

- `--receive-keys | --recv-keys keyIDs`

  从 keyserver 导入指定的 key

- `--refresh-keys`

  从 keyserver 同步 local keyring，直接理解成类似 `git pull` 的操作

- `--search-keys names`

  从 keyserver 搜索指定的 key，为了搜索精度 name 可以是邮件地址

- `--fetch-keys URIs`

  从指定 URIs 中 retrieve keys

### Input and output options

- `--armor | -a`

  以 ASCII 格式输出需求的内容，默认以 binary 格式输出

- `--output | -o file`

  将内容输出到 file

- `--with-colons`

  以便于程序读取的方式输出

  ```
  root@v2:~# gpg --with-colons -ktru::1:1680964089:0:3:1:5
  pub:-:2048:1:AE075F8C687EE529:1680856692:::-:::scESC::::::23::0:
  fpr:::::::::9CF6872151584CAD207AF4A9AE075F8C687EE529:
  uid:-::::1680856692::D786769A4820DF66198BB68C45AFBD69AB2D5005::Bobby <bobby@gmail.com>::::::::::0
  ```

  例如可以使用如下删除所有的 GPG key

  ```
  root@v2:~# keys=$(gpg --with-colons -k | sed -n '/pub/p' | awk -F : '{print($5)}')
  root@v2:~# gpg --delete-secret-keys ${keys} && gpg --delete-keys ${keys}
  ```

### Misellaneous options

- `--gen-random 0|1|2 count`

  随机生成指定 count type 大小的字节数

- `--yes`

  默认使用yes做为大多数问题的答案

- `-v | --verbose`
- `-n | --dry-run`

- `--debug`