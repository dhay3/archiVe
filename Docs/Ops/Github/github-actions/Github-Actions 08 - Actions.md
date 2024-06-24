---
createTime: 2024-06-24 15:19
tags:
  - "#Github"
  - "#CICD"
---

# Github-Actions 08 - Actions

## 0x01 Overview

可以将 action 理解成函数，可以在不同的 workflows yaml 中被调用

## 0x02 Action Market[^1]

Github 提供了 [Action Market](https://github.com/marketplace?type=actions)，用户可以在 market 找到各种各样的 actions

actions 可以通过 `project/repository@tags` 来调用
例如
```yaml
jobs:
  my_first_job:
    steps:
      - name: My first step
        uses: actions/setup-node@v4
```

### 0x02a Useful Actions

- [Metrics embed · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/metrics-embed)
	可以写出漂亮的 readme
- [Checkout · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/checkout)
	可以切换到指定 branch
- [GitHub Pages action · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/github-pages-action)
	github page 部署
- [Upload a Build Artifact · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/upload-a-build-artifact)

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Finding and customizing actions - GitHub Docs](https://docs.github.com/en/actions/learn-github-actions/finding-and-customizing-actions)