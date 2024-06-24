---
createTime: 2024-06-21 16:55
tags:
  - "#Github"
  - "#CICD"
---

# Github-Actions 06 - Events

## 0x01 Overview[^1]

events(triggers) 就是 workflows 的触发器(也可以说是条件)，通过 `on: ` 指定

## 0x02 Events

> [!NOTE] 
> 这里只记录常用的 events，具体可以使用的 events 参考 events-that-trigger-workflows[^2]

### 0x02a workflow_dispatch

> 只有 workflow files 在 default branch 时才生效

允许 workflow 手动触发，同时还提供了自定义输入参数的功能

例如
```yaml
on:
  workflow_dispatch:
    inputs:
      logLevel:
        description: 'Log level'
        required: true
        default: 'warning'
        type: choice
        options:
        - info
        - warning
        - debug
      tags:
        description: 'Test scenario tags'
        required: false
        type: boolean
      environment:
        description: 'Environment to run tests against'
        type: environment
        required: true

jobs:
  log-the-inputs:
    runs-on: ubuntu-latest
    steps:
      - run: |
          echo "Log level: $LEVEL"
          echo "Tags: $TAGS"
          echo "Environment: $ENVIRONMENT"
        env:
          LEVEL: ${{ inputs.logLevel }}
          TAGS: ${{ inputs.tags }}
          ENVIRONMENT: ${{ inputs.environment }}
```

上述 workflow 定义了 3 个参数，可以分别通过 `${{inputs.logLevel}}`，`${{inputs.tags}}`，`${{inputs.environments}}` 获取到对应的值

如果想要运行 workflows 就需要通过 Github browser 手动触发[^3]
![](https://docs.github.com/assets/cb-78157/mw-1440/images/help/actions/workflow-dispatch-inputs.webp)
或者是通过 `gh` Github CLI 来触发
```shell
gh workflow run run-tests.yml -f logLevel=warning -f tags=false -f environment=staging
```

### 0x02b schedule

> 只有 workflow files 在 default branch 时才生效

在特定 cron 触发 workflow
```yaml
on:
  schedule:
    - cron: '30 5 * * 1,3'
    - cron: '30 5 * * 2,4'

jobs:
  test_schedule:
    runs-on: ubuntu-latest
    steps:
      - name: Not on Monday or Wednesday
        if: github.event.schedule != '30 5 * * 1,3'
        run: echo "This step will be skipped on Monday and Wednesday"
      - name: Every time
        run: echo "This step will always run"
```
可以通过 [Crontab.guru - The cron schedule expression generator](https://crontab.guru/) 来生成或者校验 cron 表达式

### 0x02c push

workflow 在 push commits 时触发
push 还支持在特定 branches push 时触发
```yaml
on:
  push:
    branches:
      - 'main'
      - 'releases/**'
```
或者是特定 branches 取反
```yaml
on:
  push:
    branches-ignore:
      - 'main'
```
也可以是在特定文件 push 时触发
```yaml
on:
  push:
    paths:
      - '**.js'
```
或者是特定文件取反
```yaml
on:
  push:
    paths-ignore:
      - 'docs/**'
```
branches 和 path 可以一起使用，必须满足两者才会触发
```yaml
on:
  push:
    branches:
      - 'releases/**'
    paths:
      - '**.js'
```

### 0x02d pull_request

workflow 在 pull request 时触发

pull_reuqest 还支持在特定 branches pull reuqest 时触发
```yaml
on:
  pull_request:
    types:
      - opened
    branches:
      - 'releases/**'
```
也可以是在特定文件 pull request 时触发
```yaml
on:
  pull_request:
    paths:
      - '**.js'
```
pull_request 和 path 可以一起使用，必须满足两者才会触发
```yaml
on:
  pull_request:
    types:
      - opened
    branches:
      - 'releases/**'
    paths:
      - '**.js'
```

### 0x02e issues

> 只有 workflow files 在 default branch 时才生效

workflow 在 issues 时触发
```yaml
on:
  issues:
    types: [opened, edited, milestoned]
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Triggering a workflow - GitHub Docs](https://docs.github.com/en/actions/using-workflows/triggering-a-workflow)
[^2]:[Events that trigger workflows - GitHub Docs](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows)
[^3]:[Manually running a workflow - GitHub Docs](https://docs.github.com/en/actions/using-workflows/manually-running-a-workflow)
