---
createTime: 2024-06-21 17:02
tags:
  - "#hash1"
  - "#hash2"
---

# Github-Actions 05 - Variable GITHUB_TOKEN

## 0x01 Overview[^1]

GITHUB_TOKEN 是 github actions 的一个特殊变量，当运行 workflow 时会自动生成，在 workflow 结束后销毁。包含当前 repository 所有 scopes 权限的 token，可以被用作参数传入到一些需要 permissions 的 actions 中

可以通过 `${{secrets.GITHUB_TOKEN}}` 来调用，也可以通过 `${{github.token}}` context 来调用

## 0x02 Permissions for GITHUB_TOKEN

GITHUB_TOKEN，有 3 种预设的权限，默认会使用 permissive
1. permissive 权限最大
2. restricted 权限最小
3. PRs from public fork 只针对 fork repo PRs

| Scope               | Default access  <br>(permissive) | Default access  <br>(restricted) | Maximum access for  <br>pull requests from  <br>public forked repositories |
| ------------------- | -------------------------------- | -------------------------------- | -------------------------------------------------------------------------- |
| actions             | read/write                       | none                             | read                                                                       |
| attestations        | read/write                       | none                             | read                                                                       |
| checks              | read/write                       | none                             | read                                                                       |
| contents            | read/write                       | read                             | read                                                                       |
| deployments         | read/write                       | none                             | read                                                                       |
| id-token            | none                             | none                             | read                                                                       |
| issues              | read/write                       | none                             | read                                                                       |
| metadata            | read                             | read                             | read                                                                       |
| packages            | read/write                       | read                             | read                                                                       |
| pages               | read/write                       | none                             | read                                                                       |
| pull-requests       | read/write                       | none                             | read                                                                       |
| repository-projects | read/write                       | none                             | read                                                                       |
| security-events     | read/write                       | none                             | read                                                                       |
| statuses            | read/write                       | none                             | read                                                                       |

GITHUB_TOKEN 可以按照 enterprise, organization, repository 层级，来设置默认使用的预设权限
1. [Enforcing policies for GitHub Actions in your enterprise - GitHub Enterprise Cloud Docs](https://docs.github.com/en/enterprise-cloud@latest/admin/policies/enforcing-policies-for-your-enterprise/enforcing-policies-for-github-actions-in-your-enterprise#enforcing-a-policy-for-workflow-permissions-in-your-enterprise)
2. [Disabling or limiting GitHub Actions for your organization - GitHub Docs](https://docs.github.com/en/organizations/managing-organization-settings/disabling-or-limiting-github-actions-for-your-organization#setting-the-permissions-of-the-github_token-for-your-organization)
3. [Managing GitHub Actions settings for a repository - GitHub Docs](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-github-actions-settings-for-a-repository#setting-the-permissions-of-the-github_token-for-your-reposi)

> [!NOTE] 
> read and write permissions 为 permissive
  read repository contents and package permissions 为 restricted
   
除了使用预设的权限(permissive and restricted)，也可以通过 workflow file `permissions: `[^2] 来修改
对应 scope permission 的值可以是 read|write|none
```yaml
permissions:
  actions: read|write|none
  checks: read|write|none
  contents: read|write|none
  deployments: read|write|none
  id-token: read|write|none
  issues: read|write|none
  discussions: read|write|none
  packages: read|write|none
  pages: read|write|none
  pull-requests: read|write|none
  repository-projects: read|write|none
  security-events: read|write|none
  statuses: read|write|none
```
或者可以使用 read-all|write-all 表示所有 scopes 设置 read 或者 write permission

根据 `permissions:` 定义的位置不同，作用的范围也不同
`jobs.<job_id>.permissions` 只针对当前 job
```yaml
jobs:
  stale:
    runs-on: ubuntu-latest

    permissions:
      issues: write
      pull-requests: write

    steps:
      - uses: actions/stale@v5
```

`permissions` 针对当前 workflow 中的所有 jobs
```yaml
name: "My workflow"

on: [ push ]

permissions: read-all

jobs:
  ...
```

当使用了 `permissions: ` 时，如果对应 scope 没有明确指定权限，默认设置为 no access
例如 设置了 actions/checks 为 read，contents 为 write，那么其他没有在 workflow file 中明确声明权限的 scope 直接会置为 no access
```
permissions:
  actions: read
  checks: read
  contents: write
```

> [!warning] 
> 任何人在不同层级具有 write 权限的，都可以通过 workflow file `permission: ` 修改 GITHUB_TOKEN 可以调用的权限


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Automatic token authentication - GitHub Docs](https://docs.github.com/en/actions/security-guides/automatic-token-authentication)
[^2]:[Workflow syntax for GitHub Actions - GitHub Docs](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#permissions)