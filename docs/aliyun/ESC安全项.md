# ESC安全项

## sshd

1. 关闭密码校验和空密码登入，启用公钥校验登入

   ```
   PasswordAuthentication no
   PermitEmptyPasswords no
   PubkeyAuthentication yes
   ```

2. 关闭root用户登入

   ```
   PermitRootLogin yes
   ```

3. 