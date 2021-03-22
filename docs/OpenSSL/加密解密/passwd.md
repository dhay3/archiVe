# paswd

用于生成加密的密码

- `-crypt`

  使用crypt加密算法，缺省值

- `-1`

  使用md5加密算法

- `-5 | -6`

  使用sha256或sha12加密

- `-salt <string>`

  指定加密的salt

  ```
  root in /usr/local/\/ssl λ openssl passwd -stdin -6
  111
  $6$PuPYSfq7gn8YfQiW$ff6rL5fjZyMH3Vr5Ah/S6CZ6QSMh3KHYOLyZkM1Jx/DUgrt3PQHr/lTu2qJxs.fg7efKXz3widZ0Bc8LvsAFP0
  ```

- `-in <file>`

  从文件中读入密码

- `-stdin`

  从stdin中读入密码

  ```
  root in /usr/local/\/ssl λ openssl passwd -stdin
  111
  9ZLZm7Ll.FTi2
  ```

### 