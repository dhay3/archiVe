# Obsidian 01 - Overview

## 0x01 Overview


> [!NOTE]
> 现阶段 Obsidian 在文本编辑的体验上还比不了 Typora
>
> 还有一堆 bug 或者是说不人性的地方
>
> 1. 在 list 内的代码块 indent 不能被正确处理，包括创建 code block 以及复制文本到 code block，以及表格
> 2. 在 live preview 中修改表格复杂
> 3. 如果超链接对应的地址或者文档不存在，会自动创建对应的文档，目前不支持修改
>
> 大部分都可以通过 Plugin 来解决，但是也有不能处理的
>
> 所以个人更加倾向于使用 Typora 写文档，Obsidian 来管理文档



Obisidian 是一个强大笔记管理工具，提供了普通 Markdown 编辑器不具备的功能

- Fine-grained control						

  Decide which files and preferences you want to sync to which devic

- Version history								

  Easily track changes between revisions, with one year of version history for every note.

- Collaboration							

  Work with your team on shared files without compromising your private data.

- Plugins
- etc

写完的文档可以直接往 Hugo 发

## 0x02 Vault

Vault 是 Obsidian 中的一个概念，和 Git 的 Wokring Directory 逻辑相同，可以存储笔记/图片/音频

需要说明的一点是，Obsidian 中 vault 的设置是隔离的，在一个 vault 中设置 shortcuts 或者是安装 Plugin，都不会在另外一个 vault 中生效，目前也没有类似 Idea 的 global settings（蛋疼）

如果想要使用其他 vault 中的配置的，就需要将 vault 下的 `.obsidan` 复制到对应的 vault 中



**reference**

[1]:https://forum.obsidian.md/t/global-settings-same-settings-themes-and-plugins-across-multiple-vaults/41789/15
