---
createTime: 2024-09-18 10:21
tags:
  - "#hash1"
  - "#hash2"
---

# W3m 01 - Quick Ref

## 0x01 Preface

>  w3m is a pager/text-based WWW browser. You can browse local documents and/or documents on the WWW using a terminal emulator.

w3m 是一个终端浏览器

## 0x02 Default Key bindings

w3m 默认的 key bindings 大体上 Vim 类似

### 0x02a Motion

- h,j,k,l

	cursor left,down,up,right

- J,K

	roller up/down one line(在 Vim 中使用 Ctrl + E/Y)

- ^,$

	go to the beginning/ending of line

- w,W

	go to next/previous word(在 Vim 中使用 b go to previous word)

- g,G

	go to the first/last line(和 Vim 有略微区别)

- ESC g

	go to specified line(在 Vim 中使用 :)

- TAB

	move to next hyperlink

- Shift TAB

	move to previous hyperlink

- \[,\]

	move to the first/last hyperlink

### 0x02b Hyperlink

- Enter

	follow hyperlink(input)

> [!NOTE]
> 在 xmap keycode 中 Enter 对应 Ctrl j/Ctrl m

- u

	peek link URL

	显示 cursor pointed hyperlink 对应的 URL

- c

	peek current URL

	显示当前页面对应的 URL

- =

	display information about current document

	类似与 `curl -v`

- Ctrl h

	view history of URL

	显示历史记录

- M

	Browse current document using external browser

	使用 external browser 打开当前 URL


### 0x02c File/Stream

- Shfit u

	open URL

	等价于 firefox 中使用 Ctrl l

- V

	View new file

	查看文件，类似 Vim 中的 :tabnew 但是不能对文件做修改

- @/#

	excute shell command and load

	类似于 Vim 中的 :!

### 0x02d Buffer

- B

	Back to the previous buffer

	等价于 go back one page

- v

	view html source

	查看源码

- R

	Reload buffer

	刷新页面


### 0x03e Search

- /,?

	search forward/backward

- n/N

	search next/previous

### 0x03f Miscellany

- H

	help 

- o

	set option

	设置

- q

	quit w3m

- Q

	quit w3m with no confirmation

- Ctrl c

	cancel current operation

## 0x03 Configuration

w3m 默认不会生成配置文件，只有在使用 o(set option) 保存后才会生成配置文件[^2]

## 0x04 Customization Key bindings

w3m 支持通过 `~/.w3m/keymap` 以 `keymap <stroke> <action>` 的格式自定义 key bindings，例如

```
keymap C-o NEXT_PAGE
```

可用的 stroke 可以参考 README.keymap[^3]，所有可用的 action 可以参考 REAMDE.func[^1]。默认配置可以参考 keymap.default[^4]

> [!NOTE]
> w3m 默认没有参数可以 reset keymap，所以如果想 reset keymap 可以使用 `keymap <stroke> NULL`

## 0x05 Image Preview

如果想要使用 kitty image protocol 显示图片，可以参考 Using kittens image protocol[^5]

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***References***

- [w3m manual](https://w3m.sourceforge.net/MANUAL)
- [w3m/doc/README.keymap at master · tats/w3m · GitHub](https://github.com/tats/w3m/blob/master/doc/README.keymap)
- [w3m - ArchWiki](https://wiki.archlinux.org/title/W3m)


***FootNotes***

[^1]:[w3m/doc/README.func at master · tats/w3m · GitHub](https://github.com/tats/w3m/blob/master/doc/README.func)
[^2]:[w3m - ArchWiki](https://wiki.archlinux.org/title/W3m#Configuration)
[^3]:[w3m/doc/README.keymap at master · tats/w3m · GitHub](https://github.com/tats/w3m/blob/master/doc/README.keymap)
[^4]:[w3m/doc/keymap.default at master · tats/w3m · GitHub](https://github.com/tats/w3m/blob/master/doc/keymap.default)
[^5]:[w3m - ArchWiki](https://wiki.archlinux.org/title/W3m#Using_kittens_image_protocol)