# OpenSSH files

ref

https://man.openbsd.org/ssh.1

- `~/.ssh`

  this directory is the default location for all user-specific configuration and authentication information

- `~/.ssh/authorized_keys`

  list the ==public keys== that can be used for logging in as this user

  记录可以连接并以当前用户登录机器的客户端公钥( 只有开启公钥认证的方式才会此文件，如果以密码的方式不会显示该文件 )

- `~/.ssh/known_hosts`

  contains a list of host keys for all hosts the user has logged into that are not already in the systemwide list of known host keys

  记录当前用户登录过的 ssh 终端的公钥和信息

  ```
  [82.157.1.137]:65522 ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOVGjNw4pme6wMkRxDDWIYLTviyTYDPvlHeWkipb37FXBjZjS+3cZe4coXNngAplW0i1tTgkoA95PQxwU75+jAk=
  ```

- `~/.ssh/config`

  针对 per-user 的 ssh 客户端配置文件

- `~/.ssh/rc`

  commands in this file are executed by ssh when the user logs in, just before the user’s shell is started