---
createTime: 2024-12-03 10:42
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# OpenSSH 0 - Certificates

## 0x01 Preface

在第一次连接至 OpenSSH server 时，默认会让用户选择是否选择连接

```
$ ssh root@10.0.3.48
...
The authenticity of host '10.0.3.48 (10.0.3.48)' can't be established.
ED25519 key fingerprint is SHA256:WfL7sVGttLvi8v330i5aBcnlqoSNM5oGPaHv2TLlziE.
This key is not known by any other names.
Are you sure you want to continue connecting (yes/no/[fingerprint])?
...
```

如果用户选择 `yes` 就会将 server host key 记录到 client host key database 中(`~/.ssh/known_hosts`)。后续如果连接到同一台 server 就不需要询问用户是否选择连接

> [!note]
> 也被称为 TOFU trust on first use

但是如果 server host key 和 client host key database 中记录值不匹配就会拒绝连接

```
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@       WARNING: POSSIBLE DNS SPOOFING DETECTED!          @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
The ECDSA host key for github.com has changed,
and the key for the corresponding IP address 20.205.243.166
is unknown. This could either mean that
DNS SPOOFING is happening or the IP address for the host
and its host key have changed at the same time.
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@    WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!     @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
IT IS POSSIBLE THAT SOMEONE IS DOING SOMETHING NASTY!
Someone could be eavesdropping on you right now (man-in-the-middle attack)!
It is also possible that a host key has just been changed.
The fingerprint for the ECDSA key sent by the remote host is
SHA256:p2QAMXNIC1TJYWeIOttrVc98/R1BUFWu3/LiyKgUfQM.
Please contact your system administrator.
Add correct host key in /root/.ssh/known_hosts to get rid of this message.
Offending ECDSA key in /root/.ssh/known_hosts:5
ECDSA host key for github.com has changed and you have requested strict checking.
Host key verification failed.
```

如果确认 server host key 是正确的，这时需要手动删除 client host key database 中对应的条目，再重新连接

虽然这在一定程度上防止了 MITMA，但是谁来认证 server host key 是否真实合法呢？所以 OpenSSH 也引入了 PKI 的逻辑

> [!note]
> 使用 certificate 还可以完成 role-based access control
> 
> 类似于 `sshd_config` 中的 `Match`

和传统的 HTTPS 验证 SNI 不同，OpenSSH 除了可以验证 server 还可以验证 client

根据验证方向的不同，分为两种 certs

- host certificates

	向 client 验证 server 是否合法(传统 HTTPS 的逻辑)

- user certificates

	向 sever 验证 client 是否合法(逻辑上)

但是不管那种 certs 都需要通过 CA private key 签名，可以通过如下命令生成 CA root keypair

```
$ ssh-keygen -t rsa -b 4096 -f /etc/ssh/ssh_ca -C ssh_ca
```

这里使用 4096 bits rsa 来确保安全性

> [!note]
> 假设 CA/server/client 为 3 台不同的服务器
> CA 10.0.3.40
> server 10.0.3.41
> client 10.0.3.42

## 0x02 Host Certifcates Authentication

类似 HTTPS 中处理 TLS 的逻辑，向 clicent 验证 server 是否合法

首先 server 需要先生成自己的 keypair

> [!note]
> 通常在第一次运行 OpenSSH server 时，会自动生成，即 `sshd_config` `Hostkey` 默认值，如果没有可以通过 `ssh-keygen` 来生成 keypair

将 server public key 拷贝至 CA

```
$ scp /etc/ssh/id_ed25519.pub root@10.0.3.40:/tmp/
```

然后使用 CA root private key 对 server public key 签名(类似于 X509 中的对 CSR 签名)

```
$ ssh-keygen -I "10.0.3.41_id" -s /etc/ssh/ssh_ca -h -n "10.0.3.41" -V "always:forever" /tmp/ssh_host_ed25519_key.pub
Signed host key /tmp/ssh_host_ed25519_key-cert.pub: id "10.0.3.41_id" serial 0 for 10.0.3.41_DN valid forever
```

这里的 `/tmp/ssh_host_ed25519_key-cert.pub` 就是 10.0.3.41 的 host cert

我们可以通过如下命令来查看生成的 host cert 是否正确

```
$ ssh-keygen -Lf /tmp/ssh_host_ed25519_key-cert.pub
/tmp/ssh_host_ed25519_key-cert.pub:
        Type: ssh-ed25519-cert-v01@openssh.com host certificate
        Public key: ED25519-CERT SHA256:WfL7sVGttLvi8v330i5aBcnlqoSNM5oGPaHv2TLlziE
        Signing CA: RSA SHA256:RmpTHgtXBcXIBfGxyiEJHSD2tYiu39mM7C5sEnYOY0s (using rsa-sha2-512)
        Key ID: "10.0.3.41_id"
        Serial: 0
        Valid: forever
        Principals:
                10.0.3.41
        Critical Options: (none)
        Extensions: (none)
```

将 CA 签名的 host cert 拷贝至 server

