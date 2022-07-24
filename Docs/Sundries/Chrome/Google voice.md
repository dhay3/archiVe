# Google voice

ref:

https://www.v2ex.com/t/845214

https://support.google.com/voice#topic=1707989

https://support.google.com/voice/answer/159519?hl=en

## Digest

Google voice 是一个 VoIP 服务，好处是可以使用 US 的虚拟电话号码（这个身份不用说了）来收发短信和电话

## Rules

在注册 Google voice 前需要了解一下他的策略和规则

### 功能

google voice 会根据用户的账户来分配功能

| Users and locations                  |  Google Workspace account (managed)   | Personal account |
| :----------------------------------- | :-----------------------------------: | :--------------: |
| Maximum number of users              |           10 to unlimited*            |        1         |
| Phone number in local country/region | 13 countries/regions* (and expanding) |     US only      |
| International locations              |              Unlimited*               |      **✘**       |

| Features                                               | Google Workspace account (managed) | Personal account |
| :----------------------------------------------------- | :--------------------------------: | :--------------: |
| Forward calls to linked numbers                        |               **✔**                |      **✔**       |
| Voicemail transcripts                                  |               **✔**                |      **✔**       |
| Mobile apps: Android & iPhone or iPad                  |               **✔**                |      **✔**       |
| Don't need to verify account with another phone number |               **✔**                |      **✘**       |
| "Do not disturb" based on Calendar working hours       |               **✔**                |      **✘**       |
| Identify incoming Google Meet calls                    |               **✔**                |      **✘**       |
| Use contacts from corporate directory                  |               **✔**                |      **✘**       |
| Desk phone compatibility                               |               **✔**                |      **✘**       |
| Auto attendants                                        |               **✔**                |      **✘**       |
| Ring groups*                                           |               **✔**                |      **✘**       |
| Add credit to make calls                               |       Billed to organization       |      **✔**       |

可以从上面看到使用 google workspace 账户，能比使用个人账户 功能多出好多。google workspace 怎么玩会另外出一篇文章

### 回收策略

google voice 一段时间内（大概是90天左右）没有使用，你会收到告警信息其中包含 reclaim date（the date the number will be removed）。收到告警信息后可以通过如下任一方式来保号

1. make a call or answewr a call with your voice number
2. send a text message with your voice number
3. listen to your voicemail

## Register

### 方法一

美国IP + 美国实体号码 注册

完全免费

### 方法二

淘宝或者闲鱼直接够买，店铺一般比较隐秘。且由于市场 从 几年前的 10 人民币涨到了 60 人民币

省时省力

### 方法三

美国IP + 使用其他 VoIP 服务提供商(非google voice)注册。提一嘴如果使用 google 的 SOE 搜索出来的大部分都是 google voice 的虚拟账号，且使用互联网的上号码有一定的风险(账号通过手机号找回)。这里推荐使用 Talkatone 另外注册一个虚拟号码，然后绑定google voice

完全免费

## Reclaim Ur voice number

因为 google voice 的策略问题，voice number 在一段时间内不使用就会被回收。这里的回收直接不会注入到 VoIP 池中，默认会有 45 天的静默期，如果超过 45 天才会注入到 VoIP 池中，其他用户才可以强注该账号。

这是可以使用按照 Register 中的 方法三 免费找回回收的账号。如果不会操作可以联系我



## Permanent

由于策略问题，legacy google voice 功能从个人账号中删除了（workspace  账号保留该功能）。仅仅是入口被删除，使用 携替号转网 的功能可以解决

https://support.google.com/voice/answer/1065667?authuser=0#googlexfer

携号转网（port），指的是可以将使用的其他运营商号码转到 google voice 的功能。这样可以同时拥有 2 个 voice number，这样就可以将 voice number make permanent

携号转网需要收取 20$ 

## Calling rates

google voice 的费用比国内的运营商大都要良心。具体可以通过如下连接查询

https://voice.google.com/rates

