---
author: "0x00"
createTime: 2024-08-14 09:09
tags:
  - "#hash1"
  - "#hash2"
---

# How to Ask Questions

## 0x01 Preface

同样的一个问题，表达的方式不同，得到的回复也不同(当然你也可能得不到回复 XD)。如何问好一个问题，才可以得到你满意的回答，这非常重要

本文在 [How To Ask Questions The Smart Way](http://www.catb.org/~esr/faqs/smart-questions.html) 的基础上按照时间的先后顺序划分如下几部分

1. Before You Ask
2. When You Ask
3. After the Problem has been Solved

## 0x02 Before You Ask

> [!important] 
> 各部分的结果可能存在交集，所以逻辑上没有先后顺序，但是从结果的权威性角度出发
> $DTFE > RTFS > JTFG > RTFM > STFW > ATFL$

在提问前需要做到如下几点(至少要证明自己对问题做过解决的尝试，不要让别人觉得你只是一个 taker)

- STFW
	try to find an answer by searching the web
- RTFW
	try to find an answer by reading the manual
- ATFL
	try to find an answer by asking the LLM
- DTFW
	try to find an answer by doing the experimentation
- RTFS
	try to find an answer by reading the source code
- JTFG
	try to find an answer by joining groups

### 0x02a Try to Find an Answer by Searching the Web(STFW)

利用 SEO 搜索你想要的答案

不仅限于各种 SEO(掌握 SEO advanced search 以及 operator search 的使用是必要的，例如 [Google Searches](Docs/Sundries/Chrome/Google%20Searches.md))，也包括一些设置了 robots.txt[^1](按照 RFC 9309 的规定，SEO 需要按照 robots.txt 中的规则爬取信息，但是不是强制要求) 的 forum 或者 site(例如 Google scholar/libgen/etc.)。这个利用 SEO 搜索的过程可以简称为 STFW(Searching the Fucking Web)

例如 你想要搜索和 KDE system monitor sensors bug 相关的内容，你可以使用 Google/DuckDuckgo/Brave/etc. SEO 直接搜索关键字 `KDE system monitor sensors bug` 也可以使用 operator search 搜索 `intitle:"system monitor sensors" intext:bug` 对结果过滤

如果没有得到想要的内容，可以到 [KDiscuss](https://discuss.kde.org/) forum 搜索，也可以到 [KDE Bugtracking System](https://bugs.kde.org/) 搜索(相对更精确，前提是你知道有对应的站点)

### 0x02b Try to Find an Answer by Reading the Manual(RTFM)

使用 User Manual 搜索你想要的答案

和家电一样，大多数软件都会有 User Manual(也可能叫 Get Help/Documentation/Support/etc.)，会告诉用户该如何使用软件以及 troubleshooting。这个使用 User Manual 搜索的过程可以简称为 RTFM(Reading the Fucking Manual)

例如 你想要使用 Thunderbird 接管 Gmail，就可以在 [Thunderbird Support](https://support.mozilla.org/en-US/products/thunderbird) 页面搜索关键字 gmail

### 0x02c Try to Find an Answer by Asking the LLM(ATFL)

使用 LLM 得到你想要的答案

> [!important]
> Prompts 使用的语言也决定了 LLM reponse 的质量

LLM 是一个很好的辅助工具(回答并不一定可靠，所以作为一个辅助工具)。你可以设置场景，让 AI 分步骤回答你的问题。这个使用 LLM 搜索的过程可以简称为 ATFL(Asking the Fucking LLM)

推荐几个 LLM

- [ChatGPT](https://chatgpt.com/)
- [Claude](https://claude.ai/)
- [Gemini](https://gemini.google.com/)
- [Perplexity](https://www.perplexity.ai/) 强烈推荐

另外推荐一个在线 AI 整合网站 [Toolify](https://www.toolify.ai/)

当然如果你想要确保隐私可以使用 [Ollama](https://ollama.com/) 本地部署，或者自己使用 [huggingface](https://huggingface.co/) 中的资源自己训练

例如 你想知道 GPLv1/2/3 之间的区别，你可以使用如下 prompt，对 LLM 设置角色

```
You are an expert in open source license and copyright sphere.
```

让 LLM 列出 GPLv1/2/3 的简要说明

```
Give me the brief of GPLv1, GPLv2 and GPLv3 in layman's word.
```

最后让 LLM 列出区别

```
Differences between GPLv1, GPLv2 and GPLv3
```

### 0x02d Try to Find an Answer by Doing Experimentation(DTFE)

通过 Experimentation 得到想要的答案

做实验是最原始也是最有效的方法，但是要控制变量。如果有正常的对照组，可以设置参照因子(这个就是编程中最朴素的 if ... elif ... else ... 逻辑)。这个做 Experimentation 的过程可以简称为 DTFE(Doing the Fucking Experimentation)

例如 你现在不能通过浏览器访问你家的软路由了(不考虑 DNS)，你要先确认是 L7 应用有问题，或者是 L4 传输有问题，或者是 L3 网络有问题 还是 L2/L1 设备有问题。这里 Layer 就是变量

- 如果 设备电口正常亮灯，可以概率排除 L2/L1 设备问题，反之 就排查物理设备或者配件(例如 软路由网口/网线/交换机/光猫/etc.)
- 如果 可以从 LAN 内主机 ping 通软路由，可以概率排除 L3 网络问题，反之 就排查软路由系统 IP 地址以及路由
- 如果 可以从 LAN 内主机 netcat 软路由 80 端口，可以概率排除 L4 传输问题，反之 就排查软路由系统是否有应用监听 socket
- 那么 最后只可能是 L7 应用问题，重启 luci 即可

### 0x02e Try to Find an Answer by Reading the Source Code(RTFS)

通过阅读 Source Code 得到想要的答案

大多数符合 OSD 标准[^2]的软件都会公开代码，所以一些奇怪的 Bug 或者是非预期的现象可以通过读源码解决。通过阅读 Source code 的过程可以简称为 RTFS(Reading the Fucking Source Code)

例如 Virtualbox Guest Machine 不能通过 Vagrant 启动，但是可以通过 Virtualbox GUI 启动

[\[SOLVED\] Can't boot virtual machines via vagrant but works well with virtualbox GUI(vagrant 2.4.0/virtualbox 7.0.12) · Issue #13288 · hashicorp/vagrant · GitHub](https://github.com/hashicorp/vagrant/issues/13288)

但是阅读 Source Code 并不是一件简单的事。如果在阅读 Source Code 后你仍然搞不明白，请务必表示

```
It's hard for me to read the source code
```

### 0x02f Try to Find an Answer by Joining Groups(JTFG)

加入 Groups 得到想要的答案

软件的站点不一定会及时更新，但是通常会有一些 official 社区组群(Discord/Telegram/etc.)，管理者会在这些组群中实时发布一些 announcements 或者是 Pinned Messages 会提供非常有的信息

例如 在 Clash Verge Rev 更新 1.6.3 后，Global Extend Config 中配置如下额外全局路由规则，发现不生效(之前生效)

```yaml
prepend-rules:
	- DOMAIN-SUFFIX,4lice.com,DIRECT
	- AND,((IP-CIDR,23.94.117.188/32),(DST-PORT,65522)),DIRECT
```

这时可以加入站点提供的 telegram 组群，Pinned Messages 说明了 v1.6.2 后 Merge 配置(Global Extend Config)不再支持 `prepend-rule-providers`、`prepend-proxy-providers`、`append-rule-providers`、`append-proxy-providers`，要通过 edit profile rule 或者是 proxies/proxy groups 实现更加细粒的配置

## 0x03 Some Tips for Before You Ask

下面还有一些小技巧适用于 Before You Ask

### 0x03a Search Engine Shortcuts

> [!NOTE]
> 读者可以自行使用 [0x02 Before You Ask](#0x02%20Before%20You%20Ask) 中的规则去检索相关的使用方法

大多数 Browsers 都会提供一个 Search Engine Shortcuts，可以让你通过快捷键就调用对应的 SEO

### 0x03b Internet Archive

因为一些人为的因素，可能会导致一些相关的内容被删除(所以 Web3 才应该是社交网络的最终形式，但是这是一个 political issue)。这时我们就可以使用 [Wayback Machine](https://web.archive.org/) 等 archived/cached 站点查找存档/缓存，这里推荐使用 [GitHub - dessant/web-archives: Browser extension for viewing archived and cached versions of web pages, available for Chrome, Edge and Safari](https://github.com/dessant/web-archives) 这个插件(需要知道完整的 URI，且对应的 URI 已经有被用户或者是系统 archived/cached)

例如 你想要看 [Ryujinx](https://github.com/Ryujinx/Ryujinx) 的历史信息，但是因为 Nintendo 对 maintainers 施压，导致项目直接从 Github 上移除了，这时就可以使用上面这个插件来搜索

### 0x03c Language to Search

互联网不是巴别塔[^3]（也不会是），不同地区的用户使用不同的语言，产生的物料质量也不同。所以选择什么语言检索，也决定了结果的质量

例如 

你想要看中国大陆信创数据库的文档，使用 简中 才是最适合的，虽然是 wrappers of postgresql

然而，如果你想要看 Python 的文档，使用 English 才是最合适的，虽然提供了中文文档

## 0x04 When You Ask

> [!NOTE]
> 当然这也适用于 figure the question out on you own 

在提问题时要做到如下几点

- TODO
- TODO

### Choose a Decent Forum(Section) or Group

在合适的 Forum(Section) 或者 Group 提出你的问题

大多数 Forum 或者 Forum 中的某一个 Section，都会有一个 topic。如果在这个 Forum 或者是 Section 下提出一个 off topic question，这不能说是 rude behavior 但是不礼貌，因为没人希望自己维护的站点，被一些不相关的内容淹没。当然这也不限于 Forum(Section)，也适用于 Groups

例如 Stack Overflow 就是 [Stack Exchange](https://stackexchange.com/sites#) 中的一个 Forum，每一个 Forum 都有自己的 topic

- [Stack Overflow](https://stackoverflow.com/)
	Q&A for professional and enthusiast programmers
- [Server Fault](https://serverfault.com)
	Q&A for system and network administrators
- [Super User](https://superuser.com/)
	Q&A for computer enthusiasts and power users
- [Ask Ubuntu](https://askubuntu.com/)
	Q&A for Ubuntu users and developers

如果你在 Ask Ubuntu 问 Gentoo 相关的问题，是不是很奇怪？

### Write in Clear

> [!tip]
> 这里推荐使用 [Grammarly](https://www.grammarly.com/spell-checker) 来检查

问题要以清楚的形式表达出来，尽量使用 Forum 或者 Group 的 primary language，以 English 为例

- Structure
	一个清楚的格式，具体可以参考 [Issue](Templates/Issue.md)
- Spelling
	单词拼写不能混淆或者误拼，例如混淆了 its 和 it’s，brief 误拼成了 breif
- Punctuate
	标点符号不能乱用，例如在英文中使用了 顿号(、) 替代 逗号(,)
- Capitalize
	大小写，例如使用 mysql 替代 MySQL
- Grammatical
	语法要尽可能的无异议，例如 I've installing lsp plugin 这就是一个反例
- Anit-semi-literate
	不要用一些网络用语替代，例如 u 替代 you，2 替代 to

当然你也可能不是 Forum 或者 Group 使用的 primary language 的 native speaker，这时可以选择加上如下声明之一

- $LANGUAGE is not my native language; please excuse typing errors.
- If you speak $LANGUAGE, please email/PM me; I may need assistance translating my question.
- I am familiar with the technical terms, but some slang expressions and idioms are difficult for me.
- I've posted my question in $LANGUAGE and English. I'll be glad to translate responses, if you only use one or the other.

### Be Precise and Informative

除了这些 [“表象”的 clearly](#Write%20in%20Clear)，还有一些针对问题本身的，要尽可能地提供详细的信息(不同问题需要提供的信息也不同)

例如 你的主机会不定时地重启，那么排查的方向可能有 2 个 —— Softwares and Hardwares。在你无法自己解决的情况下要提供如下信息

1. Softwares
	- the OS and its version you used
	- the drivers and its versions you used
	- the journal logs of last crashed
	- the cron logs of last crashed
1. Hardwares
	- the hardwares you used(eg. CPU/GPU/Memory Bank/Battery)
	- the BIOS version and its settings you used
	- the usage of CPU/GPU/Memory
	- the hardwares's temperature of last crashed

上面这些信息可以通过 `inxi`/`sensors`/`journalctl` 获取。当然除了上面这些信息外，还有一些客观因素

- What were you doing when crashed(What applications opened .etc)
- Uptime of host
- .etc

如果是在网络排障中，那么就需要提供如下信息(5W1H[^4])

- Who - source and destination IP address or target
- What - symptons
- When - timeline,deadline,duration
- Where - location or position
- 
- How


### Describe your problem's symptoms in chronological order

### 0x03a Display the Fact that You have Done These Things

在问题的主体里展示你已经做过了 [Before You Ask](#0x02%20Before%20You%20Ask) 中要求的事


## 0x04 After the Problem has been Solved


- English is not my native language; please excuse typing errors.


## 0x05 Words After All

be poliot如果有正常的对照组，可以设置参照因子(这个就是编程中最朴素的 if ... elif ... else ... 逻辑)。这个做 Experimentation 的过程可以简称为 DTFE(Doing the Fucking 

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***References***

- [How To Ask Questions The Smart Way](http://www.catb.org/~esr/faqs/smart-questions.html)


***FootNotes***

[^1]:[robots.txt - Wikipedia](https://en.wikipedia.org/wiki/Robots.txt)
[^2]:[The Open Source Definition – Open Source Initiative](https://opensource.org/osd)
[^3]:[Tower of Babel - Wikipedia](https://en.wikipedia.org/wiki/Tower_of_Babel)
[^4]:[Project of How → The Kipling method](https://projectofhow.com/methods/the-kipling-method/)
