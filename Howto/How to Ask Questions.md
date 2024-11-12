---
createTime: 2024-08-14 09:09
license: cc by 4.0
tags:
  - "#hash1"
  - "#hash2"
---

# How to Ask Questions

## 0x01 Preface

> [!important]
> 本文在 [How To Ask Questions The Smart Way](http://www.catb.org/~esr/faqs/smart-questions.html) 的基础上结合自己的理解和实践有感而发

同样的一个问题，表达的方式不同，得到的回复也不同(当然你也可能得不到回复 XD)。如何问好一个问题，才可以得到你满意的回答，这非常重要

一个事物的生命周期普遍包括

1. Begin
2. Process
3. End

相对应得一个问题也有自己的生命周期

1. Before You Ask
2. When You Ask
3. After the Problem has been Solved

## 0x02 Before You Ask

> [!note] 
> 各部分的结果可能存在交集，所以逻辑上没有先后顺序，但是从结果的权威性角度出发
> $DTFE > RTFS > JTFG > RTFM > STFW > ATFL$

在提问前需要做到如下几点(至少要证明自己对问题做过解决的尝试，不要让别人觉得你只是一个 taker)

- STFW
	Try to Find an Answer by Searching the Web
- RTFW
	Try to Find an Answer by Reading the Manual
- ATFL
	Try to Find an Answer by Asking the LLM
- DTFW
	Try to Find an Answer by Doing Experimentation
- RTFS
	Try to Find an Answer by Reading the Source Code
- JTFG
	Try to Find an Answer by Joining Groups

### 0x02a Try to Find an Answer by Searching the Web(STFW)

利用 SEO 搜索你想要的答案

不仅限于各种 SEO(掌握 SEO advanced search 以及 operator search 的使用是必要的，例如 [Google Searches](../Notes/Sundries/Chrome/Google%20Searches.md)，也包括一些设置了 robots.txt[^1](按照 RFC 9309 的规定，SEO 需要按照 robots.txt 中的规则爬取信息，但是不是强制要求) 的 forum 或者 site(例如 Google scholar/libgen/etc.)。这个利用 SEO 搜索的过程可以简称为 STFW(Searching the Fucking Web)

例如 你想要搜索和 KDE system monitor sensors bug 相关的内容，你可以使用 Google/DuckDuckgo/Brave/etc. SEO 直接搜索关键字 `KDE system monitor sensors bug` 也可以使用 operator search 搜索 `intitle:"system monitor sensors" intext:bug` 对结果过滤

如果没有得到想要的内容，可以到 [KDiscuss](https://discuss.kde.org/) forum 搜索，也可以到 [KDE Bugtracking System](https://bugs.kde.org/) 搜索(相对更精确，前提是你知道有对应的站点)

### 0x02b Try to Find an Answer by Reading the Manual(RTFM)

使用 User Manual 搜索你想要的答案

和家电一样，大多数软件都会有 User Manual(也可能叫 Get Help/Documentation/Support/etc.)，会告诉用户该如何使用软件以及 troubleshooting。这个使用 User Manual 搜索的过程可以简称为 RTFM(Reading the Fucking Manual)

例如 你想要使用 Thunderbird 接管 Gmail，就可以在 [Thunderbird Support](https://support.mozilla.org/en-US/products/thunderbird) 页面搜索关键字 gmail

### 0x02c Try to Find an Answer by Asking the LLM(ATFL)

使用 LLM 得到你想要的答案

> [!note]
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

做实验是最原始也是最有效的方法，但是要控制变量。如果有正常的对照组，可以设置参照因子(这个就是编程中最朴素的 `if ... elif ... else ...` 逻辑)。这个做 Experimentation 的过程可以简称为 DTFE(Doing the Fucking Experimentation)

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

因为一些人为的因素，可能会导致一些相关的内容被删除(所以 Web3 才应该是社交网络的最终形式，但这是一个 political issue)。这时我们就可以使用 [Wayback Machine](https://web.archive.org/) 等 archived/cached 站点查找存档/缓存，这里推荐使用 [GitHub - dessant/web-archives: Browser extension for viewing archived and cached versions of web pages, available for Chrome, Edge and Safari](https://github.com/dessant/web-archives) 这个插件(需要知道完整的 URI，且对应的 URI 已经有被用户或者是系统 archived/cached)

例如 你想要看 [Ryujinx](https://github.com/Ryujinx/Ryujinx) 的历史信息，但是因为 Nintendo 对 maintainers 施压，导致项目直接从 Github 上移除了，这时就可以使用上面这个插件来搜索

### 0x03c Language to Search

互联网不是巴别塔[^3]（也不会是），不同地区的用户使用不同的语言，产生的物料质量也不同。所以选择什么语言检索，也决定了结果的质量

例如 

你想要看中国大陆信创数据库（wrappers of postgresql）的文档，使用 简中 才是最适合的

然而，如果你想要看 Qt6 的文档，使用 English 才是最合适的

## 0x04 When You Ask

> [!note]
> 当然这也适用于 figure the question on you own 

在提问题时要做到如下几点

- Choose a Community
- Read Community Rules
- a Summarized Title
- Write in Clear
- Be Precise and Informative
- How to Reproduce
- Display that You have Done Before You Ask
- Proofread Before Posting

### 0x04a Choose a Community

在合适的 Community(社区可以是 Forum/Section of Forum/Groups) 提出你的问题

大多数 Community，都会有一个 topic。如果在这个 Community 提出一个 off topic question，这不能说是 rude behavior 但是不礼貌，因为没人希望自己维护的社区，被一些不相关的内容淹没

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

### 0x04b Read Community Rules

每一个 Community 都会有自己的规则（也可能叫作 FAQ），在你提问前请先阅读 Community Rules。这是为了社区不被没有营养的内容”泛洪“，同时也是保护你自己的账号不被封禁

例如 Do 的

[FAQ - LINUX DO](https://linux.do/faq)

### 0x04b a Summarized Title 

一个精简明了的标题可以说是最重要的一部分(the first thing that potential answerers will see)。我们应该只描述发生了什么，弱化对象、空间和时间，即 5W1H[^4] 中 What

例如

[An option "Runs in web side panel" for extensions · zen-browser/desktop · Discussion #2685 · GitHub](https://github.com/zen-browser/desktop/discussions/2685)

### 0x04c Write in Clear

> [!tip]
> 主要针对内容，这里推荐使用 [Grammarly](https://www.grammarly.com/spell-checker) 来检查

问题要以清楚的形式表达出来，尽量使用社区的 primary language，以 English 为例

- Structure
	一个清楚的格式，具体可以参考 [TEM_OF_ISSUE](../Templates/TEM_OF_ISSUE.md)
- Spelling
	单词拼写不能混淆或者误拼，例如混淆了 its 和 it's，brief 误拼成了 breif
- Punctuate
	标点符号不能乱用，例如在英文中使用了 顿号(、) 替代 逗号(,)
- Capitalize
	大小写，例如使用 mysql 替代 MySQL
- Grammatical
	语法要尽可能的无异议，例如 I've installing lsp plugin 这就是一个反例
- Anit-semi-literate
	不要用一些网络用语替代，例如 u 替代 you，2 替代 to

一个清楚的格式，应该有包括如下几个原则

- describe your problem's symptoms in chronological order
	以时间顺序描述问题的症状
- volume is not precision
	内容越多并不一定越好，尽量缩减做到“小巧精悍”

当然你也可能不是 primary language native speaker，这时可以选择加上如下声明之一

- $LANGUAGE is not my native language; please excuse typing errors.
- If you speak $LANGUAGE, please email/PM me; I may need assistance translating my question.
- I am familiar with the technical terms, but some slang expressions and idioms are difficult for me.
- I've posted my question in $LANGUAGE and English. I'll be glad to translate responses, if you only use one or the other.

### 0x04d Be Precise and Informative

除了这些 [“表象”的 clearly](#0x04c%20Write%20in%20Clear)，还有一些针对问题本身的，要尽可能地提供详细的信息(不同问题需要提供的信息也不同)

例如 你的主机会不定时地重启，那么排查的方向可能有 2 个 —— Softwares and Hardwares。在你无法自己解决的情况下要提供如下信息

1. Softwares
	- the OS and its version you used
	- the drivers and its versions you used
	- the journal logs of last crashed
	- the cron logs of last crashed
2. Hardwares
	- the hardwares you used(eg. CPU/GPU/Memory Bank/Battery)
	- the BIOS version and its settings you used
	- the usage of CPU/GPU/Memory
	- the hardwares's temperature of last crashed

上面这些信息可以通过 `inxi`/`sensors`/`journalctl` 获取。当然除了上面这些信息外，还有一些客观因素

- What were you doing when crashed(What applications have opened .etc)
- Uptime of host
- .etc

如果是在网络排障中，那么就需要提供如下信息(5W1H[^4])

- Who - source and destination IP address or target
- What - symptons
- Where - location or position
- When - timeline,deadline,duration
- Why - possible reason
- How - how comes

### 0x04e How to Reproduce

如果问题是可以被复现的，要提供复现方式

例如 How to Reproduce 部分

[\`kitten icat\`  freezes  the keyboard input when \`--hold\` option is presented in TMUX runs on Kitty · kovidgoyal/kitty · Discussion #7793 · GitHub](https://github.com/kovidgoyal/kitty/discussions/7793)

### 0x04f Display that You have Done Before You Ask

在问题的主体中尽可能体现你已经做过了 [Before You Ask](#0x02%20Before%20You%20Ask) 中要求的事

例如 [Scaling do not work correctly when taking screenshot in Plasma6/Wayland · Issue #3614 · flameshot-org/flameshot · GitHub](https://github.com/flameshot-org/flameshot/issues/3614) 中 Things I have tried 所展示的(这不是一个很好的例子，没有体现 STFW/RTFM/ATFL/RTFS/JTFG)

### 0x04g Proofread Before Posting

在 Post 前对内容做校对也非常重要。如果你的问题 edited over and over again 就说明你没有考虑清楚自己的问题

## 0x05 After the Problem has been Solved

TCP/IP 中在发出 SYN 请求报文，你会收到 SYN-ACK 做为回应报文，同时发送 ACK 表示收到回应

提问也一样，下面这些规则就是第三次握手的 ACK

- Have 'FIXED' 'SOLVED' Tag in the Subject Line
- Follow up with a brief note on the solution

### 0x05a Have 'FIXED' 'SOLVED' Tag in the Subject Line

在标题添加 `[FIXED]` 或者是 `[SOLVED]` 标签，例如

[\[SOLVED\] Running guest machines show as powered off in Virtualbox - virtualbox.org](https://forums.virtualbox.org/viewtopic.php?t=111131)

> [!important]
> 要根据社区的规则，例如 Stack Exchange 就要求用户不要在 Subject 中添加 FIXED 或者是 SOLVED 等字样

### 0x05b Follow up with a brief note on the solution

在 main thread 中添加一段 solution breif 告诉其他人你的解决方式，例如

[zsh - How to start tmux as the default when terminals opened exclude the terminal in Dolphin - Super User](https://superuser.com/questions/1832872/how-to-start-tmux-as-the-default-when-terminals-opened-exclude-the-terminal-in-d/1832883#1832883)

## 0x06 Words After All

上面都是规则类的准则，下面这些是行为类的准则

### 0x06a Courtesy never hurts

在不了解对方的情况下，人们总是愿意选择和看上去谈吐谦虚礼貌的人交谈。这也适用于网络社交，你可以使用文字展现你的 Courtesy

例如 你可以在问题结尾添加类似

- "Any hints would be highly appreciated."
- "Thanks in advance."

又譬如 在阐述自己的推测时添加类似

- "Please tell me if I've missed somthing"

### 0x06b If You Can't Get An Answer

当然你也有可能会得不到任何有用的答复，这时候请不要：

- 觉得没有人能帮助你
	可能你的问题只是一个动动手就可以得到答案的
- 在同一个社区重复 repost
	repost 通常是无意义的，如果自己不能解决那么请耐心

## 0x07 Epilog

没有想好写什么

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See Also***

- [How To Ask Questions The Smart Way](http://www.catb.org/~esr/faqs/smart-questions.html)
- [The 5W1H Method: Elements & Example | SafetyCulture](https://safetyculture.com/topics/5w1h/)
- [How do I ask a good question? - Help Center - Stack Overflow](https://stackoverflow.com/help/how-to-ask)
- [How do I write a good answer? - Help Center - Stack Overflow](https://stackoverflow.com/help/how-to-answer)

***References***

[^1]:[robots.txt - Wikipedia](https://en.wikipedia.org/wiki/Robots.txt)
[^2]:[The Open Source Definition – Open Source Initiative](https://opensource.org/osd)
[^3]:[Tower of Babel - Wikipedia](https://en.wikipedia.org/wiki/Tower_of_Babel)
[^4]:[Project of How → The Kipling method](https://projectofhow.com/methods/the-kipling-method/)