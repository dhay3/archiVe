# gh gpg-key

ref
[https://cli.github.com/manual/gh_gpg-key](https://cli.github.com/manual/gh_gpg-key)
用于管理 github 上的 GPG key，需要 `admin:gpg_key` 权限

## Add
syntax
```
gh gpg-key add [<key-file>] [flags]
```
用于将 GPG key 上传到 github
假设现在有一个 trump GPG key 需要上传公钥到 github
```
λ ~/ gpg -k
/home/cpl/.gnupg/pubring.gpg
------------------------------------
pub   2048R/FABC8322 2023-04-06
uid                  Trump <trump@gmail.com>
sub   2048R/AE6A64A1 2023-04-06
```
首先需要导出 GPG key 公钥
```
λ ~/ gpg --armor --output trump.gpk --export trump
```
然后上传 GPG key 公钥，到 github 名字为 trump_gpk 
```
λ ~/ gh gpg-key add trump.gpk -t trump_gpk
✓ GPG key added to your account
```
### Optional args

- `-t | --title <string>`

指定 GPG key 以 string 的名字存储在 github 上
## List
列出 github 所有的 GPG key
```
λ ~/ gh gpg-key list 
EMAIL                                    KEY ID            PUBLIC KEY                                           ADDED         EXPIREStrump@gmail.com                          3A9CA6E6FABC8322  xsBNBGQuQTEBCADR9symL5pF...UB1cN9j7riI6GiNVABEBAAE=  2m            Never
hostlockdown@gmail.com                   A2E176C5FE41F19D  xsFNBGCzKlIBEADwAQKwps3+...hOxlp1Y0tS4K31GCwQARAQAB  May 30, 2021  Never
62749885+dhay3@users.noreply.github.com  F3A82ABD5E016AC9  xsFNBGCzUWUBEAC8y/yjwImG...yr2T/z1coV6y1E+CzwARAQAB  May 30, 2021  Never
```
## Delete
syntax
```
gh gpg-key delete <key-id> [flags]
```
用于删除 github 上指定 key-id 的 GPG key
```
λ ~/ gh gpg-key delete 3A9CA6E6FABC8322
X Sorry, your reply was invalid: You entered y
? Type 3A9CA6E6FABC8322 to confirm deletion: 3A9CA6E6FABC8322
✓ GPG key 3A9CA6E6FABC8322 deleted from your account
```
### Optional args

- `-y|--yes`
