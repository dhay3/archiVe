---
createTime: 2024-06-24 14:43
tags:
  - "#Github"
  - "#CICD"
---

# Github-Actions 09 - Workflow file

具体关键字使用，参考 Workflow syntax[^1]

```yaml
name: Compress images
on:
  pull_request:
    paths:
      - '**.jpg'
      - '**.jpeg'
      - '**.png'
      - '**.webp'
  push:
    branches:
      - master
    paths:
      - '**.jpg'
      - '**.jpeg'
      - '**.png'
      - '**.webp'
  workflow_dispatch:
  schedule:
    - cron: '00 00 * * 1'
jobs:
  build:
    name: calibreapp/image-actions
    runs-on: ubuntu-latest
    if: |
      github.event_name != 'pull_request' || 
      github.event.pull_request.head.repo.full_name == github.repository
    steps:
      - name: Checkout Branch
        uses: actions/checkout@v4
      - name: Compress Images
        id: calibre
        uses: calibreapp/image-actions@main
        with:
          githubToken: ${{ secrets.GITHUB_TOKEN }}
          compressOnly: ${{ github.event_name != 'pull_request' }}
      - name: Create Pull Request
        if: |
          github.event_name != 'pull_request' &&
          steps.calibre.outputs.markdown != ''
        uses: peter-evans/create-pull-request@v6
        with:
          title: Auto Compress Images
          branch-suffix: timestamp
          commit-message: Compress Images
          body: ${{ steps.calibre.outputs.markdown }}
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Workflow syntax for GitHub Actions - GitHub Docs](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions)