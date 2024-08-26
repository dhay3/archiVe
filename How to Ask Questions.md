---
author: "0x00"
createTime: 2024-08-14 09:09
tags:
  - "#hash1"
  - "#hash2"
---

# How to Ask Questions

## 0x01 Preface

同样的一个问题，表达的方式不同，得到的回复也不同(你也可能得不到回复 XD)。如何问好一个问题，才可以得到你满意的回答，这非常重要

本文在 [How To Ask Questions The Smart Way](http://www.catb.org/~esr/faqs/smart-questions.html) (a little bit old school)的基础上按照时间逻辑划分如下几部分

1. Before You Ask
2. When You Ask
3. After the Problem has been Solved

## 0x02 Before You Ask

> [!important] 
> 各部分的结果可能存在交集，所以逻辑上没有先后顺序，但是从权威性的角度出发
> $DTFE > RTFS > JTFG > RTFM > STFW > ATFL$

在提问前需要做到如下几点

#### 0x02a Try to Find an Answer by Searching the Web(STFW)

利用网络，不仅限于各种 SEO(必须掌握 SEO advanced search 以及 operator search 的使用)，也包括一些设置了 robots.txt[^1](按照 RFC 9309 的规定，SEO 需要按照 robots.txt 中的规则爬取信息，但是不是强制要求) 的 forum 或者 site(例如 Google scholar/libgen/etc.)。可以直接简称为 STFW(Searching the Fucking Web)

例如 你想要搜索和 KDE system monitor sensors bug 相关的内容，你可以使用 Google/DuckDuckgo/Brave/etc. SEO 直接搜索关键字 `KDE system monitor sensors bug` 也可以使用 operator search 搜索 `intitle:"system monitor sensors" intext:bug` 对结果过滤

如果没有得到想要的内容，可以到 [KDiscuss](https://discuss.kde.org/) forum 搜索，也可以到 [KDE Bugtracking System](https://bugs.kde.org/) 搜索(相对更精确，前提是你知道有对应的站点)

#### 0x02b Try to Find an Answer by Reading the Manual(RTFM)

大多数软件都会有 User Manual(也可能叫 Get Help/Documentation/Support/etc.)，会告诉用户该如何使用软件以及 troubleshooting。可以直接简称为 RTFM(Reading the Fucking Manual)

例如 你想要使用 Thunderbird 接管 Gmail，就可以在 [Thunderbird Support](https://support.mozilla.org/en-US/products/thunderbird) 页面搜索关键字 gmail

#### 0x02c Try to Find an Answer by Asking the LLM(ATFL)

LLM 是一个很好的辅助工具(回答并不一定可靠，所以作为一个辅助工具)，你可以设置场景，让 AI 分步骤回答你的问题。可以直接简称为 ATFL(Asking the Fucking LLM)

例如 你想知道 GPLv1/2/3 之间的区别，你可以使用 Claude/ChatGPT/Gemini/Perlplexity/etc. 先对 LLM 设置角色

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

#### 0x02d Try to Find an Answer by Doing Experimentation(DTFE)

做实验是最原始也是最有效的方法，但是要控制变量，如果有正常的对照组，可以设置参照因子。可以简称为 DTFE(Doing the Fucking Experimentation)

例如 你现在不能通过浏览器访问你家的软路由了(不考虑 DNS)，你要先确认是 L7 应用有问题，或者是 L4 传输有问题，或者是 L3 网络有问题 还是 L2/L1 设备有问题。这里 Layer 就是变量

如果 设备电口正常亮灯，可以概率排除 L2/L1 设备问题，反之 就排查物理设备或者配件(例如 软路由网口/网线/交换机/光猫/etc.)
如果 可以从 LAN 内主机 ping 通软路由，可以概率排除 L3 网络问题，反之 就排查软路由系统 IP 地址以及路由
如果 可以从 LAN 内主机 netcat 软路由 80 端口，可以概率排除 L4 传输问题，反之 就排查软路由系统是否有应用监听 socket
那么 最后只可能是 L7 应用问题，重启 luci 即可

#### 0x02e Try to Find an Answer by Reading the Source Code(RTFS)

大多数开源的软件都会公开代码，所以一些奇怪的 Bug 或者是非预期的现象可以通过读源码解决。可以简称为 RTFS(Reading the Fucking Source Code)

例如 Virtualbox Guest Machine 不能通过 Vagrant 启动，但是可以通过 Virtualbox GUI 启动

[\[SOLVED\] Can't boot virtual machines via vagrant but works well with virtualbox GUI(vagrant 2.4.0/virtualbox 7.0.12) · Issue #13288 · hashicorp/vagrant · GitHub](https://github.com/hashicorp/vagrant/issues/13288)

#### 0x02f Try to Find an Answer by Joining Groups(JTFG)

软件的站点不一定会及时更新，但是会有一些社区组群(Discord/Telegram/etc.)，管理者会在这些组群中实时发布一些 声明 或者是 Pinned Messages 会提供非常有的信息

例如 在 Clash Verge Rev 更新后，Global Extend Config 中配置如下额外全局路由规则，发现不生效(之前生效)

```
prepend-rules:
	- DOMAIN-SUFFIX,4lice.com,DIRECT
	- AND,((IP-CIDR,23.94.117.188/32),(DST-PORT,65522)),DIRECT
```

这时可以加入站点提供的 telegram 组群，Pinned Messages 说明了 v1.6.2 后 Merge 配置(Global Extend Config)不再支持 `prepend-rule-providers`、`prepend-proxy-providers`、`append-rule-providers`、`append-proxy-providers`，要通过 edit profile rule 或者是 proxies/proxy groups 实现更加细粒的配置

## 0x03 When You Ask

在提问题时要做到如下几点

### 0x03a Write in Clear Language

问题要以清楚的语言格式表达出来

1. Spell

   单词拼写不能混淆或者误拼，例如混淆了 its 和 it’s，brief 误拼成了 breif

2. Punctuate

   标点符号不能乱用，例如在英文中使用了 顿号(、)

3. Capitalize

   大小写



### 0x03a Display the Fact that You have Done These Things

在问题的主体里展示你已经做过了 [Before You Ask](#0x02%20Before%20You%20Ask) 中要求的事


## 0x04 After the Problem has been Solved


- English is not my native language; please excuse typing errors.


---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***References***

- [How To Ask Questions The Smart Way](http://www.catb.org/~esr/faqs/smart-questions.html)


***FootNotes***

[^1]:[robots.txt - Wikipedia](https://en.wikipedia.org/wiki/Robots.txt)
