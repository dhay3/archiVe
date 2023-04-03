# Github GPG

参考：

https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/managing-commit-signature-verification

使用GPG签名来校验用户的身份，==但是如果本地生成的密钥对被删除了，将不能上传文件==

> 注意这里需要git bash中生成gpg密钥对，才会生效

1. 创建GPG密钥对`gpg --full-generate-key`，需要记住passpharse，在`git commit`时需要使用该密码。长度必须为4096，邮箱必须是github上认证过的或是noreply email

2. 将==公钥==（`gpg --armor --export <key id>`）导入到github中

3. 配置全局参数

   ```
   git config --global user.signingkey <fingerprint>
   ```

4. 在`git commit`时指定`-S`对`commit`签名



