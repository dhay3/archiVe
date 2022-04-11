# openssl passpharse

## passphrase-options

用于指定私钥的passpharse才会被用到

- `pass:password`

  私钥的密码。不安全因为是密文的，所以可以被`ps`查看到

- `env:var`

  私钥的密码是环境变量

- `file:pathname`

  文件中的第一行是密码

- `fd:number`

  指定私钥的密钥是文件句柄

- `stdin`

  从stdin中输入密码