```
$ scp /tmp/ssh_host_ed25519_key-cert.pub root@10.0.3.41:/etc/ssh/
```

server `sshd_config` 配置 CA 签名的 host cert

```
HostCertificate /etc/ssh/ssh_host_ed25519_key-cert.pub
```

重启 OpenSSH server

```
$ systemctl restart sshd
```

client 测试连接

```
$ ssh -v root@10.0.3.41
...
debug1: Server host certificate: ssh-ed25519-cert-v01@openssh.com SHA256:WfL7sVGttLvi8v330i5aBcnlqoSNM5oGPaHv2TLlziE, serial 0 ID "10.0.3.41_id" CA ssh-rsa SHA256:RmpTHgtXBcXIBfGxyiEJHSD2tYiu39mM7C5sEnYOY0s valid forever
debug1: No matching CA found. Retry with plain key
The authenticity of host '10.0.3.41 (10.0.3.41)' can't be established.
ED25519 key fingerprint is SHA256:WfL7sVGttLvi8v330i5aBcnlqoSNM5oGPaHv2TLlziE.
ED25519 key fingerprint is MD5:56:78:3a:12:6f:7f:a6:49:07:19:be:ee:70:ab:86:45.
Are you sure you want to continue connecting (yes/no)?
```

这时可以发现 server 会回送 host certifcate 但是因为本地没有 CA public key 所以，验证失败了。和 HTTPS 的逻辑一样，要将 CA public key 存储到本地

```
$ cat /etc/ssh/ssh_ca.pub | ssh root@10.0.3.42 'xargs echo @cert-authority \*  > ~/.ssh/known_hosts'
```

client 查看 CA public key

```
$ cat ~/.ssh/known_hosts
@cert-authority * ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8RRkh2jIXT0bepGA6/UPsaVtXfEHao6Tf4vnqWZ3LPFJVaTPik/7fzmTI5zhXOX2JZg4gpCW69E+O/Z2trI4/z1F67ZTSOqap4TUS1xvWLfV5tSHjF0yG6JB696PLONzNF98XpfVbEydNdqDHkV0QcLSHnETmdhuuklOa2l3oG5hLINCONtBGYxjYedN/witiGTYKE7z8Z9nns6ymMbNOLyojkzlKvHol+w4KSyL0CuyZWdcGbmUFb5MpvBJ5OE5OG2jfySo2/84ifDHULJ2qo+LNAKqqN16d2oMfdv91OBtdOFVYUIj7avd7QxbYqTylyiNwmGMM16bBrcSYawDi+ZCASpIxpbSscF3xCEgx1Fe+cUIVVVM4p2y+52znd9T7fNcxMtCopWGGZjctM+H2moIoKYvYkAKc0CMyu7TSN1U5TsCF4zUH835EsRhaIi9T/2NEWFWAns+LRFsFmetMyOTyUdlRG/BXemqznXGkCQmafD7sB67LSsboUILjDdCdaPKwKNFSNOf89ASIUvaKtA/AmoH9td9StNoGMVuwyubXIOJfBfwfU7VekFSLET4O0eQlMYhVz/AmTusUpmYzMsnIAlgY5rvnRDOPg397JUwCvPDV2SS64gWinPhwsX95+QonkH9zE410axTuqsAfKpzHjNSEm4NE3H9OAGb2PQ== ssh_ca
```

在执行完上述操作后 client 就可以向 server 发起连接。server 会回送自己的 host cert，然后 client 通过 CA public key 来验证 server 是否是有效的

> [!note]
> 这里并不会提示 `Are you sure you want to continue connecting (yes/no)?`
> 
> 因为由 PKI 已完成信任的操作了

```
$ ssh -v root@10.0.3.41
...
debug1: Server host certificate: ssh-ed25519-cert-v01@openssh.com SHA256:WfL7sVGttLvi8v330i5aBcnlqoSNM5oGPaHv2TLlziE, serial 0 ID "10.0.3.41_id" CA ssh-rsa SHA256:RmpTHgtXBcXIBfGxyiEJHSD2tYiu39mM7C5sEnYOY0s valid forever
debug1: Host '10.0.3.41' is known and matches the ED25519-CERT host certificate.
debug1: Found CA key in /root/.ssh/known_hosts:1
debug1: rekey after 134217728 blocks
debug1: SSH2_MSG_NEWKEYS sent
debug1: expecting SSH2_MSG_NEWKEYS
debug1: SSH2_MSG_NEWKEYS received
debug1: rekey after 134217728 blocks
debug1: SSH2_MSG_EXT_INFO received
debug1: kex_input_ext_info: server-sig-algs=<rsa-sha2-256,rsa-sha2-512>
debug1: SSH2_MSG_SERVICE_ACCEPT received
debug1: Authentications that can continue: password,keyboard-interactive
debug1: Next authentication method: keyboard-interactive
debug1: Authentications that can continue: password,keyboard-interactive
debug1: Next authentication method: password
root@10.0.3.41's password:
```

## 0x03 User Certificates Authentication

向 sever 验证 client 是否合法

