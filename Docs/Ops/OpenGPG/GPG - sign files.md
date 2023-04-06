GPG 最重要的一部分就是对文件签名
假设现在有一个文件 file
```
root@v2:~# cat file
this is a file gonna to be signed
```
## 二进制签名不分离
现在需要对该文件使用 GPG 签名，就需要使用到 `-s | --sign` option
```
root@v2:~# gpg --verbose -s file
gpg: using pgp trust model
File 'file.gpg' exists. Overwrite? (y/N) y
gpg: writing to 'file.gpg'
gpg: RSA/SHA512 signature from: "0FC2114BEEB4EBF8 alice (this is a comment) <alice@yahoo.com>
```
如果没有指定 `--local-user`，GPG 默认选择 `gpg -k`中出现的第一个 key 作为签名用的 GPG key
```
root@v2:~# gpg -v --local-user test -s file
gpg: using subkey D8573B042FA8A678 instead of primary key A23605A6E42927AD
File 'file.gpg' exists. Overwrite? (y/N) y
gpg: writing to 'file.gpg'
gpg: DSA/SHA256 signature from: "D8573B042FA8A678 tester (this is a comment) <tester@qq.com>"
```
上述命令会生成一个 `.gpg` 后缀的文件包含文件内容和签名，以 binary 格式存储
```
root@v2:~# ls
file  file.gpg
```
## 明文签名不分离
如果需要生成明文的文件，可以使用 `--clearsign` (无须和 `--sign` 一起使用)
```
root@v2:~# gpg -v --local-user tester --clearsign file
gpg: writing to 'file.asc'
gpg: pinentry launched (296520 curses 1.1.0 /dev/pts/0 xterm-256color -)
gpg: RSA/SHA512 signature from: "A23605A6E42927AD tester (this is a comment) <tester@qq.com>"
```
上述命令会生成一个 `.asc` 后缀的文件包含文件内容和签名，以明文的方式显示
```
root@v2:~# ls
file  file.asc
root@v2:~# cat file.asc 
-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA512

this is a file gonna to be signed
-----BEGIN PGP SIGNATURE-----

iQHCBAEBCgAsFiEESYAZEbmEIgUfaqqGojYFpuQpJ60FAmQuki8OHHRlc3RlckBx
cS5jb20ACgkQojYFpuQpJ61aTQv+NI/ErUSlPE1yGOZtjpUgdg19eXEFM+sOxabl
7hs28k4uvw/ZisOuZ0iOqr+mw1gfk7X+o9RIwUfyNRPGvn+Jbee77RZEn97mU2RB
/AG/Zvhsh4Y0K9N2UMV4L394JKUz70zT+i/I9POlEeyDgtVTZ1yAg4P4xxLSFA+S
6a5bFfu+4tTeaUP4aDxwZ400fBI84XdkbHECgXfcmpHVUHaftS72CNftXz3nlmUR
NGMO8YoEhMIwBoS2010wYIO8GDJadI2C1eR7PtcoSTygBvfg8/qLfT5uKkslDjMe
LJYLzKPyoeZyEc/yGO4Y7RzJ/30IyNKeS+tMk++WVcqefb5Z5GRLBqmpC6X6NzTs
6twi6BC24qyYiledqZLoJHzlZFDsI76hXiWARO92eVGuEBa+QsdSLKkUOu43wO+T
GwUz13MuesuWEDU6BTgIQzwmiGgK38nj0HYMHJEpZsAXWOD95Cm0iXoyNZUWURRC
nrYlhy84Dk0YUVbxPBpCWIDNDgP/
=Xjrb
-----END PGP SIGNATURE-----
```
## 二进制签名分离
如果需要单独将签名文件和文件内容分离，可以使用 `--detach-sign`
```
root@v2:~# gpg -v --local-user tester --detach-sign  file
gpg: writing to 'file.sig'
gpg: RSA/SHA512 signature from: "A23605A6E42927AD tester (this is a comment) <tester@qq.com>"
```
上述命令会生成一个 `.sig` 文件只包含文件签名，以 binary 格式存储
```
root@v2:~# ls
file  file.sig
```
## 明文签名分离
如果需要以明文的方式显示签名的内容，并分离文件内容和签名，在二进制签名分离的基础使用 `--armor` (而不是和 `--clear-sign` 一起使用)
```
root@v2:~# gpg -v --local-user tester --detach-sign --armor  file
File 'file.asc' exists. Overwrite? (y/N) y
gpg: writing to 'file.asc'
gpg: RSA/SHA512 signature from: "A23605A6E42927AD tester (this is a comment) <tester@qq.com>"
```
上述命令会生成一个 `.sig` 文件只包含签名文件，以明文存储
```
root@v2:~# cat file.asc 
-----BEGIN PGP SIGNATURE-----

iQHCBAABCgAsFiEESYAZEbmEIgUfaqqGojYFpuQpJ60FAmQulAQOHHRlc3RlckBx
cS5jb20ACgkQojYFpuQpJ60qoAwA06lywgO4VbQwHxGkHBtF5R53b2u8BOEuL76J
FDgsrz2LXPQ/1cVepB8Ct8CyfiWTSsciJjR3LmFR5X/lC361AV0reU4F8HhIbOW7
wjNixVwUf5OYxw9Mg/kSQj2XYfuCUArX5oyIoNINSveAiJ61QFI1fK2kGa93THbx
xqPwb2ioOXup0E4r9WiaDmtTno4E20bP9CST5hpPA8SRKWHE4VkFdosOyZDkFeyF
yzOwHuraWX6fM3yxnVReFEG/zjgBCCjn+U4SPqxt+pm5TO9HcHHBJFJ5tany9wmU
CUzqwRw30ouGcpEZsMoKmkLyQmfHWXmQ/uuxJWNWRhSltNHs+7mBdwuRmJt8WyeR
l5O7wDrvWqj2HBN+4gwpLymRhzhHMH6dNUd9mXB/PHwE+gj+wipwR769P8dNov+F
J0DRwNmDBFUlnM5wrif+CXKPhup0nGh7VN1xzVuyX+juJ7tq2aZ1bxD0u3n+t0/S
LtCaAvXzfz1YUhu5pO5ej8hxfHh2
=rGd5
-----END PGP SIGNATURE-----
```
