	---
createTime: 2024-06-20 16:14
tags:
  - "#Github"
  - "#CICD"
---

# Github-Actions 01 - Helloworld

> 先看使用限制[^4]

## 0x01 Overview

> GitHub Actions is a continuous integration and continuous delivery (CI/CD) platform that allows you to automate your build, test, and deployment pipeline.

Github actions 是 github 提供的 CICD 平台，可以实现自动构建/测试/部署

## 0x02 Terms[^1]

### 0x02a Workflows

> A workflow is a configurable automated process that will run one or more jobs.

workflows 是自动化的流程，会在特定的 events 触发
workflows 需要以 yaml 的格式编写，并存储在 `.github/workflows/` 下，可以同时包含多个 workflows(one yml per workflow)，例如

> you can have one workflow to build and test pull requests, another workflow to deploy your application every time a release is created, and still another workflow that adds a label every time someone opens a new issue.

> [!important] 
> Workflows 会同时以 parallel 的方式运行
> 如果想要按照顺序运行，需要使用 `jobs.<job_id>.needs` 或者直接使用 `steps`
> 所以尽量避免过分解耦生成多个 workflow files，同一大类功能的都应该放在一个 workflow file 中

### 0x02b Events

> An event is a specific activity in a repository that triggers a workflow run.

events 就是 workflows' triggers，例如

>  an activity can originate from GitHub when someone creates a pull request, opens an issue, or pushes a commit to a repository.

### 0x02c Jobs

> A job is a set of steps in a workflow that is executed on the same runner.

一组 steps 或者是 shell scripts 的集合，有如下规则
1. steps 按照从上到下的顺序执行
2. steps 之间数据共享(因为在同一个 runner 上运行)

> [!important]
> jobs 默认以 parallel 的方式运行，如果想要按照顺序执行需要使用 `jobs.<job_id>.needs`

### 0x02d Actions

> An *action* is a custom application for the GitHub Actions platform that performs a complex but frequently repeated task.

可以将 action 理解成函数，可以被不同的 workflows ymal 中被调用

action market[^3]

### 0x02e Runner

> A runner is a server that runs your workflows when they're triggered.

runner 就是一台会运行你配置的 workflows 的服务器，需要注意如下几点
1. Larger runners are only available for organizations and enterprises using the GitHub Team or GitHub Enterprise Cloud plans.[^2]


## 0x03 Demo

1. 在 git repository 下创建 `.github/workflows/`，用于存储 workflows yaml
2. 创建 `.github/workflows/helloworld.yml`
	```
	name: helloworld
	run-name: ${{ github.actor }} is learning GitHub Actions
	on: [push]
	jobs:
	  check-bats-version:
	    runs-on: ubuntu-latest
	    steps:
	      - uses: actions/checkout@v4
	      - uses: actions/setup-node@v4
	        with:
	          node-version: '20'
	      - run: npm install -g bats
	      - run: bats -v
	
	```
3. 推送到 remote repository 即可，github 就会自动运行 workflow，运行的任务在 Actions tab 可以查看


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Understanding GitHub Actions - GitHub Docs](https://docs.github.com/en/actions/learn-github-actions/understanding-github-actions)
[^2]:[About larger runners - GitHub Docs](https://docs.github.com/en/actions/using-github-hosted-runners/about-larger-runners/about-larger-runners)
[^3]:[Finding and customizing actions - GitHub Docs](https://docs.github.com/en/actions/learn-github-actions/finding-and-customizing-actions)
[^4]:[Usage limits, billing, and administration - GitHub Docs](https://docs.github.com/en/actions/learn-github-actions/usage-limits-billing-and-administration)
