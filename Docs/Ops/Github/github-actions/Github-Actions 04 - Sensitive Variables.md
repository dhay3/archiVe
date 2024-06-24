---
createTime: 2024-06-21 14:22
tags:
  - "#Github"
  - "#CICD"
---

# Github-Actions 04 - Sensitive Variables

## 0x01 Overview

敏感的信息通过 [Github-Actions 03 - Variables](Github-Actions%2003%20-%20Variables.md) 中的方式存储，非常不合理也不安全

github 提供了 2 种安全的方法
1. repository secrects[^1]
2. environment secrects[^2]

这里只介绍 repository secrects

## 0x02 Repository secrects

> 具体配置方法参考 creating-secrets-for-an-environment[^1]

repository secrects 对当前 repository 下的所有 workflows 都生效，需要通过 `secrects` context 来获取

例如 配置了一个 SuperSecret repository secrects，那么就可以通过如下方式来调用 SuperSecret(如果没有配置默认取空值)

```yaml
steps:
  - name: Hello world action
    with: # Set the secret as an input
      super_secret: ${{ secrets.SuperSecret }}
    env: # Or as an environment variable
      super_secret: ${{ secrets.SuperSecret }}
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Using secrets in GitHub Actions - GitHub Docs](https://docs.github.com/en/actions/security-guides/using-secrets-in-github-actions?tool=webui#creating-secrets-for-a-repository)
[^2]:[Using secrets in GitHub Actions - GitHub Docs](https://docs.github.com/en/actions/security-guides/using-secrets-in-github-actions?tool=webui#creating-secrets-for-an-environment)