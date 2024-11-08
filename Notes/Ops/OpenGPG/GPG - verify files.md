# GPG - verify files

ref
[https://www.gnupg.org/faq/gnupg-faq.html#how_do_i_verify_signed_packages](https://www.gnupg.org/faq/gnupg-faq.html#how_do_i_verify_signed_packages)
[https://www.gnupg.org/gph/en/manual/x135.html](https://www.gnupg.org/gph/en/manual/x135.html)
现在需要对文件签名进行校验，不管是明文的还是 binary 的都能处理

## 签名不分离
文件内容和签名不分离
```
root@v2:~# cat file.asc 
-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA512

this is a file gonna to be signed
-----BEGIN PGP SIGNATURE-----

iQHCBAEBCgAsFiEESYAZEbmEIgUfaqqGojYFpuQpJ60FAmQul14OHHRlc3RlckBx
cS5jb20ACgkQojYFpuQpJ61v4gwApcnUjaVRw4Pl9Pnn8hBsveHxjc/kFR3l7M1P
pWTY0bRO+n59gdOF/P5g0jW7gVprd2KRaBz65dpyxAxsNAsShULLJNY+iIRiST2U
8uk1UufcwaVna1o79jFNwuE5R5f6fRmClITSWrMKaUTU8eJ12/1krPsJAaXJi8qA
LbhQGut9iMS1bA6iGJP/s3nNUcqeUm8z2Nw6JJsHyOE06OtW1GWYQCLcjoNkRuLo
0kI+1SzUzGVAIdsWb0F723LpgWZFqEMo2dszKvBQ9Le77+A5qDrcSHCprtaU3hsl
1iTvfiu45ieXe1oBtB1ek1f3Kfm4NlnsfF9HVf9dvseuHrQhIuOBtkEjHFYzsHLQ
hHMPuWmTxFH/DN+aa07ElGFL+/XCRJcfGn4bC11zdqovKevWKMZ1Qp1m0GvWi+My
/sSgR9RFGfL+i7AlUYaTPhUYIW85KdiPR7LKGmJg0F9UY4y8UWKcVxbRS70r/Kbb
07w1zspbl2I8GJxeWgQdPjPTifsI
=hGiY
-----END PGP SIGNATURE-----
```
文件内容和签名不分离时，可以使用 `--verify <signature-file>` 来校验，当时需要注意的一点是去掉 `.asc` 后的文件名不能和被签名的文件名相同，否则就会校验失败。因为当前目录中有相同的文件名时，GPG 就认为当前的文件是一个分离的签名文，所以就有问题了
```
root@v2:~# gpg --verify file.asc
gpg: Signature made Fri 07 Apr 2023 05:30:12 PM CST
gpg:                using RSA key 49801911B98422051F6AAA86A23605A6E42927AD
gpg: Good signature from "tester (this is a comment) <tester@qq.com>" [unknown]
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: 4980 1911 B984 2205 1F6A  AA86 A236 05A6 E429 27AD
gpg: WARNING: not a detached signature; file 'file' was NOT verified!
root@v2:~# cp file.asc file.asc~
root@v2:~# gpg --verify file.asc~
gpg: Signature made Fri 07 Apr 2023 05:30:12 PM CST
gpg:                using RSA key 49801911B98422051F6AAA86A23605A6E42927AD
gpg: Good signature from "tester (this is a comment) <tester@qq.com>" [unknown]
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: 4980 1911 B984 2205 1F6A  AA86 A236 05A6 E429 27AD
```
因为这种特殊的方式，所以在生成签名时 推荐采用 detach 签名
## 签名分离
文件内容和签名分离
```
root@v2:~# cat file 
this is a file gonna to be signed
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
如果签名文件去掉 `.asc` 后在当前目录中有相同名字的文件时，可以使用`--verify <detach-signature>`来校验
```
root@v2:~# gpg --verify file.asc file
gpg: Signature made Thu 06 Apr 2023 05:49:55 PM CST
gpg:                using RSA key 49801911B98422051F6AAA86A23605A6E42927AD
gpg:                issuer "tester@qq.com"
gpg: Good signature from "tester (this is a comment) <tester@qq.com>" [ultimate]
```
当需要校验的文件和签名文件去掉 `.asc` 仍不同，或者不在一个目录时，可以使用`--verify <detach-signature> <signed-file>`来校验
```
root@v2:~# gpg --verify file.asc f
gpg: Signature made Thu 06 Apr 2023 05:49:55 PM CST
gpg:                using RSA key 49801911B98422051F6AAA86A23605A6E42927AD
gpg:                issuer "tester@qq.com"
gpg: Good signature from "tester (this is a comment) <tester@qq.com>" [ultimate]
```
也可以不使用任何参数只指定分离的签名文件，GPG 会自己推测使用的命令
```
root@v2:~# gpg file.asc
gpg: WARNING: no command supplied.  Trying to guess what you mean ...
gpg: assuming signed data in 'file'
gpg: Signature made Fri 07 Apr 2023 01:50:40 PM CST
gpg:                using RSA key 49801911B98422051F6AAA86A23605A6E42927AD
gpg:                issuer "tester@qq.com"
gpg: Good signature from "tester (this is a comment) <tester@qq.com>" [ultimate]
```
