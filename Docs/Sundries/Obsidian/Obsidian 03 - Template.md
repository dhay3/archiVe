# Obsidian 03 - Template

Obsidian 一个核心的功能就是 Template，可以将文档中 boilerplate 提取出来复用。其中有一些动态的变量

| Variable    | Description                                     |
| ----------- | ----------------------------------------------- |
| `{{title}}` | Title of the active note.                       |
| `{{date}}`  | Today's date. **Default format:** `YYYY-MM-DD`. |
| `{{time}}`  | Current time. **Default format:** `HH:mm`.      |

这些变量可以用在任何地方，不仅仅在 frontmatter 中

```
---
title: "{{title}}"
author: "0x00"
createTime: {{date}} 
lastModifiedTime: {{date}}-{{time}}
draft: true
tags:
  - hashtag1
  - hashtag2
---

**references**


---
*As we enjoy great advantages from the inventions of others, we should be glad of an opportunity to serve others by any invention of ours, and this we should do freely and generously.*
  -- Benjamin Franklin
```

**referneces**

[1]:https://help.obsidian.md/Plugins/Templates
