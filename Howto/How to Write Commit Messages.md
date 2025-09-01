---
createTime: 2024-06-07 10:26
---

# How to Write Good Commit Messages

## 0x01 Preface

Commit Messages 是你在每次 git commit 时必须要填的。你可能会不知道填什么，或者随便填。例如

```
* c1b5f09 - (3 years, 4 months ago) update content - cyberPelican
* 5d09c60 - (3 years, 3 months ago) rm license - cyberPelican
* 340f176 - (3 years, 4 months ago) modified - cyberPelican
```

这些 commit messages 即不美观，也不能说明具体 commit 了什么

假设你现在想要看某一文件的变更记录，有如下两种方式
1. 在 Github 上通过文件右上角的 History 来查看
2. 在 Git 上可以通过 `git diff <commit1> <commit2>` 来查看

但是不管是 Github 还是 Git 只会显示 commit 对应的 hash 和 commit message。如果你以上述 commit message 的方式记录 commit，就很难将 commit 修改的内容和 commit 的对上。所以需要一套标准

按照个人喜好提供 2 种，更推荐 Emojial Commits(图形能表达更多意思，且更简洁)
1. Conventional Commits
2. Emojial Commits

## 0x02 Conventional Commits

这里在 Conventional Commits[^1] 的基础上总结如下规则

所有 commit messages 按照如下 EBNF

```
<type>[!]([optional scope]): <description>
```

### 0x02a Type

按照 commit 意图分为如下几类(具体可以参考 [Angular convention](https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#-commit-message-guidelines))，type 必须全小写

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
- build
	Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
	```
	build: update dockerfile
	```
- ci
	Changes to CI configuration files and scripts
	```
	ci: add release ci
	```
- test
	Adding missing tests or correcting existing tests
	```
	test: add unitest module 
	```

另外可以在 type 前使用 exclamation 表示在对应类别有重要的修改，例如

```
chore!: drop Node 6 from testing matrix
docs!: update 'how to use' in readme
```

### 0x02b Optional Scope

commits 只关联特定某领域，例如

```
docs(ai): add 'installation' for stablediffusion-webui
docs(ops)!: add 'how to install docker'
docs(lang): add polish language
```

### 0x02c Description

commit 主要做了什么，需要符合如下规则

1. Limit the subject line to 50 characters
2. Don't put a period(`.`) at the end of the subject line

```
chore: update npm dependency to latest version
```

## 0x03 Emojial Commits

> [!tip] 
> 如果想要在 terminal 中显示 emoji，可以参考 Noto Emoji Color fontconfig for Konsole[^3]

借助 gitmoji-cli[^2]

```
<intention> [scope:?] <message>
```

其中 commit title 对应 git commit 中的 message 也就是在 Github Code 页面显示的

```
gitmoji -c
? Choose a gitmoji: 📝  - Add or update documentation.
? Enter the commit title [52/48]: Add "Emojial Commits" in "How to Write Good Commits"
? Enter the commit message:
[master 303b0a1] 📝 Add "Emojial Commits" in "How to Write Good Commits"
 1 file changed, 55 insertions(+), 21 deletions(-)
```

### 0x03a Intention

> 所有的 emoji 都可以在 [gitmoji | An emoji guide for your commit messages](https://gitmoji.dev/) 可以找到
> 
> 或者使用 `plasma-emojier` 搜索

表示意图的 emoji

常用的有
- 🐛 Fix a bug 
- 🔥 Remove code or files
- ♻️ Refactor code
- 📝 Add or update documents
- ✨ Introduce new feature
- 🚧 Work in progress
- ✏️ Fix typos
- 📦️ Add or update compiled files or packages
- 📄 Add or update license
- 💥 Introduce breaking changes
- 👽️ Update code due to external API changes
- ⬆️ Upgrade dependencies
- ⬇️ Downgrade dependencies
- ⏪️ Revert changes.

可以使用 `gitmoji -l` 查看所有的 emoji

### 0x03b Scope

同 [Conventional Commits#Optional Scope](#Optional%20Scope) 相同

### 0x03c Message

同 [Description](#Description)

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See Also***

- [How to Write Better Git Commit Messages – A Step-By-Step Guide](https://www.freecodecamp.org/news/how-to-write-better-git-commit-messages/)
- [angular/CONTRIBUTING.md](https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#-commit-message-guidelines)

***References***

[^1]:[Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0-beta.4/)
[^2]:[GitHub - carloscuesta/gitmoji: An emoji guide for your commit messages. 😜](https://github.com/carloscuesta/gitmoji?tab=readme-ov-file)
[^3]:[Noto Emoji Color fontconfig for Konsole · GitHub](https://gist.github.com/IgnoredAmbience/7c99b6cf9a8b73c9312a71d1209d9bbb)


