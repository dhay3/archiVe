# git credentials

ref

https://git-scm.com/docs/git-credential

https://git-scm.com/docs/git-credential-store

https://git-scm.com/docs/git-credential-cache

## Digest

git credentials 是 git 的一个组件，用于解决重复输入用户账号和密码（credentials）的问题

有两种策略可以选择，分别是 cache 和 store

## Cache

在一段时间内存储 credentials，会失效

## Store

直接将 credentials 存储在本地，不会失效。一般存储在 `.git-credentials` 以如下格式存储

```
https://user:pass@example.com
```

如果是账号开启了 2FA 认证，pass 改用 OA 授权码