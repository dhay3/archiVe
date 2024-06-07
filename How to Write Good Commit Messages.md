---
author: "0x00"
createTime: 2024-06-07 10:26
draft: false
---

# How to Write Good Commit Messages

Commit Message 是你在每次 commit 时必须要填的。你可能会不知道填什么，或者随便填。例如

```
* c1b5f09 - (3 years, 4 months ago) update content - cyberPelican
* 5d09c60 - (3 years, 3 months ago) rm license - cyberPelican
* 340f176 - (3 years, 4 months ago) modified - cyberPelican
```

这些 commit messages 即不美观也不能说明具体 commit 了什么

## Conventional Commits

如果你想要看文件的变更记录，在 Github 上可以通过 History 来查看，在 Git 上可以通过 `git diff <commit1> <commit2>` 来查看。但是不管是 Github 还是 Git 只会显示 commit 对应的 hash 和 commit message。如果你以上述 commit message 的方式记录 commit，就很难将 commit 修改的内容和 commit 的对上。所以需要一套标准，这里在 Conventional Commits[^1] 的基础上总结如下规则

所有 commit messages 按照如下 EBNF 

```
<type>([optional scope]): <description>
```

### Type

按照 commit 意图分为如下几类，type 必须全小写

- init 
	a project has been initialed
	```
	init: [repository] created
	```
- feat 
	a new feature is introduced with the changes
	```
	feat: allow provided config object to extend other configs
	```
- fix
	a bug fix has occurred
	```
	fix: tables don't rendered in reading view 
	```
- chore
	changes that do not relate to a fix or feature and don't modify src or test files (for example updating dependencies)
	```
	chore: update dependency spring-boot 4.1 to sprint-boot 4.2
	```
- docs
	updates to documentation such as a the README or other markdown files
	```
	docs: add badges to readme
	```
- style
	changes that do not affect the meaning of the code, likely related to code formatting such as white-space, missing semi-colons, and so on.
	```
	style: remove unintended white-space
	```
- revert
	reverts a previous commit
	```
	revert: revert commit [hash1] to [hash2]
	```

另外可以在 type 前使用 exclamation 表示在对应类别有重要的修改，例如

```
chore!: drop Node 6 from testing matrix
docs!: update 'how to use' in readme
```

### Optional Scope

commits 只关联特定某领域，例如

```
docs(ai): add 'installation' for stablediffusion-webui
docs(ops)!: add 'how to install docker'
docs(lang): add polish language
```

### Description

commit 主要做了什么，需要符合如下规则

1. Limit the subject line to 50 characters
2. Don't put a period(`.`) at the end of the subject line

```
chore: update npm dependency to latest version
```

除此外还可以在 description 中加入 emoji， 常用的有
- :wrench: 🔧 用于 fix
- :memeo: 📝 用于 docs
- 

docs(ai):  add 'installation' for stablediffusion-webui

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

**references**

[^1]:[Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0-beta.4/)
[^2]:[How to Write Better Git Commit Messages – A Step-By-Step Guide](https://www.freecodecamp.org/news/how-to-write-better-git-commit-messages/)
[^3]:[Understand how to write a good commit message through memes 😉 | by Hritik Jaiswal | Medium](https://medium.com/@hritik.jaiswal/how-to-write-a-good-commit-message-9d2d533b9052)


