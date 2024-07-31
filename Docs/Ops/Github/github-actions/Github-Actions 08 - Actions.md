---
createTime: 2024-06-24 15:19
tags:
  - "#Github"
  - "#CICD"
---

# Github-Actions 08 - Actions

## 0x01 Overview


> [!warning] Title
> 需要注意的是 `uses:` 和 `run:` 不能在一个 steps 中一起使用


可以将 action 理解成函数，通过 `uses: project/repository@tags` 在不同的 workflows yaml 中被调用
```yaml
jobs:
  my_first_job:
    steps:
      - name: My first step
        uses: actions/setup-node@v4
```

## 0x02 Action Market[^1]

> 具体例子可以查看 [GitHub - dhay3/github-action-lab](https://github.com/dhay3/github-action-lab)

Github 提供了 [Action Market](https://github.com/marketplace?type=actions)，用户可以在 market 找到各种各样的 actions

### 0x02a Useful Actions

#### Essential Tools
- [Checkout · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/checkout)
	切到指定的 branch，从而让 workflow 可以获取到 branch 中的内容(可以执行 branch 中的脚本)
	你无需拥有 repository 的所有权也可以使用 checkout action 获取 repository 中的内容，可以参考
	https://github.com/MetaCubeX/meta-rules-dat/blob/master/.github/workflows/run.yml
- [Upload a Build Artifact · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/upload-a-build-artifact)
	可以将 workflow 生成的 artifact 上传到 github(会提供下载链接)
- [Download a Build Artifact · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/download-a-build-artifact)
	可以将 workflow 生成的 artifact 下载下来
- [Create Pull Request · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/create-pull-request)
	生成 PRs
- [Setup Python · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/setup-python)
	设置 Python 环境
- [Setup Java JDK · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/setup-java-jdk)
	设置 JDK 环境
- [GH Release · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/gh-release)
	生成 github releases

#### Profile/UI
- [Metrics embed · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/metrics-embed)
	可以写出漂亮的 readme
- [generate-snake-game-from-github-contribution-grid · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/generate-snake-game-from-github-contribution-grid)
	生成 git commit pacman game

#### Service
- [Hugo setup · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/hugo-setup)
	Hugo
- [Build and push Docker images · Actions · GitHub Marketplace · GitHub](https://github.com/marketplace/actions/build-and-push-docker-images)
	构建/推送 docker image 

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Finding and customizing actions - GitHub Docs](https://docs.github.com/en/actions/learn-github-actions/finding-and-customizing-actions)
