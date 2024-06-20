---
createTime: 2024-06-20 16:14
tags:
  - "#Github"
---

# Github-Actions 01 - Overview

## 0x01 Overview

> GitHub Actions is a continuous integration and continuous delivery (CI/CD) platform that allows you to automate your build, test, and deployment pipeline.

Github actions 是 github 提供的 CICD 平台，可以实现自动构建/测试/部署

## 0x02 Terms[^1]

### 0x02a Workflows

workflows 是自动化的流程，会在特定的 events 触发
workflows 需要以 yaml 的格式编写，并存储在 `.github/workflows/` 下，可以同时包含多个 workflows，例如

> you can have one workflow to build and test pull requests, another workflow to deploy your application every time a release is created, and still another workflow that adds a label every time someone opens a new issue.

### 0x02b Events

Events 就是 workflows' triggers，例如

>  an activity can originate from GitHub when someone creates a pull request, opens an issue, or pushes a commit to a repository.

### 0x03c Jobs



---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Understanding GitHub Actions - GitHub Docs](https://docs.github.com/en/actions/learn-github-actions/understanding-github-actions)



