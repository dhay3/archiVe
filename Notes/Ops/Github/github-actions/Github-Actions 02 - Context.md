---
createTime: 2024-06-21 14:04
tags:
  - "#hash1"
  - "#hash2"
---

# Github-Actions 02 - Context

> Contexts are a way to access information about workflow runs, variables, runner environments, jobs, and steps. Each context is an object that contains properties, which can be strings or other objects.

把 context 理解成 workflows built-in class，通过他可以获取 variables，runner environments, jobs 以及 steps 中的 metadata

通过 `${{ <context> }}` 来调用，具体语法 Expression[^1]

具体可以使用的 context 参考 Contexts[^2]

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Expressions - GitHub Docs](https://docs.github.com/en/actions/learn-github-actions/expressions)
[^2]:[Contexts - GitHub Docs](https://docs.github.com/en/actions/learn-github-actions/contexts)