> [!note]
> 也是 public key authentication 的一种变种，使用 user certificate 作为 public key

首先 client 需要先生成自己的 keypair

```
$ ssh-keygen -t rsa
Generating public/private rsa key pair.
Enter file in which to save the key (/root/.ssh/id_rsa):
/root/.ssh/id_rsa already exists.
Overwrite (y/n)? y
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in /root/.ssh/id_rsa.
Your public key has been saved in /root/.ssh/id_rsa.pub.
The key fingerprint is:
SHA256:5tgePvqIZK/htXv1gNvwVZHaU9gcboIVxbPWHlOUKy0 root@kingbase-3220
The key's randomart image is:
+---[RSA 2048]----+
|             o+O+|
|            o =+=|
|           . +.=*|
|            .E*B.|
|        S.   .+oo|
|       =o o .   .|
|    + o +* +     |
|   + = =+.o .    |
|    +.*=+.       |
+----[SHA256]-----+
```

将 client public key 拷贝至 CA

```
$ scp ~/.ssh/id_rsa.pub root@10.0.3.40:/tmp/
```

然后使用 CA root private key 对 client public key 签名(类似于 X509 中的对 CSR 签名)

```
$ ssh-keygen -I "root@10.0.3.42" -s /etc/ssh/ssh_ca -n root -V "always:+365d" /tmp/id_rsa.pub
Signed user key /tmp/id_rsa-cert.pub: id "root@10.0.3.42" serial 0 for root valid before 2025-12-03T09:55:12
```

> [!note]
> user certificate 必须使用 `-n` 指明 username 或者是 hostname，否则无法使用 user certificate 完成认证。`sshd` 提示 `error: Certificate lacks principal list`

这里的 `/tmp/id_rsa-cert.pub` 就是 10.0.3.42 的 user cert

我们可以通过如下命令来查看生成的 user cert 是否正确

```
$ ssh-keygen -Lf /tmp/id_rsa-cert.pub
/tmp/id_rsa-cert.pub:
        Type: ssh-rsa-cert-v01@openssh.com user certificate
        Public key: RSA-CERT SHA256:5tgePvqIZK/htXv1gNvwVZHaU9gcboIVxbPWHlOUKy0
        Signing CA: RSA SHA256:RmpTHgtXBcXIBfGxyiEJHSD2tYiu39mM7C5sEnYOY0s (using rsa-sha2-512)
        Key ID: "root@10.0.3.42"
        Serial: 0
        Valid: before 2025-12-03T09:55:12
        Principals:
                root
        Critical Options: (none)
        Extensions:
                permit-X11-forwarding
                permit-agent-forwarding
                permit-port-forwarding
                permit-pty
                permit-user-rc
```

将 CA 签名的 user cert 拷贝至 client

```
$ scp /tmp/id_rsa-cert.pub root@10.0.3.42:~/.ssh/
```

将 CA 的 public key 拷贝至 server

```
$ scp /etc/ssh/ssh_ca.pub root@10.0.3.41:/etc/ssh
```

server `sshd_config` 配置 CA 签名的 public key

```
TrustedUserCAKeys /etc/ssh/ssh_host_ed25519_key-cert.pub
```

要想使用 user certificate 还要开启 public key authentication

```
PubkeyAuthentication yes
```

重启 OpenSSH server

```
$ systemctl restart sshd
```

和 public key authentication 一样，要将 CA 签名的 user cert 拷贝至 server authorized_keys

```
$ cat ~/.ssh/id_rsa-cert.pub | ssh root@10.0.3.41 'xargs echo  >> ~/.ssh/authorized_keys'
```

监听 `ssh` 日志

```
$ journalctl -fu sshd
```

在执行完上述操作后 client 就可以向 server 发起连接。client 会发送自己的 user cert，然后 server 通过 CA public key 来验证 client 是否是有效的

```
$ ssh -v -i id_rsa root@10.0.3.41
...
debug1: Next authentication method: publickey
debug1: Offering RSA public key: id_rsa
debug1: Authentications that can continue: publickey,password,keyboard-interactive
debug1: Offering RSA-CERT public key: id_rsa
debug1: Server accepts key: pkalg ssh-rsa-cert-v01@openssh.com blen 1611
debug1: Authentication succeeded (publickey).
...
```

同时 `sshd` 日志会有如下信息

```
Dec 03 14:58:33 localhost.localdomain sshd[2668]: Accepted publickey for root from 10.0.3.220 port 51223 ssh2: RSA-CERT ID root@10.0.3.42 (serial 0) CA RSA SHA256:RmpTHgtXBcXIBfGxyiEJHSD2tYiu39mM7C5sEnYOY0s
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***

- [ssh\_config(5) - OpenBSD manual pages](https://man.openbsd.org/ssh_config)
- [How to configure SSH Certificate-Based Authentication](https://goteleport.com/blog/how-to-configure-ssh-certificate-based-authentication/)

***References***

[^1]:[ssh\_config(5) - OpenBSD manual pages](https://man.openbsd.org/ssh_config)
