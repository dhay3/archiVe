---
createTime: 2024-06-24 11:34
tags:
  - "#Github"
  - "#CICD"
---

# Github-Actions 07 - Jobs

## 0x01 Overview[^1]

jobs 是 workflows 中最核心的部分，决定了 workflow 要运行什么内容

```yaml
jobs:
  <job_id>:
    name: My first job
  <job_id>:
    name: My second job
```

其中 `jobs.<job_id> ` 必须全局唯一，`jobs.<job_id>.name` 是显示在 UI 上的

## 0x02 Execution order of jobs

jobs 默认以 parallel 的方式运行，如果想要 jobs 依次运行，需要使用如下格式定义
```yaml
jobs:
  job1:
  job2:
    needs: job1
  job3:
    needs: [job1, job2]
```

> [!NOTE]
> 有一点需要注意的时，`needs: ` 指定的 jobs，状态必须是 sccessful
> 上述例子中，只有 job1 运行成功才会运行 job2，只有 job1 和 job2 都运行成功才会运行 job3

如果只需要 jobs 运行完成，并不需要判断状态，可以使用如下方法
```
jobs:
  job1:
  job2:
    needs: job1
  job3:
    if: ${{ always() }}
    needs: [job1, job2]
```

## 0x03 Runner[^2]

每个 job 都需要定义 `runs-on`，表示 job 运行的 runner，这里只介绍 Github-hosted runners(白嫖的)[^3]

常用的 runners 如下

| **Virtual Machine** | **Processor (CPU)** | **Memory (RAM)** | **Storage (SSD)** | **Workflow label**                                                     | **Notes**                                                                                    |
| ------------------- | ------------------- | ---------------- | ----------------- | ---------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| Linux               | 4                   | 16 GB            | 14 GB             | `ubuntu-latest`, `ubuntu-24.04` [Beta], `ubuntu-22.04`, `ubuntu-20.04` | The `ubuntu-latest` label currently uses the Ubuntu 22.04 runner image.                      |
| Windows             | 4                   | 16 GB            | 14 GB             | `windows-latest`, `windows-2022`, `windows-2019`                       | The `windows-latest` label currently uses the Windows 2022 runner image.                     |
| macOS               | 3                   | 14 GB            | 14 GB             | `macos-12` or `macos-11`                                               | The `macos-11` label has been deprecated and will no longer be available after 28 June 2024. |
| macOS               | 4                   | 14 GB            | 14 GB             | `macos-13`                                                             | N/A                                                                                          |
| macOS               | 3 (M1)              | 7 GB             | 14 GB             | `macos-latest` or `macos-14`                                           | The `macos-latest` label currently uses the macOS 14 runner image.                           |

例如

```yaml
name: learn-github-actions
on: [push]
jobs:
  check-bats-version:
    runs-on: ubuntu-runners
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '14'
      - run: npm install -g bats
      - run: bats -v
```

### 0x03a Container[^4]

> 可以是 docker hub 中的任意 container

jobs 还可以运行在 container 中
```yaml
jobs:
  container-test-job:
    runs-on: ubuntu-latest
    container: node:18
```

## 0x04 Conditions

> [!warning] 
> 在 if 中通常无需使用 `${{}}`，但是如果取反，就必须使用 `${{}}`
> 例如 `if: ${{ ! startsWith(github.ref, 'refs/tags/') }}`

在 jobs 中还可以使用 flow control，来控制 jobs 只在特定情况下运行
例如 只在指定 repository 下运行
```yaml
name: example-workflow
on: [push]
jobs:
  production-deploy:
    if: github.repository == 'octo-org/octo-repo-prod'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '14'
      - run: npm install -g bats
```

## 0x05 Permissions

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Using jobs in a workflow - GitHub Docs](https://docs.github.com/en/actions/using-jobs/using-jobs-in-a-workflow)
[^2]:[Choosing the runner for a job - GitHub Docs](https://docs.github.com/en/actions/using-jobs/choosing-the-runner-for-a-job)
[^3]:[Choosing the runner for a job - GitHub Docs](https://docs.github.com/en/actions/using-jobs/choosing-the-runner-for-a-job#choosing-github-hosted-runners)
[^4]:[Running jobs in a container - GitHub Docs](https://docs.github.com/en/actions/using-jobs/running-jobs-in-a-container)